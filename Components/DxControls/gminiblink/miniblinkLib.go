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
	wkeCreateStringW  uintptr
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
	fOnTitleChanged,fOnMouseOverUrlChanged		WkeTitleChanged
	fOnURLChanged	WkeURLChangedCallback2
	fonPaintUpdatedCallback	WkePaintUpdatedCallback
	fOnPaintBitUpdated	WkePaintBitUpdatedCallback
	fOnAlertBox			WkeAlertBoxCallback
	fOnConfirmBox		WkeConfirmBoxCallback
	fOnPromptBox		WkePromptBoxCallback
	fOnNavigation		WkeNavigationCallback
	fOnCreateView		WkeCreateViewCallback
	fOnDocumentReady	WkeDocumentReadyCallback2
	fOnDownload			WkeDownloadCallback
	fNetOnResponse,fOnLoadUrlBegin		WkeNetResponseCallback
	fOnConsole			WkeConsoleCallback
	fOnLoadUrlEnd		WkeLoadUrlEndCallback
	fOnDidCreateScriptContext	WkeDidCreateScriptContextCallback
	fOnWillReleaseScriptContext	WkeWillReleaseScriptContextCallback
	fOnWillMediaLoad			WkeWillMediaLoadCallback
	fOnWindowClosing			WkeWindowClosingCallback
	fOnWindowDestroy			WkeWindowDestroyCallback
	frunNativefunc				WkeJsNativeFunction
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
	blink.wkeCreateStringW, _ = syscall.GetProcAddress(blink.libminiblink, "wkeCreateStringW")
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

func (blink *MiniBlinkLib)WkeShowWindow(wkeWebView WkeWebView, showFlag bool){
	if blink.libminiblink != 0 && blink.wkeShowWindow != 0{
		syscall.Syscall(blink.wkeShowWindow, 2, uintptr(wkeWebView), uintptr(DxCommonLib.Ord(showFlag)), 0)
	}
}

func (blink *MiniBlinkLib)WkeEnableWindow(wkeWebView WkeWebView, enabled bool){
	if blink.libminiblink != 0 && blink.wkeEnableWindow != 0{
		syscall.Syscall(blink.wkeEnableWindow, 2, uintptr(wkeWebView), uintptr(DxCommonLib.Ord(enabled)), 0)
	}
}

func (blink *MiniBlinkLib)WkeLoadURL(wkeWebView WkeWebView,url string)  {
	if blink.libminiblink != 0 && blink.wkeLoadURL != 0 {
		syscall.Syscall(blink.wkeLoadURL, 2, uintptr(wkeWebView), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(url))), 0)
	}
}

func (blink *MiniBlinkLib)WkeLoadHTML(wkeWebView WkeWebView,html string)  {
	if blink.libminiblink != 0 && blink.wkeLoadHTML != 0 {
		syscall.Syscall(blink.wkeLoadHTML, 2, uintptr(wkeWebView), uintptr(unsafe.Pointer(&html)), 0)
	}
}

func (blink *MiniBlinkLib)WkeLoadHtmlWithBaseUrl(wkeWebView WkeWebView,html,baseUrl string)  {
	if blink.libminiblink != 0 && blink.wkeLoadHtmlWithBaseUrl != 0 {
		syscall.Syscall(blink.wkeLoadHtmlWithBaseUrl, 3, uintptr(wkeWebView), uintptr(unsafe.Pointer(&html)), uintptr(unsafe.Pointer(&baseUrl)))
	}
}

func (blink *MiniBlinkLib)WkeLoadFile(wkeWebView WkeWebView,fileName string)  {
	if blink.libminiblink != 0 && blink.wkeLoadFile != 0 {
		syscall.Syscall(blink.wkeLoadFile, 2, uintptr(wkeWebView), uintptr(unsafe.Pointer(&fileName)), 0)
	}
}

func (blink *MiniBlinkLib)WkeCreateWebWindow(Type uintptr, parent int, x int, y int, width int, height int) uintptr{
	if blink.libminiblink != 0 && blink.wkeCreateWebWindow != 0 {
		ret, _, _ := syscall.Syscall6(blink.wkeCreateWebWindow, 6, Type, uintptr(parent), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
		return ret
	}
	return 0
}

func (blink *MiniBlinkLib)WkeDestroyWebWindow(webview WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeDestroyWebWindow != 0 {
		syscall.Syscall(blink.wkeDestroyWebWindow, 1,uintptr(webview),0, 0)
	}
}

func (blink *MiniBlinkLib)WkeGetWindowHandle(webview WkeWebView)syscall.Handle  {
	if blink.libminiblink != 0 && blink.wkeGetWindowHandle != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeGetWindowHandle, 1,uintptr(webview),0, 0)
		return syscall.Handle(ret)
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

func (blink *MiniBlinkLib)WkeDestroyWebView(webview WkeWebView)  {
	if blink.libminiblink != 0 && blink.wkeDestroyWebView != 0 {
		syscall.Syscall(blink.wkeDestroyWebView, 1,uintptr(webview),0, 0)
	}
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
		blink.fOnTitleChanged = callback
		syscall.Syscall(blink.wkeOnTitleChanged,3,uintptr(webView),fTitleChangeCallBack,param)
	}
}


func wkeOnMouseOverUrlChanged(webview WkeWebView,param uintptr,title uintptr)  {
	if BlinkLib.fOnMouseOverUrlChanged != nil{
		BlinkLib.fOnMouseOverUrlChanged(webview,param,BlinkLib.WkeGetString(WkeString(title)))
	}
}

func (blink *MiniBlinkLib)WkeOnMouseOverUrlChanged(webView WkeWebView,callback WkeTitleChanged,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnMouseOverUrlChanged != 0 {
		if fOnMouseOverUrlChanged == 0{
			fOnMouseOverUrlChanged = syscall.NewCallbackCDecl(wkeOnMouseOverUrlChanged)
		}
		blink.fOnMouseOverUrlChanged = callback
		syscall.Syscall(blink.wkeOnMouseOverUrlChanged,3,uintptr(webView),fOnMouseOverUrlChanged,param)
	}
}


