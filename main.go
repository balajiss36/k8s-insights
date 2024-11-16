package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/balajiss36/k8s-insights/db"
	"github.com/balajiss36/k8s-insights/misc"
	"github.com/balajiss36/k8s-insights/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := misc.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	router := gin.Default()

	client, err := db.SetupMongoDB(ctx, config)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v\n", err)
	}

	defer db.CloseConnection(ctx, client)

	handlers := routes.NewHandler(ctx, client)

	handlers.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    config.HTTPAddress,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
