### ThreadLocal

`ThreadLocal`表示线程的`局部变量`，它确保每个线程的`ThreadLocal`变量都是各自`独立`的；`ThreadLocal`适合在一个线程的处理流程中保持上下文（避免了同一参数在所有方法中传递），非常适合`web`应用存储当前请求所需要的全局变量，避免了一个线程中，横跨若干方法调用，需要给每个方法增加一个`context`参数，将`context`一路传递的麻烦



```python
# python demo
# 此框架为协程框架，如果请求数大于 workers 数（线程），则无法实现局部变量效果
import asyncio
import random
import threading

import uvicorn
from fastapi import FastAPI


app = FastAPI()
local_storage = threading.local()


@app.get("/")
async def root():
    num = random.randint(0, 100)
    t = threading.currentThread()
    local_storage.val = f"{t.ident}_{num}_abc"
    print(local_storage.val)
    await asyncio.sleep(5)
    return {"message": local_storage.val}


if __name__ == "__main__":
    uvicorn.run("demo:app", host="0.0.0.0", workers=3)

```


备注：
* golang 并不支持`ThreadLocal`，golang建议使用将context一路传递的方式
* 不适合协程框架

