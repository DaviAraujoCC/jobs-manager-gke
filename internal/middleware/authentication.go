package middleware

import (
	"net/http"

	auth "github.com/hurbcom/jobs-manager-gke/internal/auth"
	logrus "github.com/sirupsen/logrus"
)

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.Validate(r); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			logrus.Infof("Unauthorized request: %s", err)
			return
		}
		nextFunc(w, r)
	}
}
