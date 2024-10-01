module proto

go 1.22.4

require google.golang.org/protobuf v1.34.2 // direct
replace protobuf => ./cmd/protobuf
require protobuf v0.0.1