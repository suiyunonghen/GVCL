package main

import (
	_ "DxSoft/GVCL/Components"
	"DxSoft/GVCL/Components/Controls"
	_ "DxSoft/GVCL/Graphics"
	_"DxSoft/GVCL/WinApi"
	_ "fmt"
	_ "reflect"
	_"time"
	"DxSoft/GVCL/WinApi"
	"DxSoft/GVCL/Graphics"
	"DxSoft/GVCL/Components/NVisbleControls"
)

type GForm1 struct {
	controls.GForm
	Button1 *controls.GButton
	Edit1 *controls.GEdit
}


func NewForm1(app *controls.WApplication)*GForm1{
	frm := new(GForm1)
	frm.SubInit()
	frm.Button1 = controls.NewButton(frm)
	frm.Button1.SetWidth(80)
	frm.Button1.SetHeight(30)
	frm.Button1.SetLeft(frm.Width() - 90)
	frm.Button1.SetTop(frm.Height() - 80)
	frm.Button1.SetCaption("确定关闭")

	frm.Edit1 = controls.NewEdit(frm)
	frm.OnClose = func(sender interface{},closeAction *int8) {
		*closeAction = controls.CAFree
	}
	frm.Button1.OnClick = func(sender interface{}) {
		WinApi.MessageBox(frm.GetWindowHandle(),"sadf","Asdf",64)
		frm.SetModalResult(controls.MrOK)
	}
	return frm
}

func main() {
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
		WinApi.MessageBox(m.GetWindowHandle(),"菜单测试"+sender.(*NVisbleControls.GMenuItem).Caption(),"消息",64)
	}
	citem = pop.Items().AddItem("测试2")
	citem.OnClick = func(sender interface{}) {
		WinApi.MessageBox(m.GetWindowHandle(),"菜单测试"+sender.(*NVisbleControls.GMenuItem).Caption(),"消息",64)
	}

	//托盘图标
	icon := NVisbleControls.NewTrayIcon(m)

	icon.SetIcon(app.AppIcon())
	//icon.SetIcon(WinApi.LoadIcon(controls.Hinstance,uintptr(5)))
	icon.SetVisible(true)
	icon.PopupMenu = pop
	icon.OnDblClick = func(sender interface{}) {
		if !m.Visible(){
			m.Show()
		}else{
			m.SetVisible(false)
		}
	}


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
		if tmpm.ShowModal() == controls.MrOK{
			WinApi.MessageBox(tmpm.GetWindowHandle(),"程序确定退出","消息",64)
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
		r :=  new(WinApi.Rect)
		r.Left = 20
		r.Top = 20
		r.Right = 150
		r.Bottom = 150
		cvs.FillRect(r)
	}

	app.Run()
}
