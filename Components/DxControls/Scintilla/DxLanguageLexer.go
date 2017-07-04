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
		fOwnerLexer				ILanguageLexer				//属于哪个语法高亮分析器
		fKeyWordStyle			bool						//是否是语法关键字
		fName					string						//字体名称
		fSize					uint8
		OtherStyleNums			[]int						//可以同时兼顾作为其他字体的样式
	}

	//高亮词法分析器
	ILanguageLexer interface {
		Language()string
		LexerId()int
		KeyWords()string
		Update()
		Editor()*GScintilla
		SetEditor(scintilla *GScintilla)
	}

	GDxLanguageLexer struct {
		fEditor					*GScintilla
		DefFont					GDxLexerFont		//默认字体
		CommentFont				GDxLexerFont		//注释字体
		KeyWordFont				GDxLexerFont		//关键字字体
		StringFont				GDxLexerFont		//字符串字体
		NumberFont				GDxLexerFont		//数字字体
		OperatorFont			GDxLexerFont		//操作符字体
	}

	GcppLexer struct {
		GDxLanguageLexer

	}

	GGoLexer	struct {
		GDxLanguageLexer
	}
)

func (lexer *GDxLanguageLexer)Editor()*GScintilla  {
	return lexer.fEditor
}

func (lexer *GDxLanguageLexer)Update()  {
	lexer.fEditor.fInitLexer = true
}

func (lexer *GDxLanguageLexer)SetEditor(scintilla *GScintilla)  {
	lexer.fEditor = scintilla
}

func (fnt *GDxLexerFont)Init()  {
	fnt.fUseEditorBackColor = true
	fnt.fCharset = WinApi.DEFAULT_CHARSET
	fnt.fBackColor = Graphics.ClWhite
	fnt.fColor = Graphics.ClBlack
	fnt.fName = "Courier New"
	fnt.fSize = 10
	fnt.fStyleNum = 0
	fnt.fKeyWordIndex = 0
}

func (lexer *GGoLexer)Language()string  {
	return "cpp"
}


func (lexer *GGoLexer)LexerId()int  {
	return SCLEX_CPP
}

func (lexer *GGoLexer)KeyWords()string  {
	return "var const package import func return defer go select interface struct break "+
			"case continue for fallthrough else if switch goto default chan type map range "+
			"uintptr int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 bool float float32 float64"
}

func (lexer *GGoLexer)Update()  {

	if lexer.fEditor != nil && lexer.fEditor.HandleAllocated(){
		lexer.fEditor.SendEditor(SCI_SETLEXER,lexer.LexerId(),0)
		if lexer.fEditor.MarginBand.ShowCodeFlod{
			lexer.fEditor.SetProperty("fold","1")
			lexer.fEditor.SetProperty("fold.compact","0")
		}
		lexer.DefFont.InitLexFont()
		lexer.NumberFont.InitLexFont()
		lexer.CommentFont.InitLexFont()
		lexer.KeyWordFont.InitLexFont()
		lexer.StringFont.InitLexFont()
		lexer.OperatorFont.InitLexFont()
		lexer.GDxLanguageLexer.Update()
	}
}


func (lexer *GcppLexer)KeyWords()string  {
	return "and and_eq asm auto bitand bitor bool break case catch char class compl const const_cast continue default delete do double dynamic_cast else enum "+
		"explicit export extern false float for friend goto if inline int long mutable namespace new not not_eq operator or or_eq private protected public register "+
		"reinterpret_cast return short signed sizeof static static_cast struct switch template this throw true try typedef typeid typename union unsigned using "+
		"virtual void volatile wchar_t while xor xor_eq"
}

