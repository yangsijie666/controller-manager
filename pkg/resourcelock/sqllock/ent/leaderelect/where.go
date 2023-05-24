// Code generated by ent, DO NOT EDIT.

package leaderelect

import (
	"entgo.io/ent/dialect/sql"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldContainsFold(FieldID, id))
}

// ResourceVersion applies equality check predicate on the "resource_version" field. It's identical to ResourceVersionEQ.
func ResourceVersion(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldResourceVersion, v))
}

// LeaseDurationSeconds applies equality check predicate on the "lease_duration_seconds" field. It's identical to LeaseDurationSecondsEQ.
func LeaseDurationSeconds(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldLeaseDurationSeconds, v))
}

// AcquireTime applies equality check predicate on the "acquire_time" field. It's identical to AcquireTimeEQ.
func AcquireTime(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldAcquireTime, v))
}

// RenewTime applies equality check predicate on the "renew_time" field. It's identical to RenewTimeEQ.
func RenewTime(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldRenewTime, v))
}

// LeaderTransitions applies equality check predicate on the "leader_transitions" field. It's identical to LeaderTransitionsEQ.
func LeaderTransitions(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldLeaderTransitions, v))
}

// HolderIdentity applies equality check predicate on the "holder_identity" field. It's identical to HolderIdentityEQ.
func HolderIdentity(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldHolderIdentity, v))
}

// ResourceVersionEQ applies the EQ predicate on the "resource_version" field.
func ResourceVersionEQ(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldResourceVersion, v))
}

// ResourceVersionNEQ applies the NEQ predicate on the "resource_version" field.
func ResourceVersionNEQ(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNEQ(FieldResourceVersion, v))
}

// ResourceVersionIn applies the In predicate on the "resource_version" field.
func ResourceVersionIn(vs ...string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldIn(FieldResourceVersion, vs...))
}

// ResourceVersionNotIn applies the NotIn predicate on the "resource_version" field.
func ResourceVersionNotIn(vs ...string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNotIn(FieldResourceVersion, vs...))
}

// ResourceVersionGT applies the GT predicate on the "resource_version" field.
func ResourceVersionGT(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGT(FieldResourceVersion, v))
}

// ResourceVersionGTE applies the GTE predicate on the "resource_version" field.
func ResourceVersionGTE(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGTE(FieldResourceVersion, v))
}

// ResourceVersionLT applies the LT predicate on the "resource_version" field.
func ResourceVersionLT(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLT(FieldResourceVersion, v))
}

// ResourceVersionLTE applies the LTE predicate on the "resource_version" field.
func ResourceVersionLTE(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLTE(FieldResourceVersion, v))
}

// ResourceVersionContains applies the Contains predicate on the "resource_version" field.
func ResourceVersionContains(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldContains(FieldResourceVersion, v))
}

// ResourceVersionHasPrefix applies the HasPrefix predicate on the "resource_version" field.
func ResourceVersionHasPrefix(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldHasPrefix(FieldResourceVersion, v))
}

// ResourceVersionHasSuffix applies the HasSuffix predicate on the "resource_version" field.
func ResourceVersionHasSuffix(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldHasSuffix(FieldResourceVersion, v))
}

// ResourceVersionEqualFold applies the EqualFold predicate on the "resource_version" field.
func ResourceVersionEqualFold(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEqualFold(FieldResourceVersion, v))
}

// ResourceVersionContainsFold applies the ContainsFold predicate on the "resource_version" field.
func ResourceVersionContainsFold(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldContainsFold(FieldResourceVersion, v))
}

// LeaseDurationSecondsEQ applies the EQ predicate on the "lease_duration_seconds" field.
func LeaseDurationSecondsEQ(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldLeaseDurationSeconds, v))
}

// LeaseDurationSecondsNEQ applies the NEQ predicate on the "lease_duration_seconds" field.
func LeaseDurationSecondsNEQ(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNEQ(FieldLeaseDurationSeconds, v))
}

