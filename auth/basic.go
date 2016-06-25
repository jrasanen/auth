package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func init() {
	username, password := getAuthData()

	if (username == "") || (password == "") {
		panic("Username or password not found")
	} else {
		fmt.Println(username)
		fmt.Println(password)
	}
}

func getAuthData() (string, string) {
	return os.Getenv("AUTH_USER"), os.Getenv("AUTH_PASS")
}

func denyAccess(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", 401)
	return
}

func authenticate(credentials string) bool {
	if credentials == "" {
		return false
	}

	data, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		fmt.Println("error:", err)
		return false
	}

	auth := strings.Split(string(data), ":")

	username, password := getAuthData()

	return auth[0] == username && auth[1] == password
}

// AuthorizeBasic is a middleware which verifies user against set env vars
func AuthorizeBasic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authData := r.Header.Get("Authorization")
		if authData == "" {
			denyAccess(w)
		}
		s := strings.Split(authData, " ")
		if s == nil || len(s) <= 1 {
			denyAccess(w)
		}
		auth := s[1]
		if authenticate(auth) {
			next.ServeHTTP(w, r)
		} else {
			denyAccess(w)
		}
	})
}
