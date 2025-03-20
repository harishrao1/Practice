package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"userapi/config"
	"userapi/db"
	"userapi/handlers"
	"userapi/routes"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, Go API"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `pong`)
}
func main() {

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ping", pingHandler)

	cfg := config.LoadConfig()
	db.InitMySQL(cfg)

	userHandler := handlers.NewUserHandler(db.DB)
	r := routes.NewRouter(userHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// shutdown
	go func() {
		log.Println("Server started at :8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Server closed: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down gracefully...")

	db.DB.Close()
	os.Exit(0)
}
