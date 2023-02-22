package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"ticketing/ticketing/controllers"
	"ticketing/ticketing/database"
	"ticketing/ticketing/middleware"

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

	// Registration
	router.POST("/registration", controllers.CreateCustomer)

	// Login
	router.POST("/login", controllers.Login)

	// Ticket - CRU + Get Ticket By ID
	// Customer Side
	router.GET("/tickets/:id", middleware.VerifyJWT, controllers.GetTicketById)
	// Back Office
	router.POST("/bo/tickets", middleware.VerifyJWT, middleware.BackOffice, controllers.CreateTicket)
	router.PUT("/bo/tickets/:id", middleware.VerifyJWT, middleware.BackOffice, controllers.UpdateTicket)
	router.DELETE("/bo/tickets/:id", middleware.VerifyJWT, middleware.BackOffice, controllers.DeleteTicket)

	// Category - CRUD + Get All Event by CategoryId
	// Customer Side
	router.GET("/categories/:id/events", middleware.VerifyJWT, controllers.GetAllEventByCategory)
	router.GET("/categories", middleware.VerifyJWT, controllers.GetAllCategory)
	// Back Office
	router.POST("/bo/categories", middleware.VerifyJWT, middleware.BackOffice, controllers.CreateCategory)
	router.PUT("/bo/categories/:id", middleware.VerifyJWT, middleware.BackOffice, controllers.UpdateCategory)
	router.DELETE("/bo/categories/:id", middleware.VerifyJWT, middleware.BackOffice, controllers.DeleteCategory)

	// Event - CRUD + Get All Ticket by EventId
	// Customer Side
	router.GET("/events", middleware.VerifyJWT, controllers.GetAllEvent)
	router.GET("/events/:id/tickets", middleware.VerifyJWT, controllers.GetAllTicketByEventId)
	// Back Office
	router.POST("/bo/events", middleware.VerifyJWT, middleware.BackOffice, controllers.CreateEvent)
	router.PUT("/bo/events/:id", middleware.VerifyJWT, middleware.BackOffice, controllers.UpdateEvent)
	router.DELETE("/bo/events/:id", middleware.VerifyJWT, middleware.BackOffice, controllers.DeleteEvent)

	// Customer - RU (Customer side)
	router.GET("/customer/:id", middleware.VerifyJWT, controllers.GetCustomerById)
	router.PUT("/customer/:id", middleware.VerifyJWT, controllers.UpdateCustomer)
	router.GET("/customer/:id/transactions", middleware.VerifyJWT, controllers.GetTransactionByCustomerId)

	// Wallet (Customer side)
	router.GET("/wallet/:id", middleware.VerifyJWT, controllers.GetWalletInfo)
	router.PUT("/wallet/top_up", middleware.VerifyJWT, controllers.TopUpBalance)

	// Transaction - CR
	// Customer side
	router.GET("/transactions/:id", middleware.VerifyJWT, controllers.GetTransactionById)
	router.POST("/transactions", middleware.VerifyJWT, controllers.CreateTransaction)
	// Back Office
	router.GET("/bo/transactions", middleware.VerifyJWT, middleware.BackOffice, controllers.GetAllTransactions)

	router.Run(":" + os.Getenv("PORT"))

}
