package controls
import (
	"github.com/suiyunonghen/GVCL/Components"
	"reflect"
	"bytes"
	"fmt"
	"unsafe"
	"syscall"
	"strings"
	"github.com/suiyunonghen/GVCL/WinApi"
	"github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/DxCommonLib"
)

type GEdit struct {
	GWinControl
	fDefault bool
	OnChange Graphics.NotifyEvent
}

func (edt *GEdit) SubInit() {
	edt.GWinControl.SubInit()
	edt.GComponent.SubInit(edt)
}

func (edt *GEdit) CreateParams(params *Components.GCreateParams) {
	edt.GWinControl.CreateParams(params)
	edt.InitSubclassParams(params, "EDIT")
	params.Style = params.Style | WinApi.ES_AUTOHSCROLL | WinApi.ES_AUTOVSCROLL | WinApi.ES_LEFT
	params.ExStyle = params.ExStyle | WinApi.WS_EX_CLIENTEDGE
}

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

type (
	GComboBoxStyle int16
	gComboBoxStrings struct {
		DxCommonLib.GStringList
		fCombobox			*GCombobox
	}
	GCombobox struct {
		GWinControl
		fItemIndex		int
		fItems			*gComboBoxStrings
		fStyle 			GComboBoxStyle
		fListHandle		syscall.Handle
		fEditHandle 	syscall.Handle
		fDefListProc 	uintptr
		fDefEditProc 	uintptr
		fDropDownCount  int32
		OnChange 		Graphics.NotifyEvent
		OnSelect		Graphics.NotifyEvent
		OnCloseUp		Graphics.NotifyEvent
	}
)

const(
	CSDropDown GComboBoxStyle = iota
	CSSimple
	CSDropDownList
	CSOwnerDrawFixed
	CSOwnerDrawVariable
)

func (list *gComboBoxStrings)Count() int  {
	if list.fCombobox.fHandle == 0{
		return list.GStringList.Count()
	}
	return int(WinApi.SendMessage(list.fCombobox.fHandle, WinApi.CB_GETCOUNT, 0, 0))
}

func (list *gComboBoxStrings)Strings(index int) string  {
	if list.fCombobox.fHandle == 0{
		return list.GStringList.Strings(index)
	}
	l := WinApi.SendMessage(list.fCombobox.fHandle, WinApi.CB_GETLBTEXTLEN, uintptr(index), 0)
	if l == -1 {
		panic("指定的索引不存在")
	}else if l != 0{
		mp := make([]uint16, l+1)
		WinApi.SendMessage(list.fCombobox.fHandle, WinApi.CB_GETLBTEXT, uintptr(index), uintptr(unsafe.Pointer(&mp[0])))
		return syscall.UTF16ToString(mp)
	}
	return ""
}

func (list *gComboBoxStrings)SetStrings(index int, str string){
	if list.fCombobox.fHandle == 0{
		list.GStringList.SetStrings(index,str)
		return
	}
	i := list.fCombobox.GetItemIndex()
	list.Delete(index)
	list.Insert(index,str)
	list.fCombobox.SetItemIndex(i)
}

func (list *gComboBoxStrings)Text() string{
	if list.fCombobox.fHandle == 0{
		return list.GStringList.Text()
	}
	var bf bytes.Buffer
	tc := list.Count()
	for i := 0;i<tc;i++{
		bf.WriteString(list.Strings(i))
		if i < tc-1{
			bf.WriteString(`\r\n`)
		}
	}
	return bf.String()
}

func (list *gComboBoxStrings)Add(str string){
	if list.fCombobox.fHandle == 0{
		list.GStringList.Add(str)
		return
	}
	mp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
	WinApi.SendMessage(list.fCombobox.fHandle, WinApi.CB_ADDSTRING, 0, mp)
}

func (list *gComboBoxStrings)SetText(text string){
	if list.fCombobox.fHandle == 0{
		list.GStringList.SetText(text)
		return
	}
	list.Clear()
	strs := strings.Split(text, `\r\n`)
	for _,str := range strs{
		list.Add(str)
	}
}

func (list *gComboBoxStrings)LoadFromFile(fileName string){
	if list.fCombobox.fHandle == 0{
		list.GStringList.LoadFromFile(fileName)
		return
	}
}

func (list *gComboBoxStrings)SaveToFile(fileName string){
	if list.fCombobox.fHandle == 0{
		list.GStringList.SaveToFile(fileName)
		return
	}
}

func (list *gComboBoxStrings)Clear()  {
	if list.fCombobox.fHandle == 0{
		list.GStringList.Clear()
		return
	}
	WinApi.SendMessage(list.fCombobox.fHandle, WinApi.CB_RESETCONTENT, 0, 0)
}

