package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	consulapi "github.com/hashicorp/consul/api"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}

func main() {

	consulConfig := consulapi.DefaultConfig()
	consulClient, err := consulapi.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("err: %v", err)
		os.Exit(1)
	}

	agent := consulClient.Agent()
	reg := &consulapi.AgentServiceRegistration{
		Name: "webapp",
		Port: 8080,
		Tags: []string{"v1"},
	}
	if err := agent.ServiceRegister(reg); err != nil {
		log.Fatalf("err: %v", err)
	}

	http.HandleFunc("/ping/", healthcheck)

	// nil argument here specifies using the DefaultServeMux
	log.Fatal(http.ListenAndServe(":8080", nil))
}
