package main

// 使用 %w 将 error 可以包裹着其他 error
// 可以使用 errors.Is() 检查原始的 error
// 可以使用 errors.As() 将对应的原始 error 提取出来
// golang.org/x/xerrors 打印error时带上堆栈信息功能
// runtime.Stack 也可以获取堆栈信息，通常配合 panic/recover 使用

import (
	"fmt"

	"golang.org/x/xerrors"
)

var myerror = xerrors.New("myerror")

func foo() error {
	return myerror
}
func foo1() error {
	return xerrors.Errorf("foo1 : %w", foo())
}
func foo2() error {
	return xerrors.Errorf("foo2 : %w", foo1())
}
func main() {
	err := foo2()
	fmt.Printf("%v\n", err)
	// foo2 : foo1 : myerror

	fmt.Printf("%+v\n", err)
	// foo2 :
	//     main.foo2
	//         D:/demo/main.go:18
	//   - foo1 :
	//     main.foo1
	//         D:/demo/main.go:15
	//   - myerror:
	//     main.init
	//         D:/demo/main.go:9
}
