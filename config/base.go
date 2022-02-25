package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"toko/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const MODE = "release"
const PORT = ":8801"

type Server struct {
	DB     *sql.DB
	Router *gin.Engine
}

func ConnectDB() (*sql.DB, error) {
	// membuat variable yang berisi format untuk koneksi dan database
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "Secret", "127.0.0.1", "3306", "enigma_toko")
	// Instance sebuah object dari sebuah struct Value yang bertipe data map[string]interface{}
	val := url.Values{}
	// Menambahkan sebuah location dan menentukan Asia/Jakarta
	val.Add("loc", "Asia/Jakarta")
	// Membuat variable yang berisikan hasil penggabungan variable connection dengan value location
	// Fungsi Encode untuk merubah menjadi sebuah string
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	// membuka koneksi
	db, err := sql.Open(`mysql`, dsn)
	// pengecekan error apakah ada error atau tidak saat buka koneksi
	if err != nil {
		log.Fatal(err)
	}
	// melakukan test ke database apakah sudah bisa digunakan dan ada pengecekan error
	if err := db.Ping(); err != nil {
		log.Print(err)
		_, _ = fmt.Scanln()
		log.Fatal(err)
	}
	log.Println("DataBase Successfully Connected")
	// mengembalikan sebuah koneksi yang sudah terbuka
	return db, err
}

func CreateRouter() *gin.Engine {
	if MODE != "release" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// menginstance object router
	r := gin.Default()
	return r
}

func InitRouter(db *sql.DB, r *gin.Engine) *Server {
	return &Server{
		DB:     db,
		Router: r,
	}
}

func (server *Server) InitializeRouter() {
	r := server.Router.Group("api/v1")
	controllers.NewBookController(server.DB, r)
	controllers.NewCategoryController(server.DB, r)
}

func Run(r *gin.Engine) {
	fmt.Println("Listen to part 8801")
	err := r.Run(PORT)
	if err != nil {
		log.Fatal(err.Error())
	}
}
