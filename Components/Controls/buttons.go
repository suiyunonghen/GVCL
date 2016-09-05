package controls

import (
	"DxSoft/GVCL/Components"
	"DxSoft/GVCL/WinApi"
	"reflect"
	"unsafe"
)

type GButton struct {
	GWinControl
	fDefault bool
	OnClick NotifyEvent
}

func (btn *GButton) SetDefault(v bool) {
	btn.fDefault = v
}

func (btn *GButton) Default() bool {
	return btn.fDefault
}

func (btn *GButton) SubInit() {
	btn.GWinControl.SubInit()
	btn.GComponent.SubInit(btn)
}

func (btn *GButton) CreateParams(params *GCreateParams) {
	btn.GWinControl.CreateParams(params)
	btn.InitSubclassParams(params, "BUTTON")
	params.WinClassName = "GButton"
	if btn.fDefault {
		params.Style = params.Style | WinApi.BS_DEFPUSHBUTTON
	} else {
		params.Style = params.Style | WinApi.BS_PUSHBUTTON
	}
}

func (btn *GButton)CreateWindowHandle(params *GCreateParams)(result bool)  {
	btn.fHandle = WinApi.CreateWindowEx(params.ExStyle, "BUTTON",
		btn.fCaption, params.Style, params.X, params.Y,
		params.Width, params.Height, params.WndParent, 0, params.WindowClass.HInstance,
		unsafe.Pointer(params.Param))
	result = btn.fHandle !=0
	if result{
		if WinApi.IsAMD64(){
			//指定窗口过程
			btn.FDefWndProc = uintptr(WinApi.SetWindowLongPtr(btn.fHandle,WinApi.GWL_WNDPROC,int64(InitWndprocCallBack)))
		}else{
			btn.FDefWndProc = uintptr(WinApi.SetWindowLong(btn.fHandle,WinApi.GWL_WNDPROC,int(InitWndprocCallBack)))
		}
	}
	return
}


func (btn *GButton) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	result = 0
	msgDispatchNext = false
	switch msg {
	case WinApi.WM_COMMAND: //按钮事件
	    	notifycode := WinApi.HiWord(uint32(wparam))
		if notifycode == uint16(WinApi.BN_CLICKED){
			if btn.OnClick != nil{
				btn.OnClick(btn)
			}
		}
	default:
		result = WinApi.CallWindowProc(btn.FDefWndProc, btn.fHandle, msg, wparam, lparam)
	}
	return
}

func NewButton(aParent Components.IWincontrol) *GButton {
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		btn := new(GButton)
		btn.SubInit()
		btn.SetParent(aParent)
		btn.SetWidth(80)
		btn.SetVisible(true)
		btn.SetHeight(30)
		return btn
	}
	return nil
}
