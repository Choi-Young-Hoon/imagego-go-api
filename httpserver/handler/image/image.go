package image

import (
	"encoding/json"
	"imagego-go-api/database"
	"imagego-go-api/httpserver/jwt"
	"net/http"
)

func ImageHandler(res http.ResponseWriter, req *http.Request, claim *jwt.JwtUserCalim) {
	number := req.PathValue("number")

	if req.Method == http.MethodGet {
		FindImage(res, number, database.NewImage())
		return
	} else if req.Method == http.MethodDelete {
		DeleteImage(res, number, database.NewImage())
		return
	} else if req.Method == http.MethodPut {
		UpdateImage(res, req, number, database.NewImage())
		return
	}

	http.Error(res, "404 Not Found", http.StatusNotFound)
}

func FindImage(res http.ResponseWriter, number string, image database.Image) {
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

func DeleteImage(res http.ResponseWriter, number string, image database.Image) {
	err := image.DeleteById(number)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(NewImageDeleteResponse())
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonData)
}

func UpdateImage(res http.ResponseWriter, req *http.Request, number string, image database.Image) {
	request := &ImageRequest{}
	// req 에서 json 데이터를 읽어온다.
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	image.Title = request.Title
	image.Description = request.Description

	err = image.UpdateById(number)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(NewImageUpdateResponse())
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonData)
}
