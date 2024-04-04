package echo

import (
	"imagego-go-api/httpserver/jwt"
	"net/http"
)

func EchoHandler(res http.ResponseWriter, req *http.Request, claim *jwt.JwtUserCalim) {
	http.Error(res, "200 OK", http.StatusOK)
}
