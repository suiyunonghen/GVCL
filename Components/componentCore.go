package Components

import (
	"github.com/suiyunonghen/GVCL/Graphics"
	"syscall"
	"github.com/suiyunonghen/GVCL/WinApi"
)

type IComponent interface {
	GetName() string       //组件名称
	SetName(fName string)  //设置组件名称
	GetTagetData() uintptr //附加数据
	SetTagetData(fData uintptr) uintptr
	GetTag() int //附加数据
	SetTag(ftag int)
	Free() //释放数据
}

type IControl interface {
	GetColor() Graphics.GColorValue
	SetColor(c Graphics.GColorValue)
	Invalidate() //触发刷新
	Left() int32
	SetLeft(l int32)
	Top() int32
	SetTop(t int32)
	Width() int32
	SetWidth(w int32)
	Height() int32
	SetHeight(h int32)
	GetParent() IWincontrol
	Free() //释放数据
	GetDeviceContext()(WinApi.HDC,syscall.Handle)
	Visible() bool
	PaintControl(dc WinApi.HDC) //执行绘制
	Paint(cvs Graphics.ICanvas)
	SetBounds(ALeft, ATop, AWidth, AHeight int32)
	IsWindowControl()bool
	AfterParentWndCreate()
}

type IApplication interface {
	Run()
	Active() bool
}

type GCreateParams struct {
	Caption      string
	Style        uint32
	ExStyle      uint32
	X            int32
	Y            int32
	Width        int32
	Height       int32
	WndParent    syscall.Handle
	Param        uintptr
	WindowClass  WinApi.GWndClass
	WinClassName string
}

type IWincontrol interface {
	GetWindowHandle() syscall.Handle
	CreateWnd()
	DestoryWnd()
	Focused() bool
	ControlCount() int
	WControlCount() int
	InsertChildWinCtrl(ctrl IWincontrol)
	InsertControl(ctrl IControl)
	GetParent() IWincontrol
	RemoveChildWinCtrl(ctrl IWincontrol)
	RemoveControl(ctrl IControl)
	ControlExists(ctrl IControl) bool
	WindowExists(handle syscall.Handle) bool
	UpdateShowing()
	WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool)
	CreateParams(params *GCreateParams)
	PaintWindow(dc WinApi.HDC)int32
	CreateWindowHandle(params *GCreateParams)bool
	PaintBack(dc WinApi.HDC)int32
	HandleAllocated()bool
}


/**
**基本组件
**/
type GComponent struct {
	Graphics.GObject
	fname      string `json:"Name"`
	ftag       int    `json:"Tag"`
	fTagetData uintptr
}

func (cmp *GComponent) GetName() string {
	return cmp.fname
}

func (cmp *GComponent) SetName(fName string) {
	cmp.fname = fName
}

func (cmp *GComponent) GetTagetData() uintptr {
	return cmp.fTagetData
}

func (cmp *GComponent) SetTagetData(fdata uintptr) {
	cmp.fTagetData = fdata
}

func (cmp *GComponent) GetTag() int {
	return cmp.ftag
}

func (cmp *GComponent) SetTag(fdata int) {
	cmp.ftag = fdata
}

