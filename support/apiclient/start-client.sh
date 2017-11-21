#!/bin/bash
docker build -t amitsaha/apiclient .
docker run --network tags_consul_demo --dns 172.22.0.4 -ti amitsaha/apiclient sh

