package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	broadcaster = flag.Bool("broadcaster", false, "enable broadcast mode")
	port        = flag.String("port", "8080", "port to listen on")
	connections = make(map[net.Conn]int)
	connMutex   sync.RWMutex
	connCounter int64
	totalBytes  int64
)

func main() {
	flag.Parse()
	// handle --help
	if *broadcaster {
		log.Println("Broadcast mode enabled. Messages will be sent to all connected clients.")
	} else {
		log.Println("Echo mode enabled. Messages will be echoed back to the sender.")
	}

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal("error starting server:", err)
	}
	defer listener.Close()

	log.Println("tcp loopback server listening on port", *port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection:", err)
			continue
		}

		connID := int(atomic.AddInt64(&connCounter, 1))
		connMutex.Lock()
		connections[conn] = connID
		connMutex.Unlock()

		log.Printf("connection %d established", connID)
		go handleConnection(conn, connID)
	}
}

func handleConnection(conn net.Conn, connID int) {
	defer func() {
		conn.Close()
		connMutex.Lock()
		delete(connections, conn)
		connMutex.Unlock()
		log.Printf("connection %d closed", connID)
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := strings.ToLower(scanner.Text())

		if *broadcaster {
			broadcastMessage(message, conn, connID)
		} else {
			bytes := len(message) + 1
			atomic.AddInt64(&totalBytes, int64(bytes))
			fmt.Fprintf(conn, "%s\n", message)
			log.Printf("connection %d echoed %d bytes", connID, bytes)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error reading from connection %d: %v", connID, err)
	}
}

func broadcastMessage(message string, sender net.Conn, senderID int) {
	connMutex.RLock()
	defer connMutex.RUnlock()

	bytes := len(message) + 1
	for conn, connID := range connections {
		if conn != sender {
			fmt.Fprintf(conn, "%s\n", message)
			atomic.AddInt64(&totalBytes, int64(bytes))
			log.Printf("broadcasted %d bytes to connection %d", bytes, connID)
		}
	}
}
