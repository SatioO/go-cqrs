package framework

import (
	"net"
	"os"

	"github.com/satioO/scheduler/scheduler/internal/adapters/framework/grpc/pb"
	"github.com/satioO/scheduler/scheduler/internal/application"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	app application.Application
}

func NewGRPCServer(app application.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func Run(adapter *GRPCServer) {
	lstn, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		logrus.Fatalf("Failed to listen on PORT", os.Getenv("PORT"))
	}

	server := grpc.NewServer()
	pb.RegisterHealthServiceServer(server, adapter)

	logrus.Println("Starting GRPC server: " + os.Getenv("PORT"))
	if err := server.Serve(lstn); err != nil {
		logrus.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
