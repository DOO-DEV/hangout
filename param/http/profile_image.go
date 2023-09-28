package param

type SaveProfileImageRequest struct {
}

type SaveProfileImageResponse struct {
	ImageUrl string `json:"image_url"`
}

type GetPrimaryProfileImageRequest struct {
}

type GetPrimaryProfileImageResponse struct {
	Url string `json:"url"`
}

type GetAllProfileImagesRequest struct {
}

type GetAllProfileImagesResponse struct {
	Data []string `json:"data"`
}

type DeleteProfileImageRequest struct {
	ImageID string `json:"image_id"`
}

type DeleteProfileImageResponse struct {
}

type SetImageAsPrimaryRequest struct {
	ImageID string `json:"image_id"`
}

type SetImageAsPrimaryResponse struct {
}
