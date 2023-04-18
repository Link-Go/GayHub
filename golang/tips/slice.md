### slice

#### slice 线程安全
* 方法 1: 加锁
```golang
func main() {
	slc := make([]int, 0, 1000)
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()

			lock.Lock()
			defer lock.Unlock()
			slc = append(slc, a)
		}(i)
	
	}
   wg.Wait()
	fmt.Println(len(slc))
}
```

* 方法 2：channel
```golang

type ServiceData struct {
	ch   chan int // 用来 同步的channel
	data []int    // 存储数据的slice
}

func (s *ServiceData) Schedule() {
	// 从 channel 接收数据
	for i := range s.ch {
		s.data = append(s.data, i)
	}
}

func (s *ServiceData) Close() {
	// 最后关闭 channel
	close(s.ch)
}

func (s *ServiceData) AddData(v int) {
	s.ch <- v // 发送数据到 channel
}

func NewScheduleJob(size int, done func()) *ServiceData {
	s := &ServiceData{
		ch:   make(chan int, size),
		data: make([]int, 0),
	}

	go func() {
		// 并发地 append 数据到 slice
		s.Schedule()
        done()
	}()

	return s
}

func main() {
	var (
		wg sync.WaitGroup
		n  = 1000
	)
	c := make(chan struct{})

	// new 了这个 job 后，该 job 就开始准备从 channel 接收数据了
	s := NewScheduleJob(n, func() { c <- struct{}{} })

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(v int) {
			defer wg.Done()
			s.AddData(v)

		}(i)
	}

	wg.Wait()
	s.Close()
	<-c

	fmt.Println(len(s.data))
}

// 使用 channel 要注意 channel size 的问题
// 注意堵塞风险
// 可以使用无限缓存的channel - chanx
```