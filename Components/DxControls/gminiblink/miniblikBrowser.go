package gminiblink

import (
	"fmt"
	"github.com/suiyunonghen/GVCL/Components"
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"github.com/suiyunonghen/GVCL/WinApi"
	"reflect"
	"unsafe"
)

type JSBindHandle func(params ...interface{})interface{}
type JSBindFunction struct {
	ParamCount	int
	BindHandle	JSBindHandle
}

type WebConsoleEvent  func(webView WkeWebView, level WkeConsoleLevel, msg, sourceName string, sourceline uint32, stackTrace string)

type GBlinkWebBrowser struct {
	controls.GWinControl
	fbindfunctions	map[string]*JSBindFunction
	fDirectWeb		bool
	webView			WkeWebView
	OnConsole		WebConsoleEvent
}

func (webBrowser *GBlinkWebBrowser) SubInit() {
	webBrowser.GWinControl.SubInit()
	webBrowser.GComponent.SubInit(webBrowser)
}

func (webBrowser *GBlinkWebBrowser) CreateParams(params *Components.GCreateParams) {
	webBrowser.GWinControl.CreateParams(params)
	params.WinClassName = "GBlinkWebBrowser"
	params.Style = (params.Style | uint32(WinApi.WS_CLIPCHILDREN) & ^(uint32(WinApi.CS_HREDRAW) | uint32(WinApi.CS_VREDRAW)))
}

func wrapBindFuncs(es JSExecState, param uintptr) JSValue{
	bindFunc := (*JSBindFunction)(unsafe.Pointer(param))
	paramCount := BlinkLib.JsArgCount(es)
	params := make([]interface{},paramCount)
	for i := 0;i<paramCount;i++{
		jsv := BlinkLib.JsArg(es, i)
		params[i] = BlinkLib.JSValue2Interface(es,jsv)
	}
	result := bindFunc.BindHandle(params)
	return BlinkLib.Value2JSValue(es,result)
}

func (webBrowser *GBlinkWebBrowser)createWebView()  {
	if BlinkLib.libminiblink == 0 {
		return
	}
	//绑定函数
	if webBrowser.fbindfunctions != nil{
		for k,v := range webBrowser.fbindfunctions{
			if v.BindHandle != nil{
				BlinkLib.WkeJsBindFunction(k,wrapBindFuncs,uintptr(unsafe.Pointer(v)),uint(v.ParamCount))
			}
		}
	}
	if webBrowser.fDirectWeb{

	}else{
		webBrowser.webView = BlinkLib.WkeCreateWebWindow(WKE_WINDOW_TYPE_CONTROL, webBrowser.GWinControl.GetWindowHandle(), 0, 0, int(webBrowser.Width()),int(webBrowser.Height()))
	}
	if webBrowser.webView == 0 {
		return
	}
	BlinkLib.WkeOnDidCreateScriptContext(webBrowser.webView, didCreateScriptContext, uintptr(unsafe.Pointer(webBrowser)))
	BlinkLib.WkeOnWillReleaseScriptContext(webBrowser.webView, willReleaseScriptContext, uintptr(unsafe.Pointer(webBrowser)))
	BlinkLib.WkeShowWindow(webBrowser.webView, webBrowser.Visible())

	/*BlinkLib.WkeOnLoadUrlBegin(webBrowser.webView, loadUrlBegin, uintptr(unsafe.Pointer(webBrowser)))
	BlinkLib.WkeOnLoadUrlEnd(webBrowser.webView, loadUrlEnd, uintptr(unsafe.Pointer(webBrowser)))
	BlinkLib.WkeOnWindowDestroy(webBrowser.webView, onDestroyWkeWindow, uintptr(unsafe.Pointer(webBrowser)))
	BlinkLib.WkeOnWindowClosing(webBrowser.webView, onClosingWkeWindow, uintptr(unsafe.Pointer(webBrowser)))*/

	BlinkLib.WkeOnConsole(webBrowser.webView, doConsole, uintptr(unsafe.Pointer(webBrowser)))
	BlinkLib.WkeSetCspCheckEnable(webBrowser.webView, false)
}