func wkeOnURLChanged(webview WkeWebView,param uintptr,frameid WkeWebFrameHandle,url WkeString)  {
	if BlinkLib.fOnURLChanged != nil{
		BlinkLib.fOnURLChanged(webview,param,frameid,BlinkLib.WkeGetString(url))
	}
}

type WkeURLChangedCallback2	func(webView WkeWebView,param uintptr,frameid WkeWebFrameHandle,url string)

func (blink *MiniBlinkLib)WkeOnURLChanged(webView WkeWebView,callback WkeURLChangedCallback2,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnURLChanged2 != 0 {
		if fOnURLChanged == 0{
			fOnURLChanged = syscall.NewCallbackCDecl(wkeOnURLChanged)
		}
		blink.fOnURLChanged = callback
		syscall.Syscall(blink.wkeOnURLChanged2,3,uintptr(webView),fOnURLChanged,param)
	}
}


type WkePaintUpdatedCallback func(webView WkeWebView,param uintptr,dc WinApi.HDC,x,y,cx,cy int)

func wkePaintUpdatedCallback(webView WkeWebView,param uintptr,dc WinApi.HDC,x,y,cx,cy int)  {
	if BlinkLib.fonPaintUpdatedCallback!=nil{
		BlinkLib.fonPaintUpdatedCallback(webView,param,dc,x,y,cx,cy)
	}
}

func (blink *MiniBlinkLib)WkeOnPaintUpdated(webView WkeWebView,callback WkePaintUpdatedCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnPaintUpdated != 0 {
		if fonPaintUpdatedCallback == 0{
			fonPaintUpdatedCallback = syscall.NewCallbackCDecl(wkePaintUpdatedCallback)
		}
		blink.fonPaintUpdatedCallback = callback
		syscall.Syscall(blink.wkeOnPaintUpdated,3,uintptr(webView),fonPaintUpdatedCallback,param)
	}
}


type WkePaintBitUpdatedCallback func(webview WkeWebView,param uintptr,buffer uintptr,rect *WkeRect,width,height int)

func wkePaintBitUpdatedCallback(webview WkeWebView,param uintptr,buffer uintptr,rect uintptr,width,height int)  {
	if BlinkLib.fOnPaintBitUpdated != nil{
		BlinkLib.fOnPaintBitUpdated(webview,param,buffer,(*WkeRect)(unsafe.Pointer(rect)),width,height)
	}
}

func (blink *MiniBlinkLib)WkeOnPaintBitUpdated(webView WkeWebView,callback WkePaintBitUpdatedCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnPaintBitUpdated != 0 {
		if fOnPaintBitUpdated == 0{
			fOnPaintBitUpdated = syscall.NewCallbackCDecl(wkePaintBitUpdatedCallback)
		}
		blink.fOnPaintBitUpdated = callback
		syscall.Syscall(blink.wkeOnPaintBitUpdated,3,uintptr(webView),fOnPaintBitUpdated,param)
	}
}

type WkeAlertBoxCallback func(webview WkeWebView,param uintptr,msg string)
func wkeAlertBoxCallback(webview WkeWebView,param uintptr,msg WkeString)  {
	if BlinkLib.fOnAlertBox != nil{
		BlinkLib.fOnAlertBox(webview,param,BlinkLib.WkeGetString(msg))
	}
}

func (blink *MiniBlinkLib)WkeOnAlertBox(webView WkeWebView,callback WkeAlertBoxCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnAlertBox != 0 {
		if fOnAlertBox == 0{
			fOnAlertBox = syscall.NewCallbackCDecl(wkeAlertBoxCallback)
		}
		blink.fOnAlertBox = callback
		syscall.Syscall(blink.wkeOnAlertBox,3,uintptr(webView),fOnAlertBox,param)
	}
}

type WkeConfirmBoxCallback func(webview WkeWebView,param uintptr,msg string)bool
func wkeConfirmBoxCallback(webview WkeWebView,param uintptr,msg WkeString)bool  {
	if BlinkLib.fOnConfirmBox!= nil{
		return BlinkLib.fOnConfirmBox(webview,param,BlinkLib.WkeGetString(msg))
	}
	return false
}

func (blink *MiniBlinkLib)WkeOnConfirmBox(webView WkeWebView,callback WkeConfirmBoxCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnConfirmBox != 0 {
		if fOnConfirmBox == 0{
			fOnConfirmBox = syscall.NewCallbackCDecl(wkeConfirmBoxCallback)
		}
		blink.fOnConfirmBox = callback
		syscall.Syscall(blink.wkeOnConfirmBox,3,uintptr(webView),fOnConfirmBox,param)
	}
}


type WkePromptBoxCallback func(webview WkeWebView,param uintptr,msg,defaultresult,sresult string) bool
func wkePromptBoxCallback(webview WkeWebView,param uintptr,msg,defaultResult,sresult WkeString)bool  {
	if BlinkLib.fOnPromptBox != nil{
		return BlinkLib.fOnPromptBox(webview,param,BlinkLib.WkeGetString(msg),BlinkLib.WkeGetString(defaultResult),
			BlinkLib.WkeGetString(sresult))
	}
	return false
}

func (blink *MiniBlinkLib)WkeOnPromptBox(webView WkeWebView,callback WkePromptBoxCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnPromptBox != 0 {
		if fOnPromptBox == 0{
			fOnPromptBox = syscall.NewCallbackCDecl(wkePromptBoxCallback)
		}
		blink.fOnPromptBox = callback
		syscall.Syscall(blink.wkeOnPromptBox,3,uintptr(webView),fOnPromptBox,param)
	}
}

type WkeNavigationCallback func(webview WkeWebView,param uintptr,navigationType WkeNavigationType,url string)bool
func wkeNavigationCallback(webview WkeWebView,param uintptr,navigationType WkeNavigationType,url WkeString)bool  {
	if BlinkLib.fOnNavigation != nil{
		return BlinkLib.fOnNavigation(webview,param,navigationType,BlinkLib.WkeGetString(url))
	}
	return false
}

func (blink *MiniBlinkLib)WkeOnNavigation(webView WkeWebView,callback WkeNavigationCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnNavigation != 0 {
		if fOnNavigation == 0{
			fOnNavigation = syscall.NewCallbackCDecl(wkeNavigationCallback)
		}
		blink.fOnNavigation = callback
		syscall.Syscall(blink.wkeOnNavigation,3,uintptr(webView),fOnNavigation,param)
	}
}

