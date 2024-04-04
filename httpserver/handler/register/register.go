package register

import (
	"encoding/json"
	"imagego-go-api/database"
	"net/http"
)

func RegisterHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "404 Not Found", http.StatusNotFound)
	}

	registerRequest := &RegisterRequest{}
	err := json.NewDecoder(req.Body).Decode(registerRequest)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	registerResponse := registerHandler(registerRequest.Id, registerRequest.Password)
	jsonData, err := json.Marshal(registerResponse)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonData)
}

func registerHandler(id, password string) *RegisterResponse {
	user := database.NewUser()

	err := user.FindByUserId(id)
	if err == nil || user.UserId == id {
		return NewRegisterFailedResponse("이미 존재하는 사용자 계정 입니다.")
	}

	user.UserId = id
	user.UserPw = password
	err = user.Create()
	if err != nil {
		return NewRegisterFailedResponse("사용자 계정 생성에 실패하였습니다.")
	}

	return NewRegisterSuccessResponse()
}
