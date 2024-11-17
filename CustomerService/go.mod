module customerservice

go 1.20
//protoc --go-grpc_out=. pbCustomer/customer.proto

require (
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.0 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)
