package main

import "surge/app"

func main() {
	app.InitApp()
}

//rdb := redisWrapper.GetClient()
//redisWrapper.AddSortedSet(context.Background(), "u1", float64(time.Now().UnixMilli()), float64(time.Now().UnixMilli()))
//res, _ := redisWrapper.GetCount(context.Background(), "u1",
//fmt.Sprintf("%d", time.Now().UnixMilli()-60*1000*30), fmt.Sprintf("%d", time.Now().UnixMilli()))
//redisWrapper.RemoveOldElements(context.Background(), "u1",
//fmt.Sprintf("%d", time.Now().UnixMilli()-60*1000*30))
//log.Println("counts is ", res)
