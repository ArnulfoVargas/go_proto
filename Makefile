run:
	@go run ./cmd/
proto-gen:
	@protoc --go_out=./cmd/ ./protobuf/*.proto
	@protoc --go_out=./cmd/ ./protobuf/job/*.proto
	@protoc --go_out=./cmd/ ./protobuf/software/*.proto
	@protoc --go_out=./cmd/ ./protobuf/comunication/*.proto