package tests

import (
	"strings"
	"reflect"
	"strconv"
)

func FillRestPath(path string, s interface{}) (str string) {
	paths := strings.Split(path, "/")
	for i := 1; i < len(paths); i++{
		item := paths[i]
		if item[0] == ':'{
			item = getValFromStruct(item[1:],s)
		}
		str += "/" + item
	}
	return
}

func getValFromStruct(name string, s interface{}) string {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		typFiled := typ.Field(i)
		if typFiled.Tag.Get("restful") == name {
			return strWithProperType(typFiled, val.Field(i))
		}
	}
	return ""
}

func strWithProperType(typ reflect.StructField, val reflect.Value) string {
	kind := typ.Type.Kind()
	if kind == reflect.Ptr{
		if !typ.Anonymous {
			return ""
		}
		val = val.Elem()
		kind = val.Kind()
	}
	switch typ.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', 2, 64)
	case reflect.String:
		return val.String()
	}
	return ""
}