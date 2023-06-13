package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type AcBookJournal struct {
	Id        int       `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()"`
	Tid       int       `gorm:"column:tid"`
	Tname     string    `gorm:"column:tname"`

	Date     time.Time `gorm:"column:date"`
	Amount   int       `gorm:"column:amount"`
	Notes    string    `gorm:"column:notes"`
	AcBookId int       `gorm:"column:ac_book_id"`
	Uid      int       `gorm:"column:uid"`
}

var _ AcBookJournalModel = (*customAcBookJournalModel)(nil)

func NewAcBookJournalModel(db *gorm.DB) AcBookJournalModel {
	return &customAcBookJournalModel{
		db: db,
	}
}

type (
	AcBookJournalModel interface {
		Create(context.Context, *AcBookJournal) error
		FindByBookId(context.Context, int) ([]*AcBookJournal, error)
	}
	customAcBookJournalModel struct {
		db *gorm.DB
	}
)

func (c *customAcBookJournalModel) Create(ctx context.Context, acBookJournal *AcBookJournal) error {
	return c.db.Create(acBookJournal).Error
}

func (c *customAcBookJournalModel) FindByBookId(ctx context.Context, bookId int) ([]*AcBookJournal, error) {
	var journals []*AcBookJournal
	err := c.db.Where("ac_book_id = ?", bookId).Order("date desc").Find(&journals).Error
	if err != nil {
		return nil, err
	}
	return journals, nil
}
