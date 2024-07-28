package main

import (
	"context"
	"net"
	pb "postisp/postisprpc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOutpostServiceServer
}

func (s *server) InsertTextItem(ctx context.Context, req *pb.TextItemInsertReq) (*pb.TextItemInsertRes, error) {
	res := &pb.TextItemInsertRes{
		Successful: true,
	}

	return res, nil
}

func (s *server) Retrieve(ctx context.Context, req *pb.RetrieveReq) (*pb.TextItemRetrieveRes, error) {
	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  nil,
	}

	return res, nil
}

func (s *server) RetrieveOnId(ctx context.Context, req *pb.RetrieveOnIdReq) (*pb.TextItemRetrieveRes, error) {
	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  nil,
	}

	return res, nil
}

func (s *server) RetrieveOnTime(ctx context.Context, req *pb.RetrieveOnTimeReq) (*pb.TextItemRetrieveRes, error) {
	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  nil,
	}

	return res, nil
}

func (s *server) RetrieveOnCategory(ctx context.Context, req *pb.RetrieveOnCategoryReq) (*pb.TextItemRetrieveRes, error) {
	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  nil,
	}

	return res, nil
}

func (s *server) RetrieveOnIdAndCategory(ctx context.Context, req *pb.RetrieveOnIdAndCategoryReq) (*pb.TextItemRetrieveRes, error) {
	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  nil,
	}

	return res, nil
}

func (s *server) RetrieveOnTimeAndId(ctx context.Context, req *pb.RetrieveOnTimeAndIdReq) (*pb.TextItemRetrieveRes, error) {
	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  nil,
	}

	return res, nil
}

func (s *server) RetrieveOnTimeAndCategory(ctx context.Context, req *pb.RetrieveOnTimeAndCategoryReq) (*pb.TextItemRetrieveRes, error) {
	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  nil,
	}

	return res, nil
}

func runServer() {
	// TODO: Try to implement running with a mock SSL cert
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	app := &server{}

	pb.RegisterOutpostServiceServer(s, app)

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
