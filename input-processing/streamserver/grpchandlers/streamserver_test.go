package grpchandlers

import (
	"context"
	"io"
	"log"
	"net"
	"testing"
	"time"

	pb "github.com/stackpath/backend-developer-tests/input-processing/streamserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var mockListerner *bufconn.Listener

func init() {
	mockListerner = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	serverStream := NewStreamServer()
	pb.RegisterTextStreamerServer(s, serverStream)
	go func() {
		if err := s.Serve(mockListerner); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func dialer(context.Context, string) (net.Conn, error) {
	return mockListerner.Dial()
}

func TestEndtoEndStreaming(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufcon", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufcon: %v", err)
	}
	defer conn.Close()
	client := pb.NewTextStreamerClient(conn)
	// try Execute()
	data := []*pb.TextInput{
		{Message: "Hello there is no error"},
		{Message: "Hello there is no err"},
	}
	outputExpected := []string{"Hello there is no error", ""}
	runStreamerTest(t, client, data, outputExpected)
}

func runStreamerTest(t *testing.T, client pb.TextStreamerClient, data []*pb.TextInput, outputExpected []string) {
	log.Printf("Streaming %v", data)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.FindErrorWord(ctx)
	if err != nil {
		log.Fatalf("%v.Execute(ctx) = %v, %v: ", client, stream, err)
	}
	waitc := make(chan struct{})
	go func() {
		i := 0
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				log.Println("EOF")
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("Err: %v", err)
				continue
			}
			log.Printf("output: %v", result.GetMessage())
			got := result.GetMessage()
			want := outputExpected[i]
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
			i++
		}
	}()

	for _, text := range data {
		if err := stream.Send(text); err != nil {
			log.Fatalf("%v.Send(%v) = %v: ", stream, text, err)
		}
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("%v.CloseSend() got error %v, want %v", stream, err, nil)
	}
	<-waitc
}
