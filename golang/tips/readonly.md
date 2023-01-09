## 只读变量

#### 实现一
* 使用包导入的方式
* 提供读取方法，不提供设置方法；还需保证传出的对象非指针对象
* 这种方式仍可以通过反射的方式获取/设置变量
```golang
package config

import (
    "reflect"
)

type ConfigDemo struct {
    field1 string
    field2 string
}

func NewConfigDemo() *ConfigDemo {
	return &ConfigDemo{
		field1: "demo1",
		field2: "demo2",
	}
}

func (c *ConfigDemo) GetField1() string {
    return c.field1
}

func (c *ConfigDemo) GetField2() string {
    return c.field2
}


// 获取未导出变量
func getUnExportedField(ptr interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	return v
}

// 设置未导出变量
func setUnExportedStrField(ptr interface{}, fieldName string, newFieldVal interface{}) (err error) {
	// 获取非导出字段反射对象
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	// 获取非导出字段可寻址反射对象
	// 与上面的区别是：这个是可寻址的
	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()

	nv := reflect.ValueOf(newFieldVal)
	if v.Kind() != nv.Kind() {
		return fmt.Errorf("expected kind %v, got kind: %v", v.Kind(), nv.Kind())
	}
	v.Set(nv)
	return nil
}

```



#### 实现二
* 使用`mmap`的方式（文件/设备映射到内存），实际操作的是文件的读写权限
```golang
package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Protect []byte

func NewProtect(size int) (Protect, error) {
	b, err := syscall.Mmap(0, 0, size, syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_ANON | syscall.MAP_PRIVATE)
	if err != nil {
		return nil, err
	}
	return Protect(b), nil
}

func (p Protect) Free() error {
	return syscall.Munmap(p)
}

func (p Protect) ReadOnly() error {
	return syscall.Mprotect([]byte(p), syscall.PROT_READ)
}

func (p Protect) ReadWrite() error {
	return syscall.Mprotect([]byte(p), syscall.PROT_READ | syscall.PROT_WRITE)
}

func (p Protect) Pointer() unsafe.Pointer {
	return unsafe.Pointer(&p[0])
}


func main() {
	pv, err := NewProtect(8)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer pv.Free()

	p := (*int)(pv.Pointer())
	*p = 100
	fmt.Println("1:", *p)

	pv.ReadOnly()
	fmt.Println("2:", *p)

	pv.ReadWrite()
	*p += 100
	fmt.Println("3:", *p)

	pv.ReadOnly()
	*p += 100 // 报错
	fmt.Println("4:", *p) 
}
```
