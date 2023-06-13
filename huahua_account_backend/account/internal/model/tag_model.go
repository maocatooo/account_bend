package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type AcTag struct {
	ID  int `gorm:"primarykey"`
	SID int `gorm:"column:sid"`
	TID int `gorm:"column:tid"`
}

type Tag struct {
	ID          int       `gorm:"primarykey"`
	Name        string    `gorm:"column:name"`
	CreatedTime time.Time `gorm:"column:created_time;default:now()"`
	Icon        string    `gorm:"column:icon"`
	IconTp      int       `gorm:"column:icon_tp"`
	Uid         int       `gorm:"column:uid"`
	Priority    int       `gorm:"column:priority"`
}

var _ TagModel = (*customTagModel)(nil)

func NewTagModel(db *gorm.DB) TagModel {
	return &customTagModel{db: db}
}

type TagModel interface {
	Insert(ctx context.Context, tag *Tag) error
	FindByUserID(ctx context.Context, uid int) ([]*Tag, error)
	InsertAcTag(ctx context.Context, at *AcTag) error
}

type customTagModel struct {
	db *gorm.DB
}

func (c *customTagModel) Insert(ctx context.Context, tag *Tag) error {
	return c.db.Create(tag).Error
}

func (c *customTagModel) FindByUserID(ctx context.Context, uid int) ([]*Tag, error) {
	var tags []*Tag
	err := c.db.Where("uid = ?", uid).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (c *customTagModel) InsertAcTag(ctx context.Context, at *AcTag) error {
	return c.db.Create(at).Error
}
