#! /bin/bash

protoc --proto_path=taskprotobuf --go_out=plugins=grpc:pkg/taskpb tasks.proto