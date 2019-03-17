package main

import (
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"github.com/suiyunonghen/GVCL/WinApi"
	"image"
	"unsafe"
	"os"
	"image/png"
	"github.com/suiyunonghen/GVCL/Graphics"
	"fmt"
	"github.com/suiyunonghen/GVCL/Components/DxControls/WindowLessControls"
	"time"
)

func main()  {
	app := controls.NewApplication()
	m := app.CreateForm()
	m.SetLeft(200)
	m.SetTop(50)
	m.SetCaption("测试窗体")
	m.SetColor(Graphics.ClGreen)


	img := WindowLessControls.NewImage(m,Graphics.GDS_NORMAL)
	img.Picture.LoadFromFile(`E:\Delphi\SuperTech\TJDun\LocalServer\Res\001.png`)
	img.SetLeft(20)
	img.SetTop(20)
	img.SetWidth(400)
	img.SetHeight(400)

	btn := controls.NewButton(m)
	btn.SetCaption("画图")
	btn.OnClick = func(sender interface{}) {
		//f,_ := os.Open("E:\\1.bmp")
		f,_ := os.Open(`E:\Delphi\Leigod\LeigodAccelerator_Windows_New\qimiao\res\myinfoimg.png`)
		img,_ := png.Decode(f)
		f.Close()
		//绘图
		var m_pBmiPaint WinApi.BITMAPINFO
		//memdc := WinApi.CreateCompatibleDC(WinApi.GetDC(m.GetWindowHandle()))
		rgba := img.(*image.NRGBA)
		m_pBmiPaint.BmiHeader.BiBitCount = 32
		m_pBmiPaint.BmiHeader.BiClrImportant = 0
		m_pBmiPaint.BmiHeader.BiClrUsed = 0
		m_pBmiPaint.BmiHeader.BiCompression = WinApi.BI_RGB
		m_pBmiPaint.BmiHeader.BiHeight = int32(rgba.Rect.Max.Y - rgba.Rect.Min.Y)
		m_pBmiPaint.BmiHeader.BiPlanes = 1
		m_pBmiPaint.BmiHeader.BiWidth = int32(rgba.Rect.Max.X - rgba.Rect.Min.X)
		m_pBmiPaint.BmiHeader.BiXPelsPerMeter = 39
		m_pBmiPaint.BmiHeader.BiYPelsPerMeter = 39
		m_pBmiPaint.BmiHeader.BiSizeImage = uint32(len(rgba.Pix))
		m_pBmiPaint.BmiHeader.BiSize =uint32(unsafe.Sizeof(WinApi.BITMAPINFOHEADER{}))
		memdc := WinApi.GetDC(m.GetWindowHandle())
		//因为是倒着的，所以恢复对掉一下上下位置
		/*mb := make([]byte,m_pBmiPaint.BmiHeader.BiSizeImage)
		lastp := len(mb)
		for h := 0;h<252;h++{
			tmpb := mb[(251-h)*272*4:lastp]
			for w := 0;w<272*4;w++{
				 tmpb[w] = rgba.Pix[h*272*4+w]
			}
			lastp = (251-h)*272*4
		}*/
		//lastp := len(rgba.Pix)
		for h := 0;h<int(m_pBmiPaint.BmiHeader.BiHeight)/2;h++{
			for w := 0;w<int(m_pBmiPaint.BmiHeader.BiWidth);w++{
				/*rgba.Pix[(int(m_pBmiPaint.BmiHeader.BiHeight-1)-h)*
					int(m_pBmiPaint.BmiHeader.BiWidth)*4+w],rgba.Pix[h*int(m_pBmiPaint.BmiHeader.BiWidth)*4+w] = rgba.Pix[h*int(m_pBmiPaint.BmiHeader.BiWidth)*4+w],
					rgba.Pix[(int(m_pBmiPaint.BmiHeader.BiHeight-1)-h)*int(m_pBmiPaint.BmiHeader.BiWidth)*4+w]*/
				//在屏幕上显示的是RGBA，但是在读取出来的是BGRA
				/*for k := 0;k<4;k++{
					rgba.Pix[(int(m_pBmiPaint.BmiHeader.BiHeight-1)-h)*
						int(m_pBmiPaint.BmiHeader.BiWidth)*4+w+k],rgba.Pix[h*int(m_pBmiPaint.BmiHeader.BiWidth)*4+w+3-k] = rgba.Pix[h*int(m_pBmiPaint.BmiHeader.BiWidth)*4+w+3-k],
						rgba.Pix[(int(m_pBmiPaint.BmiHeader.BiHeight-1)-h)*int(m_pBmiPaint.BmiHeader.BiWidth)*4+w+k]
				}*/
				Last := (*WinApi.RGBQUAD)(unsafe.Pointer(&rgba.Pix[h*int(m_pBmiPaint.BmiHeader.BiWidth)*4+w*4]))
				first := (*WinApi.RGBQUAD)(unsafe.Pointer(&rgba.Pix[(int(m_pBmiPaint.BmiHeader.BiHeight-1)-h)*int(m_pBmiPaint.BmiHeader.BiWidth)*4+w*4]))
				tmp := *Last
				Last.RgbBlue = first.RgbRed
				Last.RgbGreen = first.RgbGreen
				Last.RgbRed = first.RgbBlue
				Last.RgbReserved = first.RgbReserved

				first.RgbBlue = tmp.RgbRed
				first.RgbGreen = tmp.RgbGreen
				first.RgbRed = tmp.RgbBlue
				first.RgbReserved = tmp.RgbReserved
			}
		}

		pb := uintptr(unsafe.Pointer(&rgba.Pix[0]))
		WinApi.SetDIBitsToDevice(memdc,80,40,uint32(m_pBmiPaint.BmiHeader.BiWidth),uint32(m_pBmiPaint.BmiHeader.BiHeight),
			0,0,0,uint(m_pBmiPaint.BmiHeader.BiHeight),pb,
			&m_pBmiPaint,0)

		//方法二
		/*m_pBmiPaint.BmiHeader.BiHeight = -1001
		m_pBmiPaint.BmiHeader.BiSizeImage = 0
		BufferDC := WinApi.CreateCompatibleDC(0)
		var BufferBits uintptr=0
		BufferBitmap := WinApi.CreateDIBSection(BufferDC, &m_pBmiPaint, WinApi.DIB_RGB_COLORS,
			&BufferBits, 0, 0)
		if BufferBits == 0{
			fmt.Println("Asdfasdf")
			return
		}
		OldBitmap := WinApi.SelectObject(BufferDC, uintptr(BufferBitmap))
		WinApi.BitBlt(BufferDC, 0, 0, 1334, 1001, memdc, 80, 40, WinApi.SRCCOPY)
		for h:=0;h<int(-m_pBmiPaint.BmiHeader.BiHeight);h++{
			curidx := h * int(m_pBmiPaint.BmiHeader.BiWidth) * 4
			for w := 0;w < int(m_pBmiPaint.BmiHeader.BiWidth);w++{
				tmp := (*WinApi.RGBQUAD)(unsafe.Pointer(uintptr(BufferBits)))
				//source := (*WinApi.RGBQUAD)(unsafe.Pointer(&rgba.Pix[curidx+w*4]))
				tmp.RgbRed = rgba.Pix[curidx+w*4]
				tmp.RgbGreen = rgba.Pix[curidx+w*4+1]
				tmp.RgbBlue = rgba.Pix[curidx+w*4+2]
				tmp.RgbReserved = 255
				BufferBits += 4
			}
		}
		//*
		WinApi.BitBlt(memdc, 80, 40, 1334, 1001, BufferDC, 0, 0, WinApi.SRCCOPY);
		WinApi.SelectObject(BufferDC, OldBitmap)
		WinApi.DeleteObject(WinApi.HGDIOBJ(BufferBitmap))
		WinApi.DeleteDC(BufferDC)*/
		WinApi.ReleaseDC(m.GetWindowHandle(),memdc)
	}


	btn.OnClick = func(sender interface{}) {
		go func() {
			str := ""
			for i := 1;i<=21;i++{
				if i < 9{
					str= fmt.Sprintf(`E:\Delphi\SuperTech\TJDun\LocalServer\Res\00%d.png`,i)
				}else{
					str = fmt.Sprintf(`E:\Delphi\SuperTech\TJDun\LocalServer\Res\0%d.png`,i)
				}
				img.Picture.LoadFromFile(str)
				time.Sleep(time.Millisecond*80)
			}
		}()


		return
		var Bitmap Graphics.GBitmap
		Bitmap.LoadFromFile(`E:\2.png`)
		//Bitmap.SetSize(300,300)
		/*f,_ := os.Open(`E:\3.bmp`)
		img,_ := bmp.Decode(f)
		f.Close()
		Bitmap.FromImage(img)*/
		memdc := WinApi.GetDC(m.GetWindowHandle())
		/*cvs := Bitmap.Canvas()
		r := WinApi.Rect{101,280,200,300}
		cvs.Brush().BrushStyle = Graphics.BSClear
		WinApi.DrawText(cvs.GetHandle(),"测试不得闲",-1,&r,WinApi.DT_CENTER | WinApi.DT_VCENTER | WinApi.DT_SINGLELINE)*/
		Bitmap.Draw(80,40,memdc)
		WinApi.ReleaseDC(m.GetWindowHandle(),memdc)


		fmt.Println(Bitmap.At(0,0))

		Bitmap.SaveToPngFile("e:\\mm.png")
		fmt.Println(Bitmap.At(0,0))
		Bitmap.Free()

	}

	app.Run()
}
