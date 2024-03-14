package response

type GetPublicToolResponse struct {
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	URL         string `json:"url"`
	Params      string `json:"params"`
	Description string `json:"description"`
	Disable     bool   `json:"disable"`
}
