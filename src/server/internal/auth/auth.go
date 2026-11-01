// Package auth defines token authentication used by the API.
package auth

import "github.com/golang-jwt/jwt/v5"

type Authenticator interface {
	GenerateToken(claims jwt.Claims) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
