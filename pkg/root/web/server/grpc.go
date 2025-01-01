package server

import (
	"fmt"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/rhoninl/sft/pkg/root/web/proto/shifu"
	"github.com/rhoninl/sft/pkg/root/web/service/shifu"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(port int) error {
	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register services
	shifuServer := shifu.NewShifuServer()
	pb.RegisterShifuServiceServer(grpcServer, shifuServer)
	reflection.Register(grpcServer)

	// Wrap gRPC server with gRPC-Web
	wrappedGrpc := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true // Allow all origins
		}),
		grpcweb.WithAllowedRequestHeaders([]string{"*"}),
	)

	// Create HTTP server
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			// Add CORS headers
			resp.Header().Set("Access-Control-Allow-Origin", "*")
			resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")

			// Handle preflight requests
			if req.Method == "OPTIONS" {
				resp.WriteHeader(http.StatusOK)
				return
			}

			// Handle gRPC-Web requests
			if wrappedGrpc.IsGrpcWebRequest(req) {
				wrappedGrpc.ServeHTTP(resp, req)
				return
			}

			// Handle normal HTTP requests
			http.Error(resp, "Unsupported request", http.StatusNotImplemented)
		}),
	}

	fmt.Printf("Starting gRPC-Web server on port %d\n", port)
	if err := httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
