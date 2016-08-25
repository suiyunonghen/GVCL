package WinApi

import (
	"syscall"
)

var (
	libversion                syscall.Handle
	fnGetFileVersionInfoW     uintptr
	fnGetFileVersionInfoA     uintptr
	fnGetFileVersionInfoSizeW uintptr
	fnGetFileVersionInfoSizeA uintptr
	fnVerFindFileW            uintptr
	fnVerFindFileA            uintptr
	fnVerInstallFileW         uintptr
	fnVerInstallFileA         uintptr
	fnVerQueryValueW          uintptr
	fnVerQueryValueA          uintptr
)

func init() {
	libversion, _ = syscall.LoadLibrary("version.dll")
	fnGetFileVersionInfoW, _ = syscall.GetProcAddress(libversion, "GetFileVersionInfoW")
	fnGetFileVersionInfoA, _ = syscall.GetProcAddress(libversion, "GetFileVersionInfoA")
	fnGetFileVersionInfoSizeW, _ = syscall.GetProcAddress(libversion, "GetFileVersionInfoSizeW")
	fnGetFileVersionInfoSizeA, _ = syscall.GetProcAddress(libversion, "GetFileVersionInfoSizeA")
	fnVerFindFileW, _ = syscall.GetProcAddress(libversion, "VerFindFileW")
	fnVerFindFileA, _ = syscall.GetProcAddress(libversion, "VerFindFileA")
	fnVerInstallFileW, _ = syscall.GetProcAddress(libversion, "VerInstallFileW")
	fnVerInstallFileA, _ = syscall.GetProcAddress(libversion, "VerInstallFileA")
	fnVerQueryValueW, _ = syscall.GetProcAddress(libversion, "VerQueryValueW")
	fnVerQueryValueA, _ = syscall.GetProcAddress(libversion, "VerQueryValueA")
}
