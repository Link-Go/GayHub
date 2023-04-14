### gRPC

* **[gRPC 官方文档中文版](https://doc.oschina.net/grpc?t=58008)**: V1.0
* **[Protocol Buffers](https://developers.google.com/protocol-buffers/docs/proto3)**

* 实用的镜像
  * git: https://github.com/rvolosatovs/docker-protobuf
  * dockerhub: https://hub.docker.com/r/rvolosatovs/protoc


* 备注说明
  * **旧版本** github.com/golang/protobuf/protoc-gen-go      生成 *.pb.go 文件
  * **新版本** google.golang.org/protobuf/cmd/protoc-gen-go  生成 *.pb.go 和 *_grpc.pb.go 两份文件，合起来就是老版本的代码
  * 目前 gRPC-go 源码中的 example 用的是新版本的生成方式 
    * [代码示例-example](https://github.com/grpc/grpc-go/tree/master/examples)
    * [高级特性-features](https://github.com/grpc/grpc-go/tree/master/examples/features)

