package test

import (
	"bytes"
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

func TestEditTask(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	client := ts.Client()

	taskResp := createTask(t, client, ts.URL, sampleTaskMap)
	taskID := taskResp.Task.ID

	editedTaskMap := map[string]any{
		"title":       "Edited title",
		"description": "Edited description",
	}

	tests := []struct {
		name       string
		taskID     string
		payload    map[string]any
		wantStatus int
		wantError  *dto.ErrorsResponse
	}{
		{
			name:       "Task not found",
			taskID:     nilID,
			payload:    editedTaskMap,
			wantStatus: http.StatusNotFound,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrTaskNotFound.Error()},
			},
		},
		{
			name:       "Invalid task ID",
			taskID:     invalidID,
			payload:    editedTaskMap,
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidTaskID.Error()},
			},
		},
		{
			name:   "Missing title",
			taskID: taskID,
			payload: map[string]any{
				"title": "     ",
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrValidation("Title", "required").Error()},
			},
		},
		{
			name:   "Deadline before creation",
			taskID: taskID,
			payload: map[string]any{
				"title":       sampleTaskMap["title"],
				"description": sampleTaskMap["description"],
				"deadline":    pastDeadline,
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrDeadlineBeforeCreation.Error()},
			},
		},
		{
			name:   "Unexpected deadline format",
			taskID: taskID,
			payload: map[string]any{
				"title":       sampleTaskMap["title"],
				"description": sampleTaskMap["description"],
				"deadline":    invalidFormatDeadline,
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidDeadlineFormat.Error()},
			},
		},
		{
			name:       "Success",
			taskID:     taskID,
			payload:    editedTaskMap,
			wantStatus: http.StatusOK,
			wantError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.payload)
			require.NoError(t, err)

			url := fmt.Sprintf("%s/tasks/%s", ts.URL, tt.taskID)

			req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(body))
			require.NoError(t, err)
			req.Header.Set("content-type", "application/json")

			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			if tt.wantError != nil {
				gotError := &dto.ErrorsResponse{}
				require.NoError(t, json.NewDecoder(resp.Body).Decode(gotError))
				assert.Equal(t, tt.wantError, gotError)
			}
		})
	}

	t.Run("Missing body", func(t *testing.T) {
		url := fmt.Sprintf("%s/tasks/%s", ts.URL, taskID)

		req, err := http.NewRequest(http.MethodPatch, url, nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		wantError := dto.ErrorsResponse{
			Errors: []string{errorsx.ErrMissingJSON.Error()},
		}

		var gotError dto.ErrorsResponse
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&gotError))

		assert.Equal(t, wantError, gotError)
	})
}
