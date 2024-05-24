package middlewares

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthGuard(r *http.Request) bool {
	res := true
	auth := r.Header.Get("Authorization")
	token := auth[6:]

	raw, err := base64.StdEncoding.DecodeString(token)

	if err != nil {
		res = false
	}

	rawStr := string(raw)

	authInfo := strings.Split(rawStr, ":")
	user := authInfo[0]
	pass := authInfo[1]

	if user != "vContract" || pass != "vContract@123" {
		res = false
	}

	return res

}