func (lexerFont *GDxLexerFont)InitLexFont()  {
	var editor *GScintilla
	if lexerFont.fOwnerLexer != nil{
		editor = lexerFont.fOwnerLexer.Editor()
	}else{
		editor = lexerFont.fEditor
	}
	if editor!=nil{
		if lexerFont.fKeyWordStyle && lexerFont.fKeyWords == ""{
			return
		}
		var b []byte
		if lexerFont.fKeyWordStyle && lexerFont.fKeyWords != ""{
			b = []byte(lexerFont.fKeyWords)
			b = append(b,0)
			editor.SendEditor(SCI_SETKEYWORDS, int(lexerFont.fKeyWordIndex), int(uintptr(unsafe.Pointer(&b[0]))))
		}
		if lexerFont.fName != ""{
			b := ([]byte)(lexerFont.fName)
			b = append(b,0)
			editor.SendEditor(SCI_STYLESETFONT,lexerFont.fStyleNum,int(uintptr(unsafe.Pointer(&b[0]))))
			if lexerFont.OtherStyleNums != nil {
				for _,style := range lexerFont.OtherStyleNums{
					editor.SendEditor(SCI_STYLESETFONT,style,int(uintptr(unsafe.Pointer(&b[0]))))
				}
			}
		}
		if lexerFont.fSize > 0{
			editor.SendEditor(SCI_STYLESETSIZE,lexerFont.fStyleNum,int(lexerFont.fSize)) //大小
			if lexerFont.OtherStyleNums != nil {
				for _,style := range lexerFont.OtherStyleNums{
					editor.SendEditor(SCI_STYLESETSIZE,style,int(lexerFont.fSize)) //大小
				}
			}
		}
		editor.SendEditor(SCI_STYLESETFORE,lexerFont.fStyleNum,int(lexerFont.fColor))//前景
		if lexerFont.OtherStyleNums != nil {
			for _,style := range lexerFont.OtherStyleNums{
				editor.SendEditor(SCI_STYLESETFORE,style,int(lexerFont.fColor))//前景
			}
		}

		if lexerFont.fUseEditorBackColor{
			editor.SendEditor(SCI_STYLESETBACK,lexerFont.fStyleNum,int(editor.GetColor()))
			if lexerFont.OtherStyleNums != nil {
				for _,style := range lexerFont.OtherStyleNums{
					editor.SendEditor(SCI_STYLESETBACK,style,int(editor.GetColor()))
				}
			}
		}else{
			editor.SendEditor(SCI_STYLESETBACK,lexerFont.fStyleNum,int(lexerFont.fBackColor))
			if lexerFont.OtherStyleNums != nil {
				for _,style := range lexerFont.OtherStyleNums{
					editor.SendEditor(SCI_STYLESETBACK,style,int(lexerFont.fBackColor))
				}
			}
		}
		boldb,italicb,underlineb,_ := lexerFont.fStyle.StyleInfo()
		bold := int(DxCommonLib.Ord(boldb))
		italic := int(DxCommonLib.Ord(italicb))
		underline := int(DxCommonLib.Ord(underlineb))
		editor.SendEditor(SCI_STYLESETBOLD,lexerFont.fStyleNum,bold)
		editor.SendEditor(SCI_STYLESETITALIC,lexerFont.fStyleNum,italic)
		editor.SendEditor(SCI_STYLESETUNDERLINE,lexerFont.fStyleNum,underline)
		editor.SendEditor(SCI_STYLESETCHARACTERSET,lexerFont.fStyleNum,int(lexerFont.fCharset))
		if lexerFont.OtherStyleNums != nil {
			for _,style := range lexerFont.OtherStyleNums{
				editor.SendEditor(SCI_STYLESETBOLD,style,bold)
				editor.SendEditor(SCI_STYLESETITALIC,style,italic)
				editor.SendEditor(SCI_STYLESETUNDERLINE,style,underline)
				editor.SendEditor(SCI_STYLESETCHARACTERSET,style,int(lexerFont.fCharset))
			}
		}
	}
}

func NewLexerFont(Lexer ILanguageLexer,StyleNum int,KeyWordIndex uint8)(result *GDxLexerFont)  {
	result = new(GDxLexerFont)
	result.Init()
	result.fOwnerLexer = Lexer
	result.fStyleNum = StyleNum
	result.fKeyWordIndex = KeyWordIndex
	return result
}

func NewGoLexer()*GGoLexer  {

	result := new(GGoLexer)

	ostyleNums := [9]int{}
	result.DefFont.Init()
	ostyleNums[8] = SCE_C_IDENTIFIER
	result.DefFont.fStyleNum = SCE_C_DEFAULT
	result.DefFont.fColor = Graphics.ClBlack
	result.DefFont.OtherStyleNums = ostyleNums[8:9]
	result.DefFont.fOwnerLexer = result

	result.CommentFont.Init()
	result.CommentFont.fOwnerLexer = result
	result.CommentFont.fColor = Graphics.RGB(128,128,128)
	result.CommentFont.fStyleNum = SCE_C_COMMENT
	//所有的注释字体，全部用这个
	ostyleNums[0] = SCE_C_COMMENTLINE
	ostyleNums[1] = SCE_C_COMMENTDOC
	ostyleNums[2] = SCE_C_COMMENTLINEDOC
	ostyleNums[3] = SCE_C_COMMENTDOCKEYWORD
	ostyleNums[4] = SCE_C_COMMENTDOCKEYWORDERROR
	result.CommentFont.OtherStyleNums = ostyleNums[:5]



	result.KeyWordFont.Init()
	result.KeyWordFont.fColor = Graphics.RGB(62,62,152)
	result.KeyWordFont.fOwnerLexer = result
	result.KeyWordFont.fStyleNum = SCE_C_WORD
	result.KeyWordFont.fKeyWordStyle = true
	result.KeyWordFont.fKeyWords = result.KeyWords()
	ostyleNums[5] = SCE_C_WORD2
	result.KeyWordFont.OtherStyleNums = ostyleNums[5:6]


	result.StringFont.Init()
	result.StringFont.fOwnerLexer = result
	result.StringFont.fStyleNum = SCE_C_STRING
	result.StringFont.fColor = Graphics.RGB(106,135,89)
	ostyleNums[6] = SCE_C_STRINGEOL
	ostyleNums[7] = SCE_C_CHARACTER
	result.StringFont.OtherStyleNums = ostyleNums[6:8]


	result.NumberFont.Init()
	result.NumberFont.fOwnerLexer = result
	result.NumberFont.fColor = Graphics.ClBlue
	result.NumberFont.fStyleNum = SCE_C_NUMBER

	result.OperatorFont.Init()
	result.OperatorFont.fOwnerLexer = result
	result.OperatorFont.fStyleNum = SCE_C_OPERATOR
	return result
}