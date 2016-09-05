package controls

import (
	"DxSoft/GVCL/Components"
	"DxSoft/GVCL/Graphics"
	"DxSoft/GVCL/WinApi"
	"reflect"
	"syscall"
	"unsafe"
	"fmt"
)

var (
	InitWndprocCallBack = syscall.NewCallback(initWndProc)
	controlAtom         WinApi.ATOM
	windowAtom          WinApi.ATOM
	Hinstance           WinApi.HINST
	MessageHandlerMap map[uint32]string
)

func init() {
	WindowAtomString := fmt.Sprintf("Go%08X", WinApi.GetCurrentProcessId())
	windowAtom = WinApi.GlobalAddAtom(WindowAtomString)
	ControlAtomString := fmt.Sprintf("Go%08X%08X", WinApi.GetModuleHandle(""), WinApi.GetCurrentProcessId())
	controlAtom = WinApi.GlobalAddAtom(ControlAtomString)
	MessageHandlerMap = make(map[uint32]string)
	MessageHandlerMap[WinApi.WM_PAINT] = "WMPaint"
}

func initWndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (result uintptr) {
	control := (*GWinControl)(unsafe.Pointer(WinApi.GetProp(hwnd, uintptr(controlAtom))))
	if control == nil {
		return WinApi.DefWindowProc(hwnd, msg, wparam, lparam)
	}
	if control.fIsForm{
		switch msg {
		case WinApi.WM_CLOSE:
			if hwnd == application.fMainForm.fHandle {
				application.fTerminate = true
				WinApi.PostQuitMessage(0)
			}
		case WinApi.WM_DESTROY:
			control.DestoryWnd()
		case WinApi.WM_SIZE:
		case WinApi.WM_COMMAND:
			if lparam==0{ //点击的是菜单或者快捷加速
				ctrlId := WinApi.LoWord(uint32(wparam))
				if WinApi.HiWord(uint32(wparam))==0{ //菜单ID
					fmt.Println("菜单ID",ctrlId)
				}else{//快捷键ID
					fmt.Println("快捷键ID",ctrlId)
				}
			}else{ //重新换算控件
				control = (*GWinControl)(unsafe.Pointer(WinApi.GetProp(syscall.Handle(lparam), uintptr(controlAtom))))
				if control == nil{
					return
				}
			}
		}
	}
	if control.fMessageHandlerMap != nil {
		if vhandler,OK := control.fMessageHandlerMap[msg];OK{
			in := make([]reflect.Value, 3)
			in[0] = reflect.ValueOf(vhandler.TargetObject)
			in[1] = reflect.ValueOf(wparam)
			in[2] = reflect.ValueOf(lparam)
			calrsult := vhandler.DispatchHandler.Call(in)
			if !calrsult[1].Bool() { //不往下传递
				return uintptr(calrsult[0].Uint())
			}
		}
	}
	if control.fRealWndprocObj == nil{
		for i := control.SubChildCount() - 1; i >= 0; i-- {
			subObj := control.SubChild(i)
			pType := reflect.TypeOf(subObj)
			if mnd, ok := pType.MethodByName("WndProc"); ok {
				pType = mnd.Func.Type()
				if pType.NumIn() == 4 && pType.NumOut() == 2 {
					control.fRealWndprocObj = new(GMsgDispatchObj)
					control.fRealWndprocObj.DispatchHandler = mnd.Func
					control.fRealWndprocObj.TargetObject = subObj
					break
				}
			}
		}
	}
	if control.fRealWndprocObj != nil{
		in := make([]reflect.Value, 4)
		in[0] = reflect.ValueOf(control.fRealWndprocObj.TargetObject)
		in[1] = reflect.ValueOf(msg)
		in[2] = reflect.ValueOf(wparam)
		in[3] = reflect.ValueOf(lparam)
		calrsult := control.fRealWndprocObj.DispatchHandler.Call(in)
		if !calrsult[1].Bool() { //不往下传递
			return uintptr(calrsult[0].Uint())
		}
	}
	if control.FDefWndProc == 0 {
		result = WinApi.DefWindowProc(hwnd, msg, wparam, lparam)
	}else{
		result = WinApi.CallWindowProc(control.FDefWndProc, hwnd, msg, wparam, lparam)
	}
	return
}

