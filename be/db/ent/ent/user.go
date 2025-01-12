// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Armatorix/GoPress/be/db/ent/ent/user"
	"github.com/Armatorix/GoPress/be/db/ext"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password ext.Password `json:"password,omitempty"`
	// Nick holds the value of the "nick" field.
	Nick string `json:"nick,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`
	// EmailConfirmationSecret holds the value of the "email_confirmation_secret" field.
	EmailConfirmationSecret ext.Password `json:"email_confirmation_secret,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Articles holds the value of the articles edge.
	Articles []*Article `json:"articles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes   [1]bool
	namedArticles map[string][]*Article
}

// ArticlesOrErr returns the Articles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ArticlesOrErr() ([]*Article, error) {
	if e.loadedTypes[0] {
		return e.Articles, nil
	}
	return nil, &NotLoadedError{edge: "articles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldPassword, user.FieldEmailConfirmationSecret:
			values[i] = new(ext.Password)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldEmail, user.FieldNick, user.FieldAvatarURL:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*ext.Password); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value != nil {
				u.Password = *value
			}
		case user.FieldNick:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nick", values[i])
			} else if value.Valid {
				u.Nick = value.String
			}
		case user.FieldAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_url", values[i])
			} else if value.Valid {
				u.AvatarURL = value.String
			}
		case user.FieldEmailConfirmationSecret:
			if value, ok := values[i].(*ext.Password); !ok {
				return fmt.Errorf("unexpected type %T for field email_confirmation_secret", values[i])
			} else if value != nil {
				u.EmailConfirmationSecret = *value
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryArticles queries the "articles" edge of the User entity.
func (u *User) QueryArticles() *ArticleQuery {
	return NewUserClient(u.config).QueryArticles(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(fmt.Sprintf("%v", u.Password))
	builder.WriteString(", ")
	builder.WriteString("nick=")
	builder.WriteString(u.Nick)
	builder.WriteString(", ")
	builder.WriteString("avatar_url=")
	builder.WriteString(u.AvatarURL)
	builder.WriteString(", ")
	builder.WriteString("email_confirmation_secret=")
	builder.WriteString(fmt.Sprintf("%v", u.EmailConfirmationSecret))
	builder.WriteByte(')')
	return builder.String()
}

// NamedArticles returns the Articles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedArticles(name string) ([]*Article, error) {
	if u.Edges.namedArticles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedArticles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedArticles(name string, edges ...*Article) {
	if u.Edges.namedArticles == nil {
		u.Edges.namedArticles = make(map[string][]*Article)
	}
	if len(edges) == 0 {
		u.Edges.namedArticles[name] = []*Article{}
	} else {
		u.Edges.namedArticles[name] = append(u.Edges.namedArticles[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
