package main

import (
	"fmt"
	_ "fmt"
	"github.com/suiyunonghen/DxCommonLib"
	_ "github.com/suiyunonghen/DxCommonLib"
	_ "github.com/suiyunonghen/GVCL/Components"
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"github.com/suiyunonghen/GVCL/Components/DxControls/Scintilla"
	"github.com/suiyunonghen/GVCL/Components/DxControls/WindowLessControls"
	"github.com/suiyunonghen/GVCL/Components/DxControls/gminiblink"
	"github.com/suiyunonghen/GVCL/Components/NVisbleControls"
	"github.com/suiyunonghen/GVCL/Graphics"
	_ "github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/GVCL/WinApi"
	_ "github.com/suiyunonghen/GVCL/WinApi"
	_ "reflect"
	_ "time"
)

type GForm1 struct {
	controls.GForm
	Button1 *controls.GButton
	Edit1   *Scintilla.GScintilla
}

func NewForm1(app *controls.WApplication) *GForm1 {
	frm := new(GForm1)
	frm.SubInit()
	frm.Button1 = controls.NewButton(frm)
	frm.Button1.SetWidth(80)
	frm.Button1.SetHeight(30)
	frm.Button1.SetCaption("确定关闭")

	frm.Edit1 = Scintilla.NewScintillaEditor(frm)
	frm.Edit1.SetColor(Graphics.RGB(184,220,220))
	frm.Edit1.MarginBand.ShowCodeFlod = true

	pop := NVisbleControls.NewPopupMenu(frm)
	TargetBookmarkItem := pop.Items().AddItem("设置书签")
	GotoBookmarkItem := pop.Items().AddItem("跳转书签")
	for i :=0;i<10;i++{
		str := fmt.Sprintf("%s%d","书签",i)
		bookmarkitem := TargetBookmarkItem.AddItem(str)
		bookmarkitem.TagData = i
		bookmarkitem.OnClick = func(sender interface{}) {
			frm.Edit1.MarginBand.SetBookmark(sender.(*NVisbleControls.GMenuItem).TagData.(int))
		}
		gotobookitem := GotoBookmarkItem.AddItem(str)
		gotobookitem.TagData = i
		gotobookitem.OnClick = func(sender interface{}) {
			frm.Edit1.MarginBand.GotoBookmark(sender.(*NVisbleControls.GMenuItem).TagData.(int))
		}
	}
	GotoBookmarkItem = pop.Items().AddItem("清空所有书签")
	GotoBookmarkItem.OnClick = func(sender interface{}) {
		frm.Edit1.MarginBand.ClearMarks()
	}
	frm.Edit1.PopupMenu = pop

	frm.Edit1.CodeLines.LineBreak = DxCommonLib.LBK_CRLF
	frm.Edit1.CodeLines.LoadFromFile("D:\\bookmarks.html")
	//("J:\\GoLibrary\\src\\github.com\\suiyunonghen\\GVCL\\Components\\componentCore.go")

	lexer := Scintilla.NewHtmlLexer()
	frm.Edit1.SetLexer(lexer)
	frm.OnClose = func(sender interface{}, closeAction *int8) {
		*closeAction = controls.CAFree
	}
	frm.OnResize = func(sender interface{}) {
		frm.Button1.SetLeft(frm.Width() - 90)
		frm.Button1.SetTop(frm.Height() - 80)
		frm.Edit1.SetWidth(frm.Width() - 20)
		frm.Edit1.SetHeight(frm.Height()-80)
	}
	frm.Button1.OnClick = func(sender interface{}) {
		str := frm.Edit1.CodeLines.Strings(0)
		WinApi.MessageBox(frm.GetWindowHandle(),"测试",str , 64)
		frm.Edit1.CodeLines.SetStrings(0,"不得闲测试")
	}
	return frm
}

type GForm2 struct {
	controls.GForm
	Browser *gminiblink.GBlinkWebBrowser
}

