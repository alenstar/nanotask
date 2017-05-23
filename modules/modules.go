package modules

import (
	"fmt"
	"github.com/alenstar/nanoweb/config"
	"github.com/alenstar/nanoweb/log"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func getDatabaseUrl() string {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// "root:123@/test?charset=utf8"
	host := config.StringN("db_host")
	dbname := config.StringN("db_name")
	user := config.StringN("db_user")
	passwd := config.StringN("db_passwd")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, passwd, host, dbname)
}

func init() {
	var err error

	engine, err = xorm.NewEngine("mysql", getDatabaseUrl())
	if err != nil {
		log.Error("xorm.NewEngine: ", err.Error())
	} else {
		engine.ShowSQL(true)
		log.Info(engine.DBMetas())
	}
}
