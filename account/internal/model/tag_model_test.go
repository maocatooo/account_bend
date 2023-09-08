package model

import (
	"account_bend/account/internal/config"
	"context"
	"testing"
)

func Test_customTagModel_FindByUserID(t *testing.T) {
	conf := config.Config{}
	conf.Mysql.DataSource = "root:zk246422@163.com@tcp(maocat.cc:3306)/account_bend?parseTime=true"
	c := &customTagModel{
		db: NewDatabase(conf),
	}
	c.FindByUserID(context.Background(), "3c85a07e9e8f423ab99303f45bde9d92")
}
