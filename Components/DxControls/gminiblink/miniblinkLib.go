package gminiblink

import (
	"github.com/suiyunonghen/DxCommonLib"
	"github.com/suiyunonghen/GVCL/WinApi"
	"syscall"
	"unsafe"
)

type MiniBlinkLib struct {
	libminiblink                                 syscall.Handle
	wkeInitialize								 uintptr
	wkeCreateWebWindow							 uintptr
	wkeShowWindow								 uintptr
	wkeLoadURL									 uintptr
	wkeVersion									 uintptr
	wkeVersionString							 uintptr
	wkeCreateWebView							 uintptr
	wkeGC										 uintptr
	wkeIsDocumentReady							 uintptr
	wkeStopLoading							 	 uintptr
	wkeReload									 uintptr
	wkeGetTitle     uintptr
	wkeGetTitleW     uintptr
	wkeResize     uintptr
	wkeGetHeight     uintptr
	wkeGetWidth     uintptr
	wkeGetContentWidth     uintptr
	wkeGetContentHeight     uintptr
	wkePaint2     uintptr
	wkePaint     uintptr
	wkeGetViewDC     uintptr
	wkeGetHostHWND     uintptr
	wkeCanGoBack     uintptr
	wkeGoBack     uintptr
	wkeCanGoForward     uintptr
	wkeGoForward		uintptr
	wkeEditorSelectAll     uintptr
	wkeEditorUnSelect     uintptr
	wkeEditorCopy     uintptr
	wkeEditorCut     uintptr
	wkeEditorDelete     uintptr
	wkeEditorUndo     uintptr
	wkeEditorRedo     uintptr
	wkeGetCookieW     uintptr
	wkeGetCookie     uintptr
	wkeSetCookie     uintptr
	wkeVisitAllCookie     uintptr
	wkePerformCookieCommand     uintptr
	wkeSetCookieEnabled     uintptr
	wkeIsCookieEnabled     uintptr
	wkeSetCookieJarPath     uintptr
	wkeSetCookieJarFullPath     uintptr
	wkeSetLocalStorageFullPath     uintptr
	wkeSetMediaVolume     uintptr
	wkeGetMediaVolume     uintptr
	wkeFireMouseEvent     uintptr
	wkeFireContextMenuEvent     uintptr
	wkeFireMouseWheelEvent     uintptr
	wkeFireKeyUpEvent     uintptr
	wkeFireKeyDownEvent     uintptr
	wkeFireKeyPressEvent     uintptr
	wkeSetFocus     uintptr
	wkeKillFocus     uintptr
	wkeGetCaretRect     uintptr
	wkeRunJS     uintptr
	wkeFireWindowsMessage     uintptr
	wkeRunJSW     uintptr
	wkeGlobalExec     uintptr
	wkeSleep     uintptr
	wkeWake     uintptr
	wkeIsAwake	uintptr
	wkeSetZoomFactor     uintptr
	wkeGetZoomFactor     uintptr
	wkeSetEditable     uintptr
	wkeOnTitleChanged     uintptr
	wkeOnMouseOverUrlChanged     uintptr
	wkeOnURLChanged2     uintptr
	wkeOnPaintUpdated     uintptr
	wkeOnPaintBitUpdated     uintptr
	wkeOnAlertBox     uintptr
	wkeOnConfirmBox     uintptr
	wkeOnPromptBox     uintptr
	wkeOnNavigation     uintptr
	wkeOnCreateView     uintptr
	wkeOnDocumentReady     uintptr
	wkeOnDocumentReady2     uintptr
	wkeOnDownload     uintptr
	wkeNetOnResponse     uintptr
	wkeOnConsole     uintptr
	wkeSetUIThreadCallback     uintptr
	wkeOnLoadUrlBegin     uintptr
	wkeOnLoadUrlEnd     uintptr
	wkeOnDidCreateScriptContext     uintptr
	wkeOnWillReleaseScriptContext     uintptr
	wkeOnWillMediaLoad     uintptr
	wkeIsMainFrame     uintptr
	wkeWebFrameGetMainFrame     uintptr
	wkeRunJsByFrame     uintptr
	wkeGetFrameUrl     uintptr
	wkeGetString     uintptr
	wkeGetStringW     uintptr
	wkeSetString     uintptr
	wkeSetStringW     uintptr
	wkeDeleteString     uintptr
	wkeSetUserKeyValue     uintptr
	wkeGetUserKeyValue     uintptr
	wkeGetCursorInfoType     uintptr
	wkeDestroyWebView     uintptr
	wkeDestroyWebWindow     uintptr
	wkeGetWindowHandle     uintptr
	wkeOnWindowClosing     uintptr
	wkeOnWindowDestroy     uintptr
	wkeEnableWindow     uintptr
	wkeMoveWindow     uintptr
	wkeResizeWindow     uintptr
	wkeMoveToCenter     uintptr
	wkeSetWindowTitle     uintptr
	wkeSetWindowTitleW     uintptr
	wkeSetDeviceParameter     uintptr
	wkeShutdown     uintptr
	wkeSetProxy     uintptr
	wkeSetViewProxy     uintptr
	wkeConfigure     uintptr
	wkeIsInitialize     uintptr
	wkeSetMemoryCacheEnable     uintptr
	wkeFinalize     uintptr
	wkeSetTouchEnabled     uintptr
	wkeSetMouseEnabled     uintptr
	wkeSetNavigationToNewWindowEnable     uintptr
	wkeSetCspCheckEnable     uintptr
	wkeSetNpapiPluginsEnabled     uintptr
	wkeSetHeadlessEnabled     uintptr
	wkeSetDebugConfig     uintptr
	wkeSetHandle     uintptr
	wkeSetHandleOffset     uintptr
	wkeSetViewSettings     uintptr
	wkeSetTransparent     uintptr
	wkeIsTransparent     uintptr
	wkeSetUserAgent     uintptr
	wkeSetUserAgentW     uintptr
	wkeGetUserAgent     uintptr
	wkeLoadW     uintptr
	wkeLoadHTML     uintptr
	wkeLoadHtmlWithBaseUrl     uintptr
	wkeLoadFile     uintptr
	wkeGetURL     uintptr
	wkeNetSetHTTPHeaderField     uintptr
	wkeNetSetMIMEType     uintptr
	wkeNetGetMIMEType     uintptr
	wkeNetSetData     uintptr
	wkeNetCancelRequest     uintptr
	wkeNetGetFavicon     uintptr
	wkeNetHoldJobToAsynCommit     uintptr
	wkeNetGetRequestMethod     uintptr
	wkeNetGetPostBody     uintptr
	wkeNetCreatePostBodyElements     uintptr
	wkeNetFreePostBodyElements     uintptr
	wkeNetCreatePostBodyElement     uintptr
	wkeNetFreePostBodyElement     uintptr
	jsArgCount     uintptr
	jsArgType     uintptr
	jsArg     uintptr
	jsTypeOf     uintptr
	jsIsNumber     uintptr
	jsIsString     uintptr
	jsIsBoolean     uintptr
	jsIsObject     uintptr
	jsIsTrue     uintptr
	jsIsFalse     uintptr
	jsToInt     uintptr
	jsToDouble     uintptr
	jsToTempStringW     uintptr
	jsToTempString     uintptr
	jsToString     uintptr
	jsToStringW     uintptr
	jsInt     uintptr
	jsString     uintptr
	jsArrayBuffer     uintptr
	jsGetArrayBuffer     uintptr
	jsEmptyObject     uintptr
	jsEvalW     uintptr
	jsEvalExW     uintptr
	jsCall     uintptr
	jsCallGlobal     uintptr
	jsGet     uintptr
	jsSet     uintptr
	jsGetGlobal     uintptr
	jsSetGlobal     uintptr
	jsGetAt     uintptr
	jsSetAt     uintptr
	jsGetKeys     uintptr
	jsGetLength     uintptr
	jsSetLength     uintptr
	jsGetWebView     uintptr
	jsGC     uintptr
	jsBindFunction     uintptr
	jsBindGetter     uintptr
	jsBindSetter     uintptr
	wkeJsBindFunction     uintptr
	jsObject     uintptr
	jsFunction     uintptr
	jsGetData     uintptr
	jsGetLastErrorIfException     uintptr
	wkeShowDevtools     uintptr
	jsFloat     uintptr
	jsDouble     uintptr
	jsBoolean     uintptr
	jsUndefined     uintptr
	jsNull     uintptr
	jsTrue     uintptr
	jsFalse     uintptr
	jsStringW     uintptr
	jsEmptyArray     uintptr
	wkeNetHookRequest     uintptr
	wkeNetStartUrlRequest     uintptr
	wkeOnLoadingFinish     uintptr
	wkeCreateWebCustomWindow     uintptr
	wkeSetCursorInfoType     uintptr

	fCookieVisitorCallBack	WkeCookieVisitorCallBack
	fOnTitleChanged		WkeTitleChanged
}

