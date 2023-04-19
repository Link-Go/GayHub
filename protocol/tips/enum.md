## enum


#### Validation
在 gRPC 中，如果客户端传递了一个服务端定义中没有的枚举值，而服务端并没有实现对其进行特殊处理或返回错误提示，可能会导致一些潜在的问题和安全隐患。原因是，gRPC 的默认行为是接受未知的枚举值，并将其视为普通整数类型进行处理

为了避免这种情况发生，建议在 proto 文件中添加 allow_alias 选项，并设置为 false。这样，在客户端发送请求时，如果请求参数中包含了未定义的枚举值，则 gRPC 会抛出 InvalidArgument 错误。例如：

```protobuf
syntax = "proto3";

package myservice;

option go_package = "myservice";

enum MyEnum {
  UNKNOWN = 0;
  VALUE1 = 1;
  VALUE2 = 2;
}

message MyRequest {
  MyEnum enum_field = 1;
}

message MyResponse {
  string message = 1;
}

service MyService {
  rpc MyMethod(MyRequest) returns (MyResponse);
}

// Set allow_alias to false
syntax = "proto3";
package myservice;
import "google/protobuf/duration.proto";

option go_package = "myservice";

enum MyEnum {
  option allow_alias = false; // disallow unknown values
  UNKNOWN = 0;
  VALUE1 = 1;
  VALUE2 = 2;
}

```

在 MyEnum 枚举定义中增加了 allow_alias 选项，并设置为 false。这样，在 MyRequest 请求消息中，如果 enum_field 字段包含了未定义的枚举值，则 gRPC 会抛出 InvalidArgument 错误

注意，将 allow_alias 选项设置为 false 可能会影响一些特定的使用场景，比如在 proto 文件中使用 proto3 JSON 格式进行编码时。因此，在实际应用中，请根据具体情况来选择是否开启或关闭该选项，并进行充分的测试和评估

适用性上，协议可以不添加allow_alias，将所有的校验都放在后端统一处理