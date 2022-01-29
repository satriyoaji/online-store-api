package seeders

import (
	//"gorm.io/gorm"
	"github.com/jinzhu/gorm"
	"log"
	"online-store-evermos/src/models"
)

var users = []models.User{
	{
		Nickname: "Satriyo Aji",
		Email:    "satriyoaji27@gmail.com",
		Password: "password",
	},
	{
		Nickname: "Arkham Cranio",
		Email:    "satriyoaji27@it.student.pens.ac.id",
		Password: "password",
	},
}

var items = []models.Item{
	{
		Title: "Item 1",
		Stock: 10,
		Price: 26000,
	},
	{
		Title: "Item 2",
		Stock: 5,
		Price: 31000,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}, &models.Item{}, &models.CartItem{}).Error
	if err != nil {
		log.Fatalf("cannot drop table : %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Item{}, &models.CartItem{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.CartItem{}).AddForeignKey("id_user", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.CartItem{}).AddForeignKey("id_item", "items(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i := range items {
		err = db.Debug().Model(&models.Item{}).Create(&items[i]).Error
		if err != nil {
			log.Fatalf("cannot seed items table: %v", err)
		}
	}
}
