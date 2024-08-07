package outpost

import (
	"context"
	"flag"
	"fmt"
	"net"
	pb "outpost/outpostrpc"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 5055, "the port to serve on")

type server struct {
	pb.UnimplementedOutpostServiceServer
	persister TextItemPersister
}

func (s *server) InsertTextItem(ctx context.Context, req *pb.TextItem) (*pb.TextItemInsertRes, error) {
	textItem := TextItem{
		Text:      req.Text,
		RefTag:    req.RefTag,
		Timestamp: req.Timestamp,
		Category:  req.Category,
	}

	if err := s.persister.Insert(textItem); err != nil {
		return nil, err
	}

	res := &pb.TextItemInsertRes{
		Successful: true,
	}

	return res, nil
}

func (s *server) Retrieve(ctx context.Context, req *pb.RetrieveReq) (*pb.TextItemRetrieveRes, error) {
	textItems, err := s.persister.Retrieve()

	if err != nil {
		return nil, err
	}

	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  textItems,
	}

	return res, nil
}

func (s *server) RetrieveOnRefTag(ctx context.Context, req *pb.RetrieveOnRefTagReq) (*pb.TextItemRetrieveRes, error) {
	textItems, err := s.persister.RetrieveOnRefTag(req.RefTag)

	if err != nil {
		return nil, err
	}

	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  textItems,
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

func (s *server) RetrieveOnRefTagAndCategory(ctx context.Context, req *pb.RetrieveOnRefTagAndCategoryReq) (*pb.TextItemRetrieveRes, error) {
	res := &pb.TextItemRetrieveRes{
		Successful: true,
		TextItems:  nil,
	}

	return res, nil
}

func (s *server) RetrieveOnTimeAndRefTag(ctx context.Context, req *pb.RetrieveOnTimeAndRefTagReq) (*pb.TextItemRetrieveRes, error) {
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

func Run(persister TextItemPersister) {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}

	// TODO: Implement with mock SSL cert
	// creds, err := credentials.NewServerTLSFromFile("", "")
	// if err != nil {
	// 	panic(err)
	// }

	// s := grpc.NewServer(grpc.Creds(creds))
	s := grpc.NewServer()

	pb.RegisterOutpostServiceServer(s, &server{persister: persister})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
