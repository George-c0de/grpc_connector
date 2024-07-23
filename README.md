# GRPC connector on Golang
## Quick start
### Install package
* go get google.golang.org/grpc
* go get google.golang.org/protobuf
### Generate .proto
* protoc -I proto proto/conn.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
### Run
* server/main.go
* client/main.go
