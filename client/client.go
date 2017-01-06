package main

import (
	"context"
	logger "github.com/Sirupsen/logrus"
	pb "grpc_test/proto/helloworld"
	"os"
	"time"
	"common"
)

const (
	ConnTimeout    = 10 * time.Second
	LOGFILE        = "log/client.log"
	ConcurrencyNum = 1

	registerPath = "/NS/TestEnv/helloclient/master"
	nodeValue = "hello client"
)

var (
	LogField logger.Fields
	LOG      *logger.Entry
)

func Init() {
	LogField = logger.Fields{
		"Server": "HelloWorld",
	}
	LOG = logger.WithFields(LogField)
	file, err := os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		LOG.Fatalf("Failed to open file[name:%s, err:%v]", "log", err)
	}

	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(file)
	logger.SetLevel(logger.DebugLevel)

	common.InitCommon(registerPath, nodeValue)
}

func Release() {
	common.Release()
}

func SendConn(chanBlack chan<- int) {
	conn, err := common.GetConnToHelloManager()
	if err != nil {
		LOG.Errorf("Failed to get conn to hello manager[err:%v]", err)
		chanBlack <- 1
		return
	}

	client := pb.NewGreeterClient(conn)

	for i := 0; i <= 10; i++ {
		time.Sleep(5 * time.Second)
		r, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World !"})
		if err != nil {
			LOG.Errorf("Failed to say hello[err:%v]", err)
			continue
		}
		if r.GetRc().ErrCode != 0 {
			LOG.Fatalf("Return err after request say hello[errInfo:%v]", r.GetRc())
		}
		logger.Printf("Recv message: %s", r.GetMessage())
	}

	chanBlack <- 1
}

func main() {
	Init()

	nowTime := time.Now()
	chanBlock := make(chan int)
	for i := 0; i < ConcurrencyNum; i++ {
		go SendConn(chanBlock)
	}

	for i := 0; i < ConcurrencyNum; i++ {
		<-chanBlock
	}

	LOG.Infof("Total cost time: %v", time.Since(nowTime))
}
