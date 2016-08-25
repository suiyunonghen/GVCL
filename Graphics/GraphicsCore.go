package Graphics

import (
	"unsafe"
)

type GColorValue uint32

func (cv GColorValue) R() byte {
	return (*GColor)(unsafe.Pointer(&cv)).R
}

func (cv GColorValue) G() byte {
	return (*GColor)(unsafe.Pointer(&cv)).G
}

func (cv GColorValue) B() byte {
	return (*GColor)(unsafe.Pointer(&cv)).B
}

func (cv GColorValue) A() byte {
	return (*GColor)(unsafe.Pointer(&cv)).A
}

type GColor struct {
	R byte
	G byte
	B byte
	A byte
}

const (
	ClRed   GColorValue = 0xFF
	ClWhite GColorValue = 0xFFFFFF
	ClBlack GColorValue = 0xFF0000
)

func RGB(R, G, B byte) (ret GColorValue) {
	ret = 0
	gcolor := Uint32ToGColor(uint32(ret))
	gcolor.R = R
	gcolor.G = G
	gcolor.B = B
	return
}

func (c *GColor) GColor2Uint32() uint32 {
	return *(*uint32)(unsafe.Pointer(c))
}

func (c *GColor) GColor2ColorValue() GColorValue {
	return GColorValue(*(*uint32)(unsafe.Pointer(c)))
}

func Uint32ToGColor(clValue uint32) *GColor {
	return (*GColor)(unsafe.Pointer(&clValue))
}
