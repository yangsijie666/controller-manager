package handler

import (
	"context"
	"github.com/yangsijie666/controller-manager/pkg/event"
	"k8s.io/client-go/util/workqueue"
)

/**
* @Author: yangsijie666
* @Date: 2023/2/2 13:59
 */

type EventHandler[T interface{}] interface {
	// Generic is called in response to an event of an unknown type or a synthetic event triggered as a cron or
	// external trigger request - e.g. reconcile Autoscaling, or a Webhook.
	Generic(context.Context, event.GenericEvent[T], workqueue.RateLimitingInterface)
}
