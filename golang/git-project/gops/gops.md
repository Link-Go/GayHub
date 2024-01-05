### gops

https://github.com/google/gops

gops是google开发的一个在你系统上列出和分析golang程序的工具。它也是一个golang的库，当你集成这个库以后，分析golang程序会变得很简单。



##### 安装

go get -u [github.com/google/gops](http://github.com/google/gops)



##### agent

在你的程序中集成 gops 的 agent

```golang
package main

import (
	"log"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{
        Addr: "listen tcp address",  // 你想监听的地址，不填的话系统会自动分配一个端口给它用。
    }); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
```



##### gops的使用

当你的程序嵌入了gops的agent，那么你就可以对你的golang程序进行诊断了。 除了诊断golang程序，gops程序也可以列出当前系统的golang程序和其系统，即时那些golang程序没有嵌入gops的agent。 注意的是`只有程序嵌入了gops的agent才可以进行程序的分析，包括程序的堆、内存信息、cpu性能分析等`



```bash
sh-5.0$ ./gops -h

gops is a tool to list and diagnose Go processes.

Usage:
  gops [flags]
  gops [command]

Examples:
  gops <cmd> <pid|addr> ...
  gops <pid> # displays process info
  gops help  # displays this help message

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  gc          Runs the garbage collector and blocks until successful.
  help        Help about any command
  memstats    Prints the allocation and garbage collection stats.
  pprof-cpu   Reads the CPU profile and launches "go tool pprof".
  pprof-heap  Reads the heap profile and launches "go tool pprof".
  process     Prints information about a Go process.
  setgc       Sets the garbage collection target percentage. To completely stop GC, set to 'off'
  stack       Prints the stack trace.
  stats       Prints runtime stats.
  trace       Runs the runtime tracer for 5 secs and launches "go tool trace".
  tree        Display parent-child tree for Go processes.
  version     Prints the Go version used to build the program.

Flags:
  -h, --help   help for gops

Use "gops [command] --help" for more information about a command.


pprof-cpu/pprof-heap 指令可以生成对应的文件，配合 go tool pprof 使用，进行分析
```



