package api

import (
	"github.com/dinzhen12306/message/user"
	"google.golang.org/grpc"
)

func Register(s grpc.ServiceRegistrar) {
	user.RegisterUserServerServer(s, new(UserRpcServer))
}
