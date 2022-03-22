package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//DBConnection -> return db instance
func DBConnection() (*gorm.DB, error) {
	USER := os.Getenv("GO_USER")
	PASS := os.Getenv("GO_PASS")
	HOST := os.Getenv("GO_HOST")
	PORT := os.Getenv("GO_PORT")
	DBNAME := os.Getenv("GO_DBNAME")
	SSLMODE := os.Getenv("GO_SSLMODE")
	TZ := os.Getenv("GO_TZ")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	// mysql
	//"gorm.io/driver/mysql"
	//url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	//return gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})

	// postgresql
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", HOST, USER, PASS, DBNAME, PORT, SSLMODE, TZ)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
}
