package main

import (
	pb "grpc_test/proto/helloworld"
	"net"
	"os"
	"sync"

	logger "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct{}

const (
	lAddr   = ":10000"
	LOGFILE = "log/server.log"
)

var (
	gTotalNum uint32
	gMutex    sync.Mutex
)

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	gMutex.Lock()
	gTotalNum += 1
	logger.WithField("Server", "HelloWorld").Infof("Start handle function sayhello, Total[%d]", gTotalNum)
	gMutex.Unlock()

	rc := pb.ResponseCode{
		ErrCode:    0,
		ErrMessage: "",
	}
	return &pb.HelloResponse{
		Rc:      &rc,
		Message: "Hello " + in.GetName(),
	}, nil
}

func Init() {
	file, err := os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.WithField("Server", "HelloWorld").Fatalf("Failed to open file[name:%s, err:%v]", "log", err)
	}

	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(file)
	logger.SetLevel(logger.DebugLevel)
}

func main() {
	Init()

	lis, err := net.Listen("tcp", lAddr)
	if err != nil {
		logger.WithField("Server", "HelloWorld").Fatalf("Failed to listen [laddr:%s, err:%v]", lAddr, err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &Server{})
	if err := server.Serve(lis); err != nil {
		logger.WithField("Server", "HelloWorld").Fatalf("Failed to serve [err:%v]", err)
	}
}
