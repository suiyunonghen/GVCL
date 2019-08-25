package gminiblink

import (
	"github.com/suiyunonghen/GVCL/Components"
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/GVCL/WinApi"
	"reflect"
	"strings"
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
	OnDocumentCompleted	Graphics.NotifyEvent

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
	result := bindFunc.BindHandle(params...)
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

	BlinkLib.WkeOnDocumentReady(webBrowser.webView,wkeOnDocumentReadycallBack,uintptr(unsafe.Pointer(webBrowser)))

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

func (webBrowser *GBlinkWebBrowser)StopLoading()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeStopLoading(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)Reload()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeReload(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)ContentWidth()int  {
	if webBrowser.webView != 0{
		return BlinkLib.WkeGetContentWidth(webBrowser.webView)
	}
	return 0
}

func (webBrowser *GBlinkWebBrowser)ContentHeight()int  {
	if webBrowser.webView != 0{
		return BlinkLib.WkeGetContentHeight(webBrowser.webView)
	}
	return 0
}

func (webBrowser *GBlinkWebBrowser)WebViewDC()WinApi.HDC  {
	if webBrowser.webView != 0{
		return BlinkLib.WkeGetViewDC(webBrowser.webView)
	}
	return 0
}

func (webBrowser *GBlinkWebBrowser)CanGoBack()bool  {
	if webBrowser.webView != 0{
		return BlinkLib.WkeCanGoBack(webBrowser.webView)
	}
	return false
}

func (webBrowser *GBlinkWebBrowser)CanGoForward()bool  {
	if webBrowser.webView != 0{
		return BlinkLib.WkeCanGoForward(webBrowser.webView)
	}
	return false
}

func (webBrowser *GBlinkWebBrowser)GoBack()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeGoBack(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)GoForward()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeGoForward(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)SelectAll()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeEditorSelectAll(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)UnSelect()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeEditorUnSelect(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)Copy()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeEditorCopy(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)Cut()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeEditorCut(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)Delete()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeEditorDelete(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)Undo()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeEditorUndo(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)Redo()  {
	if webBrowser.webView != 0{
		BlinkLib.WkeEditorRedo(webBrowser.webView)
	}
}

func (webBrowser *GBlinkWebBrowser)ExecuteJavascript(js string)bool  {
	if webBrowser.webView != 0{
		v := BlinkLib.WkeRunJS(webBrowser.webView,"try { " + js + " return 1; } catch(err){ return 0;}")
		return BlinkLib.JsIsNumber(v) && BlinkLib.JsToInt(BlinkLib.WkeGlobalExec(webBrowser.webView), v) == 1
	}
	return false
}

func (webBrowser *GBlinkWebBrowser)JsCallGlobal(funcName string,params ...interface{})interface{}  {
	if webBrowser.webView != 0{
		es := BlinkLib.WkeGlobalExec(webBrowser.webView)
		funcs := strings.Split(funcName,".")
		funcCount := len(funcs)
		jsFunc := JSValue(-1)
		for i,procName := range funcs{
			if i == 0{
				jsFunc = BlinkLib.JsGetGlobal(es,procName)
			}else{
				jsFunc = BlinkLib.JsGet(es,jsFunc,procName)
			}
			funcType := BlinkLib.JsTypeOf(jsFunc)
			if i == funcCount - 1{
				if funcType == JSTYPE_FUNCTION{
					pcount := len(params)
					if len(params) == 0{
						return BlinkLib.JSValue2Interface(es,BlinkLib.JsCallGlobal(es,jsFunc,nil,0))
					}else{
						paramsp := make([]JSValue,pcount)
						for i,v := range params {
							paramsp[i] = BlinkLib.Value2JSValue(es,v)
						}
						return BlinkLib.JSValue2Interface(es,BlinkLib.JsCallGlobal(es,jsFunc,paramsp,pcount))
					}
				}
			}else if funcType != JSTYPE_OBJECT{
				return nil
			}
		}
	}
	return nil
}

func (webBrowser *GBlinkWebBrowser)MainWebView()WkeWebView  {
	return webBrowser.webView
}

func (webBrowser *GBlinkWebBrowser)LoadFromFile(webView WkeWebView, fileName string)  {
	if webBrowser.webView != 0{
		if webView != 0{
			BlinkLib.WkeLoadFile(webView,fileName)
		}else{
			BlinkLib.WkeLoadFile(webBrowser.webView,fileName)
		}
	}
}

func (webBrowser *GBlinkWebBrowser)LoadHtml(webView WkeWebView,html,baseurlPath string)  {
	if webBrowser.webView != 0{
		if webView == 0{
			webView = webBrowser.webView
		}
		if baseurlPath != ""{
			BlinkLib.WkeLoadHtmlWithBaseUrl(webView,html,baseurlPath)
		}else{
			BlinkLib.WkeLoadHTML(webView,html)
		}
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

func wkeOnDocumentReadycallBack(webview WkeWebView, param uintptr, frameid WkeFrameHwnd){
	webBrowser := (*GBlinkWebBrowser)(unsafe.Pointer(param))
	if webBrowser.OnDocumentCompleted != nil{
		webBrowser.OnDocumentCompleted(webBrowser)
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