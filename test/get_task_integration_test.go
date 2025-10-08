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

func TestGetTask(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	client := ts.Client()

	taskResp := createTask(t, client, ts.URL, sampleTaskMap)
	taskID := taskResp.Task.ID

	tests := []struct {
		name         string
		taskID       string
		wantStatus   int
		wantResponse *dto.TaskResponse
		wantError    *dto.ErrorsResponse
	}{
		{
			name:         "Task not found",
			taskID:       nilID,
			wantStatus:   http.StatusNotFound,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrTaskNotFound.Error()},
			},
		},
		{
			name:         "Invalid task ID",
			taskID:       invalidID,
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidTaskID.Error()},
			},
		},
		{
			name:       "Success",
			taskID:     taskID,
			wantStatus: http.StatusOK,
			wantResponse: &dto.TaskResponse{
				Task: sampleTask,
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/tasks/%s", ts.URL, tt.taskID)

			resp, err := client.Get(url)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			if tt.wantError != nil {
				gotError := &dto.ErrorsResponse{}
				require.NoError(t, json.NewDecoder(resp.Body).Decode(gotError))

				assert.Equal(t, tt.wantError, gotError)
				return
			}

			gotResponse := &dto.TaskResponse{}
			require.NoError(t, json.NewDecoder(resp.Body).Decode(gotResponse))

			assert.Equal(t, tt.wantResponse.Task.Title, gotResponse.Task.Title)
			assert.Equal(t, tt.wantResponse.Task.Description, gotResponse.Task.Description)
		})
	}
}
