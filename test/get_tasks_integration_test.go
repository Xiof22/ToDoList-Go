package test

import (
	"encoding/json"
	"fmt"
	"github.com/Xiof22/ToDoList/internal/dto"
	_ "github.com/Xiof22/ToDoList/internal/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTasks(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	client := ts.Client()

	url := fmt.Sprintf("%s/tasks", ts.URL)

	t.Run("No tasks", func(t *testing.T) {

		resp, err := client.Get(url)
		require.NoError(t, err)
		defer resp.Body.Close()

		wantResponse := dto.TasksResponse{
			Count: 0,
			Tasks: []dto.Task{},
		}

		var gotResponse dto.TasksResponse
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&gotResponse))

		assert.Equal(t, gotResponse, wantResponse)
	})

	t.Run("Have task", func(t *testing.T) {
		createTask(t, client, ts.URL, sampleTaskMap)

		resp, err := client.Get(url)
		require.NoError(t, err)
		defer resp.Body.Close()

		wantResponse := dto.TasksResponse{
			Count: 1,
			Tasks: []dto.Task{sampleTask},
		}

		var gotResponse dto.TasksResponse
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&gotResponse))

		assert.Equal(t, gotResponse.Count, wantResponse.Count)
		assert.Equal(t, gotResponse.Tasks[0].Title, wantResponse.Tasks[0].Title)
		assert.Greater(t, gotResponse.Tasks[0].ID, 0)
	})
}
