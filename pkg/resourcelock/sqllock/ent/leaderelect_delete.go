// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent/leaderelect"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent/predicate"
)

// LeaderElectDelete is the builder for deleting a LeaderElect entity.
type LeaderElectDelete struct {
	config
	hooks    []Hook
	mutation *LeaderElectMutation
}

// Where appends a list predicates to the LeaderElectDelete builder.
func (led *LeaderElectDelete) Where(ps ...predicate.LeaderElect) *LeaderElectDelete {
	led.mutation.Where(ps...)
	return led
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (led *LeaderElectDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, led.sqlExec, led.mutation, led.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (led *LeaderElectDelete) ExecX(ctx context.Context) int {
	n, err := led.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (led *LeaderElectDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(leaderelect.Table, sqlgraph.NewFieldSpec(leaderelect.FieldID, field.TypeString))
	if ps := led.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, led.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	led.mutation.done = true
	return affected, err
}

// LeaderElectDeleteOne is the builder for deleting a single LeaderElect entity.
type LeaderElectDeleteOne struct {
	led *LeaderElectDelete
}

// Where appends a list predicates to the LeaderElectDelete builder.
func (ledo *LeaderElectDeleteOne) Where(ps ...predicate.LeaderElect) *LeaderElectDeleteOne {
	ledo.led.mutation.Where(ps...)
	return ledo
}

// Exec executes the deletion query.
func (ledo *LeaderElectDeleteOne) Exec(ctx context.Context) error {
	n, err := ledo.led.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{leaderelect.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ledo *LeaderElectDeleteOne) ExecX(ctx context.Context) {
	if err := ledo.Exec(ctx); err != nil {
		panic(err)
	}
}