type WkeCreateViewCallback func(webview WkeWebView,param uintptr,navigationType WkeNavigationType,url string,windowFeatures *WkeWindowFeatures) WkeWebView

func wkeCreateViewCallback(webview WkeWebView,param uintptr,navigationType WkeNavigationType,url WkeString,windowFeatures uintptr) uintptr  {
	if BlinkLib.fOnCreateView != nil{
		return uintptr(BlinkLib.fOnCreateView(webview,param,navigationType,BlinkLib.WkeGetString(url),(*WkeWindowFeatures)(unsafe.Pointer(windowFeatures))))
	}
	return 0
}

func (blink *MiniBlinkLib)WkeOnCreateView(webView WkeWebView,callback WkeCreateViewCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnCreateView != 0 {
		if fOnCreateView == 0{
			fOnCreateView = syscall.NewCallbackCDecl(wkeCreateViewCallback)
		}
		blink.fOnCreateView = callback
		syscall.Syscall(blink.wkeOnCreateView,3,uintptr(webView),fOnCreateView,param)
	}
}

type WkeDocumentReadyCallback2 func(webview WkeWebView,param uintptr,frameid WkeFrameHwnd)
func wkeDocumentReadyCallback2(webview WkeWebView,param uintptr,frameid WkeFrameHwnd)  {
	if BlinkLib.fOnDocumentReady != nil{
		BlinkLib.fOnDocumentReady(webview,param,frameid)
	}
}

func (blink *MiniBlinkLib)WkeOnDocumentReady(webView WkeWebView,callback WkeDocumentReadyCallback2,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnDocumentReady2 != 0 {
		if fOnDocumentReady == 0{
			fOnDocumentReady = syscall.NewCallbackCDecl(wkeDocumentReadyCallback2)
		}
		blink.fOnDocumentReady = callback
		syscall.Syscall(blink.wkeOnDocumentReady2,3,uintptr(webView),fOnDocumentReady,param)
	}
}

type  WkeDownloadCallback func(webView WkeWebView,param uintptr,url string) bool
func wkeDownloadCallback(webView WkeWebView,param uintptr,url uintptr)bool  {
	if BlinkLib.fOnDownload != nil{
		return BlinkLib.fOnDownload(webView,param,DxCommonLib.StringFromUtf8Pointer(url,1024))
	}
	return false
}

func (blink *MiniBlinkLib)WkeOnDownload(webView WkeWebView,callback WkeDownloadCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnDownload != 0 {
		if fOnDownload == 0{
			fOnDownload = syscall.NewCallbackCDecl(wkeDownloadCallback)
		}
		blink.fOnDownload = callback
		syscall.Syscall(blink.wkeOnDownload,3,uintptr(webView),fOnDownload,param)
	}
}


type  WkeNetResponseCallback func(webView WkeWebView,param uintptr,url string,job uintptr) bool
func wkeNetResponseCallback(webView WkeWebView,param uintptr,url uintptr,job uintptr)bool  {
	if BlinkLib.fNetOnResponse != nil{
		return BlinkLib.fNetOnResponse(webView,param,DxCommonLib.StringFromUtf8Pointer(url,1024),job)
	}
	return false
}

func (blink *MiniBlinkLib)WkeNetOnResponse(webView WkeWebView,callback WkeNetResponseCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeNetOnResponse != 0 {
		if fNetOnResponse == 0{
			fNetOnResponse = syscall.NewCallbackCDecl(wkeNetResponseCallback)
		}
		blink.fNetOnResponse = callback
		syscall.Syscall(blink.wkeNetOnResponse,3,uintptr(webView),fNetOnResponse,param)
	}
}


type  WkeConsoleCallback func(webView WkeWebView,param uintptr,level WkeConsoleLevel,msg,sourceName string,sourceline uint32,stackTrace string)
func wkeConsoleCallback(webView WkeWebView,param uintptr,level WkeConsoleLevel,msg,sourceName WkeString,sourceLine uint32,stackTrace WkeString)  {
	if BlinkLib.fOnConsole != nil{
		BlinkLib.fOnConsole(webView,param,level,DxCommonLib.StringFromUtf8Pointer(uintptr(msg),102400),
			DxCommonLib.StringFromUtf8Pointer(uintptr(sourceName),1024),sourceLine,DxCommonLib.StringFromUtf8Pointer(uintptr(stackTrace),1024000))
	}
}

func (blink *MiniBlinkLib)WkeOnConsole(webView WkeWebView,callback WkeConsoleCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnConsole != 0 {
		if fOnConsole == 0{
			fOnConsole = syscall.NewCallbackCDecl(wkeConsoleCallback)
		}
		blink.fOnConsole = callback
		syscall.Syscall(blink.wkeOnConsole,3,uintptr(webView),fOnConsole,param)
	}
}


func wkeLoadUrlBeginCallback(webView WkeWebView,param uintptr,url uintptr,job uintptr)bool  {
	if BlinkLib.fOnLoadUrlBegin != nil{
		return BlinkLib.fOnLoadUrlBegin(webView,param,DxCommonLib.StringFromUtf8Pointer(url,1024),job)
	}
	return false
}

func (blink *MiniBlinkLib)WkeOnLoadUrlBegin(webView WkeWebView,callback WkeNetResponseCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnLoadUrlBegin != 0 {
		if fOnLoadUrlBegin == 0{
			fOnLoadUrlBegin = syscall.NewCallbackCDecl(wkeLoadUrlBeginCallback)
		}
		blink.fOnLoadUrlBegin = callback
		syscall.Syscall(blink.wkeOnLoadUrlBegin,3,uintptr(webView),fOnLoadUrlBegin,param)
	}
}


type WkeLoadUrlEndCallback func(webView WkeWebView,param uintptr,url string,job uintptr,buf []byte)
func wkeLoadUrlEndCallback(webView WkeWebView,param uintptr,url uintptr,job uintptr,buf uintptr,len uintptr)  {
	if BlinkLib.fOnLoadUrlEnd != nil{
		buffer := make([]byte,len)
		DxCommonLib.CopyMemory(unsafe.Pointer(&buffer[0]),unsafe.Pointer(buf),len)
		BlinkLib.fOnLoadUrlEnd(webView,param,DxCommonLib.StringFromUtf8Pointer(url,1024),job,buffer)
	}
}

