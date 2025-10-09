package test

import (
	"bytes"
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

func TestEditTask(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	client := ts.Client()

	taskResp := createTask(t, client, ts.URL, sampleTaskMap)
	strTaskID := strconv.Itoa(taskResp.Task.ID)

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
			taskID:     "999",
			payload:    editedTaskMap,
			wantStatus: http.StatusNotFound,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Task not found"},
			},
		},
		{
			name:       "Task ID less than 1",
			taskID:     "0",
			payload:    editedTaskMap,
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Field 'ID' doesn't match the rule 'gt'"},
			},
		},
		{
			name:       "Alphameric Task ID",
			taskID:     "abc",
			payload:    editedTaskMap,
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'id'"},
			},
		},
		{
			name:   "Missing title",
			taskID: strTaskID,
			payload: map[string]any{
				"title": "     ",
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Field 'Title' doesn't match the rule 'required'"},
			},
		},
		{
			name:       "Success",
			taskID:     strTaskID,
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
		url := fmt.Sprintf("%s/tasks/%s", ts.URL, strTaskID)

		req, err := http.NewRequest(http.MethodPatch, url, nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		wantError := dto.ErrorsResponse{
			Errors: []string{"Empty JSON"},
		}

		var gotError dto.ErrorsResponse
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&gotError))

		assert.Equal(t, wantError, gotError)
	})
}
