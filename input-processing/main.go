package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	pb "github.com/stackpath/backend-developer-tests/input-processing/proto"
	"github.com/stackpath/backend-developer-tests/input-processing/streamserver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	// Create a new stream server
	streamServer := streamserver.NewStreamServer()
	// grpcListerner for rest api
	address := fmt.Sprintf("localhost:%d", 8091)
	startRestServer(streamServer, address)
	// stdIOListener for grpc server
	stdinReader := bufio.NewReader(os.Stdin)
	stdOutWrite := bufio.NewWriter(os.Stdout)
	ioPipe := streamserver.NewStdStreamServer(stdinReader, stdOutWrite)
	ioListerner := &streamserver.StandardIO{}
	go startStandarIOServer(streamServer, ioListerner)
	ioListerner.Ready(ioPipe)
	holdSignal := make(chan os.Signal, 1)
	signal.Notify(holdSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	// if system throw any termination stuff let channel handle it and cancel
	<-holdSignal

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
	streamServer pb.TextStreamerServer,
	grpcEndpoint string,
) error {
	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := pb.RegisterTextStreamerHandlerFromEndpoint(ctx, mux, grpcEndpoint, dialOptions)
	if err != nil {
		return err
	}
	port := strings.Split(grpcEndpoint, ":")[1]
	addressToServe := fmt.Sprintf(":%s", port)
	log.Infof("Rest API tied to FindError started at", grpcEndpoint)

	return http.ListenAndServe(addressToServe, mux)
}
