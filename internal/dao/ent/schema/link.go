package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("short").Optional().Unique().MaxLen(32),
		field.String("origin"),
		field.String("comment").Default(""),
		field.Time("start_time").Optional().Nillable(),
		field.Time("end_time").Optional().Nillable(),
		field.Bool("active").Default(true),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("links").Unique(),
	}
}
