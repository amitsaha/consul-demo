#!/bin/bash
docker build -t amitsaha/dnsmasq .
docker run --network=tags_consul_demo --cap-add=NET_ADMIN -ti amitsaha/dnsmasq
