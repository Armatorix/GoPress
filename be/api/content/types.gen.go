// Package content provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package content

// Article defines model for Article.
type Article struct {
	Body  string   `json:"body"`
	Id    int      `json:"id"`
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
}

// Articles defines model for Articles.
type Articles = []Article

// ErrorMsg defines model for ErrorMsg.
type ErrorMsg struct {
	Error string `json:"error"`
}

// PatchArticle defines model for PatchArticle.
type PatchArticle struct {
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
}

// PostArticle defines model for PostArticle.
type PostArticle struct {
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
}

// ArticleId defines model for articleId.
type ArticleId = string

// GetArticles defines model for GetArticles.
type GetArticles = Articles

// PatchAdminArticleArticleIdJSONRequestBody defines body for PatchAdminArticleArticleId for application/json ContentType.
type PatchAdminArticleArticleIdJSONRequestBody = PatchArticle

// PostAdminArticlesJSONRequestBody defines body for PostAdminArticles for application/json ContentType.
type PostAdminArticlesJSONRequestBody = PostArticle
