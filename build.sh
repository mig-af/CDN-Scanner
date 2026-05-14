#!/bin/bash

#linux
echo "Compilando linux"
go build -o recondomain


echo "compilando android"
GOOS=android GOARCH=arm64 go build -o recondomain-android 
