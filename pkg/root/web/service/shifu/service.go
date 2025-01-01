package shifu

import (
	"context"

	"github.com/rhoninl/sft/pkg/root/install"
	pb "github.com/rhoninl/sft/pkg/root/web/proto/shifu"
	"github.com/rhoninl/sft/pkg/utils/shifu"
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
