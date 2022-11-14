package main

import (
	"log"

	"github.com/AshiishKarhade/GO-Projects/go-fiber/database"
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
}

func setupRoutes(app *fiber.App) {
	log.Println("Setting up routes...")
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}
