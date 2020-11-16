package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/patoui/test-fiber/database"
	"github.com/patoui/test-fiber/src/book"
	"github.com/patoui/test-fiber/src/home"
)

func setupRoutes(app *fiber.App) {
	// Assets
	app.Static("/static", "./static")

	// Views
	app.Get("/", home.Index)

	// API
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()
}
