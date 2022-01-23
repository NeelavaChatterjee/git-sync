package config

import (
	"fmt"
	"log"
	"os"

	"github.com/NeelavaChatterjee/git-sync/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	conn_str := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := gorm.Open(mysql.Open(conn_str), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}

	db.AutoMigrate(&models.Track{}, &models.CommitHistory{}, &models.PollLogs{})
	return db
}
