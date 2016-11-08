package Graphics
import (
	"suiyunonghen/GVCL/WinApi"
)

type GCanvas struct {
	GObject
	fHandle WinApi.HDC
	fBrush *GBrush
	fFont  *GFont
	fPen   *GPen
}

type ICanvas interface {
	GetHandle()WinApi.HDC
	SetHandle(dc WinApi.HDC)
	MoveTo(x,y int32)
	GetClipRect(*WinApi.Rect)
	FrameRect(r *WinApi.Rect)
	FillRect(r *WinApi.Rect)
	LineTo(x,y int)
	Brush()*GBrush
	Pen()*GPen
	Font()*GFont
	CreateHandle()
	TextExtent(str string,size *WinApi.GSize)bool
	TextWidth(str string)(int32,bool)
	TextHeight(str string)(int32,bool)
}

func NewCanvas()*GCanvas  {
	cvs := new(GCanvas)
	cvs.SubInit()
	return cvs
}

func (cvs *GCanvas) SubInit() {
	cvs.GObject.SubInit(cvs)
}

func (cvs *GCanvas) TextExtent(str string,size *WinApi.GSize)bool {
	if cvs.GetHandle() != 0{
		r := new(WinApi.Rect)
		r.Right = 10000
		r.Bottom = 10000
		WinApi.DrawText(cvs.fHandle,str,-1,r,WinApi.DT_LEFT | WinApi.DT_TOP | WinApi.DT_CALCRECT)
		size.CX = r.Width()
		size.CY = r.Height()
		return  true
	}
	return false
}

func (cvs *GCanvas) TextWidth(str string)(w int32,b bool) {
	w = 0
	b = cvs.GetHandle() != 0
	if b{
		r := new(WinApi.Rect)
		WinApi.DrawText(cvs.fHandle,str,-1,r,WinApi.DT_LEFT | WinApi.DT_TOP)
		w = r.Width()
	}
	return
}

func (cvs *GCanvas) TextHeight(str string)(h int32, b bool) {
	h = 0
	b = cvs.GetHandle() != 0
	if b{
		r := new(WinApi.Rect)
		WinApi.DrawText(cvs.fHandle,str,-1,r,WinApi.DT_LEFT | WinApi.DT_TOP)
		h = r.Width()
	}
	return
}

func (cvs *GCanvas)CreateHandle()  {
	if cvs.fHandle!=0{
		cvs.RefreshCanvasState()
	}
}

func (cvs *GCanvas)RefreshCanvasState()  {
	if cvs.fBrush == nil{
		cvs.fBrush = NewBrush()
	}
	if cvs.fFont == nil{
		cvs.fFont = NewFont()
	}
	if cvs.fPen == nil{
		cvs.fPen = NewPen()
	}
	WinApi.UnrealizeObject(uintptr(cvs.fBrush.Handle))
	WinApi.SelectObject(cvs.fHandle,uintptr(cvs.fBrush.Handle))
	if cvs.fBrush.BrushStyle == BSSolid{
		WinApi.SetBkColor(cvs.fHandle, uint32(cvs.fBrush.Color))
		WinApi.SetBkMode(cvs.fHandle, WinApi.OPAQUE)
	}else{
		WinApi.SetBkColor(cvs.fHandle, ^uint32(cvs.fBrush.Color))
		WinApi.SetBkMode(cvs.fHandle, WinApi.TRANSPARENT)
	}
	WinApi.SelectObject(cvs.fHandle,uintptr(cvs.fFont.FontHandle))
	WinApi.SetTextColor(cvs.fHandle, uint32(cvs.fFont.Color))

	WinApi.SelectObject(cvs.fHandle,uintptr(cvs.fPen.Handle))
}

func (cvs *GCanvas)Destroy()  {
	if cvs.fFont != nil{
		cvs.fFont.Destroy()
		cvs.fFont = nil
	}

	if cvs.fBrush != nil{
		cvs.fBrush.Destroy()
		cvs.fBrush = nil
	}

	if cvs.fPen != nil{
		cvs.fPen.Destroy()
		cvs.fPen = nil
	}
}

func (cvs *GCanvas)GetHandle()WinApi.HDC  {
	if cvs.fHandle == 0{
		var target interface{}
		if i := cvs.SubChildCount() - 1; i >= 0{
			target = cvs.SubChild(i)
		}else{
			target = cvs
		}
		target.(ICanvas).CreateHandle()
	}else{
		cvs.RefreshCanvasState()
	}
	return cvs.fHandle
}

func (cvs *GCanvas)SetHandle(dc WinApi.HDC)  {
	if cvs.fHandle == dc{
		return
	}
	var penPos *WinApi.POINT = nil
	if cvs.fHandle != 0{
		penPos = new(WinApi.POINT)
		WinApi.GetCurrentPositionEx(cvs.fHandle,penPos)
	}
	cvs.fHandle = dc
	if cvs.fHandle !=0 && penPos != nil{
		cvs.MoveTo(int32(penPos.X),int32(penPos.Y))
	}
}

func (cvs *GCanvas)MoveTo(x,y int32)  {
	if cvs.GetHandle() != 0{
		WinApi.MoveToEx(cvs.fHandle,x,y,nil)
	}
}

func (cvs *GCanvas)GetClipRect(rect *WinApi.Rect)  {
	if cvs.GetHandle() != 0{
		WinApi.GetClipBox(cvs.fHandle,rect)
	}
}

func (cvs *GCanvas)GetPenPos(pt *WinApi.POINT)  {
	if cvs.GetHandle() !=0{
		WinApi.GetCurrentPositionEx(cvs.fHandle,pt)
	}
}

func (cvs *GCanvas)FrameRect(r *WinApi.Rect)  {
	if cvs.GetHandle() !=0{
		WinApi.FrameRect(cvs.fHandle,r,cvs.fBrush.Handle)
	}
}

func (cvs *GCanvas)FillRect(r *WinApi.Rect)  {
	if cvs.GetHandle() !=0 {
		WinApi.FillRect(cvs.fHandle,r,cvs.fBrush.Handle)
	}
}

func (cvs *GCanvas)LineTo(x,y int)  {
	if cvs.GetHandle() != 0{
		WinApi.LineTo(cvs.fHandle,x,y)
	}
}

func (cvs *GCanvas)Brush()*GBrush  {
	if cvs.fBrush == nil{
		cvs.fBrush = NewBrush()
	}
	return cvs.fBrush
}

func (cvs *GCanvas)Pen()*GPen  {
	if cvs.fPen == nil{
		cvs.fPen = NewPen()
	}
	return cvs.fPen
}

func (cvs *GCanvas)Font()*GFont  {
	if cvs.fFont == nil{
		cvs.fFont = NewFont()
	}
	return cvs.fFont
}

