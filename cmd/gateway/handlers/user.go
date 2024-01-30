package handlers

import (
	"context"

	"micro/cmd/gateway/rpc"
	"micro/cmd/gateway/vo"
	"micro/kitex_gen/userdemo"

	"github.com/cloudwego/hertz/pkg/app"
)

func CreateUser(c context.Context, ctx *app.RequestContext) {

	req := new(vo.CreateUserRequest)

	ctx.BindJSON(req)

	ucli := *rpc.GetUserService()

	_, err := ucli.CreateUser(c, &userdemo.CreateUserRequest{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
	})


	vo.SendResponse(ctx, err, nil)
	
}
