package main

import (
	_ "DxSoft/GVCL/Components"
	"DxSoft/GVCL/Components/Controls"
	_ "DxSoft/GVCL/Graphics"
	"DxSoft/GVCL/WinApi"
	_ "fmt"
	_ "reflect"
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
	b.SetCaption("确定")
	b.OnClick = func(sender interface{}) {
		WinApi.MessageBox(b.GetWindowHandle(),b.GetText(),b.GetText(),64)
	}

	b1 := controls.NewButton(m)
	b1.SetCaption("关闭")
	b1.SetLeft(100)
	b1.SetTop(40)
	b1.OnClick = func(sender interface{}) {
		m.Close()
	}

	app.Run()
}
