package tests

import "reflect"

func EqualStruct(first, second interface{}) bool {
	return equalValue(checkPtr(first), checkPtr(second))
}

func checkPtr(val interface{}) reflect.Value {
	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Ptr || rv.IsNil(){
		panic("Pointers are accepted as a models that must not empty")
	}
	return rv.Elem()
}

func formatTypeValue(rv reflect.Value) (reflect.Type, reflect.Value) {
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	return rv.Type(), rv
}

func equalValue(first, second reflect.Value) bool {
	frt, frv := formatTypeValue(first)
	srt, srv := formatTypeValue(second)
	if frt.Kind() != srt.Kind() {
		return false
	}

	for i := 0; i < frt.NumField(); i++ {
		frs := frt.Field(i)
		if _, ok := frs.Tag.Lookup("equal"); !ok {
			continue
		}

		fiv := frv.Field(i)
		if !fiv.CanSet() || frs.Anonymous {
			continue
		}

		siv := srv.Field(i)
		if fiv.Kind() == reflect.Ptr && (fiv.IsNil() || siv.IsNil()) {
			return fiv.IsNil() == siv.IsNil()
		}

		if !equalWithProperType(fiv, siv) {
			return false
		}
	}
	return true
}

func equalWithProperType(first, second reflect.Value) bool {
	frt, frv := formatTypeValue(first)
	_, srv := formatTypeValue(second)

	switch frt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return frv.Int() == srv.Int()
	case reflect.Bool:
		return frv.Bool() == srv.Bool()
	case reflect.Float32, reflect.Float64:
		return frv.Float() == srv.Float()
	case reflect.String:
		return frv.String() == srv.String()
	case reflect.Struct:
		return equalValue(frv, srv)
	case reflect.Map, reflect.Array, reflect.Slice:
		//todo:
	}
	return true
}
