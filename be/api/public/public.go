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

// GetContent implements StrictServerInterface.
func (h *handler) GetContent(
	ctx context.Context,
	request GetContentRequestObject,
) (GetContentResponseObject, error) {
	art, err := h.db.GetPublishedArticleWithAuthor(ctx, request.ArticleId)
	if err != nil {
		return GetContent500JSONResponse{}, err
	}

	return GetContent200JSONResponse{
		GetContentJSONResponse{articleFromEnt(art)},
	}, nil
}

// GetContents implements StrictServerInterface.
func (h *handler) GetContents(
	ctx context.Context,
	request GetContentsRequestObject,
) (GetContentsResponseObject, error) {
	arts, err := h.db.GetPublishedArticlesWithAuthors(ctx)
	if err != nil {
		return GetContents500JSONResponse{}, err
	}

	return GetContents200JSONResponse{
		GetContentsJSONResponse{articlesFromEnt(arts)},
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
