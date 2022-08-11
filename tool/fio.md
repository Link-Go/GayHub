## 磁盘IO测试工具 fio

```bash

# fio -direct=1 -iodepth=64 -rw=read -ioengine=libaio -bs=4k -size=10G -numjobs=1  -name=./fio.test

```
* 主要参数
    - "-direct=1"，代表采用非 buffered I/O 文件读写的方式，避免文件读写过程中内存缓冲对性能的影响。
    - "-iodepth=64"和"-ioengine=libaio"这两个参数，这里指文件读写采用异步 I/O（Async I/O）的方式，也就是进程可以发起多个 I/O 请求，并且不用阻塞地等待 I/O 的完成。稍后等 I/O 完成之后，进程会收到通知。这种异步 I/O 很重要，因为它可以极大地提高文件读写的性能。在这里我们设置了同时发出 64 个 I/O 请求。
    - "-rw=read，-bs=4k，-size=10G"，这几个参数指这个测试是个读文件测试，每次读 4KB 大小数块，总共读 10GB 的数据。
    - "-numjobs=1"，指只有一个进程 / 线程在运行。



* 常用参数
    ```txt
    参数说明：
    filename=/dev/sdb1 测试文件名称，通常选择需要测试的盘的data目录。
    direct=1 是否使用directIO，测试过程绕过OS自带的buffer，使测试磁盘的结果更真实。Linux读写的时候，内核维护了缓存，数据先写到缓存，后面再后台写到SSD。读的时候也优先读缓存里的数据。这样速度可以加快，但是一旦掉电缓存里的数据就没了。所以有一种模式叫做DirectIO，跳过缓存，直接读写SSD。 
    rw=randwrite 测试随机写的I/O
    rw=randrw 测试随机写和读的I/O
    bs=16k 单次io的块文件大小为16k
    bsrange=512-2048 同上，提定数据块的大小范围
    size=5G 每个线程读写的数据量是5GB。
    numjobs=1 每个job（任务）开1个线程，这里用了几，后面每个用-name指定的任务就开几个线程测试。所以最终线程数=任务数（几个name=jobx）* numjobs。 
    name=job1：一个任务的名字，重复了也没关系。如果fio -name=job1 -name=job2，建立了两个任务，共享-name=job1之前的参数。-name之后的就是job2任务独有的参数。 
    thread  使用pthread_create创建线程，另一种是fork创建进程。进程的开销比线程要大，一般都采用thread测试。 
    runtime=1000 测试时间为1000秒，如果不写则一直将5g文件分4k每次写完为止。
    ioengine=libaio 指定io引擎使用libaio方式。libaio：Linux本地异步I/O。请注意，Linux可能只支持具有非缓冲I/O的排队行为（设置为“direct=1”或“buffered=0”）；rbd:通过librbd直接访问CEPH Rados 
    iodepth=16 队列的深度为16.在异步模式下，CPU不能一直无限的发命令到SSD。比如SSD执行读写如果发生了卡顿，那有可能系统会一直不停的发命令，几千个，甚至几万个，这样一方面SSD扛不住，另一方面这么多命令会很占内存，系统也要挂掉了。这样，就带来一个参数叫做队列深度。
    Block Devices（RBD），无需使用内核RBD驱动程序（rbd.ko）。该参数包含很多ioengine，如：libhdfs/rdma等
    rwmixwrite=30 在混合读写的模式下，写占30%
    group_reporting 关于显示结果的，汇总每个进程的信息。
    此外
    lockmem=1g 只使用1g内存进行测试。
    zero_buffers 用0初始化系统buffer。
    nrfiles=8 每个进程生成文件的数量。
    磁盘读写常用测试点：
    1. Read=100% Ramdon=100% rw=randread (100%随机读)
    2. Read=100% Sequence=100% rw=read （100%顺序读）
    3. Write=100% Sequence=100% rw=write （100%顺序写）
    4. Write=100% Ramdon=100% rw=randwrite （100%随机写）
    5. Read=70% Sequence=100% rw=rw, rwmixread=70, rwmixwrite=30
    （70%顺序读，30%顺序写）
    6. Read=70% Ramdon=100% rw=randrw, rwmixread=70, rwmixwrite=30
    (70%随机读，30%随机写)
    ```



