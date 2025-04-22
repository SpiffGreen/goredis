package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"os"
)

var serverPort = flag.String("port", "6379", "Redis server port")

type Config struct {
	port string
}

type RedisServer struct {
	Config
	ln     net.Listener
	quitch chan struct{}
}

func NewRedisServer(cfg Config) *RedisServer {
	return &RedisServer{
		Config: cfg,
		quitch: make(chan struct{}),
	}
}

func (r *RedisServer) Run() {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", r.port))
	if err != nil {
		log.Fatalln("Error starting server", err)
	}

	defer ln.Close()

	r.ln = ln

	go r.acceptLoop()

	<-r.quitch
}

func (r *RedisServer) acceptLoop() {
	for {
		conn, err := r.ln.Accept()
		if err != nil {
			slog.Error("An error occured", "err", err)
			continue
		}

		go r.readLoop(conn)
	}
}

func (r *RedisServer) readLoop(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		// read messages from client
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Print("Error reading from client: ", err.Error())
			os.Exit(1)
		}

		msg := buf[:n]

		response := readInstructions(msg)

		conn.Write(response)
	}
}

func main() {
	flag.Parse()
	redisServer := NewRedisServer(Config{
		port: *serverPort,
	})
	fmt.Printf("Server running on port %s\n", redisServer.port)
	redisServer.Run()

}
