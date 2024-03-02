package main

import (
	"github.com/dinzhen12306/framework/grpc"
	"github.com/dinzhen12306/template/api"
	"github.com/dinzhen12306/template/config"
	_ "github.com/dinzhen12306/template/config"
	"github.com/dinzhen12306/template/model/mysql"
	grpc2 "google.golang.org/grpc"
)

func main() {
	err := config.NewNacosConfig(func() error {
		err := mysql.Init()
		if err != nil {
			return err
		}
		mysql.MigrationDatabases()
		return nil
	})
	if err != nil {
		panic(err)
	}

	err = mysql.Init()
	if err != nil {
		panic(err)
	}
	mysql.MigrationDatabases()

	err = grpc.NewGrpcRegister(8081, func(grpcServer *grpc2.Server) {
		api.Register(grpcServer)
	})
	if err != nil {
		panic(err)
	}
}
