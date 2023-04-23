## 镜像瘦身

#### 构建方式
1. 通过 `docker build` 执行 Dockerfile 里的指令来构建镜像
2. 通过 `docker commit` 在某个容器的基础上生成新的镜像


### 镜像瘦身
 

```dockerfile
FROM ubuntu:focal

ENV REDIS_VERSION=6.0.5
ENV REDIS_URL=http://download.redis.io/releases/redis-$REDIS_VERSION.tar.gz

# update source and install tools
RUN sed -i "s/archive.ubuntu.com/mirrors.aliyun.com/g; s/security.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list 
RUN apt update 
RUN apt install -y curl make gcc

# download source code and install redis
RUN curl -L $REDIS_URL | tar xzv
WORKDIR redis-$REDIS_VERSION
RUN make
RUN make install

# clean up
RUN rm  -rf /var/lib/apt/lists/* 

CMD ["redis-server"]
```

#### RUN指令合并
指令合并是最简单也是最方便的降低镜像层数的方式

```dockerfile
FROM ubuntu:focal

ENV REDIS_VERSION=6.0.5
ENV REDIS_URL=http://download.redis.io/releases/redis-$REDIS_VERSION.tar.gz

# update source and install tools
RUN sed -i "s/archive.ubuntu.com/mirrors.aliyun.com/g; s/security.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list &&\
    apt update &&\
    apt install -y curl make gcc &&\

# download source code and install redis
    curl -L $REDIS_URL | tar xzv &&\
    cd redis-$REDIS_VERSION &&\
    make &&\
    make install &&\

# clean up
    apt remove -y --auto-remove curl make gcc &&\
    apt clean &&\
    rm  -rf /var/lib/apt/lists/* 

CMD ["redis-server"]
```


#### 多阶段构建
```dockerfile
FROM ubuntu:focal AS build

ENV REDIS_VERSION=6.0.5
ENV REDIS_URL=http://download.redis.io/releases/redis-$REDIS_VERSION.tar.gz

# update source and install tools
RUN sed -i "s/archive.ubuntu.com/mirrors.aliyun.com/g; s/security.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list &&\
    apt update &&\
    apt install -y curl make gcc &&\

# download source code and install redis
    curl -L $REDIS_URL | tar xzv &&\
    cd redis-$REDIS_VERSION &&\
    make &&\
    make install 

FROM ubuntu:focal
# copy 
ENV REDIS_VERSION=6.0.5
COPY --from=build /usr/local/bin/redis* /usr/local/bin/

CMD ["redis-server"]
```
* 第一行多了As build, 为后面的COPY做准备 
* 第一阶段中没有了清理操作，因为第一阶段构建的镜像只有编译的目标文件（二进制文件或jar包）有用，其它的都无用 
* 第二阶段直接从第一阶段拷贝目标文件

多个 FROM 指令并不是为了生成多根的层关系，最后生成的镜像，仍以最后一条 FROM 为准，之前的 FROM 会被抛弃。


#### 删除RUN的缓存文件
linux中大部分包管理软件都需要更新源，该操作会带来一些缓存文件，这里记录了常用的清理方法。

* 基于debian的镜像
```bash
# 换国内源，并更新     
sed -i “s/deb.debian.org/mirrors.aliyun.com/g” /etc/apt/sources.list && apt update     
# --no-install-recommends 很有用     
apt install -y --no-install-recommends a b c && rm -rf /var/lib/apt/lists/*

```


* alpine镜像
```bash
# 换国内源，并更新     
sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories     
# --no-cache 表示不缓存     
apk add --no-cache a b c && rm -rf /var/cache/apk/*
```

* centos镜像
```bash
# 换国内源并更新
curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo && yum makecache
yum install -y a b c  && yum clean al
```