func (blink *MiniBlinkLib)WkeOnLoadUrlEnd(webView WkeWebView,callback WkeLoadUrlEndCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnLoadUrlEnd != 0 {
		if fOnLoadUrlEnd == 0{
			fOnLoadUrlEnd = syscall.NewCallbackCDecl(wkeLoadUrlEndCallback)
		}
		blink.fOnLoadUrlEnd = callback
		syscall.Syscall(blink.wkeOnLoadUrlEnd,3,uintptr(webView),fOnLoadUrlEnd,param)
	}
}

type WkeDidCreateScriptContextCallback func(webView WkeWebView,param uintptr,frameId WkeWebFrameHandle,context uintptr,extensionGroup int,worldId int)
func wkeDidCreateScriptContextCallback(webView WkeWebView,param uintptr,frameId WkeWebFrameHandle,context uintptr,extensionGroup int,worldId int)  {
	if BlinkLib.fOnDidCreateScriptContext != nil{
		BlinkLib.fOnDidCreateScriptContext(webView,param,frameId,context,extensionGroup,worldId)
	}
}

func (blink *MiniBlinkLib)WkeOnDidCreateScriptContext(webView WkeWebView,callback WkeDidCreateScriptContextCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnDidCreateScriptContext != 0 {
		if fOnDidCreateScriptContext == 0{
			fOnDidCreateScriptContext = syscall.NewCallbackCDecl(wkeDidCreateScriptContextCallback)
		}
		blink.fOnDidCreateScriptContext = callback
		syscall.Syscall(blink.wkeOnDidCreateScriptContext,3,uintptr(webView),fOnDidCreateScriptContext,param)
	}
}

type WkeWillReleaseScriptContextCallback func(webView WkeWebView,param uintptr,frameId WkeWebFrameHandle,context uintptr,worldId int)
func wkeWillReleaseScriptContextCallback(webView WkeWebView,param uintptr,frameId WkeWebFrameHandle,context uintptr,worldId int)  {
	if BlinkLib.fOnWillReleaseScriptContext != nil{
		BlinkLib.fOnWillReleaseScriptContext(webView,param,frameId,context,worldId)
	}
}

func (blink *MiniBlinkLib)WkeOnWillReleaseScriptContext(webView WkeWebView,callback WkeWillReleaseScriptContextCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnWillReleaseScriptContext != 0 {
		if fOnWillReleaseScriptContext == 0{
			fOnWillReleaseScriptContext = syscall.NewCallbackCDecl(wkeWillReleaseScriptContextCallback)
		}
		blink.fOnWillReleaseScriptContext = callback
		syscall.Syscall(blink.wkeOnWillReleaseScriptContext,3,uintptr(webView),fOnWillReleaseScriptContext,param)
	}
}

type WkeWillMediaLoadCallback func(webView WkeWebView,param uintptr,url string,info *WkeMediaLoadInfo)
func wkeWillMediaLoadCallback(webView WkeWebView,param uintptr,url uintptr,info uintptr)  {
	if BlinkLib.fOnWillMediaLoad != nil{
		BlinkLib.fOnWillMediaLoad(webView,param,DxCommonLib.StringFromUtf8Pointer(url,1024),(*WkeMediaLoadInfo)(unsafe.Pointer(info)))
	}
}

func (blink *MiniBlinkLib)WkeOnWillMediaLoad(webView WkeWebView,callback WkeWillMediaLoadCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnWillMediaLoad != 0 {
		if fOnWillMediaLoad == 0{
			fOnWillMediaLoad = syscall.NewCallbackCDecl(wkeWillMediaLoadCallback)
		}
		blink.fOnWillMediaLoad = callback
		syscall.Syscall(blink.wkeOnWillReleaseScriptContext,3,uintptr(webView),fOnWillMediaLoad,param)
	}
}


func (blink *MiniBlinkLib)WkeIsMainFrame(webView WkeWebView,frameid WkeWebFrameHandle)bool  {
	if blink.libminiblink != 0 && blink.wkeIsMainFrame != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeIsMainFrame,2,uintptr(webView),uintptr(frameid),0)
		return  ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)WkeWebFrameGetMainFrame(webView WkeWebView)WkeWebFrameHandle  {
	if blink.libminiblink != 0 && blink.wkeWebFrameGetMainFrame != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeWebFrameGetMainFrame,1,uintptr(webView),0,0)
		return  WkeWebFrameHandle(ret)
	}
	return 0
}

func (blink *MiniBlinkLib)WkeRunJsByFrame(webView WkeWebView,frameId WkeWebFrameHandle,script string,isInClosure bool)JSValue  {
	if blink.libminiblink != 0 && blink.wkeRunJsByFrame != 0 {
		ret,_,_ := syscall.Syscall6(blink.wkeRunJsByFrame,4,uintptr(webView),uintptr(frameId),uintptr(unsafe.Pointer(&script)),
			uintptr(DxCommonLib.Ord(isInClosure)),0,0)
		return  JSValue(ret)
	}
	return 0
}


func (blink *MiniBlinkLib)WkeGetFrameUrl(webView WkeWebView,frameId WkeWebFrameHandle)string  {
	if blink.libminiblink != 0 && blink.wkeGetFrameUrl != 0{
		ret,_,_ := syscall.Syscall(blink.wkeGetFrameUrl,2,uintptr(webView),uintptr(frameId),0)
		return DxCommonLib.StringFromUtf8Pointer(ret,255)
	}
	return ""
}

func (blink *MiniBlinkLib)WkeSetString(wkString WkeString,value string)  {
	if blink.libminiblink != 0 && blink.wkeSetString != 0{
		bt := DxCommonLib.FastString2Byte(value)
		syscall.Syscall(blink.wkeSetString,3,uintptr(wkString),uintptr(unsafe.Pointer(&bt[0])),uintptr(len(bt)))
	}
}

