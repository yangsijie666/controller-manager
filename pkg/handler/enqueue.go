package handler

import (
	"context"
	"github.com/yangsijie666/controller-manager/pkg/event"
	"github.com/yangsijie666/controller-manager/pkg/reconcile"
	"k8s.io/client-go/util/workqueue"
)

/**
* @Author: yangsijie666
* @Date: 2023/2/2 20:08
 */

//var enqueueLog = logf.Log.WithName("controller-runtime").WithName("eventhandler").WithName("EnqueueRequestForObject")

var _ EventHandler[string] = &EnqueueRequestForObject[string]{}

type EnqueueRequestForObject[T comparable] struct {
}

func (e *EnqueueRequestForObject[T]) Generic(_ context.Context, evt event.GenericEvent[T], q workqueue.RateLimitingInterface) {
	q.Add(reconcile.Request[T]{
		ID: evt.Object,
	})
}
