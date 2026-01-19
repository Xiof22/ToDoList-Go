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

func TestGetTasks(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	jar, _ := cookiejar.New(nil)
	client := ts.Client()
	client.Jar = jar

	createUser(t, client, ts.URL, newUserMap("GetTasks@gmail.com", "0000"))

	listResp := createList(t, client, ts.URL, sampleListMap)
	listID := listResp.List.ID
	strListID := strconv.Itoa(listID)

	tests := []struct {
		name         string
		listID       string
		wantStatus   int
		wantResponse *dto.TasksResponse
		wantError    *dto.ErrorsResponse
	}{
		{
			name:         "List not found",
			listID:       "999",
			wantStatus:   http.StatusNotFound,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"List not found"},
			},
		},
		{
			name:         "List ID less than 1",
			listID:       "0",
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Invalid list ID"},
			},
		},
		{
			name:         "Alphameric list ID",
			listID:       "abc",
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'list_id' from URL"},
			},
		},
		{
			name:       "No tasks",
			listID:     strListID,
			wantStatus: http.StatusOK,
			wantResponse: &dto.TasksResponse{
				Count: 0,
				Tasks: []dto.Task{},
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/lists/%s/tasks", ts.URL, tt.listID)

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

			gotResponse := &dto.TasksResponse{}
			require.NoError(t, json.NewDecoder(resp.Body).Decode(gotResponse))

			assert.Equal(t, gotResponse, tt.wantResponse)
		})
	}

	t.Run("Have task", func(t *testing.T) {
		createTask(t, client, ts.URL, listID, sampleTaskMap)

		url := fmt.Sprintf("%s/lists/%s/tasks", ts.URL, strListID)

		resp, err := client.Get(url)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		wantResponse := dto.TasksResponse{
			Count: 1,
			Tasks: []dto.Task{sampleTask},
		}

		var gotResponse dto.TasksResponse
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&gotResponse))

		assert.Equal(t, gotResponse.Count, wantResponse.Count)
		assert.Equal(t, gotResponse.Tasks[0].Title, wantResponse.Tasks[0].Title)
		assert.Equal(t, gotResponse.Tasks[0].Deadline, wantResponse.Tasks[0].Deadline)
		assert.Greater(t, gotResponse.Tasks[0].ID, 0)
	})
}