* fio例子
    ```txt
    [root@docker sda]# fio -ioengine=libaio -bs=4k -direct=1 -thread -rw=read -filename=/dev/sda -name="BS 4KB read test" -iodepth=16 -runtime=60
    BS 4KB read test: (g=0): rw=read, bs=(R) 4096B-4096B, (W) 4096B-4096B, (T) 4096B-4096B, ioengine=libaio, iodepth=16
    fio-3.7
    Starting 1 thread
    Jobs: 1 (f=1): [R(1)][100.0%][r=89.3MiB/s,w=0KiB/s][r=22.9k,w=0 IOPS][eta 00m:00s]
    BS 4KB read test: (groupid=0, jobs=1): err= 0: pid=18557: Thu Apr 11 13:08:11 2019
    read: IOPS=22.7k, BW=88.5MiB/s (92.8MB/s)(5313MiB/60001msec)
        slat (nsec): min=901, max=168330, avg=6932.34, stdev=1348.82
        clat (usec): min=90, max=63760, avg=698.08, stdev=240.83
        lat (usec): min=97, max=63762, avg=705.17, stdev=240.81
        clat percentiles (usec):
        |  1.00th=[  619],  5.00th=[  627], 10.00th=[  627], 20.00th=[  635],
        | 30.00th=[  635], 40.00th=[  685], 50.00th=[  717], 60.00th=[  725],
        | 70.00th=[  725], 80.00th=[  725], 90.00th=[  734], 95.00th=[  816],
        | 99.00th=[ 1004], 99.50th=[ 1020], 99.90th=[ 1057], 99.95th=[ 1057],
        | 99.99th=[ 1860]
    bw (  KiB/s): min=62144, max=91552, per=100.00%, avg=90669.02, stdev=3533.77, samples=120
    iops        : min=15536, max=22888, avg=22667.27, stdev=883.44, samples=120
    lat (usec)   : 100=0.01%, 250=0.01%, 500=0.01%, 750=93.85%, 1000=5.14%
    lat (msec)   : 2=0.99%, 4=0.01%, 10=0.01%, 50=0.01%, 100=0.01%
    cpu          : usr=5.35%, sys=23.17%, ctx=1359692, majf=0, minf=17
    IO depths    : 1=0.1%, 2=0.1%, 4=0.1%, 8=0.1%, 16=100.0%, 32=0.0%, >=64=0.0%
        submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
        complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.1%, 32=0.0%, 64=0.0%, >=64=0.0%
        issued rwts: total=1360097,0,0,0 short=0,0,0,0 dropped=0,0,0,0
        latency   : target=0, window=0, percentile=100.00%, depth=16
    
    Run status group 0 (all jobs):
    READ: bw=88.5MiB/s (92.8MB/s), 88.5MiB/s-88.5MiB/s (92.8MB/s-92.8MB/s), io=5313MiB (5571MB), run=60001-60001msec
    
    Disk stats (read/write):
    sda: ios=1357472/0, merge=70/0, ticks=949141/0, in_queue=948776, util=99.88%
    ```

io=执行了多少M的IO

bw=平均IO带宽
iops=IOPS
runt=线程运行时间
slat=提交延迟，提交该IO请求到kernel所花的时间（不包括kernel处理的时间）
clat=完成延迟, 提交该IO请求到kernel后，处理所花的时间
lat=响应时间
bw=带宽
cpu=利用率
IO depths=io队列
IO submit=单个IO提交要提交的IO数
IO complete=Like the above submit number, but for completions instead.
IO issued=The number of read/write requests issued, and how many of them were short.
IO latencies=IO完延迟的分布

io=总共执行了多少size的IO
aggrb=group总带宽
minb=最小.平均带宽.
maxb=最大平均带宽.
mint=group中线程的最短运行时间.
maxt=group中线程的最长运行时间.

ios=所有group总共执行的IO数.
merge=总共发生的IO合并数.
ticks=Number of ticks we kept the disk busy.
io_queue=花费在队列上的总共时间.
util=磁盘利用率

fio 有很多测试任务配置文件，在git工程 examples 文件夹中，我们可以使用命令行参数进行直接配置，也可以直接通过配置文件配置一次测试的内容。

更详细对fio输出说明请参考博文：[Fio Output Explained](https://tobert.github.io/post/2014-04-17-fio-output-explained.html) 