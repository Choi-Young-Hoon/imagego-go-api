package login

import (
	"encoding/json"
	"imagego-go-api/database"
	"imagego-go-api/httpserver/jwt"
	"net/http"
)

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "404 Not Found", http.StatusNotFound)
	}

	loginRequest := &LoginRequest{}
	err := json.NewDecoder(req.Body).Decode(loginRequest)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	loginResponse := loginHandler(loginRequest.Id, loginRequest.Password, req)
	jsonData, err := json.Marshal(loginResponse)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}

func loginHandler(id, password string, req *http.Request) *LoginResponse {
	user := database.NewUser()
	isSuccess, err := user.Authorize(id, password)
	if !isSuccess {
		return NewLoginFailedResponse(err.Error())
	}

	jwtToken, err := jwt.GenerateJwtToken(id, req.Host)
	if err != nil {
		return NewLoginFailedResponse(err.Error())
	}

	return NewLoginSuccessResponse(jwtToken)
}
