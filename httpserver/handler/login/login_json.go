package login

type LoginRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Result string `json:"result"`
	Reason string `json:"reason,omitempty"`
	JWT    string `json:"jwt,omitempty"`
}

func NewLoginFailedResponse(reason string) *LoginResponse {
	return &LoginResponse{
		Result: "failed",
		Reason: reason,
	}
}

func NewLoginSuccessResponse(jwt string) *LoginResponse {
	return &LoginResponse{
		Result: "success",
		JWT:    jwt,
	}
}
