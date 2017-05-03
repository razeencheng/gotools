package structs

import (
	"reflect"
	"strconv"
)

func StructToMap(m interface{}) map[string]string {
	res := map[string]string{}
	mVal := reflect.ValueOf(m).Elem()
	typ := mVal.Type()
	for i := 0; i < mVal.NumField(); i++ {
		field := mVal.Field(i)
		jsonTag := typ.Field(i).Tag.Get("json")
		hidTag := typ.Field(i).Tag.Get("need")

		var v string
		switch field.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(field.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(field.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(field.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(field.Float(), 'f', 4, 64)
		case []byte:
			v = string(field.Bytes())
		case string:
			v = field.String()
		case bool:
			if field.Bool() {
				v = "true"
			} else {
				v = "false"
			}
		}

		if hidTag == "no" && v == "" {
			continue
		}
		res[jsonTag] = v
	}
	return res
}
