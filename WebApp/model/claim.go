package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	UserId         string
	Role           string
	jwt.StandardClaims
}

// Valid implements jwt.Claims.
// func (j *JwtCustomClaims) Valid() error {
// 	panic("unimplemented")
// }
