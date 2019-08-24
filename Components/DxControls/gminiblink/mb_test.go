package gminiblink

import (
	"fmt"
	"testing"
)

func TestMiniBlinkLib_JsArg(t *testing.T) {
	BlinkLib.LoadBlink(`E:\Delphi\Controls\UI\DxSkinctrl\miniblink-190630\node.dll`)
	fmt.Println(BlinkLib.VersionString())
	fmt.Println("start")
	wkstring := BlinkLib.WkeCreateString("测试不得闲")
	fmt.Println(BlinkLib.WkeGetString(wkstring))
	BlinkLib.WkeSetStringW(wkstring,"高价收电费")
	fmt.Println(BlinkLib.WkeGetString(wkstring))
	BlinkLib.WkeDeleteString(wkstring)
}