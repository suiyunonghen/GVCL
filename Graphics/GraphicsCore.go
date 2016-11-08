package Graphics

import (
	"unsafe"
	"suiyunonghen/GVCL/WinApi"
	"syscall"
	"fmt"
)

type GColorValue uint32

func (cv GColorValue) R() byte {
	return (*GColor)(unsafe.Pointer(&cv)).R
}

func (cv GColorValue) G() byte {
	return (*GColor)(unsafe.Pointer(&cv)).G
}

func (cv GColorValue) B() byte {
	return (*GColor)(unsafe.Pointer(&cv)).B
}

func (cv GColorValue) A() byte {
	return (*GColor)(unsafe.Pointer(&cv)).A
}

type GColor struct {
	R byte
	G byte
	B byte
	A byte
}

const (
	SystemColor GColorValue = 0xFF000000
	SCROLLBAR int32 = 0
	BACKGROUND int32 = 1
	ACTIVECAPTION int32 = 2
	INACTIVECAPTION int32 = 3
	MENU int32 = 4
	WINDOW int32=5
	WINDOWFRAME int32 = 6
	MENUTEXT int32 = 7
	WINDOWTEXT int32 = 8
	CAPTIONTEXT int32 = 9
	ACTIVEBORDER int32 =10
	INACTIVEBORDER int32 = 11
	APPWORKSPACE int32 = 12
	HIGHLIGHT int32 = 13
	HIGHLIGHTTEXT int32 = 14
	BTNFACE int32 = 15
	BTNSHADOW int32 = 16
	GRAYTEXT = 17
	BTNTEXT = 18
	INACTIVECAPTIONTEXT = 19
	BTNHIGHLIGHT = 20
	c3DDKSHADOW = 21
	c3DLIGHT = 22
	cINFOTEXT = 23
	cINFOBK = 24
	cHOTLIGHT = 26
	cGRADIENTACTIVECAPTION = 27
	cGRADIENTINACTIVECAPTION = 28
	cMENUHILIGHT = 29
	cMENUBAR = 30

	ClRed   GColorValue = 0xFF
	ClWhite GColorValue = 0xFFFFFF
	ClBlack GColorValue = 0x000000
	ClGray GColorValue = 0x808080
	ClDkGray GColorValue = 0x808080
	ClLtGray GColorValue = 0xC0C0C0
	ClDarkGray GColorValue = 0xA9A9A9
)


