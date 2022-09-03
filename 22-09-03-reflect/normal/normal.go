package normal

import (
	"fmt"
	"reflect"
)

type SimpleStruct struct {
	A int
	B int
}

func populateStructReflect(in interface{}) error {
	val := reflect.ValueOf(in)
	if val.Type().Kind() != reflect.Ptr {
		return fmt.Errorf("you must pass in a pointer")
	}
	el := val.Elem()
	if el.Type().Kind() != reflect.Struct {
		return fmt.Errorf("you must pass in a pointer to struct")
	}
	fVal := el.FieldByName("B")
	fVal.SetInt(42)
	return nil
}
