package api

type Response struct {
	Version ResponseVersion `json:"version"`
}

type ResponseVersion struct {
	Version string `json:"version"`
}
