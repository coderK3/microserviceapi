package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

const (
	host     = "testmumbai.caogoix5ad9c.ap-south-1.rds.amazonaws.com"
	port     = 5432
	user     = "testmumbai"
	password = "test)(*098"
	dbname   = "vigyaa1"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	backup(db)

	// db, err := gorm.Open("sqlite3", "../test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// db.AutoMigrate(&Product{})

	// db.Create(&Product{Code: "kunak", Price: 2000})
	// // fmt.Println(&Product.Code)
	// fmt.Println(db.Table("products").Rows())
	// fmt.Println(db.Find(&product))
}

func backup(db *sql.DB) {
	file, err := os.Create("../main.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sqlCopyStatement := `COPY (SELECT * FROM vigyaa ) TO STDOUT`

	_, err := db.CopyTo(file, sqlCopyStatement)
	if err != nil {
		log.Fatal(err)

	}

}
