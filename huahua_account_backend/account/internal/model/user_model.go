package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           int       `gorm:"primarykey"`
	Username     string    `gorm:"column:username"`
	Email        string    `gorm:"column:email"`
	PasswordHash string    `gorm:"column:password_hash"`
	Openid       string    `gorm:"column:openid"`
	LastTime     time.Time `gorm:"column:last_time"`
	Avatar       string    `gorm:"column:avatar"`
}

var _ UserModel = (*customUserModel)(nil)

func NewUserModel(db *gorm.DB) UserModel {
	return &customUserModel{
		db: db,
	}
}

type (
	UserModel interface {
		FindByOpenid(context.Context, string) (*User, error)
		Insert(context.Context, *User) error
	}
	customUserModel struct {
		db *gorm.DB
	}
)

func (c *customUserModel) FindByOpenid(ctx context.Context, openid string) (*User, error) {
	var res User
	err := c.db.Where("openid = ?", openid).First(&res).Error
	if err != nil {
		return nil, ErrRecordNotFound(err)
	}
	return &res, nil
}

func (c *customUserModel) Insert(ctx context.Context, user *User) error {
	return c.db.Create(user).Error
}
