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

func TestLogout(t *testing.T) {
	ts := newTestServer(t)

	userData := map[string]any{
		"email":    "Logout@gmail.com",
		"password": "0000",
	}

	jar, _ := cookiejar.New(nil)
	client := ts.Client()
	client.Jar = jar

	createUser(t, client, ts.URL, userData)

	url := fmt.Sprintf("%s/auth/logout", ts.URL)

	tests := []struct {
		name       string
		wantStatus int
		wantError  *dto.ErrorsResponse
	}{
		{
			name:       "Success",
			wantStatus: http.StatusNoContent,
			wantError:  nil,
		},
		{
			name:       "Unauthorized",
			wantStatus: http.StatusUnauthorized,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrUnauthorized.Error()},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.Post(url, "application/json", nil)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			if tt.wantError == nil {
				return
			}

			gotError := &dto.ErrorsResponse{}
			require.NoError(t, json.NewDecoder(resp.Body).Decode(gotError))

			assert.Equal(t, tt.wantError, gotError)
		})
	}
}
