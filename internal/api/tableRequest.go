package api

import "errors"

const MAX_PAGE_SIZE = 500

var (
	TableRequestLimitError   = errors.New("Invalid limit value")
	TableRequestPageTooLarge = errors.New("Page size too large")
)

type TableRequest struct {
	Offset  int               `json:"offset"`
	Limit   int               `json:"limit"`
	Filters map[string]string `json:"filters"`
	OrderBy string            `json:"orderBy"`
	Order   string            `json:"order"`
}

func (t *TableRequest) Validate() error {

	if t.Limit == 0 {
		return TableRequestLimitError
	}

	if t.Limit > MAX_PAGE_SIZE {
		return TableRequestPageTooLarge
	}

	return nil
}
