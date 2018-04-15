package WinApi

import (
	"syscall"
	"unsafe"
	"unicode/utf16"
)



var (
	ScreenLogPixels			     int32
	libuser32                            syscall.Handle
	fnActivateKeyboardLayout             uintptr
	fnAdjustWindowRect                   uintptr
	fnAdjustWindowRectEx                 uintptr
	fnAnyPopup                           uintptr
	fnAppendMenuW                        uintptr
	fnAppendMenuA                        uintptr
	fnArrangeIconicWindows               uintptr
	fnAttachThreadInput                  uintptr
	fnBeginDeferWindowPos                uintptr
	fnBeginPaint                         uintptr
	fnBringWindowToTop                   uintptr
	fnBroadcastSystemMessageW            uintptr
	fnBroadcastSystemMessageA            uintptr
	fnCallMsgFilterW                     uintptr
	fnCallMsgFilterA                     uintptr
	fnCallNextHookEx                     uintptr
	fnCallWindowProcW                    uintptr
	fnLoadCursor                         uintptr
	fnCallWindowProcA                    uintptr
	fnCascadeWindows                     uintptr
	fnChangeWindowMessageFilter          uintptr
	fnChangeClipboardChain               uintptr
	fnChangeDisplaySettingsW             uintptr
	fnChangeDisplaySettingsA             uintptr
	fnChangeDisplaySettingsExW           uintptr
	fnChangeDisplaySettingsExA           uintptr
	fnChangeMenuW                        uintptr
	fnChangeMenuA                        uintptr
	fnCheckDlgButton                     uintptr
	fnCheckMenuItem                      uintptr
	fnCheckMenuRadioItem                 uintptr
	fnCheckRadioButton                   uintptr
	fnChildWindowFromPoint               uintptr
	fnChildWindowFromPointEx             uintptr
	fnClientToScreen                     uintptr
	fnClipCursor                         uintptr
	fnCloseClipboard                     uintptr
	fnCloseDesktop                       uintptr
	fnCloseTouchInputHandle              uintptr
	fnCloseWindow                        uintptr
	fnCloseWindowStation                 uintptr
	fnCopyAcceleratorTableW              uintptr
	fnCopyAcceleratorTableA              uintptr
	fnCopyIcon                           uintptr
	fnCopyImage                          uintptr
	fnCopyRect                           uintptr
	fnCountClipboardFormats              uintptr
	fnCreateAcceleratorTableW            uintptr
	fnCreateAcceleratorTableA            uintptr
	fnCreateCaret                        uintptr
	fnCreateCursor                       uintptr
	fnCreateDesktopW                     uintptr
	fnCreateDesktopA                     uintptr
	fnCreateDesktopExW                   uintptr
	fnCreateDesktopExA                   uintptr
	fnCreateDialogIndirectParamW         uintptr
	fnCreateDialogIndirectParamA         uintptr
	fnCreateDialogParamW                 uintptr
	fnCreateDialogParamA                 uintptr
	fnCreateIcon                         uintptr
	fnCreateIconFromResource             uintptr
	fnCreateIconFromResourceEx           uintptr
	fnCreateIconIndirect                 uintptr
	fnCreateMDIWindowW                   uintptr
	fnCreateMDIWindowA                   uintptr
	fnCreateMenu                         uintptr
	fnCreatePopupMenu                    uintptr
	fnCreateWindowStationW               uintptr
	fnCreateWindowStationA               uintptr
	fnDdeSetQualityOfService             uintptr
	fnDefDlgProcW                        uintptr
	fnDefDlgProcA                        uintptr
	fnDefFrameProcW                      uintptr
	fnDefFrameProcA                      uintptr
	fnDefMDIChildProcW                   uintptr
	fnDefMDIChildProcA                   uintptr
	fnDefWindowProcW                     uintptr
	fnDefWindowProcA                     uintptr
	fnDeferWindowPos                     uintptr
	fnDeleteMenu                         uintptr
	fnDestroyAcceleratorTable            uintptr
	fnDestroyCaret                       uintptr
	fnDestroyCursor                      uintptr
	fnDestroyIcon                        uintptr
	fnDestroyMenu                        uintptr
	fnDestroyWindow                      uintptr
	fnDialogBoxIndirectParamW            uintptr
	fnDialogBoxIndirectParamA            uintptr
	fnDialogBoxParamW                    uintptr
	fnDialogBoxParamA                    uintptr
	fnDlgDirListW                        uintptr
	fnDlgDirListA                        uintptr
	fnDlgDirListComboBoxW                uintptr
	fnDlgDirListComboBoxA                uintptr
	fnDlgDirSelectComboBoxExW            uintptr
	fnDlgDirSelectComboBoxExA            uintptr
	fnDlgDirSelectExW                    uintptr
	fnDlgDirSelectExA                    uintptr
	fnDragDetect                         uintptr
	fnDragObject                         uintptr
	fnDrawAnimatedRects                  uintptr
	fnDrawCaption                        uintptr
	fnDrawEdge                           uintptr
	fnDrawFocusRect                      uintptr
	fnDrawFrameControl                   uintptr
	fnDrawIcon                           uintptr
	fnDrawIconEx                         uintptr
	fnDrawMenuBar                        uintptr
	fnDrawStateW                         uintptr
	fnDrawStateA                         uintptr
	fnDrawTextExW                        uintptr
	fnDrawTextExA                        uintptr
	fnDrawTextW                          uintptr
	fnEmptyClipboard                     uintptr
	fnEnableMenuItem                     uintptr
	fnEnableScrollBar                    uintptr
	fnSystemParametersInfo				 uintptr
	fnEnableWindow                       uintptr
	fnEndDeferWindowPos                  uintptr
	fnEndDialog                          uintptr
	fnEndMenu                            uintptr
	fnEndPaint                           uintptr
	fnEnumChildWindows                   uintptr
	fnEnumClipboardFormats               uintptr
	fnEnumDesktopsW                      uintptr
	fnEnumDesktopsA                      uintptr
	fnEnumDesktopWindows                 uintptr
	fnEnumDisplaySettingsW               uintptr
	fnEnumDisplaySettingsA               uintptr
	fnEnumDisplayDevicesW                uintptr
	fnEnumDisplayDevicesA                uintptr
	fnEnumPropsW                         uintptr
	fnEnumPropsA                         uintptr
	fnEnumPropsExW                       uintptr
	fnEnumPropsExA                       uintptr
	fnEnumThreadWindows                  uintptr
	fnEnumWindowStationsW                uintptr
	fnEnumWindowStationsA                uintptr
	fnEnumWindows                        uintptr
	fnEqualRect                          uintptr
	fnExcludeUpdateRgn                   uintptr
	fnExitWindowsEx                      uintptr
	fnFillRect                           uintptr
	fnFlashWindow                        uintptr
	fnFlashWindowEx                      uintptr
	fnFrameRect                          uintptr
	fnFreeDDElParam                      uintptr
	fnGetActiveWindow                    uintptr
	fnGetAltTabInfoW                     uintptr
	fnGetAltTabInfoA                     uintptr
	fnGetAncestor                        uintptr
	fnGetAsyncKeyState                   uintptr
	fnGetCapture                         uintptr
	fnGetCaretBlinkTime                  uintptr
	fnGetCaretPos                        uintptr
	fnGetClassInfoW                      uintptr
	fnGetClassInfoA                      uintptr
	fnGetClassInfoExW                    uintptr
	fnGetWindowThreadProcessId			uintptr
	fnGetClassInfoExA                    uintptr
	fnGetClassNameW                      uintptr
	fnGetClassNameA                      uintptr
	fnGetClassWord                       uintptr
	fnGetClientRect                      uintptr
	fnGetClipCursor                      uintptr
	fnGetClipboardData                   uintptr
	fnGetClipboardFormatNameW            uintptr
	fnGetClipboardFormatNameA            uintptr
	fnGetClipboardSequenceNumber         uintptr
	fnGetClipboardOwner                  uintptr
	fnGetClipboardViewer                 uintptr
	fnGetComboBoxInfo                    uintptr
	fnGetCursor                          uintptr
	fnGetCursorInfo                      uintptr
	fnGetCursorPos                       uintptr
	fnGetPhysicalCursorPos               uintptr
	fnGetDC                              uintptr
	fnGetDCEx                            uintptr
	fnGetDesktopWindow                   uintptr
	fnGetDialogBaseUnits                 uintptr
	fnGetDlgCtrlID                       uintptr
	fnGetDlgItem                         uintptr
	fnGetDlgItemInt                      uintptr
	fnGetDlgItemTextW                    uintptr
	fnGetDlgItemTextA                    uintptr
	fnGetDoubleClickTime                 uintptr
	fnGetFocus                           uintptr
	fnGetForegroundWindow                uintptr
	fnGetGuiResources                    uintptr
	fnGetGUIThreadInfo                   uintptr
	fnBlockInput                         uintptr
	fnIsProcessDPIAware                  uintptr
	fnIsTouchWindow                      uintptr
	fnGetIconInfo                        uintptr
	fnGetIconInfoExW                     uintptr
	fnGetIconInfoExA                     uintptr
	fnGetInputState                      uintptr
	fnGetKBCodePage                      uintptr
	fnGetKeyNameTextW                    uintptr
	fnGetKeyNameTextA                    uintptr
	fnGetKeyState                        uintptr
	fnGetKeyboardLayout                  uintptr
	fnGetKeyboardLayoutList              uintptr
	fnGetKeyboardLayoutNameW             uintptr
	fnGetKeyboardLayoutNameA             uintptr
	fnGetMouseMovePointsEx               uintptr
	fnGetKeyboardState                   uintptr
	fnGetKeyboardType                    uintptr
	fnGetLastActivePopup                 uintptr
	fnGetLastInputInfo                   uintptr
	fnGetListBoxInfo                     uintptr
	fnGetMenu                            uintptr
	fnGetMenuBarInfo                     uintptr
	fnGetMenuCheckMarkDimensions         uintptr
	fnGetMenuContextHelpId               uintptr
	fnGetMenuDefaultItem                 uintptr
	fnGetMenuInfo                        uintptr
	fnGetMenuItemCount                   uintptr
	fnGetMenuItemID                      uintptr
	fnGetMenuItemInfoW                   uintptr
	fnGetMenuItemInfoA                   uintptr
	fnGetMenuItemRect                    uintptr
	fnGetMenuState                       uintptr
	fnGetMenuStringW                     uintptr
	fnGetMenuStringA                     uintptr
	fnGetNextDlgGroupItem                uintptr
	fnGetNextDlgTabItem                  uintptr
	fnGetWindow                          uintptr
	fnGetOpenClipboardWindow             uintptr
	fnAddClipboardFormatListener         uintptr
	fnRemoveClipboardFormatListener      uintptr
	fnGetUpdatedClipboardFormats         uintptr
	fnGetParent                          uintptr
	fnGetPriorityClipboardFormat         uintptr
	fnGetProcessWindowStation            uintptr
	fnGetPropW                           uintptr
	fnGetPropA                           uintptr
	fnGetQueueStatus                     uintptr
	fnGetScrollBarInfo                   uintptr
	fnGetScrollInfo                      uintptr
	fnGetScrollPos                       uintptr
	fnGetScrollRange                     uintptr
	fnGetSubMenu                         uintptr
	fnGetSysColor                        uintptr
	fnGetSysColorBrush                   uintptr
	fnGetSystemMenu                      uintptr
	fnGetSystemMetrics                   uintptr
	fnGetTabbedTextExtentW               uintptr
	fnGetTabbedTextExtentA               uintptr
	fnGetThreadDesktop                   uintptr
	fnGetTitleBarInfo                    uintptr
	fnGetTopWindow                       uintptr
	fnGetTouchInputInfo                  uintptr
	fnGetUpdateRect                      uintptr
	fnGetUpdateRgn                       uintptr
	fnGetUserObjectInformationW          uintptr
	fnGetUserObjectInformationA          uintptr
	fnGetUserObjectSecurity              uintptr
	fnGetWindowContextHelpId             uintptr
	fnGetWindowDC                        uintptr
	fnGetWindowDisplayAffinity           uintptr
	fnGetWindowInfo                      uintptr
	fnGetWindowModuleFileNameW           uintptr
	fnGetWindowModuleFileNameA           uintptr
	fnGetWindowPlacement                 uintptr
	fnGetWindowRect                      uintptr
	fnGetWindowRgn                       uintptr
	fnGetWindowTextW                     uintptr
	fnGetWindowTextA                     uintptr
	fnGetWindowTextLengthW               uintptr
	fnGetWindowTextLengthA               uintptr
	fnGetWindowWord                      uintptr
	fnGrayStringW                        uintptr
	fnGrayStringA                        uintptr
	fnHideCaret                          uintptr
	fnHiliteMenuItem                     uintptr
	fnImpersonateDdeClientWindow         uintptr
	fnInflateRect                        uintptr
	fnInsertMenuW                        uintptr
	fnInsertMenuA                        uintptr
	fnInsertMenuItemW                    uintptr
	fnInsertMenuItemA                    uintptr
	fnIntersectRect                      uintptr
	fnInvalidateRect                     uintptr
	fnInvalidateRgn                      uintptr
	fnInvertRect                         uintptr
	fnIsCharAlphaW                       uintptr
	fnIsCharAlphaA                       uintptr
	fnIsCharAlphaNumericW                uintptr
	fnIsCharAlphaNumericA                uintptr
	fnIsCharLowerW                       uintptr
	fnIsCharLowerA                       uintptr
	fnIsCharUpperW                       uintptr
	fnIsCharUpperA                       uintptr
	fnIsChild                            uintptr
	fnIsClipboardFormatAvailable         uintptr
	fnIsDialogMessageW                   uintptr
	fnIsDialogMessageA                   uintptr
	fnIsDlgButtonChecked                 uintptr
	fnIsIconic                           uintptr
	fnIsMenu                             uintptr
	fnIsRectEmpty                        uintptr
	fnIsWindowEnabled                    uintptr
	fnIsWindowUnicode                    uintptr
	fnIsWindowVisible                    uintptr
	fnIsZoomed                           uintptr
	fnKillTimer                          uintptr
	fnLoadAcceleratorsW                  uintptr
	fnLoadAcceleratorsA                  uintptr
	fnLoadBitmapW                        uintptr
	fnLoadBitmapA                        uintptr
	fnLoadCursorFromFileW                uintptr
	fnLoadCursorFromFileA                uintptr
	fnLoadIconW                          uintptr
	fnLoadIconA                          uintptr
	fnLoadImageW                         uintptr
	fnLoadImageA                         uintptr
	fnLoadKeyboardLayoutW                uintptr
	fnLoadKeyboardLayoutA                uintptr
	fnLoadMenuW                          uintptr
	fnLoadMenuA                          uintptr
	fnLoadMenuIndirectW                  uintptr
	fnLoadMenuIndirectA                  uintptr
	fnLockWindowUpdate                   uintptr
	fnLockWorkStation                    uintptr
	fnLookupIconIdFromDirectory          uintptr
	fnLookupIconIdFromDirectoryEx        uintptr
	fnMapDialogRect                      uintptr
	fnMapVirtualKeyW                     uintptr
	fnMapVirtualKeyA                     uintptr
	fnMapVirtualKeyExW                   uintptr
	fnMapVirtualKeyExA                   uintptr
	fnMapWindowPoints                    uintptr
	fnMenuItemFromPoint                  uintptr
	fnMessageBeep                        uintptr
	fnMessageBoxW                        uintptr
	fnMessageBoxA                        uintptr
	fnMessageBoxExW                      uintptr
	fnMessageBoxExA                      uintptr
	fnMessageBoxIndirectW                uintptr
	fnMessageBoxIndirectA                uintptr
	fnModifyMenuW                        uintptr
	fnModifyMenuA                        uintptr
	fnMoveWindow                         uintptr
	fnOemKeyScan                         uintptr
	fnOemToCharA                         uintptr
	fnOemToCharBuffA                     uintptr
	fnOemToCharW                         uintptr
	fnOemToCharBuffW                     uintptr
	fnOffsetRect                         uintptr
	fnOpenClipboard                      uintptr
	fnOpenDesktopW                       uintptr
	fnOpenDesktopA                       uintptr
	fnOpenIcon                           uintptr
	fnOpenInputDesktop                   uintptr
	fnOpenWindowStationW                 uintptr
	fnOpenWindowStationA                 uintptr
	fnPackDDElParam                      uintptr
	fnPaintDesktop                       uintptr
	fnPtInRect                           uintptr
	fnRealChildWindowFromPoint           uintptr
	fnRealGetWindowClassW                uintptr
	fnRealGetWindowClassA                uintptr
	fnRedrawWindow                       uintptr
	fnRegisterClassW                     uintptr
	fnRegisterClassA                     uintptr
	fnRegisterClassExW                   uintptr
	fnRegisterClassExA                   uintptr
	fnCreateWindowExA                    uintptr
	fnCreateWindowExW                    uintptr
	fnRegisterClipboardFormatW           uintptr
	fnRegisterClipboardFormatA           uintptr
	fnRegisterDeviceNotificationW        uintptr
	fnRegisterDeviceNotificationA        uintptr
	fnRegisterHotKey                     uintptr
	fnRegisterPowerSettingNotification   uintptr
	fnRegisterTouchWindow                uintptr
	fnRegisterWindowMessageW             uintptr
	fnRegisterWindowMessageA             uintptr
	fnReleaseCapture                     uintptr
	fnReleaseDC                          uintptr
	fnRemoveMenu                         uintptr
	fnRemovePropW                        uintptr
	fnRemovePropA                        uintptr
	fnReplyMessage                       uintptr
	fnReuseDDElParam                     uintptr
	fnScreenToClient                     uintptr
	fnLogicalToPhysicalPoint             uintptr
	fnPhysicalToLogicalPoint             uintptr
	fnScrollDC                           uintptr
	fnScrollWindow                       uintptr
	fnSendDlgItemMessageW                uintptr
	fnSendDlgItemMessageA                uintptr
	fnSendInput                          uintptr
	fnSetActiveWindow                    uintptr
	fnSetCapture                         uintptr
	fnSetCaretBlinkTime                  uintptr
	fnSetCaretPos                        uintptr
	fnSetClassWord                       uintptr
	fnSetClipboardData                   uintptr
	fnSetClipboardViewer                 uintptr
	fnSetCursor                          uintptr
	fnSetCursorPos                       uintptr
	fnSetDlgItemInt                      uintptr
	fnSetDlgItemTextW                    uintptr
	fnSetDlgItemTextA                    uintptr
	fnSetDoubleClickTime                 uintptr
	fnSetFocus                           uintptr
	fnSetForegroundWindow                uintptr
	fnAllowSetForegroundWindow           uintptr
	fnLockSetForegroundWindow            uintptr
	fnSetKeyboardState                   uintptr
	fnSetMenu                            uintptr
	fnSetMenuContextHelpId               uintptr
	fnSetMenuDefaultItem                 uintptr
	fnSetMenuInfo                        uintptr
	fnSetMenuItemBitmaps                 uintptr
	fnSetMenuItemInfoW                   uintptr
	fnSetMenuItemInfoA                   uintptr
	fnSetMessageExtraInfo                uintptr
	fnSetMessageQueue                    uintptr
	fnSetParent                          uintptr
	fnSetPhysicalCursorPos               uintptr
	fnSetProcessDPIAware                 uintptr
	fnSetProcessWindowStation            uintptr
	fnSetPropW                           uintptr
	fnSetPropA                           uintptr
	fnSetRect                            uintptr
	fnSetRectEmpty                       uintptr
	fnSetScrollInfo                      uintptr
	fnSetScrollPos                       uintptr
	fnSetScrollRange                     uintptr
	fnSetSysColors                       uintptr
	fnSetSystemCursor                    uintptr
	fnSetThreadDesktop                   uintptr
	fnSetTimer                           uintptr
	fnSetUserObjectInformationW          uintptr
	fnSetUserObjectInformationA          uintptr
	fnSetUserObjectSecurity              uintptr
	fnSetWindowContextHelpId             uintptr
	fnSetWindowDisplayAffinity           uintptr
	fnSetWindowPlacement                 uintptr
	fnSetWindowPos                       uintptr
	fnSetWindowTextW                     uintptr
	fnSetWindowTextA                     uintptr
	fnSetWindowWord                      uintptr
	fnSetWindowsHookW                    uintptr
	fnSetWindowsHookA                    uintptr
	fnSetWindowsHookExW                  uintptr
	fnSetWindowsHookExA                  uintptr
	fnSetWindowRgn                       uintptr
	fnSetWinEventHook                    uintptr
	fnShowCaret                          uintptr
	fnShowCursor                         uintptr
	fnShowOwnedPopups                    uintptr
	fnShowScrollBar                      uintptr
	fnShowWindow                         uintptr
	fnAnimateWindow                      uintptr
	fnShowWindowAsync                    uintptr
	fnSubtractRect                       uintptr
	fnSwapMouseButton                    uintptr
	fnSwitchDesktop                      uintptr
	fnSoundSentry                        uintptr
	fnTabbedTextOutW                     uintptr
	fnTabbedTextOutA                     uintptr
	fnTileWindows                        uintptr
	fnToAscii                            uintptr
	fnToAsciiEx                          uintptr
	fnTrackMouseEvent                    uintptr
	fnTrackPopupMenu                     uintptr
	fnTrackPopupMenuEx                   uintptr
	fnTranslateAcceleratorW              uintptr
	fnTranslateAcceleratorA              uintptr
	fnTranslateMDISysAccel               uintptr
	fnUnhookWindowsHook                  uintptr
	fnUnhookWindowsHookEx                uintptr
	fnUnhookWinEvent                     uintptr
	fnUnionRect                          uintptr
	fnUnloadKeyboardLayout               uintptr
	fnUnpackDDElParam                    uintptr
	fnUnregisterClassW                   uintptr
	fnUnregisterClassA                   uintptr
	fnUnregisterDeviceNotification       uintptr
	fnUnregisterPowerSettingNotification uintptr
	fnUnregisterHotKey                   uintptr
	fnUnregisterTouchWindow              uintptr
	fnUpdateWindow                       uintptr
	fnUpdateLayeredWindow                uintptr
	fnUserHandleGrantAccess              uintptr
	fnValidateRect                       uintptr
	fnValidateRgn                        uintptr
	fnVkKeyScanW                         uintptr
	fnVkKeyScanA                         uintptr
	fnVkKeyScanExW                       uintptr
	fnVkKeyScanExA                       uintptr
	fnWaitForInputIdle                   uintptr
	fnWaitMessage                        uintptr
	fnWindowFromDC                       uintptr
	fnWindowFromPoint                    uintptr
	fnPostQuitMessage                    uintptr
	fnWindowFromPhysicalPoint            uintptr
	fnwsprintfW                          uintptr
	fnwsprintfA                          uintptr
	fnwvsprintfW                         uintptr
	fnwvsprintfA                         uintptr
	fnGetWindowLongPtrW                  uintptr
	fnSetWindowLongPtrW                  uintptr
	fnGetClassLongPtrW                   uintptr
	fnSetClassLongPtrW                   uintptr
	fnGetWindowLongPtrA                  uintptr
	fnSetWindowLongPtrA                  uintptr
	fnGetClassLongPtrA                   uintptr
	fnSetClassLongPtrA                   uintptr
	fnGetWindowLongW                     uintptr
	fnSetWindowLongW                     uintptr
	fnGetClassLongW                      uintptr
	fnSetClassLongW                      uintptr
	fnGetWindowLongA                     uintptr
	fnSetWindowLongA                     uintptr
	fnGetClassLongA                      uintptr
	fnSetClassLongA                      uintptr
	fnGetGestureInfo                     uintptr
	fnGetGestureExtraArgs                uintptr
	fnCloseGestureInfoHandle             uintptr
	fnSetGestureConfig                   uintptr
	fnGetGestureConfig                   uintptr
	fnWINNLSEnableIME                    uintptr
	fnTranslateMessage                   uintptr
	fnDispatchMessageA                   uintptr
	fnDispatchMessageW                   uintptr
	fnPeekMessageW                       uintptr
	fnSendMessageW			   			 uintptr
	fnPostMessageA						 uintptr
	fnPostMessageW						 uintptr
)

