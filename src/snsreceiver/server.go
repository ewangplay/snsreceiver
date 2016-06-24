package main

import (
	"flag"
	"fmt"
	"github.com/ewangplay/jzlconfig"
	"github.com/outmana/log4jzl"
	"net/http"
	"os"
)

const (
	RETRY_MAX_COUNT = 3
)

var g_logger *log4jzl.Log4jzl
var g_config jzlconfig.JZLConfig
var g_nsqProducer *NSQProducer
var g_mysqladaptor *MysqlDBAdaptor

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [--config path_to_config_file]")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	//parse command line
	var configFile string
	flag.Usage = Usage
	flag.StringVar(&configFile, "config", "snsreceiver.conf", "specified config filename")
	flag.Parse()

	fmt.Println("config file: ", configFile)

	//read config file
	if err := g_config.Read(configFile); err == nil {
		fmt.Println(g_config)
	}

	//init logger
	var err error
	g_logger, err = log4jzl.New("snsreceiver")
	if err != nil {
		fmt.Println("Open log file fail.", err)
		os.Exit(1)
	}

	//init log level object
	g_logLevel, err = NewLogLevel()
	if err != nil {
		LOG_ERROR("创建SNSDBMgr对象失败，失败原因: %v", err)
		os.Exit(1)
	}

	//init nsq producer
	g_nsqProducer, err = NewNSQProducer()
	if err != nil {
		fmt.Println("create NSQProducer object fail.", err)
		os.Exit(1)
	}
	g_nsqProducer.Init()
	defer g_nsqProducer.Release()

	//init db adaptor
	g_mysqladaptor, err = NewMysqlDBAdaptor()
	if err != nil {
		fmt.Println("create MysqlDBAdaptor object fail.", err)
		os.Exit(1)
	}
	defer g_mysqladaptor.Release()

	//format the server listening newwork address
	var networkAddr string
	serviceIp, serviceIPIsSet := g_config.Get("service.addr")
	servicePort, servicePortIsSet := g_config.Get("service.port")
	if serviceIPIsSet && servicePortIsSet {
		networkAddr = fmt.Sprintf("%s:%s", serviceIp, servicePort)
	} else {
		networkAddr = "127.0.0.1:19090"
	}

	service := &BaseService{}
	weiboProcessor := &WeiboProcessor{}
	weixinProcessor := &WeixinProcessor{}
	router := &Router{service, map[string]JZLProcessor{
		//API每一个访问的资源需要在这里注册一个对应的服务
		"weibo":  weiboProcessor,
		"weixin": weixinProcessor,
	}}

	//启动服务
	fmt.Println("snsreceiver server working on", networkAddr)
	LOG_INFO("snsreceiver服务启动，监听地址：%v", networkAddr)

	err = http.ListenAndServe(networkAddr, router)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		os.Exit(1)
	}
}