func (blink *MiniBlinkLib)WkeSetStringW(wkString WkeString,value string)  {
	if blink.libminiblink != 0 && blink.wkeSetStringW != 0{
		utf16arr,_ := syscall.UTF16FromString(value)
		syscall.Syscall(blink.wkeSetStringW,3,uintptr(wkString),uintptr(unsafe.Pointer(&utf16arr[0])),uintptr(len(utf16arr)))
	}
}

func (blink *MiniBlinkLib)WkeCreateString(value string)WkeString  {
	if blink.libminiblink != 0 && blink.wkeCreateStringW != 0{
		utf16arr,_ := syscall.UTF16FromString(value)
		ret,_,_ := syscall.Syscall(blink.wkeCreateStringW,2,uintptr(unsafe.Pointer(&utf16arr[0])),uintptr(len(utf16arr)),0)
		return WkeString(ret)
	}
	return 0
}

func (blink *MiniBlinkLib)WkeDeleteString(wkString WkeString)  {
	if blink.libminiblink != 0 && blink.wkeDeleteString != 0{
		syscall.Syscall(blink.wkeDeleteString,1,uintptr(wkString),0,0)
	}
}

func (blink *MiniBlinkLib)WkeSetUserKeyValue(webView WkeWebView,key string,value uintptr)  {
	if blink.libminiblink != 0 && blink.wkeSetUserKeyValue != 0{
		syscall.Syscall(blink.wkeSetUserKeyValue,3,uintptr(webView),uintptr(unsafe.Pointer(&key)),value)
	}
}

func (blink *MiniBlinkLib)WkeGetUserKeyValue(webView WkeWebView,key string)uintptr  {
	if blink.libminiblink != 0 && blink.wkeGetUserKeyValue != 0{
		ret,_,_ := syscall.Syscall(blink.wkeGetUserKeyValue,2,uintptr(webView),uintptr(unsafe.Pointer(&key)),0)
		return ret
	}
	return 0
}

func (blink *MiniBlinkLib)WkeGetCursorInfoType(webView WkeWebView)WkeCursorInfoType  {
	if blink.libminiblink != 0 && blink.wkeGetCursorInfoType != 0{
		ret,_,_ := syscall.Syscall(blink.wkeGetCursorInfoType,1,uintptr(webView),0,0)
		return WkeCursorInfoType(ret)
	}
	return WkeCursorInfoPointer
}

type WkeWindowClosingCallback func(webView WkeWebView,param uintptr)bool
type WkeWindowDestroyCallback func(webView WkeWebView,param uintptr)
func wkeWindowClosingCallback(webView WkeWebView,param uintptr)bool  {
	if BlinkLib.fOnWindowClosing != nil{
		return  BlinkLib.fOnWindowClosing(webView,param)
	}
	return false
}

func wkeWindowDestroyCallback(webView WkeWebView,param uintptr)  {
	if BlinkLib.fOnWindowDestroy != nil{
		BlinkLib.fOnWindowDestroy(webView,param)
	}
}

func (blink *MiniBlinkLib)WkeOnWindowClosing(webView WkeWebView,callback WkeWindowClosingCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnWindowClosing != 0 {
		if fOnWindowClosing == 0{
			fOnWindowClosing = syscall.NewCallbackCDecl(wkeWindowClosingCallback)
		}
		blink.fOnWindowClosing = callback
		syscall.Syscall(blink.wkeOnWindowClosing,3,uintptr(webView),fOnWindowClosing,param)
	}
}

func (blink *MiniBlinkLib)WkeOnWindowDestroy(webView WkeWebView,callback WkeWindowDestroyCallback,param uintptr)  {
	if blink.libminiblink != 0 && blink.wkeOnWindowDestroy != 0 {
		if fOnWindowDestroy == 0{
			fOnWindowDestroy = syscall.NewCallbackCDecl(wkeWindowDestroyCallback)
		}
		blink.fOnWindowDestroy = callback
		syscall.Syscall(blink.wkeOnWindowDestroy,3,uintptr(webView),fOnWindowDestroy,param)
	}
}

func (blink *MiniBlinkLib)WkeMoveWindow(webView WkeWebView,x,y,w,h int){
	if blink.libminiblink != 0 && blink.wkeMoveWindow != 0 {
		syscall.Syscall6(blink.wkeMoveWindow,5,uintptr(webView),uintptr(x),uintptr(y),uintptr(w),uintptr(h),0)
	}
}

func (blink *MiniBlinkLib)WkeMoveToCenter(webView WkeWebView){
	if blink.libminiblink != 0 && blink.wkeMoveToCenter != 0 {
		syscall.Syscall(blink.wkeMoveToCenter,1,uintptr(webView),0,0)
	}
}

func (blink *MiniBlinkLib)WkeResizeWindow(webView WkeWebView,w,h int){
	if blink.libminiblink != 0 && blink.wkeResizeWindow != 0 {
		syscall.Syscall(blink.wkeResizeWindow,1,uintptr(webView),uintptr(w),uintptr(h))
	}
}

func (blink *MiniBlinkLib)WkeSetWindowTitle(webView WkeWebView,title string){
	if blink.libminiblink != 0 && blink.wkeSetWindowTitle != 0 {
		syscall.Syscall(blink.wkeSetWindowTitle,2,uintptr(webView),uintptr(unsafe.Pointer(&title)),0)
	}
}

func (blink *MiniBlinkLib)WkeSetDeviceParameter(webView WkeWebView,device,paramStr string,paramInt int,paramFloat float32){
	if blink.libminiblink != 0 && blink.wkeSetDeviceParameter != 0 {
		syscall.Syscall6(blink.wkeSetDeviceParameter,5,uintptr(webView),uintptr(unsafe.Pointer(&device)),uintptr(unsafe.Pointer(&paramStr)),
			uintptr(paramInt),uintptr(paramFloat),0)
	}
}

func (blink *MiniBlinkLib)WkeSetProxy(proxy *WkeProxy){
	if blink.libminiblink != 0 && blink.wkeSetProxy != 0 {
		syscall.Syscall(blink.wkeSetProxy,1,uintptr(unsafe.Pointer(proxy)),0,0)
	}
}

