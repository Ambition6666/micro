package main

import (
	"context"
	"micro/cmd/gateway/handlers"
	"micro/cmd/gateway/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func init() {
	rpc.InitRPC()
}

func main() {


	h := server.Default(server.WithHostPorts(":9090"))

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.POST("/user", handlers.CreateUser)

	h.GET("/chat", handlers.Chat)

	h.Spin()

}
