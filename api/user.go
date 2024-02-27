package api

import (
	"context"
	pb "github.com/dinzhen12306/message/user"
	"github.com/dinzhen12306/template/model/mysql"
)

type UserRpcServer struct {
	pb.UnimplementedUserServerServer
}

func (s *UserRpcServer) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.GetUserResp, error) {
	inMap := make(map[string]interface{})
	for k, v := range in.Where {
		inMap[k] = v
	}
	user := mysql.NewUser()
	err := user.Get(inMap)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResp{
		UserInfo: user.MysqlToPb(),
	}, nil
}

func (s *UserRpcServer) GetUsers(ctx context.Context, in *pb.GetUsersReq) (*pb.GetUsersResp, error) {
	user := mysql.NewUser()
	users, err := user.GetUsers(in.Limit, in.Size)
	if err != nil {
		return nil, err
	}
	var data []*pb.User
	for _, v := range users {
		data = append(data, v.MysqlToPb())
	}
	return &pb.GetUsersResp{
		UserInfo: data,
	}, nil
}

func (s *UserRpcServer) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	user := mysql.PbToMysql(in.UserInfo)
	err := user.Create()
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResp{
		UserInfo: user.MysqlToPb(),
	}, nil
}

func (s *UserRpcServer) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	user := mysql.PbToMysql(in.UserInfo)
	err := user.Update()
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResp{
		UserInfo: user.MysqlToPb(),
	}, nil
}

func (s *UserRpcServer) DeleteUser(ctx context.Context, in *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
	user := mysql.NewUser()
	user.Id = in.Id
	err := user.Delete()
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResp{}, nil
}
