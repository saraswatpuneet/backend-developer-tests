package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	pb "github.com/stackpath/backend-developer-tests/input-processing/proto"
	"github.com/stackpath/backend-developer-tests/input-processing/streamserver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (  
    PORT = ":9192"  
) 

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// // Read STDIN into a new buffered reader
	// reader := bufio.NewReader(os.Stdin)

	// // TODO: Look for lines in the STDIN reader that contain "error" and output them.

	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			fmt.Println(err)
	// 			os.Exit(1)
	// 		}
	// 	}
	// 	if strings.Contains(line, "error") {
	// 		fmt.Println(line)
	// 	}
	// }
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	holdSignal := make(chan os.Signal, 1)
	signal.Notify(holdSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	// Create a new stream server
	streamServer := streamserver.NewStreamServer()
	// grpcListerner for rest api
	lis, err := net.Listen("tcp", PORT)  
    if err != nil {  
        log.Fatalf("failed to listen: %v", err)  
    }  
	if err != nil {
		log.Fatal("cannot start rpc server: ", err)
	}
	err = startRestServer(ctx, streamServer, "8091", lis)
	if err != nil {
		log.Fatal("cannot start rest server: ", err)
	}
	<-holdSignal
	// stdIOListener for grpc server
	stdinReader := bufio.NewReader(os.Stdin)
	stdOutWrite := bufio.NewWriter(os.Stdout)
	ioPipe := streamserver.NewStdStreamServer(stdinReader, stdOutWrite)
	ioListerner := &streamserver.StandardIO{}
	go startStandarIOServer(streamServer, ioListerner)
	ioListerner.Ready(ioPipe)
	// if system throw any termination stuff let channel handle it and cancel

}
func startStandarIOServer(
	streamServer pb.TextStreamerServer,
	listener net.Listener,
) error {
	grpcServer := grpc.NewServer()
	pb.RegisterTextStreamerServer(grpcServer, streamServer)
	reflection.Register(grpcServer)
	log.Infof("GRPC Server tied to Std I/O started at %s", listener.Addr().String())
	return grpcServer.Serve(listener)
}

func startRestServer(
	ctx context.Context,
	streamServer pb.TextStreamerServer,
	port string,
	listener net.Listener,
) error {
	server := grpc.NewServer()
	pb.RegisterTextStreamerServer(server, streamServer)
	go func() {
		log.Fatalln(server.Serve(listener))
	}()
	gwmux, err := newGateway(ctx)
	if err != nil {
		log.Errorf("error encountered while creating gateway: %v", err)
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	log.Infof("grpc-gateway listen on localhost:%v", 8091)
	return http.ListenAndServe(":8091", mux)
}

func newGateway(ctx context.Context) (http.Handler, error) {

	opts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterTextStreamerHandlerFromEndpoint(ctx, gwmux, ":9192", opts); err != nil {
		return nil, err
	}

	return gwmux, nil
}
