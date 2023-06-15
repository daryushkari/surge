package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"surge/delivery"
)

func InitApp() {
	r := gin.Default()
	AddRideRouter(r)
	err := r.Run("localhost:8080")

	if err != nil {
		log.Fatalln("error occurred:", err)
	}
}

func AddRideRouter(r *gin.Engine) {
	userRouter := r.Group("/ride")
	{
		userRouter.POST("/save", delivery.SaveRide)
	}
}
