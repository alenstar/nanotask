package main

import (
	"fmt"
	"goworker/modules"
	"goworker/utils"
	"time"
)

type DiscussInfo struct {
	Id        int64     `xorm:"unique index pk autoincr(1)" json:"-"`
	ArticleId uint64    `xorm:"notnull unique index" json:"article_id"` // crc64(md5(string))
	DiscussId uint64    `xorm:"notnull unique index" json:"discuss_id"`
	Content   string    `xorm:"notnull" json:"content"`
	UserId    uint64    `xorm:"notnull" json:"user_id"`  // Relate to UserInfo
	ReplyId   uint64    `xorm:"notnull" json:"reply_id"` // reply_id is discuss_id
	Created   time.Time `xorm:"notnull created" json:"created"`
	Updated   time.Time `xorm:"notnull updated" json:"updated"`
}

func (d *DiscussInfo) RealId() uint64 {
	// xx_id 16bits
	// db_id 8bits
	// tb_id 8bits
	// user_id 32bits
	return (0x0000 << 32) | uint64(d.Id)
}
func (d *DiscussInfo) CalcDiscussId() uint64 {
	return uint64(utils.CRC32([]byte(utils.Md5String(fmt.Sprintf("%d@%d on %d", d.UserId, d.ArticleId, time.Now().Unix())))))
}

func init() {
	modules.DefaultEngine().Sync(new(DiscussInfo))
}

/*

添加一个对象：

curl -X POST -H 'Content-Type: application/json' -d '{"Name":"alen","Password":"Sean Plott", "Email":"alen@taobao.com"}' http://127.0.0.1:8888/discuss

返回一个相应的Id:{Id}

查询一个对象

curl -X GET http://127.0.0.1:8888/discuss/{Name}

查询全部的对象

curl -X GET http://127.0.0.1:8888/discuss

更新一个对象

curl -X PUT -H 'Content-Type: application/json' -d '{"link":"https://github.com/alenstar}' http://127.0.0.1:8888/discuss?id={Name}

删除一个对象

curl -X DELETE http://127.0.0.1:8888/discuss?id={Name}


*/