func (list *gComboBoxStrings)Insert(Index int, str string){
	if list.fCombobox.fHandle == 0{
		list.GStringList.Insert(Index,str)
		return
	}
	mp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
	WinApi.SendMessage(list.fCombobox.fHandle, WinApi.CB_INSERTSTRING, uintptr(Index), mp)
}

func (list *gComboBoxStrings)Delete(index int){
	if list.fCombobox.fHandle == 0{
		list.GStringList.Delete(index)
		return
	}
	WinApi.SendMessage(list.fCombobox.fHandle, WinApi.CB_DELETESTRING, uintptr(index), 0)
}

func (list *gComboBoxStrings)AddStrings(strs DxCommonLib.IStrings){
	if list.fCombobox.fHandle == 0{
		list.GStringList.AddStrings(strs)
		return
	}
	for i := 0;i < strs.Count();i++{
		list.Add(strs.Strings(i))
	}
}

func (list *gComboBoxStrings)IndexOf(str string) int{
	if list.fCombobox.fHandle == 0{
		return list.GStringList.IndexOf(str)
	}
	mp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
	m := -1
	return  int(WinApi.SendMessage(list.fCombobox.fHandle, WinApi.CB_FINDSTRINGEXACT, uintptr(m), mp))
}

func (list *gComboBoxStrings)AddPair(Name, Value string){
	if list.fCombobox.fHandle == 0{
		list.GStringList.AddPair(Name,Value)
		return
	}
	list.Add(fmt.Sprintf("%s=%s", Name, Value))
}

func (lst *gComboBoxStrings)IndexOfName(Name string) int{
	if lst.fCombobox.fHandle == 0{
		return lst.GStringList.IndexOfName(Name)
	}
	for i:=0;i<lst.Count();i++{
		v := lst.Strings(i)
		if eidx := strings.IndexByte(v, '='); eidx > 0 {
			bt := []byte(v)
			if string(bt[:eidx]) == Name {
				return i
			}
		}
	}
	return -1
}

func (lst *gComboBoxStrings)ValueFromIndex(index int) string{
	if lst.fCombobox.fHandle == 0{
		return lst.GStringList.ValueFromIndex(index)
	}
	tc := lst.Count()
	if tc == 0 {
		return ""
	}
	if index >= 0 && index < tc {
		str := lst.Strings(index)
		if idx := strings.IndexByte(str, '='); idx > 0 {
			bt := []byte(str)
			return DxCommonLib.FastByte2String(bt[idx+1:])
		}
		return ""
	}
	return ""
}

func (list *gComboBoxStrings)ValueByName(Name string) string{
	if list.fCombobox.fHandle == 0{
		return list.GStringList.ValueByName(Name)
	}
	for i := 0;i<list.Count();i++{
		v := list.Strings(i)
		if eidx := strings.IndexByte(v, '='); eidx > 0 {
			bt := []byte(v)
			if string(bt[:eidx]) == Name {
				return DxCommonLib.FastByte2String(bt[eidx+1:])
			}
		}
	}
	return ""
}

func (list *gComboBoxStrings)Names(Index int) string{
	if list.fCombobox.fHandle == 0{
		return list.GStringList.Names(Index)
	}
	if list.Count() == 0 {
		return ""
	}
	if Index >= 0 && Index < list.Count() {
		str := list.Strings(Index)
		if idx := strings.IndexByte(str, '='); idx > 0 {
			bt := []byte(str)
			return DxCommonLib.FastByte2String(bt[:idx])
		}
		return ""
	}
	return ""
}

func (list *gComboBoxStrings)AsSlice() []string{
	if list.fCombobox.fHandle == 0{
		return list.GStringList.AsSlice()
	}
	tc := list.Count()
	result := make([]string,tc)
	for i := 0;i<tc;i++{
		result[i] = list.Strings(i)
	}
	return result
}

func (list *gComboBoxStrings)AddSlice(strs []string){
	if list.fCombobox.fHandle == 0{
		list.GStringList.AddSlice(strs)
		return
	}
	for _,str := range strs{
		list.Add(str)
	}
}

func (cmbox *GCombobox)GetItemIndex()int{
	return int(WinApi.SendMessage(cmbox.fHandle, WinApi.CB_GETCURSEL, 0, 0))
}

func (cmbox *GCombobox)SetItemIndex(idx int){
	if cmbox.HandleAllocated(){
		if cmbox.GetItemIndex() != idx{
			WinApi.SendMessage(cmbox.fHandle, WinApi.CB_GETCURSEL, uintptr(idx), 0)
		}
	}else{
		cmbox.fItemIndex = idx
	}
}

