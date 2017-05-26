package app

import (
	"github.com/alenstar/nanoweb/modules"
	"github.com/alenstar/nanoweb/utils"
	"time"
)

type UserInfo struct {
	Id       int64     `xorm:"unique index pk autoincr(1001)"`
	UserId   uint64    `xorm:"notnull unique index autoincr(1001)"`
	Name     string    `xorm:"notnull unique index varchar(31)"`
	NickName string    `xorm:"notnull varchar(31)"`
	Password string    `xorm:"notnull varchar(63)"`
	Email    string    `xorm:"notnull unique index varchar(63)"`
	Created  time.Time `xorm:"notnull created"`
	Updated  time.Time `xorm:"notnull updated"`
}

func (u *UserInfo) RealId() uint64 {
	// xx_id 16bits
	// db_id 8bits
	// tb_id 8bits
	// user_id 32bits
	return (0x0000 << 32) | uint64(u.Id)
}

func (u *UserInfo) CalcUserId() uint64 {
	return utils.CRC64([]byte(utils.Md5String(u.Name + u.Email)))
}

func init() {
	modules.DefaultEngine().Sync(new(UserInfo))
}

/*

添加一个对象：

curl -X POST -H 'Content-Type: application/json' -d '{"Name":"alen","Password":"Sean Plott", "Email":"alen@taobao.com"}' http://127.0.0.1:8080/user

返回一个相应的Id:{Id}

查询一个对象

curl -X GET http://127.0.0.1:8888/user/{Name}

查询全部的对象

curl -X GET http://127.0.0.1:8888/user

更新一个对象

curl -X PUT -H 'Content-Type: application/json' -d '{"link":"https://github.com/alenstar}' http://127.0.0.1:8888/user?id={Name}

删除一个对象

curl -X DELETE http://127.0.0.1:8888/user?id={Name}


*/
