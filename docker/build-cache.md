## docker build 的 cache 机制



#### docker build 简介
众所周知，一个 Dockerfile 唯一的定义了一个 Docker 镜像。如此依赖，Docker 必须提供一种方式，将 Dockerfile 转换为 Docker 镜像，采用的方式就是 docker build 命令 。以如下的 Dockerfile 为例：

```dockerfile
FROM ubuntu:14.04 
RUN apt-get update 
ADD run.sh / 
VOLUME /data 
CMD ["./run.sh"]  
```

一般此 Dockerfile 的当前目录下，必须包含文件 run.sh。通过执行以下命令：
`docker build -t="my_new_image" . `
即可将当前目录下的 Dockerfile 构建成一个名为 my_new_image  的镜像，镜像的默认 tag 为 latest。对于以上的docker build 请求，Docker Daemon 新创建了 4 层镜像，除了 FROM 命令，其余的 RUN、ADD、VOLUME 以及 CMD 命令都会创建一层新的镜像。

每一次重新build，docker 都会遍历本地镜像，发现镜像与即将构建出的镜像一致时，将找到的镜像作为 cache 镜像，复用 cache 镜像作为构建结果。


#### cache 机制注意事项

1. ADD 命令与 COPY 命令

Dockerfile 没有发生任何改变，但是命令ADD run.sh /  中 Dockerfile 当前目录下的 run.sh 却发生了变化，从而将直接导致镜像层文件系统内容的更新，原则上不应该再使用 cache。那么，判断 ADD 命令或者 COPY 命令后紧接的文件是否发生变化，则成为是否延用 cache 的重要依据。Docker 采取的策略是：获取 Dockerfile 下内容（包括文件的部分 inode 信息），计算出一个唯一的 hash 值，若 hash 值未发生变化，则可以认为文件内容没有发生变化，可以使用 cache 机制；反之亦然。

2. RUN 命令存在外部依赖

一旦 RUN 命令存在外部依赖，如RUN apt-get update ，那么随着时间的推移，基于同一个基础镜像，一年的 apt-get update 和一年后的 apt-get update， 由于软件源软件的更新，从而导致产生的镜像理论上应该不同。如果继续使用 cache 机制，将存在不满足用户需求的情况。Docker 一开始的设计既考虑了外部依赖的问题，用户可以使用参数 --no-cache 确保获取最新的外部依赖，命令为docker build --no-cache -t="my_new_image" . 

3. 树状的镜像关系决定了，一次新镜像层的成功构建将导致后续的 cache 机制全部失效

即某一层的镜像层是全新的镜像层，没有使用缓存。那后面的所有镜像层都会重新生成。这也是为什么，书写 Dockerfile 时，应该将更多静态的安装、配置命令尽可能地放在 Dockerfile 的较前位置。