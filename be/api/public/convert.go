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
	blogUrl, err := url.Parse(details.BlogUrl)
	if err != nil {
		return nil, 0, err
	}

	f := feeds.Feed{
		Title: details.BlogName,
		Id:    details.BlogUrl,
		Link: &feeds.Link{
			Href: details.BlogUrl,
			Rel:  "self",
		},
		Description: details.BlogDesc,
		Author: &feeds.Author{
			Name:  details.BlogAuthor,
			Email: details.BlogEmail,
		},
		Items: rssFeedItemsFromEnt(blogUrl, ins),
	}

	rss, err := f.ToRss()
	if err != nil {
		return nil, 0, err
	}

	return strings.NewReader(rss), int64(len(rss)), nil
}

func rssFeedItemsFromEnt(blogUrl *url.URL, ins ent.Articles) []*feeds.Item {
	items := make([]*feeds.Item, len(ins))
	for i, in := range ins {
		itemUrl := blogUrl.JoinPath("articles", strconv.Itoa(in.ID))
		items[i] = rssFeedItemFromEnt(
			in,
			itemUrl,
		)
	}
	return items
}

func rssFeedItemFromEnt(in *ent.Article, itemUrl *url.URL) *feeds.Item {
	return &feeds.Item{
		Title:       in.Title,
		Link:        &feeds.Link{Href: itemUrl.String(), Rel: "self"},
		Description: in.Description,
		Created:     in.CreatedAt,
		Updated:     in.UpdatedAt,
		Id:          itemUrl.String(),
	}
}
