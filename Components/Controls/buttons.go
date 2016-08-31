package controls

import (
	"DxSoft/GVCL/Components"
	"DxSoft/GVCL/WinApi"
	"reflect"
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

func (btn *GButton) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	result = 0
	msgDispatchNext = true
	switch msg {
	case WinApi.WM_COMMAND: //按钮事件
	    	notifycode := WinApi.HiWord(uint32(wparam))
		if notifycode == uint16(WinApi.BN_CLICKED){
			if btn.OnClick != nil{
				btn.OnClick(btn)
			}
		}
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
