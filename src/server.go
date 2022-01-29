package src

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"online-store-evermos/src/controllers"
	"online-store-evermos/src/seeders"
	"os"
)

var s = controllers.Server{}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	s.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	seeders.Load(s.DB)
	s.Run(":"+os.Getenv("PORT"))
}
