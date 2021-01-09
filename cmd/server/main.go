package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-edtf-http/api"	
	"log"
	"net/http"
)

func main() {

	server_uri := flag.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	flag.Parse()

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	api_parse_handler, err := api.ParseHandler()

	if err != nil {
		log.Fatalf("Failed to API parse handler, %v", err)
	}
	
	mux := http.NewServeMux()
	mux.Handle("/api/parse", api_parse_handler)

	log.Printf("Listening on %s", s.Address())
	s.ListenAndServe(ctx, mux)
}
