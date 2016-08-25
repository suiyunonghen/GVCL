package WinApi

import (
	"syscall"
)

var (
	libmpr                              syscall.Handle
	fnMultinetGetConnectionPerformanceW uintptr
	fnMultinetGetConnectionPerformanceA uintptr
	fnWNetAddConnection2W               uintptr
	fnWNetAddConnection2A               uintptr
	fnWNetAddConnection3W               uintptr
	fnWNetAddConnection3A               uintptr
	fnWNetAddConnectionW                uintptr
	fnWNetAddConnectionA                uintptr
	fnWNetCancelConnection2W            uintptr
	fnWNetCancelConnection2A            uintptr
	fnWNetCancelConnectionW             uintptr
	fnWNetCancelConnectionA             uintptr
	fnWNetCloseEnum                     uintptr
	fnWNetConnectionDialog1W            uintptr
	fnWNetConnectionDialog1A            uintptr
	fnWNetConnectionDialog              uintptr
	fnWNetDisconnectDialog1W            uintptr
	fnWNetDisconnectDialog1A            uintptr
	fnWNetDisconnectDialog              uintptr
	fnWNetEnumResourceW                 uintptr
	fnWNetEnumResourceA                 uintptr
	fnWNetGetConnectionW                uintptr
	fnWNetGetConnectionA                uintptr
	fnWNetGetLastErrorW                 uintptr
	fnWNetGetLastErrorA                 uintptr
	fnWNetGetNetworkInformationW        uintptr
	fnWNetGetNetworkInformationA        uintptr
	fnWNetGetProviderNameW              uintptr
	fnWNetGetProviderNameA              uintptr
	fnWNetGetResourceParentW            uintptr
	fnWNetGetResourceParentA            uintptr
	fnWNetGetUniversalNameW             uintptr
	fnWNetGetUniversalNameA             uintptr
	fnWNetGetUserW                      uintptr
	fnWNetGetUserA                      uintptr
	fnWNetOpenEnumW                     uintptr
	fnWNetOpenEnumA                     uintptr
	fnWNetSetConnectionW                uintptr
	fnWNetSetConnectionA                uintptr
	fnWNetUseConnectionW                uintptr
	fnWNetUseConnectionA                uintptr
)

func init() {
	libmpr, _ = syscall.LoadLibrary("mpr.dll")
	fnMultinetGetConnectionPerformanceW, _ = syscall.GetProcAddress(libmpr, "MultinetGetConnectionPerformanceW")
	fnMultinetGetConnectionPerformanceA, _ = syscall.GetProcAddress(libmpr, "MultinetGetConnectionPerformanceA")
	fnWNetAddConnection2W, _ = syscall.GetProcAddress(libmpr, "WNetAddConnection2W")
	fnWNetAddConnection2A, _ = syscall.GetProcAddress(libmpr, "WNetAddConnection2A")
	fnWNetAddConnection3W, _ = syscall.GetProcAddress(libmpr, "WNetAddConnection3W")
	fnWNetAddConnection3A, _ = syscall.GetProcAddress(libmpr, "WNetAddConnection3A")
	fnWNetAddConnectionW, _ = syscall.GetProcAddress(libmpr, "WNetAddConnectionW")
	fnWNetAddConnectionA, _ = syscall.GetProcAddress(libmpr, "WNetAddConnectionA")
	fnWNetCancelConnection2W, _ = syscall.GetProcAddress(libmpr, "WNetCancelConnection2W")
	fnWNetCancelConnection2A, _ = syscall.GetProcAddress(libmpr, "WNetCancelConnection2A")
	fnWNetCancelConnectionW, _ = syscall.GetProcAddress(libmpr, "WNetCancelConnectionW")
	fnWNetCancelConnectionA, _ = syscall.GetProcAddress(libmpr, "WNetCancelConnectionA")
	fnWNetCloseEnum, _ = syscall.GetProcAddress(libmpr, "WNetCloseEnum")
	fnWNetConnectionDialog1W, _ = syscall.GetProcAddress(libmpr, "WNetConnectionDialog1W")
	fnWNetConnectionDialog1A, _ = syscall.GetProcAddress(libmpr, "WNetConnectionDialog1A")
	fnWNetConnectionDialog, _ = syscall.GetProcAddress(libmpr, "WNetConnectionDialog")
	fnWNetDisconnectDialog1W, _ = syscall.GetProcAddress(libmpr, "WNetDisconnectDialog1W")
	fnWNetDisconnectDialog1A, _ = syscall.GetProcAddress(libmpr, "WNetDisconnectDialog1A")
	fnWNetDisconnectDialog, _ = syscall.GetProcAddress(libmpr, "WNetDisconnectDialog")
	fnWNetEnumResourceW, _ = syscall.GetProcAddress(libmpr, "WNetEnumResourceW")
	fnWNetEnumResourceA, _ = syscall.GetProcAddress(libmpr, "WNetEnumResourceA")
	fnWNetGetConnectionW, _ = syscall.GetProcAddress(libmpr, "WNetGetConnectionW")
	fnWNetGetConnectionA, _ = syscall.GetProcAddress(libmpr, "WNetGetConnectionA")
	fnWNetGetLastErrorW, _ = syscall.GetProcAddress(libmpr, "WNetGetLastErrorW")
	fnWNetGetLastErrorA, _ = syscall.GetProcAddress(libmpr, "WNetGetLastErrorA")
	fnWNetGetNetworkInformationW, _ = syscall.GetProcAddress(libmpr, "WNetGetNetworkInformationW")
	fnWNetGetNetworkInformationA, _ = syscall.GetProcAddress(libmpr, "WNetGetNetworkInformationA")
	fnWNetGetProviderNameW, _ = syscall.GetProcAddress(libmpr, "WNetGetProviderNameW")
	fnWNetGetProviderNameA, _ = syscall.GetProcAddress(libmpr, "WNetGetProviderNameA")
	fnWNetGetResourceParentW, _ = syscall.GetProcAddress(libmpr, "WNetGetResourceParentW")
	fnWNetGetResourceParentA, _ = syscall.GetProcAddress(libmpr, "WNetGetResourceParentA")
	fnWNetGetUniversalNameW, _ = syscall.GetProcAddress(libmpr, "WNetGetUniversalNameW")
	fnWNetGetUniversalNameA, _ = syscall.GetProcAddress(libmpr, "WNetGetUniversalNameA")
	fnWNetGetUserW, _ = syscall.GetProcAddress(libmpr, "WNetGetUserW")
	fnWNetGetUserA, _ = syscall.GetProcAddress(libmpr, "WNetGetUserA")
	fnWNetOpenEnumW, _ = syscall.GetProcAddress(libmpr, "WNetOpenEnumW")
	fnWNetOpenEnumA, _ = syscall.GetProcAddress(libmpr, "WNetOpenEnumA")
	fnWNetSetConnectionW, _ = syscall.GetProcAddress(libmpr, "WNetSetConnectionW")
	fnWNetSetConnectionA, _ = syscall.GetProcAddress(libmpr, "WNetSetConnectionA")
	fnWNetUseConnectionW, _ = syscall.GetProcAddress(libmpr, "WNetUseConnectionW")
	fnWNetUseConnectionA, _ = syscall.GetProcAddress(libmpr, "WNetUseConnectionA")
}