func (blink *MiniBlinkLib)WkeSetViewProxy(webView WkeWebView, proxy *WkeProxy){
	if blink.libminiblink != 0 && blink.wkeSetViewProxy != 0 {
		syscall.Syscall(blink.wkeSetViewProxy,2,uintptr(webView),uintptr(unsafe.Pointer(proxy)),0)
	}
}

func (blink *MiniBlinkLib)WkeConfigure(settings *WkeSettings){
	if blink.libminiblink != 0 && blink.wkeConfigure != 0 {
		syscall.Syscall(blink.wkeConfigure,1,uintptr(unsafe.Pointer(settings)),0,0)
	}
}

func (blink *MiniBlinkLib)WkeIsInitialize()bool  {
	if blink.libminiblink != 0 && blink.wkeIsInitialize != 0 {
		ret,_,_ := syscall.Syscall(blink.wkeIsInitialize,0,0,0,0)
		return ret != 0
	}
	return false
}


func (blink *MiniBlinkLib)WkeSetMemoryCacheEnable(webView WkeWebView, b bool){
	if blink.libminiblink != 0 && blink.wkeSetMemoryCacheEnable != 0{
		syscall.Syscall(blink.wkeSetMemoryCacheEnable, 2, uintptr(webView), uintptr(DxCommonLib.Ord(b)), 0)
	}
}

func (blink *MiniBlinkLib)WkeSetTouchEnabled(webView WkeWebView, b bool){
	if blink.libminiblink != 0 && blink.wkeSetTouchEnabled != 0{
		syscall.Syscall(blink.wkeSetTouchEnabled, 2, uintptr(webView), uintptr(DxCommonLib.Ord(b)), 0)
	}
}

func (blink *MiniBlinkLib)WkeSetMouseEnabled(webView WkeWebView, b bool){
	if blink.libminiblink != 0 && blink.wkeSetMouseEnabled != 0{
		syscall.Syscall(blink.wkeSetMouseEnabled, 2, uintptr(webView), uintptr(DxCommonLib.Ord(b)), 0)
	}
}

func (blink *MiniBlinkLib)WkeSetNavigationToNewWindowEnable(webView WkeWebView, b bool){
	if blink.libminiblink != 0 && blink.wkeSetNavigationToNewWindowEnable != 0{
		syscall.Syscall(blink.wkeSetNavigationToNewWindowEnable, 2, uintptr(webView), uintptr(DxCommonLib.Ord(b)), 0)
	}
}

func (blink *MiniBlinkLib)WkeSetCspCheckEnable(webView WkeWebView, b bool){
	if blink.libminiblink != 0 && blink.wkeSetCspCheckEnable != 0{
		syscall.Syscall(blink.wkeSetCspCheckEnable, 2, uintptr(webView), uintptr(DxCommonLib.Ord(b)), 0)
	}
}


func (blink *MiniBlinkLib)WkeSetNpapiPluginsEnabled(webView WkeWebView, b bool){
	if blink.libminiblink != 0 && blink.wkeSetNpapiPluginsEnabled != 0{
		syscall.Syscall(blink.wkeSetNpapiPluginsEnabled, 2, uintptr(webView), uintptr(DxCommonLib.Ord(b)), 0)
	}
}

func (blink *MiniBlinkLib)WkeSetHeadlessEnabled(webView WkeWebView, b bool){
	if blink.libminiblink != 0 && blink.wkeSetHeadlessEnabled != 0{
		syscall.Syscall(blink.wkeSetHeadlessEnabled, 2, uintptr(webView), uintptr(DxCommonLib.Ord(b)), 0)
	}
}

func (blink *MiniBlinkLib)WkeSetDebugConfig(webView WkeWebView, debugString,param string){
	if blink.libminiblink != 0 && blink.wkeSetDebugConfig != 0{
		syscall.Syscall(blink.wkeSetDebugConfig, 3, uintptr(webView), uintptr(unsafe.Pointer(&debugString)), uintptr(unsafe.Pointer(&param)))
	}
}

func (blink *MiniBlinkLib)WkeSetHandle(webView WkeWebView,wnd syscall.Handle){
	if blink.libminiblink != 0 && blink.wkeSetHandle != 0{
		syscall.Syscall(blink.wkeSetHandle, 2, uintptr(webView), uintptr(wnd), 0)
	}
}

func (blink *MiniBlinkLib)WkeSetHandleOffset(webView WkeWebView,x,y int){
	if blink.libminiblink != 0 && blink.wkeSetHandleOffset != 0{
		syscall.Syscall(blink.wkeSetHandleOffset, 2, uintptr(webView), uintptr(x), uintptr(y))
	}
}

func (blink *MiniBlinkLib)WkeSetViewSettings(webView WkeWebView, settings *WkeViewSettings){
	if blink.libminiblink != 0 && blink.wkeSetViewSettings != 0{
		syscall.Syscall(blink.wkeSetViewSettings, 2, uintptr(webView), uintptr(unsafe.Pointer(settings)), 0)
	}
}

func (blink *MiniBlinkLib)WkeSetTransparent(webView WkeWebView, b bool){
	if blink.libminiblink != 0 && blink.wkeSetTransparent != 0{
		syscall.Syscall(blink.wkeSetTransparent, 2, uintptr(webView), uintptr(DxCommonLib.Ord(b)), 0)
	}
}

func (blink *MiniBlinkLib)WkeIsTransparent(webView WkeWebView) bool{
	if blink.libminiblink != 0 && blink.wkeIsTransparent != 0{
		ret,_,_ := syscall.Syscall(blink.wkeIsTransparent, 1, uintptr(webView), 0, 0)
		return ret!=0
	}
	return false
}

func (blink *MiniBlinkLib)WkeSetUserAgent(webView WkeWebView, userAgent string){
	if blink.libminiblink != 0 && blink.wkeSetUserAgent != 0{
		syscall.Syscall(blink.wkeSetUserAgent, 2, uintptr(webView), uintptr(unsafe.Pointer(&userAgent)), 0)
	}
}

func (blink *MiniBlinkLib)WkeGetUserAgent(webView WkeWebView) string{
	if blink.libminiblink != 0 && blink.wkeGetUserAgent != 0{
		ret,_,_ := syscall.Syscall(blink.wkeGetUserAgent, 1, uintptr(webView), 0, 0)
		return DxCommonLib.StringFromUtf8Pointer(ret,1024000)
	}
	return ""
}

