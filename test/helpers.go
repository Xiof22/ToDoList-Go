package test

import (
	"bytes"
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository/memory"
	"github.com/Xiof22/ToDoList/internal/router"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestServer() *httptest.Server {
	m := memory.New()
	svc := service.New(m)
	h := handlers.New(svc)
	r := router.New(h)

	return httptest.NewServer(r)
}

func createTask(t *testing.T, client *http.Client, baseURL string, taskMap map[string]any) dto.TaskResponse {
	t.Helper()

	body, err := json.Marshal(taskMap)
	require.NoError(t, err)

	resp, err := client.Post(baseURL+"/tasks", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusCreated, resp.StatusCode)

	var taskResp dto.TaskResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&taskResp))

	return taskResp
}
