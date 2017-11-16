package controls
import (
	"github.com/suiyunonghen/GVCL/Components"
	"reflect"
	"fmt"
	"github.com/suiyunonghen/GVCL/WinApi"
	"github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/DxCommonLib"
	"unsafe"
	"syscall"
	"bytes"
	"strings"
)

type (
	ListBoxStyle int16

	gListBoxStrings struct{
		DxCommonLib.GStringList
		fListBox			*GListBox
	}
	GListBox struct {
	GWinControl
	fItems		*gListBoxStrings
	OnItemClick 	Graphics.NotifyEvent
	OnItemDblClick	Graphics.NotifyEvent
	fStyle		ListBoxStyle
	}
)

const(
	LBStandard=iota
	LBOwnerDrawFixed
	LBOwnerDrawVariable
	LBVirtual
	LBVirtualOwnerDraw
)

func (list *gListBoxStrings)Count() int  {
	if list.fListBox.fHandle == 0{
		return list.GStringList.Count()
	}
	return int(WinApi.SendMessage(list.fListBox.fHandle, WinApi.LB_GETCOUNT, 0, 0))
}

func (list *gListBoxStrings)Strings(index int) string  {
	if list.fListBox.fHandle == 0{
		return list.GStringList.Strings(index)
	}
	if list.fListBox.fStyle >= LBVirtual{
		return "" //虚拟自填充数据的暂时不处理
	}else{
		l := WinApi.SendMessage(list.fListBox.fHandle, WinApi.LB_GETTEXTLEN, uintptr(index), 0)
		if l == -1 {
			panic("指定的索引不存在")
		}else if l != 0{
			mp := make([]uint16, l+1)
			WinApi.SendMessage(list.fListBox.fHandle, WinApi.LB_GETTEXT, uintptr(index), uintptr(unsafe.Pointer(&mp[0])))
			return syscall.UTF16ToString(mp)
		}
	}
	return ""
}

func (list *gListBoxStrings)SetStrings(index int, str string){
	if list.fListBox.fHandle == 0{
		list.GStringList.SetStrings(index,str)
		return
	}
	i := list.fListBox.GetItemIndex()
	tmpdata := list.fListBox.GetItemData(index)
	list.fListBox.SetItemData(index,0)
	list.Delete(index)
	list.Insert(index,str)
	list.fListBox.SetItemData(index,tmpdata)
	list.fListBox.SetItemIndex(i)
}

