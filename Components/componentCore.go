package Components

import (
	"DxSoft/GVCL/Graphics"
	"syscall"
	"reflect"
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
}

type IApplication interface {
	Run()
	Active() bool
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
}

type GDispatchObj struct {
	TargetObject interface{}
	DispatchHandler reflect.Value
}
/**
**基本组件
**/
type GComponent struct {
	fname      string `json:"Name"`
	ftag       int    `json:"Tag"`
	fTagetData uintptr
	fSubChilds []interface{} //子对象
	fRealFreeDispatch *GDispatchObj
}

func (cmp *GComponent) GetName() string {
	return cmp.fname
}

func (cmp *GComponent) SubInit(subObj interface{}) {
	if cmp.fSubChilds == nil {
		cmp.fSubChilds = make([]interface{}, 0)
	}
	for _, v := range cmp.fSubChilds {
		if v == subObj {
			return
		}
	}
	cmp.fSubChilds = append(cmp.fSubChilds, subObj)
}

func (cmp *GComponent) SubChild(idx int) interface{} {
	if cmp.fSubChilds == nil {
		return nil
	}
	if idx >= 0 && idx < len(cmp.fSubChilds) {
		return cmp.fSubChilds[idx]
	}
	return nil
}

func (cmp *GComponent) SubChildCount() int {
	if cmp.fSubChilds == nil {
		return 0
	}
	return len(cmp.fSubChilds)
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

func (cmp *GComponent)Free()  {
	//执行Destroy过程
	if cmp.fRealFreeDispatch == nil{
		for i := cmp.SubChildCount() - 1; i >= 0; i-- {
			subObj := cmp.SubChild(i)
			pType := reflect.TypeOf(subObj)
			if mnd, ok := pType.MethodByName("Destroy"); ok {
				pType = mnd.Func.Type()
				if pType.NumIn() == 1 && pType.NumOut() == 0 {
					cmp.fRealFreeDispatch = new(GDispatchObj)
					cmp.fRealFreeDispatch.DispatchHandler = mnd.Func
					cmp.fRealFreeDispatch.TargetObject = subObj
					break
				}
			}
		}
	}
	if cmp.fRealFreeDispatch!=nil{
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(cmp.fRealFreeDispatch.TargetObject)
		 cmp.fRealFreeDispatch.DispatchHandler.Call(in)
	}
}
