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

func TestLogin(t *testing.T) {
	ts := newTestServer(t)

	userData := map[string]any{
		"email":    "Login@gmail.com",
		"password": "0000",
	}

	client := ts.Client()

	createUser(t, client, ts.URL, userData)

	jar, _ := cookiejar.New(nil)
	client.Jar = jar

	url := fmt.Sprintf("%s/auth/login", ts.URL)

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
			name: "Invalid credentials(email)",
			payload: map[string]any{
				"email":    "LoremIpsum@gmail.com",
				"password": userData["password"],
			},
			wantStatus: http.StatusUnauthorized,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidCredentials.Error()},
			},
		},
		{
			name: "Invalid credentials(password)",
			payload: map[string]any{
				"email":    userData["email"],
				"password": "lorem",
			},
			wantStatus: http.StatusUnauthorized,
			wantError: &dto.ErrorsResponse{
				Errors: []string{errorsx.ErrInvalidCredentials.Error()},
			},
		},
		{
			name:       "Success",
			payload:    userData,
			wantStatus: http.StatusOK,
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
}
