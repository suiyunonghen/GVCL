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
	ClientRect()WinApi.Rect
	Enabled()bool
	SetEnabled(v bool)
	MouseEnter()
	MouseLeave()
	MouseMove(x,y int,state KeyState)
	MouseDown(button MouseButton,x,y int,state KeyState)
	MouseUp(button MouseButton,x,y int,state KeyState)
}

type MouseButton uint8

const (
	MbLeft MouseButton = iota
	MbRight
	MbMiddle
)

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
	IControl
	GetWindowHandle() syscall.Handle
	CreateWnd()
	DestoryWnd()
	Focused() bool
	ControlCount() int
	WControlCount() int
	InsertChildWinCtrl(ctrl IWincontrol)
	InsertControl(ctrl IControl)
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
	Perform(msg uint32,wparam, lparam uintptr)(result WinApi.LRESULT)
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

const(
	MK_XBUTTON1 = 0x0020
	MK_XBUTTON2 = 0x0040
)
//键盘状态
type KeyState uint32

func (state KeyState)CtrlKeyDown()bool  {
	return state & WinApi.MK_CONTROL != 0
}

func (state KeyState)LButtonDown()bool  {
	return state & WinApi.MK_LBUTTON != 0
}

func (state KeyState)MButtonDown()bool  {
	return state & WinApi.MK_MBUTTON != 0
}

func (state KeyState)RButtonDown()bool  {

	return state & WinApi.MK_RBUTTON != 0
}

func (state KeyState)ShiftKeyDown()bool  {
	return state & WinApi.MK_SHIFT != 0
}

func (state KeyState)AltKeyDown()bool  {
	return WinApi.GetKeyState(WinApi.VK_MENU) < 0
}

func (state KeyState)XBUTTON1Down()bool  {
	return MK_XBUTTON1 & state != 0
}

func (state KeyState)XBUTTON2Down()bool  {
	return MK_XBUTTON2 & state != 0
}

func (state *KeyState)SetCtrlKeyDown(v bool)  {
	*state = KeyState(uint32(*state) | WinApi.MK_CONTROL)
}

