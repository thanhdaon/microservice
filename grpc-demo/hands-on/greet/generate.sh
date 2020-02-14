#!/bin/bash

protoc greetpb/greet.proto --go_out=plugins=grpc:.