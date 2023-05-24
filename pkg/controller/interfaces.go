package controller

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yangsijie666/controller-manager/pkg/config"
	"github.com/yangsijie666/controller-manager/pkg/controller/metrics"
	"github.com/yangsijie666/controller-manager/pkg/handler"
	"github.com/yangsijie666/controller-manager/pkg/manager"
	"github.com/yangsijie666/controller-manager/pkg/predicate"
	"github.com/yangsijie666/controller-manager/pkg/ratelimiter"
	"github.com/yangsijie666/controller-manager/pkg/reconcile"
	"github.com/yangsijie666/controller-manager/pkg/source"
	"k8s.io/client-go/util/workqueue"
)

/**
* @Author: yangsijie666
* @Date: 2023/2/1 16:05
 */

type Interface[T comparable] interface {
	reconcile.Reconciler[T]

	// Watch takes events provided by a Source and uses the EventHandler to
	// enqueue reconcile.Requests in response to the events.
	//
	// Watch may be provided one or more Predicates to filter events before
	// they are given to the EventHandler.  Events will be passed to the
	// EventHandler if all provided Predicates evaluate to true.
	Watch(src source.Source[T], eventHandler handler.EventHandler[T], predicates ...predicate.Predicate[T]) error

	// Start starts the controller.  Start blocks until the context is closed or a
	// controller has an error starting.
	Start(ctx context.Context) error

	// GetLogger returns this controller logger prefilled with basic information.
	GetLogger() logr.Logger
}

// Options are the arguments for creating a new Controller.
type Options[T comparable] struct {
	config.Controller

	// Reconciler reconciles an object
	Reconciler reconcile.Reconciler[T]

	// RateLimiter is used to limit how frequently requests may be queued.
	// Defaults to MaxOfRateLimiter which has both overall and per-item rate limiting.
	// The overall is a token bucket and the per-item is exponential.
	RateLimiter ratelimiter.RateLimiter

	// LogConstructor is used to construct a logger used for this controller and passed
	// to each reconciliation via the context field.
	LogConstructor func(request *reconcile.Request[T]) logr.Logger

	MetricsRegistry *prometheus.Registry
}

// New returns a new Controller registered with the Manager.  The Manager will ensure that shared Caches have
// been synced before the Controller is Started.
func New[T comparable](name string, mgr manager.Manager, options Options[T]) (Interface[T], error) {
	c, err := NewUnmanaged(name, mgr, options)
	if err != nil {
		return nil, err
	}

	// Add the controller as a Manager components
	return c, mgr.Add(c)
}

// NewUnmanaged returns a new controller without adding it to the manager. The
// caller is responsible for starting the returned controller.
func NewUnmanaged[T comparable](name string, mgr manager.Manager, options Options[T]) (Interface[T], error) {
	if options.Reconciler == nil {
		return nil, errors.Errorf("must specify Reconciler")
	}

	if len(name) == 0 {
		return nil, errors.Errorf("must specify Name for Controller")
	}

	if options.LogConstructor == nil {
		log := mgr.GetLogger().WithValues(
			"controller", name,
		)
		options.LogConstructor = func(req *reconcile.Request[T]) logr.Logger {
			log := log
			if req != nil {
				log = log.WithValues(
					"objectID", req.ObjectID(),
				)
			}
			return log
		}
	}

	if options.MaxConcurrentReconciles <= 0 {
		options.MaxConcurrentReconciles = 1
	}

	if options.RateLimiter == nil {
		options.RateLimiter = workqueue.DefaultControllerRateLimiter()
	}

	if options.RecoverPanic == nil {
		options.RecoverPanic = mgr.GetControllerOptions().RecoverPanic
	}

	if options.MetricsRegistry == nil {
		options.MetricsRegistry = prometheus.NewRegistry()
	}
	metrics.MustRegister(options.MetricsRegistry)

	// Create controller with dependencies set
	return &Controller[T]{
		Do: options.Reconciler,
		MakeQueue: func() workqueue.RateLimitingInterface {
			return workqueue.NewRateLimitingQueueWithConfig(options.RateLimiter, workqueue.RateLimitingQueueConfig{
				Name: name,
			})
		},
		MaxConcurrentReconciles: options.MaxConcurrentReconciles,
		Name:                    name,
		LogConstructor:          options.LogConstructor,
		RecoverPanic:            options.RecoverPanic,
		LeaderElected:           options.NeedLeaderElection,
	}, nil
}
