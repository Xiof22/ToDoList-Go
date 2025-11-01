package handlers

import (
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
	"reflect"
	"strings"
)

// TrimStrings removes surrounding spaces from all string fields.
func trimStrings(s any) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return
	}

	v = v.Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.CanSet() {
			field.SetString(strings.TrimSpace(field.String()))
		}
	}
}

func pathID[T ~[16]byte](r *http.Request, key string) (T, error) {
	raw := chi.URLParam(r, key)
	parsed, err := uuid.Parse(raw)
	if err != nil {
		var zero T
		return zero, err
	}

	return T(parsed), nil
}