// LeaseDurationSecondsIn applies the In predicate on the "lease_duration_seconds" field.
func LeaseDurationSecondsIn(vs ...int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldIn(FieldLeaseDurationSeconds, vs...))
}

// LeaseDurationSecondsNotIn applies the NotIn predicate on the "lease_duration_seconds" field.
func LeaseDurationSecondsNotIn(vs ...int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNotIn(FieldLeaseDurationSeconds, vs...))
}

// LeaseDurationSecondsGT applies the GT predicate on the "lease_duration_seconds" field.
func LeaseDurationSecondsGT(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGT(FieldLeaseDurationSeconds, v))
}

// LeaseDurationSecondsGTE applies the GTE predicate on the "lease_duration_seconds" field.
func LeaseDurationSecondsGTE(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGTE(FieldLeaseDurationSeconds, v))
}

// LeaseDurationSecondsLT applies the LT predicate on the "lease_duration_seconds" field.
func LeaseDurationSecondsLT(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLT(FieldLeaseDurationSeconds, v))
}

// LeaseDurationSecondsLTE applies the LTE predicate on the "lease_duration_seconds" field.
func LeaseDurationSecondsLTE(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLTE(FieldLeaseDurationSeconds, v))
}

// AcquireTimeEQ applies the EQ predicate on the "acquire_time" field.
func AcquireTimeEQ(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldAcquireTime, v))
}

// AcquireTimeNEQ applies the NEQ predicate on the "acquire_time" field.
func AcquireTimeNEQ(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNEQ(FieldAcquireTime, v))
}

// AcquireTimeIn applies the In predicate on the "acquire_time" field.
func AcquireTimeIn(vs ...int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldIn(FieldAcquireTime, vs...))
}

// AcquireTimeNotIn applies the NotIn predicate on the "acquire_time" field.
func AcquireTimeNotIn(vs ...int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNotIn(FieldAcquireTime, vs...))
}

// AcquireTimeGT applies the GT predicate on the "acquire_time" field.
func AcquireTimeGT(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGT(FieldAcquireTime, v))
}

// AcquireTimeGTE applies the GTE predicate on the "acquire_time" field.
func AcquireTimeGTE(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGTE(FieldAcquireTime, v))
}

// AcquireTimeLT applies the LT predicate on the "acquire_time" field.
func AcquireTimeLT(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLT(FieldAcquireTime, v))
}

// AcquireTimeLTE applies the LTE predicate on the "acquire_time" field.
func AcquireTimeLTE(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLTE(FieldAcquireTime, v))
}

// RenewTimeEQ applies the EQ predicate on the "renew_time" field.
func RenewTimeEQ(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldRenewTime, v))
}

// RenewTimeNEQ applies the NEQ predicate on the "renew_time" field.
func RenewTimeNEQ(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNEQ(FieldRenewTime, v))
}

// RenewTimeIn applies the In predicate on the "renew_time" field.
func RenewTimeIn(vs ...int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldIn(FieldRenewTime, vs...))
}

// RenewTimeNotIn applies the NotIn predicate on the "renew_time" field.
func RenewTimeNotIn(vs ...int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNotIn(FieldRenewTime, vs...))
}

// RenewTimeGT applies the GT predicate on the "renew_time" field.
func RenewTimeGT(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGT(FieldRenewTime, v))
}

// RenewTimeGTE applies the GTE predicate on the "renew_time" field.
func RenewTimeGTE(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGTE(FieldRenewTime, v))
}

// RenewTimeLT applies the LT predicate on the "renew_time" field.
func RenewTimeLT(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLT(FieldRenewTime, v))
}

// RenewTimeLTE applies the LTE predicate on the "renew_time" field.
func RenewTimeLTE(v int64) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLTE(FieldRenewTime, v))
}

// LeaderTransitionsEQ applies the EQ predicate on the "leader_transitions" field.
func LeaderTransitionsEQ(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldLeaderTransitions, v))
}

