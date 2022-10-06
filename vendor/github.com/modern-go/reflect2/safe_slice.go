package reflect2

import (
	"reflect"
	"unsafe"
)

type safeSliceType struct {
	safeType
}

func (type2 *safeSliceType) SetIndex(obj any, index int, value any) {
	val := reflect.ValueOf(obj).Elem()
	elem := reflect.ValueOf(value).Elem()
	val.Index(index).Set(elem)
}

func (type2 *safeSliceType) UnsafeSetIndex(obj unsafe.Pointer, index int, value unsafe.Pointer) {
	panic("does not support unsafe operation")
}

func (type2 *safeSliceType) GetIndex(obj any, index int) any {
	val := reflect.ValueOf(obj).Elem()
	elem := val.Index(index)
	ptr := reflect.New(elem.Type())
	ptr.Elem().Set(elem)
	return ptr.Interface()
}

func (type2 *safeSliceType) UnsafeGetIndex(obj unsafe.Pointer, index int) unsafe.Pointer {
	panic("does not support unsafe operation")
}

func (type2 *safeSliceType) MakeSlice(length int, cap int) any {
	val := reflect.MakeSlice(type2.Type, length, cap)
	ptr := reflect.New(val.Type())
	ptr.Elem().Set(val)
	return ptr.Interface()
}

func (type2 *safeSliceType) UnsafeMakeSlice(length int, cap int) unsafe.Pointer {
	panic("does not support unsafe operation")
}

func (type2 *safeSliceType) Grow(obj any, newLength int) {
	oldCap := type2.Cap(obj)
	oldSlice := reflect.ValueOf(obj).Elem()
	delta := newLength - oldCap
	deltaVals := make([]reflect.Value, delta)
	newSlice := reflect.Append(oldSlice, deltaVals...)
	oldSlice.Set(newSlice)
}

func (type2 *safeSliceType) UnsafeGrow(ptr unsafe.Pointer, newLength int) {
	panic("does not support unsafe operation")
}

func (type2 *safeSliceType) Append(obj any, elem any) {
	val := reflect.ValueOf(obj).Elem()
	elemVal := reflect.ValueOf(elem).Elem()
	newVal := reflect.Append(val, elemVal)
	val.Set(newVal)
}

func (type2 *safeSliceType) UnsafeAppend(obj unsafe.Pointer, elem unsafe.Pointer) {
	panic("does not support unsafe operation")
}

func (type2 *safeSliceType) SetNil(obj any) {
	val := reflect.ValueOf(obj).Elem()
	val.Set(reflect.Zero(val.Type()))
}

func (type2 *safeSliceType) UnsafeSetNil(ptr unsafe.Pointer) {
	panic("does not support unsafe operation")
}

func (type2 *safeSliceType) LengthOf(obj any) int {
	return reflect.ValueOf(obj).Elem().Len()
}

func (type2 *safeSliceType) UnsafeLengthOf(ptr unsafe.Pointer) int {
	panic("does not support unsafe operation")
}

func (type2 *safeSliceType) Cap(obj any) int {
	return reflect.ValueOf(obj).Elem().Cap()
}

func (type2 *safeSliceType) UnsafeCap(ptr unsafe.Pointer) int {
	panic("does not support unsafe operation")
}
