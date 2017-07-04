/*
Scintilla编辑器控件的封装
Autor: 不得闲
QQ:75492895
 */
package Scintilla

import (
	"github.com/suiyunonghen/GVCL/Components"
	"github.com/suiyunonghen/GVCL/WinApi"
	"github.com/suiyunonghen/DxCommonLib"
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"reflect"
	"syscall"
	"unsafe"
	"os"
	"strings"
	"fmt"
	"github.com/suiyunonghen/GVCL/Graphics"
)



type (
	scintillaLib struct {
		fHandle			syscall.Handle
		fDirectFunction uintptr			//导出的窗口函数
	}


	GCodeLines struct {
		DxCommonLib.GStringList
		fCodeEditor			*GScintilla
	}

	GCaretPos struct {
		Line			int
		Column			int
	}

	GScintilla struct {
		controls.GWinControl
		fDirectPointer				uintptr
		CodeLines					*GCodeLines
		defFont						*GDxLexerFont
		MarginBand					*GDxMarginBand
		fInitLexer					bool		//是否初始化了语法分析器
		SelectFore					Graphics.GColorValue	//选中前景色
		SelectBack					Graphics.GColorValue	//选中背景色
		LineSpaceBelow				int				//本行下间距
		LineSpaceAbove				int				//本行上间距
		TabWidth					int				//Tab宽度
		UseTabSpaceChar				bool			//是否将Tab转成以完整的空格显示
		CaretPos					GCaretPos
		fLastCaretPos				GCaretPos
		OnCaretPosChange			Graphics.NotifyEvent
		fLexer						ILanguageLexer
	}
)

var(
	scintillaHandle		*scintillaLib
	cDSciNull  byte = 0
)


func (lines *GCodeLines)Clear()  {
	lines.GStringList.Clear()
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		lines.fCodeEditor.SendEditor(SCI_CLEARALL,0,0)
	}
}

func (lines *GCodeLines)SetStrings(index int,str string){
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		if lines.fCodeEditor.setTargetLine(index) {
			if str == ""{
				lines.fCodeEditor.SendEditor(SCI_REPLACETARGET, 0, int(uintptr(unsafe.Pointer(&cDSciNull))))
			}else{
				bt := DxCommonLib.FastString2Byte(str)
				lines.fCodeEditor.SendEditor(SCI_REPLACETARGET, len(bt), int(uintptr(unsafe.Pointer(&bt[0]))))
			}
		}
	}else{
		lines.GStringList.SetStrings(index,str)
	}
}

func (lines *GCodeLines)Count()(Result int){
	Result = 0
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		Result = lines.fCodeEditor.SendEditor(SCI_GETLINECOUNT,0,0)
		if Result == 1{
			if  lines.fCodeEditor.SendEditor(SCI_GETLENGTH,0,0) == 0 {
				Result = 0
			}
		}
	}else{
		Result = lines.GStringList.Count()
	}
	return
}

func (lines *GCodeLines)Strings(index int)string  {
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		lTextRange := GSci_TextRange{}
		lTextRange.chrg.cpMin = GSciPositionCR(lines.fCodeEditor.SendEditor(SCI_POSITIONFROMLINE, index,0))
		lTextRange.chrg.cpMax = GSciPositionCR(lines.fCodeEditor.SendEditor(SCI_GETLINEENDPOSITION, index,0))
		if lTextRange.chrg.cpMin == INVALID_POSITION || lTextRange.chrg.cpMax == INVALID_POSITION ||
			lTextRange.chrg.cpMax == lTextRange.chrg.cpMin{
			return ""
		}
		lLen := lTextRange.chrg.cpMax - lTextRange.chrg.cpMin
		lbufer := make([]byte, lLen + 1)
		lTextRange.lpstrText = &lbufer[0]
		lines.fCodeEditor.SendEditor(SCI_GETTEXTRANGE, 0, int(uintptr(unsafe.Pointer(&lTextRange))))
		return DxCommonLib.FastByte2String(lbufer[:lLen])
	}else{
		return lines.GStringList.Strings(index)
	}
	return ""
}

