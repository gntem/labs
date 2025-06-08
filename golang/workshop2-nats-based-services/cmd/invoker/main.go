package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
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

	defer nc.Close()

	var strs []string
	strs = append(strs, "foo")
	strs = append(strs, "bar")
	strs = append(strs, "bar")

	var requestData = map[string]interface{}{
		"data": strs,
	}

	requestDataJSON, marshalErr := json.Marshal(requestData)
	if marshalErr != nil {
		log.Fatal(marshalErr)
		return
	}

	msg, reqErr := nc.Request("dedup.dedup", requestDataJSON, 2*time.Second)
	if reqErr != nil {
		log.Fatal(reqErr)
		return
	}
	if msg == nil {
		log.Fatal("No response received")
		return
	}
	var responseMap map[string][]string
	err = json.Unmarshal(msg.Data, &responseMap)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(">Response received")
	responseData := responseMap["data"]
	fmt.Printf("Response: %v\n", responseData)
	fmt.Printf("Response length: %d\n", len(responseData))
}
