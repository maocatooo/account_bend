package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type AcBook struct {
	ID        int `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time `gorm:"column:created_at;default:now()"`
	Tp        int
	Uid       int
}

var _ AcBookModel = (*customAcBookModel)(nil)

func NewAcBookModel(db *gorm.DB) AcBookModel {
	return &customAcBookModel{db: db}
}

type (
	AcBookModel interface {
		FindByUid(context.Context, int) ([]*AcBook, error)
		Insert(context.Context, *AcBook) error
	}
	customAcBookModel struct {
		db *gorm.DB
	}
)

func (c *customAcBookModel) FindByUid(ctx context.Context, uid int) ([]*AcBook, error) {
	var res []*AcBook
	err := c.db.Where("uid = ?", uid).Find(&res).Error
	return res, err
}

func (c *customAcBookModel) Insert(ctx context.Context, acBook *AcBook) error {
	return c.db.Create(acBook).Error
}
