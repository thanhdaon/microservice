#!/bin/bash

protoc pb/calculator.proto --go_out=plugins=grpc:.