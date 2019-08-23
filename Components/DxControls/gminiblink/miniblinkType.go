package gminiblink

import (
	"github.com/suiyunonghen/GVCL/WinApi"
	"unsafe"
)

type WkeRect struct {
	X,Y,W,H int32
}

func (r WkeRect)ToWindowRect()WinApi.Rect  {
	return WinApi.Rect{r.X,r.Y,r.X+r.W,r.Y+r.H}
}

func (r WkeRect)Right()int32  {
	return r.X + r.W
}

func (r WkeRect)Bottom()int32  {
	return r.Y + r.H
}

func (r WkeRect)RectContactPt(p WkePoint)bool  {
	return p.X >= r.X && p.X <= r.Right() && p.Y>=r.Y && p.Y <= r.Bottom()
}

func WinRect2WkeRect(r *WinApi.Rect)WkeRect  {
	return WkeRect{r.Left,r.Top,r.Width(),r.Height()}
}

type WkePoint struct {
	X,Y int32
}

func (p *WkePoint)ToWinPoint()*WinApi.POINT  {
	return (*WinApi.POINT)(unsafe.Pointer(p))
}



type WKEMouseFlags byte
const(
	WKE_LBUTTON WKEMouseFlags =0x01
	WKE_RBUTTON WKEMouseFlags=0x02
	WKE_SHIFT WKEMouseFlags=0x04
	WKE_CONTROL WKEMouseFlags=0x08
	WKE_MBUTTON WKEMouseFlags=0x10
)

type WKEKeyFlags uint32
const(
	WKE_EXTENDED WKEKeyFlags=0x0100
	WKE_REPEAT = 0x4000
)

type WKEMouseMsg uint32
const(
	WKE_MSG_MOUSEMOVE WKEMouseMsg = 0x0200
	WKE_MSG_LBUTTONDOWN = 0x0201
	WKE_MSG_LBUTTONUP = 0x0202
	WKE_MSG_LBUTTONDBLCLK = 0x0203
	WKE_MSG_RBUTTONDOWN = 0x0204
	WKE_MSG_RBUTTONUP = 0x0205
	WKE_MSG_RBUTTONDBLCLK = 0x0206
	WKE_MSG_MBUTTONDOWN = 0x0207
	WKE_MSG_MBUTTONUP = 0x0208
	WKE_MSG_MBUTTONDBLCLK = 0x0209
	WKE_MSG_MOUSEWHEEL = 0x020A
)

//通过wkeGetCursorInfoType获得光标信息
type WkeCursorInfoType byte
const(
	WkeCursorInfoPointer WkeCursorInfoType = iota
    WkeCursorInfoCross
    WkeCursorInfoHand
    WkeCursorInfoIBeam
    WkeCursorInfoWait
    WkeCursorInfoHelp
    WkeCursorInfoEastResize
    WkeCursorInfoNorthResize
    WkeCursorInfoNorthEastResize
    WkeCursorInfoNorthWestResize
    WkeCursorInfoSouthResize
    WkeCursorInfoSouthEastResize
    WkeCursorInfoSouthWestResize
    WkeCursorInfoWestResize
    WkeCursorInfoNorthSouthResize
    WkeCursorInfoEastWestResize
    WkeCursorInfoNorthEastSouthWestResize
    WkeCursorInfoNorthWestSouthEastResize
    WkeCursorInfoColumnResize
    WkeCursorInfoRowResize
    WkeCursorInfoMiddlePanning
    WkeCursorInfoEastPanning
    WkeCursorInfoNorthPanning
    WkeCursorInfoNorthEastPanning
    WkeCursorInfoNorthWestPanning
    WkeCursorInfoSouthPanning
    WkeCursorInfoSouthEastPanning
    WkeCursorInfoSouthWestPanning
    WkeCursorInfoWestPanning
    WkeCursorInfoMove
    WkeCursorInfoVerticalText
    WkeCursorInfoCell
    WkeCursorInfoContextMenu
    WkeCursorInfoAlias
    WkeCursorInfoProgress
    WkeCursorInfoNoDrop
    WkeCursorInfoCopy
    WkeCursorInfoNone
    WkeCursorInfoNotAllowed
    WkeCursorInfoZoomIn
    WkeCursorInfoZoomOut
    WkeCursorInfoGrab
    WkeCursorInfoGrabbing
    WkeCursorInfoCustom
)

type (
	JSExecState uintptr
	JSValue int64
	PJSValue *JSValue
	WkeWebView uintptr
	WkeString uintptr
	WkeUrlRequestCallbacks struct{
		willRedirectCallback uintptr
		didReceiveResponseCallback uintptr
		didReceiveDataCallback uintptr
		didFailCallback uintptr
		didFinishLoadingCallback uintptr
	}

	WkeWebFrameHandle uintptr
	WkeFrameHwnd uintptr
	WkeCookieCommand byte
	WkeNavigationType byte
	WkeConsoleLevel byte
	WkeWindowType byte
	WkeProxyType byte
)

const (
	wkeCookieCommandClearAllCookies WkeCookieCommand = iota
	wkeCookieCommandClearSessionCookies
	wkeCookieCommandFlushCookiesToFile
	wkeCookieCommandReloadCookiesFromFile
)

