# FROM golang
FROM golang:latest

# Set timezone to Asia/Tokyo
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN apt-get update \
  && apt-get install -y build-essential git curl wget \
                        zlib1g-dev libxml2-dev libxslt1-dev \
                        openssl less libssl-dev libreadline-dev vim unzip

# Install nats
RUN go get github.com/nats-io/nats-streaming-server

# generate simbolic link for "go get"
RUN ln -s /go/pkg/linux_amd64/* /usr/local/go/pkg/linux_amd64/

