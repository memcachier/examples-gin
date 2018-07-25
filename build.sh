#!/usr/bin/env bash
set -xe

# get all of the dependencies needed
go get "github.com/gin-gonic/gin"
go get "github.com/memcachier/mc"

# create the application binary that eb uses
GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"
