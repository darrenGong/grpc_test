#!/bin/sh

protoc -I helloworld/ helloworld/hello.proto --go_out=plugins=grpc:helloworld