func (lines *GCodeLines)Text()string {
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		lLen := lines.fCodeEditor.SendEditor(SCI_GETLENGTH,0,0)
		if lLen != 0{
			lBuf := make([]byte,lLen+1)
			lines.fCodeEditor.SendEditor(SCI_GETTEXT, lLen + 1, int(uintptr(unsafe.Pointer(&lBuf[0]))))
			return DxCommonLib.FastByte2String(lBuf[:lLen])
		}
	}else{
		return lines.GStringList.Text()
	}
	return ""
}

func (lines *GCodeLines)SetText(text string){
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		if text == ""{
			lines.fCodeEditor.SendEditor(SCI_SETTEXT, 0, int(uintptr(unsafe.Pointer(&cDSciNull))))
		}else{
			bt := DxCommonLib.FastString2Byte(text)
			lines.fCodeEditor.SendEditor(SCI_SETTEXT, 0, int(uintptr(unsafe.Pointer(&bt[0]))))
		}
	}else{
		lines.GStringList.SetText(text)
	}
}
func (lines *GCodeLines)LoadFromFile(fileName string){
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		if finfo,err := os.Stat(fileName);err == nil && !finfo.IsDir(){
			if file,err := os.Open(fileName);err == nil{
				databytes := make([]byte,finfo.Size())
				file.Read(databytes)
				isUtf8 := databytes[0] == 0xEF && databytes[1] == 0xBB && databytes[2] == 0xBF
				if isUtf8{
					databytes = databytes[3:]
				}
				file.Close()
				if !isUtf8{
					if tmpbytes,err := DxCommonLib.GBK2Utf8(databytes);err == nil{
						databytes = tmpbytes
					}
				}
				lines.fCodeEditor.SendEditor(SCI_SETTEXT, 0, int(uintptr(unsafe.Pointer(&databytes[0]))))
			}
		}
	}else{
		lines.GStringList.LoadFromFile(fileName)
	}
}

func (lines *GCodeLines)SaveToFile(fileName string){
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		if file,err := os.OpenFile(fileName,os.O_CREATE | os.O_TRUNC,0644);err == nil{
			lLen := lines.fCodeEditor.SendEditor(SCI_GETLENGTH,0,0)
			if lLen != 0{
				lBuf := make([]byte,lLen+1)
				lines.fCodeEditor.SendEditor(SCI_GETTEXT, lLen + 1, int(uintptr(unsafe.Pointer(&lBuf[0]))))
				file.Write([]byte{0xEF,0xBB,0xBF})
				file.Write(lBuf[:lLen])
			}
			file.Close()
		}
	}else{
		lines.GStringList.SaveToFile(fileName)
	}
}

func (lines *GCodeLines)Add(str string){
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		lines.Insert(lines.Count(), str)
	}else{
		lines.GStringList.Add(str)
	}
}

func (lines *GCodeLines)Insert(Index int,str string){
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		lLinePos := lines.fCodeEditor.SendEditor(SCI_POSITIONFROMLINE, Index,0)
		if lLinePos == INVALID_POSITION{
			return
		}
		lines.fCodeEditor.SendEditor(SCI_SETTARGETSTART, lLinePos,0)
		lines.fCodeEditor.SendEditor(SCI_SETTARGETEND, lLinePos,0)
		if lLinePos == lines.fCodeEditor.SendEditor(SCI_GETLENGTH,0,0){
			if lLinePos != 0{
				str = fmt.Sprintf("%s%s",lines.LineBreakStr(),str)
			}
		}else{
			str = fmt.Sprintf("%s%s",str,lines.LineBreakStr())
		}
		if str == ""{
			lines.fCodeEditor.SendEditor(SCI_REPLACETARGET, 0, int(uintptr(unsafe.Pointer(&cDSciNull))))
		}else{
			bt := DxCommonLib.FastString2Byte(str)
			lines.fCodeEditor.SendEditor(SCI_REPLACETARGET, len(bt),int(uintptr(unsafe.Pointer(&bt[0]))))
		}
	}else{
		lines.GStringList.Insert(Index,str)
	}
}
func (lines *GCodeLines)Delete(index int){
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		if lines.fCodeEditor.setTargetLine(index){
			lines.fCodeEditor.SendEditor(SCI_REPLACETARGET, 0, 0)
		}
	}else{
		lines.GStringList.Delete(index)
	}
}

