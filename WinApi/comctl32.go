package WinApi

import (
	"syscall"
	"unsafe"
)

var (
	libcomctl32                    syscall.Handle
	fnInitCommonControlsEx         uintptr
	fnImageList_Create             uintptr
	fnImageList_Destroy            uintptr
	fnImageList_GetImageCount      uintptr
	fnImageList_SetImageCount      uintptr
	fnImageList_Add                uintptr
	fnImageList_ReplaceIcon        uintptr
	fnImageList_SetBkColor         uintptr
	fnImageList_GetBkColor         uintptr
	fnImageList_SetOverlayImage    uintptr
	fnImageList_Draw               uintptr
	fnImageList_Replace            uintptr
	fnImageList_AddMasked          uintptr
	fnImageList_DrawEx             uintptr
	fnImageList_DrawIndirect       uintptr
	fnImageList_Remove             uintptr
	fnImageList_GetIcon            uintptr
	fnImageList_LoadImageA         uintptr
	fnImageList_LoadImageW         uintptr
	fnImageList_Copy               uintptr
	fnImageList_BeginDrag          uintptr
	fnImageList_EndDrag            uintptr
	fnImageList_DragEnter          uintptr
	fnImageList_DragLeave          uintptr
	fnImageList_DragMove           uintptr
	fnImageList_SetDragCursorImage uintptr
	fnImageList_DragShowNolock     uintptr
	fnImageList_GetDragImage       uintptr
)

type tagINITCOMMONCONTROLSEX struct {
	dwSize uint
	dwICC  uint
}

type GInitCommonControlsEX tagINITCOMMONCONTROLSEX

func init() {
	libcomctl32, _ = syscall.LoadLibrary("comctl32.dll")
	fnInitCommonControlsEx, _ = syscall.GetProcAddress(libcomctl32, "InitCommonControlsEx")
	fnImageList_Create, _ = syscall.GetProcAddress(libcomctl32, "ImageList_Create")
	fnImageList_Destroy, _ = syscall.GetProcAddress(libcomctl32, "ImageList_Destroy")
	fnImageList_GetImageCount, _ = syscall.GetProcAddress(libcomctl32, "ImageList_GetImageCount")
	fnImageList_SetImageCount, _ = syscall.GetProcAddress(libcomctl32, "ImageList_SetImageCount")
	fnImageList_Add, _ = syscall.GetProcAddress(libcomctl32, "ImageList_Add")
	fnImageList_ReplaceIcon, _ = syscall.GetProcAddress(libcomctl32, "ImageList_ReplaceIcon")
	fnImageList_SetBkColor, _ = syscall.GetProcAddress(libcomctl32, "ImageList_SetBkColor")
	fnImageList_GetBkColor, _ = syscall.GetProcAddress(libcomctl32, "ImageList_GetBkColor")
	fnImageList_SetOverlayImage, _ = syscall.GetProcAddress(libcomctl32, "ImageList_SetOverlayImage")
	fnImageList_Draw, _ = syscall.GetProcAddress(libcomctl32, "ImageList_Draw")
	fnImageList_Replace, _ = syscall.GetProcAddress(libcomctl32, "ImageList_Replace")
	fnImageList_AddMasked, _ = syscall.GetProcAddress(libcomctl32, "ImageList_AddMasked")
	fnImageList_DrawEx, _ = syscall.GetProcAddress(libcomctl32, "ImageList_DrawEx")
	fnImageList_DrawIndirect, _ = syscall.GetProcAddress(libcomctl32, "ImageList_DrawIndirect")
	fnImageList_Remove, _ = syscall.GetProcAddress(libcomctl32, "ImageList_Remove")

	fnImageList_GetIcon, _ = syscall.GetProcAddress(libcomctl32, "ImageList_GetIcon")
	fnImageList_LoadImageA, _ = syscall.GetProcAddress(libcomctl32, "ImageList_LoadImageA")
	fnImageList_LoadImageW, _ = syscall.GetProcAddress(libcomctl32, "ImageList_LoadImageW")
	fnImageList_Copy, _ = syscall.GetProcAddress(libcomctl32, "ImageList_Copy")
	fnImageList_BeginDrag, _ = syscall.GetProcAddress(libcomctl32, "ImageList_BeginDrag")
	fnImageList_EndDrag, _ = syscall.GetProcAddress(libcomctl32, "ImageList_EndDrag")
	fnImageList_DragEnter, _ = syscall.GetProcAddress(libcomctl32, "ImageList_DragEnter")
	fnImageList_DragLeave, _ = syscall.GetProcAddress(libcomctl32, "ImageList_DragLeave")
	fnImageList_DragMove, _ = syscall.GetProcAddress(libcomctl32, "ImageList_DragMove")
	fnImageList_SetDragCursorImage, _ = syscall.GetProcAddress(libcomctl32, "ImageList_SetDragCursorImage")
	fnImageList_DragShowNolock, _ = syscall.GetProcAddress(libcomctl32, "ImageList_DragShowNolock")
	fnImageList_GetDragImage, _ = syscall.GetProcAddress(libcomctl32, "ImageList_GetDragImage")
}

func InitCommonControlsEx(ICC *GInitCommonControlsEX) bool {
	if fnInitCommonControlsEx != 0 {
		ret, _, _ := syscall.Syscall(fnInitCommonControlsEx, 1, uintptr(unsafe.Pointer(ICC)), 0, 0)
		return ret != 0
	}
	return false
}
