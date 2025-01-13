// Package content provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package content

// Article defines model for Article.
type Article struct {
	Body        string   `json:"body"`
	Description string   `json:"description"`
	Id          int      `json:"id"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
}

// Articles defines model for Articles.
type Articles = []Article

// ErrorMsg defines model for ErrorMsg.
type ErrorMsg struct {
	Error string `json:"error"`
}

// PatchArticle defines model for PatchArticle.
type PatchArticle struct {
	Body        string   `json:"body"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
}

// PostArticle defines model for PostArticle.
type PostArticle struct {
	Body        string   `json:"body"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
}

// ArticleId defines model for articleId.
type ArticleId = int

// ArticleIds defines model for articleIds.
type ArticleIds = []int

// GetArticles defines model for GetArticles.
type GetArticles = Articles

// DeleteArticlesParams defines parameters for DeleteArticles.
type DeleteArticlesParams struct {
	ArticleIds ArticleIds `form:"articleIds" json:"articleIds"`
}

// UpdateArticleJSONRequestBody defines body for UpdateArticle for application/json ContentType.
type UpdateArticleJSONRequestBody = PatchArticle

// CreateArticleJSONRequestBody defines body for CreateArticle for application/json ContentType.
type CreateArticleJSONRequestBody = PostArticle
