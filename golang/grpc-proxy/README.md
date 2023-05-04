### proxy

可以不使用 pb 文件，进行 grpc 请求的反向代理

1. grpc.UnknownServiceHandler
* 将未注册的服务进行统一处理

2. proxy.StreamDirector
* 将所有请求根据 methodName 进行区分处理

3. github.com/mwitkow/grpc-proxy/proxy
* proxy.TransparentHandler
    * s.forwardServerToClient
    * s.forwardClientToServer
