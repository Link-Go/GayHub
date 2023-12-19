### 使用 kubeadm 创建集群

https://kubernetes.io/zh-cn/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/





### 实现细节

https://kubernetes.io/zh-cn/docs/reference/setup-tools/kubeadm/implementation-details/





### 静态pod

静态Pod 是由 kubelet 拉起的。kubelet 是 Kubernetes 节点上的一个组件，负责管理容器的生命周期，包括拉起和监控静态 Pod。静态 Pod 的配置文件通常存放在节点上的特定目录中`/etc/kubernetes/manifests`，kubelet 会监视这些目录并根据其中的配置文件来创建静态 Pod。

```
# k8s 现阶段存在的静态pod
etcd.yaml  
kube-apiserver.yaml  
kube-controller-manager.yaml  
kube-scheduler.yaml
```



静态pod的启动可以没有k8s，但是需要启动的容器运行时，如：containerd





### 准入控制器

https://kubernetes.io/zh-cn/docs/reference/access-authn-authz/admission-controllers/

```
# 可以通过修改配置文件，增加不同的准入控制器
kube-apiserver.yaml  
```



