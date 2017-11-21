FROM alpine:edge
RUN apk update && apk --no-cache add dnsmasq curl bind-tools
CMD ["dnsmasq", "-k", "-S", "/consul/172.21.0.2#8600", "--log-facility=-"]
#"--listen-address", "127.0.0.1", "--bind-interfaces", "--interface", "lo", 
