package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/NeelavaChatterjee/git-sync/database"
	"github.com/NeelavaChatterjee/git-sync/routes"
	"github.com/NeelavaChatterjee/git-sync/utilities"
	"github.com/joho/godotenv"
)

func init() {
	// Loading environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Writing logs to a file
	f, err := os.OpenFile(os.Getenv("LOG_FILENAME"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	fmt.Println("Database is being initialized")

	utilities.Initialize()
	utilities.Cron.Start()
	defer utilities.Cron.Stop()

	// Connect to database
	database.Connect()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for the existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := routes.Router()

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Run our server in a goroutine so that it doesn't block
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We will accept graceful shutdowns when quit via SIGINT (Ctrl + C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl + /) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