func init() {
	libuser32, _ = syscall.LoadLibrary("user32.dll")
	fnSystemParametersInfo, _ = syscall.GetProcAddress(libuser32, "SystemParametersInfoW")
	fnLoadCursor, _ = syscall.GetProcAddress(libuser32, "LoadCursorW")
	fnPostQuitMessage, _ = syscall.GetProcAddress(libuser32, "PostQuitMessage")
	fnTranslateMessage, _ = syscall.GetProcAddress(libuser32, "TranslateMessage")
	fnDispatchMessageA, _ = syscall.GetProcAddress(libuser32, "DispatchMessageA")
	fnDispatchMessageW, _ = syscall.GetProcAddress(libuser32, "DispatchMessageW")
	fnPeekMessageW, _ = syscall.GetProcAddress(libuser32, "PeekMessageW")
	fnSendMessageW, _ = syscall.GetProcAddress(libuser32, "SendMessageW")
	fnActivateKeyboardLayout, _ = syscall.GetProcAddress(libuser32, "ActivateKeyboardLayout")
	fnAdjustWindowRect, _ = syscall.GetProcAddress(libuser32, "AdjustWindowRect")
	fnAdjustWindowRectEx, _ = syscall.GetProcAddress(libuser32, "AdjustWindowRectEx")
	fnAnyPopup, _ = syscall.GetProcAddress(libuser32, "AnyPopup")
	fnAppendMenuW, _ = syscall.GetProcAddress(libuser32, "AppendMenuW")
	fnAppendMenuA, _ = syscall.GetProcAddress(libuser32, "AppendMenuA")
	fnArrangeIconicWindows, _ = syscall.GetProcAddress(libuser32, "ArrangeIconicWindows")
	fnAttachThreadInput, _ = syscall.GetProcAddress(libuser32, "AttachThreadInput")
	fnBeginDeferWindowPos, _ = syscall.GetProcAddress(libuser32, "BeginDeferWindowPos")
	fnBeginPaint, _ = syscall.GetProcAddress(libuser32, "BeginPaint")
	fnBringWindowToTop, _ = syscall.GetProcAddress(libuser32, "BringWindowToTop")
	fnBroadcastSystemMessageW, _ = syscall.GetProcAddress(libuser32, "BroadcastSystemMessageW")
	fnBroadcastSystemMessageA, _ = syscall.GetProcAddress(libuser32, "BroadcastSystemMessageA")
	fnCallMsgFilterW, _ = syscall.GetProcAddress(libuser32, "CallMsgFilterW")
	fnCallMsgFilterA, _ = syscall.GetProcAddress(libuser32, "CallMsgFilterA")
	fnCallNextHookEx, _ = syscall.GetProcAddress(libuser32, "CallNextHookEx")
	fnCallWindowProcW, _ = syscall.GetProcAddress(libuser32, "CallWindowProcW")
	fnCallWindowProcA, _ = syscall.GetProcAddress(libuser32, "CallWindowProcA")
	fnCascadeWindows, _ = syscall.GetProcAddress(libuser32, "CascadeWindows")
	fnChangeWindowMessageFilter, _ = syscall.GetProcAddress(libuser32, "ChangeWindowMessageFilter")
	fnChangeClipboardChain, _ = syscall.GetProcAddress(libuser32, "ChangeClipboardChain")
	fnChangeDisplaySettingsW, _ = syscall.GetProcAddress(libuser32, "ChangeDisplaySettingsW")
	fnChangeDisplaySettingsA, _ = syscall.GetProcAddress(libuser32, "ChangeDisplaySettingsA")
	fnChangeDisplaySettingsExW, _ = syscall.GetProcAddress(libuser32, "ChangeDisplaySettingsExW")
	fnChangeDisplaySettingsExA, _ = syscall.GetProcAddress(libuser32, "ChangeDisplaySettingsExA")
	fnChangeMenuW, _ = syscall.GetProcAddress(libuser32, "ChangeMenuW")
	fnChangeMenuA, _ = syscall.GetProcAddress(libuser32, "ChangeMenuA")
	fnCheckDlgButton, _ = syscall.GetProcAddress(libuser32, "CheckDlgButton")
	fnCheckMenuItem, _ = syscall.GetProcAddress(libuser32, "CheckMenuItem")
	fnCheckMenuRadioItem, _ = syscall.GetProcAddress(libuser32, "CheckMenuRadioItem")
	fnCheckRadioButton, _ = syscall.GetProcAddress(libuser32, "CheckRadioButton")
	fnChildWindowFromPoint, _ = syscall.GetProcAddress(libuser32, "ChildWindowFromPoint")
	fnChildWindowFromPointEx, _ = syscall.GetProcAddress(libuser32, "ChildWindowFromPointEx")
	fnClientToScreen, _ = syscall.GetProcAddress(libuser32, "ClientToScreen")
	fnClipCursor, _ = syscall.GetProcAddress(libuser32, "ClipCursor")
	fnCloseClipboard, _ = syscall.GetProcAddress(libuser32, "CloseClipboard")
	fnCloseDesktop, _ = syscall.GetProcAddress(libuser32, "CloseDesktop")
	fnCloseTouchInputHandle, _ = syscall.GetProcAddress(libuser32, "CloseTouchInputHandle")
	fnCloseWindow, _ = syscall.GetProcAddress(libuser32, "CloseWindow")
	fnCloseWindowStation, _ = syscall.GetProcAddress(libuser32, "CloseWindowStation")
	fnCopyAcceleratorTableW, _ = syscall.GetProcAddress(libuser32, "CopyAcceleratorTableW")
	fnCopyAcceleratorTableA, _ = syscall.GetProcAddress(libuser32, "CopyAcceleratorTableA")
	fnCopyIcon, _ = syscall.GetProcAddress(libuser32, "CopyIcon")
	fnCopyImage, _ = syscall.GetProcAddress(libuser32, "CopyImage")
	fnCopyRect, _ = syscall.GetProcAddress(libuser32, "CopyRect")
	fnCountClipboardFormats, _ = syscall.GetProcAddress(libuser32, "CountClipboardFormats")
	fnCreateAcceleratorTableW, _ = syscall.GetProcAddress(libuser32, "CreateAcceleratorTableW")
	fnCreateAcceleratorTableA, _ = syscall.GetProcAddress(libuser32, "CreateAcceleratorTableA")
	fnCreateCaret, _ = syscall.GetProcAddress(libuser32, "CreateCaret")
	fnCreateCursor, _ = syscall.GetProcAddress(libuser32, "CreateCursor")
	fnCreateDesktopW, _ = syscall.GetProcAddress(libuser32, "CreateDesktopW")
	fnCreateDesktopA, _ = syscall.GetProcAddress(libuser32, "CreateDesktopA")
	fnCreateDesktopExW, _ = syscall.GetProcAddress(libuser32, "CreateDesktopExW")
	fnCreateDesktopExA, _ = syscall.GetProcAddress(libuser32, "CreateDesktopExA")
	fnCreateDialogIndirectParamW, _ = syscall.GetProcAddress(libuser32, "CreateDialogIndirectParamW")
	fnCreateDialogIndirectParamA, _ = syscall.GetProcAddress(libuser32, "CreateDialogIndirectParamA")
	fnCreateDialogParamW, _ = syscall.GetProcAddress(libuser32, "CreateDialogParamW")
	fnCreateDialogParamA, _ = syscall.GetProcAddress(libuser32, "CreateDialogParamA")
	fnCreateIcon, _ = syscall.GetProcAddress(libuser32, "CreateIcon")
	fnCreateIconFromResource, _ = syscall.GetProcAddress(libuser32, "CreateIconFromResource")
	fnCreateIconFromResourceEx, _ = syscall.GetProcAddress(libuser32, "CreateIconFromResourceEx")
	fnCreateIconIndirect, _ = syscall.GetProcAddress(libuser32, "CreateIconIndirect")
	fnCreateMDIWindowW, _ = syscall.GetProcAddress(libuser32, "CreateMDIWindowW")
	fnCreateMDIWindowA, _ = syscall.GetProcAddress(libuser32, "CreateMDIWindowA")
	fnCreateMenu, _ = syscall.GetProcAddress(libuser32, "CreateMenu")
	fnCreatePopupMenu, _ = syscall.GetProcAddress(libuser32, "CreatePopupMenu")
	fnCreateWindowStationW, _ = syscall.GetProcAddress(libuser32, "CreateWindowStationW")
	fnCreateWindowStationA, _ = syscall.GetProcAddress(libuser32, "CreateWindowStationA")
	fnDdeSetQualityOfService, _ = syscall.GetProcAddress(libuser32, "DdeSetQualityOfService")
	fnDefDlgProcW, _ = syscall.GetProcAddress(libuser32, "DefDlgProcW")
	fnDefDlgProcA, _ = syscall.GetProcAddress(libuser32, "DefDlgProcA")
	fnDefFrameProcW, _ = syscall.GetProcAddress(libuser32, "DefFrameProcW")
	fnDefFrameProcA, _ = syscall.GetProcAddress(libuser32, "DefFrameProcA")
	fnDefMDIChildProcW, _ = syscall.GetProcAddress(libuser32, "DefMDIChildProcW")
	fnDefMDIChildProcA, _ = syscall.GetProcAddress(libuser32, "DefMDIChildProcA")
	fnDefWindowProcW, _ = syscall.GetProcAddress(libuser32, "DefWindowProcW")
	fnDefWindowProcA, _ = syscall.GetProcAddress(libuser32, "DefWindowProcA")
	fnDeferWindowPos, _ = syscall.GetProcAddress(libuser32, "DeferWindowPos")
	fnDeleteMenu, _ = syscall.GetProcAddress(libuser32, "DeleteMenu")
	fnDestroyAcceleratorTable, _ = syscall.GetProcAddress(libuser32, "DestroyAcceleratorTable")
	fnDestroyCaret, _ = syscall.GetProcAddress(libuser32, "DestroyCaret")
	fnDestroyCursor, _ = syscall.GetProcAddress(libuser32, "DestroyCursor")
	fnDestroyIcon, _ = syscall.GetProcAddress(libuser32, "DestroyIcon")
	fnDestroyMenu, _ = syscall.GetProcAddress(libuser32, "DestroyMenu")
	fnDestroyWindow, _ = syscall.GetProcAddress(libuser32, "DestroyWindow")
	fnDialogBoxIndirectParamW, _ = syscall.GetProcAddress(libuser32, "DialogBoxIndirectParamW")
	fnDialogBoxIndirectParamA, _ = syscall.GetProcAddress(libuser32, "DialogBoxIndirectParamA")
	fnDialogBoxParamW, _ = syscall.GetProcAddress(libuser32, "DialogBoxParamW")
	fnDialogBoxParamA, _ = syscall.GetProcAddress(libuser32, "DialogBoxParamA")
	fnDlgDirListW, _ = syscall.GetProcAddress(libuser32, "DlgDirListW")
	fnDlgDirListA, _ = syscall.GetProcAddress(libuser32, "DlgDirListA")
	fnDlgDirListComboBoxW, _ = syscall.GetProcAddress(libuser32, "DlgDirListComboBoxW")
	fnDlgDirListComboBoxA, _ = syscall.GetProcAddress(libuser32, "DlgDirListComboBoxA")
	fnDlgDirSelectComboBoxExW, _ = syscall.GetProcAddress(libuser32, "DlgDirSelectComboBoxExW")
	fnDlgDirSelectComboBoxExA, _ = syscall.GetProcAddress(libuser32, "DlgDirSelectComboBoxExA")
	fnDlgDirSelectExW, _ = syscall.GetProcAddress(libuser32, "DlgDirSelectExW")
	fnDlgDirSelectExA, _ = syscall.GetProcAddress(libuser32, "DlgDirSelectExA")
	fnDragDetect, _ = syscall.GetProcAddress(libuser32, "DragDetect")
	fnDragObject, _ = syscall.GetProcAddress(libuser32, "DragObject")
	fnDrawAnimatedRects, _ = syscall.GetProcAddress(libuser32, "DrawAnimatedRects")
	fnDrawCaption, _ = syscall.GetProcAddress(libuser32, "DrawCaption")
	fnDrawEdge, _ = syscall.GetProcAddress(libuser32, "DrawEdge")
	fnDrawFocusRect, _ = syscall.GetProcAddress(libuser32, "DrawFocusRect")
	fnDrawFrameControl, _ = syscall.GetProcAddress(libuser32, "DrawFrameControl")
	fnDrawIcon, _ = syscall.GetProcAddress(libuser32, "DrawIcon")
	fnDrawIconEx, _ = syscall.GetProcAddress(libuser32, "DrawIconEx")
	fnDrawMenuBar, _ = syscall.GetProcAddress(libuser32, "DrawMenuBar")
	fnDrawStateW, _ = syscall.GetProcAddress(libuser32, "DrawStateW")
	fnDrawStateA, _ = syscall.GetProcAddress(libuser32, "DrawStateA")
	fnDrawTextExW, _ = syscall.GetProcAddress(libuser32, "DrawTextExW")
	fnDrawTextExA, _ = syscall.GetProcAddress(libuser32, "DrawTextExA")
	fnDrawTextW,_=syscall.GetProcAddress(libuser32, "DrawTextW")
	fnEmptyClipboard, _ = syscall.GetProcAddress(libuser32, "EmptyClipboard")
	fnEnableMenuItem, _ = syscall.GetProcAddress(libuser32, "EnableMenuItem")
	fnEnableScrollBar, _ = syscall.GetProcAddress(libuser32, "EnableScrollBar")
	fnEnableWindow, _ = syscall.GetProcAddress(libuser32, "EnableWindow")
	fnEndDeferWindowPos, _ = syscall.GetProcAddress(libuser32, "EndDeferWindowPos")
	fnEndDialog, _ = syscall.GetProcAddress(libuser32, "EndDialog")
	fnEndPaint, _ = syscall.GetProcAddress(libuser32, "EndPaint")
	fnEnumChildWindows, _ = syscall.GetProcAddress(libuser32, "EnumChildWindows")
	fnEnumClipboardFormats, _ = syscall.GetProcAddress(libuser32, "EnumClipboardFormats")
	fnEnumDesktopsW, _ = syscall.GetProcAddress(libuser32, "EnumDesktopsW")
	fnEnumDesktopsA, _ = syscall.GetProcAddress(libuser32, "EnumDesktopsA")
	fnEnumDesktopWindows, _ = syscall.GetProcAddress(libuser32, "EnumDesktopWindows")
	fnEnumDisplaySettingsW, _ = syscall.GetProcAddress(libuser32, "EnumDisplaySettingsW")
	fnEnumDisplaySettingsA, _ = syscall.GetProcAddress(libuser32, "EnumDisplaySettingsA")
	fnEnumDisplayDevicesW, _ = syscall.GetProcAddress(libuser32, "EnumDisplayDevicesW")
	fnEnumDisplayDevicesA, _ = syscall.GetProcAddress(libuser32, "EnumDisplayDevicesA")
	fnEnumPropsW, _ = syscall.GetProcAddress(libuser32, "EnumPropsW")
	fnEnumPropsA, _ = syscall.GetProcAddress(libuser32, "EnumPropsA")
	fnEnumPropsExW, _ = syscall.GetProcAddress(libuser32, "EnumPropsExW")
	fnEnumPropsExA, _ = syscall.GetProcAddress(libuser32, "EnumPropsExA")
	fnEnumThreadWindows, _ = syscall.GetProcAddress(libuser32, "EnumThreadWindows")
	fnEnumWindowStationsW, _ = syscall.GetProcAddress(libuser32, "EnumWindowStationsW")
	fnEnumWindowStationsA, _ = syscall.GetProcAddress(libuser32, "EnumWindowStationsA")
	fnEnumWindows, _ = syscall.GetProcAddress(libuser32, "EnumWindows")
	fnEqualRect, _ = syscall.GetProcAddress(libuser32, "EqualRect")
	fnExcludeUpdateRgn, _ = syscall.GetProcAddress(libuser32, "ExcludeUpdateRgn")
	fnExitWindowsEx, _ = syscall.GetProcAddress(libuser32, "ExitWindowsEx")
	fnFillRect, _ = syscall.GetProcAddress(libuser32, "FillRect")
	fnFlashWindow, _ = syscall.GetProcAddress(libuser32, "FlashWindow")
	fnFlashWindowEx, _ = syscall.GetProcAddress(libuser32, "FlashWindowEx")
	fnFrameRect, _ = syscall.GetProcAddress(libuser32, "FrameRect")
	fnFreeDDElParam, _ = syscall.GetProcAddress(libuser32, "FreeDDElParam")
	fnGetActiveWindow, _ = syscall.GetProcAddress(libuser32, "GetActiveWindow")
	fnGetAltTabInfoW, _ = syscall.GetProcAddress(libuser32, "GetAltTabInfoW")
	fnGetAltTabInfoA, _ = syscall.GetProcAddress(libuser32, "GetAltTabInfoA")
	fnGetAncestor, _ = syscall.GetProcAddress(libuser32, "GetAncestor")
	fnGetAsyncKeyState, _ = syscall.GetProcAddress(libuser32, "GetAsyncKeyState")
	fnGetCapture, _ = syscall.GetProcAddress(libuser32, "GetCapture")
	fnGetCaretBlinkTime, _ = syscall.GetProcAddress(libuser32, "GetCaretBlinkTime")
	fnGetCaretPos, _ = syscall.GetProcAddress(libuser32, "GetCaretPos")
	fnGetClassInfoW, _ = syscall.GetProcAddress(libuser32, "GetClassInfoW")
	fnGetClassInfoA, _ = syscall.GetProcAddress(libuser32, "GetClassInfoA")
	fnGetClassInfoExW, _ = syscall.GetProcAddress(libuser32, "GetClassInfoExW")
	fnGetClassInfoExA, _ = syscall.GetProcAddress(libuser32, "GetClassInfoExA")
	fnGetClassNameW, _ = syscall.GetProcAddress(libuser32, "GetClassNameW")
	fnGetClassNameA, _ = syscall.GetProcAddress(libuser32, "GetClassNameA")
	fnGetClassWord, _ = syscall.GetProcAddress(libuser32, "GetClassWord")
	fnGetClientRect, _ = syscall.GetProcAddress(libuser32, "GetClientRect")
	fnGetClipCursor, _ = syscall.GetProcAddress(libuser32, "GetClipCursor")
	fnGetClipboardData, _ = syscall.GetProcAddress(libuser32, "GetClipboardData")
	fnGetClipboardFormatNameW, _ = syscall.GetProcAddress(libuser32, "GetClipboardFormatNameW")
	fnGetClipboardFormatNameA, _ = syscall.GetProcAddress(libuser32, "GetClipboardFormatNameA")
	fnGetClipboardSequenceNumber, _ = syscall.GetProcAddress(libuser32, "GetClipboardSequenceNumber")
	fnGetClipboardOwner, _ = syscall.GetProcAddress(libuser32, "GetClipboardOwner")
	fnGetClipboardViewer, _ = syscall.GetProcAddress(libuser32, "GetClipboardViewer")
	fnGetComboBoxInfo, _ = syscall.GetProcAddress(libuser32, "GetComboBoxInfo")
	fnGetCursor, _ = syscall.GetProcAddress(libuser32, "GetCursor")
	fnGetCursorInfo, _ = syscall.GetProcAddress(libuser32, "GetCursorInfo")
	fnGetCursorPos, _ = syscall.GetProcAddress(libuser32, "GetCursorPos")
	fnGetPhysicalCursorPos, _ = syscall.GetProcAddress(libuser32, "GetPhysicalCursorPos")
	fnGetDC, _ = syscall.GetProcAddress(libuser32, "GetDC")
	fnGetDCEx, _ = syscall.GetProcAddress(libuser32, "GetDCEx")
	fnGetDesktopWindow, _ = syscall.GetProcAddress(libuser32, "GetDesktopWindow")
	fnGetDialogBaseUnits, _ = syscall.GetProcAddress(libuser32, "GetDialogBaseUnits")
	fnGetDlgCtrlID, _ = syscall.GetProcAddress(libuser32, "GetDlgCtrlID")
	fnGetDlgItem, _ = syscall.GetProcAddress(libuser32, "GetDlgItem")
	fnGetDlgItemInt, _ = syscall.GetProcAddress(libuser32, "GetDlgItemInt")
	fnGetDlgItemTextW, _ = syscall.GetProcAddress(libuser32, "GetDlgItemTextW")
	fnGetDlgItemTextA, _ = syscall.GetProcAddress(libuser32, "GetDlgItemTextA")
	fnGetDoubleClickTime, _ = syscall.GetProcAddress(libuser32, "GetDoubleClickTime")
	fnGetFocus, _ = syscall.GetProcAddress(libuser32, "GetFocus")
	fnGetForegroundWindow, _ = syscall.GetProcAddress(libuser32, "GetForegroundWindow")
	fnGetGuiResources, _ = syscall.GetProcAddress(libuser32, "GetGuiResources")
	fnGetGUIThreadInfo, _ = syscall.GetProcAddress(libuser32, "GetGUIThreadInfo")
	fnBlockInput, _ = syscall.GetProcAddress(libuser32, "BlockInput")
	fnIsProcessDPIAware, _ = syscall.GetProcAddress(libuser32, "IsProcessDPIAware")
	fnIsTouchWindow, _ = syscall.GetProcAddress(libuser32, "IsTouchWindow")
	fnGetIconInfo, _ = syscall.GetProcAddress(libuser32, "GetIconInfo")
	fnGetIconInfoExW, _ = syscall.GetProcAddress(libuser32, "GetIconInfoExW")
	fnGetIconInfoExA, _ = syscall.GetProcAddress(libuser32, "GetIconInfoExA")
	fnGetInputState, _ = syscall.GetProcAddress(libuser32, "GetInputState")
	fnGetKBCodePage, _ = syscall.GetProcAddress(libuser32, "GetKBCodePage")
	fnGetKeyNameTextW, _ = syscall.GetProcAddress(libuser32, "GetKeyNameTextW")
	fnGetKeyNameTextA, _ = syscall.GetProcAddress(libuser32, "GetKeyNameTextA")
	fnGetKeyState, _ = syscall.GetProcAddress(libuser32, "GetKeyState")
	fnGetKeyboardLayout, _ = syscall.GetProcAddress(libuser32, "GetKeyboardLayout")
	fnGetKeyboardLayoutList, _ = syscall.GetProcAddress(libuser32, "GetKeyboardLayoutList")
	fnGetKeyboardLayoutNameW, _ = syscall.GetProcAddress(libuser32, "GetKeyboardLayoutNameW")
	fnGetKeyboardLayoutNameA, _ = syscall.GetProcAddress(libuser32, "GetKeyboardLayoutNameA")
	fnGetMouseMovePointsEx, _ = syscall.GetProcAddress(libuser32, "GetMouseMovePointsEx")
	fnGetKeyboardState, _ = syscall.GetProcAddress(libuser32, "GetKeyboardState")
	fnGetKeyboardType, _ = syscall.GetProcAddress(libuser32, "GetKeyboardType")
	fnGetLastActivePopup, _ = syscall.GetProcAddress(libuser32, "GetLastActivePopup")
	fnGetLastInputInfo, _ = syscall.GetProcAddress(libuser32, "GetLastInputInfo")
	fnGetListBoxInfo, _ = syscall.GetProcAddress(libuser32, "GetListBoxInfo")
	fnGetMenu, _ = syscall.GetProcAddress(libuser32, "GetMenu")
	fnGetMenuBarInfo, _ = syscall.GetProcAddress(libuser32, "GetMenuBarInfo")
	fnGetMenuCheckMarkDimensions, _ = syscall.GetProcAddress(libuser32, "GetMenuCheckMarkDimensions")
	fnGetMenuContextHelpId, _ = syscall.GetProcAddress(libuser32, "GetMenuContextHelpId")
	fnGetMenuDefaultItem, _ = syscall.GetProcAddress(libuser32, "GetMenuDefaultItem")
	fnGetMenuInfo, _ = syscall.GetProcAddress(libuser32, "GetMenuInfo")
	fnGetMenuItemCount, _ = syscall.GetProcAddress(libuser32, "GetMenuItemCount")
	fnGetMenuItemID, _ = syscall.GetProcAddress(libuser32, "GetMenuItemID")
	fnGetMenuItemInfoW, _ = syscall.GetProcAddress(libuser32, "GetMenuItemInfoW")
	fnGetMenuItemInfoA, _ = syscall.GetProcAddress(libuser32, "GetMenuItemInfoA")
	fnGetMenuItemRect, _ = syscall.GetProcAddress(libuser32, "GetMenuItemRect")
	fnGetMenuState, _ = syscall.GetProcAddress(libuser32, "GetMenuState")
	fnGetMenuStringW, _ = syscall.GetProcAddress(libuser32, "GetMenuStringW")
	fnGetMenuStringA, _ = syscall.GetProcAddress(libuser32, "GetMenuStringA")
	fnGetNextDlgGroupItem, _ = syscall.GetProcAddress(libuser32, "GetNextDlgGroupItem")
	fnGetNextDlgTabItem, _ = syscall.GetProcAddress(libuser32, "GetNextDlgTabItem")
	fnGetWindow, _ = syscall.GetProcAddress(libuser32, "GetWindow")
	fnGetOpenClipboardWindow, _ = syscall.GetProcAddress(libuser32, "GetOpenClipboardWindow")
	fnAddClipboardFormatListener, _ = syscall.GetProcAddress(libuser32, "AddClipboardFormatListener")
	fnRemoveClipboardFormatListener, _ = syscall.GetProcAddress(libuser32, "RemoveClipboardFormatListener")
	fnGetUpdatedClipboardFormats, _ = syscall.GetProcAddress(libuser32, "GetUpdatedClipboardFormats")
	fnGetParent, _ = syscall.GetProcAddress(libuser32, "GetParent")
	fnGetPriorityClipboardFormat, _ = syscall.GetProcAddress(libuser32, "GetPriorityClipboardFormat")
	fnGetProcessWindowStation, _ = syscall.GetProcAddress(libuser32, "GetProcessWindowStation")
	fnGetPropW, _ = syscall.GetProcAddress(libuser32, "GetPropW")
	fnGetPropA, _ = syscall.GetProcAddress(libuser32, "GetPropA")
	fnGetQueueStatus, _ = syscall.GetProcAddress(libuser32, "GetQueueStatus")
	fnGetScrollBarInfo, _ = syscall.GetProcAddress(libuser32, "GetScrollBarInfo")
	fnGetScrollInfo, _ = syscall.GetProcAddress(libuser32, "GetScrollInfo")
	fnGetScrollPos, _ = syscall.GetProcAddress(libuser32, "GetScrollPos")
	fnGetScrollRange, _ = syscall.GetProcAddress(libuser32, "GetScrollRange")
	fnGetSubMenu, _ = syscall.GetProcAddress(libuser32, "GetSubMenu")
	fnGetSysColor, _ = syscall.GetProcAddress(libuser32, "GetSysColor")
	fnGetSysColorBrush, _ = syscall.GetProcAddress(libuser32, "GetSysColorBrush")
	fnGetSystemMenu, _ = syscall.GetProcAddress(libuser32, "GetSystemMenu")
	fnGetSystemMetrics, _ = syscall.GetProcAddress(libuser32, "GetSystemMetrics")
	fnGetTabbedTextExtentW, _ = syscall.GetProcAddress(libuser32, "GetTabbedTextExtentW")
	fnGetTabbedTextExtentA, _ = syscall.GetProcAddress(libuser32, "GetTabbedTextExtentA")
	fnGetThreadDesktop, _ = syscall.GetProcAddress(libuser32, "GetThreadDesktop")
	fnGetTitleBarInfo, _ = syscall.GetProcAddress(libuser32, "GetTitleBarInfo")
	fnGetTopWindow, _ = syscall.GetProcAddress(libuser32, "GetTopWindow")
	fnGetTouchInputInfo, _ = syscall.GetProcAddress(libuser32, "GetTouchInputInfo")
	fnGetUpdateRect, _ = syscall.GetProcAddress(libuser32, "GetUpdateRect")
	fnGetUpdateRgn, _ = syscall.GetProcAddress(libuser32, "GetUpdateRgn")
	fnGetUserObjectInformationW, _ = syscall.GetProcAddress(libuser32, "GetUserObjectInformationW")
	fnGetUserObjectInformationA, _ = syscall.GetProcAddress(libuser32, "GetUserObjectInformationA")
	fnGetUserObjectSecurity, _ = syscall.GetProcAddress(libuser32, "GetUserObjectSecurity")
	fnGetWindowContextHelpId, _ = syscall.GetProcAddress(libuser32, "GetWindowContextHelpId")
	fnGetWindowDC, _ = syscall.GetProcAddress(libuser32, "GetWindowDC")
	fnGetWindowDisplayAffinity, _ = syscall.GetProcAddress(libuser32, "GetWindowDisplayAffinity")
	fnGetWindowInfo, _ = syscall.GetProcAddress(libuser32, "GetWindowInfo")
	fnGetWindowModuleFileNameW, _ = syscall.GetProcAddress(libuser32, "GetWindowModuleFileNameW")
	fnGetWindowModuleFileNameA, _ = syscall.GetProcAddress(libuser32, "GetWindowModuleFileNameA")
	fnGetWindowPlacement, _ = syscall.GetProcAddress(libuser32, "GetWindowPlacement")
	fnGetWindowRect, _ = syscall.GetProcAddress(libuser32, "GetWindowRect")
	fnGetWindowRgn, _ = syscall.GetProcAddress(libuser32, "GetWindowRgn")
	fnGetWindowTextW, _ = syscall.GetProcAddress(libuser32, "GetWindowTextW")
	fnGetWindowTextA, _ = syscall.GetProcAddress(libuser32, "GetWindowTextA")
	fnGetWindowTextLengthW, _ = syscall.GetProcAddress(libuser32, "GetWindowTextLengthW")
	fnGetWindowTextLengthA, _ = syscall.GetProcAddress(libuser32, "GetWindowTextLengthA")
	fnGetWindowWord, _ = syscall.GetProcAddress(libuser32, "GetWindowWord")
	fnGrayStringW, _ = syscall.GetProcAddress(libuser32, "GrayStringW")
	fnGrayStringA, _ = syscall.GetProcAddress(libuser32, "GrayStringA")
	fnHideCaret, _ = syscall.GetProcAddress(libuser32, "HideCaret")
	fnHiliteMenuItem, _ = syscall.GetProcAddress(libuser32, "HiliteMenuItem")
	fnImpersonateDdeClientWindow, _ = syscall.GetProcAddress(libuser32, "ImpersonateDdeClientWindow")
	fnInflateRect, _ = syscall.GetProcAddress(libuser32, "InflateRect")
	fnInsertMenuW, _ = syscall.GetProcAddress(libuser32, "InsertMenuW")
	fnInsertMenuA, _ = syscall.GetProcAddress(libuser32, "InsertMenuA")
	fnInsertMenuItemW, _ = syscall.GetProcAddress(libuser32, "InsertMenuItemW")
	fnInsertMenuItemA, _ = syscall.GetProcAddress(libuser32, "InsertMenuItemA")
	fnIntersectRect, _ = syscall.GetProcAddress(libuser32, "IntersectRect")
	fnInvalidateRect, _ = syscall.GetProcAddress(libuser32, "InvalidateRect")
	fnInvalidateRgn, _ = syscall.GetProcAddress(libuser32, "InvalidateRgn")
	fnInvertRect, _ = syscall.GetProcAddress(libuser32, "InvertRect")
	fnIsCharAlphaW, _ = syscall.GetProcAddress(libuser32, "IsCharAlphaW")
	fnIsCharAlphaA, _ = syscall.GetProcAddress(libuser32, "IsCharAlphaA")
	fnIsCharAlphaNumericW, _ = syscall.GetProcAddress(libuser32, "IsCharAlphaNumericW")
	fnIsCharAlphaNumericA, _ = syscall.GetProcAddress(libuser32, "IsCharAlphaNumericA")
	fnIsCharLowerW, _ = syscall.GetProcAddress(libuser32, "IsCharLowerW")
	fnIsCharLowerA, _ = syscall.GetProcAddress(libuser32, "IsCharLowerA")
	fnIsCharUpperW, _ = syscall.GetProcAddress(libuser32, "IsCharUpperW")
	fnIsCharUpperA, _ = syscall.GetProcAddress(libuser32, "IsCharUpperA")
	fnIsChild, _ = syscall.GetProcAddress(libuser32, "IsChild")
	fnIsClipboardFormatAvailable, _ = syscall.GetProcAddress(libuser32, "IsClipboardFormatAvailable")
	fnIsDialogMessageW, _ = syscall.GetProcAddress(libuser32, "IsDialogMessageW")
	fnIsDialogMessageA, _ = syscall.GetProcAddress(libuser32, "IsDialogMessageA")
	fnIsDlgButtonChecked, _ = syscall.GetProcAddress(libuser32, "IsDlgButtonChecked")
	fnIsIconic, _ = syscall.GetProcAddress(libuser32, "IsIconic")
	fnIsMenu, _ = syscall.GetProcAddress(libuser32, "IsMenu")
	fnIsRectEmpty, _ = syscall.GetProcAddress(libuser32, "IsRectEmpty")
	fnIsWindowEnabled, _ = syscall.GetProcAddress(libuser32, "IsWindowEnabled")
	fnIsWindowUnicode, _ = syscall.GetProcAddress(libuser32, "IsWindowUnicode")
	fnIsWindowVisible, _ = syscall.GetProcAddress(libuser32, "IsWindowVisible")
	fnIsZoomed, _ = syscall.GetProcAddress(libuser32, "IsZoomed")
	fnKillTimer, _ = syscall.GetProcAddress(libuser32, "KillTimer")
	fnLoadAcceleratorsW, _ = syscall.GetProcAddress(libuser32, "LoadAcceleratorsW")
	fnLoadAcceleratorsA, _ = syscall.GetProcAddress(libuser32, "LoadAcceleratorsA")
	fnLoadBitmapW, _ = syscall.GetProcAddress(libuser32, "LoadBitmapW")
	fnLoadBitmapA, _ = syscall.GetProcAddress(libuser32, "LoadBitmapA")
	fnLoadCursorFromFileW, _ = syscall.GetProcAddress(libuser32, "LoadCursorFromFileW")
	fnLoadCursorFromFileA, _ = syscall.GetProcAddress(libuser32, "LoadCursorFromFileA")
	fnLoadIconW, _ = syscall.GetProcAddress(libuser32, "LoadIconW")
	fnLoadIconA, _ = syscall.GetProcAddress(libuser32, "LoadIconA")
	fnLoadImageW, _ = syscall.GetProcAddress(libuser32, "LoadImageW")
	fnLoadImageA, _ = syscall.GetProcAddress(libuser32, "LoadImageA")
	fnLoadKeyboardLayoutW, _ = syscall.GetProcAddress(libuser32, "LoadKeyboardLayoutW")
	fnLoadKeyboardLayoutA, _ = syscall.GetProcAddress(libuser32, "LoadKeyboardLayoutA")
	fnLoadMenuW, _ = syscall.GetProcAddress(libuser32, "LoadMenuW")
	fnLoadMenuA, _ = syscall.GetProcAddress(libuser32, "LoadMenuA")
	fnLoadMenuIndirectW, _ = syscall.GetProcAddress(libuser32, "LoadMenuIndirectW")
	fnLoadMenuIndirectA, _ = syscall.GetProcAddress(libuser32, "LoadMenuIndirectA")
	fnLockWindowUpdate, _ = syscall.GetProcAddress(libuser32, "LockWindowUpdate")
	fnLockWorkStation, _ = syscall.GetProcAddress(libuser32, "LockWorkStation")
	fnLookupIconIdFromDirectory, _ = syscall.GetProcAddress(libuser32, "LookupIconIdFromDirectory")
	fnLookupIconIdFromDirectoryEx, _ = syscall.GetProcAddress(libuser32, "LookupIconIdFromDirectoryEx")
	fnMapDialogRect, _ = syscall.GetProcAddress(libuser32, "MapDialogRect")
	fnMapVirtualKeyW, _ = syscall.GetProcAddress(libuser32, "MapVirtualKeyW")
	fnMapVirtualKeyA, _ = syscall.GetProcAddress(libuser32, "MapVirtualKeyA")
	fnMapVirtualKeyExW, _ = syscall.GetProcAddress(libuser32, "MapVirtualKeyExW")
	fnMapVirtualKeyExA, _ = syscall.GetProcAddress(libuser32, "MapVirtualKeyExA")
	fnMapWindowPoints, _ = syscall.GetProcAddress(libuser32, "MapWindowPoints")
	fnMenuItemFromPoint, _ = syscall.GetProcAddress(libuser32, "MenuItemFromPoint")
	fnMessageBeep, _ = syscall.GetProcAddress(libuser32, "MessageBeep")
	fnMessageBoxW, _ = syscall.GetProcAddress(libuser32, "MessageBoxW")
	fnMessageBoxA, _ = syscall.GetProcAddress(libuser32, "MessageBoxA")
	fnMessageBoxExW, _ = syscall.GetProcAddress(libuser32, "MessageBoxExW")
	fnMessageBoxExA, _ = syscall.GetProcAddress(libuser32, "MessageBoxExA")
	fnMessageBoxIndirectW, _ = syscall.GetProcAddress(libuser32, "MessageBoxIndirectW")
	fnMessageBoxIndirectA, _ = syscall.GetProcAddress(libuser32, "MessageBoxIndirectA")
	fnModifyMenuW, _ = syscall.GetProcAddress(libuser32, "ModifyMenuW")
	fnModifyMenuA, _ = syscall.GetProcAddress(libuser32, "ModifyMenuA")
	fnMoveWindow, _ = syscall.GetProcAddress(libuser32, "MoveWindow")
	fnOemKeyScan, _ = syscall.GetProcAddress(libuser32, "OemKeyScan")
	fnOemToCharA, _ = syscall.GetProcAddress(libuser32, "OemToCharA")
	fnOemToCharBuffA, _ = syscall.GetProcAddress(libuser32, "OemToCharBuffA")
	fnOemToCharW, _ = syscall.GetProcAddress(libuser32, "OemToCharW")
	fnGetWindowThreadProcessId, _ = syscall.GetProcAddress(libkernel32, "GetWindowThreadProcessId")
	fnOemToCharBuffW, _ = syscall.GetProcAddress(libuser32, "OemToCharBuffW")
	fnOffsetRect, _ = syscall.GetProcAddress(libuser32, "OffsetRect")
	fnOpenClipboard, _ = syscall.GetProcAddress(libuser32, "OpenClipboard")
	fnOpenDesktopW, _ = syscall.GetProcAddress(libuser32, "OpenDesktopW")
	fnOpenDesktopA, _ = syscall.GetProcAddress(libuser32, "OpenDesktopA")
	fnOpenIcon, _ = syscall.GetProcAddress(libuser32, "OpenIcon")
	fnOpenInputDesktop, _ = syscall.GetProcAddress(libuser32, "OpenInputDesktop")
	fnOpenWindowStationW, _ = syscall.GetProcAddress(libuser32, "OpenWindowStationW")
	fnOpenWindowStationA, _ = syscall.GetProcAddress(libuser32, "OpenWindowStationA")
	fnPackDDElParam, _ = syscall.GetProcAddress(libuser32, "PackDDElParam")
	fnPaintDesktop, _ = syscall.GetProcAddress(libuser32, "PaintDesktop")
	fnPtInRect, _ = syscall.GetProcAddress(libuser32, "PtInRect")
	fnRealChildWindowFromPoint, _ = syscall.GetProcAddress(libuser32, "RealChildWindowFromPoint")
	fnRealGetWindowClassW, _ = syscall.GetProcAddress(libuser32, "RealGetWindowClassW")
	fnRealGetWindowClassA, _ = syscall.GetProcAddress(libuser32, "RealGetWindowClassA")
	fnRedrawWindow, _ = syscall.GetProcAddress(libuser32, "RedrawWindow")
	fnRegisterClassW, _ = syscall.GetProcAddress(libuser32, "RegisterClassW")
	fnRegisterClassA, _ = syscall.GetProcAddress(libuser32, "RegisterClassA")
	fnRegisterClassExW, _ = syscall.GetProcAddress(libuser32, "RegisterClassExW")
	fnRegisterClassExA, _ = syscall.GetProcAddress(libuser32, "RegisterClassExA")
	fnCreateWindowExW, _ = syscall.GetProcAddress(libuser32, "CreateWindowExW")
	fnCreateWindowExA, _ = syscall.GetProcAddress(libuser32, "CreateWindowExA")
	fnRegisterClipboardFormatW, _ = syscall.GetProcAddress(libuser32, "RegisterClipboardFormatW")
	fnRegisterClipboardFormatA, _ = syscall.GetProcAddress(libuser32, "RegisterClipboardFormatA")
	fnRegisterDeviceNotificationW, _ = syscall.GetProcAddress(libuser32, "RegisterDeviceNotificationW")
	fnRegisterDeviceNotificationA, _ = syscall.GetProcAddress(libuser32, "RegisterDeviceNotificationA")
	fnRegisterHotKey, _ = syscall.GetProcAddress(libuser32, "RegisterHotKey")
	fnRegisterPowerSettingNotification, _ = syscall.GetProcAddress(libuser32, "RegisterPowerSettingNotification")
	fnRegisterTouchWindow, _ = syscall.GetProcAddress(libuser32, "RegisterTouchWindow")
	fnRegisterWindowMessageW, _ = syscall.GetProcAddress(libuser32, "RegisterWindowMessageW")
	fnRegisterWindowMessageA, _ = syscall.GetProcAddress(libuser32, "RegisterWindowMessageA")
	fnReleaseCapture, _ = syscall.GetProcAddress(libuser32, "ReleaseCapture")
	fnReleaseDC, _ = syscall.GetProcAddress(libuser32, "ReleaseDC")
	fnRemoveMenu, _ = syscall.GetProcAddress(libuser32, "RemoveMenu")
	fnRemovePropW, _ = syscall.GetProcAddress(libuser32, "RemovePropW")
	fnRemovePropA, _ = syscall.GetProcAddress(libuser32, "RemovePropA")
	fnReplyMessage, _ = syscall.GetProcAddress(libuser32, "ReplyMessage")
	fnReuseDDElParam, _ = syscall.GetProcAddress(libuser32, "ReuseDDElParam")
	fnScreenToClient, _ = syscall.GetProcAddress(libuser32, "ScreenToClient")
	fnLogicalToPhysicalPoint, _ = syscall.GetProcAddress(libuser32, "LogicalToPhysicalPoint")
	fnPhysicalToLogicalPoint, _ = syscall.GetProcAddress(libuser32, "PhysicalToLogicalPoint")
	fnScrollDC, _ = syscall.GetProcAddress(libuser32, "ScrollDC")
	fnScrollWindow, _ = syscall.GetProcAddress(libuser32, "ScrollWindow")
	fnSendDlgItemMessageW, _ = syscall.GetProcAddress(libuser32, "SendDlgItemMessageW")
	fnSendDlgItemMessageA, _ = syscall.GetProcAddress(libuser32, "SendDlgItemMessageA")
	fnSendInput, _ = syscall.GetProcAddress(libuser32, "SendInput")
	fnSetActiveWindow, _ = syscall.GetProcAddress(libuser32, "SetActiveWindow")
	fnSetCapture, _ = syscall.GetProcAddress(libuser32, "SetCapture")
	fnSetCaretBlinkTime, _ = syscall.GetProcAddress(libuser32, "SetCaretBlinkTime")
	fnSetCaretPos, _ = syscall.GetProcAddress(libuser32, "SetCaretPos")
	fnSetClassWord, _ = syscall.GetProcAddress(libuser32, "SetClassWord")
	fnSetClipboardData, _ = syscall.GetProcAddress(libuser32, "SetClipboardData")
	fnSetClipboardViewer, _ = syscall.GetProcAddress(libuser32, "SetClipboardViewer")
	fnSetCursor, _ = syscall.GetProcAddress(libuser32, "SetCursor")
	fnSetCursorPos, _ = syscall.GetProcAddress(libuser32, "SetCursorPos")
	fnSetDlgItemInt, _ = syscall.GetProcAddress(libuser32, "SetDlgItemInt")
	fnSetDlgItemTextW, _ = syscall.GetProcAddress(libuser32, "SetDlgItemTextW")
	fnSetDlgItemTextA, _ = syscall.GetProcAddress(libuser32, "SetDlgItemTextA")
	fnSetDoubleClickTime, _ = syscall.GetProcAddress(libuser32, "SetDoubleClickTime")
	fnSetFocus, _ = syscall.GetProcAddress(libuser32, "SetFocus")
	fnSetForegroundWindow, _ = syscall.GetProcAddress(libuser32, "SetForegroundWindow")
	fnAllowSetForegroundWindow, _ = syscall.GetProcAddress(libuser32, "AllowSetForegroundWindow")
	fnLockSetForegroundWindow, _ = syscall.GetProcAddress(libuser32, "LockSetForegroundWindow")
	fnSetKeyboardState, _ = syscall.GetProcAddress(libuser32, "SetKeyboardState")
	fnSetMenu, _ = syscall.GetProcAddress(libuser32, "SetMenu")
	fnSetMenuContextHelpId, _ = syscall.GetProcAddress(libuser32, "SetMenuContextHelpId")
	fnSetMenuDefaultItem, _ = syscall.GetProcAddress(libuser32, "SetMenuDefaultItem")
	fnSetMenuInfo, _ = syscall.GetProcAddress(libuser32, "SetMenuInfo")
	fnSetMenuItemBitmaps, _ = syscall.GetProcAddress(libuser32, "SetMenuItemBitmaps")
	fnSetMenuItemInfoW, _ = syscall.GetProcAddress(libuser32, "SetMenuItemInfoW")
	fnSetMenuItemInfoA, _ = syscall.GetProcAddress(libuser32, "SetMenuItemInfoA")
	fnSetMessageExtraInfo, _ = syscall.GetProcAddress(libuser32, "SetMessageExtraInfo")
	fnSetMessageQueue, _ = syscall.GetProcAddress(libuser32, "SetMessageQueue")
	fnSetParent, _ = syscall.GetProcAddress(libuser32, "SetParent")
	fnSetPhysicalCursorPos, _ = syscall.GetProcAddress(libuser32, "SetPhysicalCursorPos")
	fnSetProcessDPIAware, _ = syscall.GetProcAddress(libuser32, "SetProcessDPIAware")
	fnSetProcessWindowStation, _ = syscall.GetProcAddress(libuser32, "SetProcessWindowStation")
	fnSetPropW, _ = syscall.GetProcAddress(libuser32, "SetPropW")
	fnSetPropA, _ = syscall.GetProcAddress(libuser32, "SetPropA")
	fnSetRect, _ = syscall.GetProcAddress(libuser32, "SetRect")
	fnSetRectEmpty, _ = syscall.GetProcAddress(libuser32, "SetRectEmpty")
	fnSetScrollInfo, _ = syscall.GetProcAddress(libuser32, "SetScrollInfo")
	fnSetScrollPos, _ = syscall.GetProcAddress(libuser32, "SetScrollPos")
	fnSetScrollRange, _ = syscall.GetProcAddress(libuser32, "SetScrollRange")
	fnSetSysColors, _ = syscall.GetProcAddress(libuser32, "SetSysColors")
	fnSetSystemCursor, _ = syscall.GetProcAddress(libuser32, "SetSystemCursor")
	fnSetThreadDesktop, _ = syscall.GetProcAddress(libuser32, "SetThreadDesktop")
	fnSetTimer, _ = syscall.GetProcAddress(libuser32, "SetTimer")
	fnSetUserObjectInformationW, _ = syscall.GetProcAddress(libuser32, "SetUserObjectInformationW")
	fnSetUserObjectInformationA, _ = syscall.GetProcAddress(libuser32, "SetUserObjectInformationA")
	fnSetUserObjectSecurity, _ = syscall.GetProcAddress(libuser32, "SetUserObjectSecurity")
	fnSetWindowContextHelpId, _ = syscall.GetProcAddress(libuser32, "SetWindowContextHelpId")
	fnSetWindowDisplayAffinity, _ = syscall.GetProcAddress(libuser32, "SetWindowDisplayAffinity")
	fnSetWindowPlacement, _ = syscall.GetProcAddress(libuser32, "SetWindowPlacement")
	fnSetWindowPos, _ = syscall.GetProcAddress(libuser32, "SetWindowPos")
	fnSetWindowTextW, _ = syscall.GetProcAddress(libuser32, "SetWindowTextW")
	fnSetWindowTextA, _ = syscall.GetProcAddress(libuser32, "SetWindowTextA")
	fnSetWindowWord, _ = syscall.GetProcAddress(libuser32, "SetWindowWord")
	fnSetWindowsHookW, _ = syscall.GetProcAddress(libuser32, "SetWindowsHookW")
	fnSetWindowsHookA, _ = syscall.GetProcAddress(libuser32, "SetWindowsHookA")
	fnSetWindowsHookExW, _ = syscall.GetProcAddress(libuser32, "SetWindowsHookExW")
	fnSetWindowsHookExA, _ = syscall.GetProcAddress(libuser32, "SetWindowsHookExA")
	fnSetWindowRgn, _ = syscall.GetProcAddress(libuser32, "SetWindowRgn")
	fnSetWinEventHook, _ = syscall.GetProcAddress(libuser32, "SetWinEventHook")
	fnShowCaret, _ = syscall.GetProcAddress(libuser32, "ShowCaret")
	fnShowCursor, _ = syscall.GetProcAddress(libuser32, "ShowCursor")
	fnShowOwnedPopups, _ = syscall.GetProcAddress(libuser32, "ShowOwnedPopups")
	fnShowScrollBar, _ = syscall.GetProcAddress(libuser32, "ShowScrollBar")
	fnShowWindow, _ = syscall.GetProcAddress(libuser32, "ShowWindow")
	fnAnimateWindow, _ = syscall.GetProcAddress(libuser32, "AnimateWindow")
	fnShowWindowAsync, _ = syscall.GetProcAddress(libuser32, "ShowWindowAsync")
	fnSubtractRect, _ = syscall.GetProcAddress(libuser32, "SubtractRect")
	fnSwapMouseButton, _ = syscall.GetProcAddress(libuser32, "SwapMouseButton")
	fnSwitchDesktop, _ = syscall.GetProcAddress(libuser32, "SwitchDesktop")
	fnSoundSentry, _ = syscall.GetProcAddress(libuser32, "SoundSentry")
	fnTabbedTextOutW, _ = syscall.GetProcAddress(libuser32, "TabbedTextOutW")
	fnTabbedTextOutA, _ = syscall.GetProcAddress(libuser32, "TabbedTextOutA")
	fnTileWindows, _ = syscall.GetProcAddress(libuser32, "TileWindows")
	fnToAscii, _ = syscall.GetProcAddress(libuser32, "ToAscii")
	fnToAsciiEx, _ = syscall.GetProcAddress(libuser32, "ToAsciiEx")
	fnTrackMouseEvent, _ = syscall.GetProcAddress(libuser32, "TrackMouseEvent")
	fnTrackPopupMenu, _ = syscall.GetProcAddress(libuser32, "TrackPopupMenu")
	fnTrackPopupMenuEx, _ = syscall.GetProcAddress(libuser32, "TrackPopupMenuEx")
	fnEndMenu,_= syscall.GetProcAddress(libuser32, "EndMenu")
	fnTranslateAcceleratorW, _ = syscall.GetProcAddress(libuser32, "TranslateAcceleratorW")
	fnTranslateAcceleratorA, _ = syscall.GetProcAddress(libuser32, "TranslateAcceleratorA")
	fnTranslateMDISysAccel, _ = syscall.GetProcAddress(libuser32, "TranslateMDISysAccel")
	fnUnhookWindowsHook, _ = syscall.GetProcAddress(libuser32, "UnhookWindowsHook")
	fnUnhookWindowsHookEx, _ = syscall.GetProcAddress(libuser32, "UnhookWindowsHookEx")
	fnUnhookWinEvent, _ = syscall.GetProcAddress(libuser32, "UnhookWinEvent")
	fnUnionRect, _ = syscall.GetProcAddress(libuser32, "UnionRect")
	fnUnloadKeyboardLayout, _ = syscall.GetProcAddress(libuser32, "UnloadKeyboardLayout")
	fnUnpackDDElParam, _ = syscall.GetProcAddress(libuser32, "UnpackDDElParam")
	fnUnregisterClassW, _ = syscall.GetProcAddress(libuser32, "UnregisterClassW")
	fnUnregisterClassA, _ = syscall.GetProcAddress(libuser32, "UnregisterClassA")
	fnUnregisterDeviceNotification, _ = syscall.GetProcAddress(libuser32, "UnregisterDeviceNotification")
	fnUnregisterPowerSettingNotification, _ = syscall.GetProcAddress(libuser32, "UnregisterPowerSettingNotification")
	fnUnregisterHotKey, _ = syscall.GetProcAddress(libuser32, "UnregisterHotKey")
	fnUnregisterTouchWindow, _ = syscall.GetProcAddress(libuser32, "UnregisterTouchWindow")
	fnUpdateWindow, _ = syscall.GetProcAddress(libuser32, "UpdateWindow")
	fnUpdateLayeredWindow, _ = syscall.GetProcAddress(libuser32, "UpdateLayeredWindow")
	fnUserHandleGrantAccess, _ = syscall.GetProcAddress(libuser32, "UserHandleGrantAccess")
	fnValidateRect, _ = syscall.GetProcAddress(libuser32, "ValidateRect")
	fnValidateRgn, _ = syscall.GetProcAddress(libuser32, "ValidateRgn")
	fnVkKeyScanW, _ = syscall.GetProcAddress(libuser32, "VkKeyScanW")
	fnVkKeyScanA, _ = syscall.GetProcAddress(libuser32, "VkKeyScanA")
	fnVkKeyScanExW, _ = syscall.GetProcAddress(libuser32, "VkKeyScanExW")
	fnVkKeyScanExA, _ = syscall.GetProcAddress(libuser32, "VkKeyScanExA")
	fnWaitForInputIdle, _ = syscall.GetProcAddress(libuser32, "WaitForInputIdle")
	fnWaitMessage, _ = syscall.GetProcAddress(libuser32, "WaitMessage")
	fnWindowFromDC, _ = syscall.GetProcAddress(libuser32, "WindowFromDC")
	fnWindowFromPoint, _ = syscall.GetProcAddress(libuser32, "WindowFromPoint")
	fnWindowFromPhysicalPoint, _ = syscall.GetProcAddress(libuser32, "WindowFromPhysicalPoint")
	fnwsprintfW, _ = syscall.GetProcAddress(libuser32, "wsprintfW")
	fnwsprintfA, _ = syscall.GetProcAddress(libuser32, "wsprintfA")
	fnwvsprintfW, _ = syscall.GetProcAddress(libuser32, "wvsprintfW")
	fnwvsprintfA, _ = syscall.GetProcAddress(libuser32, "wvsprintfA")
	fnGetWindowLongPtrW, _ = syscall.GetProcAddress(libuser32, "GetWindowLongPtrW")
	fnSetWindowLongPtrW, _ = syscall.GetProcAddress(libuser32, "SetWindowLongPtrW")
	fnGetClassLongPtrW, _ = syscall.GetProcAddress(libuser32, "GetClassLongPtrW")
	fnSetClassLongPtrW, _ = syscall.GetProcAddress(libuser32, "SetClassLongPtrW")
	fnGetWindowLongPtrA, _ = syscall.GetProcAddress(libuser32, "GetWindowLongPtrA")
	fnSetWindowLongPtrA, _ = syscall.GetProcAddress(libuser32, "SetWindowLongPtrA")
	fnGetClassLongPtrA, _ = syscall.GetProcAddress(libuser32, "GetClassLongPtrA")
	fnSetClassLongPtrA, _ = syscall.GetProcAddress(libuser32, "SetClassLongPtrA")
	fnGetWindowLongW, _ = syscall.GetProcAddress(libuser32, "GetWindowLongW")
	fnSetWindowLongW, _ = syscall.GetProcAddress(libuser32, "SetWindowLongW")
	fnGetClassLongW, _ = syscall.GetProcAddress(libuser32, "GetClassLongW")
	fnSetClassLongW, _ = syscall.GetProcAddress(libuser32, "SetClassLongW")
	fnGetWindowLongA, _ = syscall.GetProcAddress(libuser32, "GetWindowLongA")
	fnSetWindowLongA, _ = syscall.GetProcAddress(libuser32, "SetWindowLongA")
	fnGetClassLongA, _ = syscall.GetProcAddress(libuser32, "GetClassLongA")
	fnSetClassLongA, _ = syscall.GetProcAddress(libuser32, "SetClassLongA")
	fnGetGestureInfo, _ = syscall.GetProcAddress(libuser32, "GetGestureInfo")
	fnGetGestureExtraArgs, _ = syscall.GetProcAddress(libuser32, "GetGestureExtraArgs")
	fnCloseGestureInfoHandle, _ = syscall.GetProcAddress(libuser32, "CloseGestureInfoHandle")
	fnSetGestureConfig, _ = syscall.GetProcAddress(libuser32, "SetGestureConfig")
	fnGetGestureConfig, _ = syscall.GetProcAddress(libuser32, "GetGestureConfig")
	fnWINNLSEnableIME, _ = syscall.GetProcAddress(libuser32, "WINNLSEnableIME")

	fnPostMessageA, _ = syscall.GetProcAddress(libuser32, "PostMessageA")
	fnPostMessageW, _ = syscall.GetProcAddress(libuser32, "PostMessageW")

	InitScreenLogPixels()
}

