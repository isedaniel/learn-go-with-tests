package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	v := value(x)

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			walk(v.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			walk(v.MapIndex(key).Interface(), fn)
		}
    case reflect.Chan:
        for r, ok := v.Recv(); ok; r, ok = v.Recv() {
            walk(r.Interface(), fn)
        }
    case reflect.Func:
        r := v.Call(nil)
        for _, res := range r {
            walk(res.Interface(), fn)
        }
	}
}

func value(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}
