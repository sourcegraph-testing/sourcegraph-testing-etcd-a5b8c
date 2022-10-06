package reflect2

import (
	"reflect"
	"unsafe"
)

type safeType struct {
	reflect.Type
	cfg *frozenConfig
}

func (type2 *safeType) New() any {
	return reflect.New(type2.Type).Interface()
}

func (type2 *safeType) UnsafeNew() unsafe.Pointer {
	panic("does not support unsafe operation")
}

func (type2 *safeType) Elem() Type {
	return type2.cfg.Type2(type2.Type.Elem())
}

func (type2 *safeType) Type1() reflect.Type {
	return type2.Type
}

func (type2 *safeType) PackEFace(ptr unsafe.Pointer) any {
	panic("does not support unsafe operation")
}

func (type2 *safeType) Implements(thatType Type) bool {
	return type2.Type.Implements(thatType.Type1())
}

func (type2 *safeType) RType() uintptr {
	panic("does not support unsafe operation")
}

func (type2 *safeType) Indirect(obj any) any {
	return reflect.Indirect(reflect.ValueOf(obj)).Interface()
}

func (type2 *safeType) UnsafeIndirect(ptr unsafe.Pointer) any {
	panic("does not support unsafe operation")
}

func (type2 *safeType) LikePtr() bool {
	panic("does not support unsafe operation")
}

func (type2 *safeType) IsNullable() bool {
	return IsNullable(type2.Kind())
}

func (type2 *safeType) IsNil(obj any) bool {
	if obj == nil {
		return true
	}
	return reflect.ValueOf(obj).Elem().IsNil()
}

func (type2 *safeType) UnsafeIsNil(ptr unsafe.Pointer) bool {
	panic("does not support unsafe operation")
}

func (type2 *safeType) Set(obj any, val any) {
	reflect.ValueOf(obj).Elem().Set(reflect.ValueOf(val).Elem())
}

func (type2 *safeType) UnsafeSet(ptr unsafe.Pointer, val unsafe.Pointer) {
	panic("does not support unsafe operation")
}

func (type2 *safeType) AssignableTo(anotherType Type) bool {
	return type2.Type1().AssignableTo(anotherType.Type1())
}
