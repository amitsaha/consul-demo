#!/bin/bash
CONSUL_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' consul)
echo "#!/bin/sh" > entrypoint.sh
echo "dnsmasq -k -S /consul/$CONSUL_IP#8600 --log-facility=-" >> entrypoint.sh
chmod +x entrypoint.sh
docker build -t amitsaha/dnsmasq .
docker run --network tags_consul_demo --cap-add NET_ADMIN --name dnsmasq amitsaha/dnsmasq
