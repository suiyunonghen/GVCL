package NVisbleControls

import (
	"suiyunonghen/GVCL/WinApi"
	"suiyunonghen/GVCL/Components"
	"unsafe"
	"syscall"
	"suiyunonghen/GVCL/Graphics"
)

type GTrayIcon struct {
	Components.GComponent
	fvisible bool
	fNotifyData WinApi.GNotifyIconData
	OnDblClick Graphics.NotifyEvent
	OnClick Graphics.NotifyEvent
	PopupMenu  *GPopupMenu
}

type GTrayIconList struct {
	trayiconWnd syscall.Handle
	trayIcons   []*GTrayIcon
}

var(
	TrayIcons  *GTrayIconList
)

func (iconlist *GTrayIconList)Destroy()  {
	if iconlist.trayIcons!=nil{
		for _,icon := range iconlist.trayIcons{
			if icon != nil{
				icon.Destroy()
			}
		}
	}
}
func (iconList *GTrayIconList)WndProc(msg uint32,IconId uint32)uintptr  {
	icon := iconList.trayIcons[IconId-1]
	switch msg {
	case WinApi.WM_MOUSEMOVE:
	case WinApi.WM_LBUTTONDOWN:
	case WinApi.WM_LBUTTONUP:
		if icon.OnClick != nil{
			icon.OnClick(icon)
		}
	case WinApi.WM_RBUTTONDOWN:
	case WinApi.WM_RBUTTONUP:
		if icon.PopupMenu != nil{
			point := new(WinApi.POINT)
			WinApi.GetCursorPos(point)
			icon.PopupMenu.PopUp(point.X,point.Y)
		}
	case WinApi.WM_LBUTTONDBLCLK, WinApi.WM_MBUTTONDBLCLK, WinApi.WM_RBUTTONDBLCLK:
		if icon.OnDblClick != nil{
			icon.OnDblClick(icon)
		}
	}
	return 0
}

func (iconList *GTrayIconList)SetIconWndProcHandle(wnd syscall.Handle,aicon WinApi.HICON)  {
	if iconList.trayiconWnd == wnd{
		return
	}
	iconList.trayiconWnd = wnd
	if iconList.trayIcons !=nil{
		for _,icon := range iconList.trayIcons{
			if icon != nil{
				icon.fNotifyData.WND = wnd
				//设为应用程序图标
				if icon.fNotifyData.HIcon == 0{
					icon.fNotifyData.HIcon = aicon
				}
				//重新显示
				if icon.fvisible{
					WinApi.Shell_NotifyIcon(WinApi.NIM_ADD,&icon.fNotifyData)
				}
			}
		}
	}
}

func (iconList *GTrayIconList)GrowTrayIcon(icon *GTrayIcon)uint32  {
	if iconList.trayIcons==nil{
		iconList.trayIcons = make([]*GTrayIcon,1)
		iconList.trayIcons[0] = icon
		return 1
	}
	for idx,icon := range iconList.trayIcons{
		if icon == nil{
			iconList.trayIcons[idx] = icon
			return uint32(idx + 1)
		}
	}
	iconList.trayIcons = append(iconList.trayIcons,icon)
	return uint32(len(iconList.trayIcons))
}

func (iconList *GTrayIconList)ReciveTrayIcon(iconId uint32)  {
	if iconList.trayIcons != nil && iconId>0 && iconId <= uint32(len(iconList.trayIcons)){
		iconList.trayIcons[iconId - 1].fNotifyData.UID = 0
		iconList.trayIcons[iconId - 1].fNotifyData.WND = 0
		iconList.trayIcons[iconId - 1] = nil
	}
}

func (ticon *GTrayIcon)SubInit()  {
	ticon.GObject.SubInit(ticon)
	ticon.fNotifyData.CbSize = uint32(unsafe.Sizeof(ticon.fNotifyData))
	ticon.fNotifyData.UFlags = WinApi.NIF_ICON | WinApi.NIF_MESSAGE
	ticon.fNotifyData.UTimeout = 10000
	ticon.fNotifyData.UCallbackMessage = WinApi.WM_SYSTEM_TRAY_MESSAGE
	ticon.fNotifyData.UFlags = ticon.fNotifyData.UFlags | WinApi.NIF_TIP
	if TrayIcons == nil{
		TrayIcons = new(GTrayIconList)
	}
	ticon.fNotifyData.UID = TrayIcons.GrowTrayIcon(ticon)
	ticon.fNotifyData.WND = TrayIcons.trayiconWnd
}

func (ticon *GTrayIcon)Destroy()  {
	ticon.SetVisible(false)
	TrayIcons.ReciveTrayIcon(ticon.fNotifyData.UID)
}

func (ticon *GTrayIcon)SetVisible(v bool)  {
	if ticon.fvisible!=v{
		ticon.fvisible = v
		if ticon.fNotifyData.WND != 0{
			if v{
				WinApi.Shell_NotifyIcon(WinApi.NIM_ADD,&ticon.fNotifyData)
			}else{
				WinApi.Shell_NotifyIcon(WinApi.NIM_DELETE,&ticon.fNotifyData)
			}
		}
	}
}

func (ticon *GTrayIcon)SetIcon(icoHandle WinApi.HICON)  {
	if ticon.fNotifyData.HIcon != icoHandle{
		ticon.fNotifyData.HIcon = icoHandle
		if ticon.fvisible && ticon.fNotifyData.HIcon != 0{
			WinApi.Shell_NotifyIcon(WinApi.NIM_MODIFY,&ticon.fNotifyData)
		}
	}
}

func NewTrayIcon(owner interface{})*GTrayIcon  {
	ret := new(GTrayIcon)
	ret.SubInit()
	ret.SetOwner(owner)
	return ret
}