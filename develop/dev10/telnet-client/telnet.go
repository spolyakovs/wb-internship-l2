package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Telnet struct {
	host    string
	port    int
	timeout time.Duration
}

func NewTelnet(host string, port int, timeout time.Duration) *Telnet {
	return &Telnet{
		host:    host,
		port:    port,
		timeout: timeout,
	}
}

func (t *Telnet) Start() error {
	fmt.Printf("Starting telnet client on %s:%v\n", t.host, t.port)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	wg := new(sync.WaitGroup)

	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigCh:
				return
			case err := <-errCh:
				if err != nil {
					fmt.Fprintf(os.Stdout, "error: %v\n", err)
					return
				}
			}
		}
	}()

	client, err := newTelnetClient(t.host, t.port, t.timeout)
	if err != nil {
		return err
	}
	defer client.close()

	go writeClient(ctx, client, errCh)
	go receiveClient(ctx, client, errCh)

	wg.Wait()
	return nil
}

func writeClient(ctx context.Context, client *telnetClient, errCh chan<- error) {
	input := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := input.ReadString('\n')
			if err != nil {
				errCh <- err
				return
			}

			if err := client.write(msg); err != nil {
				errCh <- err
				return
			}
		}

	}

}

func receiveClient(ctx context.Context, client *telnetClient, errCh chan<- error) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := client.read(); err != nil {
				errCh <- err
				return
			}
		}
	}
}
