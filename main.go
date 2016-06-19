package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

const Name = "k8s-deployment-demo"
const Version = "0.2"

var hostname, env string

type Response struct {
	Version     string
	Hostname    string
	Environment string
	Time        time.Time
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(BuildResponse())
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func BuildResponse() *Response {
	var err error
	if hostname == "" {
		hostname, err = os.Hostname()
		if err != nil {
			log.Fatalf("error getting hostname: %v", err)
		}
	}

	if env == "" {
		env = os.Getenv("DEMO_ENV")
		if env == "" {
			env = "none"
		}
	}

	return &Response{
		Version:     Version,
		Hostname:    hostname,
		Environment: env,
		Time:        time.Now(),
	}
}
