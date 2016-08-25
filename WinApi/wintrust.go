package WinApi

import (
	"syscall"
)

var (
	libwintrust            syscall.Handle
	fnWinLoadTrustProvider uintptr
	fnWinSubmitCertificate uintptr
	fnWinVerifyTrust       uintptr
)

func init() {
	libwintrust, _ = syscall.LoadLibrary("wintrust.dll")
	fnWinLoadTrustProvider, _ = syscall.GetProcAddress(libwintrust, "WinLoadTrustProvider")
	fnWinSubmitCertificate, _ = syscall.GetProcAddress(libwintrust, "WinSubmitCertificate")
	fnWinVerifyTrust, _ = syscall.GetProcAddress(libwintrust, "WinVerifyTrust")
}
