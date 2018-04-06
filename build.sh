#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -o out/api cmd/rest/main.go 
CGO_ENABLED=0 GOOS=linux go build -o out/cancelledwatcher cmd/cancelledwatcher/main.go
CGO_ENABLED=0 GOOS=linux go build -o out/tradewatcher cmd/tradewatcher/main.go
CGO_ENABLED=0 GOOS=linux go build -o out/tickworker cmd/tickworker/main.go
CGO_ENABLED=0 GOOS=linux go build -o out/marketsworker cmd/marketsworker/main.go
