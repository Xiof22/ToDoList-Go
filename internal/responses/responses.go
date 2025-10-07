package responses

import (
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	var err error
	msg, ok := data.(string)
	if ok {
		err = json.NewEncoder(w).Encode(map[string]string{
			"message": msg,
		})
	} else {
		err = json.NewEncoder(w).Encode(data)
	}

	if err != nil {
		http.Error(w, errorsx.ErrWriteJSON.Error(), http.StatusInternalServerError)
	}
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, dto.ErrorsResponse{
		Errors: formatError(err),
	})
}

func formatError(err error) []string {
	if err.Error() == "EOF" {
		return []string{errorsx.ErrMissingJSON.Error()}
	}

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{err.Error()}
	}

	messages := make([]string, len(ve))
	for i, e := range ve {
		messages[i] = errorsx.ErrValidation(e.Field(), e.Tag()).Error()
	}

	return messages
}
