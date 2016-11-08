package NVisbleControls

import (
	"suiyunonghen/GVCL/WinApi"
	"suiyunonghen/GVCL/Graphics"
	"suiyunonghen/GVCL/Components"
	"fmt"
	"syscall"
	"unsafe"
)

type IMenu interface {
	IsPopupMenu()bool
	MenuHandle()WinApi.HMENU
}

type GPopList struct {
	WindowHandle syscall.Handle
	usedIds []*GMenuItem
}

func (popList *GPopList)MenuItemById(Id uint16)*GMenuItem  {
	if popList.usedIds != nil && Id>0 && int(Id)<=len(popList.usedIds){
		return popList.usedIds[Id - 1]
	}
	return nil
}

func (popList *GPopList)setItemMenuId(item *GMenuItem)  {
	if popList.usedIds == nil{
		popList.usedIds = make([]*GMenuItem,256)
		popList.usedIds[0] = item
		item.fMenuId = 1
		return
	}
	totallen := len(popList.usedIds)
	for i:=0;i<totallen;i++{
		if popList.usedIds[i]==nil {
			popList.usedIds[i] = item
			item.fMenuId = uint16(i +1)
			return
		}
	}
	popList.usedIds = append(popList.usedIds,item)
	item.fMenuId = uint16(len(popList.usedIds))
	return
}

func (popList *GPopList)reciveMenuId(id uint16)  {
	if popList.usedIds!=nil && id>0 && id <= uint16(len(popList.usedIds)){
		popList.usedIds[id-1] = nil
	}
}
var(
	PopList *GPopList
)

type GMenu struct {
	Components.GComponent
	fItems *GMenuItem
}

func (menu *GMenu)Destroy(){
	if menu.fItems != nil{
		menu.fItems.Destroy()
	}
}

func (menu *GMenu)CreateMenu()  {
	if menu.fItems==nil{
		menu.fItems = new(GMenuItem)
		menu.fItems.SubInit()
		var targeobj interface{}
		if i := menu.SubChildCount() -1;i>=0{
			targeobj = menu.SubChild(i)
		}else{
			targeobj = menu
		}
		menu.fItems.fOwnerMenu = targeobj
	}
	if menu.fItems.GetHandle() == 0{
		panic(fmt.Sprintf("创建菜单失败，错误代码：%d",WinApi.GetLastError()))
	}
}

func (menu *GMenu)MenuHandle()WinApi.HMENU  {
	return menu.fItems.GetHandle()
}

func (menu *GMenu)IsPopupMenu()bool  {
	var targeobj interface{}
	if i := menu.SubChildCount() -1;i>=0{
		targeobj = menu.SubChild(i)
		return targeobj.(IMenu).IsPopupMenu()
	}
	return false
}

func (menu *GMenu)SubInit()  {
	menu.GObject.SubInit(menu)
}

type GPopupMenu struct {
	GMenu
}

func (menu *GPopupMenu)IsPopupMenu()bool  {
	return true
}

func (menu *GPopupMenu)SubInit()  {
	menu.GMenu.SubInit()
	menu.GObject.SubInit(menu)
}

func (menu *GPopupMenu)PopUp(x,y int32)  {
	MenuHandle := menu.MenuHandle()
	if MenuHandle!=0{
		WinApi.TrackPopupMenu(MenuHandle,WinApi.TPM_LEFTALIGN,int(x),int(y),0,PopList.WindowHandle,nil)
	}
}

func (menu *GPopupMenu)Items()*GMenuItem  {
	if menu.fItems == nil{
		menu.fItems = new(GMenuItem)
		menu.fItems.SubInit()
		var targeobj interface{}
		if i := menu.SubChildCount() -1;i>=0{
			targeobj = menu.SubChild(i)
		}else{
			targeobj = menu
		}
		menu.fItems.fOwnerMenu = targeobj
	}
	return menu.fItems
}

//菜单项目
type GMenuItem struct {
	Graphics.GObject
	menuHandle WinApi.HMENU
	fownerItem *GMenuItem
	fCaption string
	fEnabled bool
	fDefault bool
	fOwnerMenu interface{}
	fitemList []*GMenuItem
	fIndex int32
	fMenuId uint16
	fChecked bool
	OnClick Graphics.NotifyEvent
}

func (item *GMenuItem)Click()  {
	if item.OnClick!=nil{
		item.OnClick(item)
	}
}

func (item *GMenuItem)Caption() string {
	return item.fCaption
}

func NewMenuItem(menucaption string)*GMenuItem  {
	menuitem := new(GMenuItem)
	menuitem.SubInit()
	menuitem.fCaption = menucaption
	return menuitem
}

