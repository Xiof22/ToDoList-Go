package test

import (
	"encoding/json"
	"fmt"
	"github.com/Xiof22/ToDoList/internal/dto"
	_ "github.com/Xiof22/ToDoList/internal/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"testing"
)

func TestGetTask(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	jar, _ := cookiejar.New(nil)
	client := ts.Client()
	client.Jar = jar

	createUser(t, client, ts.URL, newUserMap("GetTask@gmail.com", "0000"))

	listResp := createList(t, client, ts.URL, sampleListMap)
	listID := listResp.List.ID
	strListID := strconv.Itoa(listID)

	taskResp := createTask(t, client, ts.URL, listID, sampleTaskMap)
	taskID := taskResp.Task.ID
	strTaskID := strconv.Itoa(taskID)

	tests := []struct {
		name         string
		listID       string
		taskID       string
		wantStatus   int
		wantResponse *dto.TaskResponse
		wantError    *dto.ErrorsResponse
	}{
		{
			name:         "List not found",
			listID:       "999",
			taskID:       strTaskID,
			wantStatus:   http.StatusNotFound,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"List not found"},
			},
		},

		{
			name:         "List ID less than 1",
			listID:       "0",
			taskID:       strTaskID,
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Invalid list ID"},
			},
		},
		{
			name:         "Alphameric List ID",
			listID:       "abc",
			taskID:       strTaskID,
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'list_id' from URL"},
			},
		},
		{
			name:         "Task not found",
			listID:       strListID,
			taskID:       "999",
			wantStatus:   http.StatusNotFound,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Task not found"},
			},
		},
		{
			name:         "Task ID less than 1",
			listID:       strListID,
			taskID:       "0",
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Invalid task ID"},
			},
		},
		{
			name:         "Alphameric Task ID",
			listID:       strListID,
			taskID:       "abc",
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'task_id' from URL"},
			},
		},
		{
			name:       "Success",
			listID:     strListID,
			taskID:     strTaskID,
			wantStatus: http.StatusOK,
			wantResponse: &dto.TaskResponse{
				Task: sampleTask,
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/lists/%s/tasks/%s", ts.URL, tt.listID, tt.taskID)

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
			assert.Greater(t, gotResponse.Task.ID, 0)
		})
	}
}
