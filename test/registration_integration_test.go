package test

import (
	"bytes"
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

func TestRegistration(t *testing.T) {
	ts := newTestServer(t)

	jar, _ := cookiejar.New(nil)
	client := ts.Client()
	client.Jar = jar

	userData := map[string]any{
		"email":    "Register@gmail.com",
		"password": "0000",
	}

	url := fmt.Sprintf("%s/auth/register", ts.URL)

	tests := []struct {
		name       string
		payload    map[string]any
		wantStatus int
		wantError  *dto.ErrorsResponse
	}{
		{
			name: "Missing email",
			payload: map[string]any{
				"password": userData["password"],
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrValidation("Email", "required").Error()},
			},
		},
		{
			name: "Invalid email",
			payload: map[string]any{
				"email":    invalidEmail,
				"password": userData["password"],
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrValidation("Email", "email").Error()},
			},
		},
		{
			name: "Missing password",
			payload: map[string]any{
				"email": userData["email"],
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrValidation("Password", "required").Error()},
			},
		},
		{
			name: "Too short password",
			payload: map[string]any{
				"email":    userData["email"],
				"password": tooShortPassword,
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrValidation("Password", "min").Error()},
			},
		},
		{
			name: "Too long password",
			payload: map[string]any{
				"email":    userData["email"],
				"password": tooLongPassword,
			},
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrValidation("Password", "max").Error()},
			},
		},
		{
			name:       "Success",
			payload:    userData,
			wantStatus: http.StatusCreated,
			wantError:  nil,
		},
		{
			name:       "Already authorized",
			payload:    userData,
			wantStatus: http.StatusBadRequest,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrAlreadyAuthorized.Error()},
			},
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

			if tt.wantError == nil {
				return
			}

			gotError := &dto.ErrorsResponse{}
			require.NoError(t, json.NewDecoder(resp.Body).Decode(gotError))

			assert.Equal(t, tt.wantError, gotError)
		})
	}

	t.Run("Email already registered", func(t *testing.T) {
		body, err := json.Marshal(userData)
		require.NoError(t, err)

		resp, err := http.Post(url, "application/json", bytes.NewReader(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusConflict, resp.StatusCode)

		var gotError dto.ErrorsResponse
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&gotError))

		assert.Equal(t, dto.ErrorsResponse{
			Errors: []string{errorsx.ErrEmailRegistered.Error()},
		}, gotError)
	})
}
