### CentOS环境安装Protobuf

1. 下载地址: https://github.com/protocolbuffers/protobuf/releases；最好是下载-all的包,因为里面的依赖文件比较全,不然还需要下载各种依赖,可能会遇到各种报错
2. 解压编译
*  ```bash
    tar -zxvf protobuf-all-21.9.tar.gz
    cd protobuf-21.9
    ./configure
    make
    make install
   ```
   最后两步比较的慢,耐心等待就行,完成后运行下面的命令可以看到版本的信息
   执行 protoc --version 会显示对应版本信息 libprotoc 3.21.9 说明安装成功

3. ```bash
    # 使用官方文档的demo时，需要满足前置操作
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        helloworld/helloworld.proto

    # 前置操作，保证可以找到protoc-gen-go，protoc-gen-go-grpc两个执行文件
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

    export PATH="$PATH:$(go env GOPATH)/bin"

    # 或者直接进入
    cd $(go env GOPATH)/bin
    cp protoc-gen-go /usr/local/bin/
    cp protoc-gen-go-grpc /usr/local/bin/
   ```

#### Q&A
* ```bash
    ...
    configure: error: C++ preprocessor "/lib/cpp" fails sanity check
    ...

    缺少必要的C++库

    yum install -y glibc-headers gcc-c++

  ```

