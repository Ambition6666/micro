package test

import (
	"micro/pkg/utils"
	"testing"
)

type TestConfig struct {
	Data string `json:"test"`
}

func TestConfigCenter(t *testing.T) {
	cfg := new(TestConfig)
	utils.InitConfigClient("user","test",111,"192.168.40.134", "2379",cfg)
	t.Log(cfg)
}
