package app

import (
	"github.com/alenstar/nanoweb/modules"
	"time"
)

type DiscussleInfo struct {
	Id        int64     `xorm:"unique pk autoincr(1)"`
	ArticleId uint64    `xorm:"unique index"` // crc64(md5(string))
	Content   string    `xorm:"notnull"`
	UserId    uint32    `xorm:"notnull"` // Relate to UserInfo
	ReplyId   uint32    //
	Created   time.Time `xorm:"created"`
	Updated   time.Time `xorm:"updated"`
}

func (d *DiscussleInfo) RealId() uint64 {
	// xx_id 16bits
	// db_id 8bits
	// tb_id 8bits
	// user_id 32bits
	return (0x0000 << 32) | uint64(d.Id)
}

func init() {
	modules.DefaultEngine().Sync(new(DiscussleInfo))
}
