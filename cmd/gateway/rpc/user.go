package rpc

import (
	"micro/kitex_gen/userdemo/userservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var ucli userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{"192.168.40.134:2379"})
	if err != nil {
		panic(err)
	}
	ucli, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func GetUserService() *userservice.Client {
	return &ucli
}
