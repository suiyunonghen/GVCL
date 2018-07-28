package  Graphics

import (
	"io"
	"errors"
	"github.com/suiyunonghen/GVCL/WinApi"
	"unsafe"
	"reflect"
	"image"
	"image/color"
)

var ErrUnsupportedBmp = errors.New("bmp: unsupported BMP image")


type GBitmapCanvas struct {
	*GCanvas
	fmemDc			WinApi.HDC
	oldplate		WinApi.HPALETTE
}

func (cvs *GBitmapCanvas)Free()  {
	if cvs.fmemDc != 0{
		if cvs.oldplate != 0{
			WinApi.SelectPalette(cvs.fmemDc,cvs.oldplate,true)
		}
		WinApi.DeleteDC(cvs.fmemDc)
		cvs.fmemDc = 0
	}
}

func readUint16(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

func readUint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

type   GBitmap struct {
	topDown				bool
	fBufferBmp			WinApi.HBITMAP
	fPalate				WinApi.HPALETTE
	fbmpBuffer			[]byte
	fcanvas				*GBitmapCanvas
	fBmpData			uintptr			//位图数据位置
}

func (bmp *GBitmap)BmpHeader()*WinApi.BITMAPINFOHEADER  {
	if len(bmp.fbmpBuffer)>0{
		return (*WinApi.BITMAPINFOHEADER)(unsafe.Pointer(&(bmp.fbmpBuffer[0])))
	}
	return nil
}

func (bmp *GBitmap)BmpInfo()*WinApi.BITMAPINFO  {
	if len(bmp.fbmpBuffer)>0{
		return (*WinApi.BITMAPINFO)(unsafe.Pointer(&(bmp.fbmpBuffer[0])))
	}
	return nil
}

/*func (bmp *GBitmap)Pix()[]uint8  {
	//image.Image()
}*/

func (bmp *GBitmap)Bounds()(result image.Rectangle)  {
	result.Min.X = 0
	result.Min.Y = 0
	bmpheader := bmp.BmpHeader()
	result.Max.X = int(bmpheader.BiWidth)
	result.Max.Y = int(bmpheader.BiHeight)
	return
}

func (bmp *GBitmap)ColorModel() color.Model{
	bmph := bmp.BmpHeader()
	switch bmph.BiBitCount{
	case 8:
		//return color.
	case 24,32:
		return color.RGBAModel
	}
	return nil
}


func (bmp *GBitmap)At(x, y int) color.Color{
	bmph := bmp.BmpHeader()
	var pix reflect.SliceHeader
	switch bmph.BiBitCount{
	case 8:
		//查找索引表
		pix.Data = bmp.fBmpData
		pix.Len = int(bmph.BiSizeImage)
		pix.Cap = pix.Len
		Pix := *(*[]byte)(unsafe.Pointer(&pix))
		colortable := bmp.ColorTable(nil)
		if colortable != nil{
			rgbinfo := colortable[Pix[x*y]]
			return rgbinfo
		}
	case 24:
		var rgb WinApi.RGBQUAD
		scanline := int((3*bmph.BiWidth+3)&^3) //每行的长度
		offset := uintptr(y*scanline + x*3)
		rgb = *(*WinApi.RGBQUAD)(unsafe.Pointer(bmp.fBmpData + offset))
		rgb.RgbReserved = 0xff
		return rgb
	case 32:
		var rgb WinApi.RGBQUAD
		scanline := int(bmph.BiWidth * 4)
		offset := uintptr(y*scanline + x*4)
		rgb = *(*WinApi.RGBQUAD)(unsafe.Pointer(bmp.fBmpData + offset))
		return rgb
	}
	return nil
}

func (bmp *GBitmap)Canvas()ICanvas  {
	if bmp.fcanvas == nil{
		bmpcanvas := new(GBitmapCanvas)
		bmpcanvas.fmemDc = WinApi.CreateCompatibleDC(0)
		bmpcanvas.GCanvas = NewCanvas()
		bmpcanvas.GCanvas.SetHandle(bmpcanvas.fmemDc)
		WinApi.SelectObject(bmpcanvas.fmemDc,uintptr(bmp.fBufferBmp))
		if bmp.fPalate != 0{
			bmpcanvas.oldplate = WinApi.SelectPalette(bmpcanvas.fmemDc,bmp.fPalate,true)
			WinApi.RealizePalette(bmpcanvas.fmemDc)
		}
		bmp.fcanvas = bmpcanvas
	}
	return bmp.fcanvas
}

func (bmp *GBitmap)Free()  {
	if bmp.fBufferBmp != 0{
		WinApi.DeleteObject(WinApi.HGDIOBJ(bmp.fBufferBmp))
		bmp.fBufferBmp = 0
	}
	if bmp.fPalate != 0{
		WinApi.DeleteObject(WinApi.HGDIOBJ(bmp.fPalate))
		bmp.fPalate = 0
	}
	if bmp.fcanvas != nil{
		bmp.fcanvas.Free()
	}
}
//返回颜色表
func (bmp *GBitmap)ColorTable(bmpinfo *WinApi.BITMAPINFO)[]WinApi.RGBQUAD  {
	if bmpinfo == nil{
		bmpinfo = bmp.BmpInfo()
	}
	if bmpinfo != nil && bmpinfo.BmiHeader.BiBitCount < 24{
		var ctable 	reflect.SliceHeader
		ctable.Len = 1 << bmpinfo.BmiHeader.BiBitCount
		ctable.Cap = ctable.Len
		ctable.Data = uintptr(unsafe.Pointer(&bmpinfo.BmiColors))
		return *(*[]WinApi.RGBQUAD)(unsafe.Pointer(&ctable))
	}
	return nil
}

func (bmp *GBitmap)createPalate(bmpinfo *WinApi.BITMAPINFO)  {
	if bmpinfo.BmiHeader.BiBitCount < 24{ //颜色表
		var logplate WinApi.LOGPALETTE
		colortable := bmp.ColorTable(bmpinfo)
		for i := 0;i<len(colortable);i++{

			logplate.PalEntry[i].Red = colortable[i].RgbRed
			logplate.PalEntry[i].Green = colortable[i].RgbGreen
			logplate.PalEntry[i].Blue = colortable[i].RgbBlue
			logplate.PalEntry[i].Flags = colortable[i].RgbReserved
		}
		//创建调色板
		logplate.Version = 0x300
		logplate.NumEntries = 256
		bmp.fPalate = WinApi.CreatePalette(&logplate)
		if bmp.fPalate != 0{
			if bmp.fcanvas != nil{
				oldpalte := WinApi.SelectPalette(bmp.fcanvas.fmemDc,bmp.fPalate,true)
				WinApi.RealizePalette(bmp.fcanvas.fmemDc)
				if bmp.fcanvas.oldplate == 0{
					bmp.fcanvas.oldplate = oldpalte
				}
			}

		}
	}
}

func (bmp *GBitmap)newCompatibleBmp()(error) {
	tmpdc := WinApi.GetDC(0)
	defer WinApi.ReleaseDC(0,tmpdc)
	bmpinfo := bmp.BmpInfo()
	//首先读取颜色表,然后填充颜色版
	bmp.createPalate(bmpinfo)
	bmp.fBmpData = 0
	bmp.fBufferBmp = WinApi.CreateDIBSection(tmpdc, bmpinfo, WinApi.DIB_RGB_COLORS,
		&bmp.fBmpData, 0, 0)
	if bmp.fBmpData == 0 || bmp.fBufferBmp == 0{
		return errors.New("Create Compatible Bitmap Failed")
	}
	return nil
}

func (bmp *GBitmap)decodeConfig(r io.Reader)error  {
	const (
		fileHeaderLen = 14
		infoHeaderLen = 40
	)
	var b [fileHeaderLen+infoHeaderLen]byte
	if _, err := io.ReadFull(r, b[:fileHeaderLen+infoHeaderLen]); err != nil {
		return  err
	}
	if string(b[:2]) != "BM" {
		return errors.New("bmp: invalid format")
	}
	offset := readUint32(b[10:14])
	if readUint32(b[14:18]) != infoHeaderLen {
		return ErrUnsupportedBmp
	}
	if len(bmp.fbmpBuffer) < 40{
		bmp.fbmpBuffer = make([]byte,1280)
	}
	copy(bmp.fbmpBuffer[:40],b[14:])
	bmp.topDown = false
	bmpheader := bmp.BmpHeader()
	if bmpheader.BiHeight < 0 {
		bmpheader.BiHeight, bmp.topDown = -bmpheader.BiHeight, true
	}
	if bmpheader.BiWidth < 0 || bmpheader.BiHeight < 0 {
		return ErrUnsupportedBmp
	}
	if bmpheader.BiPlanes != 1 || bmpheader.BiCompression != 0 {
		return ErrUnsupportedBmp
	}



	switch bmpheader.BiBitCount {
	case 8: //8像素
		if offset != fileHeaderLen+infoHeaderLen+256*4 {
			return  ErrUnsupportedBmp
		}
		bmpheader.BiClrUsed = 256
		_, err := io.ReadFull(r, bmp.fbmpBuffer[40:1064]) //读取颜色表
		if err != nil {
			return  err
		}
		return bmp.newCompatibleBmp()
	case 24,32:
		if offset != fileHeaderLen+infoHeaderLen {
			return  ErrUnsupportedBmp
		}
		return bmp.newCompatibleBmp()
	}
	return ErrUnsupportedBmp
}

func (bmp *GBitmap)Draw(dc WinApi.HDC,x,y int)  {
	var oldplate WinApi.HPALETTE
	if bmp.fPalate != 0{
		oldplate = WinApi.SelectPalette(dc,bmp.fPalate,true)
		WinApi.RealizePalette(dc)
	}
	canvas := bmp.Canvas()
	/*r := WinApi.Rect{20,20,100,100}
	canvas.Brush().BrushStyle = BSClear
	WinApi.DrawText(canvas.GetHandle(),"asdfasdf",-1,&r,WinApi.DT_CENTER | WinApi.DT_VCENTER | WinApi.DT_SINGLELINE)*/
	bmpHeader := bmp.BmpHeader()
	WinApi.BitBlt(dc, x, y, int(bmpHeader.BiWidth), int(bmpHeader.BiHeight), canvas.GetHandle(), 0, 0, WinApi.SRCCOPY)
	if oldplate != 0{
		WinApi.SelectPalette(dc,oldplate,true)
	}
}

func (bmp *GBitmap)decodePaletted(r io.Reader)error  {
	bmpHeader := bmp.BmpHeader()
	if bmpHeader.BiWidth == 0 || bmpHeader.BiHeight == 0 {
		return nil
	}
	var pix reflect.SliceHeader
	pix.Data = bmp.fBmpData
	pix.Len = int(bmpHeader.BiSizeImage)
	pix.Cap = pix.Len
	Pix := *(*[]byte)(unsafe.Pointer(&pix))
	if _, err := io.ReadFull(r,Pix);err!=nil{
		return err
	}
	return nil
}

func (bmp *GBitmap)decodeRGB(r io.Reader)error  {
	bmpHeader := bmp.BmpHeader()
	bmpHeader.BiSizeImage = uint32(int((3*bmpHeader.BiWidth+3)&^3) * int(bmpHeader.BiHeight))
	var pix reflect.SliceHeader
	pix.Data = bmp.fBmpData
	pix.Len = int(bmpHeader.BiSizeImage)
	pix.Cap = pix.Len
	Pix := *(*[]byte)(unsafe.Pointer(&pix))
	if _, err := io.ReadFull(r,Pix);err!=nil{
		return err
	}
	return nil
}

func (bmp *GBitmap)decodeNRGBA(r io.Reader)error  {
	return bmp.decodePaletted(r)
}

func (bmp *GBitmap)Decode(r io.Reader) error {
	if bmp.fBufferBmp != 0{
		WinApi.DeleteObject(WinApi.HGDIOBJ(bmp.fBufferBmp))
		bmp.fBufferBmp = 0
	}
	if bmp.fPalate != 0{
		WinApi.DeleteObject(WinApi.HGDIOBJ(bmp.fPalate))
		bmp.fPalate = 0
	}
	err := bmp.decodeConfig(r)
	if err != nil{
		return err
	}
	switch bmp.BmpHeader().BiBitCount {
	case 8:
		return bmp.decodePaletted(r)
	case 24:
		return bmp.decodeRGB(r)
	case 32:
		return bmp.decodeNRGBA(r)
	}
	return ErrUnsupportedBmp
}