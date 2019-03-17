package main

import (
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"github.com/suiyunonghen/GVCL/WinApi"
	"github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/DxCommonLib"
	"time"
)

func PngUi(frm *controls.GForm)  {
	var BlendFunction WinApi.BLENDFUNCTION
	DxCommonLib.Sleep(time.Second)
	BlendFunction.BlendOp = WinApi.AC_SRC_OVER
	BlendFunction.BlendFlags = 0
	BlendFunction.SourceConstantAlpha = 255
	BlendFunction.AlphaFormat = WinApi.AC_SRC_ALPHA

	handle := frm.GetWindowHandle()
	dc := WinApi.GetDC(handle)
	ptDst := WinApi.POINT{frm.Left(),frm.Top()}
	frm.SetWidth(632)
	frm.SetHeight(643)
	osize := WinApi.GSize{632,643}
	ptSrc := WinApi.POINT{}
	WinApi.SetWindowLong(handle,WinApi.GWL_STYLE,WinApi.GetWindowLong(handle,WinApi.GWL_STYLE) & (^WinApi.WS_CAPTION) & (^WinApi.WS_THICKFRAME))
	WinApi.SetWindowLong(handle,WinApi.GWL_EXSTYLE,WinApi.GetWindowLong(handle, WinApi.GWL_EXSTYLE) | WinApi.WS_EX_LAYERED)
	var tmpbmp Graphics.GBitmap

	//str = ""
	for i := 1;i<=1;i++{
		/*if i < 9{
			str= fmt.Sprintf(`E:\Delphi\SuperTech\TJDun\LocalServer\Res\00%d.png`,i)
		}else{
			str = fmt.Sprintf(`E:\Delphi\SuperTech\TJDun\LocalServer\Res\0%d.png`,i)
		}*/
		str := `e:\2.png`
		tmpbmp.LoadFromFile(str)

		//var destbmp Graphics.GBitmap
		//destbmp.SetSize(int(frm.Width()),int(frm.Height()),true)
		//cvs := destbmp.Canvas()
		//cvs.Draw(0,0,&tmpbmp)
		WinApi.UpdateLayeredWindow(handle,dc,&ptDst,&osize,tmpbmp.Canvas().GetHandle(),&ptSrc,0,&BlendFunction,WinApi.ULW_ALPHA)
		//destbmp.Free()
		DxCommonLib.Sleep(100)
	}
	WinApi.ReleaseDC(handle,dc)
}

func main()  {
	app := controls.NewApplication()
	m := app.CreateForm()
	m.SetLeft(200)
	m.SetTop(50)
	//m.SetWidth(650)
	//m.SetHeight(650)
	m.SetCaption("测试窗体")




	go PngUi(m)

	app.Run()
}