all: proto_python proto_go

proto_python: proto/data_collector.proto
	python3 -m grpc_tools.protoc -I./proto --python_out=./py/ --pyi_out=./py/ \
	--grpc_python_out=./py/ ./proto/data_collector.proto

proto_go: proto/data_collector.proto
	protoc \
	--go_out ./go/ --go_opt paths=source_relative \
	--go-grpc_out ./go/ --go-grpc_opt paths=source_relative \
	./proto/data_collector.proto