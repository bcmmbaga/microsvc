#! /bin/bash

protoc --proto_path=api/taskprotobuf --go_out=plugins=grpc:pkg/taskpb tasks.proto