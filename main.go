package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
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

func createDsn(driver string) string {
	host := os.Getenv("DB_HOST")

	if driver == "sqlite3" || driver == "sqlite" {
		return host
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	protocol := os.Getenv("DB_PROTOCOL")
	if protocol == "" {
		protocol = "tcp"
	}
	database := os.Getenv("DB_DATABASE")

	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, protocol, host, port, database)
}

func initDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	driver := os.Getenv("DB_DRIVER")
	database.DBConn, err = gorm.Open(driver, createDsn(driver))
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
