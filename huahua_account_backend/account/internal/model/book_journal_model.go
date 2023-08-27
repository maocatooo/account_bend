package model

import (
	"context"
	"gorm.io/gorm"
	"huahua_account_backend/account/internal/types"
	"huahua_account_backend/account/utils/uuid"
	"time"
)

type BookJournal struct {
	ID        string    `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	Tid       string    `gorm:"column:tid"`
	Tname     string    `gorm:"column:tname"`
	Name      string    `gorm:"column:name"`

	Date   time.Time `gorm:"column:date;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	Amount string    `gorm:"column:amount"`
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
		Update(ctx context.Context, m *BookJournal) error
		DeleteById(context.Context, string) error
		FindByBookId(context.Context, string) ([]*BookJournal, error)
		FindByTypes(context.Context, *types.Journal) ([]*BookJournal, error)
		NamePrompt(ctx context.Context, uid, tid string) ([]string, error)
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

func (c *customBookJournalModel) Update(ctx context.Context, m *BookJournal) error {
	if m.ID == "" {
		m.ID = uuid.New()
	}
	return c.db.Model(m).Updates(m).Error
}

func (c *customBookJournalModel) DeleteById(ctx context.Context, id string) error {
	if id == "" {
		return nil
	}
	return c.db.Where("id = ?", id).Delete(&BookJournal{}).Error
}

func (c *customBookJournalModel) FindByBookId(ctx context.Context, bookId string) ([]*BookJournal, error) {
	var journals []*BookJournal
	err := c.db.Where("book_id = ?", bookId).Debug().Order("date desc").Find(&journals).Error
	if err != nil {
		return nil, err
	}
	return journals, nil
}

func (c *customBookJournalModel) FindByTypes(ctx context.Context, tps *types.Journal) ([]*BookJournal, error) {
	var journals []*BookJournal
	db := c.db
	db = db.Where("book_id = ?", tps.BookID)
	if tps.Date != 0 {
		d := time.UnixMilli(int64(tps.Date))
		d1 := d.AddDate(0, 1, 0)
		db = db.Where("date >= ? and date < ?", d, d1)
	}
	err := db.Debug().Order("date desc").Find(&journals).Error
	if err != nil {
		return nil, err
	}
	return journals, nil
}

func (c *customBookJournalModel) NamePrompt(ctx context.Context, uid, tid string) ([]string, error) {
	var res []struct {
		Name string
	}
	db := c.db.Model(&BookJournal{}).Where("uid = ?", uid)
	if tid != "" {
		db = db.Where("tid = ?", tid)
	}
	err := db.Select("name, count(name) as c").Group("name").Order("c desc").Limit(10).Find(&res).Error
	if err != nil {
		return nil, err
	}
	var s []string
	for _, item := range res {
		s = append(s, item.Name)
	}

	return s, nil
}
