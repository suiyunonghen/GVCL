package WindowLessControls

import (
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/GVCL/Components"
	"reflect"
	"github.com/suiyunonghen/GVCL/WinApi"
)

type GDxImage struct {
	controls.GBaseControl
	fCaption string
	Picture	Graphics.GBitmap
	fDrawStyle	Graphics.GDrawStyle
	OnChange	Graphics.NotifyEvent
}

func (img *GDxImage)picChange(sender interface{})  {
	img.Invalidate()
}

func (img *GDxImage) SubInit() {
	img.GBaseControl.SubInit()
	img.GComponent.SubInit(img)
	img.SetTrasparent(true)
	img.SetWidth(300)
	img.SetHeight(300)
	img.fDrawStyle = Graphics.GDS_NORMAL
	img.Picture.OnChange = img.picChange
}

func (img *GDxImage)Paint(cvs Graphics.ICanvas) {
	w,h := img.Picture.Size()
	if w > 0 && h > 0{
		switch img.fDrawStyle {
		case Graphics.GDS_NORMAL:
			cvs.Draw(0,0,&img.Picture)
		case Graphics.GDS_CENTER:
			cvs.Draw((int(img.Width()) - w) /2,(int(img.Height()) - h)/2,&img.Picture)
		case Graphics.GDS_STRETCH:
			srcrect := WinApi.Rect{0,0,int32(w),int32(h)}
			destRect := WinApi.Rect{0,0,img.Width(),img.Height()}
			img.Picture.DrawToDest(srcrect,destRect,cvs.GetHandle())
		case Graphics.GDS_FILL:
			dc := cvs.GetHandle()
			srcdc := img.Picture.Canvas().GetHandle()
			starty := 0
			cwidth := int(img.Width())
			for starty < int(img.Height()){
				startx := 0
				for startx < cwidth{
					WinApi.BitBlt(dc,startx,starty,startx+w,starty+h,srcdc,0,0,WinApi.SRCCOPY)
					startx += w
				}
				starty += h
			}
		}
	}

}

func (img *GDxImage)AfterParentWndCreate(){

}

func NewImage(aParent Components.IWincontrol,drawStyle Graphics.GDrawStyle) *GDxImage {
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		img := new(GDxImage)
		img.SubInit()
		img.fDrawStyle = drawStyle
		img.SetParent(aParent)
		return img
	}
	return nil
}