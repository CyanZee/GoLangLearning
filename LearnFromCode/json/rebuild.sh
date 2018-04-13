#!/bin/bash

go build jsonTester.go

#GOOS=windows GOARCH=amd64 go build -o jsonTester64.exe jsonTester.go
#GOOS=windows GOARCH=386 go build -o jsonTester32.exe jsonTester.go
#GOOS=linux GOARCH=arm64 go build -o jsonTesterArm jsonTester.go
