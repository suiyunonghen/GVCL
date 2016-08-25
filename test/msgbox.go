package main

import (
	_ "DxSoft/GVCL/Components"
	"DxSoft/GVCL/Components/Controls"
	_ "DxSoft/GVCL/Graphics"
	_ "DxSoft/GVCL/WinApi"
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

	b1 := controls.NewButton(m)
	b1.SetCaption("取消")
	b1.SetLeft(100)
	b1.SetTop(40)

	app.Run()
	/*cv := Graphics.ClRed
	fmt.Println(reflect.ValueOf(cv).Type())
	WinApi.MessageBox(0, "测试不得闲 ", "测试不得闲", 16)*/
}
