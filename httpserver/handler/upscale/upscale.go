package upscale

import (
	"encoding/json"
	"imagego-go-api/httpserver/jwt"
	"net/http"
)

func UpscaleHandler(res http.ResponseWriter, req *http.Request, claim *jwt.JwtUserCalim) {
	if req.Method != http.MethodGet {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	number := req.PathValue("number")

	upscaleRequest := &UpscaleRequest{}
	err := json.NewDecoder(req.Body).Decode(upscaleRequest)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var upscaleResponse *UpscaleResponse
	if upscaleRequest.Type == TYPE_UPSCALE {
		upscaleResponse, err = upscaleHandler(upscaleRequest, number)
	} else {
		upscaleResponse, err = aiUpscale(upscaleRequest, number)
	}

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(upscaleResponse)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonData)
}

func aiUpscale(upscaleRequest *UpscaleRequest, number string) (*UpscaleResponse, error) {
	return nil, nil
}

func upscaleHandler(upscaleRequest *UpscaleRequest, number string) (*UpscaleResponse, error) {
	return nil, nil
}
