#!/bin/bash
docker build -t amitsaha/apiclient .
DNSMASQ_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' dnsmasq)
docker run --network tags_consul_demo --dns $DNSMASQ_IP -ti amitsaha/apiclient sh
