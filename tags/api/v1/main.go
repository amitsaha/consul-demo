package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	consulapi "github.com/hashicorp/consul/api"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there! I am v1")
}

func main() {

	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = "consul:8500"
	consulClient, err := consulapi.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("err: %v", err)
		os.Exit(1)
	}

	agent := consulClient.Agent()
	reg := &consulapi.AgentServiceRegistration{
		ID:      "apiv1",
		Address: "172.21.0.3",
		Name:    "api",
		Port:    8080,
		Tags:    []string{"v1"},
	}
	if err := agent.ServiceRegister(reg); err != nil {
		log.Fatalf("err: %v", err)
	}

	http.HandleFunc("/ping/", healthcheck)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
