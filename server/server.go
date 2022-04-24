package server

import (
	"context"
	pb "gitlab.ozon.dev/artacone/workshop-1/api"
	"gitlab.ozon.dev/artacone/workshop-1/pkg/cache"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	repository cache.ObjectRepository
	pb.UnimplementedObjectsServiceServer
}

func Run() {
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := newServer()
	//reflection.Register(grpcServer)
	pb.RegisterObjectsServiceServer(grpcServer, server)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func newServer() *server {
	s := &server{}
	s.repository = cache.New()
	return s
}

func (s *server) CreateObject(ctx context.Context, obj *pb.CreateReq) (*pb.CreateResp, error) {
	id, err := s.repository.Create(obj.GetName())

	return &pb.CreateResp{Id: id}, err
}

func (s *server) GetObject(ctx context.Context, id *pb.GetReq) (*pb.GetResp, error) {
	obj, err := s.repository.Get(id.GetId())
	return &pb.GetResp{
		Object: &pb.Object{
			Id: id.GetId(),
			Data: &pb.ObjectData{
				Name: obj.Data.Name,
				Ts:   obj.Data.Ts,
			},
		}}, err
}

func (s *server) EditObject(ctx context.Context, obj *pb.EditReq) (*pb.Empty, error) {
	err := s.repository.Edit(obj.GetId(), obj.GetName())
	return &pb.Empty{}, err
}

func (s *server) DeleteObject(ctx context.Context, id *pb.DelReq) (*pb.Empty, error) {
	err := s.repository.Delete(id.GetId())
	return &pb.Empty{}, err
}
