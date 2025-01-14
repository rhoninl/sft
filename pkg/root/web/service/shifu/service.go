package shifu

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/root/devices"
	"github.com/rhoninl/sft/pkg/root/forward"
	"github.com/rhoninl/sft/pkg/root/install"
	"github.com/rhoninl/sft/pkg/root/restart"
	"github.com/rhoninl/sft/pkg/root/uninstall"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	pb "github.com/rhoninl/sft/proto/shifu"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShifuServer struct {
	pb.UnimplementedShifuServiceServer
}

func NewShifuServer() *ShifuServer {
	return &ShifuServer{}
}

func (s *ShifuServer) CheckInstallation(ctx context.Context, req *pb.CheckInstallationRequest) (*pb.CheckInstallationResponse, error) {
	installed := true
	if err := shifu.CheckShifuInstalled(); err != nil {
		installed = false
	}

	return &pb.CheckInstallationResponse{
		Installed: installed,
	}, nil
}

func (s *ShifuServer) InstallShifu(ctx context.Context, req *pb.InstallShifuRequest) (*pb.InstallShifuResponse, error) {
	version := req.GetVersion()
	if version == "" {
		version = install.EmptyVersion
	}

	if err := install.InstallShifu(version); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to install shifu: %v", err)
	}

	return &pb.InstallShifuResponse{
		Success: true,
	}, nil
}

func (s *ShifuServer) GetAllAvailableVersions(ctx context.Context, req *pb.GetAllAvailableVersionsRequest) (*pb.GetAllAvailableVersionsResponse, error) {
	versions := shifu.GetAllAvailableVersions()
	return &pb.GetAllAvailableVersionsResponse{
		Versions: versions,
	}, nil
}

func (s *ShifuServer) UninstallShifu(ctx context.Context, req *pb.UninstallShifuRequest) (*pb.UninstallShifuResponse, error) {
	if err := uninstall.UninstallShifu(); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to uninstall shifu: %v", err)
	}

	return &pb.UninstallShifuResponse{
		Success: true,
	}, nil
}

func (s *ShifuServer) ListDevices(ctx context.Context, req *pb.ListDevicesRequest) (*pb.ListDevicesResponse, error) {
	devices, err := devices.ListDevices()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list devices: %v", err)
	}

	protoDevices := make([]*pb.Device, 0)
	for _, device := range devices {
		protoDevices = append(protoDevices, &pb.Device{
			Name:     device.Name,
			Protocol: device.Protocol,
			Address:  device.Address,
			Status:   device.Status,
			Age:      device.Age,
		})
	}

	return &pb.ListDevicesResponse{
		Devices: protoDevices,
	}, nil
}

func (s *ShifuServer) GetDeviceDetails(ctx context.Context, req *pb.GetDeviceDetailsRequest) (*pb.GetDeviceDetailsResponse, error) {
	device, err := k8s.GetAllByDeviceName(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get device details: %v", err)
	}

	var resp pb.GetDeviceDetailsResponse
	var edgedevice pb.Edgedevice
	if device.EdgeDevice.Status.EdgeDevicePhase != nil {
		edgedevice.Status = string(*device.EdgeDevice.Status.EdgeDevicePhase)
	}
	edgedevice.Sku = *device.EdgeDevice.Spec.Sku
	edgedevice.Protocol = string(*device.EdgeDevice.Spec.Protocol)
	edgedevice.Address = *device.EdgeDevice.Spec.Address
	edgedevice.Age = devices.TimeToAge(device.EdgeDevice.CreationTimestamp.Time)
	data, err := json.Marshal(device.EdgeDevice.Spec.ProtocolSettings)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal setting: %v", err)
	}
	edgedevice.Setting = string(data)
	data, err = json.Marshal(device.EdgeDevice.Spec.GatewaySettings)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal gateway: %v", err)
	}
	edgedevice.Gateway = string(data)

	apis := device.ConfigMap.Data["instructions"]
	resp.APIs = apis

	resp.Edgedevice = &edgedevice
	return &resp, nil
}

func (s *ShifuServer) ForwardPort(req *pb.ForwardPortRequest, stream pb.ShifuService_ForwardPortServer) error {
	readyChan := make(chan struct{})

	if err := forward.ForwardPort(stream.Context(), req.GetDeviceName(), req.GetDevicePort(), req.GetLocalPort(), readyChan); err != nil {
		if err := stream.Send(&pb.ForwardPortResponse{
			Success: false,
		}); err != nil {
			return err
		}
		logger.Printf("GRPC: Error: Failed to forward port: %v", err)
		return err
	}

	select {
	case <-readyChan:
		if err := stream.Send(&pb.ForwardPortResponse{
			Success: true,
		}); err != nil {
			return err
		}
	case <-stream.Context().Done():
		return stream.Context().Err()
	}

	<-stream.Context().Done()

	return nil
}

func (s *ShifuServer) RestartDeviceShifu(ctx context.Context, req *pb.RestartDeviceShifuRequest) (*pb.Empty, error) {
	if err := restart.RestartDeviceShifu(req.DeviceName); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to restart device shifu: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s *ShifuServer) DeleteDeviceShifu(ctx context.Context, req *pb.DeleteDeviceShifuRequest) (*pb.Empty, error) {
	return nil, errors.New("not implemented")
}

func (s *ShifuServer) GetAllContainerName(ctx context.Context, req *pb.GetAllContainerNameRequest) (*pb.GetAllContainerNameResponse, error) {
	// deploys, err := k8s.GetDeployByEnv(req.GetDeviceName())
	device, err := k8s.GetAllByDeviceName(req.GetDeviceName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all container name: %v", err)
	}

	pods, err := k8s.GetDeploymentPods("deviceshifu", device.Deployment.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all container name: %v", err)
	}

	if len(pods) == 0 {
		return nil, status.Errorf(codes.NotFound, "no pods found")
	}

	pod := pods[0]

	containerNames := make([]string, 0)
	for _, container := range pod.Spec.Containers {
		containerNames = append(containerNames, container.Name)
	}

	return &pb.GetAllContainerNameResponse{
		ContainerNames: containerNames,
	}, nil
}

func (s *ShifuServer) GetDeviceShifuLogs(req *pb.GetDeviceShifuLogsRequest, stream pb.ShifuService_GetDeviceShifuLogsServer) error {
	r, w := io.Pipe()
	device, err := k8s.GetAllByDeviceName(req.GetDeviceName())
	if err != nil {
		return err
	}

	go func() {
		defer w.Close()
		err := k8s.GetDeploymentLogs("deviceshifu", device.Deployment.Name, req.ContainerName, true, w)
		if err != nil {
			w.CloseWithError(err)
		}
	}()

	buffer := make([]byte, 4096)
	for {
		n, err := r.Read(buffer)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.GetDeviceShifuLogsResponse{
			Log: string(buffer[:n]),
		}); err != nil {
			return err
		}
	}
}