type MessageEventHandler func(sender interface{}, msg *WinApi.MSG, handled *bool)
type MessageDispatch func(sender interface{}, msg uint32, wparam, lparam uintptr) uintptr
type NotifyEvent func(sender interface{})

type GMsgDispatchObj struct {
	TargetObject interface{}
	DispatchHandler reflect.Value
}

type GBaseControl struct {
	Components.GComponent
	fColor             Graphics.GColorValue //Graphics.GColor
	fParent            Components.IWincontrol
	fleft              int32
	ftop               int32
	fwidth             int32
	fheight            int32
	fVisible           bool
	fMessageHandlerMap map[uint32]*GMsgDispatchObj
	OnResize NotifyEvent
}

func (ctrl *GBaseControl)BindMessageMpas()  {
	if ctrl.fMessageHandlerMap == nil{
		ctrl.fMessageHandlerMap = make(map[uint32]*GMsgDispatchObj)
	}
	hasInControl := false
	for i:=0;i<ctrl.SubChildCount()-1;i++{
		subObj := ctrl.SubChild(i)
		pType := reflect.TypeOf(subObj)
		stype := pType.String()
		if !hasInControl {
			if stype != "*controls.GBaseControl" || stype != "*GBaseControl"{
				hasInControl=true
			}else{
				continue
			}
		}
		for k,v := range MessageHandlerMap{
			if mnd, ok := pType.MethodByName(v); ok{
				pType = mnd.Func.Type()
				if pType.NumIn()==3 && pType.NumOut() == 2{
					msgobj := new(GMsgDispatchObj)
					msgobj.TargetObject = subObj
					msgobj.DispatchHandler = mnd.Func
					ctrl.fMessageHandlerMap[k] = msgobj
				}
			}
		}
	}
}



func (ctrl *GBaseControl) Left() int32 {
	return ctrl.fleft
}

func (ctrl *GBaseControl) SubInit() {
	ctrl.GComponent.SubInit(ctrl)
}

func (ctrl *GBaseControl) Top() int32 {
	return ctrl.ftop
}

func (ctrl *GBaseControl) Height() int32 {
	return ctrl.fheight
}

func (ctrl *GBaseControl) Width() int32 {
	return ctrl.fwidth
}

func (ctrl *GBaseControl) SetWidth(w int32) {
	if ctrl.fwidth != w{
		ctrl.fwidth = w
		ctrl.changeBands()
	}
}

func (ctrl *GBaseControl)changeBands()  {
	for i:=ctrl.SubChildCount()-1;i>=0;i--{
		subObj := ctrl.SubChild(i)
		pType := reflect.TypeOf(subObj)
		stype := pType.String()
		if stype == "*controls.GBaseControl" || stype == "*GBaseControl" {
			break
		}
		if mnd, ok := pType.MethodByName("SetBounds"); ok {
			pType = mnd.Func.Type()
			if pType.NumIn() == 5 && pType.NumOut() == 0 {
				in := make([]reflect.Value, 5)
				in[0] = reflect.ValueOf(subObj)
				in[1] = reflect.ValueOf(ctrl.fleft)
				in[2] = reflect.ValueOf(ctrl.ftop)
				in[3] = reflect.ValueOf(ctrl.fwidth)
				in[4] = reflect.ValueOf(ctrl.fheight)
				mnd.Func.Call(in)
			}
		}
	}
}

func (ctrl *GBaseControl) SetLeft(l int32) {
	if ctrl.fleft != l{
		ctrl.fleft = l
		ctrl.changeBands()
	}
}
func (ctrl *GBaseControl) SetTop(t int32) {
	if ctrl.ftop != t{
		ctrl.ftop = t
		ctrl.changeBands()
	}
}
func (ctrl *GBaseControl) SetHeight(h int32) {
	if ctrl.fheight != h{
		ctrl.fheight = h
		ctrl.changeBands()
	}
}

