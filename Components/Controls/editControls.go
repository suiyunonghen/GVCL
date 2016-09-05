package controls
import (
	"DxSoft/GVCL/Components"
	//"DxSoft/GVCL/WinApi"
	"reflect"
	_"fmt"
	"DxSoft/GVCL/WinApi"
	"unsafe"
)

type GEdit struct {
	GWinControl
	fDefault bool
	OnChange NotifyEvent
}

func (edt *GEdit) SubInit() {
	edt.GWinControl.SubInit()
	edt.GComponent.SubInit(edt)
}

func (edt *GEdit) CreateParams(params *GCreateParams) {
	edt.GWinControl.CreateParams(params)
	edt.InitSubclassParams(params, "EDIT")
	params.WinClassName = "GEdit"
	params.Style = params.Style | WinApi.ES_AUTOHSCROLL | WinApi.ES_AUTOVSCROLL | WinApi.ES_LEFT
	params.ExStyle = params.ExStyle | WinApi.WS_EX_CLIENTEDGE
}

func (edt *GEdit)CreateWindowHandle(params *GCreateParams)(result bool)  {
	edt.fHandle = WinApi.CreateWindowEx(params.ExStyle, "EDIT",
		edt.fCaption, params.Style, params.X, params.Y,
		params.Width, params.Height, params.WndParent, 0, params.WindowClass.HInstance,
		unsafe.Pointer(params.Param))
	result = edt.fHandle !=0
	if result{
		if WinApi.IsAMD64(){
			//指定窗口过程
			WinApi.SetWindowLongPtr(edt.fHandle,WinApi.GWL_WNDPROC,int64(InitWndprocCallBack))
		}else{
			WinApi.SetWindowLong(edt.fHandle,WinApi.GWL_WNDPROC,int(InitWndprocCallBack))
		}
	}
	return
}

//func (edt *GEdit)WMPaint(wparam,lparam uintptr)(result uintptr, msgDispatchNext bool)  {
//	result = WinApi.CallWindowProc(edt.FDefWndProc, edt.fHandle, WinApi.WM_PAINT, wparam, lparam)
//	msgDispatchNext = false
//	return
//}

func (edt *GEdit) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	msgDispatchNext = false
	if msg == WinApi.WM_COMMAND{
		notifycode := WinApi.HiWord(uint32(wparam))
		if notifycode == WinApi.EN_CHANGE{
			if edt.OnChange != nil{
				edt.OnChange(edt)
			}
		}
	}
	result = WinApi.CallWindowProc(edt.FDefWndProc, edt.fHandle, msg, wparam, lparam)
	return
}

func NewEdit(aParent Components.IWincontrol) *GEdit {
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		edt := new(GEdit)
		edt.SubInit()
		edt.SetParent(aParent)
		edt.SetWidth(80)
		edt.SetVisible(true)
		edt.SetHeight(30)
		return edt
	}
	return nil
}