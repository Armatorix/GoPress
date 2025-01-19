package db

import (
	"context"

	"github.com/Armatorix/GoPress/be/db/ent/ent"
	"github.com/Armatorix/GoPress/be/db/ent/ent/article"
)

type InsertArticle struct {
	Body        string
	Title       string
	Description string
	Released    bool
	UserId      int
}

type ArticleStats struct {
	TotalReleased int
	Total         int
}

func (db *DB) GetArticles(ctx context.Context) (ent.Articles, error) {
	return db.ArticleClient().
		Query().
		Order(ent.Desc(article.FieldID)).
		All(ctx)
}

func (db *DB) InsertArticle(ctx context.Context, v InsertArticle) error {
	return db.ArticleClient().
		Create().
		SetBody(v.Body).
		SetTitle(v.Title).
		SetDescription(v.Description).
		SetReleased(v.Released).
		SetAuthorID(v.UserId).
		Exec(ctx)
}

func (db *DB) UpdateArticle(ctx context.Context, id int, v InsertArticle) error {
	return db.ArticleClient().
		UpdateOneID(id).
		SetBody(v.Body).
		SetTitle(v.Title).
		SetDescription(v.Description).
		SetReleased(v.Released).
		SetAuthorID(v.UserId).
		Exec(ctx)
}

func (db *DB) DeleteArticles(ctx context.Context, ids ...int) error {
	_, err := db.ArticleClient().
		Delete().
		Where(article.IDIn(ids...)).
		Exec(ctx)
	return err
}

func (db *DB) PublishArticle(ctx context.Context, id int) error {
	article, err := db.ArticleClient().Get(ctx, id)
	if err != nil {
		return err
	}
	return db.ArticleClient().
		UpdateOneID(id).
		SetReleased(!article.Released).
		Exec(ctx)
}

// GetPublishedArticlesWithAuthors
func (db *DB) GetPublishedArticlesWithAuthors(ctx context.Context) (ent.Articles, error) {
	return db.ArticleClient().
		Query().
		Where(article.Released(true)).
		WithUser().
		Order(ent.Desc(article.FieldID)).
		All(ctx)
}

func (db *DB) GetPublishedArticleWithAuthor(ctx context.Context, id int) (*ent.Article, error) {
	return db.ArticleClient().
		Query().
		Where(
			article.ID(id),
			article.Released(true),
		).
		WithUser().
		Only(ctx)
}

func (db *DB) GetArticleStats(ctx context.Context) (*ArticleStats, error) {
	total, err := db.ArticleClient().
		Query().
		Count(ctx)
	if err != nil {
		return nil, err
	}
	totalReleased, err := db.ArticleClient().
		Query().
		Where(
			article.Released(true),
		).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	return &ArticleStats{
		Total:         total,
		TotalReleased: totalReleased,
	}, nil
}
