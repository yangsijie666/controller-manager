# controller-manager Project

🤖 controller-manager 受 Kubernetes controller-runtime 项目启发，将 controller-manager 剥离出来，便于通用应用快速接入 leader 选举模式。

## 特性

- 😹 支持多态
- 🎃 支持通过数据库选举
- 🌈 支持在非 Kubernetes 场景使用 controller-manager

## 安装

```shell
go get github.com/yangsijie666/controller-manager
```

## 示例

定义一个自定义 Controller，并托管于 Manager，该 Controller 每隔 `5s` 生成一个随机数并触发 Controller 的 `Reconcile` 进行处理。

以下函数创建 Manager 和 Controller，将 Controller 托管于 Manager 后，启动 Manager。

```go
package main

import (
    "context"
    "github.com/go-logr/zapr"
    "github.com/yangsijie666/controller-manager/pkg/manager"
    "github.com/yangsijie666/controller-manager/pkg/resourcelock"
    "github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock"
    "github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent"
    "go.uber.org/zap"
    "os"
    "sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

var (
    identity  string
    entClient *ent.Client
)

func init() {
    var err error

    identity, err = os.Hostname()
    if err != nil {
        panic(err)
    }

    // init ent client
    entClient, err = ent.Open(
        "sqlite3",
        "file:ent.db?mode=memory&cache=shared&_fk=1",
    )
    if err != nil {
        panic(err)
    }
    if err = entClient.Schema.Create(context.TODO()); err != nil {
        panic(err)
    }
}

func main() {
    logger, _ := zap.NewDevelopment()
    entrylog := zapr.NewLogger(logger)

    // setup manager
    entrylog.Info("setup manager...")
    mgr, err := manager.New(manager.Options{
        Logger: entrylog,
        LeaderElectionResourceLockInterface: sqllock.NewSqlLock(
            "example-controller",
            entClient,
            resourcelock.ResourceLockConfig{
                Identity:      identity,
                EventRecorder: entrylog,
            },
        ),
    })
    if err != nil {
        entrylog.Error(err, "setup manager error")
        os.Exit(1)
    }

    // setup controller
    entrylog.Info("setup controller...")
    err = NewController(mgr)
    if err != nil {
        entrylog.Error(err, "setup controller error")
        os.Exit(1)
    }

    // starting
    entrylog.Info("starting manager...")
    if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
        return
    }
}
```

以下函数定义 Source，每隔 `5s` 生成随机数，并推送至队列中，由 Manager 收到后触发 Controller 执行 `Reconcile`。

```go
package main

import (
    "context"
    "github.com/yangsijie666/controller-manager/pkg/event"
    "github.com/yangsijie666/controller-manager/pkg/handler"
    "github.com/yangsijie666/controller-manager/pkg/predicate"
    "github.com/yangsijie666/controller-manager/pkg/source"
    "k8s.io/client-go/util/workqueue"
    "math/rand"
    "time"
)

var _ source.Source[int] = &Source{}

type Source struct {
}

func (s *Source) Start(ctx context.Context, evtHdr handler.EventHandler[int], q workqueue.RateLimitingInterface, _ ...predicate.Predicate[int]) error {
    go func(ctx context.Context) {
        t := time.NewTicker(time.Second * 5)
        defer t.Stop()

        for {
            s.fetchPendingObjects(ctx, evtHdr, q)

            select {
            case <-ctx.Done():
                return
            case <-t.C:
                continue
            }
        }
    }(ctx)
    return nil
}

func (s *Source) fetchPendingObjects(ctx context.Context, hdr handler.EventHandler[int], q workqueue.RateLimitingInterface) {
    defer func() {
        if err := recover(); err != nil {
            panic(err)
        }
    }()

    rand.Seed(time.Now().UnixNano())
    hdr.Generic(ctx, event.GenericEvent[int]{
        Object: rand.Intn(100),
    }, q)
}
```

以下函数定义 Controller，在 `Reconcile` 中输出触发对象。

```go
package main

import (
    "context"
    "github.com/go-logr/logr"
    ctrlblder "github.com/yangsijie666/controller-manager/pkg/builder"
    ctrlcfg "github.com/yangsijie666/controller-manager/pkg/config"
    "github.com/yangsijie666/controller-manager/pkg/controller"
    "github.com/yangsijie666/controller-manager/pkg/handler"
    "github.com/yangsijie666/controller-manager/pkg/manager"
    "github.com/yangsijie666/controller-manager/pkg/reconcile"
    "github.com/yangsijie666/controller-manager/pkg/util"
    "k8s.io/utils/pointer"
)

const (
    Name = "ExampleController"
)

func NewController(
    mgr manager.Manager,
) error {
    reconciler := &ExampleReconciler{
        log: mgr.GetLogger(),
    }
    return add(mgr, reconciler)
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler[int]) error {
    // Create a new controller
    return ctrlblder.ControllerManagedBy[int](mgr).
        Named(Name).
        WithSource(&Source{}).
        WithEventHandler(&handler.EnqueueRequestForObject[int]{}).
        WithOptions(controller.Options[int]{
            Controller: ctrlcfg.Controller{
                MaxConcurrentReconciles: 1,
                NeedLeaderElection:      pointer.Bool(true),
            },
        }).
        Complete(r)
}

var _ reconcile.Reconciler[int] = &ExampleReconciler{}

type ExampleReconciler struct {
    log logr.Logger
}

func (e *ExampleReconciler) Reconcile(ctx context.Context, r reconcile.Request[int]) (reconcile.Result, error) {
    e.log.Info("begin reconcile", "object", r.ID, "reconcileID", util.ExtractReconcileID(ctx))
    return reconcile.Result{}, nil
}
```

## 如何贡献代码

非常感激任何的代码提交。创建 pull request 时请遵守以下规则。

1. Fork 仓库。
2. 创建自己的特性分支。
3. 提交变更。
4. Push 分支。
5. 创建新的 pull request。
