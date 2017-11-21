FROM alpine:edge
RUN apk update && apk --no-cache add dnsmasq curl bind-tools
CMD ["dnsmasq", "-k", "-S", "/consul/172.22.0.2#8600", "--log-facility=-"]
