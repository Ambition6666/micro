package main

import (
	"context"
	"micro/cmd/user/dal/db"
	"micro/cmd/user/model"
	"micro/cmd/user/pack"
	userdemo "micro/kitex_gen/userdemo"
	"micro/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (resp *userdemo.CreateUserResponse, err error) {
	// TODO: Your code here...

	resp = new(userdemo.CreateUserResponse)

	u := model.NewUser(req.UserName, utils.Encrypt(req.Password, ""), req.Email)

	err = db.CreateUser(ctx, u)

	if err != nil {
		resp.BaseResp = pack.ReturnFail("创建失败") 
		return 
	}

	resp.BaseResp = pack.ReturnSuccess("创建成功")

	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (resp *userdemo.MGetUserResponse, err error) {
	// TODO: Your code here...

	resp = new(userdemo.MGetUserResponse)	

	u, err := db.MGetUsers(ctx,req.UserIds)

	if err != nil {
		resp.BaseResp = pack.ReturnFail("获取失败")
		return  
	}

	resp.BaseResp = pack.ReturnSuccess("获取成功")
	resp.Users = pack.PackUsers(u)
	return

}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (resp *userdemo.CheckUserResponse, err error) {
	// TODO: Your code here...

	resp = new(userdemo.CheckUserResponse)

	u, err := db.QueryUser(ctx, req.Email)

	if err != nil {
		resp.BaseResp = pack.ReturnFail("查询失败")

		return
	}

	if u.Password != utils.Encrypt(req.Password, "") {
		resp.BaseResp = pack.ReturnFail("密码错误")

		return 
	}

	resp.BaseResp = pack.ReturnSuccess("校验成功")
	resp.UserId = int64(u.ID)

	return
}
