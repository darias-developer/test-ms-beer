package data

type ListResponse struct {
	Success    bool              `json:"success"`
	Currencies map[string]string `json:"currencies"`
	Error      ErrorLive         `json:"error"`
}
