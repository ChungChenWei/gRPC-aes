package main

import (
	cipher "cipher"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 3000, "The server port")
)

type server struct {
	cipher.UnimplementedCipherServiceServer
}

var AES_KEY = os.Getenv("AES_KEY")
var AES_IV = os.Getenv("AES_IV")

func (s *server) Encrypt(ctx context.Context, in *cipher.EncryptRequest) (*cipher.EncryptResponse, error) {
	return &cipher.EncryptResponse{Crypted: cipher.Encrypt(in.Plain, AES_KEY, AES_IV)}, nil
}

func (s *server) Decrypt(ctx context.Context, in *cipher.DecryptRequest) (*cipher.DecryptResponse, error) {
	return &cipher.DecryptResponse{Plain: cipher.Decrypt(in.Crypted, AES_KEY, AES_IV)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cipher.RegisterCipherServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