func (cmbox *GCombobox) CreateParams(params *Components.GCreateParams) {
	cmbox.GWinControl.CreateParams(params)
	cmbox.InitSubclassParams(params, "COMBOBOX")
	params.Style = params.Style | (WinApi.WS_VSCROLL | WinApi.CBS_HASSTRINGS | WinApi.CBS_AUTOHSCROLL)
	switch cmbox.fStyle {
	case CSDropDown: params.Style = params.Style | WinApi.CBS_DROPDOWN
	case CSSimple: params.Style = params.Style | WinApi.CBS_SIMPLE
	case CSDropDownList: params.Style = params.Style | WinApi.CBS_DROPDOWNLIST
	case CSOwnerDrawFixed: params.Style = params.Style | WinApi.CBS_DROPDOWNLIST | WinApi.CBS_OWNERDRAWFIXED
	case CSOwnerDrawVariable: params.Style = params.Style | WinApi.CBS_DROPDOWNLIST | WinApi.CBS_OWNERDRAWVARIABLE
	}
}

var(
	cmbListWndprocCallBack = syscall.NewCallback(cmbListWndProc)
	cmbEdtWndprocCallBack = syscall.NewCallback(cmbEdtWndProc)
)

func cmbListWndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (result uintptr) {
	cmbox := (*GCombobox)(unsafe.Pointer(WinApi.GetProp(hwnd, uintptr(controlAtom))))
	if cmbox == nil {
		return WinApi.DefWindowProc(hwnd, msg, wparam, lparam)
	}
	result = cmbox.comboWndProc(msg,wparam,lparam,hwnd,cmbox.fDefEditProc)
	return
}

func cmbEdtWndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (result uintptr) {
	cmbox := (*GCombobox)(unsafe.Pointer(WinApi.GetProp(hwnd, uintptr(controlAtom))))
	if cmbox == nil {
		return WinApi.DefWindowProc(hwnd, msg, wparam, lparam)
	}
	if msg == WinApi.WM_SYSCOMMAND{
		result,_ := cmbox.WndProc(msg,wparam,lparam)
		return result
	}
	result = cmbox.comboWndProc(msg,wparam,lparam,hwnd,cmbox.fDefEditProc)
	return
}

func (cmbox *GCombobox)comboWndProc(msg uint32, wparam, lparam uintptr,msghwnd syscall.Handle,defwndproc uintptr)  (result uintptr) {
	return WinApi.CallWindowProc(defwndproc,msghwnd,msg,wparam,lparam)
}

func (cmbox *GCombobox) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	result = 0
	msgDispatchNext = false
	switch msg {
	case WinApi.WM_COMMAND:
		notifycode := WinApi.HiWord(uint32(wparam))
		switch notifycode {
		case WinApi.CBN_DBLCLK:
			fmt.Println("Asdf")
		case WinApi.CBN_EDITCHANGE:
			if cmbox.OnChange != nil{
				cmbox.OnChange(cmbox)
			}
		case WinApi.CBN_SELCHANGE:
			index := cmbox.GetItemIndex()
			strptr,_ := syscall.UTF16PtrFromString(cmbox.fItems.Strings(index))
			cmbox.Perform(WinApi.WM_SETTEXT,0,uintptr(unsafe.Pointer(strptr)))
			WinApi.SendMessage(cmbox.fHandle, WinApi.CB_SHOWDROPDOWN, 0, 0)
			if cmbox.OnSelect!=nil{
				cmbox.OnSelect(cmbox)
			}else if cmbox.OnChange != nil{
				cmbox.OnChange(cmbox)
			}
		case WinApi.CBN_CLOSEUP:
			if cmbox.OnCloseUp!=nil{
				cmbox.OnCloseUp(cmbox)
			}
		default:
			result = WinApi.CallWindowProc(cmbox.FDefWndProc, cmbox.fHandle, msg, wparam, lparam)
		}
	case WinApi.WM_SETFONT:
		result = WinApi.CallWindowProc(cmbox.FDefWndProc, cmbox.fHandle, msg, wparam, lparam)
		//设置实际高度
		ItemHeight := int32(WinApi.SendMessage(cmbox.fHandle,WinApi.CB_GETITEMHEIGHT,0,0))
		WinApi.SetWindowPos(cmbox.fHandle,0, 0, 0, cmbox.fwidth, ItemHeight * cmbox.fDropDownCount +
			cmbox.fheight + 2, WinApi.SWP_NOMOVE | WinApi.SWP_NOZORDER | WinApi.SWP_NOACTIVATE | WinApi.SWP_NOREDRAW |
			WinApi.SWP_SHOWWINDOW)
	default:
		result = WinApi.CallWindowProc(cmbox.FDefWndProc, cmbox.fHandle, msg, wparam, lparam)
	}
	return
}

