package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/manveru/faker"
)

type Product struct {
	Id              int       `gorm:"column:Id"`
	Name            string    `gorm:"column:Name"`
	Price           float64   `gorm:"column:Price"`
	LastUpdatedItem time.Time `gorm:"column:LastUpdatedItem"`
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

	user := "postgres"
	password := "1234"
	host := "localhost"
	port := "5433"
	_db := "trendyol_case"
	conStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, _db, password)

	db, err := gorm.Open("postgres", conStr)
	if err != nil {
		panic(err)
	}

	return db.Debug()
}
