package main

import (
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
	retryConfig := retry.NewRetryConfig(
		retry.WithMaxAttemptTimes(10),
		retry.WithObserveDelay(20*time.Second),
		retry.WithRetryDelay(5*time.Second),
	)
	r, err := etcd.NewEtcdRegistryWithRetry([]string{"116.196.66.40:2379"}, retryConfig) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := userdemo.NewServer(new(UserServiceImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}), server.WithRegistry(r))
	err = svr.Run()
	if err != nil {
		log.Fatal(err)
	}

}
