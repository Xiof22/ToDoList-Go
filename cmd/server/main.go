package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Xiof22/ToDoList/config"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository/memory"
	"github.com/Xiof22/ToDoList/internal/router"
	"github.com/Xiof22/ToDoList/internal/server"
	"github.com/Xiof22/ToDoList/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	m := memory.New()
	svc := service.New(m)
	h := handlers.New(svc)
	r := router.New(h)
	srv := server.New(r, cfg)

	fmt.Printf("Serving start on %s...\n", cfg.Addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh

	fmt.Printf("Shutting down in %d seconds...\n", cfg.ShutdownTimeout/1000000000)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
