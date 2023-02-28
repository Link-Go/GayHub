## 网络诊断工具

背景：在请求网络资源获取缓慢或者有丢包过程中。经常会使用到网络路径探测工具。linux 下最常用的有mtr、traceroute、tracepath 等

## 原理浅析

实现数据在网络中的交互，都基于最基础的网络地址：源--->目的地址。这个地址为IP地址。为了避免无效的数据包垃圾在网络中永久存在。有8位的TTL值。

在互联网中传输过程中，IP数据包每经过一个路由器该值减一，直到为0 时将该包丢弃。

路由探测原理就是从发送TTL 值为1开始的数据包开始，每次增加1，直到该数据包能抵达目的IP地址。

未能抵达最终目的地址的数据包，当路由器将值减为0 时，会给源地址返回一个数据包告知源地址：因经过路由的数据过多，导致TTL耗尽。数据包无法到达最终目的地。工具根据中间路由节点这个数据包返回的时间戳和发出数据包时的时间戳相减，计算出中间经过的每个路由节点的耗时。并获取中间路由节点的IP地址。

## mtr

`traceroute` 命令只会做一次链路跟踪测试，而mtr命令会对链路上的相关节点做持续探测并给出相应的统计信息。所以，mtr命令能避免节点波动对测试结果的影响，所以其测试结果更正确，建议优先使用。

`mtr` 默认使用ICMP协议发送探测数据包

```bash
    mtr -u  使用udp
        -T  使用tcp
        -r  报告形式输出
        -w  完整的报告
        -P  指定端口

[root@iZwz9in4retlu6s4wr8rh8Z ~]# mtr --help
usage: mtr [-BfhvrwctglxspQomniuT46] [--help] [--version] [--report]
                [--report-wide] [--report-cycles=COUNT] [--curses] [--gtk]
                [--csv|-C] [--raw] [--xml] [--split] [--mpls] [--no-dns] [--show-ips]
                [--address interface] [--filename=FILE|-F]
                [--ipinfo=item_no|-y item_no]
                [--aslookup|-z]
                [--psize=bytes/-s bytes] [--order fields]
                [--report-wide|-w] [--inet] [--inet6] [--max-ttl=NUM] [--first-ttl=NUM]
                [--bitpattern=NUM] [--tos=NUM] [--udp] [--tcp] [--port=PORT] [--timeout=SECONDS]
                [--interval=SECONDS] HOSTNAME

[root@iZwz9in4retlu6s4wr8rh8Z ~]# mtr -w -T www.baidu.com
Start: Tue Feb 28 15:32:54 2023
HOST: iZwz9in4retlu6s4wr8rh8Z Loss%   Snt   Last   Avg  Best  Wrst StDev
  1.|-- 10.12.32.250             0.0%    10    0.8 402.7   0.5 3006. 968.9
  2.|-- 10.12.32.77             20.0%    10  1004. 502.3   0.7 1004. 536.1
  3.|-- 11.54.240.237            0.0%    10    0.6   0.6   0.4   0.7   0.0
  4.|-- 10.102.253.225           0.0%    10    3.5   1.6   1.2   3.5   0.7
  5.|-- 117.49.38.2             70.0%    10  3016. 1010.   6.8 3016. 1737.6
  6.|-- 113.96.60.21            90.0%    10  8909. 8909. 8909. 8909.   0.0
  7.|-- 58.61.162.153           80.0%    10    2.5   2.5   2.5   2.6   0.0
  8.|-- 119.147.221.237         40.0%    10    3.1 171.2   2.7 1004. 408.0
  9.|-- 113.96.5.102            60.0%    10   12.5 261.9   8.0 1013. 500.9
 10.|-- 121.14.14.162            0.0%    10    7.9 109.8   6.4 1022. 320.6
 11.|-- 14.29.117.246            0.0%    10    5.6   5.9   5.2   7.5   0.5
 12.|-- ???                     100.0    10    0.0   0.0   0.0   0.0   0.0
 13.|-- 14.215.177.38           80.0%    10  151.3 147.5 143.6 151.3   5.4

数字代表一共经过了多少次转发触发目标
Loss： 节点丢包率
Snt： 已发送数据包数。默认值是10，可以通过参数-c指定
Last：最近一次的探测延迟值
Avg：探测延迟的平均值
Best：探测延迟的最小值
Wrst：探测延迟的最大值
StDev：标准偏差。该值越大说明相应节点越不稳定
```

## traceroute



```bash

traceroute  -I 使用ICMP
            -T 使用TCP （默认80端口）
            -4 IPV4
            -6 IPV6
            -P 指定端口

[root@iZwz9in4retlu6s4wr8rh8Z ~]# traceroute --help
Usage:
  traceroute [ -46dFITnreAUDV ] [ -f first_ttl ] [ -g gate,... ] [ -i device ] [ -m max_ttl ] [ -N squeries ] [ -p port ] [ -t tos ] [ -l flow_label ] [ -w waittime ] [ -q nqueries ] [ -s src_addr ] [ -z sendwait ] [ --fwmark=num ] host [ packetlen ]


[root@iZwz9in4retlu6s4wr8rh8Z ~]# traceroute -I www.baidu.com
traceroute to www.baidu.com (14.215.177.39), 30 hops max, 60 byte packets
 1  10.12.36.70 (10.12.36.70)  0.469 ms * *
 2  * * *
 3  10.255.99.89 (10.255.99.89)  0.791 ms  0.793 ms  0.812 ms
 4  103.41.141.250 (103.41.141.250)  4.849 ms  4.810 ms  4.846 ms
 5  117.49.38.34 (117.49.38.34)  8.196 ms  8.224 ms  8.282 ms
 6  * * *
 7  58.61.162.181 (58.61.162.181)  2.517 ms  2.521 ms  2.488 ms
 8  119.147.221.233 (119.147.221.233)  2.893 ms  2.940 ms  3.178 ms
 9  * 113.96.4.14 (113.96.4.14)  10.403 ms  10.004 ms
10  94.96.135.219.broad.fs.gd.dynamic.163data.com.cn (219.135.96.94)  7.579 ms  7.472 ms  7.478 ms
11  14.29.121.190 (14.29.121.190)  5.592 ms  6.020 ms  5.672 ms
12  * * *
13  * * *
14  14.215.177.39 (14.215.177.39)  7.848 ms  7.836 ms  7.825 ms

* * * 证明某些节点没有返回、无法访问、被屏蔽、超时等情况
唯一可以保证便是第14跳收到了14.215.177.39返回的信息
```


## tracepath

功能太少，仅支持udp，不建议使用

```bash
[root@iZwz9in4retlu6s4wr8rh8Z ~]# tracepath
Usage: tracepath [-n] [-b] [-l <len>] [-p port] <destination>
```