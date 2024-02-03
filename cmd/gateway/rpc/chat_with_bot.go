package rpc

import (
	"fmt"
	"micro/pkg/utils"

	"log"
	go_protoc "micro/kitex_gen/chatwithbotdemo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ccli go_protoc.ChatServiceClient
var conn *grpc.ClientConn
var err error

func initChatWithBot() {
	grpcClient := utils.NewGrpcClient("192.168.40.134", 2379, "/chat")
	ip := grpcClient.GetRrpcServIp()
	fmt.Println(ip)

	conn = new(grpc.ClientConn)

	conn, err = grpc.Dial(ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	ccli = go_protoc.NewChatServiceClient(conn)

}

func GetChatWithBotService() go_protoc.ChatServiceClient {
	return ccli
}

func CloseConn() {
	conn.Close()
}