func (lines *GCodeLines)AddStrings(strs DxCommonLib.IStrings){
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		count := lines.Count()
		for idx := 0;idx < strs.Count();idx++{
			lines.Insert(count, strs.Strings(idx))
			count++
		}
	}else{
		lines.GStringList.AddStrings(strs)
	}
}

func (lines *GCodeLines)AddSlice(strs []string){
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		count := lines.Count()
		for idx := 0;idx < len(strs);idx++{
			lines.Insert(count, strs[idx])
			count++
		}
	}else{
		lines.GStringList.AddSlice(strs)
	}
}

func (lines *GCodeLines)AsSlice()[]string  {
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		return strings.Split(lines.Text(),lines.LineBreakStr())
	}else{
		return lines.GStringList.AsSlice()
	}
}

func (lines *GCodeLines)IndexOf(str string) int{
	if lines.fCodeEditor != nil && lines.fCodeEditor.HandleAllocated(){
		for i := 0;i<lines.Count();i++{
			if lines.Strings(i) == str{
				return i
			}
		}
		return -1
	}else{
		return lines.GStringList.IndexOf(str)
	}
}

func (Scintilla *GScintilla) SubInit() {
	Scintilla.GWinControl.SubInit()
	Scintilla.GComponent.SubInit(Scintilla)
}

func (Scintilla *GScintilla)setTargetLine(ALine int) bool  {
	lLineStart := Scintilla.SendEditor(SCI_POSITIONFROMLINE, ALine,0)
	if lLineStart == INVALID_POSITION {
		return false
	}
	lLineEnd := 0
	if lLineStart == Scintilla.SendEditor(SCI_GETLENGTH,0,0) && ALine > 0{
		lLineEnd = lLineStart
		lLineStart = Scintilla.SendEditor(SCI_GETLINEENDPOSITION, ALine - 1,0)
	}else{
		lLineEnd = lLineStart + Scintilla.SendEditor(SCI_LINELENGTH, ALine,0)
	}
	if lLineEnd == INVALID_POSITION {
		return false
	}
	Scintilla.SendEditor(SCI_SETTARGETSTART, lLineStart,0)
	Scintilla.SendEditor(SCI_SETTARGETEND, lLineEnd,0)
	return true
}
func (Scintilla *GScintilla) SetColor(c Graphics.GColorValue) {
	if Scintilla.GetColor() != c {
		Scintilla.GBaseControl.SetColor(c)
		Scintilla.SendEditor(SCI_STYLESETBACK,32,int(c)); //设置编辑区窗口背景色
	}
}
func (Scintilla *GScintilla) CreateParams(params *Components.GCreateParams) {
	Scintilla.GWinControl.CreateParams(params)
	Scintilla.InitSubclassParams(params, "SCINTILLA")
}

func (Scintilla *GScintilla)SendEditor(AMessage,WParam,LParam int)int  {
	if !Scintilla.HandleAllocated(){
		return -1
	}
	if scintillaHandle.fDirectFunction != 0{
		ret,_,_:= syscall.Syscall6(scintillaHandle.fDirectFunction,4,Scintilla.fDirectPointer,uintptr(AMessage),uintptr(WParam),uintptr(LParam),0,0)
		return int(ret)
	}else{
		return int(WinApi.SendMessage(Scintilla.GetWindowHandle(),uint(AMessage),uintptr(WParam),uintptr(LParam)))
	}

}


