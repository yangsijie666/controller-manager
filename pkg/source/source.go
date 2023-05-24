package source

import (
	"context"
	"github.com/yangsijie666/controller-manager/pkg/handler"
	"github.com/yangsijie666/controller-manager/pkg/predicate"
	"k8s.io/client-go/util/workqueue"
)

/**
* @Author: yangsijie666
* @Date: 2023/2/2 13:55
 */

type Source[T interface{}] interface {
	// Start is internal and should be called only by the Controller to register an EventHandler with the Informer
	// to enqueue reconcile.Requests.
	Start(context.Context, handler.EventHandler[T], workqueue.RateLimitingInterface, ...predicate.Predicate[T]) error
}
