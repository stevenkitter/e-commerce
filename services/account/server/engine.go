package server

import (
	"github.com/stevenkitter/skills/database"
	pb "github.com/stevenkitter/skills/services/account/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Engine struct {
	Name string
	Port string
	DB   *database.Database
}

type Server interface {
	Run() error
}

func (en *Engine) Run() error {
	tcp, err := net.Listen("tcp", en.Port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterAccountServer(s, en)
	reflection.Register(s)
	if err := s.Serve(tcp); err != nil {
		return err
	}
	return nil
}
