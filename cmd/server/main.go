package main

import (
	"fmt"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository/memory"
	"github.com/Xiof22/ToDoList/internal/router"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/Xiof22/ToDoList/config"
	_ "github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
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

	port := fmt.Sprintf(":%d", cfg.Port)
	fmt.Println("Serving start...")
	if err := http.ListenAndServe(port, r); err != nil {
		panic(err)
	}
}
