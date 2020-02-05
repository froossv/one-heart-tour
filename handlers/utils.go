package handlers

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

//GetS3 Credentials
func GetS3() *s3.S3 {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("BUCKET_REGION"))},
	)
	return s3.New(sess)
}
