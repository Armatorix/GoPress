package public

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
		Author:      in.Edges.User.Nick,
		Title:       in.Title,
		Description: in.Description,
		Body:        in.Body,
		Id:          in.ID,
		Tags:        []string{},
	}
}
