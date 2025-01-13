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

// DeleteAdminArticles implements StrictServerInterface.
func (h *handler) DeleteAdminArticles(
	ctx context.Context,
	request DeleteAdminArticlesRequestObject,
) (DeleteAdminArticlesResponseObject, error) {
	panic("unimplemented")
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

// PatchAdminArticleArticleId implements StrictServerInterface.
func (h *handler) PatchAdminArticleArticleId(
	ctx context.Context,
	request PatchAdminArticleArticleIdRequestObject,
) (PatchAdminArticleArticleIdResponseObject, error) {
	panic("unimplemented")
}

// PostAdminArticles implements StrictServerInterface.
func (h *handler) PostAdminArticles(
	ctx context.Context,
	request PostAdminArticlesRequestObject,
) (PostAdminArticlesResponseObject, error) {
	panic("unimplemented")
}

func NewHandler(db *db.DB) ServerInterface {
	return NewStrictHandler(
		&handler{
			db: db,
		},
		nil,
	)
}
