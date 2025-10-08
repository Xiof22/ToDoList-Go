package handlers

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"strconv"
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

func getURLIntParam(r *http.Request, key string) (int, error) {
	vars := mux.Vars(r)
	paramStr := vars[key]
	id, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, errors.New("ID parsing error")
	}

	return id, nil
}
