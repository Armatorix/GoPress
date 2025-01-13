package db

import (
	"context"

	"github.com/Armatorix/GoPress/be/db/ent/ent"
)

func (db *DB) GetArticles(ctx context.Context) (ent.Articles, error) {
	return db.ArticleClient().
		Query().
		All(ctx)
}
