// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Armatorix/GoPress/be/db/ent/ent/article"
	"github.com/Armatorix/GoPress/be/db/ext"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description ext.Password `json:"description,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// AuthorID holds the value of the "author_id" field.
	AuthorID string `json:"author_id,omitempty"`
	// EmailConfirmationSecret holds the value of the "email_confirmation_secret" field.
	EmailConfirmationSecret ext.Password `json:"email_confirmation_secret,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArticleQuery when eager-loading is set.
	Edges        ArticleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ArticleEdges holds the relations/edges for other nodes in the graph.
type ArticleEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	namedUser   map[string][]*User
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e ArticleEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldDescription, article.FieldEmailConfirmationSecret:
			values[i] = new(ext.Password)
		case article.FieldID:
			values[i] = new(sql.NullInt64)
		case article.FieldTitle, article.FieldBody, article.FieldAuthorID:
			values[i] = new(sql.NullString)
		case article.FieldCreatedAt, article.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case article.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case article.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldDescription:
			if value, ok := values[i].(*ext.Password); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value != nil {
				a.Description = *value
			}
		case article.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				a.Body = value.String
			}
		case article.FieldAuthorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_id", values[i])
			} else if value.Valid {
				a.AuthorID = value.String
			}
		case article.FieldEmailConfirmationSecret:
			if value, ok := values[i].(*ext.Password); !ok {
				return fmt.Errorf("unexpected type %T for field email_confirmation_secret", values[i])
			} else if value != nil {
				a.EmailConfirmationSecret = *value
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Article.
// This includes values selected through modifiers, order, etc.
func (a *Article) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Article entity.
func (a *Article) QueryUser() *UserQuery {
	return NewArticleClient(a.config).QueryUser(a)
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return NewArticleClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(fmt.Sprintf("%v", a.Description))
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(a.Body)
	builder.WriteString(", ")
	builder.WriteString("author_id=")
	builder.WriteString(a.AuthorID)
	builder.WriteString(", ")
	builder.WriteString("email_confirmation_secret=")
	builder.WriteString(fmt.Sprintf("%v", a.EmailConfirmationSecret))
	builder.WriteByte(')')
	return builder.String()
}

// NamedUser returns the User named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Article) NamedUser(name string) ([]*User, error) {
	if a.Edges.namedUser == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedUser[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Article) appendNamedUser(name string, edges ...*User) {
	if a.Edges.namedUser == nil {
		a.Edges.namedUser = make(map[string][]*User)
	}
	if len(edges) == 0 {
		a.Edges.namedUser[name] = []*User{}
	} else {
		a.Edges.namedUser[name] = append(a.Edges.namedUser[name], edges...)
	}
}

// Articles is a parsable slice of Article.
type Articles []*Article
