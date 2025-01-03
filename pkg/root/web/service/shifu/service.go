package shifu

import (
	"context"

	"github.com/rhoninl/sft/pkg/root/install"
	"github.com/rhoninl/sft/pkg/root/uninstall"
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
