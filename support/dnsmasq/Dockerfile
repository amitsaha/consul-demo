FROM alpine:edge
RUN apk update && apk --no-cache add dnsmasq curl bind-tools
COPY entrypoint.sh /entrypoint.sh
CMD ["/entrypoint.sh"]