var
(
	ClBtnFace,ClScrollBar,ClBackground,ClActiveCaption,ClInactiveCaption GColorValue
	ClMenu,ClWindow,ClWindowFrame,ClMenuText,ClWindowText,ClCaptionText GColorValue
	ClActiveBorder,ClInactiveBorder,ClAppWorkSpace,ClHighlight,ClHighlightText GColorValue
	ClBtnShadow,ClGrayText,ClBtnText,ClInactiveCaptionText,ClBtnHighlight GColorValue
	Cl3DDkShadow,Cl3DLight,ClInfoText,ClInfoBk,ClHotLight,ClGradientActiveCaption GColorValue
	ClGradientInactiveCaption,ClMenuHighlight,ClMenuBar GColorValue

	WhiteBrush WinApi.HBRUSH
	EmptyBrush WinApi.HBRUSH
	BlackBrush WinApi.HBRUSH
	LtGrayBrush WinApi.HBRUSH
	GrayBrush WinApi.HBRUSH
	DkGrayBrush WinApi.HBRUSH

	WhitePen WinApi.HPEN
	BlackPen WinApi.HPEN
	EmptyPen WinApi.HPEN
)
func init()  {
	ClBtnFace = GColorValue(WinApi.GetSysColor(BTNFACE))
	ClScrollBar= GColorValue(WinApi.GetSysColor(SCROLLBAR))
	ClBackground= GColorValue(WinApi.GetSysColor(BACKGROUND))
	ClActiveCaption= GColorValue(WinApi.GetSysColor(ACTIVECAPTION))
	ClInactiveCaption= GColorValue(WinApi.GetSysColor(INACTIVECAPTION))
	ClMenu= GColorValue(WinApi.GetSysColor(MENU))
	ClWindow= GColorValue(WinApi.GetSysColor(WINDOW))
	ClWindowFrame= GColorValue(WinApi.GetSysColor(WINDOWFRAME))
	ClMenuText= GColorValue(WinApi.GetSysColor(MENUTEXT))
	ClWindowText= GColorValue(WinApi.GetSysColor(WINDOWTEXT))
	ClCaptionText= GColorValue(WinApi.GetSysColor(CAPTIONTEXT))
	ClActiveBorder = GColorValue(WinApi.GetSysColor(ACTIVEBORDER))
	ClInactiveBorder = GColorValue(WinApi.GetSysColor(INACTIVEBORDER))
	ClAppWorkSpace = GColorValue(WinApi.GetSysColor(APPWORKSPACE))
	ClHighlight = GColorValue(WinApi.GetSysColor(HIGHLIGHT))
	ClHighlightText = GColorValue(WinApi.GetSysColor(HIGHLIGHTTEXT))
	ClBtnShadow = GColorValue(WinApi.GetSysColor(BTNSHADOW))
	ClGrayText = GColorValue(WinApi.GetSysColor(GRAYTEXT))
	ClBtnText = GColorValue(WinApi.GetSysColor(BTNTEXT))
	ClInactiveCaptionText = GColorValue(WinApi.GetSysColor(INACTIVECAPTIONTEXT))
	ClBtnHighlight = GColorValue(WinApi.GetSysColor(BTNHIGHLIGHT))
	Cl3DDkShadow  = GColorValue(WinApi.GetSysColor(c3DDKSHADOW))
	Cl3DLight  = GColorValue(WinApi.GetSysColor(c3DLIGHT))
	ClInfoText  = GColorValue(WinApi.GetSysColor(cINFOTEXT))
	ClInfoBk  = GColorValue(WinApi.GetSysColor(cINFOBK))
	ClHotLight  = GColorValue(WinApi.GetSysColor(cHOTLIGHT))
	ClGradientActiveCaption  = GColorValue(WinApi.GetSysColor(cGRADIENTACTIVECAPTION))
	ClGradientInactiveCaption  = GColorValue(WinApi.GetSysColor(cGRADIENTINACTIVECAPTION))
	ClMenuHighlight  = GColorValue(WinApi.GetSysColor(cMENUHILIGHT))
	ClMenuBar  = GColorValue(WinApi.GetSysColor(cMENUBAR))


	WhiteBrush = WinApi.HBRUSH(WinApi.GetStockObject(WinApi.WHITE_BRUSH))
	EmptyBrush  = WinApi.HBRUSH(WinApi.GetStockObject(WinApi.HOLLOW_BRUSH))
	BlackBrush  = WinApi.HBRUSH(WinApi.GetStockObject(WinApi.BLACK_BRUSH))
	LtGrayBrush  = WinApi.HBRUSH(WinApi.GetStockObject(WinApi.LTGRAY_BRUSH))
	GrayBrush  = WinApi.HBRUSH(WinApi.GetStockObject(WinApi.GRAY_BRUSH))
	DkGrayBrush  = WinApi.HBRUSH(WinApi.GetStockObject(WinApi.DKGRAY_BRUSH))

	WhitePen = WinApi.HPEN(WinApi.GetStockObject(WinApi.WHITE_PEN))
	BlackPen = WinApi.HPEN(WinApi.GetStockObject(WinApi.BLACK_PEN))
	EmptyPen = WinApi.HPEN(WinApi.GetStockObject(WinApi.NULL_PEN))
}

func RGB(R, G, B byte) (ret GColorValue) {
	ret = 0
	gcolor := Uint32ToGColor(uint32(ret))
	gcolor.R = R
	gcolor.G = G
	gcolor.B = B
	return
}

func (c *GColor) GColor2Uint32() uint32 {
	return *(*uint32)(unsafe.Pointer(c))
}

func (c *GColor) GColor2ColorValue() GColorValue {
	return GColorValue(*(*uint32)(unsafe.Pointer(c)))
}

func Uint32ToGColor(ClValue uint32) *GColor {
	return (*GColor)(unsafe.Pointer(&ClValue))
}

var(
	gdiManager *gGDIManager=new(gGDIManager)
)

type gGDIManager struct {
	fFonts map[string]*gFontData
	fBrushs map[string]*gBrushData
	fPens map[string]*gPenData
}


func (BrushMng *gGDIManager)createPenData(pdata *gPenData)  {
	pdata.createPen()
	pdata.frefCount++
	pdata.fHashCode = pdata.getHashCode()
	BrushMng.fPens[pdata.fHashCode] = pdata
}

