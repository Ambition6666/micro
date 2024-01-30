package main

import (
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/registry-etcd/retry"

	"log"
	"micro/cmd/user/dal/db"
	userdemo "micro/kitex_gen/userdemo/userservice"
)

func init() {
	db.InitDB()
}

func main() {

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")

	retryConfig := retry.NewRetryConfig(
		retry.WithMaxAttemptTimes(10),
		retry.WithObserveDelay(20*time.Second),
		retry.WithRetryDelay(5*time.Second),
	)
	r, err := etcd.NewEtcdRegistryWithRetry([]string{"192.168.40.134:2379"}, retryConfig) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := userdemo.NewServer(new(UserServiceImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}), server.WithRegistry(r), server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		log.Fatal(err)
	}

}
