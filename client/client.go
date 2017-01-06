package main

import (
	"context"
	logger "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
	pb "grpc_test/proto/helloworld"
	"os"
	"time"
)

const (
	lAddr          = ":10000"
	ConnTimeout    = 10 * time.Second
	LOGFILE        = "log/client.log"
	ConcurrencyNum = 1
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
}

func SendConn(conn *grpc.ClientConn, chanBlack chan<- int) {
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
		LOG.Infof("Recv message: %s", r.GetMessage())
	}

	chanBlack <- 1
}

func main() {
	Init()

	conn, err := grpc.Dial(lAddr, grpc.WithInsecure())
	if err != nil {
		LOG.Fatalf("Failed to dial[laddr:%s, err:%v]", lAddr, err)
	}
	defer conn.Close()

	nowTime := time.Now()
	chanBlock := make(chan int)
	for i := 0; i < ConcurrencyNum; i++ {
		go SendConn(conn, chanBlock)
	}

	for i := 0; i < ConcurrencyNum; i++ {
		<-chanBlock
	}

	LOG.Infof("Total cost time: %v", time.Since(nowTime))
}
