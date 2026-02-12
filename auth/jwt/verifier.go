package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Verifier struct {
	keyFunc jwt.Keyfunc
}

func NewHMAC(secret string) *Verifier {
	return &Verifier{
		keyFunc: func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		},
	}
}

func Verify[T jwt.Claims](v *Verifier, tokenString string) (*T, error) {
	claims := new(T)

	token, err := jwt.ParseWithClaims(
		tokenString,
		*claims,
		v.keyFunc,
	)

	if err != nil {
		return nil, ErrInvalidToken
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
