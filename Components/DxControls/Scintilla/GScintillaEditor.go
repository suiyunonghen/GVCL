/*
Scintilla编辑器控件的封装
Autor: 不得闲
QQ:75492895
 */
package Scintilla

import (
	"github.com/suiyunonghen/GVCL/Components"
	"github.com/suiyunonghen/GVCL/WinApi"
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"reflect"
	"syscall"
	"github.com/suiyunonghen/GVCL/Graphics"
)



type (
	scintillaLib struct {
		fHandle			syscall.Handle
		fDirectFunction uintptr			//导出的窗口函数
	}
	GCodeLines struct {
		fCodeEditor *GScintilla		//所属的代码编辑器
	}
	GScintilla struct {
		controls.GWinControl
		fDirectFunction 			uintptr
		fDirectPointer				uintptr
		CodeLines					GCodeLines
		defFont						*GDxLexerFont
	}
)

var(
	scintillaHandle		*scintillaLib
)

func (lines *GCodeLines)Clear()  {
	if lines.fCodeEditor != nil{
		lines.fCodeEditor.SendEditor(SCI_CLEARALL,0,0)
		/*if lines.fCodeEditor.FLanguageLexer != nil{
			lines.fCodeEditor.FLanguageLexer.ClearApiItems(False);
		}*/
	}
}


func (Scintilla *GScintilla) SubInit() {
	Scintilla.GWinControl.SubInit()
	Scintilla.GComponent.SubInit(Scintilla)
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

func (Scintilla *GScintilla)SendEditor(AMessage,WParam,LParam int)  {
	if !Scintilla.HandleAllocated(){
		return
	}
	if Scintilla.fDirectFunction != 0{
		syscall.Syscall6(Scintilla.fDirectFunction,4,Scintilla.fDirectPointer,uintptr(AMessage),uintptr(WParam),uintptr(LParam),0,0)
	}else{
		WinApi.SendMessage(Scintilla.GetWindowHandle(),uint(AMessage),uintptr(WParam),uintptr(LParam))
	}

}


func (Scintilla *GScintilla) initScintilla()  {
	//设定Utf8编码
	Scintilla.SendEditor(SCI_SETKEYSUNICODE, 1, 0)
	Scintilla.SendEditor(SCI_SETCODEPAGE, SC_CP_UTF8, 0)

	Scintilla.SendEditor(SCI_STYLESETBACK,32,int(Scintilla.GetColor())); //设置编辑区窗口背景色
	Scintilla.defFont.InitLexFont()
}

func (Scintilla *GScintilla) CreateWindowHandle(params *Components.GCreateParams)bool{
   if Scintilla.GWinControl.CreateWindowHandle(params){
	   h := Scintilla.GetWindowHandle()
	   //typedef sptr_t (*SciFnDirect)(sptr_t ptr, unsigned int iMessage, uptr_t wParam, sptr_t lParam);
	   //Scintilla.fDirectPointer = uintptr(WinApi.SendMessage(h,SCI_GETDIRECTFUNCTION, 0, 0))
	   //这个用的Cdecl方式，调用有问题，所以用导出函数Scintilla_DirectFunction
	   Scintilla.fDirectFunction = scintillaHandle.fDirectFunction
	   Scintilla.fDirectPointer = uintptr(WinApi.SendMessage(h,SCI_GETDIRECTPOINTER, 0, 0))
	   Scintilla.initScintilla()
	   return true
   }
	return false
}
func (Scintilla *GScintilla) WndProc(msg uint32, wparam, lparam uintptr) (result uintptr, msgDispatchNext bool) {
	result = 0
	msgDispatchNext = false
	switch msg {
	case WinApi.WM_CREATE: //释放的时候触发
	default:
		result = WinApi.CallWindowProc(Scintilla.FDefWndProc, Scintilla.GetWindowHandle(), msg, wparam, lparam)
	}
	return
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
		Scintilla.defFont = NewLexerFont(nil,0,0)
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