func (fntMng *gGDIManager)AllocPenData(pdata *gPenData)  {
	if fntMng.fPens == nil{
		fntMng.fPens = make(map[string]*gPenData)
		fntMng.createPenData(pdata)
		return
	}
	hash := pdata.getHashCode()
	if hash != pdata.fHashCode{ //有变动
		if pdata.fHashCode != ""{//先移除以前的
			if v,ok := fntMng.fPens[pdata.fHashCode];ok{
				v.frefCount--
				pdata.fHashCode = ""
				if v.frefCount == 0{
					if v.Handle != BlackPen && v.Handle != EmptyPen && v.Handle != WhitePen{
						WinApi.DeleteObject(uintptr(v.Handle))
					}
					delete(fntMng.fPens,hash)
				}
			}
		}
		if v,ok := fntMng.fPens[hash];ok{ //新的
			v.frefCount++
			pdata.Handle = v.Handle
			pdata.fHashCode = hash
		}else{
			fntMng.createPenData(pdata)
		}
	}
}

func (fntMsg *gGDIManager)createFontData(fntdata *gFontData)  {
	fntdata.createFont()
	fntdata.frefCount++
	fntdata.fHashCode = fntdata.getFontHashCode()
	fntMsg.fFonts[fntdata.fHashCode] = fntdata
}

func (BrushMng *gGDIManager)createBrushData(brushdata *gBrushData)  {
	brushdata.createBrush()
	brushdata.frefCount++
	brushdata.fHashCode = brushdata.getHashCode()
	BrushMng.fBrushs[brushdata.fHashCode] = brushdata
}

func (fntMng *gGDIManager)AllocFontData(fntData *gFontData)  {
	if fntMng.fFonts == nil{
		fntMng.fFonts = make(map[string]*gFontData)
		fntMng.createFontData(fntData)
		return
	}
	hash := fntData.getFontHashCode()
	if hash != fntData.fHashCode{ //有变动
		if fntData.fHashCode != ""{//先移除以前的
			if v,ok := fntMng.fFonts[fntData.fHashCode];ok{
				v.frefCount--
				fntData.fHashCode = ""
				if v.frefCount == 0{
					WinApi.DeleteObject(uintptr(v.FontHandle))
					delete(fntMng.fFonts,hash)
				}
			}
		}
		if v,ok := fntMng.fFonts[hash];ok{ //新的
			v.frefCount++
			fntData.FontHandle = v.FontHandle
			fntData.fHashCode = hash
		}else{
			fntMng.createFontData(fntData)
		}
	}
}

func (BrushMng *gGDIManager)AllocBrushData(brushData *gBrushData)  {
	if BrushMng.fBrushs == nil{
		BrushMng.fBrushs = make(map[string]*gBrushData)
		BrushMng.createBrushData(brushData)
		return
	}
	hash := brushData.getHashCode()
	if hash != brushData.fHashCode{ //有变动
		if brushData.fHashCode != ""{//先移除以前的
			if v,ok := BrushMng.fBrushs[brushData.fHashCode];ok{
				v.frefCount--
				brushData.fHashCode = ""
				if v.frefCount == 0{
					if v.Handle != EmptyBrush && v.Handle != WhiteBrush &&
						v.Handle != GrayBrush && v.Handle != DkGrayBrush &&
						v.Handle != BlackBrush && v.Handle != LtGrayBrush{
						WinApi.DeleteObject(uintptr(v.Handle))
					}
					delete(BrushMng.fBrushs,hash)
				}
			}
		}
		if v,ok := BrushMng.fBrushs[hash];ok{ //新的
			v.frefCount++
			brushData.Handle = v.Handle
			brushData.fHashCode = hash
		}else{
			BrushMng.createBrushData(brushData)
		}
	}
}

type gFontData struct{
	FontHandle syscall.Handle
	Color GColorValue
	Height int32
	Escapement int32 //角度*10
	Italic byte
	Underline byte
	StrikeOut byte
	FontName string
	Weight  int32
	frefCount int32
	fHashCode string
}

//BrushStyle
const(
	BSSolid=iota
	BSClear
	BSHorizontal
	BSVertical
	BSFDiagonal
	BSBDiagonal
	BSCross
	BSDiagCross
)

type gBrushData struct{
	Handle WinApi.HBRUSH
	Color  GColorValue
	BrushStyle uint8
	frefCount int32
	fHashCode string
}

//PenStyle
const (
	PSSolid=iota
	PSDash
	PSDot
	PSDashDot
	PSDashDotDot
	PSClear
	PSInsideFrame
	PSUserStyle
	PSAlternate
)

