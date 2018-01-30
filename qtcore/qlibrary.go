package qtcore

// /usr/include/qt/QtCore/qlibrary.h
// #include <qlibrary.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
type QLibrary struct {
	*QObject
}

func (this *QLibrary) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QLibrary) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQLibraryFromPointer(cthis unsafe.Pointer) *QLibrary {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QLibrary{bcthis0}
}
func (*QLibrary) NewFromPointer(cthis unsafe.Pointer) *QLibrary {
	return NewQLibraryFromPointer(cthis)
}

// /usr/include/qt/QtCore/qlibrary.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QLibrary) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qlibrary.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(QObject *)
func NewQLibrary(parent *QObject /*777 QObject **/) *QLibrary {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:69
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(const QString &, QObject *)
func NewQLibrary_1(fileName *QString, parent *QObject /*777 QObject **/) *QLibrary {
	var convArg0 = fileName.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:70
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(const QString &, int, QObject *)
func NewQLibrary_2(fileName *QString, verNum int, parent *QObject /*777 QObject **/) *QLibrary {
	var convArg0 = fileName.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringiP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0, verNum, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:71
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(const QString &, const QString &, QObject *)
func NewQLibrary_3(fileName *QString, version *QString, parent *QObject /*777 QObject **/) *QLibrary {
	var convArg0 = fileName.GetCthis()
	var convArg1 = version.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringS2_P7QObject", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLibrary()
func DeleteQLibrary(*QLibrary) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QFunctionPointer resolve(const char *)
func (this *QLibrary) Resolve(symbol string) unsafe.Pointer /*666*/ {
	var convArg0 = qtrt.CString(symbol)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlibrary.h:75
// index:1
// Public static Visibility=Default Availability=Available
// [8] QFunctionPointer resolve(const QString &, const char *)
func (this *QLibrary) Resolve_1(fileName *QString, symbol string) unsafe.Pointer /*666*/ {
	var convArg0 = fileName.GetCthis()
	var convArg1 = qtrt.CString(symbol)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringPKc", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}
func QLibrary_Resolve_1(fileName *QString, symbol string) unsafe.Pointer /*666*/ {
	var nilthis *QLibrary
	rv := nilthis.Resolve_1(fileName, symbol)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:76
// index:2
// Public static Visibility=Default Availability=Available
// [8] QFunctionPointer resolve(const QString &, int, const char *)
func (this *QLibrary) Resolve_2(fileName *QString, verNum int, symbol string) unsafe.Pointer /*666*/ {
	var convArg0 = fileName.GetCthis()
	var convArg2 = qtrt.CString(symbol)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringiPKc", ffiqt.FFI_TYPE_POINTER, convArg0, verNum, convArg2)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}
func QLibrary_Resolve_2(fileName *QString, verNum int, symbol string) unsafe.Pointer /*666*/ {
	var nilthis *QLibrary
	rv := nilthis.Resolve_2(fileName, verNum, symbol)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:77
// index:3
// Public static Visibility=Default Availability=Available
// [8] QFunctionPointer resolve(const QString &, const QString &, const char *)
func (this *QLibrary) Resolve_3(fileName *QString, version *QString, symbol string) unsafe.Pointer /*666*/ {
	var convArg0 = fileName.GetCthis()
	var convArg1 = version.GetCthis()
	var convArg2 = qtrt.CString(symbol)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringS2_PKc", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}
func QLibrary_Resolve_3(fileName *QString, version *QString, symbol string) unsafe.Pointer /*666*/ {
	var nilthis *QLibrary
	rv := nilthis.Resolve_3(fileName, version, symbol)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load()
func (this *QLibrary) Load() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary4loadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qlibrary.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool unload()
func (this *QLibrary) Unload() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary6unloadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qlibrary.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoaded()
func (this *QLibrary) IsLoaded() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary8isLoadedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qlibrary.h:83
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isLibrary(const QString &)
func (this *QLibrary) IsLibrary(fileName *QString) bool {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary9isLibraryERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QLibrary_IsLibrary(fileName *QString) bool {
	var nilthis *QLibrary
	rv := nilthis.IsLibrary(fileName)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QLibrary) SetFileName(fileName *QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QLibrary) FileName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlibrary.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileNameAndVersion(const QString &, int)
func (this *QLibrary) SetFileNameAndVersion(fileName *QString, verNum int) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary21setFileNameAndVersionERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, verNum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFileNameAndVersion(const QString &, const QString &)
func (this *QLibrary) SetFileNameAndVersion_1(fileName *QString, version *QString) {
	var convArg0 = fileName.GetCthis()
	var convArg1 = version.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary21setFileNameAndVersionERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QLibrary) ErrorString() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qlibrary.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QLibrary::LoadHints loadHints()
func (this *QLibrary) LoadHints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary9loadHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

type QLibrary__LoadHint = int

const QLibrary__ResolveAllSymbolsHint QLibrary__LoadHint = 1
const QLibrary__ExportExternalSymbolsHint QLibrary__LoadHint = 2
const QLibrary__LoadArchiveMemberHint QLibrary__LoadHint = 4
const QLibrary__PreventUnloadHint QLibrary__LoadHint = 8
const QLibrary__DeepBindHint QLibrary__LoadHint = 16

//  body block end