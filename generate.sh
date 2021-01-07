#!/bin/bash

protoc ./calcproto/calc.proto --go_out=plugins=grpc:.