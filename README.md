# HOW TO
1) make sure you have GO, PROTOC and compiler plugins for GO installed
2) Generate gRPC GO files by running (at helloworld folder):
```go
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto
```
3) start the server, form helloworld folder: go run server/main.go
4) start the client, form helloworld folder: go run client/main.go

# FULL DOCUMENTATION
https://grpc.io/docs/languages/go/quickstart/#prerequisites