func (citem *GMenuItem)appendToParentItem(parentItem *GMenuItem)  {
	PopList.setItemMenuId(citem)
	itemInfo := new(WinApi.GMenuItemInfo)
	itemInfo.CbSize = uint32(unsafe.Sizeof(*itemInfo))
	itemInfo.FMask = WinApi.MIIM_CHECKMARKS | WinApi.MIIM_DATA | WinApi.MIIM_ID |
		WinApi.MIIM_STATE | WinApi.MIIM_SUBMENU | WinApi.MIIM_TYPE
	itemInfo.WID = uint32(citem.fMenuId)
	var fstate uint32
	if citem.fChecked {
		fstate= WinApi.MFS_CHECKED
	}else{
		fstate= WinApi.MFS_UNCHECKED
	}
	if citem.fEnabled{
		fstate = fstate | WinApi.MFS_ENABLED
	}else{
		fstate = fstate | WinApi.MFS_DISABLED
	}
	if citem.fDefault{
		fstate = fstate | WinApi.MFS_DEFAULT
	}
	itemInfo.FState = fstate
	if citem.fCaption == "-"{
		fstate = WinApi.MFT_SEPARATOR
	}else{
		fstate =WinApi.MFT_STRING
	}
	itemInfo.DwTypeData = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(citem.fCaption)))
	itemInfo.FType = fstate
	if citem.fitemList !=nil && len(citem.fitemList)>0 {
		itemInfo.HSubMenu = citem.GetHandle()
		if itemInfo.HSubMenu == 0{
			panic(fmt.Sprintf("创建菜单失败，错误代码：%d",WinApi.GetLastError()))
		}
	}else{
		itemInfo.HSubMenu = 0
	}
	WinApi.InsertMenuItem(parentItem.menuHandle,0xFFFFFFFF,true,itemInfo)
}

func (item *GMenuItem)AddItem(caption string)(menuitem *GMenuItem)  {
	menuitem = new(GMenuItem)
	if item.fitemList == nil{
		item.fitemList = make([]*GMenuItem,0)
	}
	menuitem.SubInit()
	menuitem.fIndex = int32(len(item.fitemList))
	menuitem.fOwnerMenu= nil
	menuitem.fCaption = caption
	menuitem.fownerItem = item
	item.fitemList = append(item.fitemList,menuitem)
	if item.menuHandle != 0{
		if menuitem.fMenuId == 0{
			PopList.setItemMenuId(menuitem)
		}
		menuitem.appendToParentItem(item)
	}
	return
}


func (item *GMenuItem)SetEnabled(v bool)  {
	if item.fEnabled != v{
		item.fEnabled = v
		if item.fownerItem != nil && item.fownerItem.menuHandle != 0{
			WinApi.EnableMenuItem(item.fownerItem.menuHandle,int(item.fIndex),v)
		}
	}
}


func (item *GMenuItem)prepareItems()  {
	if item.fitemList == nil {
		return
	}
	itemInfo := new(WinApi.GMenuItemInfo)
	itemInfo.CbSize = uint32(unsafe.Sizeof(*itemInfo))
	itemInfo.FMask = WinApi.MIIM_CHECKMARKS | WinApi.MIIM_DATA | WinApi.MIIM_ID |
		WinApi.MIIM_STATE | WinApi.MIIM_SUBMENU | WinApi.MIIM_TYPE
	var fstate uint32
	if itemlen := len(item.fitemList); itemlen >0{
		for i := 0;i<itemlen;i++{
			subitem := item.fitemList[i]
			if subitem.fMenuId == 0{
				PopList.setItemMenuId(subitem)
			}
			itemInfo.WID = uint32(subitem.fMenuId)
			if subitem.fChecked {
				fstate= WinApi.MFS_CHECKED
			}else{
				fstate= WinApi.MFS_UNCHECKED
			}
			if subitem.fEnabled{
				fstate = fstate | WinApi.MFS_ENABLED
			}else{
				fstate = fstate | WinApi.MFS_DISABLED
			}
			if subitem.fDefault{
				fstate = fstate | WinApi.MFS_DEFAULT
			}
			itemInfo.FState = fstate
			if subitem.fCaption == "-"{
				fstate = WinApi.MFT_SEPARATOR
			}else{
				fstate =WinApi.MFT_STRING
			}
			itemInfo.DwTypeData = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subitem.fCaption)))
			itemInfo.FType = fstate
			if subitem.fitemList !=nil && len(subitem.fitemList)>0 {
				itemInfo.HSubMenu = subitem.GetHandle()
				if itemInfo.HSubMenu == 0{
					panic(fmt.Sprintf("创建菜单失败，错误代码：%d",WinApi.GetLastError()))
				}
			}else{
				itemInfo.HSubMenu = 0
			}
			WinApi.InsertMenuItem(item.menuHandle,0xFFFFFFFF,true,itemInfo)
		}
	}
}

func (item *GMenuItem)SubInit()   {
	item.GObject.SubInit(item)
	item.fEnabled = true
	item.fChecked = false
	item.fDefault = false
}

func (item *GMenuItem)Destroy()  {
	if item.fitemList != nil{ //回收菜单ID
		for _,v := range item.fitemList {
			PopList.reciveMenuId(v.fMenuId)
		}
	}
	if item.menuHandle != 0{
		WinApi.DestroyMenu(item.menuHandle)
		item.menuHandle = 0
	}
}

func (item *GMenuItem)GetHandle()WinApi.HMENU  {
	if item.menuHandle== 0{
		if item.fOwnerMenu != nil && item.fOwnerMenu.(IMenu).IsPopupMenu(){
			item.menuHandle = WinApi.CreatePopupMenu()
		}else{
			item.menuHandle = WinApi.CreateMenu()
		}
		//填充菜单
		item.prepareItems()
	}
	return item.menuHandle
}

func NewPopupMenu(Owner interface{})*GPopupMenu  {
	popmenu := new(GPopupMenu)
	popmenu.SubInit()
	popmenu.SetOwner(Owner)
	return popmenu
}
