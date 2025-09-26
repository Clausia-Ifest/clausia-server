package pagination

type PaginationMetaData struct {
	TotalData int64 `json:"total_data"`
	TotalPage int64 `json:"total_page"`
	Limit     int64 `json:"limit"`
	Page      int64 `json:"page"`
}

func Initialize(totalData, limit, page int64) PaginationMetaData {
	var totalPages int64

	if totalData == 0 {
		totalPages = 0
	} else {
		totalPages = totalData / limit
		if totalData%limit != 0 {
			totalPages++
		}
	}

	return PaginationMetaData{
		TotalData: totalData,
		TotalPage: totalPages,
		Limit:     limit,
		Page:      page,
	}
}
