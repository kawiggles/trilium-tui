package main

import (
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error reading environmental variables")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	token := os.Getenv("TOKEN")

	fmt.Println(token)
	fmt.Println(host)
	fmt.Println(port)
}
