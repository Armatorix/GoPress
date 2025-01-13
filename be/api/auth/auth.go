//go:generate oapi-codegen -generate types -o types.gen.go -package auth ../../../api/auth.spec.yaml
//go:generate oapi-codegen -generate echo-server,strict-server -o server.gen.go -package auth ../../../api/auth.spec.yaml

package auth

import (
	"context"
	"time"

	"github.com/Armatorix/GoPress/be/db"
	"github.com/Armatorix/GoPress/be/db/ent/ent"
	"github.com/Armatorix/GoPress/be/x/xjwt"
	"github.com/golang-jwt/jwt/v5"
)

type handler struct {
	db  *db.DB
	jwt *xjwt.Client
}

// PostAdminLogin implements StrictServerInterface.
func (h *handler) PostAdminLogin(
	ctx context.Context,
	request PostAdminLoginRequestObject,
) (PostAdminLoginResponseObject, error) {
	user, err := h.db.UserByEmailOrNick(ctx, request.Body.Email)
	if err != nil {
		return PostAdminLogin500JSONResponse{}, err
	}

	// TODO: user token
	if user.Password.String() != request.Body.Password {
		return PostAdminLogin400JSONResponse{}, nil
	}

	token, err := h.userToken(user)
	if err != nil {
		return PostAdminLogin500JSONResponse{}, err
	}
	return PostAdminLogin200JSONResponse{
		AuthTokenJSONResponse(*token),
	}, nil
}

func NewHandler(db *db.DB, jwtc *xjwt.Client) ServerInterface {
	return NewStrictHandler(
		&handler{
			db:  db,
			jwt: jwtc,
		},
		nil,
	)
}

func (h *handler) userToken(user *ent.User) (*AuthToken, error) {
	expAt := time.Now().Add(time.Hour * 24 * 7)
	jwt, err := h.jwt.Encode(xjwt.Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gopress",
			Audience:  jwt.ClaimStrings{"gopress"},
			ExpiresAt: jwt.NewNumericDate(expAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	if err != nil {
		return nil, err
	}
	return &AuthToken{
		Token:     jwt,
		ExpiresAt: expAt,
	}, nil
}
