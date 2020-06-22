package main

import (
	"cron/db"
	"cron/opt"
	"github.com/robfig/cron"
	"log"
)

func main() {

	c := cron.New()
	eID , _ := c.AddFunc(db.Cron.Spec, opt.OrderCallBack)
	//eID , _ := c.AddFunc(spec, OrderCallBackWithGroup)
	log.Println(eID)
	c.Start()

	select {}
}
