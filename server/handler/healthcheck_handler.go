package handler

import (
	"io"
	"net/http"
	"strings"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.Copy(w, strings.NewReader("OK"))
}
