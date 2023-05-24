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

/**
* @Author: yangsijie666
* @Date: 2023/5/24 14:26
 */

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
