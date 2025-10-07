package main

import (
	"fmt"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository/memory"
	"github.com/Xiof22/ToDoList/internal/router"
	"github.com/Xiof22/ToDoList/internal/service"
	_ "github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
)

func main() {
	m := memory.New()
	svc := service.New(m)
	h := handlers.New(svc)
	r := router.New(h)

	fmt.Println("Serving start...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