func (blink *MiniBlinkLib)WkeGetURL(webView WkeWebView) string{
	if blink.libminiblink != 0 && blink.wkeGetURL != 0{
		ret,_,_ := syscall.Syscall(blink.wkeGetURL, 1, uintptr(webView), 0, 0)
		return DxCommonLib.StringFromUtf8Pointer(ret,1024000)
	}
	return ""
}

func (blink *MiniBlinkLib)WkeNetSetHTTPHeaderField(jobPtr uintptr,key,value string,response bool){
	if blink.libminiblink != 0 && blink.wkeNetSetHTTPHeaderField != 0{
		syscall.Syscall6(blink.wkeNetSetHTTPHeaderField,4,jobPtr,uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(key))),
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(value))),uintptr(DxCommonLib.Ord(response)),0,0)
	}
}

func (blink *MiniBlinkLib)WkeNetSetMIMEType(jobPtr uintptr,mimetype string){
	if blink.libminiblink != 0 && blink.wkeNetSetMIMEType != 0{
		syscall.Syscall(blink.wkeNetSetMIMEType,2,jobPtr,uintptr(unsafe.Pointer(&mimetype)),0)
	}
}

func (blink *MiniBlinkLib)WkeNetSetData(jobPtr uintptr,buf []byte) {
	if blink.libminiblink != 0 && blink.wkeNetSetData != 0{
		syscall.Syscall(blink.wkeNetSetData,3,jobPtr,uintptr(unsafe.Pointer(&buf[0])),uintptr(len(buf)))
	}
}

func (blink *MiniBlinkLib)WkeNetCancelRequest(jobPtr uintptr)  {
	if blink.libminiblink != 0 && blink.wkeNetCancelRequest != 0{
		syscall.Syscall(blink.wkeNetSetData,1,jobPtr,0,0)
	}
}

