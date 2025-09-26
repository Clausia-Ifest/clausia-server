package querybuilder

type QB struct {
	//
}

type IQB interface {
	WhereConditions(s any) string
}

func New() IQB {
	return &QB{}
}
