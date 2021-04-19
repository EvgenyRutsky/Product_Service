package main

import (
	"context"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go_microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "hello-api", log.LstdFlags)

	ph := handlers.NewProduct(l)

	sm := mux.NewRouter()

	cors := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)

	server := &http.Server{
		Addr : ":9090",
		Handler: cors(sm),
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func(){
		l.Println("Server is starting")
		log.Fatal(server.ListenAndServe())
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig :=  <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(tc)
}
