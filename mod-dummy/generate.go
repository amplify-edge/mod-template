package main

// ============================================================================
// GO
// ============================================================================
// GRPC & Protobuf
//go:generate /usr/bin/env bash -c "echo 'Generating protobuf and grpc services golang'"
//go:generate protoc -I./proto/v2/ -I. --go_out=./service/go/rpc/v2 --go_opt=paths=source_relative ./proto/v2/mod_dummy_models.proto
//go:generate protoc -I./proto/v2/ -I. --go-grpc_out=./service/go/rpc/v2/ --cobra_out=./service/go/rpc/v2 --go-grpc_opt=paths=source_relative --cobra_opt=paths=source_relative ./proto/v2/mod_dummy_services.proto
//go:generate /usr/bin/env bash -c "echo 'Building example golang binaries (CLI and Server)'"
//go:generate /usr/bin/env bash -c "mkdir -p bin-all/{cli,client}/{go,flutter}"
//go:generate go build -v -o bin-all/cli/go/ ./example/cli/go/main.go

// ============================================================================
// Flutter
// ============================================================================
// GRPC & Protobuf
//go:generate /usr/bin/env bash -c "echo 'generating protobuf and grpc services for flutter/dart'"
//go:generate protoc -I./proto/v2/ -I. --dart_out=grpc:./service/flu/lib/rpc/v2/ ./proto/v2/mod_dummy_models.proto ./proto/v2/mod_dummy_services.proto
//go:generate /usr/bin/env bash -c "echo 'building flutter binary'"
// go:generate make flu-web-build
