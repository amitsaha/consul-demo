FROM golang
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/app
COPY . .
RUN dep ensure && go build main.go && go install
CMD ["/go/bin/app"]
