package handler

import (
	"imagego-go-api/database"
	"imagego-go-api/util"
)

type ImageRequest struct {
}

type ImageResponse struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

func NewImageResponseList(image []database.Image) []ImageResponse {
	config := util.GetServerConfig()

	var imagesResponseList []ImageResponse
	for _, img := range image {
		imagesResponseList = append(imagesResponseList, ImageResponse{
			Id:          img.ID,
			Title:       img.Title,
			Description: img.Description,
			ImageUrl:    config.ImageServerUrl + "/" + img.ImageName,
		})
	}
	return imagesResponseList
}

func NewImageResponse(image database.Image) ImageResponse {
	config := util.GetServerConfig()

	return ImageResponse{
		Id:          image.ID,
		Title:       image.Title,
		Description: image.Description,
		ImageUrl:    config.ImageServerUrl + "/" + image.ImageName,
	}
}
