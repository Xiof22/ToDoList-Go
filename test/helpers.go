package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Xiof22/ToDoList/config"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/middleware"
	"github.com/Xiof22/ToDoList/internal/repository/mysql"
	"github.com/Xiof22/ToDoList/internal/router"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	loc, err := time.LoadLocation(cfg.TimezoneLocation)
	if err != nil {
		fmt.Printf("Failed to load location %s: %v\n", cfg.TimezoneLocation, err)
		time.Local = time.UTC
	} else {
		time.Local = loc
	}

	cs := sessions.NewCookieStore([]byte(cfg.CookieStoreKey))
	cs.Options.Secure = false
	repo, err := mysql.New(cfg.DBDSN)
	require.NoError(t, err)
	svc := service.New(repo)
	h := handlers.New(svc, cs, cfg)
	mw := middleware.New(cs, cfg)
	r := router.New(h, mw)

	return httptest.NewServer(r)
}

func createUser(t *testing.T, client *http.Client, baseURL string, userMap map[string]any) dto.UserResponse {
	t.Helper()

	body, err := json.Marshal(userMap)
	require.NoError(t, err)

	resp, err := client.Post(baseURL+"/auth/register", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusCreated, resp.StatusCode)

	var userResp dto.UserResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&userResp))

	return userResp
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

func createTask(t *testing.T, client *http.Client, baseURL string, listID string, taskMap map[string]any) dto.TaskResponse {
	t.Helper()

	body, err := json.Marshal(taskMap)
	require.NoError(t, err)

	url := fmt.Sprintf("%s/lists/%s/tasks", baseURL, listID)

	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusCreated, resp.StatusCode)

	var taskResp dto.TaskResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&taskResp))

	return taskResp
}

func errStrings(errs ...error) []string {
	s := make([]string, len(errs))
	for i, err := range errs {
		s[i] = err.Error()
	}

	return s
}
