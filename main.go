package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/manveru/faker"
)

type Product struct {
	Id              int       `gorm:"column:Id"`
	Name            string    `gorm:"column:Name"`
	Price           float64   `gorm:"column:Price"`
	LastUpdatedTime time.Time `gorm:"column:LastUpdatedTime"`
}

func (Product) TableName() string {
	return "Products"
}

func main() {
	db := DB()
	db.AutoMigrate(&Product{})
	fake, err := faker.New("en")
	if err != nil {
		log.Fatal(err)
	}
	productModels := db.Model(&Product{})
	for i := 0; i < 1000000; i++ {
		productModels.Create(&Product{
			Id:    i,
			Name:  fake.Name(),
			Price: float64(i * rand.Intn(100))})
		fmt.Println(fmt.Sprintf("Created Product %v", i))
	}
}

func DB() *gorm.DB {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	_db := os.Getenv("DB")
	conStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, _db, password)

	db, err := gorm.Open(os.Getenv("DB_TYPE"), conStr)
	if err != nil {
		panic(err)
	}

	return db.Debug()
}
