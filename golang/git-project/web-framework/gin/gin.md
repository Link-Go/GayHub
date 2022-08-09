### Gin 

#### 核心功能
* 支持 HTTP 方法：GET、POST、PUT、PATCH、DELETE、OPTIONS
* 支持不同位置的 HTTP 参数：路径参数（path）、查询字符串参数（query）、表单参数（form）、HTTP 头参数（header）、消息体参数（body）
* 支持 HTTP 路由和路由分组
    * 路由支持精准匹配/模糊匹配
* 支持 middleware 和自定义 middleware
* 支持自定义 Log
* 支持 binding 和 validation，支持自定义 validator。可以 bind 如下参数：query、path、body、header、form
    * gin 引用 github.com/go-playground/validator/v10； 使用 binding 标签
    * gin 基于 ShouldBindWith 和 MustBindWith 这两个函数，又衍生出很多新的 Bind 函数
* 支持重定向
* 支持 basic auth middleware
* 支持自定义 HTTP 配置
* 支持优雅关闭
* 支持 HTTP2
* 支持设置和获取 cookie



#### 中间件

| 中间件            | 功能                                                         |
| :---------------- | ------------------------------------------------------------ |
| gin-jwt           | jwt中间件，实现jwt认证                                       |
| gin-swagger       | 自动生成swagger 2.0格式的 restful api 文档                   |
| cors              | 实现http请求跨域                                             |
| sessions          | 会话管理中间件                                               |
| authz             | 基于casbin的授权中间件                                       |
| pprof             | gin pprof 中间件                                             |
| go-gin-prometheus | prometheus metrics exporter                                  |
| gzip              | 支持http请求和响应的gzip压缩                                 |
| gin-limit         | http请求并发控制中间件                                       |
| requestID         | 给每个request生成uuid，并添加再返回的X-Request-ID Header中<br />如果请求头中含有这个字段，则使用此字段 |



#### 优雅关停
* Go 1.8 版本或者更新的版本，http.Server 内置的 Shutdown 方法，已经实现了优雅关闭

```golang

// +build go1.8

package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
    time.Sleep(5 * time.Second)
    c.String(http.StatusOK, "Welcome Gin Server")
  })

  srv := &http.Server{
    Addr:    ":8080",
    Handler: router,
  }

  // Initializing the server in a goroutine so that
  // it won't block the graceful shutdown handling below
  go func() {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatalf("listen: %s\n", err)
    }
  }()

  // Wait for interrupt signal to gracefully shutdown the server with
  // a timeout of 5 seconds.
  // 使用有缓存的channel，防止信号丢失
  quit := make(chan os.Signal, 1)
  // kill (no param) default send syscall.SIGTERM
  // kill -2 is syscall.SIGINT
  // kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  log.Println("Shutting down server...")

  // The context is used to inform the server it has 5 seconds to finish
  // the request it is currently handling
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  if err := srv.Shutdown(ctx); err != nil {
    log.Fatal("Server forced to shutdown:", err)
  }

  log.Println("Server exiting")
}
```

* 使用第三方库
[fvbock/endless](https://github.com/fvbock/endless)