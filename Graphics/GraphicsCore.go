package Graphics

import (
	"unsafe"
	"DxSoft/GVCL/WinApi"
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
)

var
(
	ClBtnFace,ClScrollBar,ClBackground,ClActiveCaption,ClInactiveCaption GColorValue
	ClMenu,ClWindow,ClWindowFrame,ClMenuText,ClWindowText,ClCaptionText GColorValue
	ClActiveBorder,ClInactiveBorder,ClAppWorkSpace,ClHighlight,ClHighlightText GColorValue
	ClBtnShadow,ClGrayText,ClBtnText,ClInactiveCaptionText,ClBtnHighlight GColorValue
	Cl3DDkShadow,Cl3DLight,ClInfoText,ClInfoBk,ClHotLight,ClGradientActiveCaption GColorValue
	ClGradientInactiveCaption,ClMenuHighlight,ClMenuBar GColorValue
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
	FontManager *GFontManager=new(GFontManager)
)

type GFontManager struct {
	fFonts map[string]*gFontData
}

func (fntMsg *GFontManager)createFontData(fntdata *gFontData)  {
	fntdata.createFont()
	fntdata.frefCount++
	fntMsg.fFonts[fntdata.getFontHashCode()] = fntdata
}

func (fntMng *GFontManager)AllocFontData(fntData *gFontData)  {
	if fntMng.fFonts == nil{
		fntMng.fFonts = make(map[string]*gFontData)
		fntMng.createFontData(fntData)
		return
	}
	hash := fntData.getFontHashCode()
	if v,ok := fntMng.fFonts[hash];ok && fntData.FontHandle != v.FontHandle{
		if fntData.FontHandle != 0{
			hash = fntData.getFontHashCode()
			if v,ok = fntMng.fFonts[hash];ok{
				v.frefCount--
				if v.frefCount == 0{
					delete(fntMng.fFonts,hash)
				}
			}
		}
		fntData.FontHandle = v.FontHandle
		v.frefCount++
	}else{
		fntMng.createFontData(fntData)
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
	frefCount int32
}

func (fdata *gFontData)getFontHashCode()string  {
	return fmt.Sprintf("%d&%d&%d&%d&%d&%d&%s",fdata.Color,fdata.Height,fdata.Escapement,fdata.Italic,fdata.Underline,fdata.StrikeOut,fdata.FontName)
}

func (fdata *gFontData)createFont()  {
	logfont := new(WinApi.GLOGFONT)
	logfont.Escapement = fdata.Escapement
	if fdata.FontName !=""{
		m := syscall.StringToUTF16(fdata.FontName)
		ml := len(m)
		if ml > 32{
			copy(logfont.FaceName[0:0],m[0:32])
		}else{
			copy(logfont.FaceName[0:0],m[0:ml])
		}
	}
	logfont.Height = fdata.Height
	logfont.Italic = fdata.Italic
	logfont.Underline = fdata.Underline
	logfont.StrikeOut = fdata.StrikeOut
	fdata.FontHandle = logfont.CreateFont()
}

type GFont struct {
	gFontData
	fsize int32
	fupcount int32
}

func (fnt *GFont)SetName(fontName string)  {
	if fnt.FontName !=fontName{
		fnt.FontName = fontName
		fnt.Change()
	}
}

func (fnt *GFont)Change()  {
	if fnt.fupcount == 0{
		FontManager.AllocFontData(&fnt.gFontData)
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