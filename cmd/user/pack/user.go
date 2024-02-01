package pack

import (
	"micro/cmd/user/model"
	"micro/kitex_gen/userdemo"
)

func PackUser(u *model.User) *userdemo.User {
	return &userdemo.User{
		UserId: int64(u.ID),
		UserName: u.UserName,
		Email: u.Email,
	}
}

func PackUsers(u []*model.User) []*userdemo.User {
	
	users := make([]*userdemo.User, 0)
	
	for  i := 0; i < len(u); i++ {
		users = append(users, PackUser(u[i]))
	}

	return users
	
}