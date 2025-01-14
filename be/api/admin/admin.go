//go:generate oapi-codegen -generate types -o types.gen.go -package admin ../../../api/admin.spec.yaml
//go:generate oapi-codegen -generate echo-server,strict-server -o server.gen.go -package admin ../../../api/admin.spec.yaml

package admin

import (
	"context"

	"github.com/Armatorix/GoPress/be/db"
	"github.com/Armatorix/GoPress/be/x/xecho"
)

type handler struct {
	db *db.DB
}

// UpdatePassword implements StrictServerInterface.
func (h *handler) UpdatePassword(
	ctx context.Context,
	request UpdatePasswordRequestObject,
) (UpdatePasswordResponseObject, error) {
	user := xecho.FetchUser(ctx)
	err := h.db.UpdatePassword(ctx, user.ID, request.Body.Password)
	if err != nil {
		return UpdatePassword500JSONResponse{}, err
	}
	return UpdatePassword200Response{}, nil
}

// DeleteUsers implements StrictServerInterface.
func (h *handler) DeleteUsers(
	ctx context.Context,
	request DeleteUsersRequestObject,
) (DeleteUsersResponseObject, error) {
	panic("unimplemented")
}

// GetUsers implements StrictServerInterface.
func (h *handler) GetUsers(
	ctx context.Context,
	request GetUsersRequestObject,
) (GetUsersResponseObject, error) {
	panic("unimplemented")
}

// PostAdminUsers implements StrictServerInterface.
func (h *handler) PostAdminUsers(
	ctx context.Context,
	request PostAdminUsersRequestObject,
) (PostAdminUsersResponseObject, error) {
	panic("unimplemented")
}

// UpdateUser implements StrictServerInterface.
func (h *handler) UpdateUser(
	ctx context.Context,
	request UpdateUserRequestObject,
) (UpdateUserResponseObject, error) {
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
