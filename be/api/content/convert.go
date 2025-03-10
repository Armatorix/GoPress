package content

import "github.com/Armatorix/GoPress/be/db/ent/ent"

func articlesFromEnt(ins ent.Articles) []Article {
	articles := make([]Article, len(ins))
	for i, in := range ins {
		articles[i] = articleFromEnt(in)
	}
	return articles
}

func articleFromEnt(in *ent.Article) Article {
	return Article{
		Body:        in.Body,
		Id:          in.ID,
		Title:       in.Title,
		Description: in.Description,
		Released:    in.Released,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}