//PenMode
const (
	PMBlack=iota
	PMWhite
	PMNop
	PMNot
	PMCopy
	PMNotCopy
	PMMergePenNot
	PMMaskPenNot
	PMMergeNotPen
	PMMaskNotPen
	PMMerge
	PMNotMerge
	PMMask
	PMNotMask
	PMXor
	PMNotXor
)
type gPenData struct {
	Handle WinApi.HPEN
	Color  GColorValue
	Width  int32
	Style  uint8
	frefCount int32
	fHashCode string
}

func (pdata *gPenData)getHashCode()string{
	return fmt.Sprintf("%d&%d&%d",pdata.Color,pdata.Width,pdata.Style)
}

func (pdata *gPenData)createPen()  {
	if pdata.Style == PSSolid && pdata.Width == 1{
		if pdata.Color == ClWhite{
			pdata.Handle = WhitePen
		}else if pdata.Color == ClBlack{
			pdata.Handle = BlackPen
		}else{
			pdata.Handle = 0
		}
	}else if pdata.Style == PSClear{
		pdata.Handle = EmptyPen
	}else {
		pdata.Handle = 0
	}
	if pdata.Handle != 0{
		return
	}
	logPen := new(WinApi.GLOGPEN)
	logPen.Color = uint32(pdata.Color)
	logPen.Width.X = pdata.Width
	logPen.Style = uint32(pdata.Style)
	pdata.Handle = logPen.CreatePen()
}

func (bdata *gBrushData)getHashCode()string{
   return fmt.Sprintf("%d&%d",bdata.Color,bdata.BrushStyle)
}

func (bdata *gBrushData)createBrush()  {
	if bdata.BrushStyle == BSSolid{
		switch bdata.Color {
		case ClWhite:
			bdata.Handle = WhiteBrush
		case ClGray:
			bdata.Handle = GrayBrush
		case ClDarkGray:
			bdata.Handle = DkGrayBrush
		case ClBlack:
			bdata.Handle = BlackBrush
		case ClLtGray:
			bdata.Handle = LtGrayBrush
		default:
			bdata.Handle = 0
		}

	}else if bdata.BrushStyle == BSClear{
		bdata.Handle = EmptyBrush
	}else{
		bdata.Handle = 0
	}
	if bdata.Handle !=0{
		return
	}
	logBrush := new(WinApi.GLogBrush)
	logBrush.Color = uint32(bdata.Color)
	switch bdata.BrushStyle {
	case BSClear:
		logBrush.Style = WinApi.BS_HOLLOW
	case BSSolid:
		logBrush.Style = WinApi.BS_SOLID
	default:
		logBrush.Style = WinApi.BS_HATCHED
		logBrush.Hatch = uintptr(bdata.BrushStyle - BSHorizontal)
	}
	bdata.Handle = logBrush.CreateBrush()
}

func (fdata *gFontData)getFontHashCode()string  {
	return fmt.Sprintf("%d&%d&%d&%d&%d&%d&%s%d",fdata.Color,fdata.Height,fdata.Escapement,fdata.Italic,fdata.Underline,fdata.StrikeOut,fdata.FontName,fdata.Weight)
}

func (fdata *gFontData)createFont()  {
	logfont := new(WinApi.GLOGFONT)
	logfont.Escapement = fdata.Escapement
	if fdata.FontName !=""{
		copy(logfont.FaceName[0:32],syscall.StringToUTF16(fdata.FontName))
	}
	logfont.Height = fdata.Height
	logfont.Italic = fdata.Italic
	logfont.Underline = fdata.Underline
	logfont.StrikeOut = fdata.StrikeOut
	logfont.Weight = fdata.Weight
	fdata.FontHandle = logfont.CreateFont()
}

type GFont struct {
	gFontData
	fsize int32
	fupcount int32
}

func (fnt *GFont)Destroy()  {
	if fnt.FontHandle!=0{
		if v,ok := gdiManager.fFonts[fnt.fHashCode];ok{
			v.frefCount--
			if v.frefCount == 0{
				WinApi.DeleteObject(uintptr(v.FontHandle))
				delete(gdiManager.fFonts,fnt.fHashCode)
			}
		}
		fnt.FontHandle = 0
	}
	fnt.fHashCode = ""
}

func (fnt *GFont)SetName(fontName string)  {
	if fnt.FontName !=fontName{
		fnt.FontName = fontName
		fnt.Change()
	}
}

