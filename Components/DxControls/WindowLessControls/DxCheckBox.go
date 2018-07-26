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
	fCaption string
	fAutoSize bool
	fChecked	bool
	fIsMouseIn	bool
	fisMouseDown bool
	fWordWrap	bool
	OnChange	Graphics.NotifyEvent
}

func (checkbox *GDxCheckBox) SubInit() {
	checkbox.GBaseControl.SubInit()
	checkbox.GComponent.SubInit(checkbox)
	checkbox.SetTrasparent(true)
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

func (checkBox *GDxCheckBox)SetChecked(v bool)  {
	if checkBox.fChecked != v{
		checkBox.fChecked = v
		checkBox.Invalidate()
		if checkBox.OnChange!=nil{
			checkBox.OnChange(checkBox)
		}
	}
}

func (checkbox *GDxCheckBox)Checked()bool  {
	return checkbox.fChecked
}

func (checkbox *GDxCheckBox)Paint(cvs Graphics.ICanvas)  {
	//先绘制checkBox
	//居中绘制，大小14*14
	if !checkbox.Trasparent(){
		r := WinApi.Rect{0,0,checkbox.Width(),checkbox.Height()}
		cvs.FillRect(&r)
	}
	brush := cvs.Brush()
	checkRect := WinApi.Rect{0,0,14,14}
	checkRect.Top = (checkbox.Height() - checkRect.Bottom) / 2
	checkRect.Bottom = checkRect.Top + 14
	if checkbox.fChecked{
		brush.BeginUpdate()
		if checkbox.fIsMouseIn{
			if checkbox.fisMouseDown{
				brush.Color = Graphics.RGB(56,140,245)
			}else{
				brush.Color = Graphics.RGB(69,165,255)
			}
		}else{
			brush.Color = Graphics.RGB(66,149,252)
		}
		brush.EndUpdate()
		cvs.FillRect(&checkRect)
		//绘制勾选的
		pen := cvs.Pen()
		pen.Color = Graphics.ClWhite
		pen.Change()
		cvs.MoveTo(checkRect.Left + 3,checkRect.Top+7)
		cvs.LineTo(int(checkRect.Left) + 5,int(checkRect.Top)+9)
		cvs.LineTo(int(checkRect.Left) + 11,int(checkRect.Top)+3)
		if checkbox.fIsMouseIn{
			if checkbox.fisMouseDown{
				pen.Color = Graphics.RGB(118,176,248) //阴影
			}else{
				pen.Color = Graphics.RGB(127,193,255) //阴影
			}
		}else{
			pen.Color = Graphics.RGB(125,182,253) //阴影
		}
		pen.Change()
		cvs.MoveTo(checkRect.Left + 3,checkRect.Top+8)
		cvs.LineTo(int(checkRect.Left) + 5,int(checkRect.Top)+10)
		cvs.LineTo(int(checkRect.Left) + 11,int(checkRect.Top)+4)

		if checkbox.fIsMouseIn {
			if checkbox.fisMouseDown{
				pen.Color = Graphics.RGB(102,166,247) //阴影
			}else {
				pen.Color = Graphics.RGB(112, 185, 255)
			}
		}else{
			pen.Color = Graphics.RGB(109, 173, 252)
		}
		pen.Change()

		cvs.MoveTo(checkRect.Left,checkRect.Top)
		cvs.LineTo(int(checkRect.Left),int(checkRect.Top+1))

		cvs.MoveTo(checkRect.Left,checkRect.Bottom-1)
		cvs.LineTo(int(checkRect.Left+1),int(checkRect.Bottom-1))

		cvs.MoveTo(checkRect.Right - 1,checkRect.Top)
		cvs.LineTo(int(checkRect.Right),int(checkRect.Top))

		cvs.MoveTo(checkRect.Right - 1,checkRect.Bottom-1)
		cvs.LineTo(int(checkRect.Right),int(checkRect.Bottom-1))

	}else if checkbox.fIsMouseIn{
		brush.BeginUpdate()
		brush.Color = Graphics.RGB(66,149,252)
		brush.EndUpdate()
		cvs.FrameRect(&checkRect)
	}else{
		brush.BeginUpdate()
		brush.Color = Graphics.RGB(200,206,210)
		brush.EndUpdate()
		cvs.FrameRect(&checkRect)
	}
	//绘制Caption
	if checkbox.fCaption != ""{
		checkRect.Left = checkRect.Left + 17
		checkRect.Top = 0
		checkRect.Bottom = checkbox.Height()
		checkRect.Right = checkbox.Width()
		brush := cvs.Brush()
		brush.BrushStyle = Graphics.BSClear
		brush.Change()
		var drawflags uint = WinApi.DT_LEFT | WinApi.DT_TOP | WinApi.DT_CALCRECT
		if checkbox.fWordWrap{
			drawflags = WinApi.DT_LEFT | WinApi.DT_VCENTER | WinApi.DT_WORDBREAK
		}else{
			drawflags = WinApi.DT_LEFT | WinApi.DT_VCENTER | WinApi.DT_SINGLELINE
		}
		WinApi.DrawText(cvs.GetHandle(),checkbox.fCaption,-1,&checkRect,drawflags)
	}

}


func (checkbox *GDxCheckBox)MouseDown(button Components.MouseButton,x,y int,state Components.KeyState)bool  {
	checkbox.fisMouseDown = button == Components.MbLeft
	if checkbox.fisMouseDown{
		checkbox.Invalidate()
	}
	return true
}

func (checkbox *GDxCheckBox)MouseUp(button Components.MouseButton,x,y int,state Components.KeyState)bool  {
	if checkbox.fisMouseDown{
		checkbox.fisMouseDown = false
		checkbox.fChecked = !checkbox.fChecked
		checkbox.Invalidate()
		if checkbox.OnChange!=nil{
			checkbox.OnChange(checkbox)
		}
	}
	return true
}




func (checkbox *GDxCheckBox)MouseEnter(){
	checkbox.fIsMouseIn = true
	if checkbox.Enabled(){
		checkbox.Invalidate()
	}
	checkbox.GBaseControl.MouseEnter()
}

func (checkbox *GDxCheckBox)MouseLeave(){
	checkbox.fIsMouseIn = false
	if checkbox.Enabled(){
		checkbox.Invalidate()
	}
	checkbox.GBaseControl.MouseLeave()
}

func NewCheckBox(aParent Components.IWincontrol) *GDxCheckBox {
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