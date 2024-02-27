package server

import (
	pb "github.com/dinzhen12306/message/user"
	"github.com/dinzhen12306/template/model/mysql"
)

func GetUser(where map[string]string) (*pb.User, error) {
	inMap := make(map[string]interface{})
	for k, v := range where {
		inMap[k] = v
	}
	user := mysql.NewUser()
	err := user.Get(inMap)
	if err != nil {
		return nil, err
	}
	return user.MysqlToPb(), nil
}
func GetUsers(limit, offset int64) ([]*pb.User, error) {
	user := mysql.NewUser()
	users, err := user.GetUsers(limit, offset)
	if err != nil {
		return nil, err
	}
	var data []*pb.User
	for _, v := range users {
		data = append(data, v.MysqlToPb())
	}
	return data, nil
}
func CreateUser(UserInfo *pb.User) (*pb.User, error) {
	user := mysql.PbToMysql(UserInfo)
	err := user.Create()
	if err != nil {
		return nil, err
	}
	return user.MysqlToPb(), nil
}
func UpdateUser(UserInfo *pb.User) (*pb.User, error) {
	user := mysql.PbToMysql(UserInfo)
	err := user.Update()
	if err != nil {
		return nil, err
	}
	return user.MysqlToPb(), nil
}
func DeleteUser(Id int64) error {
	user := mysql.NewUser()
	user.Id = Id
	err := user.Delete()
	if err != nil {
		return err
	}
	return nil
}
