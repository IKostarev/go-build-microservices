package main

import (
	"context"
	"go-build-microservices/working/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()

	sm.Handle("/hello", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalCh := make(chan os.Signal)

	sig := <-signalCh
	l.Println("Recieved terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