func (blink *MiniBlinkLib)LoadBlink(blinkPath string) bool{
	blink.UnLoad()
	blink.libminiblink,_ = syscall.LoadLibrary(blinkPath)
	if blink.libminiblink == 0{
		return false
	}
	blink.wkeInitialize, _ = syscall.GetProcAddress(blink.libminiblink, "wkeInitialize")
	blink.wkeCreateWebWindow, _ = syscall.GetProcAddress(blink.libminiblink, "wkeCreateWebWindow")
	blink.wkeShowWindow, _ = syscall.GetProcAddress(blink.libminiblink, "wkeShowWindow")
	blink.wkeLoadURL, _ = syscall.GetProcAddress(blink.libminiblink, "wkeLoadURLW")
	blink.wkeVersion, _ = syscall.GetProcAddress(blink.libminiblink, "wkeVersion")
	blink.wkeVersionString, _ = syscall.GetProcAddress(blink.libminiblink, "wkeVersionString")
	blink.wkeCreateWebView, _ = syscall.GetProcAddress(blink.libminiblink, "wkeCreateWebView")
	blink.wkeGC, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGC")
	blink.wkeIsDocumentReady,_= syscall.GetProcAddress(blink.libminiblink, "wkeIsDocumentReady")
	blink.wkeStopLoading,_= syscall.GetProcAddress(blink.libminiblink, "wkeStopLoading")
	blink.wkeReload,_= syscall.GetProcAddress(blink.libminiblink, "wkeReload")
	blink.wkeGetTitle, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetTitle")
	blink.wkeGetTitleW, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetTitleW")
	blink.wkeResize, _ = syscall.GetProcAddress(blink.libminiblink, "wkeResize")
	blink.wkeGetHeight, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetHeight")
	blink.wkeGetWidth, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetWidth")
	blink.wkeGetContentWidth, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetContentWidth")
	blink.wkeGetContentHeight, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetContentHeight")
	blink.wkePaint2, _ = syscall.GetProcAddress(blink.libminiblink, "wkePaint2")
	blink.wkePaint, _ = syscall.GetProcAddress(blink.libminiblink, "wkePaint")
	blink.wkeGetViewDC, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetViewDC")
	blink.wkeGetHostHWND, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetHostHWND")
	blink.wkeCanGoBack, _ = syscall.GetProcAddress(blink.libminiblink, "wkeCanGoBack")
	blink.wkeGoBack, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGoBack")
	blink.wkeCanGoForward, _ = syscall.GetProcAddress(blink.libminiblink, "wkeCanGoForward")
	blink.wkeGoForward, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGoForward")
	blink.wkeEditorSelectAll, _ = syscall.GetProcAddress(blink.libminiblink, "wkeEditorSelectAll")
	blink.wkeEditorUnSelect, _ = syscall.GetProcAddress(blink.libminiblink, "wkeEditorUnSelect")
	blink.wkeEditorCopy, _ = syscall.GetProcAddress(blink.libminiblink, "wkeEditorCopy")
	blink.wkeEditorCut, _ = syscall.GetProcAddress(blink.libminiblink, "wkeEditorCut")
	blink.wkeEditorDelete, _ = syscall.GetProcAddress(blink.libminiblink, "wkeEditorDelete")
	blink.wkeEditorUndo, _ = syscall.GetProcAddress(blink.libminiblink, "wkeEditorUndo")
	blink.wkeEditorRedo, _ = syscall.GetProcAddress(blink.libminiblink, "wkeEditorRedo")
	blink.wkeGetCookieW, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetCookieW")
	blink.wkeGetCookie, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetCookie")
	blink.wkeSetCookie, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetCookie")
	blink.wkeVisitAllCookie, _ = syscall.GetProcAddress(blink.libminiblink, "wkeVisitAllCookie")
	blink.wkePerformCookieCommand, _ = syscall.GetProcAddress(blink.libminiblink, "wkePerformCookieCommand")
	blink.wkeSetCookieEnabled, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetCookieEnabled")
	blink.wkeIsCookieEnabled, _ = syscall.GetProcAddress(blink.libminiblink, "wkeIsCookieEnabled")
	blink.wkeSetCookieJarPath, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetCookieJarPath")
	blink.wkeSetCookieJarFullPath, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetCookieJarFullPath")
	blink.wkeSetLocalStorageFullPath, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetLocalStorageFullPath")
	blink.wkeSetMediaVolume, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetMediaVolume")
	blink.wkeGetMediaVolume, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetMediaVolume")
	blink.wkeFireMouseEvent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeFireMouseEvent")
	blink.wkeFireContextMenuEvent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeFireContextMenuEvent")
	blink.wkeFireMouseWheelEvent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeFireMouseWheelEvent")
	blink.wkeFireKeyUpEvent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeFireKeyUpEvent")
	blink.wkeFireKeyDownEvent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeFireKeyDownEvent")
	blink.wkeFireKeyPressEvent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeFireKeyPressEvent")
	blink.wkeSetFocus, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetFocus")
	blink.wkeKillFocus, _ = syscall.GetProcAddress(blink.libminiblink, "wkeKillFocus")
	blink.wkeGetCaretRect, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetCaretRect")
	blink.wkeRunJS, _ = syscall.GetProcAddress(blink.libminiblink, "wkeRunJS")
	blink.wkeFireWindowsMessage, _ = syscall.GetProcAddress(blink.libminiblink, "wkeFireWindowsMessage")
	blink.wkeRunJSW, _ = syscall.GetProcAddress(blink.libminiblink, "wkeRunJSW")
	blink.wkeGlobalExec, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGlobalExec")
	blink.wkeSleep, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSleep")
	blink.wkeWake, _ = syscall.GetProcAddress(blink.libminiblink, "wkeWake")
	blink.wkeIsAwake, _ = syscall.GetProcAddress(blink.libminiblink, "wkeIsAwake")
	blink.wkeSetZoomFactor, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetZoomFactor")
	blink.wkeGetZoomFactor, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetZoomFactor")
	blink.wkeSetEditable, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetEditable")
	blink.wkeOnTitleChanged, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnTitleChanged")
	blink.wkeOnMouseOverUrlChanged, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnMouseOverUrlChanged")
	blink.wkeOnURLChanged2, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnURLChanged2")
	blink.wkeOnPaintUpdated, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnPaintUpdated")
	blink.wkeOnPaintBitUpdated, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnPaintBitUpdated")
	blink.wkeOnAlertBox, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnAlertBox")
	blink.wkeOnConfirmBox, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnConfirmBox")
	blink.wkeOnPromptBox, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnPromptBox")
	blink.wkeOnNavigation, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnNavigation")
	blink.wkeOnCreateView, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnCreateView")
	blink.wkeOnDocumentReady, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnDocumentReady")
	blink.wkeOnDocumentReady2, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnDocumentReady2")
	blink.wkeOnDownload, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnDownload")
	blink.wkeNetOnResponse, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetOnResponse")
	blink.wkeOnConsole, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnConsole")
	blink.wkeSetUIThreadCallback, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetUIThreadCallback")
	blink.wkeOnLoadUrlBegin, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnLoadUrlBegin")
	blink.wkeOnLoadUrlEnd, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnLoadUrlEnd")
	blink.wkeOnDidCreateScriptContext, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnDidCreateScriptContext")
	blink.wkeOnWillReleaseScriptContext, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnWillReleaseScriptContext")
	blink.wkeOnWillMediaLoad, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnWillMediaLoad")
	blink.wkeIsMainFrame, _ = syscall.GetProcAddress(blink.libminiblink, "wkeIsMainFrame")
	blink.wkeWebFrameGetMainFrame, _ = syscall.GetProcAddress(blink.libminiblink, "wkeWebFrameGetMainFrame")
	blink.wkeRunJsByFrame, _ = syscall.GetProcAddress(blink.libminiblink, "wkeRunJsByFrame")
	blink.wkeGetFrameUrl, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetFrameUrl")
	blink.wkeGetString, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetString")
	blink.wkeGetStringW, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetStringW")
	blink.wkeSetString, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetString")
	blink.wkeSetStringW, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetStringW")
	blink.wkeDeleteString, _ = syscall.GetProcAddress(blink.libminiblink, "wkeDeleteString")
	blink.wkeSetUserKeyValue, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetUserKeyValue")
	blink.wkeGetUserKeyValue, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetUserKeyValue")
	blink.wkeGetCursorInfoType, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetCursorInfoType")
	blink.wkeDestroyWebView, _ = syscall.GetProcAddress(blink.libminiblink, "wkeDestroyWebView")
	blink.wkeDestroyWebWindow, _ = syscall.GetProcAddress(blink.libminiblink, "wkeDestroyWebWindow")
	blink.wkeGetWindowHandle, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetWindowHandle")
	blink.wkeOnWindowClosing, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnWindowClosing")
	blink.wkeOnWindowDestroy, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnWindowDestroy")
	blink.wkeEnableWindow, _ = syscall.GetProcAddress(blink.libminiblink, "wkeEnableWindow")
	blink.wkeMoveWindow, _ = syscall.GetProcAddress(blink.libminiblink, "wkeMoveWindow")
	blink.wkeResizeWindow, _ = syscall.GetProcAddress(blink.libminiblink, "wkeResizeWindow")
	blink.wkeMoveToCenter, _ = syscall.GetProcAddress(blink.libminiblink, "wkeMoveToCenter")
	blink.wkeSetWindowTitle, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetWindowTitle")
	blink.wkeSetWindowTitleW, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetWindowTitleW")
	blink.wkeSetDeviceParameter, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetDeviceParameter")
	blink.wkeShutdown, _ = syscall.GetProcAddress(blink.libminiblink, "wkeShutdown")
	blink.wkeSetProxy, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetProxy")
	blink.wkeSetViewProxy, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetViewProxy")
	blink.wkeConfigure, _ = syscall.GetProcAddress(blink.libminiblink, "wkeConfigure")
	blink.wkeIsInitialize, _ = syscall.GetProcAddress(blink.libminiblink, "wkeIsInitialize")
	blink.wkeSetMemoryCacheEnable, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetMemoryCacheEnable")
	blink.wkeFinalize, _ = syscall.GetProcAddress(blink.libminiblink, "wkeFinalize")
	blink.wkeSetTouchEnabled, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetTouchEnabled")
	blink.wkeSetMouseEnabled, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetMouseEnabled")
	blink.wkeSetNavigationToNewWindowEnable, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetNavigationToNewWindowEnable")
	blink.wkeSetCspCheckEnable, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetCspCheckEnable")
	blink.wkeSetNpapiPluginsEnabled, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetNpapiPluginsEnabled")
	blink.wkeSetHeadlessEnabled, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetHeadlessEnabled")
	blink.wkeSetDebugConfig, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetDebugConfig")
	blink.wkeSetHandle, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetHandle")
	blink.wkeSetHandleOffset, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetHandleOffset")
	blink.wkeSetViewSettings, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetViewSettings")
	blink.wkeSetTransparent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetTransparent")
	blink.wkeIsTransparent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeIsTransparent")
	blink.wkeSetUserAgent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetUserAgent")
	blink.wkeSetUserAgentW, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetUserAgentW")
	blink.wkeGetUserAgent, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetUserAgent")
	blink.wkeLoadW, _ = syscall.GetProcAddress(blink.libminiblink, "wkeLoadW")
	blink.wkeLoadHTML, _ = syscall.GetProcAddress(blink.libminiblink, "wkeLoadHTML")
	blink.wkeLoadHtmlWithBaseUrl, _ = syscall.GetProcAddress(blink.libminiblink, "wkeLoadHtmlWithBaseUrl")
	blink.wkeLoadFile, _ = syscall.GetProcAddress(blink.libminiblink, "wkeLoadFile")
	blink.wkeGetURL, _ = syscall.GetProcAddress(blink.libminiblink, "wkeGetURL")
	blink.wkeNetSetHTTPHeaderField, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetSetHTTPHeaderField")
	blink.wkeNetSetMIMEType, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetSetMIMEType")
	blink.wkeNetGetMIMEType, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetGetMIMEType")
	blink.wkeNetSetData, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetSetData")
	blink.wkeNetCancelRequest, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetCancelRequest")
	blink.wkeNetGetFavicon, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetGetFavicon")
	blink.wkeNetHoldJobToAsynCommit, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetHoldJobToAsynCommit")
	blink.wkeNetGetRequestMethod, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetGetRequestMethod")
	blink.wkeNetGetPostBody, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetGetPostBody")
	blink.wkeNetCreatePostBodyElements, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetCreatePostBodyElements")
	blink.wkeNetFreePostBodyElements, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetFreePostBodyElements")
	blink.wkeNetCreatePostBodyElement, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetCreatePostBodyElement")
	blink.wkeNetFreePostBodyElement, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetFreePostBodyElement")
	blink.jsArgCount, _ = syscall.GetProcAddress(blink.libminiblink, "jsArgCount")
	blink.jsArgType, _ = syscall.GetProcAddress(blink.libminiblink, "jsArgType")
	blink.jsArg, _ = syscall.GetProcAddress(blink.libminiblink, "jsArg")
	blink.jsTypeOf, _ = syscall.GetProcAddress(blink.libminiblink, "jsTypeOf")
	blink.jsIsNumber, _ = syscall.GetProcAddress(blink.libminiblink, "jsIsNumber")
	blink.jsIsString, _ = syscall.GetProcAddress(blink.libminiblink, "jsIsString")
	blink.jsIsBoolean, _ = syscall.GetProcAddress(blink.libminiblink, "jsIsBoolean")
	blink.jsIsObject, _ = syscall.GetProcAddress(blink.libminiblink, "jsIsObject")
	blink.jsIsTrue, _ = syscall.GetProcAddress(blink.libminiblink, "jsIsTrue")
	blink.jsIsFalse, _ = syscall.GetProcAddress(blink.libminiblink, "jsIsFalse")
	blink.jsToInt, _ = syscall.GetProcAddress(blink.libminiblink, "jsToInt")
	blink.jsToDouble, _ = syscall.GetProcAddress(blink.libminiblink, "jsToDouble")
	blink.jsToTempStringW, _ = syscall.GetProcAddress(blink.libminiblink, "jsToTempStringW")
	blink.jsToTempString, _ = syscall.GetProcAddress(blink.libminiblink, "jsToTempString")
	blink.jsToString, _ = syscall.GetProcAddress(blink.libminiblink, "jsToString")
	blink.jsToStringW, _ = syscall.GetProcAddress(blink.libminiblink, "jsToStringW")
	blink.jsInt, _ = syscall.GetProcAddress(blink.libminiblink, "jsInt")
	blink.jsString, _ = syscall.GetProcAddress(blink.libminiblink, "jsString")
	blink.jsArrayBuffer, _ = syscall.GetProcAddress(blink.libminiblink, "jsArrayBuffer")
	blink.jsGetArrayBuffer, _ = syscall.GetProcAddress(blink.libminiblink, "jsGetArrayBuffer")
	blink.jsEmptyObject, _ = syscall.GetProcAddress(blink.libminiblink, "jsEmptyObject")
	blink.jsEvalW, _ = syscall.GetProcAddress(blink.libminiblink, "jsEvalW")
	blink.jsEvalExW, _ = syscall.GetProcAddress(blink.libminiblink, "jsEvalExW")
	blink.jsCall, _ = syscall.GetProcAddress(blink.libminiblink, "jsCall")
	blink.jsCallGlobal, _ = syscall.GetProcAddress(blink.libminiblink, "jsCallGlobal")
	blink.jsGet, _ = syscall.GetProcAddress(blink.libminiblink, "jsGet")
	blink.jsSet, _ = syscall.GetProcAddress(blink.libminiblink, "jsSet")
	blink.jsGetGlobal, _ = syscall.GetProcAddress(blink.libminiblink, "jsGetGlobal")
	blink.jsSetGlobal, _ = syscall.GetProcAddress(blink.libminiblink, "jsSetGlobal")
	blink.jsGetAt, _ = syscall.GetProcAddress(blink.libminiblink, "jsGetAt")
	blink.jsSetAt, _ = syscall.GetProcAddress(blink.libminiblink, "jsSetAt")
	blink.jsGetKeys, _ = syscall.GetProcAddress(blink.libminiblink, "jsGetKeys")
	blink.jsGetLength, _ = syscall.GetProcAddress(blink.libminiblink, "jsGetLength")
	blink.jsSetLength, _ = syscall.GetProcAddress(blink.libminiblink, "jsSetLength")
	blink.jsGetWebView, _ = syscall.GetProcAddress(blink.libminiblink, "jsGetWebView")
	blink.jsGC, _ = syscall.GetProcAddress(blink.libminiblink, "jsGC")
	blink.jsBindFunction, _ = syscall.GetProcAddress(blink.libminiblink, "jsBindFunction")
	blink.jsBindGetter, _ = syscall.GetProcAddress(blink.libminiblink, "jsBindGetter")
	blink.jsBindSetter, _ = syscall.GetProcAddress(blink.libminiblink, "jsBindSetter")
	blink.wkeJsBindFunction, _ = syscall.GetProcAddress(blink.libminiblink, "wkeJsBindFunction")
	blink.jsObject, _ = syscall.GetProcAddress(blink.libminiblink, "jsObject")
	blink.jsFunction, _ = syscall.GetProcAddress(blink.libminiblink, "jsFunction")
	blink.jsGetData, _ = syscall.GetProcAddress(blink.libminiblink, "jsGetData")
	blink.jsGetLastErrorIfException, _ = syscall.GetProcAddress(blink.libminiblink, "jsGetLastErrorIfException")
	blink.wkeShowDevtools, _ = syscall.GetProcAddress(blink.libminiblink, "wkeShowDevtools")
	blink.jsFloat, _ = syscall.GetProcAddress(blink.libminiblink, "jsFloat")
	blink.jsDouble, _ = syscall.GetProcAddress(blink.libminiblink, "jsDouble")
	blink.jsBoolean, _ = syscall.GetProcAddress(blink.libminiblink, "jsBoolean")
	blink.jsUndefined, _ = syscall.GetProcAddress(blink.libminiblink, "jsUndefined")
	blink.jsNull, _ = syscall.GetProcAddress(blink.libminiblink, "jsNull")
	blink.jsTrue, _ = syscall.GetProcAddress(blink.libminiblink, "jsTrue")
	blink.jsFalse, _ = syscall.GetProcAddress(blink.libminiblink, "jsFalse")
	blink.jsStringW, _ = syscall.GetProcAddress(blink.libminiblink, "jsStringW")
	blink.jsEmptyArray, _ = syscall.GetProcAddress(blink.libminiblink, "jsEmptyArray")
	blink.wkeNetHookRequest, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetHookRequest")
	blink.wkeNetStartUrlRequest, _ = syscall.GetProcAddress(blink.libminiblink, "wkeNetStartUrlRequest")
	blink.wkeOnLoadingFinish, _ = syscall.GetProcAddress(blink.libminiblink, "wkeOnLoadingFinish")
	blink.wkeCreateWebCustomWindow, _ = syscall.GetProcAddress(blink.libminiblink, "wkeCreateWebCustomWindow")
	blink.wkeSetCursorInfoType, _ = syscall.GetProcAddress(blink.libminiblink, "wkeSetCursorInfoType")

	syscall.Syscall(blink.wkeInitialize,0,0,0,0)
	return true
}


