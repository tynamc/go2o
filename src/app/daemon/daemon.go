/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-01-08 21:01
 * description :
 * history :
 */

package daemon

import (
	"flag"
	"fmt"
	"github.com/jrsix/gof"
	"go2o/src/core"
	"go2o/src/core/service/dps"
	"strings"
)

const (
	ServeMail  string = "mail"
	ServeOrder string = "order"
)

var (
	appCtx     *core.MainApp
	allService = []string{ServeMail, ServeOrder}
)

// 运行
func Run(ctx gof.App) {
	if ctx != nil {
		appCtx = ctx.(*core.MainApp)
	} else {
		appCtx = getAppCtx("app.conf")
	}
	bootService(allService, false)
}

// 自定义参数运行
func FlagRun() {
	var conf string
	var debug bool
	var trace bool
	var service string
	var serviceArr []string
	flag.StringVar(&conf, "conf", "app.conf", "")
	flag.BoolVar(&debug, "debug", true, "")
	flag.BoolVar(&trace, "trace", true, "")
	flag.StringVar(&service, "service", strings.Join(allService, ","), "")

	flag.Parse()

	appCtx = getAppCtx(conf)
	appCtx.Init(debug, trace)
	gof.CurrentApp = appCtx

	dps.Init(appCtx)

	if service == "all" {
		serviceArr = allService
	} else {
		serviceArr = strings.Split(service, ",")
	}

	bootService(serviceArr, true)
}

func getAppCtx(conf string) *core.MainApp {
	return core.NewMainApp(conf)
}

func bootService(arr []string, standOne bool) {
	fmt.Println("[ Go2o][ Daemon][ Booted] - Daemon service is running.")
	for _, v := range arr {
		switch v {
		case ServeMail:
			fmt.Println("[ Go2o][ Daemon][ Booted] - mail daemon running")
			go startMailQueue()
		case ServeOrder:
			fmt.Println("[ Go2o][ Daemon][ Booted] - order daemon running")
			go orderDaemon()
		}
	}
	if standOne {
		var ch chan int = make(chan int)
		<-ch
	}
}

func recoverDaemon() {

}
