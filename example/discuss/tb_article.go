package main

import (
	"github.com/alenstar/nanotask/modules"
	"github.com/alenstar/nanotask/utils"
	"time"
)

type ArticleInfo struct {
	Id        int64     `xorm:"unique index pk autoincr" json:"-"`
	ArticleId uint64    `xorm:"notnull unique index" json:"article_id"` // crc64(md5(string))
	UserId    uint64    `xorm:"notnull" json:"user_id"`                 // Relate to UserInfo
	Title     string    `xorm:"notnull varchar(63)" json:"title"`
	Author    string    `xorm:"notnull varchar(31)" json:"author"`
	Content   string    `xorm:"notnull" json:"content"`
	Link      string    `xorm:"notnull varchar(63)" json:"link"`
	Created   time.Time `xorm:"notnull created" json:"created"`
	Updated   time.Time `xorm:"notnull updated" json:"updated"`
}

func (a *ArticleInfo) RealId() uint64 {
	// xx_id 16bits
	// db_id 8bits
	// tb_id 8bits
	// user_id 32bits
	return (0x0000 << 32) | uint64(a.Id)
}

func (a *ArticleInfo) CalcArticleId() uint64 {
	return uint64(utils.CRC32([]byte(utils.Md5String(a.Title + a.Author))))
}

func init() {
	modules.DefaultEngine().Sync(new(ArticleInfo))
}

/*

添加一个对象：

curl -X POST -H 'Content-Type: application/json' -d '{"title":"Sean Plott", "author":"alen", "link":"https://alenstar.github.io/"}' http://127.0.0.1:8888/article

返回一个相应的ArticleId:{ArticleId}

查询一个对象

curl -X GET http://127.0.0.1:8888/article?id={ArticleId}

查询全部的对象

curl -X GET http://127.0.0.1:8888/article

更新一个对象

curl -X PUT -H 'Content-Type: application/json' -d '{"link":"https://github.com/alenstar}' http://127.0.0.1:8888/article?id={ArticleId}

删除一个对象

curl -X DELETE http://127.0.0.1:8888/article?id={ArticleId}


*/
