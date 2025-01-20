package public

import (
	"io"
	"net/url"
	"strconv"
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

func rssFeedFromEnt(details Config, ins ent.Articles) (r io.Reader, n int64, err error) {
	items, err := rssFeedItemsFromEnt(details, ins)
	if err != nil {
		return nil, 0, err
	}
	f := feeds.Feed{
		Title: details.BlogName,
		Link: &feeds.Link{
			Href: details.BlogUrl,
		},
		Description: details.BlogDesc,
		Author: &feeds.Author{
			Name:  details.BlogAuthor,
			Email: details.BlogEmail,
		},
		Items: items,
	}

	rss, err := f.ToRss()
	if err != nil {
		return nil, 0, err
	}

	return strings.NewReader(rss), int64(len(rss)), nil
}

func rssFeedItemsFromEnt(details Config, ins ent.Articles) ([]*feeds.Item, error) {
	url, err := url.Parse(details.BlogUrl)
	if err != nil {
		return nil, err
	}
	items := make([]*feeds.Item, len(ins))
	for i, in := range ins {
		itemUrl := *url
		url.JoinPath("articles", strconv.Itoa(in.ID))
		items[i] = rssFeedItemFromEnt(
			in,
			itemUrl,
		)
	}
	return items, nil
}

func rssFeedItemFromEnt(in *ent.Article, url url.URL) *feeds.Item {
	return &feeds.Item{
		Title:       in.Title,
		Link:        &feeds.Link{Href: url.String()},
		Description: in.Description,
		Created:     in.CreatedAt,
		Updated:     in.UpdatedAt,
	}
}
