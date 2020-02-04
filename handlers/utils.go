package handlers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//GetDB Get Database Instance
func GetDB() *gorm.DB {
	dburl := os.Getenv("DATABASE_URL")
	psqlInfo, _ := pq.ParseURL(dburl)
	db, _ := gorm.Open("postgres", psqlInfo)
	return db
}

//CreateJWTString Util
func CreateJWTString(Username string) (string, time.Time) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(os.Getenv("JWTSIGNINGKEY"))
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString, expirationTime
}
