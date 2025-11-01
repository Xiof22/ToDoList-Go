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

func TestEditList(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	client := ts.Client()

	listResp := createList(t, client, ts.URL, sampleListMap)
	strListID := strconv.Itoa(listResp.List.ID)

	editedListMap := map[string]any{
		"title":       "Edited title",
		"description": "Edited description",
	}

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
			payload:    editedListMap,
			wantStatus: http.StatusNotFound,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"List not found"},
			},
		},
		{
			name:       "List ID less than 1",
			listID:     "0",
			payload:    editedListMap,
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Field 'ListID' doesn't match the rule 'gt'"},
			},
		},
		{
			name:       "Alphameric List ID",
			listID:     "abc",
			payload:    editedListMap,
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'list_id'"},
			},
		},
		{
			name:   "Missing title",
			listID: strListID,
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
			listID:     strListID,
			payload:    editedListMap,
			wantStatus: http.StatusOK,
			wantError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.payload)
			require.NoError(t, err)

			url := fmt.Sprintf("%s/lists/%s", ts.URL, tt.listID)

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
		url := fmt.Sprintf("%s/lists/%s", ts.URL, strListID)

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
