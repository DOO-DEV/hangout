package param

type GetAllProfileImagesRequest struct {
}

type GetAllProfileImagesResponse struct {
	Data []string `json:"data"`
}

type DeleteProfileImage struct {
	ImageID string `json:"image_id"`
}

type SetImageAsPrimaryRequest struct {
	ImageID string `json:"image_id"`
}

type SetImageAsPrimaryResponse struct {
}
