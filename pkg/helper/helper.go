package helper

type PagedResponse struct {
	Data     interface{} `json:"data"`
	Page     int         `json:"page"`
	Size     int         `json:"size"`
	Next     bool        `json:"next"`
	Previous bool        `json:"previous"`
	Total    int         `json:"total"`
}

func GetLimitOffset(page, size int) (limit int, offset int) {
	if page == 0 || size == 0 {
		// using -1 to disable gorm size and offset in case page and size not set
		size = -1
		offset = -1
		return size, offset
	}
	offset = (page - 1) * size
	return size, offset
}
