package upscale

const TYPE_AI_UPSCALE = "ai_upscale"
const TYPE_UPSCALE = "upscale"

type UpscaleRequest struct {
	Type string `json:"type"`

	// AI upscaling 요청 시에만 사용
	Scale int `json:"scale,omitempty"`

	// Upscale 요청 시에만 사용
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

func NewUpscaleSuccessResponse(downloadUrl string) *UpscaleResponse {
	return &UpscaleResponse{
		Result:      "success",
		DownloadUrl: downloadUrl,
	}
}

func NewUpscaleFailedResponse(reason string) *UpscaleResponse {
	return &UpscaleResponse{
		Result: "failed",
		Reason: reason,
	}
}

type UpscaleResponse struct {
	Result      string `json:"result"`
	Reason      string `json:"reason,omitempty"`
	DownloadUrl string `json:"url,omitempty"`
}
