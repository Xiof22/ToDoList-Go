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

func TestGetList(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	jar, _ := cookiejar.New(nil)
	client := ts.Client()
	client.Jar = jar

	createUser(t, client, ts.URL, newUserMap("GetList@gmail.com", "0000"))

	listResp := createList(t, client, ts.URL, sampleListMap)
	listID := listResp.List.ID

	tests := []struct {
		name         string
		listID       string
		wantStatus   int
		wantResponse *dto.ListResponse
		wantError    *dto.ErrorsResponse
	}{
		{
			name:         "List not found",
			listID:       nilID,
			wantStatus:   http.StatusNotFound,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrListNotFound.Error()},
			},
		},
		{
			name:         "Invalid list ID",
			listID:       invalidID,
			wantStatus:   http.StatusBadRequest,
			wantResponse: nil,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidListID.Error()},
			},
		},
		{
			name:       "Success",
			listID:     listID,
			wantStatus: http.StatusOK,
			wantResponse: &dto.ListResponse{
				List: sampleList,
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

			assert.Equal(t, tt.wantResponse.List.Title, gotResponse.List.Title)
			assert.Equal(t, tt.wantResponse.List.Description, gotResponse.List.Description)
		})
	}
}
