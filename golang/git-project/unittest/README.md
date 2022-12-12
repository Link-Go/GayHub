## 单元测试

* TestMain 函数
    * TestMain是一个特殊的函数（相当于 main 函数），测试用例在执行时，会先执行TestMain函数，然后可以在TestMain中调用m.Run()函数执行普通的测试函数。在m.Run()函数前面我们可以编写准备逻辑，在m.Run()后面我们可以编写清理逻辑。
    ```golang
        func TestMain(m *testing.M) {
            fmt.Println("do some setup")
            m.Run()
            fmt.Println("do some cleanup")
        }
    ```
* 单元测试框架： **[testify](https://github.com/stretchr/testify)**
    * assert package: 友好的代码断言工具包
    * require package: 提供全局函数
    * mock package: 插桩
    * suite package: 面向对象方式的单元测试
* 单元测试生成工具： **[gotests工具](https://github.com/cweill/gotests)**
* Mock 测试
    * **[mock仓库](https://github.com/golang/mock)**
    ```bash
        $ go get github.com/golang/mock/gomock
        $ go install github.com/golang/mock/mockgen  // 通过mock工具生成mock代码
    ```
    * 其他mock工具
        * [sqlmock](https://github.com/DATA-DOG/go-sqlmock)
        * [httpmock](https://github.com/jarcoal/httpmock)
        * [monkey](https://github.com/bouk/monkey)