@rm -r ./build>nul 2>nul
@cd ./start
@go build -o ../build/prts-backend.exe
@cd ../build
@prts-backend.exe