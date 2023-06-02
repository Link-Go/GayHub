## 限制 goroutine num

```golang
func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer func() {
                wg.Done()
                <-ch
            }() 
			log.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}
```