package app

import (
	"github.com/alenstar/nanoweb/modules"
	"github.com/alenstar/nanoweb/utils"
	"time"
)

type ArticleInfo struct {
	Id        int64     `xorm:"unique pk autoincr" json:"-"`
	ArticleId uint64    `xorm:"unique index" json:"-"` // crc64(md5(string))
	Title     string    `xorm:"notnull varchar(63)" json:"title"`
	Author    string    `xorm:"notnull varchar(31)" json:"author"`
	Content   string    `json:"content"`
	Link      string    `xorm:"notnull varchar(63)" json:"link"`
	Created   time.Time `xorm:"created" json:"-"`
	Updated   time.Time `xorm:"updated" json:"-"`
}

func (a *ArticleInfo) RealId() uint64 {
	// xx_id 16bits
	// db_id 8bits
	// tb_id 8bits
	// user_id 32bits
	return (0x0000 << 32) | uint64(a.Id)
}

func (a *ArticleInfo) CalcArticleId() uint64 {
	// xx_id 16bits
	// db_id 8bits
	// tb_id 8bits
	// user_id 32bits
	return utils.CRC64([]byte(utils.Md5String(a.Title + a.Author)))
}

func init() {
	modules.DefaultEngine().Sync(new(ArticleInfo))
}

/*

添加一个对象：

curl -X POST -H 'Content-Type: application/json' -d 'json={"Title":"Sean Plott", "author":"alen", "link":"https://alenstar.github.io/"}' http://127.0.0.1:8080/article

返回一个相应的ArticleId:{ArticleId}

查询一个对象

curl -X GET http://127.0.0.1:8888/article?id={ArticleId}

查询全部的对象

curl -X GET http://127.0.0.1:8888/article

更新一个对象

curl -X PUT -d 'json={"Link":"https://github.com/alenstar}' http://127.0.0.1:8888/article?id={ArticleId}

删除一个对象

curl -X DELETE http://127.0.0.1:8888/article?id={ArticleId}


*/
