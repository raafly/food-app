package config

import (
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("!fo0dS3cr3tKeY%")

type JWTClaim struct {
	Username	string
	jwt.RegisteredClaims
}