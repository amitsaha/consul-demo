package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	consulapi "github.com/hashicorp/consul/api"
)

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there! I am v2")
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
		ID:      "apiv2",
		Address: GetLocalIP(),
		Name:    "api",
		Port:    8080,
		Tags:    []string{"v2"},
	}
	if err := agent.ServiceRegister(reg); err != nil {
		log.Fatalf("err: %v", err)
	}

	http.HandleFunc("/ping/", healthcheck)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
