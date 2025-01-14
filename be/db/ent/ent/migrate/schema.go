// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "body", Type: field.TypeString, Default: ""},
		{Name: "released", Type: field.TypeBool, Default: false},
		{Name: "author_id", Type: field.TypeInt},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "nick", Type: field.TypeString, Default: ""},
		{Name: "avatar_url", Type: field.TypeString, Default: ""},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// ArticleUserColumns holds the columns for the "article_user" table.
	ArticleUserColumns = []*schema.Column{
		{Name: "article_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ArticleUserTable holds the schema information for the "article_user" table.
	ArticleUserTable = &schema.Table{
		Name:       "article_user",
		Columns:    ArticleUserColumns,
		PrimaryKey: []*schema.Column{ArticleUserColumns[0], ArticleUserColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "article_user_article_id",
				Columns:    []*schema.Column{ArticleUserColumns[0]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "article_user_user_id",
				Columns:    []*schema.Column{ArticleUserColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		UsersTable,
		ArticleUserTable,
	}
)

func init() {
	ArticleUserTable.ForeignKeys[0].RefTable = ArticlesTable
	ArticleUserTable.ForeignKeys[1].RefTable = UsersTable
}
