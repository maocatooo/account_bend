package model

import (
	"context"
	"gorm.io/gorm"
	"huahua_account_backend/account/utils/uuid"
	"time"
)

type BookJournal struct {
	ID        string    `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	Tid       string    `gorm:"column:tid"`
	Tname     string    `gorm:"column:tname"`

	Date   time.Time `gorm:"column:date;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	Amount int       `gorm:"column:amount"`
	Record string    `gorm:"column:record"`
	BookID string    `gorm:"column:book_id"`
	Uid    string    `gorm:"column:uid"`
}

var _ BookJournalModel = (*customBookJournalModel)(nil)

func NewBookJournalModel(db *gorm.DB) BookJournalModel {
	return &customBookJournalModel{
		db: db,
	}
}

type (
	BookJournalModel interface {
		Create(context.Context, *BookJournal) error
		FindByBookId(context.Context, string) ([]*BookJournal, error)
	}
	customBookJournalModel struct {
		db *gorm.DB
	}
)

func (c *customBookJournalModel) Create(ctx context.Context, m *BookJournal) error {
	if m.ID == "" {
		m.ID = uuid.New()
	}
	return c.db.Create(m).Error
}

func (c *customBookJournalModel) FindByBookId(ctx context.Context, bookId string) ([]*BookJournal, error) {
	var journals []*BookJournal
	err := c.db.Where("book_id = ?", bookId).Order("created_at desc").Find(&journals).Error
	if err != nil {
		return nil, err
	}
	return journals, nil
}
