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
	"testing"
)

func TestUncompleteTask(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	client := ts.Client()

	taskResp := createTask(t, client, ts.URL, sampleTaskMap)
	taskID := taskResp.Task.ID

	failTests := []struct {
		name       string
		taskID     string
		wantStatus int
		wantError  dto.ErrorsResponse
	}{
		{
			name:       "Task not found",
			taskID:     nilID,
			wantStatus: http.StatusNotFound,
			wantError: dto.ErrorsResponse{
				Errors: []string{errorsx.ErrTaskNotFound.Error()},
			},
		},
		{
			name:       "Invalid task ID",
			taskID:     invalidID,
			wantStatus: http.StatusBadRequest,
			wantError: dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidTaskID.Error()},
			},
		},
		{
			name:       "Task is already uncompleted",
			taskID:     taskID,
			wantStatus: http.StatusBadRequest,
			wantError: dto.ErrorsResponse{
				Errors: []string{errorsx.ErrAlreadyUncompleted.Error()},
			},
		},
	}

	for _, tt := range failTests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/tasks/%s/uncomplete", ts.URL, tt.taskID)

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
		taskResp := createTask(t, client, ts.URL, sampleTaskMap)
		taskID := taskResp.Task.ID

		taskURL := fmt.Sprintf("%s/tasks/%s", ts.URL, taskID)

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