// LeaderTransitionsNEQ applies the NEQ predicate on the "leader_transitions" field.
func LeaderTransitionsNEQ(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNEQ(FieldLeaderTransitions, v))
}

// LeaderTransitionsIn applies the In predicate on the "leader_transitions" field.
func LeaderTransitionsIn(vs ...int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldIn(FieldLeaderTransitions, vs...))
}

// LeaderTransitionsNotIn applies the NotIn predicate on the "leader_transitions" field.
func LeaderTransitionsNotIn(vs ...int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNotIn(FieldLeaderTransitions, vs...))
}

// LeaderTransitionsGT applies the GT predicate on the "leader_transitions" field.
func LeaderTransitionsGT(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGT(FieldLeaderTransitions, v))
}

// LeaderTransitionsGTE applies the GTE predicate on the "leader_transitions" field.
func LeaderTransitionsGTE(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGTE(FieldLeaderTransitions, v))
}

// LeaderTransitionsLT applies the LT predicate on the "leader_transitions" field.
func LeaderTransitionsLT(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLT(FieldLeaderTransitions, v))
}

// LeaderTransitionsLTE applies the LTE predicate on the "leader_transitions" field.
func LeaderTransitionsLTE(v int) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLTE(FieldLeaderTransitions, v))
}

// HolderIdentityEQ applies the EQ predicate on the "holder_identity" field.
func HolderIdentityEQ(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEQ(FieldHolderIdentity, v))
}

// HolderIdentityNEQ applies the NEQ predicate on the "holder_identity" field.
func HolderIdentityNEQ(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNEQ(FieldHolderIdentity, v))
}

// HolderIdentityIn applies the In predicate on the "holder_identity" field.
func HolderIdentityIn(vs ...string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldIn(FieldHolderIdentity, vs...))
}

// HolderIdentityNotIn applies the NotIn predicate on the "holder_identity" field.
func HolderIdentityNotIn(vs ...string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldNotIn(FieldHolderIdentity, vs...))
}

// HolderIdentityGT applies the GT predicate on the "holder_identity" field.
func HolderIdentityGT(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGT(FieldHolderIdentity, v))
}

// HolderIdentityGTE applies the GTE predicate on the "holder_identity" field.
func HolderIdentityGTE(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldGTE(FieldHolderIdentity, v))
}

// HolderIdentityLT applies the LT predicate on the "holder_identity" field.
func HolderIdentityLT(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLT(FieldHolderIdentity, v))
}

// HolderIdentityLTE applies the LTE predicate on the "holder_identity" field.
func HolderIdentityLTE(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldLTE(FieldHolderIdentity, v))
}

// HolderIdentityContains applies the Contains predicate on the "holder_identity" field.
func HolderIdentityContains(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldContains(FieldHolderIdentity, v))
}

// HolderIdentityHasPrefix applies the HasPrefix predicate on the "holder_identity" field.
func HolderIdentityHasPrefix(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldHasPrefix(FieldHolderIdentity, v))
}

// HolderIdentityHasSuffix applies the HasSuffix predicate on the "holder_identity" field.
func HolderIdentityHasSuffix(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldHasSuffix(FieldHolderIdentity, v))
}

// HolderIdentityEqualFold applies the EqualFold predicate on the "holder_identity" field.
func HolderIdentityEqualFold(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldEqualFold(FieldHolderIdentity, v))
}

// HolderIdentityContainsFold applies the ContainsFold predicate on the "holder_identity" field.
func HolderIdentityContainsFold(v string) predicate.LeaderElect {
	return predicate.LeaderElect(sql.FieldContainsFold(FieldHolderIdentity, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LeaderElect) predicate.LeaderElect {
	return predicate.LeaderElect(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LeaderElect) predicate.LeaderElect {
	return predicate.LeaderElect(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LeaderElect) predicate.LeaderElect {
	return predicate.LeaderElect(func(s *sql.Selector) {
		p(s.Not())
	})
}
