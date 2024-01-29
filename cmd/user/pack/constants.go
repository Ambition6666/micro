package pack

import (
	userdemo "micro/kitex_gen/userdemo"
	"micro/pkg/consts"
	"time"
)

func ReturnSuccess(msg string) *userdemo.BaseResp {
	return &userdemo.BaseResp{
		StatusCode: consts.Success,
		StatusMessage: msg,
		ServiceTime: time.Now().Unix(),
	}
}

func ReturnFail(msg string) *userdemo.BaseResp {
	return &userdemo.BaseResp{
		StatusCode: consts.Fail,
		StatusMessage: msg,
		ServiceTime: time.Now().Unix(),
	}
}