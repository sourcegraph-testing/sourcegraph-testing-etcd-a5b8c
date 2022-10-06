package reflect2

import (
	"reflect"
	"unsafe"
)

type eface struct {
	rtype unsafe.Pointer
	data  unsafe.Pointer
}

func unpackEFace(obj any) *eface {
	return (*eface)(unsafe.Pointer(&obj))
}

func packEFace(rtype unsafe.Pointer, data unsafe.Pointer) any {
	var i any
	e := (*eface)(unsafe.Pointer(&i))
	e.rtype = rtype
	e.data = data
	return i
}

type UnsafeEFaceType struct {
	unsafeType
}

func newUnsafeEFaceType(cfg *frozenConfig, type1 reflect.Type) *UnsafeEFaceType {
	return &UnsafeEFaceType{
		unsafeType: *newUnsafeType(cfg, type1),
	}
}

func (type2 *UnsafeEFaceType) IsNil(obj any) bool {
	if obj == nil {
		return true
	}
	objEFace := unpackEFace(obj)
	assertType("Type.IsNil argument 1", type2.ptrRType, objEFace.rtype)
	return type2.UnsafeIsNil(objEFace.data)
}

func (type2 *UnsafeEFaceType) UnsafeIsNil(ptr unsafe.Pointer) bool {
	if ptr == nil {
		return true
	}
	return unpackEFace(*(*any)(ptr)).data == nil
}

func (type2 *UnsafeEFaceType) Indirect(obj any) any {
	objEFace := unpackEFace(obj)
	assertType("Type.Indirect argument 1", type2.ptrRType, objEFace.rtype)
	return type2.UnsafeIndirect(objEFace.data)
}

func (type2 *UnsafeEFaceType) UnsafeIndirect(ptr unsafe.Pointer) any {
	return *(*any)(ptr)
}
