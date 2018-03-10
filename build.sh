#!/usr/bin/env node

go build -o out/api cmd/rest/main.go 
go build -o out/cancelledwatcher cmd/cancelledwatcher/main.go
go build -o out/tradewatcher md/tradewatcher/main.go
go build -o out/tickworker cmd/tickworker/main.go
