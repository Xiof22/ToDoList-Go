package test

import (
	"encoding/json"
	"fmt"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	_ "github.com/Xiof22/ToDoList/internal/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func TestUncompleteTask(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	jar, _ := cookiejar.New(nil)
	client := ts.Client()
	client.Jar = jar

	createUser(t, client, ts.URL, newUserMap("UncompleteTask@gmail.com", "0000"))

	listResp := createList(t, client, ts.URL, sampleListMap)
	listID := listResp.List.ID

	taskResp := createTask(t, client, ts.URL, listID, sampleTaskMap)
	taskID := taskResp.Task.ID

	failTests := []struct {
		name       string
		listID     string
		taskID     string
		wantStatus int
		wantError  dto.ErrorsResponse
	}{
		{
			name:       "List not found",
			listID:     nilID,
			taskID:     taskID,
			wantStatus: http.StatusNotFound,
			wantError: dto.ErrorsResponse{
				Errors: []string{errorsx.ErrListNotFound.Error()},
			},
		},

		{
			name:       "Invalid list ID",
			listID:     invalidID,
			taskID:     taskID,
			wantStatus: http.StatusBadRequest,
			wantError: dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidListID.Error()},
			},
		},
		{
			name:       "Task not found",
			listID:     listID,
			taskID:     nilID,
			wantStatus: http.StatusNotFound,
			wantError: dto.ErrorsResponse{
				Errors: []string{errorsx.ErrTaskNotFound.Error()},
			},
		},
		{
			name:       "Invalid task ID",
			listID:     listID,
			taskID:     invalidID,
			wantStatus: http.StatusBadRequest,
			wantError: dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidTaskID.Error()},
			},
		},
		{
			name:       "Task is already uncompleted",
			listID:     listID,
			taskID:     taskID,
			wantStatus: http.StatusBadRequest,
			wantError: dto.ErrorsResponse{
				Errors: []string{errorsx.ErrAlreadyUncompleted.Error()},
			},
		},
	}

	for _, tt := range failTests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/lists/%s/tasks/%s/uncomplete", ts.URL, tt.listID, tt.taskID)

			req, err := http.NewRequest(http.MethodPatch, url, nil)
			require.NoError(t, err)

			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			var gotError dto.ErrorsResponse
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&gotError))

			assert.Equal(t, tt.wantError, gotError)
		})
	}

	t.Run("Success", func(t *testing.T) {
		taskResp := createTask(t, client, ts.URL, listID, sampleTaskMap)
		taskID := taskResp.Task.ID

		taskURL := fmt.Sprintf("%s/lists/%s/tasks/%s", ts.URL, listID, taskID)

		req, err := http.NewRequest(http.MethodPatch, taskURL+"/complete", nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusNoContent, resp.StatusCode)

		req, err = http.NewRequest(http.MethodPatch, taskURL+"/uncomplete", nil)
		require.NoError(t, err)

		resp2, err := client.Do(req)
		require.NoError(t, err)
		defer resp2.Body.Close()

		assert.Equal(t, http.StatusNoContent, resp2.StatusCode)

		resp3, err := client.Get(taskURL)
		require.NoError(t, err)
		defer resp3.Body.Close()

		var gotResponse dto.TaskResponse
		require.NoError(t, json.NewDecoder(resp3.Body).Decode(&gotResponse))

		assert.Equal(t, false, gotResponse.Task.IsCompleted)
	})
}
