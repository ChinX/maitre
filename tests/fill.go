package tests

import (
	"reflect"
	"strconv"
	"strings"
)

func FillPath(path string, s interface{}) (str string) {
	rv := checkPtr(s)
	paths := strings.Split(strings.Trim(path, "/"), "/")
	for i := 0; i < len(paths); i++ {
		item := paths[i]
		if item[0] == ':' {
			item = getValByName(item[1:], rv)
		}
		str += "/" + item
	}
	return
}

func getValByName(name string, rv reflect.Value) string {
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		rsi := rt.Field(i)
		if rsi.Anonymous || rsi.Tag.Get("path") != name {
			continue
		}
		rvi := rv.Field(i)
		if rvi.Kind() != reflect.Ptr || !rvi.IsNil() && rvi.CanSet() {
			str := valWithProperType(rvi)

			return str
		}
	}
	return ""
}

func valWithProperType(rv reflect.Value) string {
	rt, rv := formatTypeValue(rv)

	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'f', 2, 64)
	case reflect.String:
		return rv.String()
	}
	return ""
}
