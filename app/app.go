package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"surge/config"
	"surge/cronjob"
	"surge/delivery"
	natsBroker "surge/pkg/nats"
	postgresql "surge/pkg/postgis"
	"surge/pkg/redis"
	"time"
)

func InitApp() {
	err := config.InitCnf("config.json")
	cnf := config.GetCnf()
	if err != nil {
		log.Fatalln("error occurred in reading config:", err)
	}
	redisWrapper.InitClient(cnf.Redis)

	err = natsBroker.Connect(cnf)
	if err != nil {
		log.Fatalln("error occurred in connecting to nats:", err)
	}

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

	exit := HandleCronJobs(cnf)

	r := gin.Default()
	AddRideRouter(r)
	err = r.Run(cnf.ExternalExpose.Rest)
	if err != nil {
		log.Fatalln("error occurred:", err)
	}
	exit <- true
}

func HandleCronJobs(cnf *config.Config) chan bool {
	exit := make(chan bool)
	ticker := time.NewTicker(time.Millisecond * time.Duration(cnf.RequestLiveTime))
	go func() {
		for {
			select {
			case <-exit:
				ticker.Stop()
				log.Println("exiting program")
				return
			case <-ticker.C:
				cronjob.RemoveOldRequest()
			}
		}
	}()
	return exit
}

func AddRideRouter(r *gin.Engine) {
	userRouter := r.Group("/ride")
	{
		userRouter.POST("/save", delivery.SaveRide)
	}
}
