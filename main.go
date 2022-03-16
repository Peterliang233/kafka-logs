package main

import (
	"fmt"
	"github.com/Peterliang233/kafka-log/config"
	"github.com/Peterliang233/kafka-log/kafka"
	"github.com/Peterliang233/kafka-log/taillog"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var iniFile = new(config.AppConfig)

func run() {
	for {
		select {
		case line := <-taillog.ReadChan():
			fmt.Println(line)
			kafka.SendToKafka(iniFile.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
func main() {
	err := ini.MapTo(iniFile, "./config/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	err = kafka.Init([]string{iniFile.Address})
	if err != nil {
		fmt.Println("init kafka error", err)
		return
	}
	fmt.Println("init kafka success")
	err = taillog.Init(iniFile.Path)
	if err != nil {
		fmt.Println("init tail file error", err)
		return
	}
	fmt.Println("init tail file success")
	run()
}
