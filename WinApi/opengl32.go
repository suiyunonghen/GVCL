package WinApi

import (
	"syscall"
)

var (
	libopengl32                 syscall.Handle
	fnwglCopyContext            uintptr
	fnwglCreateContext          uintptr
	fnwglCreateLayerContext     uintptr
	fnwglDeleteContext          uintptr
	fnwglDescribeLayerPlane     uintptr
	fnwglGetCurrentContext      uintptr
	fnwglGetCurrentDC           uintptr
	fnwglGetLayerPaletteEntries uintptr
	fnwglMakeCurrent            uintptr
	fnwglRealizeLayerPalette    uintptr
	fnwglSetLayerPaletteEntries uintptr
	fnwglShareLists             uintptr
	fnwglSwapLayerBuffers       uintptr
	fnwglSwapMultipleBuffers    uintptr
	fnwglUseFontBitmapsW        uintptr
	fnwglUseFontOutlinesW       uintptr
	fnwglUseFontBitmapsA        uintptr
	fnwglUseFontOutlinesA       uintptr
)

func init() {
	libopengl32, _ = syscall.LoadLibrary("opengl32.dll")
	fnwglCopyContext, _ = syscall.GetProcAddress(libopengl32, "wglCopyContext")
	fnwglCreateContext, _ = syscall.GetProcAddress(libopengl32, "wglCreateContext")
	fnwglCreateLayerContext, _ = syscall.GetProcAddress(libopengl32, "wglCreateLayerContext")
	fnwglDeleteContext, _ = syscall.GetProcAddress(libopengl32, "wglDeleteContext")
	fnwglDescribeLayerPlane, _ = syscall.GetProcAddress(libopengl32, "wglDescribeLayerPlane")
	fnwglGetCurrentContext, _ = syscall.GetProcAddress(libopengl32, "wglGetCurrentContext")
	fnwglGetCurrentDC, _ = syscall.GetProcAddress(libopengl32, "wglGetCurrentDC")
	fnwglGetLayerPaletteEntries, _ = syscall.GetProcAddress(libopengl32, "wglGetLayerPaletteEntries")
	fnwglMakeCurrent, _ = syscall.GetProcAddress(libopengl32, "wglMakeCurrent")
	fnwglRealizeLayerPalette, _ = syscall.GetProcAddress(libopengl32, "wglRealizeLayerPalette")
	fnwglSetLayerPaletteEntries, _ = syscall.GetProcAddress(libopengl32, "wglSetLayerPaletteEntries")
	fnwglShareLists, _ = syscall.GetProcAddress(libopengl32, "wglShareLists")
	fnwglSwapLayerBuffers, _ = syscall.GetProcAddress(libopengl32, "wglSwapLayerBuffers")
	fnwglSwapMultipleBuffers, _ = syscall.GetProcAddress(libopengl32, "wglSwapMultipleBuffers")
	fnwglUseFontBitmapsW, _ = syscall.GetProcAddress(libopengl32, "wglUseFontBitmapsW")
	fnwglUseFontOutlinesW, _ = syscall.GetProcAddress(libopengl32, "wglUseFontOutlinesW")
	fnwglUseFontBitmapsA, _ = syscall.GetProcAddress(libopengl32, "wglUseFontBitmapsA")
	fnwglUseFontOutlinesA, _ = syscall.GetProcAddress(libopengl32, "wglUseFontOutlinesA")
}