func (ctrl *GBaseControl) GetColor() Graphics.GColorValue {
	return ctrl.fColor
}

func (ctrl *GBaseControl) SetColor(c Graphics.GColorValue) {
	ctrl.fColor = c
}

func (ctrl *GBaseControl)BoundsRect()*WinApi.Rect  {
	result := new(WinApi.Rect)
	result.Left = ctrl.fleft
	result.Top = ctrl.ftop
	result.Right = ctrl.fleft + ctrl.fwidth
	result.Bottom = ctrl.ftop + ctrl.fheight
	return result
}

func (ctrl *GBaseControl) Invalidate() {
	if ctrl.fParent!=nil {
		handle := ctrl.fParent.GetWindowHandle()
		if handle!=0{
			r := ctrl.BoundsRect()
			WinApi.InvalidateRect(handle,r,false)
		}
	}
}

func (ctrl *GBaseControl)Visible()bool  {
	return ctrl.fVisible
}

func (ctrl *GBaseControl)SetBounds(ALeft, ATop, AWidth, AHeight int32)  {
	ctrl.fleft = ALeft
	ctrl.ftop= ATop
	ctrl.fwidth = AWidth
	ctrl.fheight = AHeight
	ctrl.Invalidate()
}


func (ctrl *GBaseControl) GetParent() Components.IWincontrol {
	return ctrl.fParent
}

func (ctrl *GBaseControl) SetParent(AParent Components.IWincontrol) {
	if ctrl.fParent != AParent {
		if ctrl.fParent != nil {
			ctrl.fParent.RemoveControl(ctrl)
		}
		ctrl.fParent = AParent
		ctrl.fParent.InsertControl(ctrl)
	}
}


func (ctrl *GBaseControl) SetVisible(v bool) {
	ctrl.fVisible = v
}

type GCreateParams struct {
	Caption      string
	Style        uint32
	ExStyle      uint32
	X            int32
	Y            int32
	Width        int32
	Height       int32
	WndParent    syscall.Handle
	Param        uintptr
	WindowClass  WinApi.GWndClass
	WinClassName string
}

type GWinControl struct {
	GBaseControl
	fControls           []Components.IControl
	fWincontrols        []Components.IWincontrol
	fHandle             syscall.Handle
	fIsForm             bool
	fCaption            string
	FDefWndProc         uintptr
	fRealWndprocObj     *GMsgDispatchObj
}

func (ctrl *GWinControl) SubInit() {
	ctrl.GBaseControl.SubInit()
	ctrl.GComponent.SubInit(ctrl)
}

func (ctrl *GWinControl)Perform(msg uint32,wparam, lparam uintptr)(result WinApi.LRESULT)  {
	if ctrl.fHandle != 0{
		return WinApi.LRESULT(initWndProc(ctrl.fHandle,msg,wparam,lparam))
	}
	if ctrl.fRealWndprocObj == nil{
		for i := ctrl.SubChildCount() - 1; i >= 0; i-- {
			subObj := ctrl.SubChild(i)
			pType := reflect.TypeOf(subObj)
			if mnd, ok := pType.MethodByName("WndProc"); ok {
				pType = mnd.Func.Type()
				if pType.NumIn() == 4 && pType.NumOut() == 2 {
					ctrl.fRealWndprocObj = new(GMsgDispatchObj)
					ctrl.fRealWndprocObj.DispatchHandler = mnd.Func
					ctrl.fRealWndprocObj.TargetObject = subObj
					break
				}
			}
		}
	}
	if ctrl.fRealWndprocObj != nil{
		in := make([]reflect.Value, 4)
		in[0] = reflect.ValueOf(ctrl.fRealWndprocObj.TargetObject)
		in[1] = reflect.ValueOf(msg)
		in[2] = reflect.ValueOf(wparam)
		in[3] = reflect.ValueOf(lparam)
		calrsult := ctrl.fRealWndprocObj.DispatchHandler.Call(in)
		if !calrsult[1].Bool() { //不往下传递
			return WinApi.LRESULT(calrsult[0].Uint())
		}
	}
	return 0
}


