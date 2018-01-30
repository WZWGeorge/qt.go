package qtwidgets

// /usr/include/qt/QtWidgets/qtreeview.h
// #include <qtreeview.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTreeView struct {
	*QAbstractItemView
}

func (this *QTreeView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemView.GetCthis()
	}
}
func (this *QTreeView) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemView = NewQAbstractItemViewFromPointer(cthis)
}
func NewQTreeViewFromPointer(cthis unsafe.Pointer) *QTreeView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QTreeView{bcthis0}
}
func (*QTreeView) NewFromPointer(cthis unsafe.Pointer) *QTreeView {
	return NewQTreeViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtreeview.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTreeView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeView(QWidget *)
func NewQTreeView(parent *QWidget /*777 QWidget **/) *QTreeView {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeViewC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtreeview.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTreeView()
func DeleteQTreeView(*QTreeView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)
func (this *QTreeView) SetModel(model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = model.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRootIndex(const QModelIndex &)
func (this *QTreeView) SetRootIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)
func (this *QTreeView) SetSelectionModel(selectionModel *qtcore.QItemSelectionModel /*777 QItemSelectionModel **/) {
	var convArg0 = selectionModel.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QHeaderView * header()
func (this *QTreeView) Header() *QHeaderView /*777 QHeaderView **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView6headerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeader(QHeaderView *)
func (this *QTreeView) SetHeader(header *QHeaderView /*777 QHeaderView **/) {
	var convArg0 = header.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9setHeaderEP11QHeaderView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoExpandDelay()
func (this *QTreeView) AutoExpandDelay() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15autoExpandDelayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoExpandDelay(int)
func (this *QTreeView) SetAutoExpandDelay(delay int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setAutoExpandDelayEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), delay)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int indentation()
func (this *QTreeView) Indentation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11indentationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndentation(int)
func (this *QTreeView) SetIndentation(i int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14setIndentationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetIndentation()
func (this *QTreeView) ResetIndentation() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16resetIndentationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rootIsDecorated()
func (this *QTreeView) RootIsDecorated() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15rootIsDecoratedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRootIsDecorated(_Bool)
func (this *QTreeView) SetRootIsDecorated(show bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setRootIsDecoratedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool uniformRowHeights()
func (this *QTreeView) UniformRowHeights() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView17uniformRowHeightsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUniformRowHeights(_Bool)
func (this *QTreeView) SetUniformRowHeights(uniform bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView20setUniformRowHeightsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), uniform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool itemsExpandable()
func (this *QTreeView) ItemsExpandable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15itemsExpandableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemsExpandable(_Bool)
func (this *QTreeView) SetItemsExpandable(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setItemsExpandableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool expandsOnDoubleClick()
func (this *QTreeView) ExpandsOnDoubleClick() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView20expandsOnDoubleClickEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpandsOnDoubleClick(_Bool)
func (this *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView23setExpandsOnDoubleClickEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnViewportPosition(int)
func (this *QTreeView) ColumnViewportPosition(column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView22columnViewportPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnWidth(int)
func (this *QTreeView) ColumnWidth(column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11columnWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnWidth(int, int)
func (this *QTreeView) SetColumnWidth(column int, width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14setColumnWidthEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnAt(int)
func (this *QTreeView) ColumnAt(x int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8columnAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isColumnHidden(int)
func (this *QTreeView) IsColumnHidden(column int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14isColumnHiddenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnHidden(int, _Bool)
func (this *QTreeView) SetColumnHidden(column int, hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setColumnHiddenEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHeaderHidden()
func (this *QTreeView) IsHeaderHidden() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14isHeaderHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeaderHidden(_Bool)
func (this *QTreeView) SetHeaderHidden(hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setHeaderHiddenEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRowHidden(int, const QModelIndex &)
func (this *QTreeView) IsRowHidden(row int, parent *qtcore.QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowHidden(int, const QModelIndex &, _Bool)
func (this *QTreeView) SetRowHidden(row int, parent *qtcore.QModelIndex, hide bool) {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFirstColumnSpanned(int, const QModelIndex &)
func (this *QTreeView) IsFirstColumnSpanned(row int, parent *qtcore.QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFirstColumnSpanned(int, const QModelIndex &, _Bool)
func (this *QTreeView) SetFirstColumnSpanned(row int, parent *qtcore.QModelIndex, span bool) {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, span)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExpanded(const QModelIndex &)
func (this *QTreeView) IsExpanded(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10isExpandedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpanded(const QModelIndex &, _Bool)
func (this *QTreeView) SetExpanded(index *qtcore.QModelIndex, expand bool) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setExpandedERK11QModelIndexb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, expand)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortingEnabled(_Bool)
func (this *QTreeView) SetSortingEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17setSortingEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortingEnabled()
func (this *QTreeView) IsSortingEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16isSortingEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnimated(_Bool)
func (this *QTreeView) SetAnimated(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setAnimatedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimated()
func (this *QTreeView) IsAnimated() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10isAnimatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAllColumnsShowFocus(_Bool)
func (this *QTreeView) SetAllColumnsShowFocus(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView22setAllColumnsShowFocusEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool allColumnsShowFocus()
func (this *QTreeView) AllColumnsShowFocus() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView19allColumnsShowFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrap(_Bool)
func (this *QTreeView) SetWordWrap(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setWordWrapEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wordWrap()
func (this *QTreeView) WordWrap() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8wordWrapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTreePosition(int)
func (this *QTreeView) SetTreePosition(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setTreePositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:133
// index:0
// Public Visibility=Default Availability=Available
// [4] int treePosition()
func (this *QTreeView) TreePosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView12treePositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:135
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void keyboardSearch(const QString &)
func (this *QTreeView) KeyboardSearch(search *qtcore.QString) {
	var convArg0 = search.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14keyboardSearchERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:137
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &)
func (this *QTreeView) VisualRect(index *qtcore.QModelIndex) *qtcore.QRect /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QTreeView) ScrollTo(index *qtcore.QModelIndex, hint int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &)
func (this *QTreeView) IndexAt(p *qtcore.QPoint) *qtcore.QModelIndex /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView7indexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:140
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex indexAbove(const QModelIndex &)
func (this *QTreeView) IndexAbove(index *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10indexAboveERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:141
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex indexBelow(const QModelIndex &)
func (this *QTreeView) IndexBelow(index *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10indexBelowERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:143
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()
func (this *QTreeView) DoItemsLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13doItemsLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset()
func (this *QTreeView) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortByColumn(int, Qt::SortOrder)
func (this *QTreeView) SortByColumn(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void sortByColumn(int)
func (this *QTreeView) SortByColumn_1(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:149
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void selectAll()
func (this *QTreeView) SelectAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9selectAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expanded(const QModelIndex &)
func (this *QTreeView) Expanded(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8expandedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapsed(const QModelIndex &)
func (this *QTreeView) Collapsed(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9collapsedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hideColumn(int)
func (this *QTreeView) HideColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10hideColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showColumn(int)
func (this *QTreeView) ShowColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10showColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expand(const QModelIndex &)
func (this *QTreeView) Expand(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView6expandERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapse(const QModelIndex &)
func (this *QTreeView) Collapse(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8collapseERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeColumnToContents(int)
func (this *QTreeView) ResizeColumnToContents(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView22resizeColumnToContentsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expandAll()
func (this *QTreeView) ExpandAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9expandAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapseAll()
func (this *QTreeView) CollapseAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11collapseAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expandToDepth(int)
func (this *QTreeView) ExpandToDepth(depth int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13expandToDepthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:167
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnResized(int, int, int)
func (this *QTreeView) ColumnResized(column int, oldSize int, newSize int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13columnResizedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, oldSize, newSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:168
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnCountChanged(int, int)
func (this *QTreeView) ColumnCountChanged(oldCount int, newCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18columnCountChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), oldCount, newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:169
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnMoved()
func (this *QTreeView) ColumnMoved() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11columnMovedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:170
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void reexpand()
func (this *QTreeView) Reexpand() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8reexpandEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:171
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void rowsRemoved(const QModelIndex &, int, int)
func (this *QTreeView) RowsRemoved(parent *qtcore.QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11rowsRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:175
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)
func (this *QTreeView) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:176
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsInserted(const QModelIndex &, int, int)
func (this *QTreeView) RowsInserted(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:177
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsAboutToBeRemoved(const QModelIndex &, int, int)
func (this *QTreeView) RowsAboutToBeRemoved(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:180
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int horizontalOffset()
func (this *QTreeView) HorizontalOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16horizontalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:181
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int verticalOffset()
func (this *QTreeView) VerticalOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14verticalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)
func (this *QTreeView) SetSelection(rect *qtcore.QRect, command int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:184
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &)
func (this *QTreeView) VisualRegionForSelection(selection *qtcore.QItemSelection) *qtgui.QRegion /*123*/ {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:187
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QTreeView) TimerEvent(event *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QTreeView) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:190
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawTree(QPainter *, const QRegion &)
func (this *QTreeView) DrawTree(painter *qtgui.QPainter /*777 QPainter **/, region *qtgui.QRegion) {
	var convArg0 = painter.GetCthis()
	var convArg1 = region.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawRow(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QTreeView) DrawRow(painter *qtgui.QPainter /*777 QPainter **/, options *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg0 = painter.GetCthis()
	var convArg1 = options.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:194
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawBranches(QPainter *, const QRect &, const QModelIndex &)
func (this *QTreeView) DrawBranches(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, index *qtcore.QModelIndex) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:198
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QTreeView) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:199
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QTreeView) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:200
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QTreeView) MouseDoubleClickEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:201
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QTreeView) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:202
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QTreeView) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:204
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)
func (this *QTreeView) DragMoveEvent(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:206
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)
func (this *QTreeView) ViewportEvent(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()
func (this *QTreeView) UpdateGeometries() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16updateGeometriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:210
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint()
func (this *QTreeView) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16viewportSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int sizeHintForColumn(int)
func (this *QTreeView) SizeHintForColumn(column int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView17sizeHintForColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:213
// index:0
// Protected Visibility=Default Availability=Available
// [4] int indexRowSizeHint(const QModelIndex &)
func (this *QTreeView) IndexRowSizeHint(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:214
// index:0
// Protected Visibility=Default Availability=Available
// [4] int rowHeight(const QModelIndex &)
func (this *QTreeView) RowHeight(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView9rowHeightERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:216
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void horizontalScrollbarAction(int)
func (this *QTreeView) HorizontalScrollbarAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView25horizontalScrollbarActionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &)
func (this *QTreeView) IsIndexHidden(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void selectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QTreeView) SelectionChanged(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) {
	var convArg0 = selected.GetCthis()
	var convArg1 = deselected.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QTreeView) CurrentChanged(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

//  body block end