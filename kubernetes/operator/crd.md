## CRD  资源自定义
#### 学习资料
kubebuilder
* https://book-v1.book.kubebuilder.io/
* https://book.kubebuilder.io/

kubernetes
* https://kubernetes.io/zh-cn/docs/home/
<br/>

#### 前置概念
**[CRD](https://kubernetes.io/zh-cn/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/)**

**[Controller](https://kubernetes.io/zh-cn/docs/concepts/architecture/controller/)**

**[Operator](https://kubernetes.io/zh-cn/docs/concepts/extend-kubernetes/operator/)**
* Operator = CRD + Controller
<br/>

#### 使用kubebuilder生成代码
**[Quick-Start](https://book.kubebuilder.io/quick-start.html)**
<br/>

#### 定义
CRD的定义推荐使用 **[kubebuilder](https://book.kubebuilder.io/introduction)**，k8s社区主流使用的代码生成脚手架工具，可以通过定义的types.go生成CRD的yaml和控制器框架代码。所以这里主要给出使用kubebuilder工具时的一些要求。

控制器模型中要求对象具备两个字段，一个为 Spec，一个为 Status。控制器的职责为不断进行 Reconcile，达成 Spec 与 Status 的一致。
<br/>

###### 建议：
在 kubernetes 中，但你需要同时更新 Spec 和 Status 时，要记得，这是两个操作（update 和 updateStatus），并不是直接更新一个 CR 就可以完成，这两个操作是无法做到原子性的（即同时成功/失败）

因此在设计 CRD 的时候，尽可能让自己的业务逻辑每次仅修改 Spec/Status，通过不断触发事件的 Reconcil 去达成 Spec（期待状态）与 Status（当前状态）的一致
<br/>

###### 备注：
```
"sigs.k8s.io/cluster-api/util/patch"
这个工具是在代码层面将 Spec 和 Status 的更新操作合并了，底层操作还是分两步更新


# kubectl 更新指令
kubectl edit xxx
kubectl edit xxx --subresource=status
```
<br/>

#### 字段命名
字段在 types.go 中命名符合 Go 语言规范，同时要求生成的 CRD 中字段取值与 OpenAPI 要求一致。
[example](https://book-v1.book.kubebuilder.io/basics/simple_resource)

```
// ContainerSet creates a new Deployment running multiple replicas of a single container with the given
// image.
// +k8s:openapi-gen=true
// +resource:path=containersets
type ContainerSet struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    // spec contains the desired behavior of the ContainerSet
    Spec   ContainerSetSpec   `json:"spec,omitempty"`

    // status contains the last observed state of the ContainerSet
    Status ContainerSetStatus `json:"status,omitempty"`
}

// ContainerSetSpec defines the desired state of ContainerSet
type ContainerSetSpec struct {
  // replicas is the number of replicas to maintain
  Replicas *int32 `json:"replicas,omitempty"`

  // image is the container image to run.  Image must have a tag.
  // +kubebuilder:validation:Pattern=.+:.+
  Image string `json:"image,omitempty"`
}

// ContainerSetStatus defines the observed state of ContainerSet
type ContainerSetStatus struct {
  HealthyReplicas *int32 `json:"healthyReplicas,omitempty"`
}
```
<br/>

#### 字段校验
CRD直接对UI时，CRD作为对外openAPI的一部分，也要考虑精确性。使用kubebuilder中特性可以让kube-apiserver帮助检查CR取值的合法性。

具体参考：https://book.kubebuilder.io/reference/markers/crd-validation
1.25之后，社区也支持更灵活的检验规则，具体可见：https://kubernetes.io/blog/2022/09/23/crd-validation-rules-beta/
<br/>

#### 文件名
项目中有多个CRD时，每个CRD使用 name_types.go风格进行命名。
<br/>

#### 资源客户端生成
不同的资源需要不同的客户端去对接操作，`client-go` 提供了多种原生的客户端与内置资源供使用者调用，从方便的角度上看，还是建议 CRD 资源生成自己的资源客户端来操作 kubernetes 资源

客户端推荐使用 code-generator 工具生成。由于 code-generator 工具对目录结构要求与 kubebuilder 不同，所以推荐以下的api目录布局方式。

code-generator 中定义的 types.go 中不定义具体的 Spec 字段，而是直接引用 kubebuilder 使用的目录。这样不会有重复的 Spec 类型定义，一旦 CRD 定义变化，不用手动更新 code-generator 目录下的内容。

```
.
├── example1                                // code-generator 使用目录
│   └── v1beta1
│       ├── doc.go
│       ├── register.go
│       ├── types.go
│       └── zz_generated.deepcopy.go
├── example2                              // code-generator 使用目录
│   └── v1beta1
│       ├── doc.go
│       ├── register.go
│       ├── types.go
│       └── zz_generated.deepcopy.go
└── v1beta1                                 // kubebuilder 使用目录
    ├── groupversion_info.go
    ├── example1_types.go
    ├── example2_types.go
    └── zz_generated.deepcopy.go
```
<br/>

#### 展示有意义的字段
kubebuilder 默认生成的CRD，通过 kubectl get CR 进行查看时，看不到什么有用的信息。它支持通过printcolumn添加要输出的字段。

示例：
```
// +kubebuilder:printcolumn:name="ID",type="string",JSONPath=".status.id"
```

添加这个声明生成的 CRD，通过 `kubectl get XX` 展示资源列表时，会将 ID 字段展示出来，方便开发和排查问题，避免需要 `kubectl get xx -o yaml`

强制要求在types.go中添加这一注释，为CRD展示有意义的字段，方便排查问题。
<br/>

#### tips
##### 更新
kubebuilder对应的controller-runtime提供了两种更新CR的接口，r.Client.Update和r.Client.Patch。
区别如下

| 操作   | 原理                                                         | 使用场景                                                     |
| ------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Update | 客户端提交整个对象，kube-apiserver比较resource version是否冲突，冲突时拒绝请求，即有乐观锁保护。 | 对并发安全要求高。<br/><br/>多个客户端可能并发修改一对象时，Update可以起到保护作用。 |
| Patch  | 客户端提交一个Patch，kube-apiserver根据Patch策略直接修改对象。 | 对性能要求高。<br/><br/>不会有多客户端并发修改。<br/>有并发修改场景，但不会并发修改同一个字段，这样也是安全的。 |

kubernetes patch 策略：https://kubernetes.io/zh-cn/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/
