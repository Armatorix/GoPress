//go:generate oapi-codegen -generate types -o types.gen.go -package content ../../../api/content.spec.yaml
//go:generate oapi-codegen -generate echo-server,strict-server -o server.gen.go -package content ../../../api/content.spec.yaml

package content

import (
	"context"

	"github.com/Armatorix/GoPress/be/db"
)

type handler struct {
	db *db.DB
}

// CreateArticle implements StrictServerInterface.
func (h *handler) CreateArticle(
	ctx context.Context,
	request CreateArticleRequestObject,
) (CreateArticleResponseObject, error) {
	err := h.db.InsertArticle(ctx, db.InsertArticle{
		Body:  request.Body.Body,
		Title: request.Body.Title,
	})
	if err != nil {
		return CreateArticle500JSONResponse{}, err
	}

	return CreateArticle200Response{}, nil
}

// UpdateArticle implements StrictServerInterface.
func (h *handler) UpdateArticle(ctx context.Context, request UpdateArticleRequestObject) (UpdateArticleResponseObject, error) {
	err := h.db.UpdateArticle(ctx, request.ArticleId, db.InsertArticle{
		Body:  request.Body.Body,
		Title: request.Body.Title,
	})
	if err != nil {
		return UpdateArticle500JSONResponse{}, err
	}

	return UpdateArticle200Response{}, nil
}

// DeleteAdminArticles implements StrictServerInterface.
func (h *handler) DeleteArticles(
	ctx context.Context,
	request DeleteArticlesRequestObject,
) (DeleteArticlesResponseObject, error) {
	err := h.db.DeleteArticles(ctx, request.Params.ArticleIds...)
	if err != nil {
		return DeleteArticles500JSONResponse{}, err
	}

	return DeleteArticles200Response{}, nil
}

// GetArticles implements StrictServerInterface.
func (h *handler) GetArticles(
	ctx context.Context,
	request GetArticlesRequestObject,
) (GetArticlesResponseObject, error) {
	articles, err := h.db.GetArticles(ctx)
	if err != nil {
		return GetArticles500JSONResponse{}, err
	}

	return GetArticles200JSONResponse{
		GetArticlesJSONResponse: articlesFromEnt(articles),
	}, nil
}

func NewHandler(db *db.DB) ServerInterface {
	return NewStrictHandler(
		&handler{
			db: db,
		},
		nil,
	)
}
