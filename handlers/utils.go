package handlers

import (
	"os"

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
