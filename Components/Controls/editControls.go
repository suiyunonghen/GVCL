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
		fItemIndex	int
		fItems		*gComboBoxStrings
		fStyle 		GComboBoxStyle
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

func NewCombobox(aParent Components.IWincontrol) *GCombobox {
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		cmb := new(GCombobox)
		cmb.SubInit()
		cmb.fwidth = 145
		cmb.fVisible = true
		cmb.fheight = 25
		cmb.fColor = Graphics.ClWhite
		cmb.SetParent(aParent)
		return cmb
	}
	return nil
}