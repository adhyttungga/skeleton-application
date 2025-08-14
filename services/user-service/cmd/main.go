package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adhyttungga/skeleton-application/services/user-service/config"
	"github.com/adhyttungga/skeleton-application/services/user-service/routes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	log.Println("Server initialized...")
	log.Printf("MONGO_URI: %s", config.Config.Database.URI)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create database connection
	conn := config.Connect(ctx)
	db := conn.Database(config.Config.Database.DatabaseName)
	defer conn.Disconnect(ctx)

	// Create a unique index on the "email" field
	userColl := db.Collection("users")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}}, // 1 for ascending
		Options: options.Index().SetUnique(true),
	}

	_, err := userColl.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Error creating unique index: %v", err)
	}
	log.Println("Unique index created on 'email' field")

	router := routes.NewRouter(db)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Config.Port),
		Handler: router,
	}

	// Gracefully handle shutdown
	go func() {
		// Service connection
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen: %v\n", err)
		}
	}()

	log.Printf("Listening and serving HTTP on %s", server.Addr)

	// Wait for interrupt signal to gracefully shutdown the server
	// with timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT (Ctrl + C)
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it (no chance to recover, not recommended)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown: %s\n", err)
	}

	// Catching ctx.Done(). Timeout 5 seconds
	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 seconds reached, shutting down User Service.")
	default:
	}

	log.Println("User Service shutdown completed gracefully")
}
