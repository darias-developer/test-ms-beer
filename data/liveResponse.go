package data

type LiveResponse struct {
	Success bool               `json:"success"`
	Quotes  map[string]float32 `json:"quotes"`
	Error   ErrorLive          `json:"error"`
}

type ErrorLive struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}
