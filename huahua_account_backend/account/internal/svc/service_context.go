package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"huahua_account_backend/account/internal/config"
	"huahua_account_backend/account/internal/model"
)

type ServiceContext struct {
	Config           config.Config
	UserModel        model.UserModel
	BookModel        model.BookModel
	TagModel         model.TagModel
	BookJournalModel model.BookJournalModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := NewDatabase(c)
	return &ServiceContext{
		Config:           c,
		UserModel:        model.NewUserModel(db),
		BookModel:        model.NewBookModel(db),
		TagModel:         model.NewTagModel(db),
		BookJournalModel: model.NewBookJournalModel(db),
	}
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
