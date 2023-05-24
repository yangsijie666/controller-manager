// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent/leaderelect"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	leaderelectFields := schema.LeaderElect{}.Fields()
	_ = leaderelectFields
	// leaderelectDescResourceVersion is the schema descriptor for resource_version field.
	leaderelectDescResourceVersion := leaderelectFields[1].Descriptor()
	// leaderelect.DefaultResourceVersion holds the default value on creation for the resource_version field.
	leaderelect.DefaultResourceVersion = leaderelectDescResourceVersion.Default.(string)
	// leaderelectDescID is the schema descriptor for id field.
	leaderelectDescID := leaderelectFields[0].Descriptor()
	// leaderelect.IDValidator is a validator for the "id" field. It is called by the builders before save.
	leaderelect.IDValidator = leaderelectDescID.Validators[0].(func(string) error)
}