func (Scintilla *GScintilla) initScintilla()  {
	//设定Utf8编码
	Scintilla.SendEditor(SCI_SETKEYSUNICODE, 1, 0)
	Scintilla.SendEditor(SCI_SETCODEPAGE, SC_CP_UTF8, 0)
	Scintilla.SendEditor(SCI_STYLESETBACK,32,int(Scintilla.GetColor())) //设置编辑区窗口背景色
	Scintilla.SendEditor(SCI_SETEOLMODE,32,int(Scintilla.CodeLines.LineBreak)) //设定换行模式
	//设置选择的前景色，背景色
	Scintilla.SendEditor(SCI_SETSELFORE,1,int(Scintilla.SelectFore))
	Scintilla.SendEditor(SCI_SETSELBACK,1,int(Scintilla.SelectBack))
	//设置行间距
	Scintilla.SendEditor(SCI_SETEXTRADESCENT,Scintilla.LineSpaceBelow,0)
	Scintilla.SendEditor(SCI_SETEXTRAASCENT,Scintilla.LineSpaceAbove,0)
	//Tab设置
	Scintilla.SendEditor(SCI_SETTABWIDTH,Scintilla.TabWidth,0)
	Scintilla.SendEditor(SCI_SETUSETABS,int(DxCommonLib.Ord(!Scintilla.UseTabSpaceChar)),0)
	Scintilla.SendEditor(SCI_SETTABINDENTS,1,0) //可以使用Tab进行缩进块
	if Scintilla.fLexer != nil{
		Scintilla.fLexer.Update()
	}else{
		Scintilla.defFont.InitLexFont()
	}
	Scintilla.MarginBand.Update()
}

func (Scintilla *GScintilla) CreateWindowHandle(params *Components.GCreateParams)bool{
   if Scintilla.GWinControl.CreateWindowHandle(params){
	   Scintilla.CodeLines.fCodeEditor = Scintilla
	   h := Scintilla.GetWindowHandle()
	   //typedef sptr_t (*SciFnDirect)(sptr_t ptr, unsigned int iMessage, uptr_t wParam, sptr_t lParam);
	   //Scintilla.fDirectPointer = uintptr(WinApi.SendMessage(h,SCI_GETDIRECTFUNCTION, 0, 0))
	   //这个用的Cdecl方式，调用有问题，所以用导出函数Scintilla_DirectFunction
	   //Scintilla.fDirectFunction = scintillaHandle.fDirectFunction
	   Scintilla.fDirectPointer = uintptr(WinApi.SendMessage(h,SCI_GETDIRECTPOINTER, 0, 0))
	   Scintilla.initScintilla()
	   //然后从CodeLines中加载内容信息
	   Scintilla.CodeLines.SetText(Scintilla.CodeLines.GStringList.Text())
	   Scintilla.CodeLines.GStringList.Clear()
	   return true
   }
	return false
}

func (Scintilla *GScintilla)Clear()  {
	Scintilla.CodeLines.Clear()
	Scintilla.MarginBand.ClearMarks()
}

func (Scintilla *GScintilla)ClearSelection()  {
	Scintilla.SendEditor(SCI_CLEAR,0,0)
}

func (Scintilla *GScintilla)ClearUndo()  {
	Scintilla.SendEditor(SCI_EMPTYUNDOBUFFER,0,0)
}

func (Scintilla *GScintilla)CopyToClipboard()  {
	Scintilla.SendEditor(SCI_COPY,0,0)
}

func (Scintilla *GScintilla)CutToClipboard()  {
	Scintilla.SendEditor(SCI_CUT,0,0)
}

func (Scintilla *GScintilla)Undo()  {
	Scintilla.SendEditor(SCI_UNDO,0,0)
}

func (Scintilla *GScintilla)Redo()  {
	Scintilla.SendEditor(SCI_REDO,0,0)
}

func (Scintilla *GScintilla)SelectAll()  {
	Scintilla.SendEditor(SCI_SELECTALL,0,0)
}

func (Scintilla *GScintilla)CanUndo()bool  {
	return Scintilla.SendEditor(SCI_CANUNDO,0,0)==1
}


func (Scintilla *GScintilla)CanRedo()bool  {
	return Scintilla.SendEditor(SCI_CANREDO,0,0) == 1
}

func (Scintilla *GScintilla)GoToLine(line int)  {
	Scintilla.SendEditor(SCI_GOTOLINE,line,0)
}

