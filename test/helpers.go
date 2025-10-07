package test

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository/memory"
	"github.com/Xiof22/ToDoList/internal/router"
	"github.com/Xiof22/ToDoList/internal/service"
	"net/http/httptest"
	"testing"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	m := memory.New()
	svc := service.New(m)
	h := handlers.New(svc)
	r := router.New(h)

	return httptest.NewServer(r)
}
