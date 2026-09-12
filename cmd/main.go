package main

import (
	"log"

	"CarSellsProject/internal/api"
)

func main() {
	if err := api.InitServer(); err != nil {
		log.Fatal(err)
	}
}
