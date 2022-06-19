package pagination

// Request ...
type Request struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// Response ...
type Response struct {
	Limit       int   `json:"limit"`
	TotalPage   uint  `json:"total_page"`
	TotalRows   int64 `json:"total_rows"`
	CurrentPage int   `json:"current_page"`
}
