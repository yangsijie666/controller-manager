package reconcile

import (
	"context"
	"fmt"
	"time"
)

/**
* @Author: yangsijie666
* @Date: 2023/2/1 16:36
 */

type Result struct {
	// Requeue tells the Controller to requeue the reconcile key.  Defaults to false.
	Requeue bool

	// RequeueAfter if greater than 0, tells the Controller to requeue the reconcile key after the Duration.
	// Implies that Requeue is true, there is no need to set Requeue to true at the same time as RequeueAfter.
	RequeueAfter time.Duration
}

type Request[T comparable] struct {
	ID T
}

func (r *Request[T]) ObjectID() string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("%v", r.ID)
}

// IsZero returns true if this result is empty.
func (r *Result) IsZero() bool {
	if r == nil {
		return true
	}
	return *r == Result{}
}

type Reconciler[T comparable] interface {
	// Reconcile performs a full reconciliation for the object referred to by the Request.
	// The Controller will requeue the Request to be processed again if an error is non-nil or
	// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
	Reconcile(context.Context, Request[T]) (Result, error)
}
