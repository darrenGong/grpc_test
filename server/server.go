package main

import (
	pb "grpc_test/proto/helloworld"
	"net"
	"os"

	logger "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"common"
	"strings"
	"fmt"
)

type Server struct{}

const (
	LOGFILE = "log/server.log"
	PORT = 10000

	registerPath = "/NS/TestEnv/hellomanager/master"
)

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	//logger.WithField("Server", "HelloWorld").Infof("Start handle function sayhello, Total[%d]", gTotalNum)

	rc := pb.ResponseCode{
		ErrCode:    0,
		ErrMessage: "",
	}
	return &pb.HelloResponse{
		Rc:      &rc,
		Message: "Hello " + in.GetName(),
	}, nil
}

func Init(laddr string) {
	file, err := os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.WithField("Server", "HelloWorld").Fatalf("Failed to open file[name:%s, err:%v]", "log", err)
	}

	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(file)
	logger.SetLevel(logger.DebugLevel)

	common.InitCommon(registerPath, laddr)
}

func Release() {
	common.Release()
}


func main() {
	ipAddr, err := common.GetIPV4Addr("wlan")
	if err != nil {
		log.Fatalf("Failed to get ip addr[%v]", err)
	}
	lAddr := strings.Join([]string{ipAddr, fmt.Sprintf("%d", PORT)}, ":")

	Init(lAddr)

	lis, err := net.Listen("tcp", lAddr)
	if err != nil {
		logger.WithField("Server", "HelloWorld").Fatalf("Failed to listen [laddr:%s, err:%v]", lAddr, err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &Server{})
	if err := server.Serve(lis); err != nil {
		logger.WithField("Server", "HelloWorld").Fatalf("Failed to serve [err:%v]", err)
	}

	Release()
}
