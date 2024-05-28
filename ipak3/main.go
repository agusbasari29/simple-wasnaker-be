package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simpel-wasnaker/ipak3/database"
	"simpel-wasnaker/ipak3/database/seeders"
	"simpel-wasnaker/ipak3/entity"
	"simpel-wasnaker/ipak3/router"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDBConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)
	db.AutoMigrate(&entity.Role{}, &entity.Contact{}, &entity.Company{}, &entity.User{})

	g := gin.Default()

	router.DefineAuthRoutes(g)
	router.DefineSecureRoutes(g)

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	g.GET("/insertfakedata", func(ctx *gin.Context) {
		seeders.Seeders()
		ctx.JSON(http.StatusOK, gin.H{"message": "Data inserted"})
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// g.Run(":" + port) => revision 1

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: g,
	}

	// Graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}

		log.Println("Server exiting")
	}()

	// Start the server
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
