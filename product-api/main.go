package main

import (
	"context"
	"go-build-microservices-product-api/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	productHandler := handlers.NewProducts(logger)

	serveMux := http.NewServeMux()

	serveMux.Handle("/", productHandler)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalCh := make(chan os.Signal)

	sig := <-signalCh
	logger.Println("Recieved terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
