package handlers

import (
	"github.com/dgrijalva/jwt-go"
)

//User Struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Response Generic
type Response struct {
	Message string `json:"message"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
