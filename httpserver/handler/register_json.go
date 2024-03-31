package handler

type RegisterRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Result string `json:"result"`
	Reason string `json:"reason,omitempty"`
}

func NewRegisterFailedResponse(reason string) *RegisterResponse {
	return &RegisterResponse{
		Result: "failed",
		Reason: reason,
	}
}

func NewRegisterSuccessResponse() *RegisterResponse {
	return &RegisterResponse{
		Result: "success",
	}
}
