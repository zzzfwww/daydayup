package unsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)

type SimpleStruct struct {
	A int
	B int
}

var unsafeCache = make(map[reflect.Type]uintptr)

type intface struct {
	typ   unsafe.Pointer
	value unsafe.Pointer
}

func populateStruct(in *SimpleStruct) {
	in.B = 42
}

func populateStructUnsafe(in interface{}) error {
	typ := reflect.TypeOf(in)

	offset, ok := unsafeCache[typ]
	if !ok {
		if typ.Kind() != reflect.Ptr {
			return fmt.Errorf("you must pass in a pointer")
		}
		if typ.Elem().Kind() != reflect.Struct {
			return fmt.Errorf("you must pass in a pointer to a struct")
		}
		f, ok := typ.Elem().FieldByName("B")
		if !ok {
			return fmt.Errorf("struct does not have field B")
		}
		if f.Type.Kind() != reflect.Int {
			return fmt.Errorf("field B should be an int")
		}
		offset = f.Offset
		unsafeCache[typ] = offset
	}
	structPtr := (*intface)(unsafe.Pointer(&in)).value
	*(*int)(unsafe.Pointer(uintptr(structPtr) + offset)) = 42
	return nil
}

var unsafeCache2 = make(map[uintptr]uintptr)

func populateStructUnsafe2(in interface{}) error {
	inf := (*intface)(unsafe.Pointer(&in))

	offset, ok := unsafeCache2[uintptr((inf.typ))]
	if !ok {
		typ := reflect.TypeOf(in)
		if typ.Kind() != reflect.Ptr {
			return fmt.Errorf("you must pass in a pointer")
		}
		if typ.Elem().Kind() != reflect.Struct {
			return fmt.Errorf("you must pass in a pointer to a struct ")
		}
		f, ok := typ.Elem().FieldByName("B")
		if !ok {
			return fmt.Errorf("struct does not have field B")
		}
		if f.Type.Kind() != reflect.Int {
			return fmt.Errorf("field B should be an int")
		}
		offset = f.Offset
		unsafeCache2[uintptr(inf.typ)] = offset
	}
	*(*int)(unsafe.Pointer(uintptr(inf.value) + offset)) = 42
	return nil
}

type typeDescriptor uintptr

func describeType(in interface{}) (typeDescriptor, error) {
	typ := reflect.TypeOf(in)
	if typ.Kind() != reflect.Ptr {
		return 0, fmt.Errorf("you must pass in a pointer")
	}
	if typ.Elem().Kind() != reflect.Struct {
		return 0, fmt.Errorf("you must pass in a pointer to a struct")
	}
	f, ok := typ.Elem().FieldByName("B")
	if !ok {
		return 0, fmt.Errorf("struct does not have field B")
	}
	if f.Type.Kind() != reflect.Int {
		return 0, fmt.Errorf("field B should be an int")
	}
	return typeDescriptor(f.Offset), nil
}

func populateStructUnsafe3(in interface{}, ti typeDescriptor) error {
	structPtr := (*intface)(unsafe.Pointer(&in)).value
	*(*int)(unsafe.Pointer(uintptr(structPtr) + uintptr(ti))) = 42
	return nil
}
