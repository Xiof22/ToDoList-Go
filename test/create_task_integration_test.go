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
	"net/http/cookiejar"
	"strconv"
	"testing"
)

func TestCreateTask(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	jar, _ := cookiejar.New(nil)
	client := ts.Client()
	client.Jar = jar

	createUser(t, client, ts.URL, newUserMap("CreateTask@gmail.com", "0000"))

	listResp := createList(t, client, ts.URL, sampleListMap)
	strListID := strconv.Itoa(listResp.List.ID)

	tests := []struct {
		name       string
		listID     string
		payload    map[string]any
		wantStatus int
		wantError  *dto.ErrorsResponse
	}{
		{
			name:       "List not found",
			listID:     "999",
			payload:    sampleTaskMap,
			wantStatus: http.StatusNotFound,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"List not found"},
			},
		},
		{
			name:       "List ID less than 1",
			listID:     "0",
			payload:    sampleTaskMap,
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Invalid list ID"},
			},
		},
		{
			name:       "Alphameric List ID",
			listID:     "abc",
			payload:    sampleTaskMap,
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'list_id' from URL"},
			},
		},
		{
			name:   "Missing title",
			listID: strListID,
			payload: map[string]any{
				"title": "       ",
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Field 'Title' doesn't match the rule 'required'"},
			},
		},
		{
			name:   "Deadline before creation",
			listID: strListID,
			payload: map[string]any{
				"title":       sampleTaskMap["title"],
				"description": sampleTaskMap["description"],
				"deadline":    "2004-07-12 16:59:21",
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Field 'Deadline' doesn't match the rule 'future_or_empty'"},
			},
		},
		{
			name:   "Unexpected deadline format",
			listID: strListID,
			payload: map[string]any{
				"title":       sampleTaskMap["title"],
				"description": sampleTaskMap["description"],
				"deadline":    "The 7-th of December 2030 year",
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Unexpected deadline format"},
			},
		},
		{
			name:       "Success",
			listID:     strListID,
			payload:    sampleTaskMap,
			wantStatus: http.StatusCreated,
			wantError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.payload)
			require.NoError(t, err)

			url := fmt.Sprintf("%s/lists/%s/tasks", ts.URL, tt.listID)

			resp, err := client.Post(url, "application/json", bytes.NewReader(body))
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
		url := fmt.Sprintf("%s/lists/%s/tasks", ts.URL, strListID)

		req, err := http.NewRequest(http.MethodPost, url, nil)
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
