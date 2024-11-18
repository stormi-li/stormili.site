package database

import (
	"comment/models"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiserd-v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlConfigName = "mysql"
var redisAddr = "118.25.196.166:3934"
var redisPassword = "12982397StrongPassw0rd"
var DB *gorm.DB

func InitDatabase() {
	var err error

	address, data := omiserd.NewClient(&redis.Options{Addr: redisAddr, Password: redisPassword}, omiserd.Config).NewDiscover().DiscoverByWeight(MysqlConfigName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", data["username"], data["password"], address, data["database"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		DB = db
	} else {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&models.Comment{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database connected and migrated successfully.")
}