func (Scintilla *GScintilla)GoToPos(pos int)  {
	Scintilla.SendEditor(SCI_GOTOPOS,pos,0)
}

func (Scintilla *GScintilla)GoToCaretPos(ACaretPos GCaretPos)  {
	imax := Scintilla.CodeLines.Count() - 1
	if ACaretPos.Line > imax{
		ACaretPos.Line = imax
	}
	StartPos := Scintilla.SendEditor(SCI_POSITIONFROMLINE,ACaretPos.Line,0)
	Endpos := Scintilla.SendEditor(SCI_GETLINEENDPOSITION,ACaretPos.Line,0)
	StartPos += ACaretPos.Column
	if StartPos > Endpos {
		StartPos = Endpos
	}
	Scintilla.SendEditor(SCI_GOTOPOS,StartPos,0)
}

func (Scintilla *GScintilla)StyleClearAll()  {
	//清空样式
	Scintilla.SendEditor(SCI_STYLECLEARALL, 0, 0)
	Scintilla.defFont.InitLexFont()
}

func (Scintilla *GScintilla) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	result = 0
	msgDispatchNext = false
	switch msg {
	case WinApi.WM_DESTROY: //释放的时候触发
		//保存数据到List
		//Scintilla.CodeLines.GStringList.SetText(Scintilla.CodeLines.Text())
		Scintilla.CodeLines.fCodeEditor = nil
		Scintilla.MarginBand.fcodeEditor = nil
		Scintilla.fDirectPointer = 0
	case WinApi.WM_ERASEBKGND:
		result = 1
	case WinApi.WM_GETDLGCODE:
		result = WinApi.CallWindowProc(Scintilla.FDefWndProc, Scintilla.GetWindowHandle(), msg, wparam, lparam)
		result = result | WinApi.DLGC_WANTARROWS | WinApi.DLGC_WANTCHARS
		result = result | WinApi.DLGC_WANTTAB
		result = result | WinApi.DLGC_WANTALLKEYS
	case WinApi.WM_COMMAND:
		result = WinApi.CallWindowProc(Scintilla.FDefWndProc, Scintilla.GetWindowHandle(), msg, wparam, lparam)
		notifycode := WinApi.HiWord(uint32(wparam))
		switch notifycode {
		case SCEN_CHANGE:
			if Scintilla.MarginBand.ShowLineNum{
				str := fmt.Sprintf("_%d",Scintilla.CodeLines.Count())
				Scintilla.SendEditor(SCI_SETMARGINWIDTHN,int(Scintilla.MarginBand.fLineNumIndex), Scintilla.MarginBand.MarginTextLen(str)) //页边长度
			}
		case SCEN_SETFOCUS:
		case SCEN_KILLFOCUS:
		}
	case WinApi.WM_NOTIFY:
		PSCNotify := (*GSCNotification)(unsafe.Pointer(lparam))
		switch PSCNotify.NotifyHeader.NMCode {
		case SCN_STYLENEEDED://对于自定义语法的样式，这里会触发
		case SCN_CHARADDED://字符变动的时候触发，自动完成列表和代码提示可以在这里设定
		case SCN_SAVEPOINTREACHED,SCN_SAVEPOINTLEFT:  //SCI_SETSAVEPOINT, SCI_GETMODIFY消息触发
		case SCN_MODIFYATTEMPTRO:  //只读时，用户输入会触发：
		case SCN_DOUBLECLICK:
		case SCN_UPDATEUI:
			startPos := Scintilla.SendEditor(SCI_GETCURRENTPOS,0,0) //取得当前位置
			Scintilla.CaretPos.Line = Scintilla.SendEditor(SCI_LINEFROMPOSITION,startPos,0)
			Scintilla.CaretPos.Column = Scintilla.SendEditor(SCI_GETCOLUMN,startPos,0)
			if Scintilla.fLastCaretPos.Line != Scintilla.CaretPos.Line || Scintilla.CaretPos.Column != Scintilla.fLastCaretPos.Column{
				Scintilla.fLastCaretPos.Line = Scintilla.CaretPos.Line
				Scintilla.fLastCaretPos.Column = Scintilla.CaretPos.Column
				if Scintilla.OnCaretPosChange != nil{
					Scintilla.OnCaretPosChange(Scintilla)
				}
			}
		case SCN_MARGINCLICK:
			Scintilla.MarginBand.BandClick(PSCNotify.Position,PSCNotify.Margin,PSCNotify.Modifiers)
		case SCN_MARGINRIGHTCLICK:
			if Scintilla.MarginBand.OnRightClick != nil{
				Scintilla.MarginBand.OnRightClick(Scintilla.MarginBand,PSCNotify.Position,PSCNotify.Margin,PSCNotify.Modifiers)
			}
		case SCN_NEEDSHOWN:
		case SCN_USERLISTSELECTION:
		case SCN_CALLTIPCLICK:
		case SCN_AUTOCSELECTION:
		case SCN_AUTOCCANCELLED:
		case SCN_AUTOCCHARDELETED:
		case SCN_AUTOCCOMPLETED:
			
		}
	default:
		result = WinApi.CallWindowProc(Scintilla.FDefWndProc, Scintilla.GetWindowHandle(), msg, wparam, lparam)
	}
	return
}


