// Package to handle grpc stream server
package streamserver

import (
	"context"
	"io"
	"strings"

	pb "github.com/stackpath/backend-developer-tests/input-processing/proto"

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
		if strings.Contains(currentMessage, "error") {
			log.Infof("found error word: %s", currentMessage)
			err = inputStream.Send(&pb.ErrorWordLines{
				Message: currentMessage,
			})
			if err != nil {
				currentErr := status.Errorf(codes.Unknown, "cannot send stream response: %v", err)
				log.Errorf("%v", currentErr)
				return currentErr
			}
		}
	}
	return nil
}

func StreamErrorManager(ctx context.Context) error {
	err := ctx.Err()
	switch err {
	case context.Canceled:
		currentErr := status.Errorf(codes.Canceled, "request is canceled", err)
		log.Errorf("%v", currentErr)
		return currentErr
	case context.DeadlineExceeded:
		currentErr := status.Errorf(codes.DeadlineExceeded, "request is canceled", err)
		log.Errorf("%v", currentErr)
		return currentErr
	default:
		currentErr := status.Errorf(codes.Unknown, "error caught by streaming function", err)
		log.Errorf("%v", currentErr)
		return currentErr
	}
}
