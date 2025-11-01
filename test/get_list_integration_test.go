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

func TestGetList(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	client := ts.Client()

	listResp := createList(t, client, ts.URL, sampleListMap)
	strListID := strconv.Itoa(listResp.List.ID)

	tests := []struct {
		name         string
		listID       string
		wantStatus   int
		wantResponse *dto.ListResponse
		wantError    *dto.ErrorsResponse
	}{
		{
			name:       "List not found",
			listID:     "999",
			wantStatus: http.StatusNotFound,
			wantResponse: &dto.ListResponse{
				List: nil,
			},
			wantError: nil,
		},
		{
			name:         "List ID less than 1",
			listID:       "0",
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Field 'ID' doesn't match the rule 'gt'"},
			},
		},
		{
			name:         "Alphameric List ID",
			listID:       "abc",
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Failed to parse 'list_id'"},
			},
		},
		{
			name:       "Success",
			listID:     strListID,
			wantStatus: http.StatusOK,
			wantResponse: &dto.ListResponse{
				List: &sampleList,
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/lists/%s", ts.URL, tt.listID)

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

			gotResponse := &dto.ListResponse{}
			require.NoError(t, json.NewDecoder(resp.Body).Decode(gotResponse))

			if tt.wantResponse.List == nil {
				assert.Equal(t, tt.wantResponse, gotResponse)
				return
			}

			assert.Equal(t, tt.wantResponse.List.Title, gotResponse.List.Title)
			assert.Equal(t, tt.wantResponse.List.Description, gotResponse.List.Description)
			assert.Greater(t, gotResponse.List.ID, 0)
		})
	}
}
