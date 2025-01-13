// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Armatorix/GoPress/be/db/ent/ent/article"
	"github.com/Armatorix/GoPress/be/db/ent/ent/predicate"
	"github.com/Armatorix/GoPress/be/db/ent/ent/user"
)

// ArticleUpdate is the builder for updating Article entities.
type ArticleUpdate struct {
	config
	hooks     []Hook
	mutation  *ArticleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ArticleUpdate builder.
func (au *ArticleUpdate) Where(ps ...predicate.Article) *ArticleUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ArticleUpdate) SetUpdatedAt(t time.Time) *ArticleUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableUpdatedAt(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// SetTitle sets the "title" field.
func (au *ArticleUpdate) SetTitle(s string) *ArticleUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableTitle(s *string) *ArticleUpdate {
	if s != nil {
		au.SetTitle(*s)
	}
	return au
}

// ClearTitle clears the value of the "title" field.
func (au *ArticleUpdate) ClearTitle() *ArticleUpdate {
	au.mutation.ClearTitle()
	return au
}

// SetDescription sets the "description" field.
func (au *ArticleUpdate) SetDescription(s string) *ArticleUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableDescription(s *string) *ArticleUpdate {
	if s != nil {
		au.SetDescription(*s)
	}
	return au
}

// ClearDescription clears the value of the "description" field.
func (au *ArticleUpdate) ClearDescription() *ArticleUpdate {
	au.mutation.ClearDescription()
	return au
}

// SetBody sets the "body" field.
func (au *ArticleUpdate) SetBody(s string) *ArticleUpdate {
	au.mutation.SetBody(s)
	return au
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableBody(s *string) *ArticleUpdate {
	if s != nil {
		au.SetBody(*s)
	}
	return au
}

// SetReleased sets the "released" field.
func (au *ArticleUpdate) SetReleased(b bool) *ArticleUpdate {
	au.mutation.SetReleased(b)
	return au
}

// SetNillableReleased sets the "released" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableReleased(b *bool) *ArticleUpdate {
	if b != nil {
		au.SetReleased(*b)
	}
	return au
}

// SetAuthorID sets the "author_id" field.
func (au *ArticleUpdate) SetAuthorID(s string) *ArticleUpdate {
	au.mutation.SetAuthorID(s)
	return au
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableAuthorID(s *string) *ArticleUpdate {
	if s != nil {
		au.SetAuthorID(*s)
	}
	return au
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (au *ArticleUpdate) AddUserIDs(ids ...int) *ArticleUpdate {
	au.mutation.AddUserIDs(ids...)
	return au
}

// AddUser adds the "user" edges to the User entity.
func (au *ArticleUpdate) AddUser(u ...*User) *ArticleUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return au.AddUserIDs(ids...)
}

// Mutation returns the ArticleMutation object of the builder.
func (au *ArticleUpdate) Mutation() *ArticleMutation {
	return au.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (au *ArticleUpdate) ClearUser() *ArticleUpdate {
	au.mutation.ClearUser()
	return au
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (au *ArticleUpdate) RemoveUserIDs(ids ...int) *ArticleUpdate {
	au.mutation.RemoveUserIDs(ids...)
	return au
}

// RemoveUser removes "user" edges to User entities.
func (au *ArticleUpdate) RemoveUser(u ...*User) *ArticleUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return au.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArticleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArticleUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArticleUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *ArticleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ArticleUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *ArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(article.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
	}
	if au.mutation.TitleCleared() {
		_spec.ClearField(article.FieldTitle, field.TypeString)
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.SetField(article.FieldDescription, field.TypeString, value)
	}
	if au.mutation.DescriptionCleared() {
		_spec.ClearField(article.FieldDescription, field.TypeString)
	}
	if value, ok := au.mutation.Body(); ok {
		_spec.SetField(article.FieldBody, field.TypeString, value)
	}
	if value, ok := au.mutation.Released(); ok {
		_spec.SetField(article.FieldReleased, field.TypeBool, value)
	}
	if value, ok := au.mutation.AuthorID(); ok {
		_spec.SetField(article.FieldAuthorID, field.TypeString, value)
	}
	if au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   article.UserTable,
			Columns: article.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedUserIDs(); len(nodes) > 0 && !au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   article.UserTable,
			Columns: article.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   article.UserTable,
			Columns: article.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArticleUpdateOne is the builder for updating a single Article entity.
type ArticleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ArticleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ArticleUpdateOne) SetUpdatedAt(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableUpdatedAt(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// SetTitle sets the "title" field.
func (auo *ArticleUpdateOne) SetTitle(s string) *ArticleUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableTitle(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetTitle(*s)
	}
	return auo
}

// ClearTitle clears the value of the "title" field.
func (auo *ArticleUpdateOne) ClearTitle() *ArticleUpdateOne {
	auo.mutation.ClearTitle()
	return auo
}

// SetDescription sets the "description" field.
func (auo *ArticleUpdateOne) SetDescription(s string) *ArticleUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableDescription(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetDescription(*s)
	}
	return auo
}

// ClearDescription clears the value of the "description" field.
func (auo *ArticleUpdateOne) ClearDescription() *ArticleUpdateOne {
	auo.mutation.ClearDescription()
	return auo
}

// SetBody sets the "body" field.
func (auo *ArticleUpdateOne) SetBody(s string) *ArticleUpdateOne {
	auo.mutation.SetBody(s)
	return auo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableBody(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetBody(*s)
	}
	return auo
}

// SetReleased sets the "released" field.
func (auo *ArticleUpdateOne) SetReleased(b bool) *ArticleUpdateOne {
	auo.mutation.SetReleased(b)
	return auo
}

// SetNillableReleased sets the "released" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableReleased(b *bool) *ArticleUpdateOne {
	if b != nil {
		auo.SetReleased(*b)
	}
	return auo
}

// SetAuthorID sets the "author_id" field.
func (auo *ArticleUpdateOne) SetAuthorID(s string) *ArticleUpdateOne {
	auo.mutation.SetAuthorID(s)
	return auo
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableAuthorID(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetAuthorID(*s)
	}
	return auo
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (auo *ArticleUpdateOne) AddUserIDs(ids ...int) *ArticleUpdateOne {
	auo.mutation.AddUserIDs(ids...)
	return auo
}

// AddUser adds the "user" edges to the User entity.
func (auo *ArticleUpdateOne) AddUser(u ...*User) *ArticleUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return auo.AddUserIDs(ids...)
}

// Mutation returns the ArticleMutation object of the builder.
func (auo *ArticleUpdateOne) Mutation() *ArticleMutation {
	return auo.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (auo *ArticleUpdateOne) ClearUser() *ArticleUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (auo *ArticleUpdateOne) RemoveUserIDs(ids ...int) *ArticleUpdateOne {
	auo.mutation.RemoveUserIDs(ids...)
	return auo
}

// RemoveUser removes "user" edges to User entities.
func (auo *ArticleUpdateOne) RemoveUser(u ...*User) *ArticleUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return auo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the ArticleUpdate builder.
func (auo *ArticleUpdateOne) Where(ps ...predicate.Article) *ArticleUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArticleUpdateOne) Select(field string, fields ...string) *ArticleUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Article entity.
func (auo *ArticleUpdateOne) Save(ctx context.Context) (*Article, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArticleUpdateOne) SaveX(ctx context.Context) *Article {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArticleUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *ArticleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ArticleUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *ArticleUpdateOne) sqlSave(ctx context.Context) (_node *Article, err error) {
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Article.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for _, f := range fields {
			if !article.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != article.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(article.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
	}
	if auo.mutation.TitleCleared() {
		_spec.ClearField(article.FieldTitle, field.TypeString)
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.SetField(article.FieldDescription, field.TypeString, value)
	}
	if auo.mutation.DescriptionCleared() {
		_spec.ClearField(article.FieldDescription, field.TypeString)
	}
	if value, ok := auo.mutation.Body(); ok {
		_spec.SetField(article.FieldBody, field.TypeString, value)
	}
	if value, ok := auo.mutation.Released(); ok {
		_spec.SetField(article.FieldReleased, field.TypeBool, value)
	}
	if value, ok := auo.mutation.AuthorID(); ok {
		_spec.SetField(article.FieldAuthorID, field.TypeString, value)
	}
	if auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   article.UserTable,
			Columns: article.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedUserIDs(); len(nodes) > 0 && !auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   article.UserTable,
			Columns: article.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   article.UserTable,
			Columns: article.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Article{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
