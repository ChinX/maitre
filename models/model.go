package models

type FieldType string

const (
	DEFAULT                  FieldType = "default"
	UPPER                    FieldType = "upper"
	LOWER                    FieldType = "lower"
	NUMBER                   FieldType = "number"
	UNDERLINE                FieldType = "underline"
	MIDDLELINE               FieldType = "middleline"
	SYMBOL                   FieldType = "symbol"
	CHINESE                  FieldType = "chinese"
	PRE_UPPER                FieldType = "pre_upper"
	PRE_LOWER                FieldType = "pre_lower"
	PRE_NUMBER               FieldType = "pre_number"
	PRE_UNDERLINE            FieldType = "pre_underline"
	PRE_MIDDLELINE           FieldType = "pre_middleline"
	PRE_DOUBLE_UNDERLINE     FieldType = "pre_double_underline"
	PRE_DOUBLE_MIDDLELINE    FieldType = "pre_double_middleline"
	PRE_YMBOL                FieldType = "inside_symbol"
	PRE_CHINESE              FieldType = "inside_chinese"
	INSIDE_UPPER             FieldType = "inside_upper"
	INSIDE_LOWER             FieldType = "inside_lower"
	INSIDE_NUMBER            FieldType = "inside_number"
	INSIDE_UNDERLINE         FieldType = "inside_underline"
	INSIDE_MIDDLELINE        FieldType = "inside_middleline"
	INSIDE_DOUBLE_UNDERLINE  FieldType = "inside_double_underline"
	INSIDE_DOUBLE_MIDDLELINE FieldType = "inside_double_middleline"
	INSIDE_SYMBOL            FieldType = "inside_symbol"
	INSIDE_CHINESE           FieldType = "inside_chinese"
	SUF_UPPER                FieldType = "suf_upper"
	SUF_LOWER                FieldType = "suf_lower"
	SUF_NUMBER               FieldType = "suf_number"
	SUF_UNDERLINE            FieldType = "suf_underline"
	SUF_MIDDLELINE           FieldType = "suf_middleline"
	SUF_DOUBLE_UNDERLINE     FieldType = "suf_double_underline"
	SUF_DOUBLE_MIDDLELINE    FieldType = "suf_double_middleline"
	SUF_SYMBOL               FieldType = "suf_symbol"
	SUF_CHINESE              FieldType = "suf_chinese"
)

var fieldList = map[string]map[FieldType]string{}

func Register(typ FieldType, field, val string) {
	fields, ok := fieldList[field]
	if !ok {
		fields = make(map[FieldType]string)
		fieldList[field] = fields
	}

	fields[typ] = val
}

func GetFieldValue(typ FieldType, field string) (val string) {
	if fields, ok := fieldList[field]; ok {
		val = fields[typ]
	}
	return
}
