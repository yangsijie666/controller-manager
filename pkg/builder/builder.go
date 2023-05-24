package builder

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/yangsijie666/controller-manager/pkg/controller"
	"github.com/yangsijie666/controller-manager/pkg/handler"
	"github.com/yangsijie666/controller-manager/pkg/manager"
	"github.com/yangsijie666/controller-manager/pkg/predicate"
	"github.com/yangsijie666/controller-manager/pkg/reconcile"
	"github.com/yangsijie666/controller-manager/pkg/source"
)

/**
* @Author: yangsijie666
* @Date: 2023/2/2 17:04
 */

// Builder builds a Controller.
type Builder[T comparable] struct {
	src              source.Source[T]
	eventHandler     handler.EventHandler[T]
	mgr              manager.Manager
	globalPredicates []predicate.Predicate[T]
	ctrl             controller.Interface[T]
	ctrlOptions      controller.Options[T]
	name             string
}

func ControllerManagedBy[T comparable](m manager.Manager) *Builder[T] {
	return &Builder[T]{mgr: m}
}

// WithEventFilter sets the event filters, to filter which create/update/delete/generic events eventually
// trigger reconciliations.  For example, filtering on whether the resource version has changed.
// Given predicate is added for all watched objects.
// Defaults to the empty list.
func (blder *Builder[T]) WithEventFilter(p predicate.Predicate[T]) *Builder[T] {
	blder.globalPredicates = append(blder.globalPredicates, p)
	return blder
}

func (blder *Builder[T]) WithSource(src source.Source[T]) *Builder[T] {
	blder.src = src
	return blder
}

func (blder *Builder[T]) WithEventHandler(eventHandler handler.EventHandler[T]) *Builder[T] {
	blder.eventHandler = eventHandler
	return blder
}

// WithOptions overrides the controller options use in doController. Defaults to empty.
func (blder *Builder[T]) WithOptions(options controller.Options[T]) *Builder[T] {
	blder.ctrlOptions = options
	return blder
}

// WithLogConstructor overrides the controller options's LogConstructor.
func (blder *Builder[T]) WithLogConstructor(logConstructor func(*reconcile.Request[T]) logr.Logger) *Builder[T] {
	blder.ctrlOptions.LogConstructor = logConstructor
	return blder
}

// Named sets the name of the controller to the given name.  The name shows up
// in metrics, among other things, and thus should be a prometheus compatible name
// (underscores and alphanumeric characters only).
//
// By default, controllers are named using the lowercase version of their kind.
func (blder *Builder[T]) Named(name string) *Builder[T] {
	blder.name = name
	return blder
}

// Complete builds the Application Controller.
func (blder *Builder[T]) Complete(r reconcile.Reconciler[T]) error {
	_, err := blder.Build(r)
	return err
}

func (blder *Builder[T]) doController(r reconcile.Reconciler[T]) (err error) {
	ctrlOptions := blder.ctrlOptions
	if ctrlOptions.Reconciler == nil {
		ctrlOptions.Reconciler = r
	}

	// Setup concurrency.
	if ctrlOptions.MaxConcurrentReconciles == 0 {
		ctrlOptions.MaxConcurrentReconciles = 1
	}

	if blder.name == "" {
		return errors.New("controllerName not set")
	}
	controllerName := blder.name

	// Setup the logger.
	if ctrlOptions.LogConstructor == nil {
		log := blder.mgr.GetLogger().WithValues(
			"controller", controllerName,
		)

		ctrlOptions.LogConstructor = func(req *reconcile.Request[T]) logr.Logger {
			log := log
			if req != nil {
				log = log.WithValues(
					"objectID", req.ObjectID(),
				)
			}
			return log
		}
	}

	// Build the controller and return.
	blder.ctrl, err = controller.New(controllerName, blder.mgr, ctrlOptions)
	return err
}

func (blder *Builder[T]) doWatch() error {
	// Reconcile type
	if blder.src == nil {
		return errors.New("must provide a source")
	}
	if blder.eventHandler == nil {
		return errors.New("must provide a eventHandler")
	}
	if err := blder.ctrl.Watch(blder.src, blder.eventHandler, blder.globalPredicates...); err != nil {
		return err
	}
	return nil
}

// Build builds the Application Controller and returns the Controller it created.
func (blder *Builder[T]) Build(r reconcile.Reconciler[T]) (controller.Interface[T], error) {
	if r == nil {
		return nil, errors.New("must provide a non-nil Reconciler")
	}
	if blder.mgr == nil {
		return nil, errors.Errorf("must provide a non-nil Manager")
	}

	// Set the ControllerManagedBy
	if err := blder.doController(r); err != nil {
		return nil, err
	}

	// Set the Watch
	if err := blder.doWatch(); err != nil {
		return nil, err
	}

	return blder.ctrl, nil
}
