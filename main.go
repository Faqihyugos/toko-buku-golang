package main

import (
	"log"
	"toko/config"
)

func main() {
	db,err := config.ConnectDB()

	// apabila ada error
	if err != nil {
		log.Fatal(err.Error())
	}

	r := config.CreateRouter()

	config.InitRouter(db, r).InitializeRouter()

	config.Run(r)
}

