package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpccrawler/grpccrawler/crawler"
	pb "grpccrawler/grpccrawler/instagram"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

type instagramServer struct {
	pb.UnimplementedInstagramServer
}

func (s *instagramServer) GetFollowersInfo(ctx context.Context, in *pb.GetFollowersInfoRequest) (*pb.GetFollowersInfoResponse, error) {
	username := in.Username
	fmt.Println("GetFollowersInfo for username: ", username)

	profile, error := crawler.InstagramCrawlerFactory().GetProfileInfo(username)

	return &pb.GetFollowersInfoResponse{Username: username, Followers: profile.Followers, Following: profile.Following}, error
}

func newServer() *instagramServer {
	server := &instagramServer{}
	return server
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = data.Path("x509/server_cert.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = data.Path("x509/server_key.pem")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials: %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }
	fmt.Print("Starting server on port ", *port)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterInstagramServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
