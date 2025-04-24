package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"os"
	"strings"
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

	// Setup data persistence
	aof, err := NewAof("db.aof")
	if err != nil {
		log.Fatalln("Error starting server", err)
	}
	defer aof.Close()

	aof.Read(func(value Value) {
		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			slog.Error("Corrupt db.aof file")
			return
		}

		handler(args)
	})

	go r.acceptLoop(aof)

	<-r.quitch
}

func (r *RedisServer) acceptLoop(aof *Aof) {
	for {
		conn, err := r.ln.Accept()
		if err != nil {
			slog.Error("An error occured", "err", err)
			continue
		}

		go r.readLoop(conn, aof)
	}
}

func (r *RedisServer) readLoop(conn net.Conn, aof *Aof) {
	defer conn.Close()

	for {
		err := readInstructions(conn, aof)
		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Error("Error reading from client: ", "err", err)
			os.Exit(1)
		}

	}
}

func main() {
	flag.Parse()
	redisServer := NewRedisServer(Config{
		port: *serverPort,
	})
	slog.Info(fmt.Sprintf("Server running on port %s\n", redisServer.port))
	redisServer.Run()

}
