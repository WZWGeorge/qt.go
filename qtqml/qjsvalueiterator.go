package qtqml

// /usr/include/qt/QtQml/qjsvalueiterator.h
// #include <qjsvalueiterator.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QJSValueIterator struct {
	*qtrt.CObject
}

func (this *QJSValueIterator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJSValueIterator) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQJSValueIteratorFromPointer(cthis unsafe.Pointer) *QJSValueIterator {
	return &QJSValueIterator{&qtrt.CObject{cthis}}
}
func (*QJSValueIterator) NewFromPointer(cthis unsafe.Pointer) *QJSValueIterator {
	return NewQJSValueIteratorFromPointer(cthis)
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJSValueIterator(const QJSValue &)
func NewQJSValueIterator(value *QJSValue) *QJSValueIterator {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QJSValueIteratorC2ERK8QJSValue", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJSValueIterator()
func DeleteQJSValueIterator(*QJSValueIterator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QJSValueIteratorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNext()
func (this *QJSValueIterator) HasNext() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QJSValueIterator7hasNextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:60
// index:0
// Public Visibility=Default Availability=Available
// [1] bool next()
func (this *QJSValueIterator) Next() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QJSValueIterator4nextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QJSValueIterator) Name() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QJSValueIterator4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue value()
func (this *QJSValueIterator) Value() *QJSValue /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QJSValueIterator5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end