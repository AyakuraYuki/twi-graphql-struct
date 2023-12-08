# Setup

1. Set `${GOBIN}` environment
2. Add `${GOBIN}` to `${PATH}`
3. Install the following plugins
    ```shell
    go install github.com/bufbuild/buf/cmd/buf@latest
    go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@latest
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install github.com/favadi/protoc-go-inject-tag@latest
    ```
