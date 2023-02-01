

## 压缩神器

压缩可执行文件，如：go编译后文件

### upx 安装

> UPX is an advanced executable file compressor. UPX will typically reduce the file size of programs and DLLs by around 50%-70%, thus reducing disk space, network load times, download times and other distribution and storage costs.

[upx](https://github.com/upx/upx) 是一个常用的压缩动态库和可执行文件的工具，通常可减少 50-70% 的体积。

upx 的安装方式非常简单，我们可以直接从 [github](https://github.com/upx/upx/releases/) 下载最新的 release 版本，支持 Windows 和 Linux，在 Ubuntu 或 Mac 可以直接使用包管理工具安装。例如 Mac 下可以直接使用 brew：

```bash
$ brew install upx
$ upx
Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2020
UPX 3.96        Markus Oberhumer, Laszlo Molnar & John Reiser   Jan 23rd 2020

Usage: upx [-123456789dlthVL] [-qvfk] [-o file] file..
...
Type 'upx --help' for more detailed help.
```

### 使用 upx

upx 有很多参数，`upx --help` 查阅，但最重要的则是压缩率，`1-9`，`1` 代表最低压缩率，`9` 代表最高压缩率。

接下来，我们看一下，如果只使用 upx 压缩，二进制的体积可以减小多少呢。

```bash
$ go build -o server main.go && upx -9 server
        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
  10253684 ->   5210128   50.81%   macho/amd64   server 
$ ls -lh server
-rwxr-xr-x  1 dj  staff   5.0M Dec  8 00:45 server
```

可以看到，使用 upx 后，可执行文件的体积从 9.8M 缩小到了 5M，缩小了 50%。

### upx 的原理

upx 压缩后的程序和压缩前的程序一样，无需解压仍然能够正常地运行，这种压缩方法称之为带壳压缩，压缩包含两个部分：

- 在程序开头或其他合适的地方插入解压代码；
- 将程序的其他部分压缩。

执行时，也包含两个部分：

- 首先执行的是程序开头的插入的解压代码，将原来的程序在内存中解压出来；
- 再执行解压后的程序。

也就是说，upx 在程序执行时，会有额外的解压动作，不过这个耗时几乎可以忽略。

如果对编译后的体积没什么要求的情况下，可以不使用 upx 来压缩。一般在服务器端独立运行的后台服务，无需压缩体积。
