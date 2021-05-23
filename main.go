package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shorturl/cache"
	"shorturl/cronjob"
	"shorturl/router"
	"time"
)

func init() {
	cronjob.Start()
}

func main() {
	defer cronjob.Stop()
	defer cache.DeleteAll()

	log.Println("Server start...")

	router := router.InitRouter()

	service := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := service.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	log.Println("Server listen...")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go func() {
		cancel()
	}()

	if err := service.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shutdown http service. Error: ", err.Error())
	}

	log.Println("Shutdown...")
}
