package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/Armatorix/GoPress/be/db/ext"
)

type Article struct {
	ent.Schema
}

func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Immutable(),
		field.String("title").
			Optional(),
		field.String("description").
			GoType(ext.Password("")).
			Optional(),
		field.String("body").
			Default(""),
		field.String("author_id"),
		field.String("email_confirmation_secret").
			GoType(ext.Password("")).
			Optional(),
	}
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
	}
}
