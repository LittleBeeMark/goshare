package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/mineralres/goshare/pkg/pb"
)

// 在独立进程中运行策略
func main() {
	cl := pb.NewUserManagerService("go.micro.srv.ucenter", client.DefaultClient)
	resp, err := cl.UserLogin(context.Background(), &pb.ReqUserLogin{})
	log.Println(resp, err)
}
