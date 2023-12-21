### Controller-runtime

#### 控制器并发数设置

并发数用来配置控制器最多同步处理不同CR数量，即在一个控制器服务进程中，同一时间一个对象只会在一个协程中处理

同一个对象的多次调谐是串行执行，不同对象是并发执行

对于短时间出现的多个事件，会进行优化处理，并不是每次事件都会触发调谐，可能是中间态的某个事件，或者最后一个事件触发调谐

* 将某个`CR`的状态按一定时间间隔（10s）进行修改，每次调谐耗时1m

  * 状态修改：Deploying --> DeployFailed --> Updating --> UpdateFailed --> Deleting

  * 实际触发了调谐只有 Deploying，Deleting 这两个状态的事件 

    

```golang
import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

// SetupWithManager sets up the controller with the Manager.
func (r *PluginInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&example.ExampleReconciler{}).
   	    WithOptions(controller.Options{MaxConcurrentReconciles: 100}).
		Complete(r)
}


通过 MaxConcurrentReconciles 开启并发数
```



#### 开启选举

deployment部署的控制器（一般都是这种部署形式）使用`leader election`控制真正工作的实例数量。

原因

* deployment升级时有双实例场景，防止事件重复消费
* deployment业务处在多实例场景

```golang
import (
		ctrl "sigs.k8s.io/controller-runtime"
		)

mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		LeaderElection:                true
})
```



#### 备注

更多用法项目参考：https://github.com/kubernetes-sigs/cluster-api