func NewForm2(app *controls.WApplication) *GForm2 {
	frm := new(GForm2)
	frm.SubInit()
	frm.Browser = gminiblink.NewBlinkWebBrowser(frm,BindFunctions)
	frm.OnShow = func(sender interface{}) {
		//frm.Browser.Navigate("www.baidu.com")
		frm.Browser.LoadFromFile(0,`E:\GoLib\src\github.com\suiyunonghen\GVCL\Components\DxControls\gminiblink\test.html`)
	}
	frm.OnClose = func(sender interface{}, closAction *int8) {
		*closAction = controls.CAFree
	}
	frm.OnResize = func(sender interface{}) {
		r := frm.ClientRect()
		frm.Browser.SetWidth(r.Width())
		frm.Browser.SetHeight(r.Height())
	}
	frm.Browser.OnConsole = func(webView gminiblink.WkeWebView, level gminiblink.WkeConsoleLevel, msg, sourceName string, sourceline uint32, stackTrace string) {
		fmt.Println(level)
		fmt.Println("ConsoleMsg：",msg)
		fmt.Println("ConsolesourceName：",sourceName,"，sourceline：",sourceline)
		fmt.Println("stackTrace:",stackTrace)
	}
	frm.Browser.OnDocumentCompleted = func(sender interface{}) {
		fmt.Println("页面加载完成")
	}
	return frm
}

var(
	BindFunctions	map[string]*gminiblink.JSBindFunction
)

