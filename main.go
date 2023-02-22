package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	controllers2 "ticketing/controllers"
	database2 "ticketing/database"
	middleware2 "ticketing/middleware"
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

	database2.DbMigrate(DB)

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(DB)

	// Router GIN
	router := gin.Default()

	// Registration
	router.POST("/registration", controllers2.CreateCustomer)

	// Login
	router.POST("/login", controllers2.Login)

	// Ticket - CRU + Get Ticket By ID
	// Customer Side
	router.GET("/tickets/:id", middleware2.VerifyJWT, controllers2.GetTicketById)
	// Back Office
	router.POST("/bo/tickets", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.CreateTicket)
	router.PUT("/bo/tickets/:id", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.UpdateTicket)
	router.DELETE("/bo/tickets/:id", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.DeleteTicket)

	// Category - CRUD + Get All Event by CategoryId
	// Customer Side
	router.GET("/categories/:id/events", middleware2.VerifyJWT, controllers2.GetAllEventByCategory)
	router.GET("/categories", middleware2.VerifyJWT, controllers2.GetAllCategory)
	// Back Office
	router.POST("/bo/categories", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.CreateCategory)
	router.PUT("/bo/categories/:id", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.UpdateCategory)
	router.DELETE("/bo/categories/:id", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.DeleteCategory)

	// Event - CRUD + Get All Ticket by EventId
	// Customer Side
	router.GET("/events", middleware2.VerifyJWT, controllers2.GetAllEvent)
	router.GET("/events/:id/tickets", middleware2.VerifyJWT, controllers2.GetAllTicketByEventId)
	// Back Office
	router.POST("/bo/events", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.CreateEvent)
	router.PUT("/bo/events/:id", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.UpdateEvent)
	router.DELETE("/bo/events/:id", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.DeleteEvent)

	// Customer - RU (Customer side)
	router.GET("/customer/:id", middleware2.VerifyJWT, controllers2.GetCustomerById)
	router.PUT("/customer/:id", middleware2.VerifyJWT, controllers2.UpdateCustomer)
	router.GET("/customer/:id/transactions", middleware2.VerifyJWT, controllers2.GetTransactionByCustomerId)

	// Wallet (Customer side)
	router.GET("/wallet/:id", middleware2.VerifyJWT, controllers2.GetWalletInfo)
	router.PUT("/wallet/top_up", middleware2.VerifyJWT, controllers2.TopUpBalance)

	// Transaction - CR
	// Customer side
	router.GET("/transactions/:id", middleware2.VerifyJWT, controllers2.GetTransactionById)
	router.POST("/transactions", middleware2.VerifyJWT, controllers2.CreateTransaction)
	// Back Office
	router.GET("/bo/transactions", middleware2.VerifyJWT, middleware2.BackOffice, controllers2.GetAllTransactions)

	router.Run(":" + os.Getenv("PORT"))

}
