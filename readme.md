protoc --go_out=plugins=grpc:. your_proto_file.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pb/*.proto