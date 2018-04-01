package WindowLessControls

import (
	"github.com/suiyunonghen/GVCL/WinApi"
	"github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"github.com/suiyunonghen/GVCL/Components"
	"reflect"
)

type GDxCheckBox struct {
	controls.GBaseControl
	fTrasparent bool
	fCaption string
	fAutoSize bool
}

func (checkbox *GDxCheckBox) SubInit() {
	checkbox.GBaseControl.SubInit()
	checkbox.GComponent.SubInit(checkbox)
	checkbox.SetWidth(100)
	checkbox.SetHeight(20)
}


func (checkbox *GDxCheckBox)calcSize()  {
	cvs := Graphics.NewCanvas()
	fhandle := checkbox.GetTargetCanvas(cvs)
	if fhandle == 0{
		return
	}
	sz := new(WinApi.GSize)
	oldr := new(WinApi.Rect)
	cvs.RefreshCanvasState()
	oldr.Right = checkbox.Width()
	oldr.Bottom = checkbox.Height()
	cvs.Font().Assign(&checkbox.Font)
	if cvs.TextExtent(checkbox.fCaption,sz){
		checkbox.SetWidth(sz.CX + 18); //加上选择区域
		if sz.CY < 16{
			sz.CY = 16
		}
		if oldr.Right < sz.CX{
			oldr.Right = sz.CX
		}
		if oldr.Bottom < sz.CY{
			oldr.Bottom = sz.CY
		}
		if oldr.Right != sz.CX || oldr.Bottom != sz.CY{
			oldr.OffsetRect(int(checkbox.Left()),int(checkbox.Top()))
			if fhandle != 0{
				WinApi.InvalidateRect(fhandle,oldr,false)
			}
		}
	}
	dc := cvs.GetHandle()
	cvs.SetHandle(0)
	WinApi.ReleaseDC(fhandle,dc)
	cvs.Destroy()
}

func (checkbox *GDxCheckBox)SetAutoSize(v bool)  {
	if checkbox.fAutoSize != v{
		checkbox.fAutoSize = v
		if checkbox.fAutoSize{
			if pc := checkbox.GetParent();pc!=nil && pc.HandleAllocated(){
				checkbox.calcSize()
			}
		}
	}
}

func (checkbox *GDxCheckBox)AfterParentWndCreate()  {
	if checkbox.fAutoSize{
		checkbox.calcSize()
	}
}


func (checkbox *GDxCheckBox)SetCaption(cap string)  {
	if checkbox.fCaption != cap{
		checkbox.fCaption = cap
		if checkbox.fAutoSize{
			checkbox.calcSize()
		}
	}
}

func (checkbox *GDxCheckBox)Paint(cvs Graphics.ICanvas)  {
	//先绘制checkBox
	//居中绘制，大小14*14
	checkRect := WinApi.Rect{0,0,14,14}
	checkRect.Top = (checkbox.Height() - checkRect.Bottom) / 2
	checkRect.Bottom = checkRect.Top + 14

}

func NewLabel(aParent Components.IWincontrol) *GDxCheckBox {
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		checkbox := new(GDxCheckBox)
		checkbox.SubInit()
		checkbox.SetParent(aParent)
		return checkbox
	}
	return nil
}