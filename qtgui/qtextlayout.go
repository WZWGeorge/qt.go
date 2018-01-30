package qtgui

// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QTextLayout struct {
	*qtrt.CObject
}

func (this *QTextLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextLayout) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTextLayoutFromPointer(cthis unsafe.Pointer) *QTextLayout {
	return &QTextLayout{&qtrt.CObject{cthis}}
}
func (*QTextLayout) NewFromPointer(cthis unsafe.Pointer) *QTextLayout {
	return NewQTextLayoutFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextlayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout()
func NewQTextLayout() *QTextLayout {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:109
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QString &)
func NewQTextLayout_1(text *qtcore.QString) *QTextLayout {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:110
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QString &, const QFont &, QPaintDevice *)
func NewQTextLayout_2(text *qtcore.QString, font *QFont, paintdevice *QPaintDevice /*777 QPaintDevice **/) *QTextLayout {
	var convArg0 = text.GetCthis()
	var convArg1 = font.GetCthis()
	var convArg2 = paintdevice.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:111
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QTextBlock &)
func NewQTextLayout_3(b *QTextBlock) *QTextLayout {
	var convArg0 = b.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextLayout()
func DeleteQTextLayout(*QTextLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QTextLayout) SetFont(f *QFont) {
	var convArg0 = f.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:115
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
func (this *QTextLayout) Font() *QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawFont(const QRawFont &)
func (this *QTextLayout) SetRawFont(rawFont *QRawFont) {
	var convArg0 = rawFont.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout10setRawFontERK8QRawFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QTextLayout) SetText(string *qtcore.QString) {
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QTextLayout) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextOption(const QTextOption &)
func (this *QTextLayout) SetTextOption(option *QTextOption) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout13setTextOptionERK11QTextOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] const QTextOption & textOption()
func (this *QTextLayout) TextOption() *QTextOption {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout10textOptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreeditArea(int, const QString &)
func (this *QTextLayout) SetPreeditArea(position int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout14setPreeditAreaEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), position, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] int preeditAreaPosition()
func (this *QTextLayout) PreeditAreaPosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout19preeditAreaPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QString preeditAreaText()
func (this *QTextLayout) PreeditAreaText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout15preeditAreaTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearAdditionalFormats()
func (this *QTextLayout) ClearAdditionalFormats() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout22clearAdditionalFormatsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearFormats()
func (this *QTextLayout) ClearFormats() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout12clearFormatsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheEnabled(_Bool)
func (this *QTextLayout) SetCacheEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout15setCacheEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cacheEnabled()
func (this *QTextLayout) CacheEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12cacheEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QTextLayout) SetCursorMoveStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout18setCursorMoveStyleEN2Qt15CursorMoveStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:154
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CursorMoveStyle cursorMoveStyle()
func (this *QTextLayout) CursorMoveStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout15cursorMoveStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginLayout()
func (this *QTextLayout) BeginLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout11beginLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endLayout()
func (this *QTextLayout) EndLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout9endLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearLayout()
func (this *QTextLayout) ClearLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout11clearLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:160
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLine createLine()
func (this *QTextLayout) CreateLine() *QTextLine /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout10createLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:162
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineCount()
func (this *QTextLayout) LineCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout9lineCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:163
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLine lineAt(int)
func (this *QTextLayout) LineAt(i int) *QTextLine /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout6lineAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:164
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLine lineForTextPosition(int)
func (this *QTextLayout) LineForTextPosition(pos int) *QTextLine /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout19lineForTextPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValidCursorPosition(int)
func (this *QTextLayout) IsValidCursorPosition(pos int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout21isValidCursorPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:171
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextCursorPosition(int, enum QTextLayout::CursorMode)
func (this *QTextLayout) NextCursorPosition(oldPos int, mode int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout18nextCursorPositionEiNS_10CursorModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), oldPos, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:172
// index:0
// Public Visibility=Default Availability=Available
// [4] int previousCursorPosition(int, enum QTextLayout::CursorMode)
func (this *QTextLayout) PreviousCursorPosition(oldPos int, mode int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout22previousCursorPositionEiNS_10CursorModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), oldPos, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:173
// index:0
// Public Visibility=Default Availability=Available
// [4] int leftCursorPosition(int)
func (this *QTextLayout) LeftCursorPosition(oldPos int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout18leftCursorPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), oldPos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:174
// index:0
// Public Visibility=Default Availability=Available
// [4] int rightCursorPosition(int)
func (this *QTextLayout) RightCursorPosition(oldPos int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout19rightCursorPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), oldPos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawCursor(QPainter *, const QPointF &, int)
func (this *QTextLayout) DrawCursor(p *QPainter /*777 QPainter **/, pos *qtcore.QPointF, cursorPosition int) {
	var convArg0 = p.GetCthis()
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cursorPosition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:179
// index:1
// Public Visibility=Default Availability=Available
// [-2] void drawCursor(QPainter *, const QPointF &, int, int)
func (this *QTextLayout) DrawCursor_1(p *QPainter /*777 QPainter **/, pos *qtcore.QPointF, cursorPosition int, width int) {
	var convArg0 = p.GetCthis()
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cursorPosition, width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:181
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position()
func (this *QTextLayout) Position() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)
func (this *QTextLayout) SetPosition(p *qtcore.QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout11setPositionERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:184
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QTextLayout) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:186
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minimumWidth()
func (this *QTextLayout) MinimumWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12minimumWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maximumWidth()
func (this *QTextLayout) MaximumWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12maximumWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(int)
func (this *QTextLayout) SetFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout8setFlagsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

type QTextLayout__CursorMode = int

const QTextLayout__SkipCharacters QTextLayout__CursorMode = 0
const QTextLayout__SkipWords QTextLayout__CursorMode = 1

//  body block end