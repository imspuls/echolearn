package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB //database

// const projectDirName = "echolearn"

func init() {
	// re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	// cwd, _ := os.Getwd()
	// rootPath := re.Find([]byte(cwd))

	// e := godotenv.Load(string(rootPath) + `/.env`) //Load .env file
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName) //Build connection string

	conn, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Print(err)
	}

	Db = conn
}

//GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return Db
}