func (blink *MiniBlinkLib)JsArgCount(es JSExecState)int {
	if blink.libminiblink != 0 && blink.jsArgCount != 0{
		ret,_,_ := syscall.Syscall(blink.jsArgCount,1,uintptr(es),0,0)
		return int(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsArgType(es JSExecState,argIdx int)JSType {
	if blink.libminiblink != 0 && blink.jsArgType != 0{
		ret,_,_ := syscall.Syscall(blink.jsArgType,2,uintptr(es),uintptr(argIdx),0)
		return JSType(ret)
	}
	return  JSTYPE_UNDEFINED
}

func (blink *MiniBlinkLib)JsArg(es JSExecState,argIdx int)JSValue {
	if blink.libminiblink != 0 && blink.jsArg != 0{
		ret,_,_ := syscall.Syscall(blink.jsArg,2,uintptr(es),uintptr(argIdx),0)
		return JSValue(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsTypeOf(v JSValue)JSType  {
	if blink.libminiblink != 0 && blink.jsTypeOf != 0{
		ret,_,_ := syscall.Syscall(blink.jsTypeOf,1,uintptr(v),0,0)
		return JSType(ret)
	}
	return JSTYPE_UNDEFINED
}

func (blink *MiniBlinkLib)JsIsNumber(v JSValue)bool  {
	if blink.libminiblink != 0 && blink.jsIsNumber != 0{
		ret,_,_ := syscall.Syscall(blink.jsIsNumber,1,uintptr(v),0,0)
		return ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)JsIsString(v JSValue)bool  {
	if blink.libminiblink != 0 && blink.jsIsString != 0{
		ret,_,_ := syscall.Syscall(blink.jsIsString,1,uintptr(v),0,0)
		return ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)JsIsBoolean(v JSValue)bool  {
	if blink.libminiblink != 0 && blink.jsIsBoolean != 0{
		ret,_,_ := syscall.Syscall(blink.jsIsBoolean,1,uintptr(v),0,0)
		return ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)JsIsObject(v JSValue)bool  {
	if blink.libminiblink != 0 && blink.jsIsObject != 0{
		ret,_,_ := syscall.Syscall(blink.jsIsObject,1,uintptr(v),0,0)
		return ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)JsIsTrue(v JSValue)bool  {
	if blink.libminiblink != 0 && blink.jsIsTrue != 0{
		ret,_,_ := syscall.Syscall(blink.jsIsTrue,1,uintptr(v),0,0)
		return ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)JsIsFalse(v JSValue)bool  {
	if blink.libminiblink != 0 && blink.jsIsFalse != 0{
		ret,_,_ := syscall.Syscall(blink.jsIsFalse,1,uintptr(v),0,0)
		return ret != 0
	}
	return false
}

func (blink *MiniBlinkLib)JsToInt(es JSExecState,v JSValue)int {
	if blink.libminiblink != 0 && blink.jsToInt != 0{
		ret,_,_ := syscall.Syscall(blink.jsToInt,2,uintptr(es),uintptr(v),0)
		return int(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsToDouble(es JSExecState,v JSValue)float64 {
	if blink.libminiblink != 0 && blink.jsToDouble != 0{
		ret,_,_ := syscall.Syscall(blink.jsToDouble,2,uintptr(es),uintptr(v),0)
		return float64(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsToTempString(es JSExecState,v JSValue)string {
	if blink.libminiblink != 0 && blink.jsToTempString != 0{
		ret,_,_ := syscall.Syscall(blink.jsToTempString,2,uintptr(es),uintptr(v),0)
		return DxCommonLib.StringFromUtf8Pointer(ret,1024000)
	}
	return  ""
}

func (blink *MiniBlinkLib)JsToString(es JSExecState,v JSValue)string {
	if blink.libminiblink != 0 && blink.jsToString != 0{
		ret,_,_ := syscall.Syscall(blink.jsToString,2,uintptr(es),uintptr(v),0)
		return DxCommonLib.StringFromUtf8Pointer(ret,1024000)
	}
	return  ""
}

func (blink *MiniBlinkLib)JsInt(v int)JSValue {
	if blink.libminiblink != 0 && blink.jsInt != 0{
		ret,_,_ := syscall.Syscall(blink.jsInt,1,uintptr(v),0,0)
		return JSValue(ret)
	}
	return  0
}


func (blink *MiniBlinkLib)JsString(es JSExecState, v string)JSValue {
	if blink.libminiblink != 0 && blink.jsString != 0{
		ret,_,_ := syscall.Syscall(blink.jsString,2,uintptr(es),uintptr(unsafe.Pointer(&v)),0)
		return JSValue(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsArrayBuffer(es JSExecState, v []byte)JSValue {
	if blink.libminiblink != 0 && blink.jsArrayBuffer != 0{
		ret,_,_ := syscall.Syscall(blink.jsArrayBuffer,3,uintptr(es),uintptr(unsafe.Pointer(&v[0])),uintptr(len(v)))
		return JSValue(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsGetArrayBuffer(es JSExecState,v JSValue)*WkeMemBuf  {
	if blink.libminiblink != 0 && blink.jsGetArrayBuffer != 0{
		ret,_,_ := syscall.Syscall(blink.jsGetArrayBuffer,2,uintptr(es),uintptr(v),0)
		return (*WkeMemBuf)(unsafe.Pointer(ret))
	}
	return  nil
}

func (blink *MiniBlinkLib)JsEmptyObject(es JSExecState)JSValue {
	if blink.libminiblink != 0 && blink.jsEmptyObject != 0{
		ret,_,_ := syscall.Syscall(blink.jsEmptyObject,1,uintptr(es),0,0)
		return JSValue(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsEvalW(es JSExecState,jsScript string)JSValue {
	if blink.libminiblink != 0 && blink.jsEvalW != 0{
		ret,_,_ := syscall.Syscall(blink.jsEvalW,2,uintptr(es),uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(jsScript))),0)
		return JSValue(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsEvalExW(es JSExecState,jsScript string,isInClosure bool)JSValue {
	if blink.libminiblink != 0 && blink.jsEvalExW != 0{
		ret,_,_ := syscall.Syscall(blink.jsEvalExW,3,uintptr(es),uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(jsScript))),uintptr(DxCommonLib.Ord(isInClosure)))
		return JSValue(ret)
	}
	return  0
}


func (blink *MiniBlinkLib)JsCall(es JSExecState,jsfunc,thisValue JSValue,args []JSValue,argCount int)JSValue {
	if blink.libminiblink != 0 && blink.jsCall != 0{
		if len(args) == 0{
			ret,_,_ := syscall.Syscall6(blink.jsCall,5,uintptr(es),uintptr(jsfunc),uintptr(thisValue),0,0,0)
			return JSValue(ret)
		}
		ret,_,_ := syscall.Syscall6(blink.jsCall,5,uintptr(es),uintptr(jsfunc),uintptr(thisValue),uintptr(unsafe.Pointer(&args[0])),uintptr(argCount),0)
		return JSValue(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsCallGlobal(es JSExecState,jsfunc JSValue,args []JSValue,argCount int)JSValue {
	if blink.libminiblink != 0 && blink.jsCallGlobal != 0{
		if len(args) == 0{
			ret,_,_ := syscall.Syscall6(blink.jsCallGlobal,4,uintptr(es),uintptr(jsfunc),0,0,0,0)
			return JSValue(ret)
		}
		ret,_,_ := syscall.Syscall6(blink.jsCallGlobal,4,uintptr(es),uintptr(jsfunc),uintptr(unsafe.Pointer(&args[0])),uintptr(argCount),0,0)
		return JSValue(ret)
	}
	return  0
}

func (blink *MiniBlinkLib)JsGetWebView(es JSExecState) WkeWebView {
	if blink.libminiblink != 0 && blink.jsGetWebView != 0{
		ret,_,_ := syscall.Syscall(blink.jsGetWebView,1,uintptr(es),0,0)
		return WkeWebView(ret)
	}
	return 0
}


type WkeJsNativeFunction func(es JSExecState,param uintptr) JSValue
func wrapJsFuncs(es JSExecState,param uintptr)JSValue  {
	if BlinkLib.frunNativefunc != nil{
		return BlinkLib.frunNativefunc(es,param)
	}
	return 0
}

func (blink *MiniBlinkLib)WkeJsBindFunction(funcName string,nativeFunction WkeJsNativeFunction,param uintptr,argCount uint)  {
	if blink.libminiblink != 0 && blink.wkeJsBindFunction != 0{
		if frunNativefunc == 0{
			frunNativefunc = syscall.NewCallbackCDecl(wrapJsFuncs)
		}
		blink.frunNativefunc = nativeFunction
		syscall.Syscall6(blink.wkeJsBindFunction,4,uintptr(unsafe.Pointer(&funcName)),frunNativefunc,param,uintptr(argCount),0,0)
	}
}

func (blink *MiniBlinkLib)JsGetLastErrorIfException(es JSExecState)*JsExceptionInfo  {
	if blink.libminiblink != 0 && blink.jsGetLastErrorIfException != 0{
		ret,_,_ := syscall.Syscall(blink.jsGetLastErrorIfException,1,uintptr(es),0,0)
		return (*JsExceptionInfo)(unsafe.Pointer(ret))
	}
	return nil
}

var (
	BlinkLib	MiniBlinkLib
	fwkeCookieVisitorCallBack	uintptr
	fTitleChangeCallBack		uintptr
	fOnMouseOverUrlChanged		uintptr
	fOnURLChanged				uintptr
	fonPaintUpdatedCallback		uintptr
	fOnPaintBitUpdated			uintptr
	fOnAlertBox					uintptr
	fOnConfirmBox				uintptr
	fOnPromptBox				uintptr
	fOnNavigation				uintptr
	fOnCreateView				uintptr
	fOnDocumentReady			uintptr
	fOnDownload					uintptr
	fNetOnResponse				uintptr
	fOnConsole					uintptr
	fOnLoadUrlBegin				uintptr
	fOnLoadUrlEnd				uintptr
	fOnDidCreateScriptContext	uintptr
	fOnWillReleaseScriptContext	uintptr
	fOnWillMediaLoad			uintptr
	fOnWindowClosing			uintptr
	fOnWindowDestroy			uintptr
	frunNativefunc				uintptr
)

func init()  {
	BlinkLib.LoadBlink("node.dll")
}
