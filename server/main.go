package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tmcnicol/multi-language-webserver/api"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	log.Println("Starting Server")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	name := strings.TrimPrefix(message, "/")

	log.Println("Handling request:", name)

	response := sendToServer(name)

	w.Write([]byte(response))
}

func sendToServer(name string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewHandlerClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	r, err := c.HandleGeneric(ctx, &api.Request{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r.Message
}
