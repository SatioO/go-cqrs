package framework

import (
	"context"

	"github.com/satioO/scheduler/scheduler/internal/adapters/framework/grpc/pb"
)

func (grpca *GRPCServer) GetHealthInfo(ctx context.Context, req *pb.HealthRequest) (*pb.HealthReply, error) {
	err := grpca.app.Commands.OpenAccount.Handle()

	if err != nil {
		return nil, err
	}

	return &pb.HealthReply{
		Message: "Command Executed",
	}, nil
}
