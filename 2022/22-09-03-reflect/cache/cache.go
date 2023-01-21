package cache

import (
	"fmt"
	"reflect"
)

type SimpleStruct struct {
	A int
	B int
}

var cache = make(map[reflect.Type][]int)

func populateStructReflectCache(in interface{}) error {
	typ := reflect.TypeOf(in)
	index, ok := cache[typ]
	if !ok {
		if typ.Kind() != reflect.Ptr {
			return fmt.Errorf("you must pass in a pointer")
		}
		if typ.Elem().Kind() != reflect.Struct {
			return fmt.Errorf("you must pass in a pointer to struct")
		}
		f, ok := typ.Elem().FieldByName("B")
		if !ok {
			return fmt.Errorf("struct does not have field B")
		}
		index = f.Index
		cache[typ] = index
	}
	val := reflect.ValueOf(in)
	el := val.Elem()
	fVal := el.FieldByIndex(index)
	fVal.SetInt(42)
	return nil
}
