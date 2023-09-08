package model

import (
	"account_bend/account/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"testing"
)

var db *gorm.DB

func TestMain(t *testing.M) {

	os.Exit(t.Run())
}

func TestEmpty(t *testing.T) {
	conf := config.Config{}
	conf.Mysql.DataSource = "root:123456@tcp(192.168.163.121:3306)/account_bend?parseTime=true"
	db = NewDatabase(conf)

	db.AutoMigrate(
		&User{},
		&Book{},
		&Tag{},
		&BookJournal{},
		&TagRel{},
	)

}

func NewDatabase(cfg config.Config) *gorm.DB {
	dsn := cfg.Mysql.DataSource
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		panic(err)
	}
	return db
}