func (webBrowser *GBlinkWebBrowser) CreateWindowHandle(params *Components.GCreateParams)bool{
	if webBrowser.GWinControl.CreateWindowHandle(params){
		webBrowser.createWebView()
		return true
	}
	return false
}

func (webBrowser *GBlinkWebBrowser)destroyWebView()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeDestroyWebWindow(webBrowser.webView)
		webBrowser.webView = 0
	}
}

func (webBrowser *GBlinkWebBrowser) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	result = 0
	msgDispatchNext = true
	switch msg {
	case WinApi.WM_DESTROY:
		webBrowser.destroyWebView()
	case WinApi.WM_SIZE:
		BlinkLib.WkeResizeWindow(webBrowser.webView, int(webBrowser.Width()),int(webBrowser.Height()))
	}
	//result,msgDispatchNext = webBrowser.GWinControl.WndProc(msg,wparam,lparam)
	return
}

func (webBrowser *GBlinkWebBrowser)Navigate(url string)  {
	if webBrowser.HandleAllocated(){
		fmt.Println(url)
		BlinkLib.WkeLoadURL(webBrowser.webView,url)
	}
}

func (webBrowser *GBlinkWebBrowser)didCreateScriptContext(webView WkeWebView,frameId WkeWebFrameHandle, context uintptr, extensionGroup int, worldId int){

}

func didCreateScriptContext(webView WkeWebView, param uintptr, frameId WkeWebFrameHandle, context uintptr, extensionGroup int, worldId int){
	webBrowser := (*GBlinkWebBrowser)(unsafe.Pointer(param))
	webBrowser.didCreateScriptContext(webView,frameId,context,extensionGroup,worldId)
	return
}

func (webBrowser *GBlinkWebBrowser)willReleaseScriptContext(webView WkeWebView,frameId WkeWebFrameHandle, context uintptr, worldId int){

}

func willReleaseScriptContext(webView WkeWebView, param uintptr, frameId WkeWebFrameHandle, context uintptr, worldId int){
	webBrowser := (*GBlinkWebBrowser)(unsafe.Pointer(param))
	webBrowser.willReleaseScriptContext(webView,frameId,context,worldId)
}


func doConsole(webView WkeWebView, param uintptr, level WkeConsoleLevel, msg, sourceName string, sourceline uint32, stackTrace string){
	webBrowser := (*GBlinkWebBrowser)(unsafe.Pointer(param))
	if webBrowser.OnConsole != nil{
		webBrowser.OnConsole(webView,level,msg,sourceName,sourceline,stackTrace)
	}
}
/*func loadUrlBegin(webView WkeWebView, param uintptr, url string, job uintptr) bool{
	webBrowser := (*GBlinkWebBrowser)(unsafe.Pointer(param))

}

func loadUrlEnd(webView WkeWebView, param uintptr, url string, job uintptr, buf []byte){

}*/

func onDestroyWkeWindow(webView WkeWebView, param uintptr){

}

func onClosingWkeWindow(webView WkeWebView, param uintptr)bool{
	return true
}

func NewBlinkWebBrowser(aParent Components.IWincontrol,BindFunctions map[string]*JSBindFunction) *GBlinkWebBrowser {
	if BlinkLib.libminiblink == 0{
		return nil
	}
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		webBrowser := new(GBlinkWebBrowser)
		webBrowser.SubInit()
		webBrowser.fDirectWeb = false
		webBrowser.fbindfunctions = BindFunctions
		webBrowser.SetParent(aParent)
		webBrowser.SetWidth(200)
		webBrowser.SetVisible(true)
		webBrowser.SetHeight(300)
		return webBrowser
	}
	return nil
}