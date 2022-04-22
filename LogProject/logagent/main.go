package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"logagent/kafka"
	tailfilesource "logagent/tailFile"
	"time"
)

type AgentConfig struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	Topic    string `ini:"topic"`
	ChanSize int64  `ini:"chan_size"`
}

type CollectConfig struct {
	LogfilePath string `ini:"log_path"`
}

func run() (err error) {
	for {
		line, ok := <-tailfilesource.TailObj.Lines // tail工具将日志读出
		if !ok {
			logrus.Warn("file is reopening")
			time.Sleep(time.Second * 100)
		}
		fmt.Println("msg:", line.Text)
		msg := &sarama.ProducerMessage{} //新建一个sarama的生产者消息，用于封装tail读出的日志，并推至kafka
		msg.Topic = "web_log"
		msg.Value = sarama.StringEncoder(line.Text)
		//丢入通道（sarama）
		kafka.MsgChan <- msg
	}
}

//日志搜集的客户端 类似的开源项目还有filebeat
//1、收集指定目录下的log（tail）
//2、将log推送至kafka（sarama）

func main() {
	var configObj = new(AgentConfig)
	err := ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		logrus.Error("loadError:", err)
		return
	}
	fmt.Printf("%#v\n", configObj)
	//0.读配置文件（go-ini）
	//1.初始化
	err = kafka.InitKafka([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Error("kafka init failed :", err)
		return
	}
	logrus.Info("kafka init success!")

	err = tailfilesource.Init(configObj.CollectConfig.LogfilePath)
	if err != nil {
		logrus.Error("tail init failed,error:", err)
	}
	logrus.Info("tail init success!")
	//2.根据配置文件中的日志路径 使用tail去收集日志

	//3.利用sarama把日志送至kafka
	err = run()
}
