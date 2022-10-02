package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type telnetClient struct {
	addr       string
	conn       net.Conn
	connReader *bufio.Reader
}

func newTelnetClient(host string, port int, timeout time.Duration) (*telnetClient, error) {
	addr := fmt.Sprintf("%s:%v", host, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	return &telnetClient{
		addr:       addr,
		conn:       conn,
		connReader: bufio.NewReader(conn),
	}, nil
}

func (tc *telnetClient) write(msg string) error {
	_, err := tc.conn.Write([]byte(msg))
	return err
}

func (tc *telnetClient) read() error {
	msg, err := tc.connReader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Fprint(os.Stdout, msg)
	return nil
}

func (tc *telnetClient) close() error {
	return tc.conn.Close()
}
