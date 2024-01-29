package handlers

import (
	"context"
	"fmt"
	"micro/cmd/gateway/rpc"
	"micro/cmd/gateway/vo"
	"micro/kitex_gen/userdemo"

	"github.com/cloudwego/hertz/pkg/app"
)

func CreateUser(c context.Context, ctx *app.RequestContext) {

	info := new(vo.CreateUserRequest)

	ctx.BindJSON(info)

	ucli := *rpc.GetUserService()

	resp, err := ucli.CreateUser(c, &userdemo.CreateUserRequest{
		UserName: info.UserName,
		Password: info.Password,
		Email:    info.Email,
	})

	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
	}

	fmt.Println(resp)

	ctx.JSON(200, "")
}
