package dto

type Document struct {
	Hash     string `json:"hash"`
	URL      string `json:"url"`
	Content  string `json:"content"`
	MetaData string `json:"meta_data"`
	Category string `json:"category"`
}

type ExtractDocumentResponse struct {
	MetaData string `json:"meta_data"`
}

type GetDocumentParams struct {
	Hash string `db:"hash"`
}