func (ctrl *GWinControl) SetParent(AParent Components.IWincontrol) {
	if ctrl.fParent != AParent {
		if ctrl.fParent != nil {
			ctrl.fParent.RemoveChildWinCtrl(ctrl)
		}
		if AParent != nil {
			hasWincontrol := false
			pType := reflect.TypeOf(AParent)
			if pType.Kind() == reflect.Ptr {
				_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
			}
			if !hasWincontrol {
				return
			}
		}
		ctrl.fParent = AParent
		if AParent != nil {
			ctrl.fParent.InsertChildWinCtrl(ctrl)
		}
	}
}

func (ctrl *GWinControl)SetBounds(ALeft, ATop, AWidth, AHeight int32)  {
	ctrl.fleft = ALeft
	ctrl.ftop= ATop
	ctrl.fwidth = AWidth
	ctrl.fheight = AHeight
	if ctrl.fHandle != 0{
		WinApi.SetWindowPos(ctrl.fHandle, 0, ALeft, ATop, AWidth, AHeight,
			WinApi.SWP_NOZORDER + WinApi.SWP_NOACTIVATE)
	}
}

func (ctrl *GWinControl) SetCaption(v string) {
	ctrl.fCaption = v
	if ctrl.fHandle != 0 {
		WinApi.SetWindowText(ctrl.fHandle, v)
	}
}

func (ctrl *GWinControl) GetText() string {
	if ctrl.fHandle == 0 {
		return ctrl.fCaption
	} else {
		return WinApi.GetWindowText(ctrl.fHandle)
	}
}

func (ctrl *GWinControl) GetWindowHandle() syscall.Handle {
	return ctrl.fHandle
}

func (ctrl *GWinControl) Invalidate() {
	if ctrl.fHandle != 0{
		WinApi.InvalidateRect(ctrl.fHandle,nil,true)
		//刷新GraphicControl控件
		for _,v := range ctrl.fControls{
			v.Invalidate()
		}
	}
}

func (ctrl *GWinControl) RemoveChildWinCtrl(Actrl Components.IWincontrol) {
	if ctrl.fWincontrols != nil {
		for k, v := range ctrl.fWincontrols {
			if v == Actrl {
				kk := k + 1
				ctrl.fWincontrols = append(ctrl.fWincontrols[:k], ctrl.fWincontrols[kk:]...)
			}
		}
	}
}

func (ctrl *GWinControl) RemoveControl(Actrl Components.IControl) {
	if ctrl.fControls != nil {
		for k, v := range ctrl.fControls {
			if v == Actrl {
				kk := k + 1
				ctrl.fControls = append(ctrl.fControls[:k], ctrl.fControls[kk:]...)
			}
		}
	}
}

func (ctrl *GWinControl) ControlExists(Actrl Components.IControl) bool {
	pType := reflect.TypeOf(Actrl)
	if pType.Kind() == reflect.Ptr {
		if _, ok := pType.Elem().FieldByName("GWinControl"); ok {
			value := reflect.ValueOf(Actrl).Elem().FieldByName("GWinControl")
			if value.IsValid() && value.CanInterface() {
				wincontrol := value.Interface().(*GWinControl)
				if ctrl.fWincontrols != nil {
					for _, wctrl := range ctrl.fWincontrols {
						if wctrl.GetWindowHandle() == wincontrol.GetWindowHandle() {
							return true
						}
					}
				}
			}
		}
	}
	if ctrl.fControls != nil {
		for _, v := range ctrl.fControls {
			if v == Actrl {
				return true
			}
		}
	}
	return false
}

func (ctrl *GWinControl) WindowExists(Actrlhandle syscall.Handle) bool {
	if ctrl.fWincontrols != nil {
		for _, v := range ctrl.fWincontrols {
			if v.GetWindowHandle() == Actrlhandle {
				return true
			}
		}
	}
	return false
}

