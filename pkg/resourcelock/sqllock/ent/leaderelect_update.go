// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent/leaderelect"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent/predicate"
)

// LeaderElectUpdate is the builder for updating LeaderElect entities.
type LeaderElectUpdate struct {
	config
	hooks    []Hook
	mutation *LeaderElectMutation
}

// Where appends a list predicates to the LeaderElectUpdate builder.
func (leu *LeaderElectUpdate) Where(ps ...predicate.LeaderElect) *LeaderElectUpdate {
	leu.mutation.Where(ps...)
	return leu
}

// SetResourceVersion sets the "resource_version" field.
func (leu *LeaderElectUpdate) SetResourceVersion(s string) *LeaderElectUpdate {
	leu.mutation.SetResourceVersion(s)
	return leu
}

// SetNillableResourceVersion sets the "resource_version" field if the given value is not nil.
func (leu *LeaderElectUpdate) SetNillableResourceVersion(s *string) *LeaderElectUpdate {
	if s != nil {
		leu.SetResourceVersion(*s)
	}
	return leu
}

// SetLeaseDurationSeconds sets the "lease_duration_seconds" field.
func (leu *LeaderElectUpdate) SetLeaseDurationSeconds(i int) *LeaderElectUpdate {
	leu.mutation.ResetLeaseDurationSeconds()
	leu.mutation.SetLeaseDurationSeconds(i)
	return leu
}

// AddLeaseDurationSeconds adds i to the "lease_duration_seconds" field.
func (leu *LeaderElectUpdate) AddLeaseDurationSeconds(i int) *LeaderElectUpdate {
	leu.mutation.AddLeaseDurationSeconds(i)
	return leu
}

// SetAcquireTime sets the "acquire_time" field.
func (leu *LeaderElectUpdate) SetAcquireTime(i int64) *LeaderElectUpdate {
	leu.mutation.ResetAcquireTime()
	leu.mutation.SetAcquireTime(i)
	return leu
}

// AddAcquireTime adds i to the "acquire_time" field.
func (leu *LeaderElectUpdate) AddAcquireTime(i int64) *LeaderElectUpdate {
	leu.mutation.AddAcquireTime(i)
	return leu
}

// SetRenewTime sets the "renew_time" field.
func (leu *LeaderElectUpdate) SetRenewTime(i int64) *LeaderElectUpdate {
	leu.mutation.ResetRenewTime()
	leu.mutation.SetRenewTime(i)
	return leu
}

// AddRenewTime adds i to the "renew_time" field.
func (leu *LeaderElectUpdate) AddRenewTime(i int64) *LeaderElectUpdate {
	leu.mutation.AddRenewTime(i)
	return leu
}

// SetLeaderTransitions sets the "leader_transitions" field.
func (leu *LeaderElectUpdate) SetLeaderTransitions(i int) *LeaderElectUpdate {
	leu.mutation.ResetLeaderTransitions()
	leu.mutation.SetLeaderTransitions(i)
	return leu
}

// AddLeaderTransitions adds i to the "leader_transitions" field.
func (leu *LeaderElectUpdate) AddLeaderTransitions(i int) *LeaderElectUpdate {
	leu.mutation.AddLeaderTransitions(i)
	return leu
}

// SetHolderIdentity sets the "holder_identity" field.
func (leu *LeaderElectUpdate) SetHolderIdentity(s string) *LeaderElectUpdate {
	leu.mutation.SetHolderIdentity(s)
	return leu
}

