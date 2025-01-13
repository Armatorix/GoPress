package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/Armatorix/GoPress/be/db/ext"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Immutable(),
		field.String("email"),
		field.String("password").
			GoType(ext.Password("")).
			Optional(),
		field.String("nick").
			Default(""),
		field.String("avatar_url").
			Default(""),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("articles", Article.Type).
			Ref("user"),
	}
}
