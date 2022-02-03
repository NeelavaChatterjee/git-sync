package database

import (
	"log"
	"time"

	"github.com/NeelavaChatterjee/git-sync/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	// conn_str := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	time.Sleep(time.Second * 30)
	conn_str := "root:1234@tcp(mysql-db)/git_sync"
	db, err := gorm.Open(mysql.Open(conn_str), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Database connected!!")
	Db = db
	db.AutoMigrate(&models.Track{}, &models.CommitHistory{}, &models.PollLogs{})
}
