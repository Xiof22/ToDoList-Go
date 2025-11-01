package test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func createList(t *testing.T, client *http.Client, baseURL string, listMap map[string]any) dto.ListResponse {
	t.Helper()

	body, err := json.Marshal(listMap)
	require.NoError(t, err)

	resp, err := client.Post(baseURL+"/lists", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusCreated, resp.StatusCode)

	var listResp dto.ListResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&listResp))

	return listResp
}

func createTask(t *testing.T, client *http.Client, baseURL string, strListID int, taskMap map[string]any) dto.TaskResponse {
	t.Helper()

	body, err := json.Marshal(taskMap)
	require.NoError(t, err)

	url := fmt.Sprintf("%s/lists/%d/tasks", baseURL, strListID)

	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusCreated, resp.StatusCode)

	var taskResp dto.TaskResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&taskResp))

	return taskResp
}
