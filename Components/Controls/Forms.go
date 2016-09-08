package controls

import (
	"DxSoft/GVCL/Components"
	"DxSoft/GVCL/WinApi"
	"DxSoft/GVCL/Graphics"
	_"fmt"
	"syscall"
)

var (
	application *WApplication
	unActiveFormList *unActiveFormNode
)
const(
	CANone=iota
	CAHide
	CAFree
	CAMinimize
)

type OnFormCloseEvent func(sender interface{},closAction *int8)

type GForm struct {
	GWinControl
	isSetModalresult bool
	fModalResult int8
	fisModalform bool
	OnClose OnFormCloseEvent
	OnCreate NotifyEvent
}

const(
	MrNone = 0
	MrOK = WinApi.IDOK
	MrCancle = WinApi.IDCANCEL
	MrYes = WinApi.IDYES
	MrNo = WinApi.IDNO
	MrIgnore = WinApi.IDIGNORE
	MrTryAgin = WinApi.IDTRYAGAIN
	MrRetry = WinApi.IDRETRY
	MrAbort = WinApi.IDABORT
	MrClose = WinApi.IDCLOSE
	MrContinue = WinApi.IDCONTINUE
)
func (frm *GForm) SubInit() {
	frm.GWinControl.SubInit()
	frm.GComponent.SubInit(frm)
	frm.BindMessageMpas()
	frm.fIsForm = true
	frm.fColor = Graphics.ClBtnFace
	frm.fwidth = 400
	frm.fheight = 300
	if frm.OnCreate!=nil{
		frm.OnCreate(frm)
	}
}

type unActiveFormNode struct {
	formwnd syscall.Handle
	prev *unActiveFormNode
}
type WApplication struct {
	Components.GComponent
	fMainForm    *GForm
	fForms       []*GForm
	fRunning     bool
	ShowMainForm bool
	fTerminate   bool
	OnMessage    MessageEventHandler
	fChildObj    interface{}
	ActiveForm  *GForm
}

func (app *WApplication) Run() {
	application = app
	app.fRunning = true
	if app.fMainForm != nil {
		if app.ShowMainForm {
			app.fMainForm.SetVisible(true)
		}
		for {
			app.HandleMessage()
			if app.fTerminate {
				break
			}
		}
	}
	app.fRunning = false
}



func (frm *GForm) CreateParams(params *Components.GCreateParams) {
	frm.GWinControl.CreateParams(params)
	nstyle := int32(^(WinApi.WS_CHILD | WinApi.WS_GROUP | WinApi.WS_TABSTOP))
	params.Style = uint32(int32(params.Style) & nstyle)
	if application.fMainForm != nil{
		params.WndParent = application.fMainForm.fHandle
	}
	params.WinClassName = "GForm"
	params.Style = params.Style | WinApi.WS_OVERLAPPEDWINDOW | WinApi.WS_CAPTION | WinApi.WS_THICKFRAME | WinApi.WS_MINIMIZEBOX | WinApi.WS_MAXIMIZEBOX | WinApi.WS_SYSMENU
	nstyle = ^(WinApi.CS_HREDRAW | WinApi.CS_VREDRAW)
	params.WindowClass.Style = uint32(int32(params.WindowClass.Style) & nstyle)
	if frm == application.fMainForm{
		params.ExStyle = params.ExStyle | WinApi.WS_EX_APPWINDOW
	}
}

func (frm *GForm)PaintWindow(dc WinApi.HDC)int32{
	return frm.GWinControl.PaintWindow(dc)
}

func (frm *GForm)PaintBack(dc WinApi.HDC)int32  {
	//frm.GWinControl.paintBackGround(dc)
	return 1
}

func (frm *GForm)Close()  {
	if application.fMainForm == frm{
		application.fTerminate = true
		WinApi.PostQuitMessage(0)
	}else{
		WinApi.SendMessage(frm.fHandle,WinApi.WM_CLOSE,0,0)
		frm.fVisible = false
	}
}

func (frm *GForm)Show()  {
	frm.isSetModalresult = false
	frm.fisModalform = false
	frm.SetVisible(true)
	WinApi.SetWindowPos(frm.fHandle, WinApi.HWND_TOP, 0, 0, 0, 0,
		WinApi.SWP_NOMOVE + WinApi.SWP_NOSIZE);
}


