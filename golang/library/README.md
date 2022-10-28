## 一些可能有用的 golang 开源库
- **[jsonparser](https://github.com/buger/jsonparser)**
    * 直接在字节层面进行操作 json 字符串，效率更高
    * 不需要创建结构实现对应的 marshal 和 unmarshal 操作
    * 可以在其基础上进行封装，实现默认值的设置（类似python dict.get(key, default)）