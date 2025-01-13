// Code generated by ent, DO NOT EDIT.

package article

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Armatorix/GoPress/be/db/ent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldUpdatedAt, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldDescription, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldBody, v))
}

// AuthorID applies equality check predicate on the "author_id" field. It's identical to AuthorIDEQ.
func AuthorID(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldAuthorID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Article {
	return predicate.Article(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Article {
	return predicate.Article(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Article {
	return predicate.Article(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Article {
	return predicate.Article(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldDescription, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldBody, v))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldBody, v))
}

// AuthorIDEQ applies the EQ predicate on the "author_id" field.
func AuthorIDEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldAuthorID, v))
}

// AuthorIDNEQ applies the NEQ predicate on the "author_id" field.
func AuthorIDNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldAuthorID, v))
}

// AuthorIDIn applies the In predicate on the "author_id" field.
func AuthorIDIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldAuthorID, vs...))
}

// AuthorIDNotIn applies the NotIn predicate on the "author_id" field.
func AuthorIDNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldAuthorID, vs...))
}

// AuthorIDGT applies the GT predicate on the "author_id" field.
func AuthorIDGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldAuthorID, v))
}

// AuthorIDGTE applies the GTE predicate on the "author_id" field.
func AuthorIDGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldAuthorID, v))
}

// AuthorIDLT applies the LT predicate on the "author_id" field.
func AuthorIDLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldAuthorID, v))
}

// AuthorIDLTE applies the LTE predicate on the "author_id" field.
func AuthorIDLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldAuthorID, v))
}

// AuthorIDContains applies the Contains predicate on the "author_id" field.
func AuthorIDContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldAuthorID, v))
}

// AuthorIDHasPrefix applies the HasPrefix predicate on the "author_id" field.
func AuthorIDHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldAuthorID, v))
}

// AuthorIDHasSuffix applies the HasSuffix predicate on the "author_id" field.
func AuthorIDHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldAuthorID, v))
}

// AuthorIDEqualFold applies the EqualFold predicate on the "author_id" field.
func AuthorIDEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldAuthorID, v))
}

// AuthorIDContainsFold applies the ContainsFold predicate on the "author_id" field.
func AuthorIDContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldAuthorID, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Article) predicate.Article {
	return predicate.Article(sql.NotPredicates(p))
}
