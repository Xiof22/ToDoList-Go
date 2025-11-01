package dto

type ListsResponse struct {
	Count int    `json:"lists_count"`
	Lists []List `json:"lists"`
}
