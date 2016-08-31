package controls

import (
	"DxSoft/GVCL/Components"
	"DxSoft/GVCL/WinApi"
	"DxSoft/GVCL/Graphics"
)

var (
	application *WApplication
)

type GForm struct {
	GWinControl
}

func (frm *GForm) SubInit() {
	frm.GWinControl.SubInit()
	frm.GComponent.SubInit(frm)
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


func formCreateWndOkHandler(ctrl interface{}) {
	WinApi.UpdateWindow(ctrl.(*GWinControl).fHandle)
	WinApi.ShowWindow(ctrl.(*GWinControl).fHandle, WinApi.SW_SHOWNORMAL)
}

func (frm *GForm) CreateParams(params *GCreateParams) {
	frm.GWinControl.CreateParams(params)
	nstyle := ^(WinApi.WS_CHILD | WinApi.WS_GROUP | WinApi.WS_TABSTOP)
	params.Style = uint32(int(params.Style) & nstyle)
	params.WinClassName = "GForm"
	params.Style = params.Style | WinApi.WS_OVERLAPPEDWINDOW | WinApi.WS_CAPTION | WinApi.WS_THICKFRAME | WinApi.WS_MINIMIZEBOX | WinApi.WS_MAXIMIZEBOX | WinApi.WS_SYSMENU
	nstyle = ^(WinApi.CS_HREDRAW | WinApi.CS_VREDRAW)
	params.WindowClass.Style = uint32(int(params.WindowClass.Style) & nstyle)
	params.ExStyle = params.ExStyle | WinApi.WS_EX_APPWINDOW
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

func NewForm() *GForm {
	frm := new(GForm)
	frm.SubInit()
	frm.fIsForm = true
	frm.fColor = Graphics.ClBtnFace
	frm.fwidth = 400
	frm.fheight = 300
	frm.SetCreateWndOkHandler(formCreateWndOkHandler)
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
