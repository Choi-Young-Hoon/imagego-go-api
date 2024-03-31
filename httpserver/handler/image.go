package handler

import (
	"encoding/json"
	"imagego-go-api/database"
	"net/http"
)

func ImageHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "404 Not Found", http.StatusNotFound)
	}
	number := req.PathValue("number")

	image := database.NewImage()
	err := image.FindById(number)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(NewImageResponse(image))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonData)
}
