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
	case reflect.Map:
		for _, keys := range val.MapKeys() {
			Walk(val.MapIndex(keys).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if V, ok := val.Recv(); ok {
				Walk(V.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Int:
		fn(fmt.Sprint(val.Int()))
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			Walk(res.Interface(), fn)
		}
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