func main() {

	gminiblink.BlinkLib.LoadBlink(`E:\GoLib\miniblink64_20190810\miniblink64\node.dll`)
	app := controls.NewApplication()
	m := app.CreateForm()
	m.SetLeft(200)
	m.SetTop(50)
	m.SetCaption("测试窗体")

	lbl := controls.NewLabel(m)
	lbl.SetCaption("说明 ")
	lbl.SetTrasparent(true)


	//lbl.SetAutoSize(true)
	lbl.SetColor(Graphics.ClRed)
	lbl.SetTop(40)


	BindFunctions = make(map[string]*gminiblink.JSBindFunction)
	func1 := new(gminiblink.JSBindFunction)
	func1.BindHandle = func(params ...interface{}) interface{} {
		h := app.MainForm().GetWindowHandle()
		WinApi.MessageBox(h,params[0].(string),"asdf",64)
		return 0
	}
	func1.ParamCount = 1
	BindFunctions["GoMsgBox"] = func1

	func1 = new(gminiblink.JSBindFunction)
	func1.BindHandle = func(params ...interface{}) interface{} {
		sum := 0
		for _,v := range params{
			sum += v.(int)
		}
		return sum
	}
	func1.ParamCount = 0
	BindFunctions["Gosum"] = func1



	//菜单
	pop := NVisbleControls.NewPopupMenu(m)
	tmpitem := pop.Items().AddItem("测试1")
	citem := tmpitem.AddItem("子测试1")
	citem.OnClick = func(sender interface{}) {
		if AMajor, AMinor, ABuild, ok := WinApi.GetProductVersion("D:\\DevTools\\Microsoft VS Code\\Code.exe"); ok {
			st := fmt.Sprintf("%d.%d.%d", AMajor, AMinor, ABuild)
			WinApi.MessageBox(m.GetWindowHandle(), st+sender.(*NVisbleControls.GMenuItem).Caption(), "消息", 64)
		} else {
			WinApi.MessageBox(m.GetWindowHandle(), "菜单测试"+sender.(*NVisbleControls.GMenuItem).Caption(), "消息", 64)
		}
	}
	citem = pop.Items().AddItem("注册表")
	citem.OnClick = func(sender interface{}) {
		reg := NVisbleControls.NewRegistry(0)
		reg.SetRootKey(WinApi.HKEY_LOCAL_MACHINE)
		if reg.OpenKey("SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", false) {
			if reg.ValueExists("SynTPEnh") {
				WinApi.MessageBox(m.GetWindowHandle(), "SynTPEnh自动启动: "+reg.ReadString("SynTPEnh"), "消息", 64)
			}
			WinApi.MessageBox(m.GetWindowHandle(), "打开注册表测试"+sender.(*NVisbleControls.GMenuItem).Caption(), "消息", 64)
		}
		reg.Free()
	}

	//托盘图标
	icon := NVisbleControls.NewTrayIcon(m)
	icon.SetIcon(app.AppIcon())
	//icon.SetIcon(WinApi.LoadIcon(controls.Hinstance,uintptr(5)))
	icon.SetVisible(true)
	icon.PopupMenu = pop
	icon.OnDblClick = func(sender interface{}) {
		if !m.Visible() {
			m.Show()
		} else {
			m.SetVisible(false)
		}
	}

	m.PopupMenu = pop

	e := controls.NewEdit(m)
	e.SetName("Edit1")
	e.SetLeft(10)
	e.SetWidth(100)
	e.SetHeight(20)
	e.SetCaption("测试")
	b := controls.NewButton(m)
	b.SetDefault(true)
	b.SetLeft(120)
	b.SetCaption("创建窗体")
	b.OnClick = func(sender interface{}) {
		tmpm := NewForm1(app)
		tmpm.SetCaption(e.GetText())
		if tmpm.ShowModal() == controls.MrOK {
			WinApi.MessageBox(tmpm.GetWindowHandle(), "程序确定退出", "消息", 64)
		}


	}

	bBrowser  := controls.NewButton(m)
	bBrowser.SetLeft(240)
	bBrowser.SetCaption("miniBlink")
	bBrowser.OnClick = func(sender interface{}) {
		tmpm := NewForm2(app)
		tmpm.SetCaption("测试MiniBlink")
		tmpm.ShowModal()
	}

	b1 := controls.NewButton(m)
	b1.SetCaption("关闭")
	b1.Font.BeginUpdate()
	b1.Font.SetSize(10)
	b1.Font.Underline = 1
	b1.Font.SetBold(true)
	b1.Font.EndUpdate()
	b1.SetLeft(100)
	b1.SetTop(40)


	lstBox	:= controls.NewListBox(m)
	lstBox.SetTop(80)
	lstBox.SetLeft(40)
	//lstBox.SetWidth(200)
	//lstBox.SetHeight(200)
	listitems := lstBox.Items()
	listitems.Add("asfasf")
	listitems.Add("测试二恶")

	lstBox.OnItemDblClick = func(sender interface{}) {
		mb := sender.(*controls.GListBox)
		WinApi.MessageBox(mb.GetWindowHandle(),mb.Items().Strings(mb.GetItemIndex()),"消息",64)
	}

	checkbox := WindowLessControls.NewCheckBox(m)
	checkbox.SetLeft(lbl.Left()+10)
	checkbox.SetTop(lstBox.Top()+lstBox.Height()+10)
	checkbox.SetCaption("测试选择框")
	checkbox.SetChecked(true)

	b1.OnClick = func(sender interface{}) {
		/*cvs := new(controls.GControlCanvas)
		cvs.SubInit()
		cvs.SetControl(m)
		brsh := cvs.Brush()
		brsh.Color = Graphics.ClRed
		brsh.BrushStyle = Graphics.BSCross
		brsh.Change()
		r := new(WinApi.Rect)
		r.Left = 20
		r.Top = 20
		r.Right = 150
		r.Bottom = 150
		cvs.FillRect(r)*/
		lstBox.Font.BeginUpdate()
		lstBox.Font.SetName("宋体")
		lstBox.Font.SetBold(true)
		lstBox.Font.EndUpdate()

		lbl.Font.BeginUpdate()
		lbl.Font.Italic = 1
		lbl.Font.SetBold(true)
		lbl.Font.EndUpdate()
	}

	cmb := controls.NewCombobox(m,controls.CSDropDownList)
	cmb.SetTop(80)
	cmb.SetLeft(180)
	listitems = cmb.Items()
	for i := 0;i<10;i++{
		listitems.Add("asfasf")
		listitems.Add("测试二恶")
	}
	cmb.SetItemIndex(1)

	app.Run()
}
