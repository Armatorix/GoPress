package openai

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type Config struct {
	ApiKey string `env:"OPENAI_API_KEY"`
}
type Client struct {
	ApiKey string
	c      *openai.Client
}

func New(cfg Config) *Client {
	return &Client{
		ApiKey: cfg.ApiKey,
		c:      openai.NewClient(cfg.ApiKey),
	}
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func (c *Client) generateArticle(ctx context.Context, topic string, contentTypePrompt string) (*Article, error) {
	if c.ApiKey == "" {
		return nil, errors.New("missing API key")
	}

	resp, err := c.c.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: "You are a content provider that generates articles on a topic provided by a user." +
					"Write it as a human, do not include headlines like 'conclusion' or 'summary'." +
					"It will be used as a part of a blog post - make nuce styling for it." +
					"Description and body are type of html and the content of them should be wrapped in div." +
					"Reply in a format of a json object with keys 'title', 'description', and 'body'." +
					contentTypePrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: topic,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("missing response")
	}

	msg := resp.Choices[0].Message.Content
	idx := strings.IndexByte(msg, '{')
	if idx != -1 {
		msg = msg[idx:]
	}
	idx = strings.LastIndexByte(msg, '}')
	if idx != -1 {
		msg = msg[:idx+1]
	}

	var article Article
	if err := json.Unmarshal([]byte(msg), &article); err != nil {
		return nil, err
	}
	return &article, nil
}

func (c *Client) GenerateHTMLArticle(ctx context.Context, topic string) (*Article, error) {
	return c.generateArticle(ctx, topic, "Use default HTML tags for styling.")
}

func (c *Client) GenerateTailwindCssArticle(ctx context.Context, topic string) (*Article, error) {
	return c.generateArticle(ctx, topic, "Use Tailwind CSS for styling.")
}