func (list *gListBoxStrings)Text() string{
	if list.fListBox.fHandle == 0{
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

func (list *gListBoxStrings)Add(str string){
	if list.fListBox.fHandle == 0{
		list.GStringList.Add(str)
		return
	}
	if list.fListBox.fStyle < LBVirtual{
		mp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
		WinApi.SendMessage(list.fListBox.fHandle, WinApi.LB_ADDSTRING, 0, mp)
	}
}

func (list *gListBoxStrings)SetText(text string){
	if list.fListBox.fHandle == 0{
		list.GStringList.SetText(text)
		return
	}
	list.Clear()
	strs := strings.Split(text, `\r\n`)
	for _,str := range strs{
		list.Add(str)
	}
}

func (list *gListBoxStrings)LoadFromFile(fileName string){
	if list.fListBox.fHandle == 0{
		list.GStringList.LoadFromFile(fileName)
		return
	}
}

func (list *gListBoxStrings)SaveToFile(fileName string){
	if list.fListBox.fHandle == 0{
		list.GStringList.SaveToFile(fileName)
		return
	}
}

func (list *gListBoxStrings)Clear()  {
	if list.fListBox.fHandle == 0{
		list.GStringList.Clear()
		return
	}
	if list.fListBox.fStyle < LBVirtual{
		WinApi.SendMessage(list.fListBox.fHandle, WinApi.LB_RESETCONTENT, 0, 0)
	}
}

func (list *gListBoxStrings)Insert(Index int, str string){
	if list.fListBox.fHandle == 0{
		list.GStringList.Insert(Index,str)
		return
	}
	if list.fListBox.fStyle < LBVirtual{
		mp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
		WinApi.SendMessage(list.fListBox.fHandle, WinApi.LB_INSERTSTRING, uintptr(Index), mp)
	}
}

func (list *gListBoxStrings)Delete(index int){
	if list.fListBox.fHandle == 0{
		list.GStringList.Delete(index)
		return
	}
	WinApi.SendMessage(list.fListBox.fHandle, WinApi.LB_DELETESTRING, uintptr(index), 0)
}

func (list *gListBoxStrings)AddStrings(strs DxCommonLib.IStrings){
	if list.fListBox.fHandle == 0{
		list.GStringList.AddStrings(strs)
		return
	}
	for i := 0;i < strs.Count();i++{
		list.Add(strs.Strings(i))
	}
}

func (list *gListBoxStrings)IndexOf(str string) int{
	if list.fListBox.fHandle == 0{
		return list.GStringList.IndexOf(str)
	}
	if list.fListBox.fStyle < LBVirtual{
		mp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
		m := -1
		return  int(WinApi.SendMessage(list.fListBox.fHandle, WinApi.LB_FINDSTRINGEXACT, uintptr(m), mp))
	}
	return -1
}

func (list *gListBoxStrings)AddPair(Name, Value string){
	if list.fListBox.fHandle == 0{
		list.GStringList.AddPair(Name,Value)
		return
	}
	list.Add(fmt.Sprintf("%s=%s", Name, Value))
}

func (lst *gListBoxStrings)IndexOfName(Name string) int{
	if lst.fListBox.fHandle == 0{
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

func (lst *gListBoxStrings)ValueFromIndex(index int) string{
	if lst.fListBox.fHandle == 0{
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

func (list *gListBoxStrings)ValueByName(Name string) string{
	if list.fListBox.fHandle == 0{
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

func (list *gListBoxStrings)Names(Index int) string{
	if list.fListBox.fHandle == 0{
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

func (list *gListBoxStrings)AsSlice() []string{
	if list.fListBox.fHandle == 0{
		return list.GStringList.AsSlice()
	}
	tc := list.Count()
	result := make([]string,tc)
	for i := 0;i<tc;i++{
		result[i] = list.Strings(i)
	}
	return result
}

func (list *gListBoxStrings)AddSlice(strs []string){
	if list.fListBox.fHandle == 0{
		list.GStringList.AddSlice(strs)
		return
	}
	for _,str := range strs{
		list.Add(str)
	}
}

func (lstbox *GListBox) SubInit() {
	if lstbox.fItems == nil{
		lstbox.fItems = &gListBoxStrings{fListBox:lstbox}
	}
	lstbox.fVisible = true
	lstbox.GWinControl.SubInit()
	lstbox.GComponent.SubInit(lstbox)
}



func (lstbox *GListBox)GetItemIndex()int{
	return int(WinApi.SendMessage(lstbox.fHandle, WinApi.LB_GETCURSEL, 0, 0))
}

func (lstbox *GListBox)SetItemIndex(idx int){
	if lstbox.GetItemIndex() != idx{
		WinApi.SendMessage(lstbox.fHandle, WinApi.LB_SETCURSEL, uintptr(idx), 0)
	}
}

func (lstbox *GListBox) CreateParams(params *Components.GCreateParams) {
	lstbox.GWinControl.CreateParams(params)
	lstbox.InitSubclassParams(params, "LISTBOX")
	var lstyle int32=0
	switch lstbox.fStyle {
	case LBStandard:
		lstyle = 0
	case LBOwnerDrawVariable:
		lstyle = WinApi.LBS_OWNERDRAWVARIABLE
	default:
		lstyle = WinApi.LBS_OWNERDRAWFIXED
	}
	params.Style = params.Style | (WinApi.WS_HSCROLL | WinApi.WS_VSCROLL | WinApi.LBS_HASSTRINGS | WinApi.LBS_NOTIFY) |
		uint32(lstyle) | WinApi.LBS_USETABSTOPS | WinApi.WS_BORDER | WinApi.WS_TABSTOP
	params.ExStyle = params.ExStyle | WinApi.WS_EX_CLIENTEDGE
	lstyle = ^(WinApi.CS_HREDRAW | WinApi.CS_VREDRAW)
	params.WindowClass.Style = params.WindowClass.Style & uint32(lstyle)
}

func (lstbox *GListBox)GetItemData(index int) int{
	return int(WinApi.SendMessage(lstbox.fHandle, WinApi.LB_GETITEMDATA, uintptr(index), 0))
}

func (lstbox *GListBox)SetItemData(index,AData int)  {
	WinApi.SendMessage(lstbox.fHandle, WinApi.LB_SETITEMDATA, uintptr(index), uintptr(AData))
}

func (lstbox *GListBox)Items()DxCommonLib.IStrings{
	return lstbox.fItems
}

func (lstbox *GListBox) CreateWindowHandle(params *Components.GCreateParams)bool{
	if lstbox.GWinControl.CreateWindowHandle(params){
		for i := 0;i<lstbox.fItems.GStringList.Count();i++{
			lp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lstbox.fItems.GStringList.Strings(i))))
			WinApi.SendMessage(lstbox.fHandle, WinApi.LB_ADDSTRING, 0, lp)
		}
		lstbox.fItems.GStringList.Clear()
		return true
	}
	return false
}

func (lstbox *GListBox) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	result = 0
	msgDispatchNext = false
	switch msg {
	case WinApi.WM_COMMAND:
		notifycode := WinApi.HiWord(uint32(wparam))
		switch notifycode {
		case WinApi.LBN_SELCHANGE:
			if lstbox.OnItemClick != nil{
				lstbox.OnItemClick(lstbox)
			}
		case WinApi.LBN_DBLCLK:
			if lstbox.OnItemDblClick != nil{
				lstbox.OnItemDblClick(lstbox)
			}
		case WinApi.LBN_SELCANCEL:
			if lstbox.OnItemClick != nil{
				lstbox.OnItemClick(lstbox)
			}
		}
	default:
		result = WinApi.CallWindowProc(lstbox.FDefWndProc, lstbox.GetWindowHandle(), msg, wparam, lparam)
	}
	return
}

func NewListBox(aParent Components.IWincontrol) *GListBox {
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		lstbox := new(GListBox)
		lstbox.SubInit()
		lstbox.SetColor(Graphics.ClWhite)
		lstbox.SetParent(aParent)
		return lstbox
	}
	return nil
}