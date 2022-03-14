package main

import (
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"

	"log"
	"toko/config"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, err := config.ConnectDB()

	// apabila ada error
	if err != nil {
		log.Fatal(err.Error())
	}

	r := config.CreateRouter()

	/*
		mekanisme untuk memberi tahu browser,
		apakah sebuah request yang di-dispatch dari aplikasi web domain lain atau origin lain,
		ke aplikasi web kita itu diperbolehkan atau tidak.
		Jika aplikasi kita tidak mengijinkan maka akan muncul error,
		dan request pasti digagalkan oleh browser. */
	r.Use(cors.AllowAll())

	/*
		Init Router digunakan untuk mengisi value Server struct
		Initial Routes digunakan untuk membuat router group
	*/
	config.InitRouter(db, r).InitializeRouter()

	config.Run(r)
}
