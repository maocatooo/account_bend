package model

import (
	"account_bend/account/utils/uuid"
	"context"
	"gorm.io/gorm"
	"time"
)

type TagRel struct {
	ID  string `gorm:"primarykey"`
	SID string `gorm:"column:sid"`
	TID string `gorm:"column:tid"`
}

type Tag struct {
	ID        string    `gorm:"primarykey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	Icon      string    `gorm:"column:icon"`
	IconTp    int       `gorm:"column:icon_tp"`
	Uid       string    `gorm:"column:uid"`
	Priority  int       `gorm:"column:priority"`
}

var _ TagModel = (*customTagModel)(nil)

func NewTagModel(db *gorm.DB) TagModel {
	return &customTagModel{db: db}
}

type TagModel interface {
	Create(ctx context.Context, tag *Tag) error
	Updates(ctx context.Context, tag *Tag, updates map[string]any) error
	FindByUserID(ctx context.Context, uid string) ([]*Tag, error)
	CreateTagRel(ctx context.Context, at *TagRel) error
}

type customTagModel struct {
	db *gorm.DB
}

func (c *customTagModel) Create(ctx context.Context, tag *Tag) error {
	if tag.ID == "" {
		tag.ID = uuid.New()
	}
	return c.db.Create(tag).Error
}

func (c *customTagModel) FindByUserID(ctx context.Context, uid string) ([]*Tag, error) {
	var tags []*Tag
	err := c.db.Debug().Select("tag.*, count( book_journal.tid) as ct").Model(Tag{}).
		Joins("left join  book_journal on book_journal.tid = tag.id ").
		Where("tag.uid = ? ", uid).Group("tag.id").Order("ct Desc, tag.created_at Desc").Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (c *customTagModel) Updates(ctx context.Context, tag *Tag, updates map[string]any) error {
	return c.db.Model(tag).Where("id = ?", tag.ID).Updates(updates).Error

}

func (c *customTagModel) CreateTagRel(ctx context.Context, at *TagRel) error {
	return c.db.Create(at).Error
}
