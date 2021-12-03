package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Int("creator_id"),
		field.String("title"),
		field.String("description").Optional(),
		field.String("priority").Default("5"),
		field.String("complexity").Default("5"),
		field.Time("hard_deadline").Optional(),
		field.Time("soft_deadline").Optional(),
		field.String("status").Default("123"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).
			Ref("tasks").
			Field("creator_id").Required().
			Unique(),
	}
}
