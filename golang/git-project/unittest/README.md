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
* 单元测试框架：
	*	**[testify](https://github.com/stretchr/testify)**
	    * assert package: 友好的代码断言工具包
	    * require package: 提供全局函数
	    * mock package: 插桩
	    * suite package: 面向对象方式的单元测试
	*	**[goconvey](https://github.com/smartystreets/goconvey)** 
		* import goconvey包时，前面加点号"."，以减少冗余的代码；
		* 测试函数的名字必须以Test开头，而且参数类型必须为*testing.T；
		* 每个测试用例必须使用Convey语句包裹起来，推荐使用Convey语句的嵌套，即一个函数有一个或多个测试函数，一个测试函数嵌套两层、三层或四层Convey语句；
		* Convey语句的第三个参数习惯以闭包的形式实现，在闭包中通过So语句完成断言；
		* 使用GoConvey框架的 Web 界面特性，作为命令行的补充；
		* 在适当的场景下使用SkipConvey函数或SkipSo函数；
		* 当测试中有需要时，可以定制断言函数。
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
		* [gomonkey](https://github.com/agiledragon/gomonkey/v2) (推荐使用)
			* 支持为一个函数打一个桩
			* 支持为一个函数打一个特定的桩序列
			* 支持为一个成员方法打一个桩
			* 支持为一个成员方法打一个特定的桩序列
			* 支持为一个函数变量打一个桩
			* 支持为一个函数变量打一个特定的桩序列
			* 支持为一个接口打桩
			* 支持为一个接口打一个特定的桩序列
			* 支持为一个全局变量打一个桩