func (cmbox *GCombobox)DroppedDown()bool{
	if cmbox.fHandle !=0{
		return WinApi.SendMessage(cmbox.fHandle, WinApi.CB_GETDROPPEDSTATE, 0, 0)!=0
	}
	return false
}

func (cmbox *GCombobox)SetDroppedDown(v bool){
	if cmbox.fHandle!=0{
		WinApi.SendMessage(cmbox.fHandle, WinApi.CB_SHOWDROPDOWN, uintptr(DxCommonLib.Ord(v)), 0)
		r := cmbox.ClientRect()
		WinApi.InvalidateRect(cmbox.fHandle,&r,true)
	}
}

func (cmbox *GCombobox) CreateWindowHandle(params *Components.GCreateParams)bool{
	if cmbox.GWinControl.CreateWindowHandle(params){
		for i := 0;i<cmbox.fItems.GStringList.Count();i++{
			lp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(cmbox.fItems.GStringList.Strings(i))))
			WinApi.SendMessage(cmbox.fHandle, WinApi.CB_ADDSTRING, 0, lp)
		}
		if cmbox.fItemIndex < cmbox.fItems.GStringList.Count(){
			WinApi.SendMessage(cmbox.fHandle, WinApi.CB_SETCURSEL, uintptr(cmbox.fItemIndex), 0)
		}
		cmbox.fItems.GStringList.Clear()
		if cmbox.fStyle == CSDropDown || cmbox.fStyle == CSSimple{
			if ChildHandle := WinApi.GetWindow(cmbox.fHandle, WinApi.GW_CHILD);ChildHandle != 0{
				if cmbox.fStyle == CSSimple{
					cmbox.fListHandle = ChildHandle
					WinApi.SetProp(ChildHandle, uintptr(controlAtom), uintptr(unsafe.Pointer(cmbox)))
					if WinApi.IsAMD64(){
						//指定窗口过程
						cmbox.fDefListProc = uintptr(WinApi.SetWindowLongPtr(ChildHandle,WinApi.GWL_WNDPROC,int64(cmbListWndprocCallBack)))
					}else{
						cmbox.fDefListProc = uintptr(WinApi.SetWindowLong(ChildHandle,WinApi.GWL_WNDPROC,int(cmbListWndprocCallBack)))
					}
					ChildHandle = WinApi.GetWindow(ChildHandle, WinApi.GW_HWNDNEXT)
				}
				cmbox.fEditHandle = ChildHandle
				WinApi.SetProp(ChildHandle, uintptr(controlAtom), uintptr(unsafe.Pointer(cmbox)))
				if WinApi.IsAMD64(){
					//指定窗口过程
					cmbox.fDefEditProc = uintptr(WinApi.SetWindowLongPtr(ChildHandle,WinApi.GWL_WNDPROC,int64(cmbEdtWndprocCallBack)))
				}else{
					cmbox.fDefEditProc = uintptr(WinApi.SetWindowLong(ChildHandle,WinApi.GWL_WNDPROC,int(cmbEdtWndprocCallBack)))
				}
			}
		}
		return true
	}
	return false
}

func (cmbox *GCombobox)Items()DxCommonLib.IStrings  {
	return cmbox.fItems
}

func (cmbox *GCombobox) SubInit() {
	if cmbox.fItems == nil{
		cmbox.fItems = &gComboBoxStrings{fCombobox:cmbox}
	}
	cmbox.GWinControl.SubInit()
	cmbox.GComponent.SubInit(cmbox)
}

func (cmbox *GCombobox)SetDropDownCount(v int)  {
	if cmbox.fDropDownCount != int32(v){
		cmbox.fDropDownCount = int32(v)
		if cmbox.fHandle != 0{
			ItemHeight := int32(WinApi.SendMessage(cmbox.fHandle,WinApi.CB_GETITEMHEIGHT,0,0))
			WinApi.SetWindowPos(cmbox.fHandle,0, 0, 0, cmbox.fwidth, ItemHeight * cmbox.fDropDownCount +
				cmbox.fheight + 2, WinApi.SWP_NOMOVE | WinApi.SWP_NOZORDER | WinApi.SWP_NOACTIVATE | WinApi.SWP_NOREDRAW |
				WinApi.SWP_SHOWWINDOW)
		}
	}
}


func NewCombobox(aParent Components.IWincontrol,mstyle GComboBoxStyle) *GCombobox {
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		cmb := new(GCombobox)
		cmb.SubInit()
		cmb.fwidth = 145
		cmb.fDropDownCount = 8
		cmb.fVisible = true
		cmb.fStyle = mstyle
		cmb.fheight = 25
		cmb.fColor = Graphics.ClWhite
		cmb.SetParent(aParent)
		return cmb
	}
	return nil
}