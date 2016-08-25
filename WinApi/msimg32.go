package WinApi

import (
	"syscall"
)

var (
	libmsimg32       syscall.Handle
	fnAlphaBlend     uintptr
	fnAlphaDIBBlend  uintptr
	fnGradientFill   uintptr
	fnTransparentBlt uintptr
)

func init() {
	libmsimg32, _ = syscall.LoadLibrary("msimg32.dll")
	fnAlphaBlend, _ = syscall.GetProcAddress(libmsimg32, "AlphaBlend")
	fnAlphaDIBBlend, _ = syscall.GetProcAddress(libmsimg32, "AlphaDIBBlend")
	fnGradientFill, _ = syscall.GetProcAddress(libmsimg32, "GradientFill")
	fnTransparentBlt, _ = syscall.GetProcAddress(libmsimg32, "TransparentBlt")
}
