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

func TestGetTask(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	jar, _ := cookiejar.New(nil)
	client := ts.Client()
	client.Jar = jar

	createUser(t, client, ts.URL, newUserMap("GetTask@gmail.com", "0000"))

	listResp := createList(t, client, ts.URL, sampleListMap)
	listID := listResp.List.ID

	taskResp := createTask(t, client, ts.URL, listID, sampleTaskMap)
	taskID := taskResp.Task.ID

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
			listID:       nilID,
			taskID:       taskID,
			wantStatus:   http.StatusNotFound,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrListNotFound.Error()},
			},
		},

		{
			name:         "Invalid list ID",
			listID:       invalidID,
			taskID:       taskID,
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidListID.Error()},
			},
		},
		{
			name:         "Task not found",
			listID:       listID,
			taskID:       nilID,
			wantStatus:   http.StatusNotFound,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrTaskNotFound.Error()},
			},
		},
		{
			name:         "Invalid task ID",
			listID:       listID,
			taskID:       invalidID,
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidTaskID.Error()},
			},
		},
		{
			name:       "Success",
			listID:     listID,
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
		})
	}
}
