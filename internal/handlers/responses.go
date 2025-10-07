package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
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
		http.Error(w, "JSON writing error", http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, dto.ErrorsResponse{
		Errors: formatError(err),
	})
}

func formatError(err error) []string {
	if err.Error() == "EOF" {
		return []string{"Empty JSON"}
	}

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{err.Error()}
	}

	messages := make([]string, len(ve))
	for i, e := range ve {
		msg := fmt.Sprintf("Field '%s' doesn't match the rule '%s'", e.Field(), e.Tag())
		messages[i] = msg
	}

	return messages
}
