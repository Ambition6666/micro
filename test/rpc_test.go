package test

import (
	"micro/cmd/gateway/rpc"
	"testing"
)

func TestXxx(t *testing.T) {
	rpc.InitRPC()
	rpc.CloseConn()
}
