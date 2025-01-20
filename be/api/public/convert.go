package public

import (
	"io"
	"strings"

	"github.com/Armatorix/GoPress/be/db/ent/ent"
	"github.com/gorilla/feeds"
)

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
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}

type blogDetails struct {
	name        string
	link        string
	description string
	authorName  string
	authorEmail string
}

func rssFeedFromEnt(details blogDetails, ins ent.Articles) (r io.Reader, n int64, err error) {
	f := feeds.Feed{
		Title: details.name,
		Link: &feeds.Link{
			Href: details.link,
		},
		Description: details.description,
		Author: &feeds.Author{
			Name:  details.authorName,
			Email: details.authorEmail,
		},
		Items: rssFeedItemsFromEnt(ins),
	}

	rss, err := f.ToRss()
	if err != nil {
		return nil, 0, err
	}

	return strings.NewReader(rss), int64(len(rss)), nil
}

func rssFeedItemsFromEnt(ins ent.Articles) []*feeds.Item {
	items := make([]*feeds.Item, len(ins))
	for i, in := range ins {
		items[i] = rssFeedItemFromEnt(in)
	}
	return items
}

func rssFeedItemFromEnt(in *ent.Article) *feeds.Item {
	return &feeds.Item{
		Title:       in.Title,
		Link:        &feeds.Link{Href: "http://example.com/article/1"},
		Description: in.Description,
		Created:     in.CreatedAt,
		Updated:     in.UpdatedAt,
	}
}