const (WKE_NAVIGATION_TYPE_LINKCLICK WkeNavigationType = iota    //点击a标签触发
    WKE_NAVIGATION_TYPE_FORMSUBMITTE        //点击form触发
    WKE_NAVIGATION_TYPE_BACKFORWARD        //前进后退触发
    WKE_NAVIGATION_TYPE_RELOAD             //重新加载触发
    WKE_NAVIGATION_TYPE_FORMRESUBMITT
    WKE_NAVIGATION_TYPE_OTHER
)

const(
	WkeLevelDebug WkeConsoleLevel= 4
	WkeLevelLog WkeConsoleLevel= 1
	WkeLevelInfo WkeConsoleLevel= 5
	WkeLevelWarning WkeConsoleLevel= 2
	WkeLevelError WkeConsoleLevel= 3
	WkeLevelRevokedError WkeConsoleLevel= 6
	WkeLevelLast WkeConsoleLevel= WkeLevelInfo
)

const (
	WKE_WINDOW_TYPE_POPUP WkeWindowType=iota
	WKE_WINDOW_TYPE_TRANSPARENT
	WKE_WINDOW_TYPE_CONTROL
)

const (
	WKE_PROXY_NONE WkeProxyType=iota
	WKE_PROXY_HTTP
	WKE_PROXY_SOCKS4
	WKE_PROXY_SOCKS4A
	WKE_PROXY_SOCKS5
	WKE_PROXY_SOCKS5HOSTNAME
)

type (
	WkeProxy struct {
		Type WkeProxyType
		HostName [100]byte
		Port uint16
		UserName [50]byte
		PassWord [50]byte
	}

	WkeSettings struct {
		Proxy WkeProxy
		Mask	uint32
	}

	WkeViewSettings struct {
		Size 	int32
		BackColor	uint32
	}

	WkeMemBuf struct {
		Size	int32
		Data	uintptr
		Len		uint
	}

	WkeWindowFeatures struct {
		X,Y,W,H	int32
		MenuBarVisible	bool
		StatusBarVisible	bool
		ToolBarVisible	bool
		LocationBarVisible	bool
		ScrollBarVisible	bool
		Resizeable	bool
		FullScreen	bool
	}

	WkeMediaLoadInfo struct {
		Size,W,H	int32
		Duration	float64
	}

	WkeRequestType byte
	WkeHttBodyElementType byte
)

const(
	kWkeRequestTypeInvalidation WkeRequestType=iota
    kWkeRequestTypeGet
    kWkeRequestTypePost
    kWkeRequestTypePut
)

const(
	wkeHttBodyElementTypeData WkeHttBodyElementType=iota
    wkeHttBodyElementTypeFile
)

type(
	WkePostBodyElement struct {
		Size	int32
		EType	WkeHttBodyElementType
		Data	*WkeMemBuf
		FilePath	WkeString
		FileStart	int64
		FileLength	int64
	}

	WkePostBodyElements struct {
		Size	int32
		Element	**WkePostBodyElement
		ElementSize	uint
		IsDirty	bool
	}
	JSType	byte
)

const(
	JSTYPE_NUMBER JSType=iota
    JSTYPE_STRING
    JSTYPE_BOOLEAN
    JSTYPE_OBJECT
    JSTYPE_FUNCTION
    JSTYPE_UNDEFINED
    JSTYPE_ARRAY
    JSTYPE_NULL
)

type(
	JsKeys struct {
		Len	uint16
		Keys	uintptr
	}

	JSData struct {
		TypeName	[100]byte
		PropertyGet	uintptr
		PropertySet	uintptr
		Finalize	uintptr
		CallAsFunction	uintptr
		TargetData	uintptr
	}

	JsExceptionInfo struct {
		Msg		uintptr
		SourceLine	uintptr
		scriptResourceName	uintptr
		LineNumber	int
		StartPosition	int
		EndPosition	int
		StartColumn	int
		EndColumn	int
		CallStackString	uintptr
	}
	WkeLoadingResult byte
	WkeSettingMask byte
)

const(
	WKE_LOADING_SUCCEEDED WkeLoadingResult = iota
    WKE_LOADING_FAILED
    WKE_LOADING_CANCELED
)

const(
	WKE_SETTING_PROXY WkeSettingMask=1
	WKE_SETTING_PAINTCALLBACK_IN_OTHER_THREAD WkeSettingMask = 1 << 2
)

func GetMouseFlags(mouseKeys uint16)uint32  {
	result := uint32(0)
	if mouseKeys & WinApi.MK_SHIFT != 0{
		result = uint32(WKE_SHIFT)
	}
	if mouseKeys & WinApi.MK_CONTROL != 0{
		result = result | uint32(WKE_CONTROL)
	}
	if mouseKeys & WinApi.MK_LBUTTON != 0{
		result = result | uint32(WKE_LBUTTON)
	}

	if mouseKeys & WinApi.MK_RBUTTON != 0{
		result = result | uint32(WKE_RBUTTON)
	}
	if mouseKeys & WinApi.MK_MBUTTON != 0{
		result = result | uint32(WKE_MBUTTON)
	}
	return result
}