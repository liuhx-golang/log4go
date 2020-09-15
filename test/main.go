package main

import (
	"flag"
	"time"

	log4go "github.com/liuhx-golang/log4go"
)

var confPath = flag.String("conf", "./conf/", "The conf path.")

func main() {
	flag.Parse()
	Log4goConf := *confPath + "log4g.xml"
	log4go.LoggerConfiguration(Log4goConf)
	log4go.GetLogger().Info("开始%v","记录")
	time.Sleep(time.Duration(2)*time.Second)
}