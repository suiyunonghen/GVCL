# GVCL
Golang Windows GUI Bindings Like Delphi VCL   
Windows的GUI界面Go语言封装，目标是实现一个类似Delphi VCL版本效果的界面库   
# 使用方法  
##### `msgbox.go`

```go
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
	app := new(controls.WApplication)
	app.ShowMainForm = true
	m := app.CreateForm()
	m.SetLeft(200)
	m.SetTop(50)
	m.SetCaption("测试窗体")

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
	b1.SetLeft(100)
	b1.SetTop(40)
	b1.OnClick = func(sender interface{}) {
		b.SetVisible(!b.Visible())
		b.SetCaption("测试")
	}
	app.Run()
}
```

##### Create Manifest `test.manifest`

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
    <assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
        <assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="SomeFunkyNameHere" type="win32"/>
        <dependency>
            <dependentAssembly>
                <assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
            </dependentAssembly>
        </dependency>
    </assembly>
```

Then either compile the manifest using the [rsrc tool](https://github.com/akavel/rsrc), like this:

	go get github.com/akavel/rsrc
	rsrc -manifest test.manifest -o rsrc.syso

or rename the `test.manifest` file to `test.exe.manifest` and distribute it with the application instead.

使用Go Build 建立文件，会将这个manifest资源包含进去，以使用系统皮肤

### Font使用   

```go
	b1.Font.BeginUpdate()
	b1.Font.SetSize(10)
	b1.Font.Underline = 1
	b1.Font.EndUpdate()
```
###  增加弹出菜单组件   
菜单使用方法
```go
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
	m.PopupMenu = pop
```
目前整体组件框架已经具备雏形，要增加其他组件库按照扩展的Button和Edit以及Label增加则可，下一步做完托盘就暂时放一段落

###  增加托盘组件完成，使用方法
```go
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
	//托盘图标，结合弹出菜单
	icon := NVisbleControls.NewTrayIcon(m)
	icon.SetIcon(app.AppIcon()) //设置托盘图标
	//icon.SetIcon(WinApi.LoadIcon(controls.Hinstance,uintptr(5))) 
	icon.SetVisible(true)
	icon.PopupMenu = pop //设置托盘的右键弹出菜单
	icon.OnDblClick = func(sender interface{}) {
		if !m.Visible(){
			m.Show()
		}else{
			m.SetVisible(false)
		}
	}
```
运行效果
![运行效果](https://github.com/suiyunonghen/GVCL/blob/master/test/GVCL.png)
