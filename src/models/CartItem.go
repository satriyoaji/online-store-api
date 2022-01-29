package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type CartItem struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	IDUser    uint64    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;size:25; not null" json:"id_user"`
	User	  User		`gorm:"foreignKey:IDUser"`
	IDItem    uint64    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;size:25; not null" json:"id_item"`
	Item	  Item		`gorm:"foreignKey:IDItem"`
	Qty       uint64    `gorm:"size:25; not null" json:"qty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *CartItem) Prepare() {
	p.ID = 0
	p.IDUser = 0
	p.IDItem = 0
	p.Qty = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *CartItem) Validate() error {
	if p.IDItem < 1 {
		return errors.New("required ID Item")
	}
	if p.IDUser < 1 {
		return errors.New("required ID User")
	}
	if p.Qty < 1 {
		return errors.New("required quantity")
	}
	return nil
}

func (p *CartItem) AddToCart(db *gorm.DB) (*CartItem, error) {
	item := &Item{}
	it, err := item.FindItemByID(db, p.IDItem)
	if err != nil {
		return &CartItem{}, err
	}

	user := &User{}
	_, err = user.FindUserById(db, p.IDUser)
	if err != nil {
		return &CartItem{}, err
	}

	if it.Stock < 1 {
		return &CartItem{}, errors.New("stock is 0")
	}

	err = db.Debug().Model(&CartItem{}).Create(&p).Error

	if err != nil {
		return &CartItem{}, err
	}

	_, err = item.DecreaseItemQty(db, p.IDItem)
	if err != nil {
		return &CartItem{}, err
	}

	return p, nil
}

func (p *CartItem) FindAllCartByUID(db *gorm.DB, uid uint64) (*[]CartItem, error) {
	var carts []CartItem
	err := db.Debug().Model(&CartItem{}).Where("id_user = ?", uid).Limit(100).Find(&carts).Error
	if err != nil {
		return &[]CartItem{}, err
	}

	return &carts, nil
}
