package api

type InRequest struct {
	Source  RequestSource     `json:"source"`
	Version map[string]string `json:"version"`
}

type OutRequest struct {
	Params  OutParams         `json:"params"`
	Source  RequestSource     `json:"source"`
	Version map[string]string `json:"version"`
}

type RequestSource struct {
	Files []RequestSourceFile `json:"files"`
}

type RequestSourceFile struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
}

type OutParams struct {
}
