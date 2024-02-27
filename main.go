package main

import (
	"github.com/dinzhen12306/framework/grpc"
	"github.com/dinzhen12306/message/user"
	"github.com/dinzhen12306/template/api"
	_ "github.com/dinzhen12306/template/config"
	grpc2 "google.golang.org/grpc"
)

func main() {
	err := grpc.NewGrpcRegister(8081, func(grpcServer *grpc2.Server) {
		user.RegisterUserServerServer(grpcServer, new(api.UserRpcServer))
	})
	if err != nil {
		panic(err)
	}
}
