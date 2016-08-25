package WinApi

import (
	"syscall"
)

var (
	libwtsapi32                        syscall.Handle
	fnWTSRegisterSessionNotification   uintptr
	fnWTSUnRegisterSessionNotification uintptr
)

func init() {
	libwtsapi32, _ = syscall.LoadLibrary("wtsapi32.dll")
	if libwtsapi32 == 0 {
		return
	}
	fnWTSRegisterSessionNotification, _ = syscall.GetProcAddress(libwtsapi32, "WTSRegisterSessionNotification")
	fnWTSUnRegisterSessionNotification, _ = syscall.GetProcAddress(libwtsapi32, "WTSUnRegisterSessionNotification")
}
