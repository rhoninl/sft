package server

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rhoninl/sft/pkg/root/web/service/shifu"
	pb "github.com/rhoninl/sft/proto/shifu"
	assets "github.com/rhoninl/sft/webview"
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

	// Setup static file server
	contentStatic, err := fs.Sub(assets.Content, "build")
	if err != nil {
		return fmt.Errorf("failed to setup static file server: %v", err)
	}
	fileServer := http.FileServer(http.FS(contentStatic))

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

			// Handle static files and SPA routing
			if strings.Contains(req.Header.Get("Accept"), "text/html") {
				// Serve index.html for all HTML requests (SPA routing)
				req.URL.Path = "/"
			}
			fileServer.ServeHTTP(resp, req)
		}),
	}

	fmt.Printf("Starting server on port %d\n", port)
	if err := httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
