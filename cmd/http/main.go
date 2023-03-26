package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snehankitrd/pet-pro/internal/pet-pro/delivery/menu"
	"github.com/snehankitrd/pet-pro/internal/pet-pro/infra/dataprovider"
	"github.com/snehankitrd/pet-pro/internal/pet-pro/usecase"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	dataProvider := dataprovider.NewService(db)
	menuService := menu.NewService(dataProvider)

	router := gin.Default()

	router.GET("/api/v1/menu", menuService.GetMenu)
	router.GET("/api/v1/menu/item/:id", menuService.GetMenuItem)
	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

var db map[string]usecase.Item = map[string]usecase.Item{
	"123": {Id: "123", Name: "pasta", Price: 120, Note: "mild", Type: "italian"},
	"345": {Id: "345", Name: "noodles", Price: 100, Note: "spicy", Type: "asian"},
	"876": {Id: "876", Name: "pavbhaji", Price: 80, Note: "hot and spicy", Type: "indian"},
	"455": {Id: "455", Name: "currywurst", Price: 30, Note: "medium spicy", Type: "german"},
	"785": {Id: "785", Name: "doner", Price: 50, Note: "spicy", Type: "turkish"},
}
