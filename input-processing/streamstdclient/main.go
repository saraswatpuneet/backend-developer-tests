package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	pb "github.com/stackpath/backend-developer-tests/input-processing/streamstdclient/proto"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	PORT = ":9192"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()
	transportOption := grpc.WithInsecure()

	conn, err := grpc.Dial(":9192", transportOption)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}
	streamObj := pb.NewTextStreamerClient(conn)
	stream, err := streamObj.FindErrorWord(ctx)
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}
	go runReaderRouting(reader, stream)
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Fatalf("server stream closed or terminated or client terminated the session%v", err)
			break
		}
		if err != nil {
			log.Fatalf("receive stream error %v", err)
			break
		}
		log.Infof("%v", resp.Message)
	}
}

func runReaderRouting(reader *bufio.Reader, stream pb.TextStreamer_FindErrorWordClient) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Error(err)
				os.Exit(1)
			}
		}
		if err = stream.Send(&pb.TextInput{Message: line}); err != nil {
			log.Fatalf("failed to send data to grpc server: %v", err)
		}
	}
}
