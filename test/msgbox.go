package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/suiyunonghen/GVCL/Components"
	"github.com/suiyunonghen/GVCL/Components/Controls"
	"github.com/suiyunonghen/GVCL/Components/DxControls/Scintilla"
	"github.com/suiyunonghen/GVCL/Components/NVisbleControls"
	"github.com/suiyunonghen/GVCL/Graphics"
	_ "github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/GVCL/WinApi"
	_ "github.com/suiyunonghen/GVCL/WinApi"
	_ "reflect"
	_ "time"
	_"github.com/suiyunonghen/DxCommonLib"
	"github.com/suiyunonghen/DxCommonLib"
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
	frm.Edit1.CodeLines.LoadFromFile("J:\\GoLibrary\\src\\github.com\\suiyunonghen\\GVCL\\Components\\componentCore.go")

	frm.OnClose = func(sender interface{}, closeAction *int8) {
		*closeAction = controls.CAFree
	}
	frm.OnResize = func(sender interface{}) {
		frm.Button1.SetLeft(frm.Width() - 90)
		frm.Button1.SetTop(frm.Height() - 80)
		frm.Edit1.SetWidth(frm.Width())
		frm.Edit1.SetHeight(frm.Height()-80)
	}
	frm.Button1.OnClick = func(sender interface{}) {
		str := frm.Edit1.CodeLines.Strings(0)
		WinApi.MessageBox(frm.GetWindowHandle(),"测试",str , 64)
		frm.Edit1.CodeLines.SetStrings(0,"不得闲测试")
	}
	return frm
}

func main() {
	/*lst := DxCommonLib.GStringList{}
	lst.LineBreak = "\r\n"
	lst.LoadFromFile("I:\\平面类资料\\test.txt")
	lst.Add("不得闲测试内容")
	lst.Add("不得闲测试内容")
	lst.Add("不得闲测试内容")
	lst.Add("不得闲测试内容")
	lst.Add("不得闲测试内容啊手动阀手动阀")
	lst.Insert(2,"插入的内容")
	fmt.Println(lst.Strings(0))
	lst.SaveToFile("I:\\平面类资料\\test.txt")*/
	app := controls.NewApplication()
	m := app.CreateForm()
	m.SetLeft(200)
	m.SetTop(50)
	m.SetCaption("测试窗体")

	lbl := controls.NewLabel(m)
	lbl.SetCaption("说明 ")

	lbl.SetAutoSize(true)
	lbl.SetColor(Graphics.ClRed)
	lbl.SetTop(40)

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

	b1 := controls.NewButton(m)
	b1.SetCaption("关闭")
	b1.Font.BeginUpdate()
	b1.Font.SetSize(10)
	b1.Font.Underline = 1
	b1.Font.SetBold(true)
	b1.Font.EndUpdate()
	b1.SetLeft(100)
	b1.SetTop(40)
	b1.OnClick = func(sender interface{}) {
		cvs := new(controls.GControlCanvas)
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
		cvs.FillRect(r)
	}

	app.Run()
}
