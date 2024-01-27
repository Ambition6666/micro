package main

import (
	"log"
	userdemo "micro/kitex_gen/userdemo/userservice"
)

func main() {
	svr := userdemo.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
