## 定时器



#### timer

* timer 指定时间触发，只触发一次


```golang
package main

import "time"

func main() {
	timer := time.NewTimer(time.Second)
	<-timer.C
	// do something
}
```



#### ticker

* ticker 指定触发时间间隔，循环触发

```golang
package main

import "time"

func main() {
	ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

	for range ticker.C {
		// do something
	}
}
```



#### 比较难看，且存在瞬时误差的写法

```golang
package main

import "time"

func main() {
	for {
		time.Sleep(time.Second)
		// do something
	}
}
```