func (fnt *GFont)Change()  {
	if fnt.fupcount == 0{
		gdiManager.AllocFontData(&fnt.gFontData)
	}
}

func (fnt *GFont)SetSize(sz int32)  {
	if fnt.fsize != sz{
		fnt.fsize = sz
		fnt.Height = -WinApi.MulDiv(sz, WinApi.ScreenLogPixels, 72)
		fnt.Change()
	}
}

func (fnt *GFont)BeginUpdate()  {
	fnt.fupcount ++
}

func (fnt *GFont)EndUpdate()  {
	fnt.fupcount--
	if fnt.fupcount<=0{
		fnt.fupcount = 0
		fnt.Change()
	}
}

func (fnt *GFont)SetBold(v bool)  {
	if fnt.Weight == WinApi.FW_BOLD != v{
		if v{
			fnt.Weight = WinApi.FW_BOLD
		}else{
			fnt.Weight = WinApi.FW_NORMAL
		}
		fnt.Change()
	}
}

func (fnt *GFont)Bold()bool  {
	return fnt.Weight == WinApi.FW_BOLD
}

type GBrush struct {
	gBrushData
	fupcount int32
}


func (brush *GBrush)BeginUpdate()  {
	brush.fupcount ++
}

func (brush *GBrush)Change()  {
	if brush.fupcount == 0{
		gdiManager.AllocBrushData(&brush.gBrushData)
	}
}

func (brush *GBrush)EndUpdate()  {
	brush.fupcount--
	if brush.fupcount<=0{
		brush.fupcount = 0
		brush.Change()
	}
}

func (brush *GBrush)Destroy()  {
	if brush.Handle!=0{
		if v,ok := gdiManager.fBrushs[brush.fHashCode];ok{
			v.frefCount--
			if v.frefCount == 0{
				WinApi.DeleteObject(uintptr(v.Handle))
				delete(gdiManager.fBrushs,brush.fHashCode)
			}
		}
		brush.Handle = 0
	}
	brush.fHashCode = ""
}

func (fnt *GFont)Assign(newfnd *GFont)  {
	fnt.BeginUpdate()
	defer fnt.EndUpdate()
	fnt.Color = newfnd.Color
	fnt.Escapement = newfnd.Escapement
	fnt.FontName = newfnd.FontName
	fnt.fsize = newfnd.fsize
	fnt.Height = newfnd.Height
	fnt.Weight = newfnd.Weight
	fnt.Italic = newfnd.Italic
	fnt.StrikeOut = newfnd.StrikeOut
	fnt.Underline = newfnd.Underline
}

func (pen *GPen)Assign(newpen *GPen)  {
	pen.BeginUpdate()
	defer pen.EndUpdate()
	pen.Color = newpen.Color
	pen.PenMode = newpen.PenMode
	pen.Style = newpen.Style
	pen.Width = newpen.Width
}

func (brush *GBrush)Assign(newBrush *GBrush)  {
	brush.BeginUpdate()
	defer brush.EndUpdate()
	brush.Color = newBrush.Color
	brush.BrushStyle = newBrush.BrushStyle
}


type GPen struct {
	gPenData
	fupcount int32
	PenMode uint8
}

func (pen *GPen)BeginUpdate()  {
	pen.fupcount ++
}

func (pen *GPen)Change()  {
	if pen.fupcount == 0{
		gdiManager.AllocPenData(&pen.gPenData)
	}
}

func (pen *GPen)EndUpdate()  {
	pen.fupcount--
	if pen.fupcount<=0{
		pen.fupcount = 0
		pen.Change()
	}
}

func (pen *GPen)Destroy()  {
	if pen.Handle!=0{
		if v,ok := gdiManager.fPens[pen.fHashCode];ok{
			v.frefCount--
			if v.frefCount == 0{
				WinApi.DeleteObject(uintptr(v.Handle))
				delete(gdiManager.fPens,pen.fHashCode)
			}
		}
		pen.Handle = 0
	}
	pen.fHashCode = ""
}


func NewBrush()*GBrush{
	brush := new(GBrush)
	brush.Color = ClWhite
	brush.BrushStyle = BSSolid
	brush.Change()
	return brush
}

func NewFont()*GFont{
	fnt := new(GFont)
	fnt.Color = ClWhite
	fnt.FontName = "宋体"
	fnt.SetSize(9)
	return fnt
}


func NewPen()*GPen{
	pen := new(GPen)
	pen.Color = ClBlack
	pen.Style = PSSolid
	pen.Width = 1
	pen.Change()
	return pen
}