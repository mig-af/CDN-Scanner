#!/bin/bash

#linux
echo "Compilando linux"
go build -o cdnscanner


echo "compilando android"
GOOS=android GOARCH=arm64 go build -o cdnscanner-android 
