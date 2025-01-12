//go:generate oapi-codegen -generate types -o types.gen.go -package auth ../../../api/auth.spec.yaml
//go:generate oapi-codegen -generate echo-server,strict-server -o server.gen.go -package auth ../../../api/auth.spec.yaml

package auth

import (
	"context"

	"github.com/Armatorix/GoPress/be/db"
)

type handler struct {
	db *db.DB
}

// PostAdminLogin implements StrictServerInterface.
func (h *handler) PostAdminLogin(ctx context.Context, request PostAdminLoginRequestObject) (PostAdminLoginResponseObject, error) {
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
