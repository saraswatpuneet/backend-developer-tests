package grpchandlers

import (
	"fmt"
	"net"
	"sync"
)

type StandardIO struct {
	closed   bool
	wait     sync.WaitGroup
	onlyConn net.Conn
}

func (lis *StandardIO) Ready(conn net.Conn) {
	lis.onlyConn = conn
	lis.wait.Done()
}

func (lis *StandardIO) Accept() (net.Conn, error) {
	lis.wait.Add(1)
	fmt.Println("accepting, waiting for only conn")
	lis.wait.Wait()
	fmt.Println("accepting, only conn ready")
	return lis.onlyConn, nil
}

func (lis *StandardIO) Close() error {
	lis.closed = true
	return nil
}

func (lis *StandardIO) Addr() net.Addr {
	return NewStdInAddress("listener")
}