func DoDisableWindow(hwnd syscall.Handle,Data uintptr) uintptr  {
	if WinApi.IsWindowVisible(hwnd) && WinApi.IsWindowEnabled(hwnd){
		oldp := unActiveFormList
		unActiveFormList = new(unActiveFormNode)
		unActiveFormList.formwnd = hwnd
		unActiveFormList.prev = oldp
		WinApi.EnableWindow(hwnd,false)
	}
	return 1
}

func EnableDisableWindow()  {
	if unActiveFormList != nil{
		for{
			WinApi.EnableWindow(unActiveFormList.formwnd,true)
			unActiveFormList = unActiveFormList.prev
			if unActiveFormList == nil{
				break
			}
		}
	}
}

func (frm *GForm)ShowModal() int8 {
	//首先应该将所有后置的窗体都EnableFalse
	OldActiveForm := application.ActiveForm
	var oldActiveWnd syscall.Handle
	if OldActiveForm == nil{
		oldActiveWnd = WinApi.GetActiveWindow()
	}else{
		oldActiveWnd = OldActiveForm.fHandle
	}
	frm.fModalResult = MrNone
	enumproc := syscall.NewCallback(DoDisableWindow)
	WinApi.EnumThreadWindows(WinApi.GetCurrentThreadID(),enumproc,0)
	frm.Show()
	frm.fisModalform = true
	for{
		application.HandleMessage()
		if frm.fModalResult != MrNone{
			break
		}
	}
	EnableDisableWindow()
	WinApi.SetActiveWindow(oldActiveWnd)
	if OldActiveForm != nil{
		application.ActiveForm = OldActiveForm
	}
	return frm.fModalResult
}

func (frm *GForm)SetModalResult(v int8)  {
	frm.fModalResult = v
	frm.isSetModalresult = true
	frm.Close()
}

func (frm *GForm)ModalResult()int8  {
	return frm.fModalResult
}

func (frm *GForm) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	result = 0
	msgDispatchNext = true
	switch msg {
	case WinApi.WM_CLOSE:
		msgDispatchNext = false
		var closeAction int8 = CAHide
		if frm.OnClose != nil{
			frm.OnClose(frm,&closeAction)
		}
		switch closeAction {
		case CANone:
			if frm.fisModalform {
				frm.SetVisible(false)
			}
			return
		case CAHide:
			frm.SetVisible(false)
		case CAFree:
			msgDispatchNext = true
		case CAMinimize:
			if frm.fisModalform {
				frm.SetVisible(false)
				return
			}
			WinApi.ShowWindow(frm.GetWindowHandle(),WinApi.SW_SHOWMINIMIZED)
		}
		if frm.isSetModalresult{

		}else{
			frm.fModalResult = MrClose
		}
	case WinApi.WM_SETFOCUS:
		application.ActiveForm = frm
	default:
		result,msgDispatchNext = frm.GWinControl.WndProc(msg,wparam,lparam)
	}
	return
}

func NewForm() *GForm {
	frm := new(GForm)
	frm.SubInit()
	return frm
}

func (app *WApplication) CreateForm() *GForm {
	frm := NewForm()
	if app.fMainForm == nil {
		app.fMainForm = frm
	}
	app.fForms = append(app.fForms, frm)
	return frm
}

//消息处理
func (app *WApplication) HandleMessage() {
	msg := new(WinApi.MSG)
	if !app.ProcessMessage(msg) {
		app.idleMsg(msg)
	}
}

func (app *WApplication) ProcessMessage(msg *WinApi.MSG) bool {
	result := false
	defer func() {
		if r := recover(); r != nil {
			//处理异常
			println("异常")
		}
	}()
	if msg.PeekMessage(0, 0, 0, WinApi.PM_REMOVE) {
		result = true
		if msg.Message != WinApi.WM_QUIT {
			handled := false
			if app.OnMessage != nil {
				app.OnMessage(app, msg, &handled)
			}
			if !handled {
				msg.TranslateMessage()
				msg.DispatchMessage()
			} else {
				app.idleMsg(msg)
			}
		} else {
			//执行完成
			app.doneApp()
		}
	}
	return result
}

func (app *WApplication) idleMsg(msg *WinApi.MSG) {
	WinApi.WaitMessage()
}

func (app *WApplication) doneApp() {
	WinApi.GlobalDeleteAtom(windowAtom)
	WinApi.GlobalDeleteAtom(controlAtom)
}
