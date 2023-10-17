### ln

Linux ln（英文全拼：link files）命令是一个非常重要命令，它的功能是为某一个文件在另外一个位置建立一个同步的链接。

当我们需要在不同的目录，用到相同的文件时，我们不需要在每一个需要的目录下都放一个必须相同的文件，我们只要在某个固定的目录，放上该文件，然后在 其它的目录下用ln命令链接（link）它就可以，不必重复的占用磁盘空间。



#### 语法

```bash
 ln [参数][源文件或目录][目标文件或目录]
 [-bdfinsvF] [-S backup-suffix] [-V {numbered,existing,simple}]
 [--help] [--version] [--]
  
Usage: ln [OPTION]... [-T] TARGET LINK_NAME   (1st form)
  or:  ln [OPTION]... TARGET                  (2nd form)
  or:  ln [OPTION]... TARGET... DIRECTORY     (3rd form)
  or:  ln [OPTION]... -t DIRECTORY TARGET...  (4th form)
In the 1st form, create a link to TARGET with the name LINK_NAME.
In the 2nd form, create a link to TARGET in the current directory.
In the 3rd and 4th forms, create links to each TARGET in DIRECTORY.
```



##### 功能

Linux文件系统中，有所谓的链接(link)，我们可以将其视为档案的别名，而链接又可分为两种 : 硬链接(hard link)与软链接(symbolic link)，硬链接的意思是一个档案可以有多个名称，而软链接的方式则是产生一个特殊的档案，该档案的内容是指向另一个档案的位置。硬链接是存在同一个文件系统中，而软链接却可以跨越不同的文件系统。

* 文件大小
  * 硬链接：硬链接与原始文件共享相同的数据块，因此它们的文件大小相同。删除任何一个硬链接都不会影响其他硬链接或原始文件。这里的文件大小相同是因为他们的`inode`号相同，指向同一份文件，因此在使用`ll -h`查看的时候会看到硬链接跟原文件的大小一样，因为 `ll –h` 或者 `ls –h` 这命令进行统计文件总大小的时候并不是从磁盘进行统计的，而是根据文件属性中的大小叠加得来的。而硬链接的文件属性中的大小就是就是inode号对应的数据块的大小，所以total中进行统计就把各个文件属性中的大小加起来作为总和，这种统计是不标准的。真正的查看某个文件夹占用磁盘空间大小命令是：`du –h` 这个命令是从磁盘上进行统计，不会被文件的属性中大小影响，使用该指令查看会发现实际使用的磁盘大小只有源文件大小。
  * 软链接：软链接本身是一个文件，它包含指向原始文件的路径。软链接的文件大小通常比原始文件要小得多，因为它只存储了路径信息。删除原始文件会导致软链接失效。
* `Inode` （`stat filePath` 查询 `inode`）
  * 硬链接：硬链接与原始文件具有相同的`inode`号码，它们在文件系统中被视为同一文件的不同名称。因此，对于硬链接和原始文件，`inode`号码是相同的。
  * 软链接：软链接有自己的`inode`号码，它是一个指向原始文件的路径。软链接和原始文件具有不同的`inode`号码。



##### 命令参数

#### 命令参数

**必要参数**：

- --backup[=CONTROL] 备份已存在的目标文件
- -b 类似 **--backup** ，但不接受参数
- -d 允许超级用户制作目录的硬链接
- -f 强制执行
- -i 交互模式，文件存在则提示用户是否覆盖
- -n 把符号链接视为一般目录
- -s 软链接(符号链接)
- -v 显示详细的处理过程

**选择参数**：

- -S "-S<字尾备份字符串> "或 "--suffix=<字尾备份字符串>"
- -V "-V<备份方式>"或"--version-control=<备份方式>"
- --help 显示帮助信息
- --version 显示版本信息



##### 实例

``````
# 链接路径存在，会在指定路径下创建一个同名链接
ln -sv boilerplate.go.txt /root/op/a/b/c/
‘/root/op/a/b/c/boilerplate.go.txt’ -> ‘boilerplate.go.txt’

# 链接路径(c)不存在，创建c并以c作为链接名称
ln -sv boilerplate.go.txt /root/op/a/b/c
‘/root/op/a/b/c’ -> ‘boilerplate.go.txt’

``````

