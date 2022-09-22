package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Email holds the schema definition for the Email entity.
type Email struct {
	ent.Schema
}

// Fields of the Email.
func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.String("from"),
		field.String("subject"),
		field.String("body"),
		field.Time("sentAt"),
		field.Bool("archived").Default(false),
		field.Bool("read").Default(false),
	}
}

// Edges of the Email.
func (Email) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "user" of type `User`
		// and reference it to the "emails" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("user", User.Type).
			Ref("emails").
			// setting the edge to unique, ensure
			// that an email can have only one owner.
			Unique(),
	}
}
