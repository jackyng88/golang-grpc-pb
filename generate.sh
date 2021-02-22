#!/bin/bash

protoc -I=greet/greetpb --go_out=greet/greetpb --go-grpc_out=greet/greetpb greet.proto

protoc -I=calculator/calculatorpb --go_out=calculator/calculatorpb --go-grpc_out=calculator/calculatorpb calculator.proto
