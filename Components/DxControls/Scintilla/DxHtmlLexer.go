package Scintilla

import(
	"github.com/suiyunonghen/GVCL/Graphics"
)
const(
	SCE_H_DEFAULT = 0
	SCE_H_TAG = 1
	SCE_H_TAGUNKNOWN = 2
	SCE_H_ATTRIBUTE = 3
	SCE_H_ATTRIBUTEUNKNOWN = 4
	SCE_H_NUMBER = 5
	SCE_H_DOUBLESTRING = 6
	SCE_H_SINGLESTRING = 7
	SCE_H_OTHER = 8
	SCE_H_COMMENT = 9
	SCE_H_ENTITY = 10
	SCE_H_TAGEND = 11
	SCE_H_XMLSTART = 12
	SCE_H_XMLEND = 13
	SCE_H_SCRIPT = 14
	SCE_H_ASP = 15
	SCE_H_ASPAT = 16
	SCE_H_CDATA = 17
	SCE_H_QUESTION = 18
	SCE_H_VALUE = 19
	SCE_H_XCCOMMENT = 20
	SCE_H_SGML_DEFAULT = 21
	SCE_H_SGML_COMMAND = 22
	SCE_H_SGML_1ST_PARAM = 23
	SCE_H_SGML_DOUBLESTRING = 24
	SCE_H_SGML_SIMPLESTRING = 25
	SCE_H_SGML_ERROR = 26
	SCE_H_SGML_SPECIAL = 27
	SCE_H_SGML_ENTITY = 28
	SCE_H_SGML_COMMENT = 29
	SCE_H_SGML_1ST_PARAM_COMMENT = 30
	SCE_H_SGML_BLOCK_DEFAULT = 31
	SCE_HJ_START = 40
	SCE_HJ_DEFAULT = 41
	SCE_HJ_COMMENT = 42
	SCE_HJ_COMMENTLINE = 43
	SCE_HJ_COMMENTDOC = 44
	SCE_HJ_NUMBER = 45
	SCE_HJ_WORD = 46
	SCE_HJ_KEYWORD = 47
	SCE_HJ_DOUBLESTRING = 48
	SCE_HJ_SINGLESTRING = 49
	SCE_HJ_SYMBOLS = 50
	SCE_HJ_STRINGEOL = 51
	SCE_HJ_REGEX = 52
	SCE_HJA_START = 55
	SCE_HJA_DEFAULT = 56
	SCE_HJA_COMMENT = 57
	SCE_HJA_COMMENTLINE = 58
	SCE_HJA_COMMENTDOC = 59
	SCE_HJA_NUMBER = 60
	SCE_HJA_WORD = 61
	SCE_HJA_KEYWORD = 62
	SCE_HJA_DOUBLESTRING = 63
	SCE_HJA_SINGLESTRING = 64
	SCE_HJA_SYMBOLS = 65
	SCE_HJA_STRINGEOL = 66
	SCE_HJA_REGEX = 67
	SCE_HB_START = 70
	SCE_HB_DEFAULT = 71
	SCE_HB_COMMENTLINE = 72
	SCE_HB_NUMBER = 73
	SCE_HB_WORD = 74
	SCE_HB_STRING = 75
	SCE_HB_IDENTIFIER = 76
	SCE_HB_STRINGEOL = 77
	SCE_HBA_START = 80
	SCE_HBA_DEFAULT = 81
	SCE_HBA_COMMENTLINE = 82
	SCE_HBA_NUMBER = 83
	SCE_HBA_WORD = 84
	SCE_HBA_STRING = 85
	SCE_HBA_IDENTIFIER = 86
	SCE_HBA_STRINGEOL = 87
	SCE_HP_START = 90
	SCE_HP_DEFAULT = 91
	SCE_HP_COMMENTLINE = 92
	SCE_HP_NUMBER = 93
	SCE_HP_STRING = 94
	SCE_HP_CHARACTER = 95
	SCE_HP_WORD = 96
	SCE_HP_TRIPLE = 97
	SCE_HP_TRIPLEDOUBLE = 98
	SCE_HP_CLASSNAME = 99
	SCE_HP_DEFNAME = 100
	SCE_HP_OPERATOR = 101
	SCE_HP_IDENTIFIER = 102
	SCE_HPHP_COMPLEX_VARIABLE = 104
	SCE_HPA_START = 105
	SCE_HPA_DEFAULT = 106
	SCE_HPA_COMMENTLINE = 107
	SCE_HPA_NUMBER = 108
	SCE_HPA_STRING = 109
	SCE_HPA_CHARACTER = 110
	SCE_HPA_WORD = 111
	SCE_HPA_TRIPLE = 112
	SCE_HPA_TRIPLEDOUBLE = 113
	SCE_HPA_CLASSNAME = 114
	SCE_HPA_DEFNAME = 115
	SCE_HPA_OPERATOR = 116
	SCE_HPA_IDENTIFIER = 117
	SCE_HPHP_DEFAULT = 118
	SCE_HPHP_HSTRING = 119
	SCE_HPHP_SIMPLESTRING = 120
	SCE_HPHP_WORD = 121
	SCE_HPHP_NUMBER = 122
	SCE_HPHP_VARIABLE = 123
	SCE_HPHP_COMMENT = 124
	SCE_HPHP_COMMENTLINE = 125
	SCE_HPHP_HSTRING_VARIABLE = 126
	SCE_HPHP_OPERATOR = 127
)

