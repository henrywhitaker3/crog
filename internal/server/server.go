package server

import (
	"net"

	"github.com/henrywhitaker3/crog/internal/config"
	"github.com/henrywhitaker3/crog/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedCrogServer

	cfg        *config.Config
	listener   *net.Listener
	grpcServer *grpc.Server
}

func New(cfg *config.Config) (*Server, error) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	serv := Server{
		cfg:        cfg,
		grpcServer: grpcServer,
	}

	pb.RegisterCrogServer(grpcServer, serv)

	return &serv, nil
}

func (s *Server) Listen() error {
	lis, err := net.Listen("tcp", s.cfg.Server.Listen)
	if err != nil {
		return err
	}
	s.listener = &lis

	if err := s.grpcServer.Serve(*s.listener); err != nil {
		return err
	}

	return nil
}

func (s *Server) Close() {
	s.grpcServer.GracefulStop()
	lis := *s.listener
	lis.Close()
}
