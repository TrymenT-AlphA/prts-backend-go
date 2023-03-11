#!/usr/bin/bash
rm -r ./build
cd ./start
go build -o ../build/prts-backend
cd ../build
./prts-backend