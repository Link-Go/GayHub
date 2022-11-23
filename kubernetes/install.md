## k8s 部署安装文档

* 安装环境

  * 基于centos7

* 可能存在问题

  * 安装 kubeadm 提示: No package kubeadm available

    * 源文件，清理/备份 /etc/yum.repos.d/kubernetes.repo 文件

    * ```bash
      cat <<EOF > /etc/yum.repos.d/kubernetes.repo
      [kubernetes]
      name=Kubernetes
      baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
      enabled=1
      gpgcheck=0
      repo_gpgcheck=0
      gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
      http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
      EOF
      ```

  * Kubeadm初始化报错：`[ERROR CRI]: container runtime is not running`

    * 确保容器运行时已经下载且正常启动，如：docker
    * 1.20 及以上版本已经不再支持docker，需要使用其他的容器运行时，如：podman
    * 如果使用docker作为容器运行时，请指定下载的 kubeadm 的版本

  * Kubeadm初始化报错: `[ERROR KubeletVersion]: the kubelet version is higher than the control plane version. This is not a supported version skew and may lead to a malfunctional cluster`

    * 下载低版本的 kubeadm 可能会导致一并下载的 kubelet, kubectl，cni 等版本出现不兼容的情况
    * 使用docker的情况下，建议使用  `yum install -y kubelet-1.18.5 kubeadm-1.18.5 kubectl-1.18.5 `  1.18.5 版本

  * `kubeadm init /proc/sys/net/bridge/bridge-nf-call-iptables contents are not set to 1`  需要开启网桥模式

    *  ```bash
       echo 1 > /proc/sys/net/bridge/bridge-nf-call-iptables
       echo 1 > /proc/sys/net/bridge/bridge-nf-call-ip6tables
       ```

  * 无法下载k8s镜像
    * ```bash
      # 指定镜像地址
      kubeadm init --image-repository registry.aliyuncs.com/google_containers 
      ```

    * 直接使用上述指令，下面操作过于繁琐
    * kubeadm config images pull  下载镜像失败；国内正常访问不到k8s.cgr.io，可以替换阿里加速镜像地址：registry.aliyuncs.com/google_containers，执行如下命令

      * ```bash
        docker pull registry.aliyuncs.com/google_containers/kube-apiserver:v1.18.20
        docker pull registry.aliyuncs.com/google_containers/kube-controller-manager:v1.18.20
        docker pull registry.aliyuncs.com/google_containers/kube-scheduler:v1.18.20
        docker pull registry.aliyuncs.com/google_containers/kube-proxy:v1.18.20
        docker pull registry.aliyuncs.com/google_containers/pause:3.2
        docker pull registry.aliyuncs.com/google_containers/etcd:3.4.3-0
        docker pull registry.aliyuncs.com/google_containers/coredns:1.6.7
        ```

    * 接下来给镜像重命名，使其和原kubeadm需要的镜像名称一致

      * ```bash
        docker tag registry.aliyuncs.com/google_containers/kube-apiserver:v1.18.20 k8s.gcr.io/kube-apiserver:v1.18.20
        docker tag registry.aliyuncs.com/google_containers/kube-controller-manager:v1.18.20 k8s.gcr.io/kube-controller-manager:v1.18.20
        docker tag registry.aliyuncs.com/google_containers/kube-scheduler:v1.18.20 k8s.gcr.io/kube-scheduler:v1.18.20
        docker tag registry.aliyuncs.com/google_containers/kube-proxy:v1.18.20 k8s.gcr.io/kube-proxy:v1.18.20
        docker tag registry.aliyuncs.com/google_containers/pause:3.2 k8s.gcr.io/pause:3.2
        docker tag registry.aliyuncs.com/google_containers/etcd:3.4.3-0 k8s.gcr.io/etcd:3.4.3-0
        docker tag registry.aliyuncs.com/google_containers/coredns:1.6.7 k8s.gcr.io/coredns:1.6.7
        
        ```

    * 再删除掉从阿里云下载的镜像

      * ```bash
        docker rmi registry.aliyuncs.com/google_containers/kube-apiserver:v1.18.20
        docker rmi registry.aliyuncs.com/google_containers/kube-controller-manager:v1.18.20
        docker rmi registry.aliyuncs.com/google_containers/kube-scheduler:v1.18.20
        docker rmi registry.aliyuncs.com/google_containers/kube-proxy:v1.18.20
        docker rmi registry.aliyuncs.com/google_containers/pause:3.2
        docker rmi registry.aliyuncs.com/google_containers/etcd:3.4.3-0
        docker rmi registry.aliyuncs.com/google_containers/coredns:1.6.7
        ```

    * 最后再执行初始化指令 kubeadm init 
    
      * ```
        Your Kubernetes control-plane has initialized successfully!
        
        To start using your cluster, you need to run the following as a regular user:
        
          mkdir -p $HOME/.kube
          sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
          sudo chown $(id -u):$(id -g) $HOME/.kube/config
        
        You should now deploy a pod network to the cluster.
        Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
          https://kubernetes.io/docs/concepts/cluster-administration/addons/
        
        ```
    
      * 按照要求执行指令，并安装网络插件（不安装网络插件，通过 `kubectl get nodes` 会发现唯一的节点处于`notReady` 状态， `kubectl describe node nodename` 查看节点状态 ）
    
      * 安装网络插件
    
        * weave: `kubectl apply -f https://github.com/weaveworks/weave/releases/download/v2.8.1/weave-daemonset-k8s.yaml`
        * flannel: ``kubectl apply -f https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml`
        * 链接文件建议保存下来，毕竟链接有可能失效
        * 之后重新查看 ``kubectl get nodes`` 节点处于ready 状态证明ok 
    
  * 主节点默认不允许运行用户pod
  
    * 依靠的是 Kubernetes 的 Taint/Toleration 机制（污点机制）
    * `kubectl taint nodes --all node-role.kubernetes.io/master-` 去除污点，可以通过`kubectl describe node nodename` 查看 `taints` 字段，去除后，该字段会置为 `<none>`
  
  * 安装可视化工具
  
    * `https://github.com/kubernetes/dashboard` 
    * [可视化工具](../platform/k8s-page-visual.md)
  
  * 部署容器存储插件
  
    * rook 项目 `https://rook.io/docs/rook/v1.6/ceph-quickstart.html `

