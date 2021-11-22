package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	pb "github.com/stackpath/backend-developer-tests/input-processing/proto"

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

	conn, err := grpc.DialContext(ctx, fmt.Sprintf(":%s", PORT), transportOption)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}
	stream, err := pb.NewTextStreamerClient(conn).FindErrorWord(ctx, transportOption)
	if err != nil {
		log.Fatal(err)
	}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
			if err = stream.Send(&pb.TextInput{Message: line}); err != nil {
				log.Fatalf("failed to send data to grpc server: %v", err)
			}
		}
	}
}
