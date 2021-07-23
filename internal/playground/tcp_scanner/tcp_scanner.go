package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

const MAX_TCP_PORTS = 65535

// Helper function to check if a port is open
func isOpen(host string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

// Based on the "Writing TCP scanner in Go" article (https://developer20.com/tcp-scanner-in-go/).
//
// TCP connection handshake has 3 steps (simplified):
//
// 	1. Client sends `syn` package. If the client gets a timeout here, it may
//	   mean that the port is behind a firewall.
//	2. Server answers with `syn-ack` when the port is opened, otherwise it
//	   responds with `rst` package.
//	3. Client has to send another packet called `ack`. Connection is
//	   established.
//
func main() {
	// Command line options
	hostname := flag.String("hostname", "", "name of the host to test")
	startPort := flag.Int("start", 0, "port to start scanning")
	endPort := flag.Int("end", MAX_TCP_PORTS, "port to stop scanning")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeout")
	flag.Parse()

	// Let's use concurrency to quickly scan which ports are open
	var openPorts []int
	wg := &sync.WaitGroup{}

	// We need to use a mutex to handle writing to the openPorts list
	mtx := &sync.Mutex{}

	// TODO: Goroutine pool?
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mtx.Lock()
				openPorts = append(openPorts, p)
				mtx.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()

	fmt.Printf("Opened ports: %v\n", openPorts)
}
