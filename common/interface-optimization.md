## interface optimization

* 1. 批量思想：批量操作数据库
* 2. 异步思想：耗时操作
    * 调用方在该次操作中只关注调用的成功与否，并不在意调用的具体逻辑实现
        * 如：销毁某个对象，对于调用方来说，他只关注通知是否发出，后台是否成功收到。剩下的销毁操作，释放资源的逻辑，调用方其实并不关注，此时便可以异步触发销毁操作
    * 存在有其他的机制保证调用方接收到该消息，如：定时器
        * 如：渲染文件等任务进程，触发任务后，调用方可能会进行其他的操作，不能因为任务的耗时导致其他操作无法进行/等待的时间过长。这个时候便可以使用异步触发任务的方式。再由定时器定时查询任务的完成状态，返回给调用方
* 3. 空间换时间思想：使用缓存
    * 使用缓存需要保证对数据的时效性不敏感，是一些热点数据，会多次频繁使用的数据。如果对数据的实效性要求过高，且使用之后就丢弃，则不建议使用缓存
* 4. 预存思想：提前使用缓存，在初始化的时候就将数据加载到缓存中
    * 如：小说数据
* 5. 事件回调的方式：可以参考下epoll的实现思路
    * 类似第二点的定时器，只不过这个事件通知是由调用方在被调用方注册了一个函数，被调用方完成任务后主动通知调用方；定时器是由调用方多次轮询查看任务状态
    * select, poll, epoll的优劣点自行百度
* 6. 池化思想：预分配/重复使用
    * 如：http-conn对象可以在服务启动的时候预先初始化，使用完成后，下一次请求复用该连接对象
    * 可以参考golang/net/http/client.go Clent:57 的说明与使用
    可以参考golang/net/http/transport.go Transport:95 的说明与使用
* 7. 远程调用由同步改为异步
    * 如：爬虫的时候，多个url请求存在大量的io阻塞，这个时候应该将同步改为异步
* 8. 锁的粒度要尽量的细
    * 防止锁的力度过大，锁内代码效率不大，多个请求竞争锁，导致接口过慢
    * 防止代码异常，锁一直得不到释放，出现死锁的现象
* 9. sql 优化
* 10. 代码逻辑结构
    * 针对具体业务具体分析