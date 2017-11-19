
```
$ docker-compose up
```

```
$ curl 127.0.0.1:8500/v1/catalog/services
{
    "api": [
        "v2",
        "v1"
    ],
    "consul": []
}
```

docker run --network tags_consul_demo -p 53:53/tcp -p 53:53/udp --cap-add=NET_ADMIN andyshinn/dnsmasq:2.75 -S /consul/172.21.0.2#8600
