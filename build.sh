#!/usr/bin/env bash

go build -o out/api cmd/api/main.go
go build -o out/cancelledwatcher cmd/cancelledwatcher/main.go
go build -o out/tradewatcher cmd/tradewatcher/main.go
go build -o out/tickworker cmd/tickworker/main.go
go build -o out/marketsworker cmd/marketsworker/main.go
go build -o out/balancewatcher cmd/balancewatcher/main.go
