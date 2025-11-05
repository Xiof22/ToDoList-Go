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

func TestDeleteTask(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	client := ts.Client()

	taskResp := createTask(t, client, ts.URL, sampleTaskMap)
	strTaskID := strconv.Itoa(taskResp.Task.ID)

	tests := []struct {
		name       string
		taskID     string
		wantStatus int
		wantError  *dto.ErrorsResponse
	}{
		{
			name:       "Task not found",
			taskID:     "999",
			wantStatus: http.StatusNotFound,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Task not found"},
			},
		},
		{
			name:       "Task ID less than 1",
			taskID:     "0",
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Field 'ID' doesn't match the rule 'gt'"},
			},
		},
		{
			name:       "Alphameric Task ID",
			taskID:     "abc",
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'id'"},
			},
		},
		{
			name:       "Success",
			taskID:     strTaskID,
			wantStatus: http.StatusNoContent,
			wantError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/tasks/%s", ts.URL, tt.taskID)

			req, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			if tt.wantError != nil {
				gotError := &dto.ErrorsResponse{}
				require.NoError(t, json.NewDecoder(resp.Body).Decode(gotError))

				assert.Equal(t, tt.wantError, gotError)
				return
			}

			resp2, err := client.Get(url)
			require.NoError(t, err)
			defer resp2.Body.Close()

			assert.Equal(t, http.StatusNotFound, resp2.StatusCode)
		})
	}
}
