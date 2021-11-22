// This file wraps standand stream server with managed stream server
// Acts like piping server to input stream and output stream
package streamserver

import (
	"io"
	"net"
	"time"
)

type StdStreamServer struct {
	inReader      io.Reader
	outWriter     io.Writer
	isClosed      bool
	localAddress  *StdInAddress
	remoteAddress *StdInAddress
}

func NewStdStreamServer(inReader io.Reader, outWriter io.Writer) *StdStreamServer {
	return &StdStreamServer{inReader, outWriter, false, NewStdInAddress("local_add"), NewStdInAddress("remote_add")}
}

func (s *StdStreamServer) LocalAddr() net.Addr {
	return s.localAddress
}

func (s *StdStreamServer) RemoteAddr() net.Addr {
	return s.remoteAddress
}

func (s *StdStreamServer) Read(b []byte) (n int, err error) {
	return s.inReader.Read(b)
}

func (s *StdStreamServer) Write(b []byte) (n int, err error) {
	return s.outWriter.Write(b)
}

func (s *StdStreamServer) Close() error {
	s.isClosed = true
	return nil
}

func (s *StdStreamServer) SetDeadline(t time.Time) error {
	return nil
}

func (s *StdStreamServer) SetReadDeadline(t time.Time) error {
	return nil
}

func (s *StdStreamServer) SetWriteDeadline(t time.Time) error {
	return nil
}
