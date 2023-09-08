package model

import (
	"account_bend/account/utils/uuid"
	"context"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        string `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	Tp        int
	Uid       string
}

var _ BookModel = (*customBookModel)(nil)

func NewBookModel(db *gorm.DB) BookModel {
	return &customBookModel{db: db}
}

type (
	BookModel interface {
		FindByUid(context.Context, string) ([]*Book, error)
		Create(context.Context, *Book) error
	}
	customBookModel struct {
		db *gorm.DB
	}
)

func (c *customBookModel) FindByUid(ctx context.Context, uid string) ([]*Book, error) {
	var res []*Book
	err := c.db.Where("uid = ?", uid).Find(&res).Error
	return res, err
}

func (c *customBookModel) Create(ctx context.Context, Book *Book) error {
	Book.ID = uuid.New()
	return c.db.Create(Book).Error
}
