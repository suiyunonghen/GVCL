package  Graphics

import (
	"io"
	"errors"
	"github.com/suiyunonghen/GVCL/WinApi"
	"unsafe"
	"reflect"
	"image"
	"image/color"
	"os"
	"image/png"
	"image/jpeg"
	"golang.org/x/image/bmp"
	"fmt"
)

var (
	ErrUnsupportedBmp = errors.New("bmp: unsupported BMP image")
	ErrNotBmpFormat = errors.New("bmp: invalid format")
)


const (
	fileHeaderLen = 14
	infoHeaderLen = 40
)

var(
	RGBQUADSize = unsafe.Sizeof(WinApi.RGBQUAD{})
)

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
	cvs.GCanvas.Free()
}

func readUint16(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

func readUint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

type   GBitmap struct {
	fBufferBmp			WinApi.HBITMAP
	fPalate				WinApi.HPALETTE
	fbmpBuffer			[]byte
	fcanvas				*GBitmapCanvas
	fBmpData			uintptr			//位图数据位置
	fcolorMode			color.Model
}

func (bmp *GBitmap)BmpHeader()*WinApi.BITMAPINFOHEADER  {
	if bmp.fbmpBuffer == nil{
		bmp.fbmpBuffer = make([]byte,1280)
		(*WinApi.BITMAPINFOHEADER)(unsafe.Pointer(&(bmp.fbmpBuffer[0]))).BiSize = infoHeaderLen
	}
	return (*WinApi.BITMAPINFOHEADER)(unsafe.Pointer(&(bmp.fbmpBuffer[0])))
}

func (bmp *GBitmap)BmpInfo()*WinApi.BITMAPINFO  {
	if bmp.fbmpBuffer == nil{
		bmp.fbmpBuffer = make([]byte,1280)
	}
	return (*WinApi.BITMAPINFO)(unsafe.Pointer(&(bmp.fbmpBuffer[0])))
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
	if result.Max.Y < 0{
		result.Max.Y = -result.Max.Y
	}
	return
}

func (bmp *GBitmap)ColorModel() color.Model{
	if bmp.fcolorMode != nil{
		return bmp.fcolorMode
	}
	bmph := bmp.BmpHeader()
	switch bmph.BiBitCount{
	case 8:
		pcm := make(color.Palette, 256)
		colortable := bmp.ColorTable(nil)
		for i := range pcm {
			pcm[i] = color.RGBA{colortable[i].RgbRed, colortable[i].RgbGreen, colortable[i].RgbBlue, 0xFF}
		}
		bmp.fcolorMode = pcm
		return bmp.fcolorMode
	case 24,32:
		return color.RGBAModel
	}
	return nil
}


func (bmp *GBitmap)At(x, y int) color.Color{
	bmph := bmp.BmpHeader()
	if bmp.BmpHeader().BiHeight > 0{ //排列是从下到上
		y = int(bmp.BmpHeader().BiHeight) - y - 1
	}
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
		if x == 0 && y == 1000{
			fmt.Println("Asdf")
		}
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
		bmp.fcanvas = nil
	}
	if bmp.fPalate != 0{
		WinApi.DeleteObject(WinApi.HGDIOBJ(bmp.fPalate))
		bmp.fPalate = 0
	}
	bmp.fbmpBuffer = nil
	bmp.fBmpData = 0
	bmp.fcolorMode = nil
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

func createPalate(bmpinfo *WinApi.BITMAPINFO)(result WinApi.HPALETTE)  {
	result = 0
	if bmpinfo.BmiHeader.BiBitCount < 24{ //颜色表
		var logplate WinApi.LOGPALETTE
		datalen := 1 << bmpinfo.BmiHeader.BiBitCount
		for i := 0;i< datalen;i++{
			Rgbquad := (*WinApi.RGBQUAD)(unsafe.Pointer( uintptr(unsafe.Pointer(&bmpinfo.BmiColors))+RGBQUADSize))
			logplate.PalEntry[i].Red = Rgbquad.RgbRed
			logplate.PalEntry[i].Green = Rgbquad.RgbGreen
			logplate.PalEntry[i].Blue = Rgbquad.RgbBlue
			logplate.PalEntry[i].Flags = Rgbquad.RgbReserved
		}
		//创建调色板
		logplate.Version = 0x300
		logplate.NumEntries = uint16(datalen)
		result = WinApi.CreatePalette(&logplate)
	}
	return result
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

func (bmp *GBitmap)Size()(w,h int)  {
	bmph := bmp.BmpHeader()
	w = int(bmph.BiWidth)
	h = int(bmph.BiHeight)
	if h < 0{
		h = -h
	}
	return
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
	var b [fileHeaderLen+infoHeaderLen]byte
	if _, err := io.ReadFull(r, b[:fileHeaderLen+infoHeaderLen]); err != nil {
		return  err
	}
	if string(b[:2]) != "BM" {
		return ErrNotBmpFormat
	}
	offset := readUint32(b[10:14])
	if readUint32(b[14:18]) != infoHeaderLen {
		return ErrUnsupportedBmp
	}
	if len(bmp.fbmpBuffer) < 40{
		bmp.fbmpBuffer = make([]byte,1280)
	}
	copy(bmp.fbmpBuffer[:40],b[14:])
	bmpheader := bmp.BmpHeader()
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


func (bmp *GBitmap)DrawToDest(srcRect WinApi.Rect,destRect WinApi.Rect,destDc WinApi.HDC){

}

func (bmp *GBitmap)Draw(x,y int,dc WinApi.HDC)  {
	var oldplate WinApi.HPALETTE
	if bmp.fPalate != 0{
		oldplate = WinApi.SelectPalette(dc,bmp.fPalate,true)
		WinApi.RealizePalette(dc)
	}
	bmpHeader := bmp.BmpHeader()
	w,h := bmp.Size()
	if bmpHeader.BiBitCount < 32{
		canvas := bmp.Canvas()
		WinApi.BitBlt(dc, x, y, w,h, canvas.GetHandle(), 0, 0, WinApi.SRCCOPY)
		if oldplate != 0{
			WinApi.SelectPalette(dc,oldplate,true)
		}
	}else{ //具备有透明通道，需要绘制透明通道
		BufferBits := uintptr(0)
		var bmpinfo WinApi.BITMAPINFO
		bmpinfo.BmiHeader = *bmpHeader
		BufferDC := WinApi.CreateCompatibleDC(0)
		if BufferDC == 0{
			return
		}
		BufferBitmap := WinApi.CreateDIBSection(BufferDC, &bmpinfo, WinApi.DIB_RGB_COLORS,
			&BufferBits, 0, 0)
		if BufferBitmap == 0 || BufferBits == 0{
			if BufferBitmap != 0 {
				WinApi.DeleteObject(WinApi.HGDIOBJ(BufferBitmap))
			}
			WinApi.DeleteDC(BufferDC)
		}
		OldBitmap := WinApi.SelectObject(BufferDC, uintptr(BufferBitmap))
		WinApi.BitBlt(BufferDC, 0, 0, w, h, dc, x, y, WinApi.SRCCOPY)

		//根据透明色重新计算颜色
		pixsize := unsafe.Sizeof(WinApi.RGBQUAD{})
		srcData := bmp.fBmpData
		destData := BufferBits
		for row := 0;row < h;row++{
			for col := 0;col < w;col++{
				srcrgba := (*WinApi.RGBQUAD)(unsafe.Pointer(srcData))
				destrgba := (*WinApi.RGBQUAD)(unsafe.Pointer(destData))
				if srcrgba.RgbReserved != 255{ //根据透明度融合背景色
					destrgba.RgbRed = byte((0x7F + int(srcrgba.RgbReserved)*int(srcrgba.RgbRed) + int(destrgba.RgbRed)*int(^srcrgba.RgbReserved)) / 0xFF)
					destrgba.RgbGreen = byte((0x7F + int(srcrgba.RgbReserved)*int(srcrgba.RgbGreen) + int(destrgba.RgbGreen)*int(^srcrgba.RgbReserved)) / 0xFF)
					destrgba.RgbBlue = byte((0x7F + int(srcrgba.RgbReserved)*int(srcrgba.RgbBlue) + int(destrgba.RgbBlue)*int(^srcrgba.RgbReserved)) / 0xFF)
					destrgba.RgbReserved = byte(^((0x7F + int(^destrgba.RgbReserved)*int(^srcrgba.RgbReserved)) / 0xFF))
				}else{
					*destrgba = *srcrgba
				}
				srcData += pixsize
				destData += pixsize
			}
		}
		WinApi.BitBlt(dc,x,y,w,h,BufferDC,0,0,WinApi.SRCCOPY)
		WinApi.SelectObject(BufferDC,OldBitmap)
		WinApi.DeleteDC(BufferDC)
		WinApi.DeleteObject(WinApi.HGDIOBJ(BufferBitmap))
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
	bmp.Free()
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

func (bmp *GBitmap)SetSize(w,h int)error  {
	screenDC := WinApi.GetDC(0)
	tmpdc := WinApi.CreateCompatibleDC(0)
	defer func() {
		WinApi.DeleteDC(tmpdc)
		WinApi.ReleaseDC(0,screenDC)
	}()

	oldbmphandle := bmp.fBufferBmp
	fbmpBuffer := make([]byte,1280)
	nbmpinfo := (*WinApi.BITMAPINFO)(unsafe.Pointer(&(fbmpBuffer[0])))
	if oldbmphandle != 0{
		nbmpinfo.BmiHeader = *bmp.BmpHeader()
	}else{
		nbmpinfo.BmiHeader.BiSize = infoHeaderLen
		nbmpinfo.BmiHeader.BiBitCount = 24
		nbmpinfo.BmiHeader.BiPlanes = 1
		nbmpinfo.BmiHeader.BiCompression = WinApi.BI_RGB
	}
	nbmpinfo.BmiHeader.BiHeight = int32(h)
	nbmpinfo.BmiHeader.BiWidth = int32(w)
	newpalate := WinApi.HPALETTE(0)
	//设置调色板
	if bmp.fPalate != 0{
		copy(fbmpBuffer[fileHeaderLen+infoHeaderLen:],bmp.fbmpBuffer[fileHeaderLen+infoHeaderLen:fileHeaderLen+infoHeaderLen+1024])
		newpalate = createPalate(nbmpinfo)
	}

	fBmpData := uintptr(0)

	fBufferBmp := WinApi.CreateDIBSection(tmpdc, nbmpinfo, WinApi.DIB_RGB_COLORS,
		&fBmpData, 0, 0)
	if fBufferBmp == 0 || fBufferBmp == 0{
		return errors.New("Create Compatible Bitmap Failed")
	}
	WinApi.SelectObject(tmpdc,uintptr(fBufferBmp))
	WinApi.PatBlt(tmpdc,0,0,w,h,WinApi.WHITENESS)
	if oldbmphandle != 0{
		WinApi.BitBlt(tmpdc,0,0,w,h,bmp.Canvas().GetHandle(),0,0,WinApi.SRCCOPY)
	}
	bmp.Free()

	bmp.fPalate = newpalate
	bmp.fbmpBuffer = fbmpBuffer
	bmp.fBufferBmp = fBufferBmp
	bmp.fBmpData = fBmpData
	return nil
}

func (bmp *GBitmap)LoadFromFile(filename string) error {
	f,err := os.Open(filename)
	if err != nil{
		return err
	}
	defer f.Close()
	err = bmp.Decode(f)
	var img image.Image
	if err != nil{
		//判断其他的格式
		var b [4]byte
		f.Seek(0,io.SeekStart)
		_,err = f.Read(b[:])
		if b[0]== 0x89 && b[1] == 'P' && b[2] == 'N' && b[3] == 'G'{
			f.Seek(0,io.SeekStart)
			img,err = png.Decode(f)
			if err == nil{
				return bmp.FromImage(img)
			}
		}else if b[0] == 0xFF && b[1] == 0xD8{
			f.Seek(-2,io.SeekEnd)
			f.Read(b[:2])
			if b[0] == 0xFF && b[1] == 0xD9{ //jpg
				f.Seek(0,io.SeekStart)
				img,err = jpeg.Decode(f)
				if err == nil{
					return bmp.FromImage(img)
				}
			}
		}

	}
	return err
}

func (bmp *GBitmap)from32imageData(w,h int, pix []byte)error  {
	bmp.Free()
	bmpInfo := bmp.BmpInfo()
	bmpInfo.BmiHeader.BiSize = infoHeaderLen
	bmpInfo.BmiHeader.BiBitCount = 32
	bmpInfo.BmiHeader.BiPlanes = 1
	bmpInfo.BmiHeader.BiCompression = WinApi.BI_RGB
	bmpInfo.BmiHeader.BiHeight = int32(-h)
	bmpInfo.BmiHeader.BiWidth = int32(w)
	bmpInfo.BmiHeader.BiSizeImage = uint32(bmpInfo.BmiHeader.BiHeight*bmpInfo.BmiHeader.BiWidth * 4)
	//填充内容数据
	err := bmp.newCompatibleBmp()
	if err != nil{
		return err
	}
	sourcdata := uintptr(unsafe.Pointer(&pix[0]))
	destdata := bmp.fBmpData
	for row := 0;row < int(-bmpInfo.BmiHeader.BiHeight);row++ {
		for col := 0; col < int(bmpInfo.BmiHeader.BiWidth); col++ {
			srgb := (*color.RGBA)(unsafe.Pointer(sourcdata))
			drgn := (*WinApi.RGBQUAD)(unsafe.Pointer(destdata))
			drgn.RgbRed = srgb.R
			drgn.RgbGreen = srgb.G
			drgn.RgbBlue = srgb.B
			drgn.RgbReserved = srgb.A
			sourcdata += RGBQUADSize
			destdata += RGBQUADSize
		}
	}
	return nil
}

func (bmp *GBitmap)fromYcbcr(img *image.YCbCr) error {
	return nil
}

func (bmp *GBitmap)FromImage(img image.Image)error  {
	switch v := img.(type) {
	case *image.RGBA:
		err := bmp.from32imageData(v.Rect.Max.X - v.Rect.Min.X,v.Rect.Max.Y - v.Rect.Min.Y, v.Pix)
		if err == nil{
			bmp.fcolorMode = img.ColorModel()
		}
		return err
	case *image.NRGBA:
		err := bmp.from32imageData(v.Rect.Max.X - v.Rect.Min.X,v.Rect.Max.Y - v.Rect.Min.Y,v.Pix)
		if err == nil{
			bmp.fcolorMode = img.ColorModel()
		}
		return err
	case *image.YCbCr:
		return bmp.fromYcbcr(v)
	case *image.Gray:

	case *image.Paletted:
		//构建8位位图
		bmp.Free()
		bmpInfo := bmp.BmpInfo()
		bmpInfo.BmiHeader.BiSize = infoHeaderLen
		bmpInfo.BmiHeader.BiBitCount = 8
		bmpInfo.BmiHeader.BiPlanes = 1
		bmpInfo.BmiHeader.BiCompression = WinApi.BI_RGB
		bmpInfo.BmiHeader.BiClrUsed = 256
		bmpInfo.BmiHeader.BiHeight = int32(v.Rect.Min.Y - v.Rect.Max.Y)
		bmpInfo.BmiHeader.BiWidth = int32(v.Rect.Max.X - v.Rect.Min.X)
		bmpInfo.BmiHeader.BiSizeImage = uint32(len(v.Pix))
		//填充调色板
		platdata := uintptr(unsafe.Pointer(&bmpInfo.BmiColors))
		for i := 0;i< len(v.Palette);i++{
			r,g,b,a := v.Palette[i].RGBA()
			rgba := (*WinApi.RGBQUAD)(unsafe.Pointer(platdata))
			rgba.RgbRed = byte(r)
			rgba.RgbGreen = byte(g)
			rgba.RgbBlue = byte(b)
			rgba.RgbReserved = byte(a)
			platdata+=RGBQUADSize
		}
		err := bmp.newCompatibleBmp()
		if err != nil{
			return err
		}
		scanline := bmpInfo.BmiHeader.BiWidth / 4
		if bmpInfo.BmiHeader.BiWidth % 4 > 0{
			scanline = (scanline + 1) * 4
		}
		h := int(-bmpInfo.BmiHeader.BiHeight)
		destdata := bmp.fBmpData
		for i := 0;i<h;i++{
			p := v.Pix[i*int(bmpInfo.BmiHeader.BiWidth) : i*int(bmpInfo.BmiHeader.BiWidth)+int(bmpInfo.BmiHeader.BiWidth)]
			WinApi.CopyMemory(unsafe.Pointer(destdata),unsafe.Pointer(&p[0]),int(bmpInfo.BmiHeader.BiWidth))
			destdata += uintptr(scanline)
		}
		bmp.fcolorMode = img.ColorModel()
		return nil
	}
	return errors.New("cannot Convert to BMP")
}

func (bmpobj *GBitmap)Encode(w io.Writer) error {
	return bmp.Encode(w,bmpobj)
}

func (bmpobj *GBitmap)EncodePNG(w io.Writer)error  {
	return png.Encode(w,bmpobj)
}

func (bmpobj *GBitmap)SaveToFile(fileName string)error  {
	if file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC, 0644); err == nil {
		defer file.Close()
		return bmpobj.Encode(file)
	}else{
		return err
	}
}

func (bmpobj *GBitmap)SaveToPngFile(fileName string)error  {
	if file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC, 0644); err == nil {
		defer file.Close()
		return bmpobj.EncodePNG(file)
	}else{
		return err
	}
}