package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	todo "github.com/CSSE6400/2026-p1-keys-i/internal/todo"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// get PORT number from env
	port := flag.String("p", "8080", "port to listen on")
	flag.Parse()

	router := todo.NewRouter()

	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: router,

		// timeout limits
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,

		// header abuse protection
		MaxHeaderBytes: 1 << 20, // 1 MiB
	}

	// starting a go routine to manage the server
	go func() {
		log.Printf("listening on 127.0.0.1:%s", *port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// shutting down the routine
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("shutting down...")
	_ = srv.Shutdown(ctx)
}
