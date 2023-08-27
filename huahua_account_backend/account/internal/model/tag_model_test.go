package model

import (
	"context"
	"huahua_account_backend/account/internal/config"
	"testing"
)

func Test_customTagModel_FindByUserID(t *testing.T) {
	conf := config.Config{}
	conf.Mysql.DataSource = "root:zk246422@163.com@tcp(maocat.cc:3306)/huahua_account_backend?parseTime=true"
	c := &customTagModel{
		db: NewDatabase(conf),
	}
	c.FindByUserID(context.Background(), "3c85a07e9e8f423ab99303f45bde9d92")
}
