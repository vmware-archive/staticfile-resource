package api

type InRequest struct {
	Source  RequestSource     `json:"source"`
	Version map[string]string `json:"version"`
}

type RequestSource struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
}
