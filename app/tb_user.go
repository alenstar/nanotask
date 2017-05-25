package app

import (
	"github.com/alenstar/nanoweb/modules"
	"time"
)

type UserInfo struct {
	Id       int64     `xorm:"unique pk autoincr(1001)"`
	Name     string    `xorm:"unique index varchar(31)"`
	NickName string    `xorm:"notnull varchar(31)"`
	Password string    `xorm:"notnull varchar(63)"`
	Email    string    `xorm:"unique index notnull varchar(63)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

func (u *UserInfo) RealId() uint64 {
	// xx_id 16bits
	// db_id 8bits
	// tb_id 8bits
	// user_id 32bits
	return (0x0000 << 32) | uint64(u.Id)
}

func init() {
	modules.DefaultEngine().Sync(new(UserInfo))
}

/*

添加一个对象：

curl -X POST -d 'json={"Name":"alen","Password":"Sean Plott", "Email":"alen@taobao.com"}' http://127.0.0.1:8080/user

返回一个相应的Id:{Id}

查询一个对象

curl -X GET http://127.0.0.1:8888/user/{Name}

查询全部的对象

curl -X GET http://127.0.0.1:8888/user

更新一个对象

curl -X PUT -d '{json="Link":"https://github.com/alenstar}' http://127.0.0.1:8888/user?id={Name}

删除一个对象

curl -X DELETE http://127.0.0.1:8888/user?id={Name}


*/
