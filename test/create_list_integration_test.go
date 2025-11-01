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
	"testing"
)

func TestCreateList(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	client := ts.Client()

	url := fmt.Sprintf("%s/lists", ts.URL)

	tests := []struct {
		name       string
		payload    map[string]any
		wantStatus int
		wantError  *dto.ErrorsResponse
	}{
		{
			name: "Missing title",
			payload: map[string]any{
				"title": "       ",
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{"Field 'Title' doesn't match the rule 'required'"},
			},
		},
		{
			name:       "Success",
			payload:    sampleListMap,
			wantStatus: http.StatusCreated,
			wantError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.payload)
			require.NoError(t, err)

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
