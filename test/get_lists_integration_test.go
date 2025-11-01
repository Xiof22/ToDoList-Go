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

func TestGetLists(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	client := ts.Client()

	url := fmt.Sprintf("%s/lists", ts.URL)

	t.Run("No lists", func(t *testing.T) {

		resp, err := client.Get(url)
		require.NoError(t, err)
		defer resp.Body.Close()

		wantResponse := dto.ListsResponse{
			Count: 0,
			Lists: []dto.List{},
		}

		var gotResponse dto.ListsResponse
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&gotResponse))

		assert.Equal(t, gotResponse, wantResponse)
	})

	t.Run("Have list", func(t *testing.T) {
		createList(t, client, ts.URL, sampleListMap)

		resp, err := client.Get(url)
		require.NoError(t, err)
		defer resp.Body.Close()

		wantResponse := dto.ListsResponse{
			Count: 1,
			Lists: []dto.List{sampleList},
		}

		var gotResponse dto.ListsResponse
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&gotResponse))

		assert.Equal(t, gotResponse.Count, wantResponse.Count)
		assert.Equal(t, gotResponse.Lists[0].Title, wantResponse.Lists[0].Title)
		assert.Greater(t, gotResponse.Lists[0].ID, 0)
	})
}
