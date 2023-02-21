package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"ticketing/ticketing/controllers"
	"ticketing/ticketing/database"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENV Config
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load file environment")
	} else {
		fmt.Println("Successfully load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(DB)

	// Router GIN
	router := gin.Default()

	// Ticket - CRU + Get Ticket By ID
	// Customer Side
	router.GET("/tickets/:id", controllers.GetTicketById)
	// Back Office
	router.POST("/bo/tickets", controllers.CreateTicket)
	router.PUT("/bo/tickets/:id", controllers.UpdateTicket)
	router.DELETE("/bo/tickets/:id", controllers.DeleteTicket)

	// Category - CRUD + Get All Event by CategoryId
	// Customer Side
	router.GET("/categories/:id/events", controllers.GetAllCategory)
	// Back Office
	router.GET("/bo/categories", controllers.GetAllCategory)
	router.POST("/bo/categories", controllers.CreateCategory)
	router.PUT("/bo/categories/:id", controllers.UpdateCategory)
	router.DELETE("/bo/categories/:id", controllers.DeleteCategory)

	// Event - CRUD + Get All Ticket by EventId
	// Customer Side
	router.GET("/events", controllers.GetAllEvent)
	router.GET("/events/:id/tickets", controllers.GetAllTicketByEventId)
	// Back Office
	router.POST("/bo/events", controllers.CreateEvent)
	router.PUT("/bo/events/:id", controllers.UpdateEvent)
	router.DELETE("/bo/events/:id", controllers.DeleteEvent)

	// Customer - CRU (Customer side)
	router.GET("/customer/:id", controllers.GetCustomerById)
	router.POST("/customer", controllers.CreateCustomer)
	router.PUT("/customer/:id", controllers.UpdateCustomer)

	// Wallet (Customer side)
	router.GET("/wallet/:id", controllers.GetWalletInfo)
	router.PUT("/wallet/top_up", controllers.TopUpBalance)

	// Transaction - CR
	// Customer side
	router.GET("/transaction/:id", controllers.GetTransactionById)
	router.POST("/transaction", controllers.CreateTransaction)
	// get trans by cust id
	// get trans by date

	// Back Office
	// get all trans

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}

}
