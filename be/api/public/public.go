//go:generate oapi-codegen -generate types -o types.gen.go -package public ../../../api/public.spec.yaml
//go:generate oapi-codegen -generate echo-server,strict-server -o server.gen.go -package public ../../../api/public.spec.yaml

package public

import (
	"context"

	"github.com/Armatorix/GoPress/be/db"
)

type handler struct {
	db *db.DB
}

// GetArticle implements StrictServerInterface.
func (h *handler) GetArticle(
	ctx context.Context,
	request GetArticleRequestObject,
) (GetArticleResponseObject, error) {
	art, err := h.db.GetPublishedArticleWithAuthor(ctx, request.ArticleId)
	if err != nil {
		return GetArticle500JSONResponse{}, err
	}

	return GetArticle200JSONResponse{
		GetArticleJSONResponse{articleFromEnt(art)},
	}, nil
}

// GetArticles implements StrictServerInterface.
func (h *handler) GetArticles(
	ctx context.Context,
	request GetArticlesRequestObject,
) (GetArticlesResponseObject, error) {
	arts, err := h.db.GetPublishedArticlesWithAuthors(ctx)
	if err != nil {
		return GetArticles500JSONResponse{}, err
	}

	return GetArticles200JSONResponse{
		GetArticlesJSONResponse{articlesFromEnt(arts)},
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
