package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func Validate(r *http.Request) error {

	var SecretKey = viper.GetString("SECRET_KEY")

	reqToken := r.Header.Get("Authorization")

	if reqToken == "" {
		return errors.New("no token provided")
	} else {

		reqToken = strings.Replace(reqToken, "Bearer ", "", -1)

	}

	if reqToken != SecretKey {
		return errors.New("password doesn't match")
	}

	return nil
}
