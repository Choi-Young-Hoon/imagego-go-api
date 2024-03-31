package handler

import (
	"imagego-go-api/httpserver/jwt"
	"net/http"
)

func EchoHandler(res http.ResponseWriter, req *http.Request) {
	bearer := req.Header.Values("Authorization")
	// Bearer 에서 jwt 토큰을 추출
	jwtToken := bearer[0][7:]

	_, err := jwt.ValidateJwtToken(jwtToken)
	if err != nil {
		http.Error(res, err.Error(), http.StatusUnauthorized)
		return
	}

	http.Error(res, "200 OK", http.StatusOK)
}
