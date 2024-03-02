package mysql

import (
	"github.com/dinzhen12306/message/user"
	"github.com/dinzhen12306/template/golbe"
	"time"
)

type User struct {
	Id       int64  `xorm:"pk autoincr"`
	Username string `xorm:"unique"`
	Password string
	Phone    string
	Email    string
	Sex      int64
	Age      int64
	Created  time.Time  `xorm:"created"`
	Updated  time.Time  `xorm:"updated"`
	Deleted  *time.Time `xorm:"deleted"`
}

func NewUser() *User {
	return new(User)
}

func (u *User) Create() (err error) {
	_, err = golbe.XDB.Insert(u)
	return
}
func (u *User) Get(where map[string]interface{}) (err error) {
	_, err = golbe.XDB.Where(where).Get(u)
	return
}

func (u *User) GetUsers(limit, offset int64) (users []User, err error) {
	err = golbe.XDB.Limit(int(limit), int(offset)).Find(&users)
	return
}

func (u *User) Update() (err error) {
	_, err = golbe.XDB.Where("id = ?", u.Id).Update(u)
	return
}
func (u *User) Delete() (err error) {
	_, err = golbe.XDB.Table(u).Where("id = ?", u.Id).Delete()
	return
}
func (u *User) MysqlToPb() *user.User {
	return &user.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
		Phone:    u.Phone,
		Email:    u.Email,
		Sex:      user.Sex(u.Sex),
		Age:      u.Age,
	}
}
func PbToMysql(users *user.User) *User {
	return &User{
		Id:       users.Id,
		Username: users.Username,
		Password: users.Password,
		Phone:    users.Phone,
		Email:    users.Email,
		Sex:      int64(users.Sex),
		Age:      users.Age,
	}
}
