package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Request struct {
	ID        uint           `gorm:"primaryKey"`
	Logs      pq.StringArray `gorm:"type:text[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Request{}); err != nil {
		panic(err)
	}
	fmt.Println("Migrated")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
