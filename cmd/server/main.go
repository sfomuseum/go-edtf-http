package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-edtf-http/api"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	"net/http"
	"os"
)

func main() {

	fs := flagset.NewFlagSet("server")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "HTTP server for exposing sfomuseum/go-edtf-http handlers.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flagset.Parse(fs)

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	api_parse_handler, err := api.ParseHandler()

	if err != nil {
		log.Fatalf("Failed to API parse handler, %v", err)
	}

	api_valid_handler, err := api.IsValidHandler()

	if err != nil {
		log.Fatalf("Failed to API is valid handler, %v", err)
	}

	api_matches_handler, err := api.MatchesHandler()

	if err != nil {
		log.Fatalf("Failed to API is matches handler, %v", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/api/parse", api_parse_handler)
	mux.Handle("/api/valid", api_valid_handler)
	mux.Handle("/api/matches", api_matches_handler)

	log.Printf("Listening on %s", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

}
