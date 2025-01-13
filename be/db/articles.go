package db

import (
	"context"

	"github.com/Armatorix/GoPress/be/db/ent/ent"
	"github.com/Armatorix/GoPress/be/db/ent/ent/article"
)

type InsertArticle struct {
	Body  string
	Title string
}

func (db *DB) GetArticles(ctx context.Context) (ent.Articles, error) {
	return db.ArticleClient().
		Query().
		All(ctx)
}

func (db *DB) InsertArticle(ctx context.Context, v InsertArticle) error {
	return db.ArticleClient().
		Create().
		SetBody(v.Body).
		SetTitle(v.Title).
		Exec(ctx)
}

func (db *DB) UpdateArticle(ctx context.Context, id int, v InsertArticle) error {
	return db.ArticleClient().
		UpdateOneID(id).
		SetBody(v.Body).
		SetTitle(v.Title).
		Exec(ctx)
}

func (db *DB) DeleteArticles(ctx context.Context, ids ...int) error {
	_, err := db.ArticleClient().
		Delete().
		Where(article.IDIn(ids...)).
		Exec(ctx)
	return err
}
