package api

import (
	"context"
	pb "github.com/dinzhen12306/message/user"
	"github.com/dinzhen12306/template/server"
)

type UserRpcServer struct {
	pb.UnimplementedUserServerServer
}

func (s *UserRpcServer) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.GetUserResp, error) {
	user, err := server.GetUser(in.Where)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResp{
		UserInfo: user,
	}, nil
}

func (s *UserRpcServer) GetUsers(ctx context.Context, in *pb.GetUsersReq) (*pb.GetUsersResp, error) {
	data, err := server.GetUsers(in.Limit, in.Size)
	if err != nil {
		return nil, err
	}
	return &pb.GetUsersResp{
		UserInfo: data,
	}, nil
}

func (s *UserRpcServer) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	user, err := server.CreateUser(in.UserInfo)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResp{
		UserInfo: user,
	}, nil
}

func (s *UserRpcServer) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	user, err := server.UpdateUser(in.UserInfo)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResp{
		UserInfo: user,
	}, nil
}

func (s *UserRpcServer) DeleteUser(ctx context.Context, in *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
	err := server.DeleteUser(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResp{}, nil
}