func (blink *MiniBlinkLib)UnLoad()  {
	if blink.libminiblink != 0{
		syscall.FreeLibrary(blink.libminiblink)
		blink.fCookieVisitorCallBack = nil
	}
}


func (blink *MiniBlinkLib)WkeIsDocumentReady(webView WkeWebView)bool  {
	if blink.libminiblink != 0 && blink.wkeIsDocumentReady != 0{
		ret,_,_ := syscall.Syscall(blink.wkeIsDocumentReady,1,uintptr(webView),0,0)
		return ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)WkeStopLoading(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeStopLoading != 0{
		syscall.Syscall(blink.wkeStopLoading,1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeReload(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeReload != 0{
		syscall.Syscall(blink.wkeReload,1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)Version()uint32  {
	if blink.libminiblink != 0 && blink.wkeVersion != 0{
		ret,_,_ := syscall.Syscall(blink.wkeVersion,0,0,0,0)
		return uint32(ret)
	}
	return 0
}

func (blink *MiniBlinkLib)VersionString()string  {
	if blink.libminiblink != 0 && blink.wkeVersionString != 0{
		ret,_,_ := syscall.Syscall(blink.wkeVersionString,0,0,0,0)
		return DxCommonLib.StringFromUtf8Pointer(ret,255)
	}
	return ""
}




func (blink *MiniBlinkLib)WkeGetTitle(webview WkeWebView) string {
	if blink.libminiblink != 0 {
		if blink.wkeGetTitleW != 0{
			ret,_,_ := syscall.Syscall(blink.wkeGetTitleW,1,uintptr(webview),0,0)
			return DxCommonLib.StringFromUtf16Pointer(ret,1024)
		}else if blink.wkeGetTitle != 0{
			ret,_,_ := syscall.Syscall(blink.wkeGetTitle,1,uintptr(webview),0,0)
			return DxCommonLib.StringFromUtf8Pointer(ret,1024)
		}
	}
	return ""
}

func (blink *MiniBlinkLib)WkeGetCookie(webview WkeWebView) string {
	if blink.libminiblink != 0 {
		if blink.wkeGetTitleW != 0{
			ret,_,_ := syscall.Syscall(blink.wkeGetCookieW,1,uintptr(webview),0,0)
			return DxCommonLib.StringFromUtf16Pointer(ret,1024)
		}else if blink.wkeGetTitle != 0{
			ret,_,_ := syscall.Syscall(blink.wkeGetCookie,1,uintptr(webview),0,0)
			return DxCommonLib.StringFromUtf8Pointer(ret,1024)
		}
	}
	return ""
}

func (blink *MiniBlinkLib)WkeSetCookie(webview WkeWebView,url,cookie string)  {
	if blink.libminiblink != 0 && blink.wkeSetCookie != 0{
		syscall.Syscall(blink.wkeSetCookie, 3, uintptr(webview), uintptr(unsafe.Pointer(&url)), uintptr(unsafe.Pointer(&cookie)))
	}
}

func (blink *MiniBlinkLib)WkeShowWindow(wkeWebView uintptr, showFlag bool){
	if blink.libminiblink != 0 && blink.wkeShowWindow != 0{
		syscall.Syscall(blink.wkeShowWindow, 2, wkeWebView, 1, 0)
	}
}

func (blink *MiniBlinkLib)WkeLoadURL(wkeWebView WkeWebView,url string)  {
	if blink.libminiblink != 0 && blink.wkeLoadURL != 0 {
		syscall.Syscall(blink.wkeLoadURL, 2, uintptr(wkeWebView), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(url))), 0)
	}
}

func (blink *MiniBlinkLib)WkeCreateWebWindow(Type uintptr, parent int, x int, y int, width int, height int) uintptr{
	if blink.libminiblink != 0 && blink.wkeCreateWebWindow != 0 {
		ret, _, _ := syscall.Syscall6(blink.wkeCreateWebWindow, 6, Type, uintptr(parent), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
		return ret
	}
	return 0
}

func (blink *MiniBlinkLib)WkeCreateWebView()WkeWebView  {
	if blink.libminiblink != 0 && blink.wkeCreateWebView != 0 {
		ret, _, _ := syscall.Syscall(blink.wkeCreateWebView, 0,0,0, 0)
		return WkeWebView(ret)
	}
	return 0
}

func (blink *MiniBlinkLib)WkeGC(webView WkeWebView,delayMs uint)  {
	if blink.libminiblink != 0 && blink.wkeGC != 0 {
		syscall.Syscall(blink.wkeGC, 2,uintptr(webView),uintptr(delayMs), 0)
	}
}

func (blink *MiniBlinkLib)WkeResize(webView WkeWebView,width,height int)  {
	if blink.libminiblink != 0 && blink.wkeResize != 0 {
		syscall.Syscall(blink.wkeResize, 3,uintptr(webView),uintptr(width), uintptr(height))
	}
}

func (blink *MiniBlinkLib)WkeGetHeight(webView WkeWebView)int  {
	if blink.libminiblink != 0 && blink.wkeGetHeight != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGetHeight, 1,uintptr(webView),0,0)
		return int(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)WkeGetWidth(webView WkeWebView)int  {
	if blink.libminiblink != 0 && blink.wkeGetWidth != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGetWidth, 1,uintptr(webView),0,0)
		return int(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)WkeGetContentHeight(webView WkeWebView)int  {
	if blink.libminiblink != 0 && blink.wkeGetContentHeight != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGetContentHeight, 1,uintptr(webView),0,0)
		return int(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)WkeGetContentWidth(webView WkeWebView)int  {
	if blink.libminiblink != 0 && blink.wkeGetContentWidth != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGetContentWidth, 1,uintptr(webView),0,0)
		return int(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)WkePaint2(webView WkeWebView,bitsbuffer uintptr,bufwidth,bufheight,xDst,yDst,w,h,xsrc,ysrc int,bCopyAlpha bool)  {
	if blink.libminiblink != 0 && blink.wkePaint2 != 0 {
		syscall.Syscall12(blink.wkePaint2,11,uintptr(webView),bitsbuffer,uintptr(bufwidth),uintptr(bufheight),uintptr(xDst),uintptr(yDst),
			uintptr(w),uintptr(h),uintptr(xsrc),uintptr(ysrc),uintptr(DxCommonLib.Ord(bCopyAlpha)),0)
	}
}

func (blink *MiniBlinkLib)WkePaint(webView WkeWebView,bitsbuffer uintptr,pitch int) {
	if blink.libminiblink != 0 && blink.wkePaint != 0 {
		syscall.Syscall(blink.wkePaint, 3,uintptr(webView),bitsbuffer,uintptr(pitch))
	}
}

func (blink *MiniBlinkLib)WkeGetViewDC(webview WkeWebView)WinApi.HDC  {
	if blink.libminiblink != 0 && blink.wkeGetViewDC != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGetViewDC, 1,uintptr(webview),0,0)
		return WinApi.HDC(ret)
	}
	return 0
}

func (blink *MiniBlinkLib)WkeGetHostHWND(webView WkeWebView)syscall.Handle  {
	if blink.libminiblink != 0 && blink.wkeGetHostHWND != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGetHostHWND, 1,uintptr(webView),0,0)
		return syscall.Handle(ret)
	}
	return 0
}

func (blink *MiniBlinkLib)WkeCanGoBack(webView WkeWebView)bool  {
	if blink.libminiblink != 0 && blink.wkeCanGoBack != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeCanGoBack, 1,uintptr(webView),0,0)
		return ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)WkeCanGoForward(webView WkeWebView)bool  {
	if blink.libminiblink != 0 && blink.wkeCanGoForward != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeCanGoForward, 1,uintptr(webView),0,0)
		return ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)WkeGoBack(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeGoBack != 0 {
		syscall.Syscall(blink.wkeGoBack, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeGoForward(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeGoForward != 0 {
		syscall.Syscall(blink.wkeGoForward, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeEditorSelectAll(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeEditorSelectAll != 0 {
		syscall.Syscall(blink.wkeEditorSelectAll, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeEditorUnSelect(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeEditorUnSelect != 0 {
		syscall.Syscall(blink.wkeEditorUnSelect, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeEditorCopy(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeEditorCopy != 0 {
		syscall.Syscall(blink.wkeEditorCopy, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeEditorCut(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeEditorCut != 0 {
		syscall.Syscall(blink.wkeEditorCut, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeEditorDelete(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeEditorDelete != 0 {
		syscall.Syscall(blink.wkeEditorDelete, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeEditorUndo(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeEditorUndo != 0 {
		syscall.Syscall(blink.wkeEditorUndo, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeEditorRedo(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeEditorRedo != 0 {
		syscall.Syscall(blink.wkeEditorRedo, 1,uintptr(webView),0,0)
	}
}


type WkeCookieVisitorCallBack func(params uintptr,name,value,domain,path string,secure,httpOnly int,expires *int)bool

func wkeCookieVisitorCallBack(params uintptr,name,value,domain,path uintptr,secure,httpOnly int,expires *int)bool  {
	if BlinkLib.fCookieVisitorCallBack != nil{
		return BlinkLib.fCookieVisitorCallBack(params,DxCommonLib.StringFromUtf8Pointer(name,1024),
			DxCommonLib.StringFromUtf8Pointer(value,1024),DxCommonLib.StringFromUtf8Pointer(domain,1024),
			DxCommonLib.StringFromUtf8Pointer(path,1024),secure,httpOnly,expires)
	}
	return false
}

func (blink *MiniBlinkLib)WkeVisitAllCookie(params uintptr,callback WkeCookieVisitorCallBack)  {
	if blink.libminiblink != 0 && blink.wkeVisitAllCookie != 0 {
		if fwkeCookieVisitorCallBack == 0{
			fwkeCookieVisitorCallBack = syscall.NewCallbackCDecl(wkeCookieVisitorCallBack)
		}
		blink.fCookieVisitorCallBack = callback
		syscall.Syscall(blink.wkeVisitAllCookie, 2,uintptr(params),fwkeCookieVisitorCallBack,0)
	}
}

func (blink *MiniBlinkLib)WkePerformCookieCommand(command WkeCookieCommand)  {
	if blink.libminiblink != 0 && blink.wkePerformCookieCommand != 0 {
		syscall.Syscall(blink.wkePerformCookieCommand, 1,uintptr(command),0,0)
	}
}

func (blink *MiniBlinkLib)WkeSetCookieEnabled(webView WkeWebView,enabled bool)  {
	if blink.libminiblink != 0 && blink.wkeSetCookieEnabled != 0 {
		syscall.Syscall(blink.wkeSetCookieEnabled, 2,uintptr(webView),uintptr(DxCommonLib.Ord(enabled)),0)
	}
}

func (blink *MiniBlinkLib)WkeIsCookieEnabled(webView WkeWebView)bool  {
	if blink.libminiblink != 0 && blink.wkeIsCookieEnabled != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeIsCookieEnabled, 1,uintptr(webView),0,0)
		return ret!=0
	}
	return false
}

func (blink *MiniBlinkLib)WkeSetCookieJarPath(webView WkeWebView,path string)  {
	if blink.libminiblink != 0 && blink.wkeSetCookieJarPath != 0 {
		utf16ptr,_ := syscall.UTF16PtrFromString(path)
		syscall.Syscall(blink.wkeSetCookieJarPath, 2,uintptr(webView),uintptr(unsafe.Pointer(utf16ptr)),0)
	}
}

func (blink *MiniBlinkLib)WkeSetCookieJarFullPath(webView WkeWebView,path string)  {
	if blink.libminiblink != 0 && blink.wkeSetCookieJarFullPath != 0 {
		utf16ptr,_ := syscall.UTF16PtrFromString(path)
		syscall.Syscall(blink.wkeSetCookieJarFullPath, 2,uintptr(webView),uintptr(unsafe.Pointer(utf16ptr)),0)
	}
}

func (blink *MiniBlinkLib)WkeSetLocalStorageFullPath(webView WkeWebView,path string)  {
	if blink.libminiblink != 0 && blink.wkeSetLocalStorageFullPath != 0 {
		utf16ptr,_ := syscall.UTF16PtrFromString(path)
		syscall.Syscall(blink.wkeSetLocalStorageFullPath, 2,uintptr(webView),uintptr(unsafe.Pointer(utf16ptr)),0)
	}
}

func (blink *MiniBlinkLib)WkeSetMediaVolume(webView WkeWebView,volume float32)  {
	if blink.libminiblink != 0 && blink.wkeSetMediaVolume != 0 {
		syscall.Syscall(blink.wkeSetMediaVolume, 2,uintptr(webView),uintptr(volume),0)
	}
}

func (blink *MiniBlinkLib)WkeFireMouseEvent(webView WkeWebView,message uint,x,y int,flags uint)bool  {
	if blink.libminiblink != 0 && blink.wkeFireMouseEvent != 0 {
		ret,_,_ := syscall.Syscall6(blink.wkeFireMouseEvent, 5,uintptr(webView),uintptr(message),uintptr(x),uintptr(y),uintptr(flags),0)
		return ret!=0
	}
	return false
}

/*func (blink *MiniBlinkLib)wkeFireContextMenuEvent()  {

}*/

func (blink *MiniBlinkLib)WkeFireMouseWheelEvent(webView WkeWebView,x,y, delta int,flags uint)bool  {
	if blink.libminiblink != 0 && blink.wkeFireMouseWheelEvent != 0 {
		ret,_,_ := syscall.Syscall6(blink.wkeFireMouseWheelEvent, 5,uintptr(webView),uintptr(x),uintptr(y),uintptr(delta),uintptr(flags),0)
		return ret!=0
	}
	return false
}

func (blink *MiniBlinkLib)WkeFireKeyUpEvent(webView WkeWebView,virtualKeyCode,flags uint,systemKey bool)bool  {
	if blink.libminiblink != 0 && blink.wkeFireKeyUpEvent != 0 {
		ret,_,_ := syscall.Syscall6(blink.wkeFireKeyUpEvent, 4,uintptr(webView),uintptr(virtualKeyCode),uintptr(flags),
			uintptr(DxCommonLib.Ord(systemKey)),0,0)
		return ret!=0
	}
	return false
}

func (blink *MiniBlinkLib)WkeFireKeyDownEvent(webView WkeWebView,virtualKeyCode,flags uint,systemKey bool)bool  {
	if blink.libminiblink != 0 && blink.wkeFireKeyDownEvent != 0 {
		ret,_,_ := syscall.Syscall6(blink.wkeFireKeyDownEvent, 4,uintptr(webView),uintptr(virtualKeyCode),uintptr(flags),
			uintptr(DxCommonLib.Ord(systemKey)),0,0)
		return ret!=0
	}
	return false
}

func (blink *MiniBlinkLib)WkeFireWindowsMessage(webView WkeWebView,hWnd syscall.Handle, message uint,wParam,lParam uintptr,result *WinApi.LRESULT)bool  {
	if blink.libminiblink != 0 && blink.wkeFireWindowsMessage != 0 {
		ret,_,_ := syscall.Syscall6(blink.wkeFireWindowsMessage, 5,uintptr(webView),uintptr(hWnd),
			uintptr(message),wParam,lParam,uintptr(unsafe.Pointer(result)))
		return ret!=0
	}
	return false
}

func (blink *MiniBlinkLib)WkeSetFocus(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeSetFocus != 0 {
		syscall.Syscall(blink.wkeSetFocus, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeKillFocus(webView WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeKillFocus != 0 {
		syscall.Syscall(blink.wkeKillFocus, 1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeGetCaretRect(webView WkeWebView)WkeRect  {
	if blink.libminiblink != 0 && blink.wkeGetCaretRect != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGetCaretRect, 1,uintptr(webView),0,0)
		return *(*WkeRect)(unsafe.Pointer(ret))
	}
	return WkeRect{}
}

func (blink *MiniBlinkLib)WkeRunJS(webView WkeWebView,script string)JSValue  {
	if blink.libminiblink != 0 && blink.wkeRunJSW != 0 {
		utf16ptr,_ := syscall.UTF16PtrFromString(script)
		ret,_,_ := syscall.Syscall(blink.wkeRunJSW, 2,uintptr(webView),uintptr(unsafe.Pointer(utf16ptr)),0)
		return JSValue(ret)
	}
	return 0
}

func (blink *MiniBlinkLib)WkeGlobalExec(webView WkeWebView)JSExecState  {
	if blink.libminiblink != 0 && blink.wkeGlobalExec != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGlobalExec, 1,uintptr(webView),0,0)
		return JSExecState(ret)
	}
	return 0
}

func (blink *MiniBlinkLib)WkeSetZoomFactor(webView WkeWebView,factor float32)  {
	if blink.libminiblink != 0 && blink.wkeSetZoomFactor != 0 {
		syscall.Syscall(blink.wkeSetZoomFactor, 2,uintptr(webView),uintptr(factor),0)
	}
}

func (blink *MiniBlinkLib)WkeGetZoomFactor(webView WkeWebView) float32  {
	if blink.libminiblink != 0 && blink.wkeGetZoomFactor != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGetZoomFactor, 2,uintptr(webView),0,0)
		return float32(ret)
	}
	return 1
}

func (blink *MiniBlinkLib)WkeGetString(str WkeString)string  {
	if blink.libminiblink!= 0 && blink.wkeGetStringW!=0{
		ret,_,_ := syscall.Syscall(blink.wkeGetStringW,1,uintptr(str),0,0)
		return DxCommonLib.StringFromUtf16Pointer(ret,1000000000)
	}
	return ""
}

type WkeTitleChanged func(webView WkeWebView,param uintptr,title string)

func wkeTitleChangeCallBack(webview WkeWebView,param uintptr,title uintptr)  {
	if BlinkLib.fOnTitleChanged != nil{
		BlinkLib.fOnTitleChanged(webview,param,BlinkLib.WkeGetString(WkeString(title)))
	}
}

func (blink *MiniBlinkLib)WkeOnTitleChanged(webView WkeWebView,callback WkeTitleChanged,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnTitleChanged != 0 {
		if fTitleChangeCallBack == 0{
			fTitleChangeCallBack = syscall.NewCallbackCDecl(wkeTitleChangeCallBack)
		}
		syscall.Syscall(blink.wkeOnTitleChanged,3,uintptr(webView),fTitleChangeCallBack,param)
	}
}

var (
	BlinkLib	MiniBlinkLib
	fwkeCookieVisitorCallBack	uintptr
	fTitleChangeCallBack		uintptr
)

func init()  {
	BlinkLib.LoadBlink("node.dll")
}
