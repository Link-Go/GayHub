## Command

#### 版本对象

Kubernetes 版本支持的所有对象

```bash
kubectl api-resources
```

获取更为详细的信息

```bash
kubectl api-resources -o wide
```

`-o` 有多种参数形式，具体使用 `--help` 查看





#### 资源（字段）解释

```bash
kubectl explain pod
kubectl explain pod.metadata
kubectl explain pod.spec
kubectl explain pod.spec.containers
```





#### 资源详情

```bash
 kubectl describe -n kube-system pod coredns-66bff467f8-4zc5l
 
 # -n 指定命名空间，不指定默认使用default
 # pod 资源类型
 # coredns-66bff467f8-4zc5l 资源名称 
```





#### YAML 样板示例

* pod（run 默认资源为pod）

```bash
kubectl run ngx --image=nginx:alpine --dry-run=client -o yaml
```

* 其他资源的 YAML 样板示例

```bash
kubectl create job echo-job --image=busybox --dry-run=client -o yaml
# kubectl create --help 查看支持的资源
# job 指定资源类型
# echo-job 指定资源名称
# --image=busybox 指定镜像
```



