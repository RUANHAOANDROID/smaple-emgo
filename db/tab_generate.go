package db

import (
	"fmt"
	"github.com/gohouse/converter"
	"os"
)

func Generate() {
	os.ReadFile("sql/table.sql")
	err := converter.NewTable2Struct().
		SavePath("./model.go").
		Dsn("用户名:密码@tcp(IP:端口号)/数据库名?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		Table("表名").
		Run()
	fmt.Println(err)

}
