package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// A heavy-dependency service: the AWS SDK and gRPC are linked in, as in a real file service with an internal API.
func main() {
	client := s3.New(s3.Options{Region: "auto", Credentials: aws.AnonymousCredentials{}})
	grpcServer := grpc.NewServer()
	healthpb.RegisterHealthServer(grpcServer, health.NewServer())
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"service": "files", "s3_region": client.Options().Region, "grpc_services": len(grpcServer.GetServiceInfo())})
	})
	_ = context.Background()
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
