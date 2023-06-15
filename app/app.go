package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"surge/config"
	"surge/delivery"
	postgresql "surge/pkg/postgis"
	"surge/pkg/redis"
)

func InitApp() {
	err := config.InitCnf("config.json")
	cnf := config.GetCnf()
	if err != nil {
		log.Fatalln("error occurred in reading config:", err)
	}
	redisWrapper.InitClient(cnf.Redis)

	err = postgresql.Init(cnf)
	if err != nil {
		log.Fatalln("error occurred connecting database:", err)
	}
	pdb := postgresql.Get()

	err = migrate(pdb)
	if err != nil {
		log.Fatalln("error occurred in migration", err)
	}
	Seed(pdb)

	r := gin.Default()
	AddRideRouter(r)
	err = r.Run(cnf.ExternalExpose.Rest)
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
