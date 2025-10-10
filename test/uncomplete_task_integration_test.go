package test

import (
	"encoding/json"
	"fmt"
	"github.com/Xiof22/ToDoList/internal/dto"
	_ "github.com/Xiof22/ToDoList/internal/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"strconv"
	"testing"
)

func TestUncompleteTask(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	client := ts.Client()

	taskResp := createTask(t, client, ts.URL, sampleTaskMap)
	taskID := taskResp.Task.ID
	strTaskID := strconv.Itoa(taskID)

	failTests := []struct {
		name       string
		taskID     string
		wantStatus int
		wantError  dto.ErrorsResponse
	}{
		{
			name:       "Task not found",
			taskID:     "999",
			wantStatus: http.StatusNotFound,
			wantError: dto.ErrorsResponse{
				Errors: []string{"Task not found"},
			},
		},
		{
			name:       "Task ID less than 1",
			taskID:     "0",
			wantStatus: http.StatusBadRequest,
			wantError: dto.ErrorsResponse{
				Errors: []string{"Field 'ID' doesn't match the rule 'gt'"},
			},
		},
		{
			name:       "Alphameric Task ID",
			taskID:     "abc",
			wantStatus: http.StatusBadRequest,
			wantError: dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'id'"},
			},
		},
		{
			name:       "Task is already uncompleted",
			taskID:     strTaskID,
			wantStatus: http.StatusBadRequest,
			wantError: dto.ErrorsResponse{
				Errors: []string{"Task is already uncompleted"},
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
		strTaskID := strconv.Itoa(taskID)

		taskURL := fmt.Sprintf("%s/tasks/%s", ts.URL, strTaskID)

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