type (
	GHtmlLexer		struct {
		fEditor					*GScintilla
		fHtmlFonts				[21]*GDxLexerFont	//Html字体设置 SCE_H_DEFAULT---SCE_H_XCCOMMENT
		fsgmlFonts				[11]*GDxLexerFont	//sgml字体设置 SCE_H_SGML_DEFAULT---SCE_H_SGML_ERROR
		fJsFonts				[13]*GDxLexerFont	//JavaScript字体设置 SCE_HJ_START --  SCE_HJ_REGEX   ,SCE_HJA_START-- SCE_HJA_REGEX
		fVbFonts				[8]*GDxLexerFont		//VBScript字体设置 SCE_HB_START --  SCE_HB_STRINGEOL ,SCE_HBA_START -- SCE_HBA_STRINGEOL
		fPyFonts				[13]*GDxLexerFont	//Python		SCE_HP_START -- SCE_HP_IDENTIFIER ,  SCE_HPA_START -- SCE_HPA_IDENTIFIER
		fPhpFonts				[11]*GDxLexerFont		//PHP  		SCE_HPHP_DEFAULT -- SCE_HPHP_COMPLEX_VARIABLE
	}
)

func (lexer *GHtmlLexer)Editor()*GScintilla  {
	return lexer.fEditor
}

func (lexer *GHtmlLexer)Update()  {
	if lexer.fEditor != nil && lexer.fEditor.HandleAllocated() {
		lexer.fEditor.SendEditor(SCI_SETLEXER, SCLEX_HTML, 0)
		if lexer.fEditor.MarginBand.ShowCodeFlod {
			lexer.fEditor.SetProperty("fold", "1")
			lexer.fEditor.SetProperty("fold.compact", "1")
			lexer.fEditor.SetProperty("fold.preprocessor", "1")
			lexer.fEditor.SetProperty("fold.scriptcomments", "1")
			lexer.fEditor.SetProperty("fold.scriptheredocs", "1")
		}
		//初始化字体信息
		for idx := SCE_H_DEFAULT;idx<=SCE_H_XCCOMMENT;idx++{
			lexer.fHtmlFonts[idx-SCE_H_DEFAULT].InitLexFont()
		}
		for idx := SCE_H_SGML_DEFAULT;idx<=SCE_H_SGML_ERROR;idx++{
			lexer.fsgmlFonts[idx-SCE_H_SGML_DEFAULT].InitLexFont()
		}

		for idx := SCE_HJ_START;idx<=SCE_HJ_REGEX;idx++{
			lexer.fJsFonts[idx-SCE_HJ_START].InitLexFont()
		}

		for idx := SCE_HB_START;idx<=SCE_HB_STRINGEOL;idx++{
			lexer.fVbFonts[idx-SCE_HB_START].InitLexFont()
		}

		for idx := SCE_HP_START;idx<=SCE_HP_IDENTIFIER;idx++{
			lexer.fPyFonts[idx-SCE_HP_START].InitLexFont()
		}

		for idx := SCE_HPHP_DEFAULT;idx<=SCE_HPHP_OPERATOR;idx++ {
			lexer.fPhpFonts[idx-SCE_HPHP_DEFAULT].InitLexFont()
		}
		lexer.fEditor.fInitLexer = true
	}
}

func (lexer *GHtmlLexer)SetEditor(scintilla *GScintilla)  {
	lexer.fEditor = scintilla
}

func (lexer *GHtmlLexer)Language()string  {
	return "HTML"
}


func (lexer *GHtmlLexer)LexerId()int  {
	return SCLEX_HTML
}