func ExecuteMethod(obj interface{}, MethodName string, params []interface{})(exeOk bool) {
	pType := reflect.TypeOf(obj)
	exeOk = false
	if mnd, ok := pType.MethodByName(MethodName); ok {
		pType = mnd.Func.Type()
		paramLen := len(params) + 1
		if pType.NumIn() == paramLen && pType.NumOut() == 0 {
			in := make([]reflect.Value, paramLen)
			in[0] = reflect.ValueOf(obj)
			for i := 0; i < paramLen-1; i++ {
				in[i+1] = reflect.ValueOf(params[i])
			}
			mnd.Func.Call(in)
			exeOk = true
		}
	}
	return
}

func (ctrl *GWinControl) CreateParams(params *GCreateParams) {
	params.Style = WinApi.WS_CHILD | WinApi.WS_CLIPSIBLINGS | WinApi.WS_TABSTOP
	params.ExStyle = WinApi.WS_EX_CONTROLPARENT
	params.X = ctrl.fleft
	params.Y = ctrl.ftop
	params.Width = ctrl.fwidth
	params.Height = ctrl.fheight
	if ctrl.fParent != nil {
		params.WndParent = ctrl.fParent.GetWindowHandle()
	}
	params.WindowClass.Style = WinApi.CS_VREDRAW + WinApi.CS_HREDRAW + WinApi.CS_DBLCLKS
	if Hinstance == 0 {
		Hinstance = WinApi.HINST(WinApi.GetModuleHandle(""))
	}
	params.WindowClass.HCursor = WinApi.LoadCursor(0, (*uint16)(unsafe.Pointer((uintptr)(WinApi.IDC_ARROW))))
	params.WindowClass.HInstance = Hinstance
}

func (ctrl *GWinControl) CreateWnd() {
	params := new(GCreateParams)
	stype := ""
	for i := ctrl.SubChildCount() - 1; i >= 0; i-- {
		subObj := ctrl.SubChild(i)
		pType := reflect.TypeOf(subObj)
		stype = pType.String()
		if stype == "*controls.GWinControl" || stype == "*GWinControl" {
			break
		}
		//模拟执行继承功能的操作
		if mnd, ok := pType.MethodByName("CreateParams"); ok {
			pType = mnd.Func.Type()
			if pType.NumIn() == 2 && pType.NumOut() == 0 {
				in := make([]reflect.Value, 2)
				in[0] = reflect.ValueOf(subObj)
				in[1] = reflect.ValueOf(params)
				mnd.Func.Call(in)
				break
			}
		}
	}
	ctrl.FDefWndProc = params.WindowClass.FnWndProc

	if exeOk,retv := ctrl.ExecuteChildMethod("CreateWindowHandle",params);exeOk && retv.Bool(){
		WinApi.SetProp(ctrl.fHandle, uintptr(controlAtom), uintptr(unsafe.Pointer(ctrl)))
		WinApi.SetProp(ctrl.fHandle, uintptr(windowAtom), uintptr(unsafe.Pointer(ctrl)))
		if ctrl.fParent!=nil{
			WinApi.SetWindowPos(ctrl.fHandle, WinApi.HWND_TOP, 0, 0, 0, 0,
				WinApi.SWP_NOMOVE + WinApi.SWP_NOSIZE + WinApi.SWP_NOACTIVATE)
		}
		//ctrl.Perform(WinApi.WM_SETFONT, ctrl.fontHandle, 1);
	}else{
		panic("Create Handle Failed")
	}
}

func (ctrl *GWinControl) DestoryWnd() {
	WinApi.RemoveProp(ctrl.fHandle, uintptr(controlAtom))
	WinApi.RemoveProp(ctrl.fHandle, uintptr(windowAtom))
	ctrl.fHandle = 0
	ctrl.fRealWndprocObj = nil
}

