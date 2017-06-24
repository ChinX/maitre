package tests

import "reflect"

func EqualStruct(s, v interface{}) bool {
	return equalValue(reflect.ValueOf(s), reflect.ValueOf(v))
}

func convNoptr(val reflect.Value) (reflect.Value, reflect.Type) {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val, val.Type()
}

func equalValue(sValue, vValue reflect.Value) bool {
	var typ reflect.Type
	sValue, typ = convNoptr(sValue)
	vValue, _ = convNoptr(vValue)
	for i := 0; i < typ.NumField(); i++ {
		typFiled := typ.Field(i)
		valFiled := sValue.Field(i)
		if !valFiled.CanSet() {
			continue
		}
		if _, ok := typFiled.Tag.Lookup("equal"); !ok {
			continue
		}
		equal := equalWithProperType(typFiled, valFiled, vValue.Field(i))
		if !equal {
			return false
		}
	}
	return true
}

func equalWithProperType(typeField reflect.StructField, sValue, vValue reflect.Value) bool {
	kind := typeField.Type.Kind()
	if kind == reflect.Ptr{
		if !typeField.Anonymous {
			return true
		}
		sValue = sValue.Elem()
		vValue = vValue.Elem()
		kind = sValue.Kind()
	}
	switch typeField.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return sValue.Int() == vValue.Int()
	case reflect.Bool:
		return sValue.Bool() == vValue.Bool()
	case reflect.Float32, reflect.Float64:
		return sValue.Float() == vValue.Float()
	case reflect.String:
		return sValue.String() == vValue.String()
	case reflect.Struct:
		return equalValue(sValue, vValue)
	case reflect.Map, reflect.Array, reflect.Slice:
		//todo:
	}
	return true
}
