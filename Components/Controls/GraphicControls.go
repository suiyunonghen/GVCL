package controls

import (
	"github.com/suiyunonghen/GVCL/WinApi"
	"github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/GVCL/Components"
	"reflect"
)

type GLabel struct {
	GBaseControl
	fTrasparent bool
	fWordWrap  bool
	fCaption string
	fAutoSize bool
}

func (lbl *GLabel) SubInit() {
	lbl.GBaseControl.SubInit()
	lbl.GComponent.SubInit(lbl)
	lbl.fwidth = 100
	lbl.fheight = 20
	lbl.fVisible = true
	lbl.fColor = Graphics.ClBtnFace
}

func (lbl *GLabel)calcSize()  {
	cvs := Graphics.NewCanvas()
	fhandle := lbl.GetTargetCanvas(cvs)
	if fhandle == 0{
		return
	}
	sz := new(WinApi.GSize)
	oldr := new(WinApi.Rect)
	cvs.RefreshCanvasState()
	oldr.Right = lbl.Width()
	oldr.Bottom = lbl.Height()
	cvs.Font().Assign(&lbl.Font)
	if cvs.TextExtent(lbl.fCaption,sz){
		lbl.fwidth = sz.CX
		lbl.fheight = sz.CY
		if oldr.Right < sz.CX{
			oldr.Right = sz.CX
		}
		if oldr.Bottom < sz.CY{
			oldr.Bottom = sz.CY
		}
		if oldr.Right != sz.CX || oldr.Bottom != sz.CY{
			oldr.OffsetRect(int(lbl.fleft),int(lbl.ftop))
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

func (lbl *GLabel)SetAutoSize(v bool)  {
	if lbl.fAutoSize != v{
		lbl.fAutoSize = v
		if lbl.fAutoSize{
			lbl.calcSize()
		}
	}
}

func (lbl *GLabel)AfterParentWndCreate()  {
	if lbl.fAutoSize{
		lbl.calcSize()
	}
}

func (lbl *GLabel)Paint(cvs Graphics.ICanvas)  {
	r := WinApi.Rect{0,0,lbl.Width(),lbl.Height()}
	var drawflags uint = WinApi.DT_LEFT | WinApi.DT_TOP | WinApi.DT_CALCRECT
	if !lbl.fTrasparent{
		cvs.FillRect(&r)
	}
	if lbl.fWordWrap{
		drawflags = WinApi.DT_CENTER | WinApi.DT_VCENTER | WinApi.DT_WORDBREAK
	}else{
		drawflags = WinApi.DT_CENTER | WinApi.DT_VCENTER | WinApi.DT_SINGLELINE
	}
	if lbl.fCaption != ""{
		WinApi.DrawText(cvs.GetHandle(),lbl.fCaption,-1,&r,drawflags)
	}
}

func (lbl *GLabel)SetTrasparent(v bool)  {
	if lbl.fTrasparent != v{
		lbl.fTrasparent = v
		lbl.Invalidate()
	}
}

func (lbl *GLabel)SetCaption(cap string)  {
	if lbl.fCaption != cap{
		lbl.fCaption = cap
		if lbl.fAutoSize{
			lbl.calcSize()
		}
	}
}

func NewLabel(aParent Components.IWincontrol) *GLabel {
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		lbl := new(GLabel)
		lbl.SubInit()
		lbl.SetParent(aParent)
		return lbl
	}
	return nil
}
