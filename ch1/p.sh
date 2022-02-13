#!/bin/sh 
protoc -I./proto --go_out=plugins=grpc:api proto/*.proto