func (ctrl *GWinControl) Focused() bool {
	if ctrl.fHandle != 0 {
		fhandle := WinApi.GetFocus()
		return fhandle == ctrl.fHandle
	}
	return false
}

func (ctrl *GWinControl) ControlCount() int {
	if ctrl.fControls == nil {
		return 0
	}
	return len(ctrl.fControls)
}

func (ctrl *GWinControl) Controls(idx int) Components.IControl {
	if idx < 0 || idx > len(ctrl.fControls)-1 {
		return nil
	}
	return ctrl.fControls[idx]
}

func (ctrl *GWinControl) WControlCount() int {
	if ctrl.fWincontrols == nil {
		return 0
	}
	return len(ctrl.fWincontrols)
}

func (ctrl *GWinControl) InsertChildWinCtrl(wctrl Components.IWincontrol) {
	if ctrl.fWincontrols == nil {
		ctrl.fWincontrols = make([]Components.IWincontrol, 0)
	}
	ctrl.fWincontrols = append(ctrl.fWincontrols, wctrl)
}

func (ctrl *GWinControl) InsertControl(cctrl Components.IControl) {
	if ctrl.fControls == nil {
		ctrl.fControls = make([]Components.IControl, 0)
	}
	ctrl.fControls = append(ctrl.fControls, cctrl)
}

func (ctrl *GWinControl) WControls(idx int) Components.IWincontrol {
	if idx < 0 || idx > len(ctrl.fWincontrols)-1 {
		return nil
	}
	return ctrl.fWincontrols[idx]
}

func (ctrl *GWinControl) SetVisible(v bool) {
	if ctrl.fVisible != v {
		ctrl.fVisible = v
		if ctrl.fParent != nil {
			if ctrl.fHandle != 0 {
				var showflags uint32
				if v {
					showflags = WinApi.ShowFlagsVisible
				} else {
					showflags = WinApi.ShowFlagsHide
				}
				WinApi.SetWindowPos(ctrl.fHandle, 0, 0, 0, 0, 0, showflags)
			} else {
				tmparent := ctrl.GetParent()
				for {
					if tmparent.GetParent() == nil {
						break
					}
					tmparent = tmparent.GetParent()
				}
				tmparent.UpdateShowing()
			}
		} else if ctrl.fIsForm {
			ctrl.UpdateShowing()
		}
	}
}

func (ctrl *GWinControl)CreateWindowHandle(params *GCreateParams)bool  {
	tmpWndClass := new(WinApi.GWndClass)
	classRegisted := WinApi.GetClassInfo(Hinstance, params.WinClassName, tmpWndClass)
	if !classRegisted || tmpWndClass.FnWndProc != InitWndprocCallBack {
		if classRegisted {
			WinApi.UnregisterClass(params.WinClassName, params.WindowClass.HInstance)
		}
		params.WindowClass.LpszClassName = syscall.StringToUTF16Ptr(params.WinClassName)
		params.WindowClass.FnWndProc = InitWndprocCallBack
		if atom := WinApi.RegisterClass(&params.WindowClass); atom != 0 {
			classRegisted = true
		} else {
			classRegisted = false
		}
	}
	if classRegisted {
		ctrl.fHandle = WinApi.CreateWindowEx(params.ExStyle, params.WinClassName,
			ctrl.fCaption, params.Style, params.X, params.Y,
			params.Width, params.Height, params.WndParent, 0, params.WindowClass.HInstance,
			unsafe.Pointer(params.Param))
	}
	return ctrl.fHandle !=0
}

