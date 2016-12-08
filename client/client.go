package main

import (
	"os"
	logger "github.com/Sirupsen/logrus"
	"net"
	"time"
	pb "go_test/grpc_test/proto/helloworld"
	"context"
)

const (
	lAddr = ":10000"
	ConnTimeout = 10 * time.Second
)

var (
	LogField logger.Fields
	LOG *logger.Entry
)

func Init() {
	LogField = logger.Fields{
		"Server": "HelloWorld",
	}
	LOG = logger.WithFields(LogField)
	file, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		LOG.Fatalf("Failed to open file[name:%s, err:%v]", "log", err)
	}

	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(file)
	logger.SetLevel(logger.DebugLevel)
}

func main() {
	Init()

	conn, err := net.DialTimeout("tcp", lAddr, ConnTimeout)
	if err != nil {
		LOG.Fatalf("Failed to dial[laddr:%s, err:%v]", lAddr, err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	r, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World !"})
	if err != nil {
		LOG.Fatalf("Failed to say hello[err:%v]", err)
	}
	if r.GetRc().ErrCode != 0 {
		LOG.Fatalf("Return err after request say hello[errInfo:%v]", r.GetRc())
	}

	LOG.Infof("Recv message: %s", r.GetMessage())
}
