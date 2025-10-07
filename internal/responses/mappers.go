package responses

import "net/http"

func MapError(err error) int {
	switch {

	default:
		return http.StatusInternalServerError

	}
}
