package entity

type Pagination struct {
	CurrentPage   int `json:"current_page"`
	TotalPages    int `json:"total_pages"`
	TotalElements int `json:"total_elements"`
}
