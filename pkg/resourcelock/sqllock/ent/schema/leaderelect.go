package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

/**
* @Author: yangsijie666
* @Date: 2023/1/30 17:52
 */

type LeaderElect struct {
	ent.Schema
}

// Fields of the Ak.
func (LeaderElect) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("name").
			NotEmpty().
			Unique(),
		field.String("resource_version").Default("0").Comment("版本号"),
		field.Int("lease_duration_seconds"),
		field.Int64("acquire_time"),
		field.Int64("renew_time"),
		field.Int("leader_transitions"),
		field.String("holder_identity"),
	}
}

func (LeaderElect) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "leader_election",
		},
	}
}
