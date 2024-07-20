package models

type Page struct {
	Result        interface{} `json:"result"`
	Page          int         `json:"page"`
	Size          int         `json:"size"`
	TotalElements int         `json:"total_elements"`
}

type SortField struct {
	SortBy    string
	Direction string
}
