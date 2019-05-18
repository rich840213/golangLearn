package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	//透過執行autoload.go的init()，所以以下註解
	/*err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}*/

	log.Println(os.Getenv("DB_PASSWORD"))
}
