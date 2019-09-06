package pages

import "net/http"

func EmptyHandler(_ *http.Request) (interface{}, int) {
	return struct{}{}, 200
}
