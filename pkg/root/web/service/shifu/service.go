package shifu

import (
	"bufio"
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/root/devices"
	"github.com/rhoninl/sft/pkg/root/forward"
	"github.com/rhoninl/sft/pkg/root/install"
	"github.com/rhoninl/sft/pkg/root/restart"
	"github.com/rhoninl/sft/pkg/root/uninstall"
	"github.com/rhoninl/sft/pkg/terminal"
	"github.com/rhoninl/sft/pkg/utils/jsonhelper"
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
	data, err := jsonhelper.MarshalAll(device.EdgeDevice.Spec.ProtocolSettings)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal setting: %v", err)
	}
	edgedevice.Setting = string(data)
	data, err = jsonhelper.MarshalAll(device.EdgeDevice.Spec.GatewaySettings)
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

func (s *ShifuServer) ExecuteCommand(req *pb.CommandRequest, stream pb.ShifuService_ExecuteCommandServer) error {
	ctx := stream.Context()

	// Extract working directory from command
	var workDir string
	if strings.HasPrefix(req.Command, "cd ") {
		// Special handling for cd command
		return handleCdCommand(req.Command, stream)
	}

	// Create command
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(ctx, "cmd", "/c", req.Command)
	} else {
		cmd = exec.CommandContext(ctx, "sh", "-c", req.Command)
	}

	// Set working directory
	if workDir != "" {
		cmd.Dir = workDir
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	// Use WaitGroup to ensure all output is processed
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		streamOutput(stdout, stream, false)
	}()

	go func() {
		defer wg.Done()
		streamOutput(stderr, stream, true)
	}()

	// Wait for command completion and all output processing
	err = cmd.Wait()
	wg.Wait()
	return err
}

func handleCdCommand(command string, stream pb.ShifuService_ExecuteCommandServer) error {
	dir := strings.TrimSpace(strings.TrimPrefix(command, "cd "))
	if dir == "" {
		// If no directory specified, use HOME directory
		dir = os.Getenv("HOME")
		if runtime.GOOS == "windows" {
			dir = os.Getenv("USERPROFILE")
		}
	}

	// Expand path
	expandedDir, err := filepath.Abs(dir)
	if err != nil {
		stream.Send(&pb.CommandResponse{
			Output:  err.Error(),
			IsError: true,
		})
		return err
	}

	// Check if directory exists
	if _, err := os.Stat(expandedDir); err != nil {
		stream.Send(&pb.CommandResponse{
			Output:  err.Error(),
			IsError: true,
		})
		return err
	}

	// Change working directory
	if err := os.Chdir(expandedDir); err != nil {
		stream.Send(&pb.CommandResponse{
			Output:  err.Error(),
			IsError: true,
		})
		return err
	}

	// Send new working directory
	stream.Send(&pb.CommandResponse{
		Output:  expandedDir,
		IsError: false,
	})

	return nil
}

func streamOutput(reader io.Reader, stream pb.ShifuService_ExecuteCommandServer, isError bool) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if err := stream.Send(&pb.CommandResponse{
			Output:  text,
			IsError: isError,
		}); err != nil {
			return
		}
	}
}

func (s *ShifuServer) GetCompletions(ctx context.Context, req *pb.CompletionRequest) (*pb.CompletionResponse, error) {
	// Get current working directory
	currentDir := req.GetCurrentDir()
	if currentDir == "" {
		currentDir = "/"
	}

	// Get partial input
	partial := req.GetPartial()

	// Use terminal package to get completion suggestions
	completions := terminal.GetCompletions(partial, currentDir)

	return &pb.CompletionResponse{
		Completions: completions,
	}, nil
}
