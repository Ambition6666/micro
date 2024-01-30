package vo

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	c "micro/pkg/consts"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx *app.RequestContext, err error,  Data any) {
	if err != nil {
		errno := c.ConvertErr(err)
		ctx.JSON(consts.StatusOK, Response{
			Code: errno.ErrCode,
			Message: errno.ErrMsg,
			Data: nil,
		})
		return
	}
	
	ctx.JSON(consts.StatusOK, Response{
		Code: c.Success.ErrCode,
		Message: c.Success.ErrMsg,
		Data: Data,
	})
}
