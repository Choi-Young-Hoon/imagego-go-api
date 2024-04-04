package image

import (
	"encoding/json"
	"imagego-go-api/database"
	"imagego-go-api/httpserver/jwt"
	"net/http"
)

func ImageAllHandler(res http.ResponseWriter, req *http.Request, claim *jwt.JwtUserCalim) {
	// Get 요청일 때만 처리
	if req.Method != http.MethodGet {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// jwt token에서 user_id를 가져온다.
	userId := claim.UserId
	images, responseCode, err := findImages(userId)
	if err != nil {
		http.Error(res, err.Error(), responseCode)
		return
	}

	imageResponseList := NewImageResponseList(images)
	response, err := json.Marshal(imageResponseList)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(responseCode)
	res.Write(response)
}

func findImages(userId string) ([]database.Image, int, error) {
	image := database.NewImage()
	images, err := image.FindByUserId(userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return images, http.StatusOK, nil
}
