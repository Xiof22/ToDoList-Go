package test

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository/memory"
	"github.com/Xiof22/ToDoList/internal/router"
	"github.com/Xiof22/ToDoList/internal/service"
	"net/http/httptest"
)

func newTestServer() *httptest.Server {
	m := memory.New()
	svc := service.New(m)
	h := handlers.New(svc)
	r := router.New(h)

	return httptest.NewServer(r)
}
