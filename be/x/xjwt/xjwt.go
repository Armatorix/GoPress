package xjwt

import "github.com/golang-jwt/jwt/v5"

type Client struct {
	secret []byte
}

type Claims struct {
	jwt.RegisteredClaims
	UserId     int  `json:"userId"`
	IsVerified bool `json:"isVerified"`
}

func New(secret []byte) *Client {
	return &Client{
		secret: secret,
	}
}

func (c *Client) Encode(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(c.secret)
}

func (c *Client) Decode(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return c.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