func (ctrl *GWinControl) UpdateShowing() {
	if !ctrl.fVisible {
		if ctrl.fIsForm{
			WinApi.ShowWindow(ctrl.fHandle,WinApi.SW_HIDE)
		}
		return
	}
	hasCreatewnd := false
	if ctrl.fHandle == 0 {
		ctrl.CreateWnd()
		hasCreatewnd = true
	}
	if ctrl.fWincontrols != nil {
		for i := 0; i < len(ctrl.fWincontrols); i++ {
			ctrl.fWincontrols[i].UpdateShowing()
		}
	}
	if hasCreatewnd {
		if ctrl.fIsForm {
			if ctrl.fHandle == application.fMainForm.fHandle{
				WinApi.UpdateWindow(ctrl.fHandle)
			}
			WinApi.ShowWindow(ctrl.fHandle, WinApi.SW_SHOWNORMAL)
		}else if !ctrl.fIsForm && ctrl.fParent != nil {
			if ctrl.fVisible {
				WinApi.SetWindowPos(ctrl.fHandle, 0, 0, 0, 0, 0,  WinApi.ShowFlagsVisible)
			} else {
				WinApi.SetWindowPos(ctrl.fHandle, 0, 0, 0, 0, 0,  WinApi.ShowFlagsHide)
			}
		}
	}
}

func (ctrl *GWinControl)paintBackGround(dc WinApi.HDC)uintptr  {
	if exeok,ret := ctrl.ExecuteChildMethod("PaintBack",dc);exeok{
		return uintptr(ret.Int())
	}
	Brush := WinApi.CreateSolidBrush(uint32(ctrl.fColor))
	r := new(WinApi.Rect)
	r.Right = ctrl.Width()
	r.Bottom =ctrl.Height()
	WinApi.FillRect(dc,r,Brush)
	WinApi.DeleteObject(uintptr(Brush))
	return  1
}

func (ctrl *GWinControl)WMPaint(wparam,lparam uintptr)(result uintptr, msgDispatchNext bool)  {
	result = ctrl.PaintHandler(0)
	msgDispatchNext = result == 0
	return
}

func (ctrl *GWinControl) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	result = 0
	msgDispatchNext = true
	switch msg {
	//case WinApi.WM_PAINT:
	//	//pstruct := new(WinApi.GPaintStruct)
	//	//dc := pstruct.BeginPaint(ctrl.fHandle)
	//	//defer pstruct.EndPaint(ctrl.fHandle)
	//	result = ctrl.PaintHandler(0)
	//	msgDispatchNext = result == 0
	case WinApi.WM_ERASEBKGND:
		//背景绘制
	        result = ctrl.paintBackGround(WinApi.HDC(wparam))
		msgDispatchNext = result == 0
	case WinApi.WM_SIZE:
		rect := new(WinApi.Rect)
		if WinApi.IsIconic(ctrl.fHandle){

		}else{
			WinApi.GetWindowRect(ctrl.fHandle,rect)
		}
		LPoint := new(WinApi.POINT)
		if WinApi.GetWindowLong(ctrl.fHandle,WinApi.GWL_STYLE) & WinApi.WS_CHILD !=0 {
			//换算真实的位置
			ParentHandle := syscall.Handle(WinApi.GetWindowLong(ctrl.fHandle, WinApi.GWL_HWNDPARENT))
			if ParentHandle != 0{
				if WinApi.GetWindowLong(ParentHandle, WinApi.GWL_EXSTYLE) & WinApi.WS_EX_LAYOUTRTL != 0 {
					rect.Left,rect.Right = rect.Right,rect.Left
				}
				LPoint.X = rect.Left
				LPoint.Y = rect.Top
				WinApi.ScreenToClient(ctrl.fHandle,LPoint)
				rect.Left = LPoint.X
				rect.Top = LPoint.Y
				LPoint.X,LPoint.Y = rect.Right,rect.Bottom
				WinApi.ScreenToClient(ctrl.fHandle,LPoint)
				rect.Bottom,rect.Right = LPoint.Y,LPoint.X
			}
		}
		ctrl.fleft = rect.Left
		ctrl.ftop = rect.Top
		ctrl.fwidth = rect.Width()
		ctrl.fheight = rect.Height()
		msgDispatchNext = false
	}
	return
}