func BoolToUint(value bool) uint8 {
	if value {
		return 1
	}

	return 0
}

func GetCapture() syscall.Handle  {
	r, _, _ := syscall.Syscall(fnGetCapture,0,0,0,0)
	return syscall.Handle(r)
}

func ReleaseCapture()bool  {
	r, _, _ := syscall.Syscall(fnReleaseCapture,0,0,0,0)
	return r != 0
}

func SetCapture(hwnd syscall.Handle)syscall.Handle  {
	r, _, _ := syscall.Syscall(fnSetCapture,1,uintptr(hwnd),0,0)
	return syscall.Handle(r)
}

func MessageBox(hwnd syscall.Handle, lpText string, lpCaption string, uType uint) int {
	r, _, _ := syscall.Syscall6(fnMessageBoxW, 4, uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpText))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpCaption))),
		uintptr(uType), 0, 0)
	return int(r)
}

func RegisterWindowMessage(msgInfo string) uint  {
	ret, _, _ := syscall.Syscall(fnRegisterWindowMessageW,1,uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(msgInfo))),0,0)
	return uint(ret)
}

func CreateWindow(ClassName, WindowName string, dwStyle uint32, x, y, width, height int, hWndParent syscall.Handle, hMenu HMENU, hInstance syscall.Handle, lpParam uintptr) (result syscall.Handle) {
	ret, _, _ := syscall.Syscall12(fnCreateWindowExW, 11, 0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(ClassName))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(WindowName))),
		uintptr(dwStyle), uintptr(x), uintptr(y), uintptr(width), uintptr(height),
		uintptr(hWndParent), uintptr(hMenu), uintptr(hInstance), lpParam)
	result = syscall.Handle(ret)
	return
}

