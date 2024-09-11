package grpc

import (
	"context"
	"fmt"
	"github.com/fmo/hexagonal-blog/config"
	"github.com/fmo/hexagonal-blog/golang/post"
	"github.com/fmo/hexagonal-blog/internal/ports"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Adapter struct {
	api    ports.APIPorts
	port   int
	server *grpc.Server
	post.UnimplementedPostServer
}

func NewAdapter(api ports.APIPorts, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run(ctx context.Context) {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()

	a.server = grpcServer

	post.RegisterPostServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
