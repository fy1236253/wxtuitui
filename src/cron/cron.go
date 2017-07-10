package cron

import (
	"g"
	"log"
	"mq"
)

var (
	wechatWorkers []*mq.Consumer
)

func Start() {
	StartToken()
	CheckToken()
	wechatWorkers = make([]*mq.Consumer, g.Config().Worker.Wechat)

	for i := range wechatWorkers {
		wechatWorkers[i] = WechatConsumer()
		go wechatWorkers[i].StartUp()
	}

	log.Println("cron.Start ok")

}

func Stop() {
	for i := range wechatWorkers {
		wechatWorkers[i].Shutdown()
	}

	log.Println("cron.Stop ok")
}
