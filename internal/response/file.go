package response

type UploadFileResponse struct {
	Path       string `json:"path,omitempty"`
	SourceName string `json:"name"`
}
