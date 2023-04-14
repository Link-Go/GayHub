## protoc-gen-validate
自动生成校验代码，可集成至grpc

* **[github](https://github.com/bufbuild/protoc-gen-validate)**: V1.0

* 前置条件
    * [install](../install.md)
    * [protocol-readme](../README.md)

#### Build from source
```bash
go get -d github.com/envoyproxy/protoc-gen-validate

or 

git clone github.com/bufbuild/protoc-gen-validate

then

cd protoc-gen-validate && make build
```


#### example
```protobuf
syntax = "proto3";

package examplepb;
option go_package = "/;examplepb";
import "validate/validate.proto";

message Person {
  uint64 id = 1 [(validate.rules).uint64.gt = 999];

  string email = 2 [(validate.rules).string.email = true];

  string name = 3 [(validate.rules).string = {
    pattern:   "^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$",
    max_bytes: 256,
  }];

  Location home = 4 [(validate.rules).message.required = true];

  message Location {
    double lat = 1 [(validate.rules).double = {gte: -90,  lte: 90}];
    double lng = 2 [(validate.rules).double = {gte: -180, lte: 180}];
  }
}
```

#### 常见问题
* Import "validate/validate.proto" was not found or had errors
* 需要复制`validate/validate.proto`到工作区以使其工作
* issue: https://github.com/bufbuild/protoc-gen-validate/issues/368
```bash
protoc \
    -I=/your-proto-folder \
    -I=${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
    --go_opt=paths=source_relative \
    --go_out=/your-code-folder \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_out=/your-code-folder \
    --validate_out=paths=source_relative,lang=go:/your-code-folder \
    your-proto.proto
```


