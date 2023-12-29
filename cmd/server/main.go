package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/353solutions/rides/pb"
)

func main() {
	addr := ":9292"

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("error: can't listen - %s", err)
	}

	srv := grpc.NewServer()
	var u Rides
	pb.RegisterRidesServer(srv, &u)

	log.Printf("info: server ready on %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s", err)
	}
}

type Rides struct {
	pb.UnimplementedRidesServer
}

func (r *Rides) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	// TODO: Valid request
	resp := &pb.StartResponse{
		Id: req.Id,
	}
	
	// TODO: Work
	return resp, nil
}
