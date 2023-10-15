package steam

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// Query defines from  the url.Values.
type Query url.Values

// Parse parses the given data to a Query.
// Only parses the elements with query tag.
func (q *Query) Parse(data interface{}) error {

	typ := reflect.TypeOf(data)
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil
		}
		value = value.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < value.NumField(); i++ {
		tag := strings.Split(typ.Field(i).Tag.Get("query"), ",")[0]
		v := value.Field(i)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				continue
			}
			v = v.Elem()
		}

		(*url.Values)(q).Set(tag, fmt.Sprintf("%v", v))
	}

	return nil
}
