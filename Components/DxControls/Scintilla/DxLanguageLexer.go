/*
Scintilla编辑器语法分析器包装
Autor: 不得闲
QQ:75492895
 */
package Scintilla

import (
	"github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/GVCL/WinApi"
	"github.com/suiyunonghen/DxCommonLib"
	"unsafe"
)

type(
	//HighlightFont
	GDxLexerFont struct {
		fCharset				byte
		fColor					Graphics.GColorValue		//前景色
		fBackColor				Graphics.GColorValue		//背景色
		fStyleNum				int							//高亮样式
		fStyle					Graphics.GFontStyles		//字体样式
		fKeyWords				string						//高亮关键字
		fKeyWordIndex			uint8						//所对应的关键字高亮索引
		fEditor					*GScintilla
		fUseEditorBackColor		bool						//是否使用编辑器背景色
		fOwnerLexer				*GDxLanguageLexer			//属于哪个语法高亮分析器
		FKeyWordStyle			bool						//是否是语法关键字
		fName					string						//字体名称
		fSize					uint8
	}

	//高亮词法分析器
	GDxLanguageLexer struct {
		fLexerId				int
		fEditor					*GScintilla
	}
)

func (lexerFont *GDxLexerFont)InitLexFont()  {
	var editor *GScintilla
	if lexerFont.fOwnerLexer != nil{
		editor = lexerFont.fOwnerLexer.fEditor
	}else{
		editor = lexerFont.fEditor
	}
	if editor!=nil{
		if lexerFont.FKeyWordStyle && lexerFont.fKeyWords == ""{
			return
		}
		var b []byte
		if lexerFont.FKeyWordStyle && lexerFont.fKeyWords != ""{
			b = []byte(lexerFont.fKeyWords)
			b = append(b,0)
			editor.SendEditor(SCI_SETKEYWORDS, int(lexerFont.fKeyWordIndex), int(uintptr(unsafe.Pointer(&b[0]))))
		}
		if lexerFont.fName != ""{
			//需要先转换为GBK
			b,_ = DxCommonLib.GBKString(lexerFont.fName)
			b = append(b,0)
			editor.SendEditor(SCI_STYLESETFONT,lexerFont.fStyleNum,int(uintptr(unsafe.Pointer(&b[0]))))
		}
		if lexerFont.fSize > 0{
			editor.SendEditor(SCI_STYLESETSIZE,lexerFont.fStyleNum,int(lexerFont.fSize)) //大小
		}
		editor.SendEditor(SCI_STYLESETFORE,lexerFont.fStyleNum,int(lexerFont.fColor))//前景
		if lexerFont.fUseEditorBackColor{
			editor.SendEditor(SCI_STYLESETBACK,lexerFont.fStyleNum,int(editor.GetColor()))
		}else{
			editor.SendEditor(SCI_STYLESETBACK,lexerFont.fStyleNum,int(lexerFont.fBackColor))
		}
		editor.SendEditor(SCI_STYLESETBOLD,lexerFont.fStyleNum,int(DxCommonLib.Ord(lexerFont.fStyle.Bold())))
		editor.SendEditor(SCI_STYLESETITALIC,lexerFont.fStyleNum,int(DxCommonLib.Ord(lexerFont.fStyle.Italic())))
		editor.SendEditor(SCI_STYLESETUNDERLINE,lexerFont.fStyleNum,int(DxCommonLib.Ord(lexerFont.fStyle.Underline())))
		editor.SendEditor(SCI_STYLESETCHARACTERSET,lexerFont.fStyleNum,int(lexerFont.fCharset))
	}
}

func NewLexerFont(Lexer *GDxLanguageLexer,StyleNum int,KeyWordIndex uint8)(result *GDxLexerFont)  {
	result = new(GDxLexerFont)
	result.fUseEditorBackColor = true
	result.fCharset = WinApi.DEFAULT_CHARSET
	result.fBackColor = Graphics.ClWhite
	result.fColor = Graphics.ClBlack
	result.fName = "Courier New"
	result.fSize = 10
	result.fOwnerLexer = Lexer
	result.fStyleNum = StyleNum
	result.fKeyWordIndex = KeyWordIndex
	return result
}