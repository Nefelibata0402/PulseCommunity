protoc --proto_path=../ --proto_path=. --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative article_model.proto article_service.proto
