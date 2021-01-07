#!/bin/bash

protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative calc.proto