package WinApi

import (
	"syscall"
	"time"
	"unsafe"
)

type HMENU uintptr
type HINST uintptr
type ATOM uint16
type HICON uintptr
type HCURSOR uintptr
type HBRUSH uintptr
type HPEN uintptr
type HKL uintptr
type HDC uintptr
type HOOK uintptr
type LRESULT int
type HBITMAP uintptr
type HGDIOBJ uintptr
type HPALETTE	uintptr

func HiWord(L uint32) uint16 {
	return uint16(L >> 16)
}

func LoWord(L uint32)uint16  {
	return uint16(L)
}

type POINT struct {
	X, Y int32
}

func (pt *POINT)PtInRect(r *Rect)bool{
	/*ret,_,_ := syscall.Syscall(fnPtInRect,2,uintptr(unsafe.Pointer(r)),uintptr(unsafe.Pointer(pt)),0)
	return ret!=0*/
	return pt.X >= r.Left && pt.X <= r.Right && pt.Y>=r.Top && pt.Y <= r.Bottom
}

type Rect struct {
	Left,Top int32
	Right,Bottom int32
}

type GSize struct {
	CX,CY int32
}

func (rect *Rect)Width()int32{
	return rect.Right-rect.Left
}

func (rect *Rect)Height()  int32{
	return rect.Bottom-rect.Top
}

func (rect *Rect)PtInRect(pt *POINT) bool{
	/*ret,_,_ := syscall.Syscall(fnPtInRect,3,uintptr(unsafe.Pointer(rect)),uintptr(pt.X),uintptr(pt.Y))
	return ret!=0*/
	return pt.X >= rect.Left && pt.X <= rect.Right && pt.Y>=rect.Top && pt.Y <= rect.Bottom
}

func (rect *Rect)GetClientRect(hWnd syscall.Handle)bool  {
	ret,_,_ := syscall.Syscall(fnGetClientRect,2,uintptr(hWnd),uintptr(unsafe.Pointer(rect)),0)
	return ret!=0
}

func (rect *Rect)GetWindowRect(hWnd syscall.Handle)bool  {
	ret,_,_ := syscall.Syscall(fnGetWindowRect,2,uintptr(hWnd),uintptr(unsafe.Pointer(rect)),0)
	return ret!=0
}

func (rect *Rect)AdjustWindowRect(dwStyle uint32,bMenu bool)bool  {
	ret,_,_ := syscall.Syscall(fnAdjustWindowRect,3,uintptr(unsafe.Pointer(rect)),uintptr(dwStyle),uintptr(BoolToUint(bMenu)))
	return ret!=0
}

func (rect *Rect)AdjustWindowRectEx(dwStyle uint32,bMenu bool,dwExStyle uint32)bool  {
	ret,_,_ := syscall.Syscall6(fnAdjustWindowRect,4,uintptr(unsafe.Pointer(rect)),uintptr(dwStyle),uintptr(BoolToUint(bMenu)),
		uintptr(dwExStyle),0,0)
	return ret!=0
}

func (rc *Rect)InflateRect(dx,dy int)bool  {
	ret,_,_:=syscall.Syscall(fnInflateRect,3,uintptr(unsafe.Pointer(rc)),uintptr(dx),uintptr(dy))
	return ret!=0
}

func (rc *Rect)SetRectEmpty()bool  {
	ret,_,_:=syscall.Syscall(fnSetRectEmpty,1,uintptr(unsafe.Pointer(rc)),0,0)
	return ret!=0
}

func (rc *Rect)IsRectEmpty()bool{
	ret,_,_:=syscall.Syscall(fnIsRectEmpty,1,uintptr(unsafe.Pointer(rc)),0,0)
	return ret!=0
}

func (rc *Rect)OffsetRect(dx,dy int)bool{
	ret,_,_:=syscall.Syscall(fnOffsetRect,3,uintptr(unsafe.Pointer(rc)),uintptr(dx),uintptr(dy))
	return ret!=0
}

type tagWNDCLASSW struct {
	Style         uint32
	FnWndProc     uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINST
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
}

type GWndClass tagWNDCLASSW

type tagWNDCLASSEXW struct {
	CbSize        uint32
	Style         uint32
	FnWndProc     uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINST
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

type GWndClassEx tagWNDCLASSEXW

func (wndClass *GWndClass) GetClassInfo(lpClassName string) bool {
	var ret uintptr
	if lpClassName!=""{
		ret, _, _ = syscall.Syscall(fnGetClassInfoW, 3,
			uintptr(wndClass.HInstance),uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpClassName))),
			uintptr(unsafe.Pointer(wndClass)))
	}else{
		ret, _, _ = syscall.Syscall(fnGetClassInfoW, 3,
			uintptr(wndClass.HInstance), uintptr(unsafe.Pointer(wndClass.LpszClassName)),
			uintptr(unsafe.Pointer(wndClass)))
	}
	return ret != 0
}

