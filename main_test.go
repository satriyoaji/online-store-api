package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"online-store-evermos/src/seeders"
	"os"
	"testing"

	"online-store-evermos/src/controllers"
)

var s = controllers.Server{}

func TestConnection(t *testing.T) {
	//t.Parallel()

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

func TestCreateNewOrder(t *testing.T) {
	fmt.Println("-- Testing Create new order --")

	values := map[string]int{"id_user" : 1, "id_item" : 1, "qty" : 1}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatalf("Error getting env, not comming through %v", errEnv)
	} else {
		fmt.Println("We are getting the env values")
	}

	resp, err := http.Post("http://localhost:"+os.Getenv("PORT")+"/orders", "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	var responseBody map[string]interface{}
	var jsonString []byte
	json.NewDecoder(resp.Body).Decode(&responseBody)
	// convert map to json
	if responseBody["data"] != nil {
		jsonString, _ = json.Marshal(responseBody["data"])
	} else {
		jsonString, _ = json.Marshal(responseBody["error"])
	}
	fmt.Println("Response JSON: ")
	fmt.Println(string(jsonString))

}
