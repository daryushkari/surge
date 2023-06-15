package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"surge/delivery"
	"surge/pkg/redis"
)

func InitApp() {
	redis.InitClient("127.0.0.1:6379")

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