func RegisterClass(wndClass *GWndClass) ATOM {
	ret, _, _ := syscall.Syscall(fnRegisterClassW, 1,
		uintptr(unsafe.Pointer(wndClass)),
		0,
		0)
	return ATOM(ret)
}

func RegisterClassEx(wndClass *GWndClassEx) ATOM {
	ret, _, _ := syscall.Syscall(fnRegisterClassExW, 1,
		uintptr(unsafe.Pointer(wndClass)),
		0,
		0)
	return ATOM(ret)
}

func GetFocus() syscall.Handle {
	ret, _, _ := syscall.Syscall(fnGetFocus, 0,
		0,
		0,
		0)
	return syscall.Handle(ret)
}

func SetFocus(focusControl syscall.Handle) syscall.Handle {
	ret, _, _ := syscall.Syscall(fnSetFocus, 1,
		uintptr(focusControl),
		0,
		0)
	return syscall.Handle(ret)
}

func SystemParametersInfo(uiAction, uiParam uint,pvParam uintptr,fWinIni uint)bool  {
	ret, _, _ := syscall.Syscall6(fnSystemParametersInfo, 4,
		uintptr(uiAction), uintptr(uiParam), pvParam,uintptr(fWinIni),0,0)
	return ret!=0
}

func GetScreenWorkArea()*Rect  {
	var Result Rect
	if SystemParametersInfo(SPI_GETWORKAREA, 0, uintptr(unsafe.Pointer(&Result)), 0){
		return &Result
	}
	return nil
}

