## python定时器Timer



### 非异步实现

##### 单次执行
* threading 中的 Timer，通过事件注册的方式实现定时任务的执行
    * threading.Timer.start() 只会触发一次


```python
class Timer(Thread):
    """Call a function after a specified number of seconds:

            t = Timer(30.0, f, args=None, kwargs=None)
            t.start()
            t.cancel()     # stop the timer's action if it's still waiting

    """

    def __init__(self, interval, function, args=None, kwargs=None):
        Thread.__init__(self)
        self.interval = interval
        self.function = function
        self.args = args if args is not None else []
        self.kwargs = kwargs if kwargs is not None else {}
        self.finished = Event()

    def cancel(self):
        """Stop the timer if it hasn't finished yet."""
        self.finished.set()

    def run(self):
        self.finished.wait(self.interval)
        if not self.finished.is_set():
            self.function(*self.args, **self.kwargs)
        self.finished.set()
```



##### 循环执行
* 不建议方式，重复注册 Timer；每次循环，系统都要创建一个线程，然后再回收，interval小时开销很大

```python
def hello(name, string):
        print(f"hello : {name} ,nice to : {string}")
        
t = Timer(10, hello, ("怎料事与愿违", "不愿染是与非"))
t.start()
```


* 重写 Timer.run 方法

```python
class RepeatingTimer(Timer):
    def run(self):
        while not self.finished.is_set():
            self.function(*self.args, **self.kwargs)
            self.finished.wait(self.interval)

class UseTimer:
    def __init__(self, interval, function_name, *args, **kwargs):
        """
        :param interval:时间间隔
        :param function_name:可调用的对象
        :param args:args和kwargs作为function_name的参数
        """
        self.timer = RepeatingTimer(interval, function_name, *args, **kwargs)

    def timer_start(self):
        self.timer.start()

    def timer_cancle(self):
        self.timer.cancel()
```



### 异步实现

```python
    asyncio.sleep(interval) # 核心代码
```

[官方文档参考](https://docs.python.org/zh-cn/3.7/library/asyncio.html)
