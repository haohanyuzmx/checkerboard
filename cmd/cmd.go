package main

import "time"

func main()  {
	hander.Init()
	mq.Init()
	db.InitDB()
	for i := 0; i < 20; i++ {
		go mq.ConsumePlayerMessage()
	}
	for i := 0; i < 10; i++ {
		go mq.ConsumeResultMessage()
	}

}

