package main

import (
	_ "DxSoft/GVCL/Components"
	"DxSoft/GVCL/Components/Controls"
	_ "DxSoft/GVCL/Graphics"
	_"DxSoft/GVCL/WinApi"
	_ "fmt"
	_ "reflect"
	_"time"
)

func main() {
	app := new(controls.WApplication)
	app.ShowMainForm = true
	m := app.CreateForm()
	m.SetLeft(200)
	m.SetTop(50)
	m.SetCaption("测试窗体")

	b := controls.NewButton(m)
	b.SetDefault(true)
	b.SetCaption("创建窗体")
	b.OnClick = func(sender interface{}) {
		tmpm := app.CreateForm()
		tmpm.SetLeft(300)
		tmpm.SetTop(200)
		tmpm.SetCaption("阿斯顿发送")
		tmpm.ShowModal()
	}

	b1 := controls.NewButton(m)
	b1.SetCaption("关闭")
	b1.SetLeft(100)
	b1.SetTop(40)
	b1.OnClick = func(sender interface{}) {
		//m.Close()
		b.SetVisible(!b.Visible())
		b.SetCaption("测试")
	}

	app.Run()
}