func (Scintilla *GScintilla) SetProperty(AKey,AValue string){
	if AKey != ""{
		btkey := ([]byte)(AKey)
		btkey = append(btkey,0)
		if AValue != ""{
			btv := ([]byte)(AValue)
			btv = append(btv,0)
			Scintilla.SendEditor(SCI_SETPROPERTY, int(uintptr(unsafe.Pointer(&btkey[0]))), int(uintptr(unsafe.Pointer(&btv[0]))))
		}else{
			Scintilla.SendEditor(SCI_SETPROPERTY, int(uintptr(unsafe.Pointer(&btkey[0]))), int(uintptr(unsafe.Pointer(&cDSciNull))))
		}
	}
}

func (Scintilla *GScintilla)SetLexer(lexer ILanguageLexer)  {
	if Scintilla.fLexer == lexer{
		return
	}
	if Scintilla.fLexer != nil{
		Scintilla.fLexer.SetEditor(nil)
	}
	if lexer == nil{
		Scintilla.StyleClearAll()
	}
	Scintilla.fLexer = lexer
	if lexer!=nil{
		lexer.SetEditor(Scintilla)
		lexer.Update()
	}
}

func NewScintillaEditor(aParent Components.IWincontrol) *GScintilla {
	if scintillaHandle.fHandle == 0{
		return nil
	}
	pType := reflect.TypeOf(aParent)
	hasWincontrol := false
	if pType.Kind() == reflect.Ptr {
		_, hasWincontrol = pType.Elem().FieldByName("GWinControl")
	}
	if hasWincontrol {
		Scintilla := new(GScintilla)
		Scintilla.SubInit()
		Scintilla.SetParent(aParent)
		Scintilla.SetWidth(200)
		Scintilla.SetVisible(true)
		Scintilla.SetHeight(300)
		Scintilla.SelectBack = Graphics.ClHighlight
		Scintilla.SelectFore = Graphics.ClHighlightText
		Scintilla.LineSpaceBelow = 1
		Scintilla.TabWidth = 8
		Scintilla.CodeLines = new(GCodeLines)
		Scintilla.defFont = NewLexerFont(nil,0,0)
		Scintilla.MarginBand = newMarginBand(Scintilla)
		Scintilla.defFont.fEditor = Scintilla
		return Scintilla
	}
	return nil
}

/*func freeScintillaHandle(){
	syscall.FreeLibrary(scintillaHandle.fHandle)
}*/

//初始化这个动态库
func init()  {
	scintillaHandle = new(scintillaLib)
	if scintillaHandle.fHandle,_ = syscall.LoadLibrary("SciLexer.dll");scintillaHandle.fHandle!=0{
		scintillaHandle.fDirectFunction, _ = syscall.GetProcAddress(scintillaHandle.fHandle, "Scintilla_DirectFunction")
		//runtime.SetFinalizer(scintillaHandle.ref,freeScintillaHandle)
	}
}