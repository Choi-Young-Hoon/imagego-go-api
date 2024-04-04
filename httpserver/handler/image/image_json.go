package image

import (
	"imagego-go-api/database"
	"imagego-go-api/util"
)

type ImageRequest struct {
	// Update요청 (PUT) 시에만 사용
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type ImageResponse struct {
	Result      string `json:"result"`
	Id          uint   `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"imageUrl,omitempty"`
}

func NewImageResponseList(image []database.Image) []ImageResponse {
	config := util.GetServerConfig()

	var imagesResponseList []ImageResponse
	for _, img := range image {
		imagesResponseList = append(imagesResponseList, ImageResponse{
			Result:      "success",
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
		Result:      "success",
		Id:          image.ID,
		Title:       image.Title,
		Description: image.Description,
		ImageUrl:    config.ImageServerUrl + "/" + image.ImageName,
	}
}

func NewImageDeleteResponse() ImageResponse {
	return ImageResponse{
		Result: "success",
	}
}

func NewImageUpdateResponse() ImageResponse {
	return ImageResponse{
		Result: "success",
	}
}
