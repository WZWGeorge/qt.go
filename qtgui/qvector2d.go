package qtgui

// /usr/include/qt/QtGui/qvector2d.h
// #include <qvector2d.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QVector2D struct {
	*qtrt.CObject
}
type QVector2D_ITF interface {
	QVector2D_PTR() *QVector2D
}

func (ptr *QVector2D) QVector2D_PTR() *QVector2D { return ptr }

func (this *QVector2D) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVector2D) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVector2DFromPointer(cthis unsafe.Pointer) *QVector2D {
	return &QVector2D{&qtrt.CObject{cthis}}
}
func (*QVector2D) NewFromPointer(cthis unsafe.Pointer) *QVector2D {
	return NewQVector2DFromPointer(cthis)
}

// /usr/include/qt/QtGui/qvector2d.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D()

/*
Constructs a null vector, i.e. with coordinates (0, 0).
*/
func (*QVector2D) NewForInherit() *QVector2D {
	return NewQVector2D()
}
func NewQVector2D() *QVector2D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:60
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(Qt::Initialization)

/*
Constructs a null vector, i.e. with coordinates (0, 0).
*/
func (*QVector2D) NewForInherit1(arg0 int) *QVector2D {
	return NewQVector2D1(arg0)
}
func NewQVector2D1(arg0 int) *QVector2D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:61
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(float, float)

/*
Constructs a null vector, i.e. with coordinates (0, 0).
*/
func (*QVector2D) NewForInherit2(xpos float32, ypos float32) *QVector2D {
	return NewQVector2D2(xpos, ypos)
}
func NewQVector2D2(xpos float32, ypos float32) *QVector2D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2Eff", qtrt.FFI_TYPE_POINTER, xpos, ypos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:62
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(const QPoint &)