func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(fnTranslateMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret != 0
}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.Syscall(fnDispatchMessageW, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret
}

func PeekMessage(lpMsg *MSG, hWnd syscall.Handle, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := syscall.Syscall6(fnPeekMessageW, 5,
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg),
		0)

	return ret != 0
}

func WaitMessage() bool {
	ret, _, _ := syscall.Syscall6(fnWaitMessage, 0,
		0,
		0,
		0,
		0,
		0,
		0)
	return ret != 0
}

func CallWindowProc(lpPrevWndFunc uintptr, hWnd syscall.Handle, Msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(fnCallWindowProcW, 5,
		lpPrevWndFunc,
		uintptr(hWnd),
		uintptr(Msg),
		wParam,
		lParam,
		0)

	return ret
}

func DefWindowProc(hWnd syscall.Handle, Msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(fnDefWindowProcW, 4,
		uintptr(hWnd),
		uintptr(Msg),
		wParam,
		lParam,
		0,
		0)
	return ret
}

func GetKeyState(nVirtKey int)int16  {
	ret,_,_ := syscall.Syscall(fnGetKeyState,1,uintptr(nVirtKey),0,0)
	return int16(ret)
}

func SetWindowPos(hWnd, hWndInsertAfter syscall.Handle, x, y, cx, cy int32, uFlags uint32) bool {
	ret, _, _ := syscall.Syscall9(fnSetWindowPos, 7,
		uintptr(hWnd),
		uintptr(hWndInsertAfter),
		uintptr(x), uintptr(y),
		uintptr(cx), uintptr(cy),
		uintptr(uFlags),
		0, 0)
	return ret != 0
}

