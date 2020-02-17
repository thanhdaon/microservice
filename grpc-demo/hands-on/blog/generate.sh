#/bin/bash

protoc pb/blog.proto --go_out=plugins=grpc:.