func (wndClass *GWndClassEx) GetClassInfoEx(lpClassName string) bool {
	wndClass.CbSize = uint32(unsafe.Sizeof(*wndClass))
	ret, _, _ := syscall.Syscall(fnGetClassInfoExW, 3,
		uintptr(wndClass.HInstance), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpClassName))),
		uintptr(unsafe.Pointer(wndClass)))
	return ret != 0
}

func (wndClass *GWndClass)RegisterClass() ATOM {
	ret, _, _ := syscall.Syscall(fnRegisterClassW, 1,
		uintptr(unsafe.Pointer(wndClass)),
		0,
		0)
	return ATOM(ret)
}


func (wndClass *GWndClassEx)RegisterClassEx() ATOM {
	wndClass.CbSize = uint32(unsafe.Sizeof(*wndClass))
	ret, _, _ := syscall.Syscall(fnRegisterClassExW, 1,
		uintptr(unsafe.Pointer(wndClass)),
		0,
		0)
	return ATOM(ret)
}


type MSG struct {
	HWnd    syscall.Handle
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

func (msg *MSG)TranslateMessage() bool {
	ret, _, _ := syscall.Syscall(fnTranslateMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret != 0
}

func (msg *MSG)DispatchMessage() uintptr {
	ret, _, _ := syscall.Syscall(fnDispatchMessageW, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret
}

func (lpMsg *MSG)PeekMessage(hWnd syscall.Handle, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := syscall.Syscall6(fnPeekMessageW, 5,
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg),
		0)

	return ret != 0
}

func (lpmsg *MSG)CallMsgFilter(nCode int32)bool{
	ret,_,_:=syscall.Syscall(fnCallMsgFilterW,2,uintptr(unsafe.Pointer(lpmsg)),uintptr(nCode),0)
	return ret!=0
}

type GPaintStruct struct {
	Hdc HDC
	Erase bool
	PaintRect Rect
	Restore bool
	IncUpdate bool
	RgbReserved [32]byte
}

func (pstruct *GPaintStruct)BeginPaint(hWnd syscall.Handle)HDC{
	ret, _, _ := syscall.Syscall(fnBeginPaint, 2, uintptr(hWnd),uintptr(unsafe.Pointer(pstruct)),0)
	return  HDC(ret)
}

func (pstruct *GPaintStruct)EndPaint(hWnd syscall.Handle) bool{
	ret, _, _ := syscall.Syscall(fnEndPaint, 2, uintptr(hWnd),uintptr(unsafe.Pointer(pstruct)),0)
	return  ret!=0
}

type GdevicemodeW struct {
	DmDeviceName  [32]uint16//array[0..CCHDEVICENAME - 1] of WideChar;
	DmSpecVersion uint16
	DmDriverVersion uint16
	DmSize uint16
	DmDriverExtra uint16
	DmFields uint16
	DmOrientation int16
	DmPaperSize int16
	DmPaperLength int16
	DmPaperWidth int16
	DmScale int16
	DmCopies int16
	DmDefaultSource int16
	DmPrintQuality int16
	DmColor int16
	DmDuplex int16
	DmYResolution int16
	DmTTOption int16
	DmCollate int16
	DmFormName [32]uint16
	DmLogPixels uint16
	DmBitsPerPel uint32
	DmPelsWidth uint32
	DmPelsHeight uint32
	DmDisplayFlags uint32
	DmDisplayFrequency uint32
	DmICMMethod uint32
	DmICMIntent uint32
	DmMediaType uint32
	DmDitherType uint32
	DmICCManufacturer uint32
	DmICCModel uint32
	DmPanningWidth uint32
	DmPanningHeight uint32
}

type GDrawTextParams struct {
	CbSize uint32
	TabLength int32
	LeftMargin int32
	RightMargin int32
	LengthDrawn uint32
}

type GLOGFONT struct {
	Height int32
	Width  int32
	Escapement int32
	Orientation int32
	Weight int32
	Italic byte
	Underline byte
	StrikeOut byte
	CharSet byte
	OutPrecision byte
	ClipPrecision byte
	Quality byte
	PitchAndFamily byte
	FaceName [32]uint16
}

type GMenuItemInfo struct{
	CbSize uint32
	FMask  uint32
	FType  uint32
	FState uint32
	WID    uint32
	HSubMenu HMENU
	HbmpChecked HBITMAP
	HbmpUnchecked HBITMAP
	DwItemData uintptr
	DwTypeData uintptr
	Cch  uint32
	HbmpItem HBITMAP
}

func (logFont *GLOGFONT)CreateFont()syscall.Handle  {
	ret,_,_ := syscall.Syscall(fnCreateFontIndirectW,1,uintptr(unsafe.Pointer(logFont)),0,0)
	return syscall.Handle(ret)
}

type GLogBrush struct {
	Style uint32
	Color uint32
	Hatch uintptr
}

func (logbrush *GLogBrush)CreateBrush()HBRUSH  {
	ret,_,_ := syscall.Syscall(fnCreateBrushIndirect,1,uintptr(unsafe.Pointer(logbrush)),0,0)
	return HBRUSH(ret)
}

type GLOGPEN struct {
	Style uint32
	Width POINT
	Color uint32
}

func (logpen *GLOGPEN)CreatePen()HPEN  {
	ret,_,_ := syscall.Syscall(fnCreatePenIndirect,1,uintptr(unsafe.Pointer(logpen)),0,0)
	return HPEN(ret)
}

func PChar(str string)*uint16  {
	return syscall.StringToUTF16Ptr(str)
}

func PCharUintptr(str string)uintptr  {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

type tagNMHDR struct{
	HwndFrom		syscall.Handle
	IdFrom			uintptr
	NMCode			int
}

type GNMHDR tagNMHDR


type CIEXYZ struct {
	CiexyzX, CiexyzY, CiexyzZ int32 // FXPT2DOT30
}

type CIEXYZTRIPLE struct {
	CiexyzRed, CiexyzGreen, CiexyzBlue CIEXYZ
}

type BITMAPINFOHEADER struct {
	BiSize		uint32
	BiWidth		int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

type BITMAPV4HEADER struct {
	BITMAPINFOHEADER
	BV4RedMask    uint32
	BV4GreenMask  uint32
	BV4BlueMask   uint32
	BV4AlphaMask  uint32
	BV4CSType     uint32
	BV4Endpoints  CIEXYZTRIPLE
	BV4GammaRed   uint32
	BV4GammaGreen uint32
	BV4GammaBlue  uint32
}

type BITMAPV5HEADER struct {
	BITMAPV4HEADER
	BV5Intent      uint32
	BV5ProfileData uint32
	BV5ProfileSize uint32
	BV5Reserved    uint32
}

type RGBQUAD struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

func (rgb RGBQUAD)RGBA() (r, g, b, a uint32)  {
	return uint32(rgb.RgbRed)<< 8,uint32(rgb.RgbGreen)<< 8,uint32(rgb.RgbBlue)<< 8,uint32(rgb.RgbReserved)<< 8
}

type PALETTEENTRY struct {
	Red			byte
	Green		byte
	Blue		byte
	Flags		byte
}

//为了兼容go自带的image接口，然后做相关的encode，所以，必须将所获分量全部左移8位
func (rgb PALETTEENTRY)RGBA() (r, g, b, a uint32)  {
	return uint32(rgb.Red)<< 8,uint32(rgb.Green)<< 8,uint32(rgb.Blue)<< 8,uint32(rgb.Flags)<< 8
}

type LOGPALETTE struct {
	Version		uint16
	NumEntries	uint16
	PalEntry	[256]PALETTEENTRY
}


type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors RGBQUAD
}

type BITMAPINFO8 struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [256]RGBQUAD
}

type BITMAP struct {
	BmType       int32
	BmWidth      int32
	BmHeight     int32
	BmWidthBytes int32
	BmPlanes     uint16
	BmBitsPixel  uint16
	BmBits       unsafe.Pointer
}

type DIBSECTION struct {
	DsBm        BITMAP
	DsBmih      BITMAPINFOHEADER
	DsBitfields [3]uint32
	DshSection  syscall.Handle
	DsOffset    uint32
}


type BLENDFUNCTION struct {
	BlendOp			byte
	BlendFlags		byte
	SourceConstantAlpha	byte
	AlphaFormat		byte
}

type GSystemTime struct {
	Year		uint16
	Month		uint16
	DayOfWeek	uint16
	Day			uint16
	Hour		uint16
	Minute		uint16
	Second		uint16
	Milliseconds uint16
}

func (sysTime GSystemTime)Time()time.Time  {
	return time.Date(int(sysTime.Year),time.Month(sysTime.Month),int(sysTime.Day),int(sysTime.Hour),int(sysTime.Minute),
		int(sysTime.Second),int(time.Duration(sysTime.Milliseconds) * time.Millisecond),time.Local)
}

func (sysTime *GSystemTime)From(t time.Time)  {
	year,month,day := t.Date()
	sysTime.Year = uint16(year)
	sysTime.Month = uint16(month)
	sysTime.Day = uint16(day)
	year,m,day := t.Clock()
	sysTime.Hour = uint16(year)
	sysTime.Minute = uint16(m)
	sysTime.Second = uint16(day)
	sysTime.Milliseconds = uint16(time.Duration(t.Nanosecond()) / time.Millisecond)
}