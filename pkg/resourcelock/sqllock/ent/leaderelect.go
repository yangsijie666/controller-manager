// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent/leaderelect"
)

// LeaderElect is the model entity for the LeaderElect schema.
type LeaderElect struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 版本号
	ResourceVersion string `json:"resource_version,omitempty"`
	// LeaseDurationSeconds holds the value of the "lease_duration_seconds" field.
	LeaseDurationSeconds int `json:"lease_duration_seconds,omitempty"`
	// AcquireTime holds the value of the "acquire_time" field.
	AcquireTime int64 `json:"acquire_time,omitempty"`
	// RenewTime holds the value of the "renew_time" field.
	RenewTime int64 `json:"renew_time,omitempty"`
	// LeaderTransitions holds the value of the "leader_transitions" field.
	LeaderTransitions int `json:"leader_transitions,omitempty"`
	// HolderIdentity holds the value of the "holder_identity" field.
	HolderIdentity string `json:"holder_identity,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LeaderElect) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case leaderelect.FieldLeaseDurationSeconds, leaderelect.FieldAcquireTime, leaderelect.FieldRenewTime, leaderelect.FieldLeaderTransitions:
			values[i] = new(sql.NullInt64)
		case leaderelect.FieldID, leaderelect.FieldResourceVersion, leaderelect.FieldHolderIdentity:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LeaderElect fields.
func (le *LeaderElect) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case leaderelect.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				le.ID = value.String
			}
		case leaderelect.FieldResourceVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_version", values[i])
			} else if value.Valid {
				le.ResourceVersion = value.String
			}
		case leaderelect.FieldLeaseDurationSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lease_duration_seconds", values[i])
			} else if value.Valid {
				le.LeaseDurationSeconds = int(value.Int64)
			}
		case leaderelect.FieldAcquireTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field acquire_time", values[i])
			} else if value.Valid {
				le.AcquireTime = value.Int64
			}
		case leaderelect.FieldRenewTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field renew_time", values[i])
			} else if value.Valid {
				le.RenewTime = value.Int64
			}
		case leaderelect.FieldLeaderTransitions:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field leader_transitions", values[i])
			} else if value.Valid {
				le.LeaderTransitions = int(value.Int64)
			}
		case leaderelect.FieldHolderIdentity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field holder_identity", values[i])
			} else if value.Valid {
				le.HolderIdentity = value.String
			}
		default:
			le.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LeaderElect.
// This includes values selected through modifiers, order, etc.
func (le *LeaderElect) Value(name string) (ent.Value, error) {
	return le.selectValues.Get(name)
}

// Update returns a builder for updating this LeaderElect.
// Note that you need to call LeaderElect.Unwrap() before calling this method if this LeaderElect
// was returned from a transaction, and the transaction was committed or rolled back.
func (le *LeaderElect) Update() *LeaderElectUpdateOne {
	return NewLeaderElectClient(le.config).UpdateOne(le)
}

// Unwrap unwraps the LeaderElect entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (le *LeaderElect) Unwrap() *LeaderElect {
	_tx, ok := le.config.driver.(*txDriver)
	if !ok {
		panic("ent: LeaderElect is not a transactional entity")
	}
	le.config.driver = _tx.drv
	return le
}

// String implements the fmt.Stringer.
func (le *LeaderElect) String() string {
	var builder strings.Builder
	builder.WriteString("LeaderElect(")
	builder.WriteString(fmt.Sprintf("id=%v, ", le.ID))
	builder.WriteString("resource_version=")
	builder.WriteString(le.ResourceVersion)
	builder.WriteString(", ")
	builder.WriteString("lease_duration_seconds=")
	builder.WriteString(fmt.Sprintf("%v", le.LeaseDurationSeconds))
	builder.WriteString(", ")
	builder.WriteString("acquire_time=")
	builder.WriteString(fmt.Sprintf("%v", le.AcquireTime))
	builder.WriteString(", ")
	builder.WriteString("renew_time=")
	builder.WriteString(fmt.Sprintf("%v", le.RenewTime))
	builder.WriteString(", ")
	builder.WriteString("leader_transitions=")
	builder.WriteString(fmt.Sprintf("%v", le.LeaderTransitions))
	builder.WriteString(", ")
	builder.WriteString("holder_identity=")
	builder.WriteString(le.HolderIdentity)
	builder.WriteByte(')')
	return builder.String()
}

// LeaderElects is a parsable slice of LeaderElect.
type LeaderElects []*LeaderElect
