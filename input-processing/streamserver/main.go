package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	pb "github.com/stackpath/backend-developer-tests/input-processing/streamserver/proto"
	"github.com/stackpath/backend-developer-tests/input-processing/streamserver/grpchandlers"
	"google.golang.org/grpc"
)

const (
	PORT = ":9192"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	holdSignal := make(chan os.Signal, 1)
	signal.Notify(holdSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	// Create a new stream server
	streamServer := grpchandlers.NewStreamServer()
	// grpcListerner for rest api
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("cannot start rpc server listener: ", err)
	}
	log.Infof("grpc endpoint listen on localhost:%v", PORT)
	go startRestServer(ctx, streamServer, lis)
	// if system throw any termination stuff let channel handle it and cancel
	<-holdSignal
}

func startRestServer(
	ctx context.Context,
	streamServer pb.TextStreamerServer,
	listener net.Listener,
) error {
	server := grpc.NewServer()
	pb.RegisterTextStreamerServer(server, streamServer)
	go func() {
		log.Fatalln(server.Serve(listener))
	}()
	gwmux, err := generateGetwayLayer(ctx)
	if err != nil {
		log.Errorf("error encountered while creating gateway: %v", err)
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	log.Infof("grpc-gateway listen on localhost:%v", 8091)
	return http.ListenAndServe(":8091", mux)
}

func generateGetwayLayer(ctx context.Context) (http.Handler, error) {

	opts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterTextStreamerHandlerFromEndpoint(ctx, gwmux, ":9192", opts); err != nil {
		return nil, err
	}

	return gwmux, nil
}
