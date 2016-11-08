package WinApi

import (
	"syscall"
	"unsafe"
)

type GNotifyIconData struct {
	CbSize  uint32
	WND  syscall.Handle
	UID   uint32
	UFlags uint32
	UCallbackMessage uint32
	HIcon  HICON
	SZTip	[128]uint16
	DWState uint32
	DWStateMask uint32
	SZInfo  [256]uint16
	UTimeout uint32
	SZInfoTitle  [64]uint16
	DWInfoFlags  uint32
	GUIDItem  syscall.GUID
	HBalloonIcon HICON
}

const(
	NIM_ADD         = 0X00000000
	NIM_MODIFY      = 0X00000001
	NIM_DELETE      = 0X00000002
	NIM_SETFOCUS    = 0X00000003
	NIM_SETVERSION  = 0X00000004
	NOTIFYICON_VERSION = 3
	NOTIFYICON_VERSION_4 = 4	
	NIF_MESSAGE     = 0X00000001
	NIF_ICON        = 0X00000002
	NIF_TIP         = 0X00000004
	NIF_STATE       = 0X00000008
	NIF_INFO        = 0X00000010
	NIF_GUID = 0X00000020
	NIF_REALTIME = 0X00000040
	NIF_SHOWTIP = 0X00000080
	NIS_HIDDEN = 0X00000001
	NIS_SHAREDICON = 0X00000002	
	NIIF_NONE       = 0X00000000
	NIIF_INFO       = 0X00000001
	NIIF_WARNING    = 0X00000002
	NIIF_ERROR      = 0X00000003
	NIIF_USER       = 0X00000004
	NIIF_ICON_MASK  = 0X0000000F
	NIIF_NOSOUND    = 0X00000010
	NIIF_LARGE_ICON = 0X00000020
	NIIF_RESPECT_QUIET_TIME = 0X00000080

	NIN_SELECT      = 0X0400
	NINF_KEY        =  0X1
	NIN_KEYSELECT   = NIN_SELECT | NINF_KEY

	NIN_BALLOONSHOW       = 0X0400 + 2
	NIN_BALLOONHIDE       = 0X0400 + 3
	NIN_BALLOONTIMEOUT    = 0X0400 + 4
	NIN_BALLOONUSERCLICK  = 0X0400 + 5
	NIN_POPUPOPEN         = 0X0400 + 6
	NIN_POPUPCLOSE        = 0X0400 + 7
	WM_SYSTEM_TRAY_MESSAGE = WM_USER + 1
)

var(
	fnShell_NotifyIcon uintptr
	libshell32 syscall.Handle
	fnShellExecuteW uintptr
)

func init()  {
	libshell32, _ = syscall.LoadLibrary("shell32.dll")
	fnShell_NotifyIcon, _ = syscall.GetProcAddress(libshell32, "Shell_NotifyIconW")
	fnShellExecuteW,_ = syscall.GetProcAddress(libshell32, "ShellExecuteW")
}

func Shell_NotifyIcon(dwMessage uint32, lpData *GNotifyIconData)bool{
	ret,_,_ := syscall.Syscall(fnShell_NotifyIcon,2,uintptr(dwMessage),uintptr(unsafe.Pointer(lpData)),0)
	return ret != 0
}

func ShellExecute(hwnd syscall.Handle,Operation, FileName, Parameters,Directory uintptr,ShowCmd int)HINST  {
	ret,_,_ := syscall.Syscall6(fnShellExecuteW,6,uintptr(hwnd),Operation, FileName, Parameters,Directory,uintptr(ShowCmd))
	return HINST(ret)
}