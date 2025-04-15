package reflection

import (
	"fmt"
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := getvalue(x)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			Walk(val.Field(i).Interface(), fn)
		}
	case reflect.Int:
		fn(fmt.Sprint(val.Int()))
	case reflect.Func:

	case reflect.String:
		fn(val.String())
	}
}

func getvalue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
