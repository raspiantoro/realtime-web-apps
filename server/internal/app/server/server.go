package server

import (
	"context"
	"log"
	"net"

	"github.com/raspiantoro/realtime-web-apps/server/internal/app/appcontext"
	trafficPB "github.com/raspiantoro/realtime-web-apps/server/pkg/traffic/v1"
	"google.golang.org/grpc"
)

//Server define struct for server instance
type Server struct{}

//NewServer get new server instance
func NewServer() Server {
	return Server{}
}

//RunServer run server for grpc service
func (s *Server) RunServer(ctx context.Context, service appcontext.Service, host string, port string) (err error) {
	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	trafficPB.RegisterTrafficServer(server, &service.Traffic)

	go func() {
		select {
		case <-ctx.Done():
			server.GracefulStop()
			log.Println("gRPC server has been shutdown")
		}
	}()

	log.Printf("gRPC serve at %s:%s", host, port)
	return server.Serve(listen)
}
