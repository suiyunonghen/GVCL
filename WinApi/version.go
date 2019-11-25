package WinApi

import (
	"fmt"
	"syscall"
	"unsafe"
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

type GVSFixedFileInfo struct{
	DWSignature uint32
	DWStrucVersion uint32
	DWFileVersionMS uint32
	DWFileVersionLS  uint32
	DWProductVersionMS uint32
	DWProductVersionLS  uint32
	DWFileFlagsMask  uint32
	DWFileFlags  uint32
	DWFileOS  uint32
	DWFileType  uint32
	DWFileSubtype  uint32
	DWFileDateMS  uint32
	DWFileDateLS  uint32
}

type GVersion struct {
	Major		uint16
	Minor		uint16
	SmallVer	uint16
	Build		uint16
}

func (v GVersion)String() string {
	return fmt.Sprintf("%d.%d.%d.%d",v.Major,v.Minor,v.SmallVer, v.Build)
}

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

func GetFileVersionInfoSize(fileName *uint16,lpdwHandle *uint32)uint32  {
	ret,_,_ :=syscall.Syscall(fnGetFileVersionInfoSizeW,2,uintptr(unsafe.Pointer(fileName)),uintptr(unsafe.Pointer(lpdwHandle)),0)
	return uint32(ret)
}

func GetFileVersionInfo(fileName *uint16,dwHandle, dwLen uint32,lpData unsafe.Pointer)bool  {
	ret,_,_ :=syscall.Syscall6(fnGetFileVersionInfoW,4,uintptr(unsafe.Pointer(fileName)),
		uintptr(dwHandle),uintptr(dwLen),uintptr(lpData),0,0)
	return ret !=0
}

func VerQueryValue(pBlock uintptr,lpSubBlock *uint16,lplpBuffer uintptr,puLen *uint32)bool  {
	ret,_,_ :=syscall.Syscall6(fnVerQueryValueW,4,pBlock,
		uintptr(unsafe.Pointer(lpSubBlock)),uintptr(unsafe.Pointer(lplpBuffer)),uintptr(unsafe.Pointer(puLen)),0,0)
	return ret !=0
}

func GetProductVersion(fileName string)(fileVersion GVersion,prodVersion GVersion,ok bool)  {
	ok = false
	fileptr := syscall.StringToUTF16Ptr(fileName)
	var wnd uint32
	if InfoSize := GetFileVersionInfoSize(fileptr,&wnd);InfoSize != 0{
		verbuffer := make([]byte,InfoSize)
		var FI *GVSFixedFileInfo
		var VerSize uint32
		if GetFileVersionInfo(fileptr,wnd,InfoSize,unsafe.Pointer(&verbuffer[0]))&&
		   VerQueryValue(uintptr(unsafe.Pointer(&verbuffer[0])),syscall.StringToUTF16Ptr("\\"),uintptr(unsafe.Pointer(&FI)),
			   &VerSize){
			fileVersion.Major = HiWord(FI.DWFileVersionMS)
			fileVersion.Minor = LoWord(FI.DWFileVersionMS)
			fileVersion.SmallVer = HiWord(FI.DWFileVersionLS)
			fileVersion.Build = LoWord(FI.DWFileVersionLS)

			prodVersion.Major = HiWord(FI.DWProductVersionMS)
			prodVersion.Minor = LoWord(FI.DWProductVersionMS)
			prodVersion.SmallVer = HiWord(FI.DWProductVersionLS)
			prodVersion.Build = LoWord(FI.DWProductVersionLS)
			ok = true
			return
		}
	}
	return
}

