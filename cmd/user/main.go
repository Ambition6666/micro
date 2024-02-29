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
	zap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"os"
    "github.com/cloudwego/kitex/pkg/klog"
	// "micro/pkg/utils"
)

func init() {
	klog.SetLogger(zap.NewLogger())
	klog.SetLevel(klog.LevelDebug)
	db.InitDB()
}

func main() {

	//确认日志位置
	f, err := os.OpenFile("./log/output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    klog.SetOutput(f)

	klog.Infof("日志注册成功")

	serviceName := "user"

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

	// cfg := make(map[string]any)

	// configSuite := utils.InitConfigClient(serviceName, "test", 111, &cfg)

	svr := userdemo.NewServer(new(UserServiceImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}), server.WithRegistry(r), server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}

}
