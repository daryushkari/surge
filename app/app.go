package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitApp() {
	r := gin.Default()
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatalln("error occurred:", err)
	}
}
