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

	"github.com/NeelavaChatterjee/git-sync/controllers"
	"github.com/NeelavaChatterjee/git-sync/database"
	"github.com/NeelavaChatterjee/git-sync/models"
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
	// TODO Remove comments to use
	database.Connect()
	// utilities.GetBranches("platform9", "pf9-kube")
	// utilities.GetCommits("platform9", "pf9-kube")
	// utilities.GetCommit("platform9", "pf9-kube", "897a05fabf282cdbdc5828b804cc216042172dd8")
	// utilities.GetFileContents("platform9", "pf9-kube", "atherton", "README.md")
	poll_interval, err := time.ParseDuration("1h")
	if err != nil {
		panic(err)
	}
	track_new := &models.Track{
		Owner:        "platform9",
		Repository:   "pf9-kube",
		Branch:       "atherton",
		PollInterval: poll_interval.String(),
		IsTracked:    true,
	}
	controllers.CreateTrackEntry(track_new)
	track, err := controllers.FindTrack("pf9-kube", "atherton")
	if err != nil {
		panic(err)
	}
	controllers.Poll(track)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for the existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// r := mux.NewRouter()
	router := routes.Router()

	srv := &http.Server{
		Addr: "127.0.0.1:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
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

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Does not block if no connections, but will otherwise wait until the timeout deadline
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
