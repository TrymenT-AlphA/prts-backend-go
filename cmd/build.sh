#!/usr/bin/bash
go mod download
cd ./start
go build -o ../build/prts-backend