func (lexer *GHtmlLexer)KeyWords(keywordIndex int)string  {
	switch keywordIndex {
	case 2://JavaScript关键字
		return "abstract boolean break byte case catch char class const continue debugger default delete do double else enum export extends final "+
		"finally float for function goto if implements import in instanceof int interface long native new package private protected public "+
		"return short static super switch synchronized this throw throws transient try typeof var void volatile while with"
	case 3://VBScript
		return "and begin case call continue do each else elseif end erase error event exit false for function get gosub goto if implement in load loop lset me mid new next "+
		"not nothing on or property raiseevent rem resume return rset select set stop sub then to true unload until wend while with withevents attribute alias as boolean "+
		"byref byte byval const compare currency date declare dim double enum explicit friend global integer let lib long module object option optional preserve private property "+
		"public redim single static string type variant"
	case 4:	//Python
		return "and as assert break class continue def del elif else except exec finally for from global if import in is lambda None not or pass "+
		"print raise return try while with yield"
	case 5://PHP
		return "and argv as argc break case cfunction class continue declare default do die echo else elseif empty enddeclare endfor endforeach "+
		"endif endswitch endwhile e_all e_parse e_error e_warning eval exit extends false for foreach function global http_cookie_vars http_get_vars http_post_vars "+
		"http_post_files http_env_vars http_server_vars if include include_once list new not null old_function or parent php_os php_self php_version print "+
		"require require_once return static switch stdclass this true var xor virtual while __file__ __line__ __sleep __wakeup"
	default:
		return "a abbr acronym address applet area b base basefont bdo big blockquote body br button caption center cite code col colgroup "+
			"dd del dfn dir div dl dt em fieldset font form frame frameset h1 h2 h3 h4 h5 h6 head hr html i iframe img input ins isindex kbd "+
			"label legend li link map menu meta noframes noscript object ol optgroup option p param pre q s samp script select small span strike strong style "+
			"sub sup table tbody td textarea tfoot th thead title tr tt u ul var xml xmlns abbr accept-charset accept accesskey action align alink alt archive axis "+
			"background bgcolor border cellpadding cellspacing char charoff charset checked cite class classid clear codebase codetype color cols colspan compact content coords "+
			"data datafld dataformatas datapagesize datasrc datetime declare defer dir disabled enctype event face for frame frameborder headers height href hreflang hspace http-equiv "+
			"id ismap label lang language leftmargin link longdesc marginwidth marginheight maxlength media method multiple name nohref noresize noshade nowrap object onblur onchange onclick ondblclick onfocus "+
			"onkeydown onkeypress onkeyup onload onmousedown onmousemove onmouseover onmouseout onmouseup onreset onselect onsubmit onunload profile prompt readonly rel rev rows rowspan rules "+
			"scheme scope selected shape size span src standby start style summary tabindex target text title topmargin type usemap valign value valuetype version vlink vspace width "+
			"text password checkbox radio submit reset file hidden image public !doctype"
	}
}


func NewHtmlLexer()*GHtmlLexer  {
	result := new(GHtmlLexer)
	for idx := SCE_H_DEFAULT;idx<=SCE_H_XCCOMMENT;idx++{
		result.fHtmlFonts[idx-SCE_H_DEFAULT] = NewLexerFont(result,idx,0)
		if idx == SCE_H_TAG || idx==SCE_H_ATTRIBUTE{
			result.fHtmlFonts[idx].fKeyWordIndex = 1
			if idx == SCE_H_TAG{
				result.fHtmlFonts[idx].fColor = Graphics.RGB(128,0,128)
			}else{
				result.fHtmlFonts[idx].fColor = Graphics.ClBlack
			}
			result.fHtmlFonts[idx].fStyle.SetBold(true)
		}else{
			switch idx {
			case SCE_H_TAGEND:
				result.fHtmlFonts[idx].fColor = Graphics.RGB(128,0,128)
			case SCE_H_NUMBER:
				result.fHtmlFonts[idx].fColor = Graphics.ClBlue
			case SCE_H_DOUBLESTRING,SCE_H_SINGLESTRING,SCE_H_VALUE:
				result.fHtmlFonts[idx].fColor = Graphics.RGB(59,59,255)
			case SCE_H_COMMENT:
				result.fHtmlFonts[idx].fColor = Graphics.RGB(59,59,255)
			case SCE_H_CDATA:
				result.fHtmlFonts[idx].fColor = Graphics.RGB(106,135,89)
			case SCE_H_QUESTION:
				result.fHtmlFonts[idx].fColor = Graphics.ClRed
			}
		}
	}
	for idx := SCE_H_SGML_DEFAULT;idx<=SCE_H_SGML_ERROR;idx++{
		result.fsgmlFonts[idx-SCE_H_SGML_DEFAULT] = NewLexerFont(result,idx,0)
	}

	for idx := SCE_HJ_START;idx<=SCE_HJ_REGEX;idx++{
		result.fJsFonts[idx-SCE_HJ_START] = NewLexerFont(result,idx,0)
		//SCE_HJA_START
		if idx == SCE_HJ_WORD || idx==SCE_HJ_KEYWORD{
			result.fJsFonts[idx-SCE_HJ_START].fKeyWordIndex = 2
			//result.fHtmlFonts[idx]
		}
	}

	for idx := SCE_HB_START;idx<=SCE_HB_STRINGEOL;idx++{
		result.fVbFonts[idx-SCE_HB_START] = NewLexerFont(result,idx,0)
		if idx == SCE_HB_WORD{
			result.fVbFonts[idx-SCE_HB_START].fKeyWordIndex = 3
		}
	}

	for idx := SCE_HP_START;idx<=SCE_HP_IDENTIFIER;idx++{
		result.fPyFonts[idx-SCE_HP_START] = NewLexerFont(result,idx,0)
		if idx == SCE_HP_WORD{
			result.fPyFonts[idx-SCE_HP_START].fKeyWordIndex = 4
		}
	}

	for idx := SCE_HPHP_DEFAULT;idx<=SCE_HPHP_OPERATOR;idx++{
		result.fPhpFonts[idx-SCE_HPHP_DEFAULT] = NewLexerFont(result,idx,0)
		if idx == SCE_HPHP_WORD{
			result.fPhpFonts[idx-SCE_HPHP_DEFAULT].fKeyWordIndex = 5
		}
	}
	return result
}