/*
Constructs a null vector, i.e. with coordinates (0, 0).
*/
func (*QVector2D) NewForInherit3(point qtcore.QPoint_ITF) *QVector2D {
	return NewQVector2D3(point)
}
func NewQVector2D3(point qtcore.QPoint_ITF) *QVector2D {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPoint_PTR() != nil {
		convArg0 = point.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2ERK6QPoint", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:63
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QVector2D(const QPointF &)

/*
Constructs a null vector, i.e. with coordinates (0, 0).
*/
func (*QVector2D) NewForInherit4(point qtcore.QPointF_ITF) *QVector2D {
	return NewQVector2D4(point)
}
func NewQVector2D4(point qtcore.QPointF_ITF) *QVector2D {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2ERK7QPointF", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:65
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QVector2D(const QVector3D &)

/*
Constructs a null vector, i.e. with coordinates (0, 0).
*/
func (*QVector2D) NewForInherit5(vector QVector3D_ITF) *QVector2D {
	return NewQVector2D5(vector)
}
func NewQVector2D5(vector QVector3D_ITF) *QVector2D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector3D_PTR() != nil {
		convArg0 = vector.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2ERK9QVector3D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:68
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QVector2D(const QVector4D &)

/*
Constructs a null vector, i.e. with coordinates (0, 0).
*/
func (*QVector2D) NewForInherit6(vector QVector4D_ITF) *QVector2D {
	return NewQVector2D6(vector)
}
func NewQVector2D6(vector QVector4D_ITF) *QVector2D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector4D_PTR() != nil {
		convArg0 = vector.QVector4D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DC2ERK9QVector4D", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVector2D)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the x and y coordinates are set to 0.0, otherwise returns false.
*/
func (this *QVector2D) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qvector2d.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float x() const

/*
Returns the x coordinate of this point.

See also setX() and y().
*/
func (this *QVector2D) X() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [4] float y() const

/*
Returns the y coordinate of this point.

See also setY() and x().
*/
func (this *QVector2D) Y() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(float)

/*
Sets the x coordinate of this point to the given x coordinate.

See also x() and setY().
*/
func (this *QVector2D) SetX(x float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2D4setXEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(float)

/*
Sets the y coordinate of this point to the given y coordinate.

See also y() and setX().
*/
func (this *QVector2D) SetY(y float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2D4setYEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] float & operator[](int)

/*

 */
func (this *QVector2D) Operator_get_index(i int) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cpretval2go("float32", rv).(float32) // 3331
}

// /usr/include/qt/QtGui/qvector2d.h:80
// index:1
// Public Visibility=Default Availability=Available
// [4] float operator[](int) const

/*

 */
func (this *QVector2D) Operator_get_index1(i int) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2DixEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] float length() const

/*
Returns the length of the vector from the origin.

See also lengthSquared() and normalized().
*/
func (this *QVector2D) Length() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D6lengthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] float lengthSquared() const

/*
Returns the squared length of the vector from the origin. This is equivalent to the dot product of the vector with itself.

See also length() and dotProduct().
*/
func (this *QVector2D) LengthSquared() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D13lengthSquaredEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D normalized() const

/*
Returns the normalized unit vector form of this vector.

If this vector is null, then a null vector is returned. If the length of the vector is very close to 1, then the vector will be returned as-is. Otherwise the normalized form of the vector of length 1 will be returned.

See also length() and normalize().
*/
func (this *QVector2D) Normalized() *QVector2D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize()

/*
Normalizes the currect vector in place. Nothing happens if this vector is a null vector or the length of the vector is very close to 1.

See also length() and normalized().
*/
func (this *QVector2D) Normalize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2D9normalizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToPoint(const QVector2D &) const

/*
Returns the distance from this vertex to a point defined by the vertex point.

This function was introduced in  Qt 5.1.

See also distanceToLine().
*/
func (this *QVector2D) DistanceToPoint(point QVector2D_ITF) float32 {
	var convArg0 unsafe.Pointer
	if point != nil && point.QVector2D_PTR() != nil {
		convArg0 = point.QVector2D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D15distanceToPointERKS_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] float distanceToLine(const QVector2D &, const QVector2D &) const

/*
Returns the distance that this vertex is from a line defined by point and the unit vector direction.

If direction is a null vector, then it does not define a line. In that case, the distance from point to this vertex is returned.

This function was introduced in  Qt 5.1.

See also distanceToPoint().
*/
func (this *QVector2D) DistanceToLine(point QVector2D_ITF, direction QVector2D_ITF) float32 {
	var convArg0 unsafe.Pointer
	if point != nil && point.QVector2D_PTR() != nil {
		convArg0 = point.QVector2D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if direction != nil && direction.QVector2D_PTR() != nil {
		convArg1 = direction.QVector2D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D14distanceToLineERKS_S1_", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qvector2d.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D & operator+=(const QVector2D &)

/*

 */
func (this *QVector2D) Operator_add_equal(vector QVector2D_ITF) *QVector2D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector2D_PTR() != nil {
		convArg0 = vector.QVector2D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DpLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D & operator-=(const QVector2D &)

/*

 */
func (this *QVector2D) Operator_minus_equal(vector QVector2D_ITF) *QVector2D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector2D_PTR() != nil {
		convArg0 = vector.QVector2D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DmIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D & operator*=(float)

/*

 */
func (this *QVector2D) Operator_mul_equal(factor float32) *QVector2D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DmLEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:94
// index:1
// Public Visibility=Default Availability=Available
// [8] QVector2D & operator*=(const QVector2D &)

/*

 */
func (this *QVector2D) Operator_mul_equal1(vector QVector2D_ITF) *QVector2D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector2D_PTR() != nil {
		convArg0 = vector.QVector2D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DmLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QVector2D & operator/=(float)

/*

 */
func (this *QVector2D) Operator_div_equal(divisor float32) *QVector2D {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DdVEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), divisor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:96
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QVector2D & operator/=(const QVector2D &)

/*

 */
func (this *QVector2D) Operator_div_equal1(vector QVector2D_ITF) *QVector2D {
	var convArg0 unsafe.Pointer
	if vector != nil && vector.QVector2D_PTR() != nil {
		convArg0 = vector.QVector2D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DdVERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector2DFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector2D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [4] float dotProduct(const QVector2D &, const QVector2D &)

/*
Returns the dot product of v1 and v2.
*/
func (this *QVector2D) DotProduct(v1 QVector2D_ITF, v2 QVector2D_ITF) float32 {
	var convArg0 unsafe.Pointer
	if v1 != nil && v1.QVector2D_PTR() != nil {
		convArg0 = v1.QVector2D_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if v2 != nil && v2.QVector2D_PTR() != nil {
		convArg1 = v2.QVector2D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2D10dotProductERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}
func QVector2D_DotProduct(v1 QVector2D_ITF, v2 QVector2D_ITF) float32 {
	var nilthis *QVector2D
	rv := nilthis.DotProduct(v1, v2)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:114
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D toVector3D() const

/*
Returns the 3D form of this 2D vector, with the z coordinate set to zero.

See also toVector4D() and toPoint().
*/
func (this *QVector2D) ToVector3D() *QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D10toVector3DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:117
// index:0
// Public Visibility=Default Availability=Available
// [16] QVector4D toVector4D() const

/*
Returns the 4D form of this 2D vector, with the z and w coordinates set to zero.

See also toVector3D() and toPoint().
*/
func (this *QVector2D) ToVector4D() *QVector4D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D10toVector4DEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVector4DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVector4D)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint toPoint() const

/*
Returns the QPoint form of this 2D vector.

See also toPointF() and toVector3D().
*/
func (this *QVector2D) ToPoint() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D7toPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qvector2d.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QPointF toPointF() const

/*
Returns the QPointF form of this 2D vector.

See also toPoint() and toVector3D().
*/
func (this *QVector2D) ToPointF() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QVector2D8toPointFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

func DeleteQVector2D(this *QVector2D) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QVector2DD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10639() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
