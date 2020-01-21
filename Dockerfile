FROM golang:alpine

ADD main.go /tmp/

RUN  apk add --no-cache --update git \
  && go get github.com/garyburd/redigo/redis github.com/gorilla/mux \
  && CGO_ENABLED=0 GOOS=linux go build -o /usr/bin/go-http-listener /tmp/main.go

FROM alpine

RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin/
COPY --from=0 /usr/bin/go-http-listener .

ENTRYPOINT ["/usr/bin/go-http-listener"]