func LoadCursor(hInstance HINST, lpCursorName *uint16) HCURSOR {
	ret, _, _ := syscall.Syscall(fnLoadCursor, 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpCursorName)),
		0)
	return HCURSOR(ret)
}

func GetClassInfo(hInstance HINST, lpClassName string, lpWndClass *GWndClass) bool {
	ret, _, _ := syscall.Syscall(fnGetClassInfoW, 3,
		uintptr(hInstance), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpClassName))),
		uintptr(unsafe.Pointer(lpWndClass)))
	return ret != 0
}

func GetClassInfoEx(hInstance HINST, lpClassName string, lpWndClass *GWndClassEx) bool {
	ret, _, _ := syscall.Syscall(fnGetClassInfoExW, 3,
		uintptr(hInstance), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpClassName))),
		uintptr(unsafe.Pointer(lpWndClass)))
	return ret != 0
}

func GetWindowThreadProcessId(hWnd syscall.Handle,dwProcessId *uint32) uint32 {
	r1, _, _ := syscall.Syscall(fnGetWindowThreadProcessId, 2, uintptr(hWnd), uintptr(unsafe.Pointer(dwProcessId)), 0)
	return uint32(r1)
}

func UnregisterClass(lpClassName string, hinst HINST) bool {
	ret, _, _ := syscall.Syscall(fnUnregisterClassW, 2, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpClassName))),
		uintptr(hinst), 0)
	return ret != 0
}

