package dto

type (
	FileExitsRequest struct {
		Path string `param:"path" valid:"required"`
	}

	FileExitsResponse struct {
		Exists bool `json:"exists"`
	}

	FileGetRequest struct {
		Path string `param:"path" valid:"required"`
	}
)