func (ctrl *GWinControl)ExecuteChildMethod(MethodName string, params ...interface{})(exeOk bool,execResult reflect.Value)  {
	exeOk = false
	for i:=ctrl.SubChildCount()-1;i>=0;i--{
		subObj := ctrl.SubChild(i)
		pType := reflect.TypeOf(subObj)
		stype := pType.String()
		if stype == "*controls.GWinControl" || stype == "*GWinControl" {
			break
		}
		if mnd, ok := pType.MethodByName(MethodName); ok {
			pType = mnd.Func.Type()
			paramLen := len(params) + 1
			if pType.NumIn() == paramLen && pType.NumOut() == 1 {
				in := make([]reflect.Value, paramLen)
				in[0] = reflect.ValueOf(subObj)
				for argNum, arg := range params{
					in[argNum+1] = reflect.ValueOf(arg)
				}
				execResult = mnd.Func.Call(in)[0]
				exeOk = true
			}
		}
	}
	return
}



//绘制控件的函数
func (ctrl *GWinControl)PaintWindow(dc WinApi.HDC)int32{
	//执行默认的绘制函数
	Brush := WinApi.CreateSolidBrush(uint32(Graphics.ClBtnFace))
	r := new(WinApi.Rect)
	r.Right = ctrl.Width()
	r.Bottom =ctrl.Height()
	WinApi.FillRect(dc,r,Brush)
	WinApi.DeleteObject(uintptr(Brush))
	return 0
}

func (ctrl *GWinControl)PaintHandler(dc WinApi.HDC)uintptr  {
	if dc == 0{
		pstruct := new(WinApi.GPaintStruct)
		dc =pstruct.BeginPaint(ctrl.fHandle)
		defer pstruct.EndPaint(ctrl.fHandle)
	}
	graphicControlCount :=ctrl.ControlCount();
	if graphicControlCount==0{
		//找到最后一个有Paint
		if exeOk,ret := ctrl.ExecuteChildMethod("PaintWindow",dc);exeOk{
			return uintptr(ret.Int())
		}
	}else{
		//裁剪掉GraphcControls的区域不绘制，然后
		SaveIndex := WinApi.SaveDC(dc)
		Clip := WinApi.SIMPLEREGION
		//defer WinApi.RestoreDC(dc,SaveIndex)
		for i := 0;i<graphicControlCount;i++{
			ictrl := ctrl.Controls(i)
			Clip =WinApi.ExcludeClipRect(dc,int(ictrl.Left()),int(ictrl.Top()),
				int(ictrl.Left()+ictrl.Width()),int(ictrl.Top()+ictrl.Height()))
			if Clip == WinApi.NULLREGION {
				break
			}
		}
		if Clip!=WinApi.NULLREGION{
			//找到最后一个有Paint
			if exeok,ret := ctrl.ExecuteChildMethod("PaintWindow",dc);exeok{
				return uintptr(ret.Int())
			}
		}
		WinApi.RestoreDC(dc,SaveIndex)
		//绘制GraphicControls
		ctrl.PaintGraphicControls(dc)
	}
	return 1
}

func (ctrl *GWinControl)PaintGraphicControls(dc WinApi.HDC)  {
	fmt.Println("PaintGraphicControls")
}

func (ctrl *GWinControl) InitSubclassParams(Params *GCreateParams, subclassName string) {
	if subclassName == "" {
		return
	}
	SaveInstance := Params.WindowClass.HInstance
	if Hinstance == 0 {
		Hinstance = WinApi.HINST(WinApi.GetModuleHandle(""))
	}
	if !WinApi.GetClassInfo(Hinstance, subclassName, &Params.WindowClass) &&
		!WinApi.GetClassInfo(0, subclassName, &Params.WindowClass) {
		WinApi.GetClassInfo(SaveInstance, subclassName, &Params.WindowClass)
	}
	Params.WindowClass.HInstance = SaveInstance
	CS_OFF := uint32(WinApi.CS_OWNDC | WinApi.CS_CLASSDC | WinApi.CS_PARENTDC | WinApi.CS_GLOBALCLASS)
	CS_ON := uint32(WinApi.CS_VREDRAW | WinApi.CS_HREDRAW)
	Params.WindowClass.Style = Params.WindowClass.Style & ^CS_OFF | CS_ON
}
