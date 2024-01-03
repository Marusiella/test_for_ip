package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

type Request struct {
	gorm.Model
	Ip        string
	Timestamp time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Request{})

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		db.Create(&Request{Ip: c.IP(), Timestamp: time.Now()})
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":5454"))
}
