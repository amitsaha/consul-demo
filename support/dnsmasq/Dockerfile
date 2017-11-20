FROM alpine:edge
RUN apk update && apk --no-cache add dnsmasq curl bind-tools
COPY 10-consul /etc/dnsmasq.d/
ENTRYPOINT ["dnsmasq", "-k"]
