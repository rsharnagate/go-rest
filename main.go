package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rsharnagate/go-rest/routes"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9000"
	}

	r := chi.NewRouter()

	r.Mount("/sample", routes.SampleRoutes())

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			fmt.Println("shutting down server")

			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			srv.Shutdown(ctx)

			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():
			}
		}
	}()

	fmt.Println("starting server on port: " + port)
	srv.ListenAndServe()
}
