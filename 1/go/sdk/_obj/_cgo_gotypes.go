// Code generated by cmd/cgo; DO NOT EDIT.

package sdk

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_People = _Ctype_struct_People

type _Ctype__GoBytes_ []byte

type _Ctype__GoString_ string

type _Ctype_char int8

type _Ctype_int int32

type _Ctype_intgo = _Ctype_ptrdiff_t

type _Ctype_long int64

type _Ctype_ptrdiff_t = _Ctype_long

type _Ctype_struct_People struct {
	name	*_Ctype_char
	content	*_Ctype_uchar
	payload	*_Ctype_char
	age	_Ctype_int
	_	[4]byte
}

type _Ctype_uchar uint8

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})


func _Cfunc_CBytes(b []byte) unsafe.Pointer {
	p := _cgo_cmalloc(uint64(len(b)))
	pp := (*[1<<30]byte)(p)
	copy(pp[:], b)
	return p
}

func _Cfunc_CString(s string) *_Ctype_char {
	p := _cgo_cmalloc(uint64(len(s)+1))
	pp := (*[1<<30]byte)(p)
	copy(pp[:], s)
	pp[len(s)] = 0
	return (*_Ctype_char)(p)
}

//go:linkname _cgo_runtime_gobytes runtime.gobytes
func _cgo_runtime_gobytes(unsafe.Pointer, int) []byte

func _Cfunc_GoBytes(p unsafe.Pointer, l _Ctype_int) []byte {
	return _cgo_runtime_gobytes(p, int(l))
}

//go:linkname _cgo_runtime_gostring runtime.gostring
func _cgo_runtime_gostring(*_Ctype_char) string

func _Cfunc_GoString(p *_Ctype_char) string {
	return _cgo_runtime_gostring(p)
}
//go:cgo_import_static _cgo_dfc451a857ba_Cfunc_free_people
//go:linkname __cgofn__cgo_dfc451a857ba_Cfunc_free_people _cgo_dfc451a857ba_Cfunc_free_people
var __cgofn__cgo_dfc451a857ba_Cfunc_free_people byte
var _cgo_dfc451a857ba_Cfunc_free_people = unsafe.Pointer(&__cgofn__cgo_dfc451a857ba_Cfunc_free_people)

//go:cgo_unsafe_args
func _Cfunc_free_people(p0 *_Ctype_struct_People) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_dfc451a857ba_Cfunc_free_people, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_dfc451a857ba_Cfunc_len_of_uchar
//go:linkname __cgofn__cgo_dfc451a857ba_Cfunc_len_of_uchar _cgo_dfc451a857ba_Cfunc_len_of_uchar
var __cgofn__cgo_dfc451a857ba_Cfunc_len_of_uchar byte
var _cgo_dfc451a857ba_Cfunc_len_of_uchar = unsafe.Pointer(&__cgofn__cgo_dfc451a857ba_Cfunc_len_of_uchar)

//go:cgo_unsafe_args
func _Cfunc_len_of_uchar(p0 *_Ctype_uchar) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_dfc451a857ba_Cfunc_len_of_uchar, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_dfc451a857ba_Cfunc_new_people
//go:linkname __cgofn__cgo_dfc451a857ba_Cfunc_new_people _cgo_dfc451a857ba_Cfunc_new_people
var __cgofn__cgo_dfc451a857ba_Cfunc_new_people byte
var _cgo_dfc451a857ba_Cfunc_new_people = unsafe.Pointer(&__cgofn__cgo_dfc451a857ba_Cfunc_new_people)

//go:cgo_unsafe_args
func _Cfunc_new_people() (r1 *_Ctype_struct_People) {
	_cgo_runtime_cgocall(_cgo_dfc451a857ba_Cfunc_new_people, uintptr(unsafe.Pointer(&r1)))
	if _Cgo_always_false {
	}
	return
}
//go:cgo_export_dynamic Echo
//go:linkname _cgoexp_dfc451a857ba_Echo _cgoexp_dfc451a857ba_Echo
//go:cgo_export_static _cgoexp_dfc451a857ba_Echo
//go:nosplit
//go:norace
func _cgoexp_dfc451a857ba_Echo(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_dfc451a857ba_Echo
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_dfc451a857ba_Echo(p0 *_Ctype_struct_People) (r0 *_Ctype_struct_People) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return Echo(p0)
}

//go:cgo_import_static _cgo_dfc451a857ba_Cfunc__Cmalloc
//go:linkname __cgofn__cgo_dfc451a857ba_Cfunc__Cmalloc _cgo_dfc451a857ba_Cfunc__Cmalloc
var __cgofn__cgo_dfc451a857ba_Cfunc__Cmalloc byte
var _cgo_dfc451a857ba_Cfunc__Cmalloc = unsafe.Pointer(&__cgofn__cgo_dfc451a857ba_Cfunc__Cmalloc)

//go:linkname runtime_throw runtime.throw
func runtime_throw(string)

//go:cgo_unsafe_args
func _cgo_cmalloc(p0 uint64) (r1 unsafe.Pointer) {
	_cgo_runtime_cgocall(_cgo_dfc451a857ba_Cfunc__Cmalloc, uintptr(unsafe.Pointer(&p0)))
	if r1 == nil {
		runtime_throw("runtime: C malloc failed")
	}
	return
}
