#!/bin/bash

cd ../proto
protoc --go_out=plugins=grpc:. *.proto
cd -
go run main.go
