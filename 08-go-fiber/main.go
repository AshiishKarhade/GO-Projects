package main

import (
	"github.com/jinzhu/gorm"
	"log"

	"github.com/AshiishKarhade/GO-Projects/go-fiber/database"
	"github.com/AshiishKarhade/GO-Projects/go-fiber/lead"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	initDB()
	defer database.DBConn.Close()
	setupRoutes(app)
	app.Listen(3000)
}

func initDB() {
	log.Println("Initialising DB...")
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect db")
	}
	log.Println("DB Connection successfully!")
	database.DBConn.AutoMigrate(&lead.Lead{})
	log.Println("Database migration successful")
}

func setupRoutes(app *fiber.App) {
	log.Println("Setting up routes...")
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}
