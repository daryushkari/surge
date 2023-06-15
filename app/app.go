package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"surge/config"
	"surge/delivery"
	"surge/pkg/redis"
)

func InitApp() {
	err := config.InitCnf("config.json")
	cnf := config.GetCnf()
	if err != nil {
		log.Fatalln("error occurred in reading config:", err)
	}
	log.Println(cnf.Redis)
	redisWrapper.InitClient(cnf.Redis)

	r := gin.Default()
	AddRideRouter(r)
	err = r.Run("localhost:8080")
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
