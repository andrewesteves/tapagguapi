package transformers

// CommonTransformer struct
type CommonTransformer struct {
	Current string `json:"current"`
	Prev    string `json:"prev"`
	Next    string `json:"next"`
	Total   string `json:"total"`
}
