package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"cmd/main/main.go/pkg/main/controllers"

	// "github.com/ivedi/simple-go-redis/pkg/main/controllers"
	"github.com/gorilla/mux"
)

var redisController = controllers.SimpleRedisController{}

func initialize() {
	redisController.Initialize()
}

func createServer() *http.Server {
	router := mux.NewRouter()
	router.Headers("Content-Type", "application/text")
	router.HandleFunc("/set/{key}", redisController.Set).Methods("POST")
	router.HandleFunc("/get/{key}", redisController.Get).Methods("GET")
	router.HandleFunc("/flush", redisController.Flush).Methods("DELETE")

	server := &http.Server{
		Handler: router,
		Addr:    ":10000",
	}

	return server
}

func onShutdown(server *http.Server, doBeforeShutdown func()) {
	// Creating a waiting group that waits until the graceful shutdown procedure is done
	var wg sync.WaitGroup
	wg.Add(1)

	// This goroutine is running in parallels to the main one
	go func() {
		// creating a channel to listen for signals, like SIGINT
		stop := make(chan os.Signal, 1)
		// subcribing to interruption signals
		signal.Notify(stop, os.Interrupt)
		// this blocks until the signal is received
		<-stop
		// initiating the shutdown
		err := server.Shutdown(context.Background())
		// can't do much here except for logging any errors
		if err != nil {
			log.Printf("error during shutdown: %v\n", err)
		}
		// notifying the main goroutine that we are done
		wg.Done()
	}()

	log.Println("listening on port: 10000")
	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		log.Println("Commencing server shutdown...")
		doBeforeShutdown()
		wg.Wait()
		log.Println("server was gracefully shutdown.")
	} else {
		log.Printf("server error: %v\n", err)
	}
}

func main() {
	initialize()
	server := createServer()
	onShutdown(server, redisController.Close)
}
