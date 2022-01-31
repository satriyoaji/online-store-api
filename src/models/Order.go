package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	helper "online-store-evermos/src/helpers"
	"time"
)

type Order struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	IDUser    uint64    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;size:25; not null" json:"id_user"`
	//User	  User		`gorm:"foreignKey:IDUser"`
	IDItem    uint64    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;size:25; not null" json:"id_item"`
	//Item	  Item		`gorm:"foreignKey:IDItem"`
	Qty       uint64    `gorm:"size:25; not null" json:"qty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Order) Prepare() {
	p.ID = 0
	p.IDUser = 0
	p.IDItem = 0
	p.Qty = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Order) Validate() error {
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

func (p *Order) AddOrder(db *gorm.DB) (*Order, error) {
	tx := db.Begin()
	defer helper.CommitOrRollback(tx)

	item := &Item{}
	it, err := item.FindItemByID(db, p.IDItem)
	if err != nil {
		return &Order{}, err
	}
	if it.Stock < 1 || it.Stock <= p.Qty {
		return &Order{}, errors.New("stock is unavailable !")
	}

	user := &User{}
	_, err = user.FindUserById(db, p.IDUser)
	if err != nil {
		return &Order{}, err
	}

	// reduce item stock depends on order qty
	_, err = item.DecreaseItemQty(db, p.IDItem, p.Qty)
	if err != nil {
		return &Order{}, err
	}

	// create new Order
	err = db.Debug().Model(&Order{}).Create(&p).Error
	if err != nil {
		fmt.Println("errr: ", err.Error())
		return &Order{}, err
	}

	return p, nil
}

func (p *Order) FindAllOrderByUID(db *gorm.DB, uid uint64) (*[]Order, error) {
	var orders []Order
	err := db.Debug().Model(&Order{}).Where("id_user = ?", uid).Limit(100).Find(&orders).Error
	if err != nil {
		return &[]Order{}, err
	}

	return &orders, nil
}
