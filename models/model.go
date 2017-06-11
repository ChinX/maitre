package models

type FieldType string

const (
	DEFAULT                        FieldType = "default"
	UPPER                          FieldType = "upper"
	LOWER                          FieldType = "lower"
	NUMBER                         FieldType = "number"
	SYMBOL                         FieldType = "symbol"
	CHINESE                        FieldType = "chinese"
	UNDERLINE                      FieldType = "underline"
	MIDDLELINE                     FieldType = "middleline"
	DOUBLE_UNDERLINE               FieldType = "double_underline"
	DOUBLE_MIDDLELINE              FieldType = "double_middleline"
	DOUBLE_UNDER_MIDDLE            FieldType = "double_under_middle"
	DOUBLE_MIDDLE_UNDER            FieldType = "double_middle_under"
	MIX_UNDER_MIDDLE               FieldType = "mix_under_middle"
	MIX_MIDDLE_UNDER               FieldType = "mix_middle_under"
	MIX_DOUBLE_UNDER_MIDDLE        FieldType = "mix_double_under_middle"
	MIX_DOUBLE_MIDDLE_UNDER        FieldType = "mix_double_middle_under"
	PRE_UPPER                      FieldType = "pre_upper"
	PRE_LOWER                      FieldType = "pre_lower"
	PRE_NUMBER                     FieldType = "pre_number"
	PRE_YMBOL                      FieldType = "pre_ymbol"
	PRE_CHINESE                    FieldType = "pre_chinese"
	PRE_UNDERLINE                  FieldType = "pre_underline"
	PRE_MIDDLELINE                 FieldType = "pre_middleline"
	PRE_DOUBLE_UNDERLINE           FieldType = "pre_double_underline"
	PRE_DOUBLE_MIDDLELINE          FieldType = "pre_double_middleline"
	PRE_DOUBLE_UNDER_MIDDLE        FieldType = "pre_double_under_middle"
	PRE_DOUBLE_MIDDLE_UNDER        FieldType = "pre_double_middle_under"
	PRE_MIX_UNDER_MIDDLE           FieldType = "pre_mix_under_middle"
	PRE_MIX_MIDDLE_UNDER           FieldType = "pre_mix_middle_under"
	PRE_MIX_DOUBLE_UNDER_MIDDLE    FieldType = "pre_mix_double_under_middle"
	PRE_MIX_DOUBLE_MIDDLE_UNDER    FieldType = "pre_mix_double_middle_under"
	INSIDE_UPPER                   FieldType = "inside_upper"
	INSIDE_LOWER                   FieldType = "inside_lower"
	INSIDE_NUMBER                  FieldType = "inside_number"
	INSIDE_SYMBOL                  FieldType = "inside_symbol"
	INSIDE_CHINESE                 FieldType = "inside_chinese"
	INSIDE_UNDERLINE               FieldType = "inside_underline"
	INSIDE_MIDDLELINE              FieldType = "inside_middleline"
	INSIDE_DOUBLE_UNDERLINE        FieldType = "inside_double_underline"
	INSIDE_DOUBLE_MIDDLELINE       FieldType = "inside_double_middleline"
	INSIDE_DOUBLE_UNDER_MIDDLE     FieldType = "inside_double_under_middle"
	INSIDE_DOUBLE_MIDDLE_UNDER     FieldType = "inside_double_middle_under"
	INSIDE_MIX_UNDER_MIDDLE        FieldType = "inside_mix_under_middle"
	INSIDE_MIX_MIDDLE_UNDER        FieldType = "inside_mix_middle_under"
	INSIDE_MIX_DOUBLE_UNDER_MIDDLE FieldType = "inside_mix_double_under_middle"
	INSIDE_MIX_DOUBLE_MIDDLE_UNDER FieldType = "inside_mix_double_middle_under"
	SUF_UPPER                      FieldType = "suf_upper"
	SUF_LOWER                      FieldType = "suf_lower"
	SUF_NUMBER                     FieldType = "suf_number"
	SUF_SYMBOL                     FieldType = "suf_symbol"
	SUF_CHINESE                    FieldType = "suf_chinese"
	SUF_UNDERLINE                  FieldType = "suf_underline"
	SUF_MIDDLELINE                 FieldType = "suf_middleline"
	SUF_DOUBLE_UNDERLINE           FieldType = "suf_double_underline"
	SUF_DOUBLE_MIDDLELINE          FieldType = "suf_double_middleline"
	SUF_DOUBLE_UNDER_MIDDLE        FieldType = "suf_double_under_middle"
	SUF_DOUBLE_MIDDLE_UNDER        FieldType = "suf_double_middle_under"
	SUF_MIX_UNDER_MIDDLE           FieldType = "suf_mix_under_middle"
	SUF_MIX_MIDDLE_UNDER           FieldType = "suf_mix_middle_under"
	SUF_MIX_DOUBLE_UNDER_MIDDLE    FieldType = "suf_mix_double_under_middle"
	SUF_MIX_DOUBLE_MIDDLE_UNDER    FieldType = "suf_mix_double_middle_under"
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
