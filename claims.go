package models

import "github.com/dgrijalva/jwt-go"

type (
	UserClaims struct {
		User string `json:"user"`
		Role string `json:"role"`
		jwt.StandardClaims
	}
)