func SetProp(hwd syscall.Handle, propName uintptr, propData uintptr) bool {
	ret, _, _ := syscall.Syscall(fnSetPropW, 3, uintptr(hwd), propName, propData)
	return ret != 0
}

func GetProp(hwd syscall.Handle, propName uintptr) uintptr {
	ret, _, _ := syscall.Syscall(fnGetPropW, 2, uintptr(hwd), propName, 0)
	return uintptr(ret)
}

func RemoveProp(hWnd syscall.Handle, propName uintptr) syscall.Handle {
	ret, _, _ := syscall.Syscall(fnRemovePropW, 2, uintptr(hWnd), propName, 0)
	return syscall.Handle(ret)
}

func CreateWindowEx(dwExStyle uint32, lpClassName, lpWindowName string, dwStyle uint32, x, y, nWidth, nHeight int32, hWndParent syscall.Handle, hMenu HMENU, hInstance HINST, lpParam unsafe.Pointer) syscall.Handle {
	ret, _, _ := syscall.Syscall12(fnCreateWindowExW, 12,
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpClassName))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpWindowName))),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam))
	return syscall.Handle(ret)
}

func CreateWindowExptr(dwExStyle uint32, lpClassName, lpWindowName *uint16, dwStyle uint32, x, y, nWidth, nHeight int32, hWndParent syscall.Handle, hMenu HMENU, hInstance HINST, lpParam unsafe.Pointer) syscall.Handle {
	ret, _, _ := syscall.Syscall12(fnCreateWindowExW, 12,
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam))
	return syscall.Handle(ret)
}

func ShowWindow(hWnd syscall.Handle, nCmdShow int32) bool {
	ret, _, _ := syscall.Syscall(fnShowWindow, 2,
		uintptr(hWnd),
		uintptr(nCmdShow),
		0)

	return ret != 0
}

func UpdateWindow(hwnd syscall.Handle) bool {
	ret, _, _ := syscall.Syscall(fnUpdateWindow, 1,
		uintptr(hwnd),
		0,
		0)

	return ret != 0
}
func WindowFromPoint(Point POINT) syscall.Handle {
	ret, _, _ := syscall.Syscall(fnWindowFromPoint, 2,
		uintptr(Point.X),
		uintptr(Point.Y),
		0)

	return syscall.Handle(ret)
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(fnPostQuitMessage, 1,
		uintptr(exitCode),
		0,
		0)
}

func SetWindowText(hwnd syscall.Handle, txt string) bool {
	ret, _, _ := syscall.Syscall(fnSetWindowTextW, 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(txt))),
		0)

	return ret != 0
}

func GetWindowText(hwnd syscall.Handle) string {
	retlen, _, _ := syscall.Syscall(fnGetWindowTextLengthW, 1,
		uintptr(hwnd),
		0,
		0)
	mp := make([]uint16, retlen+1)
	retlen, _, _ = syscall.Syscall(fnGetWindowTextW, 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&mp[0])),
		retlen+1)
	return syscall.UTF16ToString(mp)
}

func GetWindowLong(hwnd syscall.Handle, nIdex int32) int {
	ret, _, _ := syscall.Syscall(fnGetWindowLongW, 2, uintptr(hwnd), uintptr(nIdex), 0)
	return int(ret)
}

func ActivateKeyboardLayout(hkl HKL, Flags uint) HKL {
	ret, _, _ := syscall.Syscall(fnActivateKeyboardLayout, 2, uintptr(hkl), uintptr(Flags), 0)
	return HKL(ret)
}

func AnyPopup()bool  {
	ret, _, _ := syscall.Syscall(fnAnyPopup, 0, 0, 0, 0)
	return  ret!=0
}

func AppendMenu(hMenu HMENU,uFlags uint, uIDNewItem uintptr, lpNewItem string) bool{
	ret, _, _ := syscall.Syscall6(fnAppendMenuW, 4, uintptr(hMenu),
		uintptr(uFlags),uIDNewItem,uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpNewItem))),0,0)
	return  ret!=0
}

func ArrangeIconicWindows(hwnd syscall.Handle)uint  {
	ret, _, _ := syscall.Syscall(fnArrangeIconicWindows, 1, uintptr(hwnd), 0, 0)
	return  uint(ret)
}

func AttachThreadInput(idAttach, idAttachTo uint32,fAttach bool)uint  {
	ret, _, _ := syscall.Syscall(fnAttachThreadInput, 3, uintptr(idAttach), uintptr(idAttachTo), uintptr(BoolToUint(fAttach)))
	return  uint(ret)
}

func BeginDeferWindowPos(nNumWindows int)syscall.Handle  {
	ret, _, _ := syscall.Syscall(fnBeginDeferWindowPos, 1, uintptr(nNumWindows),0,0)
	return  syscall.Handle(ret)
}

func BeginPaint(hWnd syscall.Handle, lpPaint *GPaintStruct) HDC{
	ret, _, _ := syscall.Syscall(fnBeginPaint, 2, uintptr(hWnd),uintptr(unsafe.Pointer(lpPaint)),0)
	return  HDC(ret)
}

func EndPaint(hWnd syscall.Handle,pstruct *GPaintStruct) bool{
	ret, _, _ := syscall.Syscall(fnEndPaint, 2, uintptr(hWnd),uintptr(unsafe.Pointer(pstruct)),0)
	return  ret!=0
}

func BringWindowToTop(hwnd syscall.Handle) bool{
	ret, _, _ := syscall.Syscall(fnBringWindowToTop, 1, uintptr(hwnd),0,0)
	return  ret!=0
}

func BroadcastSystemMessage(Flags uint32, Recipients uintptr, uiMessage uint, wParam uintptr, lParam uintptr)int32{
	ret, _, _ := syscall.Syscall6(fnBroadcastSystemMessageW, 5, uintptr(Flags),Recipients,
		uintptr(uiMessage),wParam,lParam,0)
	return  int32(ret)
}

func CallMsgFilter(lpmsg *MSG,nCode int32)bool{
	ret,_,_:=syscall.Syscall(fnCallMsgFilterW,2,uintptr(unsafe.Pointer(lpmsg)),uintptr(nCode),0)
	return ret!=0
}

func CallNextHookEx(hhk HOOK,nCode int,wParam,lParam uintptr)  LRESULT{
	ret,_,_:=syscall.Syscall6(fnCallNextHookEx,4,uintptr(hhk),uintptr(nCode),wParam,lParam,0,0)
	return LRESULT(ret)
}


func CascadeWindows(hwndParent syscall.Handle,wHow uint,lpRect *Rect,ckids uint,lpkids unsafe.Pointer) uint16 {
	ret,_,_:=syscall.Syscall6(fnCascadeWindows,5,uintptr(hwndParent),uintptr(wHow),
		uintptr(unsafe.Pointer(lpRect)),uintptr(ckids),uintptr(lpkids),0)
	return uint16(ret)
}

func ChangeWindowMessageFilter(message uint,dwFlag uint32)bool  {
	ret,_,_:=syscall.Syscall(fnChangeWindowMessageFilter,2,uintptr(message),uintptr(dwFlag),0)
	return ret!=0
}

func ChangeClipboardChain(hWndRemove, hWndNewNext syscall.Handle)bool  {
	ret,_,_:=syscall.Syscall(fnChangeClipboardChain,2,uintptr(hWndRemove),uintptr(hWndNewNext),0)
	return ret!=0
}

func ChangeDisplaySettings(lpDevMode *GdevicemodeW,dwFlags uint32)int32  {
	ret,_,_ := syscall.Syscall(fnChangeDisplaySettingsW,2,uintptr(unsafe.Pointer(lpDevMode)),uintptr(dwFlags),0)
	return int32(ret)
}
func ChangeDisplaySettingsEx(lpszDeviceName string,lpDevMode *GdevicemodeW,wnd syscall.Handle,dwFlags uint32,lParam uintptr)int32  {
	ret,_,_ := syscall.Syscall6(fnChangeDisplaySettingsExW,5,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpszDeviceName))),
		uintptr(unsafe.Pointer(lpDevMode)),uintptr(wnd), uintptr(dwFlags),lParam,0)
	return int32(ret)
}

func ChangeMenu(hMenu HMENU,cmd uint,lpszNewItem string,cmdInsert uint,flags uint)bool{
	ret,_,_ := syscall.Syscall6(fnChangeMenuW,5,
		uintptr(hMenu),uintptr(cmd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpszNewItem))),
		uintptr(cmdInsert),uintptr(flags),0)
	return ret!=0
}


func CheckDlgButton(hdlg syscall.Handle,nIdButton int32,uCheck uint)bool  {
	ret,_,_:=syscall.Syscall(fnCheckDlgButton,3,uintptr(hdlg),uintptr(nIdButton),uintptr(uCheck))
	return ret!=0
}

func DeleteMenu(hMenu HMENU,uPosition, uFlags uint)bool  {
	ret,_,_:=syscall.Syscall(fnDeleteMenu,3,uintptr(hMenu),uintptr(uPosition),uintptr(uFlags))
	return ret!=0
}

func DestroyCaret()bool{
	ret,_,_:=syscall.Syscall(fnDestroyCaret,0,0,0,0)
	return ret!=0
}

func DestroyCursor(cursor HCURSOR)bool{
	ret,_,_:=syscall.Syscall(fnDestroyCursor,1,uintptr(cursor),0,0)
	return ret!=0
}

func DestroyIcon(icon HICON)bool{
	ret,_,_:=syscall.Syscall(fnDestroyIcon,1,uintptr(icon),0,0)
	return ret!=0
}

func DestroyMenu(menu HMENU)bool{
	ret,_,_:=syscall.Syscall(fnDestroyMenu,1,uintptr(menu),0,0)
	return ret!=0
}

func DestroyWindow(wnd syscall.Handle)bool{
	ret,_,_:=syscall.Syscall(fnDestroyWindow,1,uintptr(wnd),0,0)
	return ret!=0
}

func DrawCaption(p1 syscall.Handle,p2 HDC,p3 *Rect,p4 uint)bool  {
	ret,_,_:=syscall.Syscall6(fnDrawCaption,4,uintptr(p1),uintptr(p2),uintptr(unsafe.Pointer(p3)),uintptr(p4),0,0)
	return ret!=0
}

func DrawEdge(hdc HDC,qrc *Rect,edge uint,grfFlags uint)bool  {
	ret,_,_:=syscall.Syscall6(fnDrawEdge,4,uintptr(hdc),uintptr(unsafe.Pointer(qrc)),uintptr(edge),uintptr(grfFlags),0,0)
	return ret!=0
}

func DrawFocusRect(hdc HDC,rpc *Rect)bool  {
	ret,_,_:=syscall.Syscall(fnDrawFocusRect,2,uintptr(hdc),uintptr(unsafe.Pointer(rpc)),0)
	return ret!=0
}


func DrawFrameControl(dc HDC,rect *Rect,utype,ustate uint) bool {
	ret,_,_:=syscall.Syscall6(fnDrawFrameControl,4,uintptr(dc),uintptr(unsafe.Pointer(rect)),uintptr(utype),uintptr(ustate),0,0)
	return ret!=0
}

//function DrawIcon(hDC: HDC; X, Y: Integer; hIcon: HICON): BOOL; stdcall;
func DrawIcon(dc HDC,x,y int,hIcon HICON) bool {
	ret,_,_:=syscall.Syscall6(fnDrawIcon,4,uintptr(dc),uintptr(x),uintptr(y),uintptr(hIcon),0,0)
	return ret!=0
}

