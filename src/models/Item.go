package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Item struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Price     uint64    `gorm:"size:25;not null" json:"price"`
	Stock     uint64    `gorm:"size:25;not null" json:"stock"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Item) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Price = 0
	p.Stock = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Item) Validate() error {
	if p.Title == "" {
		return errors.New("required title")
	}
	if p.Price < 1 {
		return errors.New("required price more than 0")
	}
	if p.Stock < 1 {
		return errors.New("required stock more than 0")
	}
	return nil
}

func (p *Item) CreateItem(db *gorm.DB) (*Item, error) {
	err := db.Debug().Model(&Item{}).Create(&p).Error

	if err != nil {
		return &Item{}, err
	}

	return p, nil
}

func (p *Item) FindAllItem(db *gorm.DB) (*[]Item, error) {
	var posts []Item
	err := db.Debug().Model(&Item{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Item{}, err
	}

	return &posts, nil
}

func (p *Item) FindItemByID(db *gorm.DB, id uint64) (*Item, error) {
	err := db.Debug().Model(&Item{}).Where("id = ?", id).Take(&p).Error
	if err != nil {
		return &Item{}, err
	}

	return p, nil
}

func (p *Item) UpdateItem(db *gorm.DB) (*Item, error) {
	err := db.Debug().Model(&Item{}).Where("id = ?", p.ID).Updates(Item{
		Title:     p.Title,
		Stock:     p.Stock,
		Price:     p.Price,
		UpdatedAt: time.Time{},
	}).Error

	if err != nil {
		return &Item{}, err
	}

	return p, nil
}

func (p *Item) DeleteItem(db *gorm.DB, id uint64) (int64, error) {
	db = db.Debug().Model(&Item{}).Where("id = ?", id).Take(&Item{}).Delete(&Item{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("post not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

func (p *Item) DecreaseItemQty(db *gorm.DB, id uint64) (*Item, error) {
	err := db.Debug().Model(&Item{}).Where("id = ?", id).Updates(Item{
		Stock:     p.Stock - 1,
		UpdatedAt: time.Time{},
	}).Error

	if err != nil {
		return &Item{}, err
	}

	return p, nil
}
