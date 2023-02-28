## 一些可能有用的 golang 开源库
- **[jsonparser](https://github.com/buger/jsonparser)**
    * 直接在字节层面进行操作 json 字符串，效率更高
    * 不需要创建结构实现对应的 marshal 和 unmarshal 操作
    * 可以在其基础上进行封装，实现默认值的设置（类似python dict.get(key, default)）

- **[endless](https://github.com/fvbock/endless/)**
    * 实现 http 服务热更新
    * 文章：https://grisha.org/blog/2014/06/03/graceful-restart-in-golang/
    * 原理：
        * 服务器要拒绝新的连接请求，但要保持已有的连接
        * 父进程 fork 一个子进程，将socket交给子进程，由子进程接收处理新的请求
        * 通知父进程关闭，父进程优雅退出（处理完原先的请求链接）
        * init 进程接管子进程

- **[go-spew](https://github.com/davecgh/go-spew)**
    * 变量数据结构调试利器 go-spew
    * 可以将`struct`字段所有的信息，所有层级具体的数据结构，包含类型、字段、字段类型、字段值等信息输出到指定位置

- **[lo](https://github.com/samber/lo)**
    * 为高效循环而创建迭代器的函数
    * 可以更好的处理一个数据集合，map, slice 等迭代遍历操作
    * version >= go 1.18

- **[mapstructure](https://github.com/mitchellh/mapstructure)**
    * `map to struct`
    * 在某些数据流（json）中，你不太清楚数据的具体结构，需要通过其中的某个字段（type）来判断数据的结构。可以将该数据转为`map[string]interface{}`，识别后再转换为`struct`
    * demo
    ```golang
    package main

    import (
        "fmt"

        "github.com/mitchellh/mapstructure"
    )

    func ExampleDecode() {
        type Person struct {
            Name   string
            Age    int
            Emails []string
            Extra  map[string]string
        }

        // This input can come from anywhere, but typically comes from
        // something like decoding JSON where we're not quite sure of the
        // struct initially.
        input := map[string]interface{}{
            "name":   "Mitchell",
            "age":    91,
            "emails": []string{"one", "two", "three"},
            "extra": map[string]string{
                "twitter": "mitchellh",
            },
        }

        var result Person
        err := mapstructure.Decode(input, &result)
        if err != nil {
            panic(err)
        }

        fmt.Printf("%#v", result)
        // Output:
        // mapstructure.Person{Name:"Mitchell", Age:91, Emails:[]string{"one", "two", "three"}, Extra:map[string]string{"twitter":"mitchellh"}}
    }
    ```

- **[gin-dump](https://github.com/tpkeeper/gin-dump)**
    * 输出出req ,res的header和body内容，方便观察请求和相应结果
    * 代码实现了body数据无法多次去除的逻辑

- **[automaxprocs](https://github.com/uber-go/automaxprocs)**
    * 自动设置GOMAXPROCS以匹配 Linux 容器 CPU 配额
    * [详细说明](./automaxprocs.md)

- **[allocate](https://github.com/cjrd/allocate)**
    * 初始化struct内各字段为默认零值而非nil值
