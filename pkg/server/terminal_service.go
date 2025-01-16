package server

import (
	"context"

	pb "github.com/edgenesis/shifu/pkg/proto/shifu"
	"github.com/edgenesis/shifu/pkg/terminal"
)

// ... 其他现有的代码 ...

func (s *shifuServer) GetCompletions(ctx context.Context, req *pb.CompletionRequest) (*pb.CompletionResponse, error) {
	completions := terminal.GetCompletions(req.GetPartial(), req.GetCurrentDir())

	return &pb.CompletionResponse{
		Completions: completions,
	}, nil
}
