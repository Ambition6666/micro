package utils

import (
	"math/rand"
	"time"
)

// 产生随机数
func Randnum(scope int) int {
	rand.New(rand.NewSource(time.Now().Unix()))
	i := rand.Intn(scope)
	return i
}