// Mutation returns the LeaderElectMutation object of the builder.
func (leu *LeaderElectUpdate) Mutation() *LeaderElectMutation {
	return leu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (leu *LeaderElectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, leu.sqlSave, leu.mutation, leu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (leu *LeaderElectUpdate) SaveX(ctx context.Context) int {
	affected, err := leu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (leu *LeaderElectUpdate) Exec(ctx context.Context) error {
	_, err := leu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (leu *LeaderElectUpdate) ExecX(ctx context.Context) {
	if err := leu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (leu *LeaderElectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(leaderelect.Table, leaderelect.Columns, sqlgraph.NewFieldSpec(leaderelect.FieldID, field.TypeString))
	if ps := leu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := leu.mutation.ResourceVersion(); ok {
		_spec.SetField(leaderelect.FieldResourceVersion, field.TypeString, value)
	}
	if value, ok := leu.mutation.LeaseDurationSeconds(); ok {
		_spec.SetField(leaderelect.FieldLeaseDurationSeconds, field.TypeInt, value)
	}
	if value, ok := leu.mutation.AddedLeaseDurationSeconds(); ok {
		_spec.AddField(leaderelect.FieldLeaseDurationSeconds, field.TypeInt, value)
	}
	if value, ok := leu.mutation.AcquireTime(); ok {
		_spec.SetField(leaderelect.FieldAcquireTime, field.TypeInt64, value)
	}
	if value, ok := leu.mutation.AddedAcquireTime(); ok {
		_spec.AddField(leaderelect.FieldAcquireTime, field.TypeInt64, value)
	}
	if value, ok := leu.mutation.RenewTime(); ok {
		_spec.SetField(leaderelect.FieldRenewTime, field.TypeInt64, value)
	}
	if value, ok := leu.mutation.AddedRenewTime(); ok {
		_spec.AddField(leaderelect.FieldRenewTime, field.TypeInt64, value)
	}
	if value, ok := leu.mutation.LeaderTransitions(); ok {
		_spec.SetField(leaderelect.FieldLeaderTransitions, field.TypeInt, value)
	}
	if value, ok := leu.mutation.AddedLeaderTransitions(); ok {
		_spec.AddField(leaderelect.FieldLeaderTransitions, field.TypeInt, value)
	}
	if value, ok := leu.mutation.HolderIdentity(); ok {
		_spec.SetField(leaderelect.FieldHolderIdentity, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, leu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{leaderelect.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	leu.mutation.done = true
	return n, nil
}

// LeaderElectUpdateOne is the builder for updating a single LeaderElect entity.
type LeaderElectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LeaderElectMutation
}

// SetResourceVersion sets the "resource_version" field.
func (leuo *LeaderElectUpdateOne) SetResourceVersion(s string) *LeaderElectUpdateOne {
	leuo.mutation.SetResourceVersion(s)
	return leuo
}

// SetNillableResourceVersion sets the "resource_version" field if the given value is not nil.
func (leuo *LeaderElectUpdateOne) SetNillableResourceVersion(s *string) *LeaderElectUpdateOne {
	if s != nil {
		leuo.SetResourceVersion(*s)
	}
	return leuo
}

// SetLeaseDurationSeconds sets the "lease_duration_seconds" field.
func (leuo *LeaderElectUpdateOne) SetLeaseDurationSeconds(i int) *LeaderElectUpdateOne {
	leuo.mutation.ResetLeaseDurationSeconds()
	leuo.mutation.SetLeaseDurationSeconds(i)
	return leuo
}

// AddLeaseDurationSeconds adds i to the "lease_duration_seconds" field.
func (leuo *LeaderElectUpdateOne) AddLeaseDurationSeconds(i int) *LeaderElectUpdateOne {
	leuo.mutation.AddLeaseDurationSeconds(i)
	return leuo
}

// SetAcquireTime sets the "acquire_time" field.
func (leuo *LeaderElectUpdateOne) SetAcquireTime(i int64) *LeaderElectUpdateOne {
	leuo.mutation.ResetAcquireTime()
	leuo.mutation.SetAcquireTime(i)
	return leuo
}

// AddAcquireTime adds i to the "acquire_time" field.
func (leuo *LeaderElectUpdateOne) AddAcquireTime(i int64) *LeaderElectUpdateOne {
	leuo.mutation.AddAcquireTime(i)
	return leuo
}

// SetRenewTime sets the "renew_time" field.
func (leuo *LeaderElectUpdateOne) SetRenewTime(i int64) *LeaderElectUpdateOne {
	leuo.mutation.ResetRenewTime()
	leuo.mutation.SetRenewTime(i)
	return leuo
}

// AddRenewTime adds i to the "renew_time" field.
func (leuo *LeaderElectUpdateOne) AddRenewTime(i int64) *LeaderElectUpdateOne {
	leuo.mutation.AddRenewTime(i)
	return leuo
}

// SetLeaderTransitions sets the "leader_transitions" field.
func (leuo *LeaderElectUpdateOne) SetLeaderTransitions(i int) *LeaderElectUpdateOne {
	leuo.mutation.ResetLeaderTransitions()
	leuo.mutation.SetLeaderTransitions(i)
	return leuo
}

// AddLeaderTransitions adds i to the "leader_transitions" field.
func (leuo *LeaderElectUpdateOne) AddLeaderTransitions(i int) *LeaderElectUpdateOne {
	leuo.mutation.AddLeaderTransitions(i)
	return leuo
}

// SetHolderIdentity sets the "holder_identity" field.
func (leuo *LeaderElectUpdateOne) SetHolderIdentity(s string) *LeaderElectUpdateOne {
	leuo.mutation.SetHolderIdentity(s)
	return leuo
}

// Mutation returns the LeaderElectMutation object of the builder.
func (leuo *LeaderElectUpdateOne) Mutation() *LeaderElectMutation {
	return leuo.mutation
}

// Where appends a list predicates to the LeaderElectUpdate builder.
func (leuo *LeaderElectUpdateOne) Where(ps ...predicate.LeaderElect) *LeaderElectUpdateOne {
	leuo.mutation.Where(ps...)
	return leuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (leuo *LeaderElectUpdateOne) Select(field string, fields ...string) *LeaderElectUpdateOne {
	leuo.fields = append([]string{field}, fields...)
	return leuo
}

// Save executes the query and returns the updated LeaderElect entity.
func (leuo *LeaderElectUpdateOne) Save(ctx context.Context) (*LeaderElect, error) {
	return withHooks(ctx, leuo.sqlSave, leuo.mutation, leuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (leuo *LeaderElectUpdateOne) SaveX(ctx context.Context) *LeaderElect {
	node, err := leuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (leuo *LeaderElectUpdateOne) Exec(ctx context.Context) error {
	_, err := leuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (leuo *LeaderElectUpdateOne) ExecX(ctx context.Context) {
	if err := leuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (leuo *LeaderElectUpdateOne) sqlSave(ctx context.Context) (_node *LeaderElect, err error) {
	_spec := sqlgraph.NewUpdateSpec(leaderelect.Table, leaderelect.Columns, sqlgraph.NewFieldSpec(leaderelect.FieldID, field.TypeString))
	id, ok := leuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LeaderElect.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := leuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, leaderelect.FieldID)
		for _, f := range fields {
			if !leaderelect.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != leaderelect.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := leuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := leuo.mutation.ResourceVersion(); ok {
		_spec.SetField(leaderelect.FieldResourceVersion, field.TypeString, value)
	}
	if value, ok := leuo.mutation.LeaseDurationSeconds(); ok {
		_spec.SetField(leaderelect.FieldLeaseDurationSeconds, field.TypeInt, value)
	}
	if value, ok := leuo.mutation.AddedLeaseDurationSeconds(); ok {
		_spec.AddField(leaderelect.FieldLeaseDurationSeconds, field.TypeInt, value)
	}
	if value, ok := leuo.mutation.AcquireTime(); ok {
		_spec.SetField(leaderelect.FieldAcquireTime, field.TypeInt64, value)
	}
	if value, ok := leuo.mutation.AddedAcquireTime(); ok {
		_spec.AddField(leaderelect.FieldAcquireTime, field.TypeInt64, value)
	}
	if value, ok := leuo.mutation.RenewTime(); ok {
		_spec.SetField(leaderelect.FieldRenewTime, field.TypeInt64, value)
	}
	if value, ok := leuo.mutation.AddedRenewTime(); ok {
		_spec.AddField(leaderelect.FieldRenewTime, field.TypeInt64, value)
	}
	if value, ok := leuo.mutation.LeaderTransitions(); ok {
		_spec.SetField(leaderelect.FieldLeaderTransitions, field.TypeInt, value)
	}
	if value, ok := leuo.mutation.AddedLeaderTransitions(); ok {
		_spec.AddField(leaderelect.FieldLeaderTransitions, field.TypeInt, value)
	}
	if value, ok := leuo.mutation.HolderIdentity(); ok {
		_spec.SetField(leaderelect.FieldHolderIdentity, field.TypeString, value)
	}
	_node = &LeaderElect{config: leuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, leuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{leaderelect.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	leuo.mutation.done = true
	return _node, nil
}
