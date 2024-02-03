package handlers

import (
	"context"

	"micro/cmd/gateway/rpc"
	"micro/cmd/gateway/vo"
	"micro/kitex_gen/chatwithbotdemo"

	"github.com/cloudwego/hertz/pkg/app"
)

func Chat(c context.Context, ctx *app.RequestContext) {

	msg := ctx.Query("msg")

	ccli := rpc.GetChatWithBotService()

	res, err := ccli.Chat(c, &chatwithbotdemo.ChatRequest{
		Msg:    msg,
		UserID: 1,
	})

	vo.SendResponse(ctx, err, res)
}
