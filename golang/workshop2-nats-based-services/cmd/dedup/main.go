package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {

	url, exists := os.LookupEnv("NATS_URL")
	if !exists {
		url = "nats://0.0.0.0:4222"
	} else {
		url = strings.TrimSpace(url)
	}

	if strings.TrimSpace(url) == "" {
		url = nats.DefaultURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	srv, err := micro.AddService(nc, micro.Config{
		Name:        "dedup",
		Version:     "1.0.0",
		Description: "deduplication service",
		QueueGroup:  "dedup-group",
	})

	fmt.Printf("Created service: %s (%s)\n", srv.Info().Name, srv.Info().ID)

	if err != nil {
		log.Fatal(err)
		return
	}

	root := srv.AddGroup("dedup")

	root.AddEndpoint("dedup", micro.HandlerFunc(dedupHandler))

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status": "OK"}`))
		})
		log.Println("HTTP healthcheck server listening on :9001")
		if err := http.ListenAndServe(":9001", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Blocking call to keep the service running (simulate service work)
	select {}
}

func dedupHandler(req micro.Request) {
	var incoming map[string][]string
	err := json.Unmarshal(req.Data(), &incoming)
	if err != nil {
		req.Error("dedup.error", "Failed to unmarshal request data", []byte(err.Error()))
		return
	}

	unique := make(map[string]struct{})
	for _, item := range incoming["data"] {
		unique[item] = struct{}{}
	}
	var result []string
	for item := range unique {
		result = append(result, item)
	}
	responseMap := make(map[string][]string)
	responseMap["data"] = result
	response, err := json.Marshal(responseMap)
	if err != nil {
		req.Error("dedup.error", "Failed to marshal response data", []byte(err.Error()))
		return
	}
	req.Respond(response)
	fmt.Printf("Received request: %s\n", string(req.Data()))
	fmt.Printf("Sending response: %s\n", string(response))
}
