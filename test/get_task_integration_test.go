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

func TestGetTask(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	client := ts.Client()

	taskResp := createTask(t, client, ts.URL, sampleTaskMap)
	strTaskID := strconv.Itoa(taskResp.Task.ID)

	tests := []struct {
		name         string
		taskID       string
		wantStatus   int
		wantResponse *dto.TaskResponse
		wantError    *dto.ErrorsResponse
	}{
		{
			name:       "Task not found",
			taskID:     "999",
			wantStatus: http.StatusNotFound,
			wantResponse: &dto.TaskResponse{
				Task: nil,
			},
			wantError: nil,
		},
		{
			name:         "Task ID less than 1",
			taskID:       "0",
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Field 'ID' doesn't match the rule 'gt'"},
			},
		},
		{
			name:         "Alphameric Task ID",
			taskID:       "abc",
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'id'"},
			},
		},
		{
			name:       "Success",
			taskID:     strTaskID,
			wantStatus: http.StatusOK,
			wantResponse: &dto.TaskResponse{
				Task: &sampleTask,
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

			if tt.wantResponse.Task == nil {
				assert.Equal(t, tt.wantResponse, gotResponse)
				return
			}

			assert.Equal(t, tt.wantResponse.Task.Title, gotResponse.Task.Title)
			assert.Equal(t, tt.wantResponse.Task.Description, gotResponse.Task.Description)
			assert.Greater(t, gotResponse.Task.ID, 0)
		})
	}
}
