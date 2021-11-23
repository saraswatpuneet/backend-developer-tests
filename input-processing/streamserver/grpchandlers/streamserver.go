// Package to handle grpc stream server
package grpchandlers

import (
	"context"
	"io"
	"strings"

	pb "github.com/stackpath/backend-developer-tests/input-processing/streamserver/proto"

	// import logrus
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// streamserver is used to implement streamserver.TextStreamerServer.
type StreamServer struct {
	pb.UnimplementedTextStreamerServer
}

// NewStreamServer creates a new streamserver
func NewStreamServer() *StreamServer {
	return &StreamServer{}
}

// FindErrorWord implements streamserver.FindErrorWord
func (s *StreamServer) FindErrorWord(inputStream pb.TextStreamer_FindErrorWordServer) error {
	log.Infof("FindErrorWord called")
	for {
		err := StreamErrorManager(inputStream.Context())
		if err != nil {
			return err
		}
		req, err := inputStream.Recv()
		if err == io.EOF {
			log.Info("no more data is streaming: break the connection")
			break
		}
		if err != nil {
			currentErr := status.Errorf(codes.Unknown, "cannot receive stream request: %v", err)
			log.Errorf("%v", currentErr)
			return currentErr
		}
		currentMessage := req.GetMessage()
		// check if message contains word "error" using strings package
		// split currentMessage into lines
		lines := strings.Split(currentMessage, "\n")
		// iterate through lines
		for _, line := range lines {
			// check if line contains word "error"
			if strings.Contains(line, "error") {
				// send error message to client
				err = inputStream.Send(&pb.ErrorWordLines{Message: line})
				if err != nil {
					currentErr := status.Errorf(codes.Unknown, "cannot send stream response: %v", err)
					log.Errorf("%v", currentErr)
					return currentErr
				}
			}
		}
	}
	return nil
}

func StreamErrorManager(ctx context.Context) error {
	err := ctx.Err()
	switch err {
	case context.Canceled:
		currentErr := status.Errorf(codes.Canceled, "request is canceled %v", err)
		log.Errorf("%v", currentErr)
		return currentErr
	case context.DeadlineExceeded:
		currentErr := status.Errorf(codes.DeadlineExceeded, "request is canceled: %v", err)
		log.Errorf("%v", currentErr)
		return currentErr
	default:
		if err != nil {
			currentErr := status.Errorf(codes.Unknown, "error caught by streaming function %v", err)
			log.Errorf("%v", currentErr)
			return currentErr
		}
	}
	return nil
}