func DrawIconEx(dc HDC,xLeft, yTop int,hIcon HICON,cxWidth, cyWidth int,istepIfAniCur uint,hbrFlickerFreeDraw HBRUSH,diFlags uint) bool {
	ret,_,_:=syscall.Syscall9(fnDrawIcon,4,uintptr(dc),uintptr(xLeft),uintptr(yTop),uintptr(hIcon),uintptr(cxWidth),
		uintptr(cyWidth),uintptr(istepIfAniCur),uintptr(hbrFlickerFreeDraw),uintptr(diFlags))
	return ret!=0
}

func DrawMenuBar(hWnd syscall.Handle)bool  {
	ret,_,_:=syscall.Syscall(fnDrawMenuBar,1,uintptr(hWnd),0,0)
	return ret!=0
}

func DrawTextEx(dc HDC,text string,textlen int,textrect *Rect,dwDTFormat uint,DTParams *GDrawTextParams)int  {
	ret,_,_:=syscall.Syscall6(fnDrawTextExW,6,uintptr(dc),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(textlen),uintptr(unsafe.Pointer(textrect)),
		uintptr(dwDTFormat),uintptr(unsafe.Pointer(DTParams)))
	return int(ret)
}

func DrawText(dc HDC,text string,textlen int,textrect *Rect,uFormat uint)int  {
	ret,_,_:=syscall.Syscall6(fnDrawTextW,6,uintptr(dc),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(textlen),uintptr(unsafe.Pointer(textrect)),
		uintptr(uFormat),0)
	return int(ret)
}

func FrameRect(hdc HDC,lprc *Rect,hbr HBRUSH)int  {
	ret,_,_:=syscall.Syscall(fnFrameRect,3,uintptr(hdc),uintptr(unsafe.Pointer(lprc)),uintptr(hbr))
	return int(ret)
}

func InvertRect(hdc HDC,rc *Rect)bool  {
	ret,_,_:=syscall.Syscall(fnFrameRect,2,uintptr(hdc),uintptr(unsafe.Pointer(rc)),0)
	return ret!=0
}

func FillRect(hdc HDC,lprc *Rect,hbr HBRUSH)int  {
	ret,_,_:=syscall.Syscall(fnFillRect,3,uintptr(hdc),uintptr(unsafe.Pointer(lprc)),uintptr(hbr))
	return int(ret)
}

func InflateRect(rc *Rect,dx,dy int)bool  {
	ret,_,_:=syscall.Syscall(fnInflateRect,3,uintptr(unsafe.Pointer(rc)),uintptr(dx),uintptr(dy))
	return ret!=0
}


func EnableWindow(hWnd syscall.Handle, bEnable bool) bool {
	ret,_,_:=syscall.Syscall(fnEnableWindow,2,uintptr(hWnd),uintptr(BoolToUint(bEnable)),0)
	return ret!=0
}

func CreateSolidBrush(color uint32)HBRUSH  {
	ret,_,_:=syscall.Syscall(fnCreateSolidBrush,1,uintptr(color),0,0)
	return HBRUSH(ret)
}

func SaveDC(dc HDC)int  {
	ret,_,_:=syscall.Syscall(fnSaveDC,1,uintptr(dc),0,0)
	return int(ret)
}

func RestoreDC(dc HDC,restIndex int)bool  {
	ret,_,_:=syscall.Syscall(fnRestoreDC,2,uintptr(dc),uintptr(restIndex),0)
	return ret!=0
}

func ExcludeClipRect(dc HDC,LeftRect, TopRect, RightRect, BottomRect int) int{
	ret,_,_:=syscall.Syscall6(fnExcludeClipRect,5,uintptr(dc),uintptr(LeftRect),
		uintptr(TopRect),uintptr(RightRect),uintptr(BottomRect),0)
	return int(ret)
}

func DeleteObject(Hgdiobj uintptr)bool  {
	ret,_,_:=syscall.Syscall(fnDeleteObject,1,Hgdiobj,0,0)
	return ret!=0
}

func IsIconic(hwnd syscall.Handle)bool  {
	ret,_,_:=syscall.Syscall(fnIsIconic,1,uintptr(hwnd),0,0)
	return ret!=0
}

func GetWindowRect(hwnd syscall.Handle,lpRect *Rect)bool  {
	ret,_,_:=syscall.Syscall(fnGetWindowRect,2,uintptr(hwnd),uintptr(unsafe.Pointer(lpRect)),0)
	return ret!=0
}

func ScreenToClient(hwnd syscall.Handle,lpoint *POINT)bool  {
	ret,_,_:=syscall.Syscall(fnScreenToClient,2,uintptr(hwnd),uintptr(unsafe.Pointer(lpoint)),0)
	return ret!=0
}

func ClientToScreen(hwnd syscall.Handle,lpoint *POINT)bool  {
	ret,_,_:=syscall.Syscall(fnClientToScreen,2,uintptr(hwnd),uintptr(unsafe.Pointer(lpoint)),0)
	return ret!=0
}

func GetSysColor(coloridx int32)uint32  {
	ret,_,_:=syscall.Syscall(fnGetSysColor,1,uintptr(coloridx),0,0)
	return uint32(ret)
}

func GetSysColorBrush(coloridx int32)HBRUSH  {
	ret,_,_:=syscall.Syscall(fnGetSysColorBrush,1,uintptr(coloridx),0,0)
	return HBRUSH(ret)
}

func SendMessage(hwnd syscall.Handle,msg uint,wparam,lparam uintptr) LRESULT {
	ret,_,_:=syscall.Syscall6(fnSendMessageW,1,uintptr(hwnd),uintptr(msg),wparam,lparam,0,0)
	return LRESULT(ret)
}


func PostMessage(hwnd syscall.Handle,msg uint,wparam,lparam uintptr) bool {
	ret,_,_:=syscall.Syscall6(fnPostMessageW,1,uintptr(hwnd),uintptr(msg),wparam,lparam,0,0)
	return ret != 0
}


func InvalidateRect(hwnd syscall.Handle,r *Rect,bErase bool)bool  {
	ret,_,_:=syscall.Syscall(fnInvalidateRect,3,uintptr(hwnd),uintptr(unsafe.Pointer(r)),uintptr(BoolToUint(bErase)))
	return ret!=0
}

func SetActiveWindow(hwnd syscall.Handle) syscall.Handle {
	ret,_,_:=syscall.Syscall(fnSetActiveWindow,1,uintptr(hwnd),0,0)
	return syscall.Handle(ret)
}

func GetActiveWindow() syscall.Handle {
	ret,_,_:=syscall.Syscall(fnGetActiveWindow,0,0,0,0)
	return syscall.Handle(ret)
}

func EnumThreadWindows(dwThreadId uint32,lpfn uintptr,lParam uintptr)bool  {
	ret,_,_:=syscall.Syscall(fnEnumThreadWindows,3,uintptr(dwThreadId),lpfn,lParam)
	return ret!=0
}

func GetCurrentThreadID() uint32  {
	ret,_,_:=syscall.Syscall(fnGetCurrentThreadId,0,0,0,0)
	return uint32(ret)
}

func GetCurrentThread() syscall.Handle  {
	ret,_,_:=syscall.Syscall(fnGetCurrentThread,0,0,0,0)
	return syscall.Handle(ret)
}

func IsWindowVisible(hwnd syscall.Handle)bool  {
	ret,_,_:=syscall.Syscall(fnIsWindowVisible,1,uintptr(hwnd),0,0)
	return ret!=0
}

func IsWindowEnabled(hwnd syscall.Handle)bool  {
	ret,_,_:=syscall.Syscall(fnIsWindowEnabled,1,uintptr(hwnd),0,0)
	return ret!=0
}


func SetWindowLong(hWnd syscall.Handle,nIndex int,dwNewLong int)int  {
	ret,_,_:=syscall.Syscall(fnSetWindowLongW,3,uintptr(hWnd),uintptr(nIndex),uintptr(dwNewLong))
	return int(ret)
}

func SetWindowLongPtr(hWnd syscall.Handle,nIndex int,dwNewLong int64)int64  {
	ret,_,_:=syscall.Syscall(fnSetWindowLongPtrW,3,uintptr(hWnd),uintptr(nIndex),uintptr(dwNewLong))
	return int64(ret)
}

func UTF16Ptr2String(ptr *uint16,ptrLen uint)string  {
	if ptr != nil {
		us := make([]uint16, 0, ptrLen)
		for p := uintptr(unsafe.Pointer(ptr)); ; p += 2 {
			u := *(*uint16)(unsafe.Pointer(p))
			if u == 0 || uint(p - uintptr(unsafe.Pointer(ptr))) >= ptrLen{
				return string(utf16.Decode(us))
			}
			us = append(us, u)
		}
	}
	return ""
}

func GetDeviceCaps(dc HDC,Index int)int  {
	ret,_,_:=syscall.Syscall(fnGetDeviceCaps,2,uintptr(dc),uintptr(Index),0)
	return int(ret)
}

func GetDC(wnd syscall.Handle)HDC  {
	ret,_,_:=syscall.Syscall(fnGetDC,1,uintptr(wnd),0,0)
	return HDC(ret)
}

func ReleaseDC(wnd syscall.Handle,DC HDC)int  {
	ret,_,_:=syscall.Syscall(fnReleaseDC,2,uintptr(wnd),uintptr(DC),0)
	return int(ret)
}


func InitScreenLogPixels()  {
	DC := GetDC(0)
	ScreenLogPixels = int32(GetDeviceCaps(DC, LOGPIXELSY))
	ReleaseDC(0,DC)
}

func GetStockObject(Index int)uintptr  {
	ret,_,_:=syscall.Syscall(fnGetStockObject,1,uintptr(Index),0,0)
	return ret
}

func GetSubMenu(hMenu HMENU,npos int)HMENU  {
	ret,_,_:=syscall.Syscall(fnGetSubMenu,2,uintptr(hMenu),uintptr(npos),0)
	return HMENU(ret)
}

func EnableMenuItem(hMenu HMENU,npos int,uEnable bool)bool  {
	var nflag uint=MF_BYPOSITION
	if uEnable{
		nflag = nflag | MF_ENABLED
	}else{
		nflag = nflag | MF_DISABLED
	}
	ret,_,_:=syscall.Syscall(fnEnableMenuItem,3,uintptr(hMenu),uintptr(npos),uintptr(nflag))
	return ret!=0
}


func GetMenuItemID(hmenu HMENU,npos int32)uint  {
	ret,_,_:=syscall.Syscall(fnGetMenuItemID,2,uintptr(hmenu),uintptr(npos),0)
	return uint(ret)
}

func GetMenuItemCount(hMenu HMENU) int {
	ret,_,_:=syscall.Syscall(fnGetMenuItemCount,2,uintptr(hMenu),0,0)
	return int(ret)
}

func GetMenuString(hMenu HMENU,npos int32)string  {
	var uflag uint=MF_BYPOSITION
	ret,_,_:=syscall.Syscall6(fnGetMenuStringW,5,uintptr(hMenu),uintptr(npos),0,0,uintptr(uflag),0)
	mp := make([]uint16,ret+1)
	ret,_,_ = syscall.Syscall6(fnGetMenuStringW,5,uintptr(hMenu),uintptr(npos),uintptr(unsafe.Pointer(&mp[0])),0,uintptr(uflag),0)
	return syscall.UTF16ToString(mp)
}

func CreateMenu()HMENU  {
	ret,_,_:=syscall.Syscall(fnCreateMenu,0,0,0,0)
	return HMENU(ret)
}

func CreatePopupMenu()HMENU  {
	ret,_,_:=syscall.Syscall(fnCreatePopupMenu,0,0,0,0)
	return HMENU(ret)
}


func GetSystemMenu(hWnd syscall.Handle,bRevert bool)HMENU  {
	ret,_,_:=syscall.Syscall(fnGetSystemMenu,2,uintptr(hWnd),uintptr(BoolToUint(bRevert)),0)
	return HMENU(ret)
}

func AppendMenuW(hMenu HMENU,lpNewItem string,MenuId uint16)bool  {
	var uflags uint=MF_STRING
	ret,_,_:=syscall.Syscall6(fnAppendMenuW,4,uintptr(hMenu),uintptr(uflags),uintptr(MenuId),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpNewItem))),0,0)
	return ret!=0
}

func TrackPopupMenu(hMenu HMENU,uflags uint,x,y,nReserved int,hwnd syscall.Handle,prRect *Rect)bool  {
	ret,_,_:=syscall.Syscall9(fnTrackPopupMenu,7,uintptr(hMenu),uintptr(uflags),uintptr(x),
		uintptr(y),uintptr(nReserved),
		uintptr(hwnd),uintptr(unsafe.Pointer(prRect)),0,0)
	return ret!=0
}

func InsertMenuItem(hmenu HMENU,nposID uint,ispos bool,iteminfo *GMenuItemInfo)bool  {
	ret,_,_:=syscall.Syscall6(fnInsertMenuItemW,4,uintptr(hmenu),uintptr(nposID),uintptr(BoolToUint(ispos)),
		uintptr(unsafe.Pointer(iteminfo)),0,0)
	return ret!=0
}

func LoadIcon(hinstance HINST,iconName uintptr) HICON {
	ret,_,_:=syscall.Syscall(fnLoadIconW,2,uintptr(hinstance),iconName,0)
	return HICON(ret)
}


func LoadBitmap(hinstance HINST,lpBitmapName string) HBITMAP {
	ret,_,_:=syscall.Syscall(fnLoadBitmapW,2,uintptr(hinstance),uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpBitmapName))),0)
	return HBITMAP(ret)
}

func GetCursorPos(pt *POINT)bool  {
	ret,_,_:=syscall.Syscall(fnGetCursorPos,1,uintptr(unsafe.Pointer(pt)),0,0)
	return ret != 0
}

func EndMenu()bool  {
	ret,_,_:=syscall.Syscall(fnEndMenu,0,0,0,0)
	return ret != 0
}

func SetForegroundWindow(hWnd syscall.Handle)bool  {
	ret,_,_:=syscall.Syscall(fnSetForegroundWindow,1,uintptr(hWnd),0,0)
	return ret != 0
}

func GetWindow(hWnd syscall.Handle,uCmd uint)syscall.Handle  {
	ret,_,_:=syscall.Syscall(fnGetWindow,2,uintptr(hWnd),uintptr(uCmd),0)
	return syscall.Handle(ret)
}