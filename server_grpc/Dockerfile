# FROM golang
FROM golang:latest

# Set timezone to Asia/Tokyo
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN apt-get update \
  && apt-get install -y build-essential git curl wget \
                        zlib1g-dev libxml2-dev libxslt1-dev \
                        openssl less libssl-dev libreadline-dev vim unzip

# Install protobuf
# https://github.com/grpc/grpc-docker-library
RUN mkdir -p /tmp/protoc && \
  curl -L https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip > /tmp/protoc/protoc.zip && \
  cd /tmp/protoc && \
  unzip protoc.zip && \
  cp /tmp/protoc/bin/protoc /usr/local/bin && \
  chmod go+rx /usr/local/bin/protoc && \
  cd /tmp && \
  rm -r /tmp/protoc

# Install go protobuf generator
RUN go get google.golang.org/grpc
RUN go get github.com/golang/protobuf/protoc-gen-go

# generate simbolic link for "go get"
RUN ln -s /go/pkg/linux_amd64/* /usr/local/go/pkg/linux_amd64/

