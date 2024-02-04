package test

import (
	"context"
	"micro/cmd/gateway/rpc"
	"micro/kitex_gen/chatwithbotdemo"
	"testing"
)

func TestXxx(t *testing.T) {
	rpc.InitRPC()

	c := *rpc.GetChatWithBotService()

	res, err := c.Chat(context.Background(), &chatwithbotdemo.ChatRequest{
		Msg:    "666",
		UserID: 1,
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}
