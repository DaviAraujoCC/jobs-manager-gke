package handler

import (
	"io"
	"net/http"
	"strings"
)

// TODO: Implement this function
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.Copy(w, strings.NewReader("OK"))
}
