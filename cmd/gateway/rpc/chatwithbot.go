package rpc

import (
	"micro/kitex_gen/chatwithbotdemo/chatservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var ccli chatservice.Client

func initChatWithBot() {
	r, err := etcd.NewEtcdResolver([]string{"192.168.40.134:2379"})
	if err != nil {
		panic(err)
	}
	ccli, err = chatservice.NewClient("chat", client.WithResolver(r), client.WithTransportProtocol(transport.GRPC))
	if err != nil {
		panic(err)
	}
}

func GetChatWithBotService() *chatservice.Client {
	return &ccli
}
