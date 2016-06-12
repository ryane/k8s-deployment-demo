package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

const Name = "k8s-deployment-demo"
const Version = "0.1"

type Response struct {
	Version     string
	Hostname    string
	Environment string
	Time        time.Time
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("error getting hostname: %v", err)
	}

	env := os.Getenv("DEMO_ENV")
	if env == "" {
		env = "none"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&Response{
			Version:     Version,
			Hostname:    hostname,
			Environment: env,
			Time:        time.Now(),
		})
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
