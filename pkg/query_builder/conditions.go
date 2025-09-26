package querybuilder

import (
	"fmt"
	"reflect"
	"strings"
)

func (qb *QB) WhereConditions(s any) string {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	r := make([]string, 0)
	for i := range v.NumField() {
		if !reflect.ValueOf(v.Field(i).Interface()).IsZero() {
			tag := t.Field(i).Tag.Get("db")
			r = append(r, fmt.Sprintf("%s = :%s", tag, tag))
		}
	}

	return strings.Trim(strings.Join(r, " AND "), " ")
}
