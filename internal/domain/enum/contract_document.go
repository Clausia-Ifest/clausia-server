package enum

type DocumentCategory int8

const (
	DocumentCategoryAdministration DocumentCategory = 0
	DocumentCategoryLegal          DocumentCategory = 1
	DocumentCategoryTechnical      DocumentCategory = 2
	DocumentCategoryFinancial      DocumentCategory = 3
)

var (
	DocumentCategoryMap = map[DocumentCategory]string{
		DocumentCategoryAdministration: "Administration Document",
		DocumentCategoryLegal:          "Legal Document",
		DocumentCategoryTechnical:      "Technical Document",
		DocumentCategoryFinancial:      "Financial Document",
	}
)

func (s DocumentCategory) String() string {
	if val, ok := DocumentCategoryMap[s]; ok {
		return val
	}

	return ""
}
