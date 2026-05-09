package types

type BasePaginatedResult struct {
	Page       int64 `json:"page"`
	Count      int64 `json:"count"`
	TotalPages int64 `json:"total_pages"`
	TotalCount int64 `json:"total_count"`
}

type BasePaginationRequest struct {
	Page  int64
	Limit int64
}
