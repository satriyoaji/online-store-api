package middlewares

import (
	"fmt"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
