package models

type FieldTyp string

const (
	PRE_UPPER             FieldTyp = "pre_upper"
	PRE_LOWER             FieldTyp = "pre_lower"
	PRE_NUMBER            FieldTyp = "pre_number"
	PRE_UNDERLINE         FieldTyp = "pre_underline"
	PRE_MIDDLELINE        FieldTyp = "pre_middleline"
	PRE_DOUBLE_UNDERLINE  FieldTyp = "pre_double_underline"
	PRE_DOUBLE_MIDDLELINE FieldTyp = "pre_double_middleline"
	PRE_SYMBOL            FieldTyp = "pre_symbol"
	PRE_CHINESE           FieldTyp = "pre_chinese"
)

var fieldList = map[string]map[FieldTyp]string{}

func Register(typ FieldTyp, field, val string) {
	fields, ok := fieldList[field]
	if !ok {
		fields = make(map[FieldTyp]string)
		fieldList[field] = fields
	}

	fields[typ] = val
}

func GetFieldValue(typ FieldTyp, field string) (val string) {
	if fields, ok := fieldList[field]; ok {
		val = fields[typ]
	}
	return
}
