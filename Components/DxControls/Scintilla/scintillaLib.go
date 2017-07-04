package Scintilla

import(
	"github.com/suiyunonghen/GVCL/WinApi"
)
type(
	GSciChar = byte
	GSciPositionCR = int32

	GSci_CharacterRange struct {
		cpMin			GSciPositionCR
		cpMax			GSciPositionCR
	}

	GSci_TextRange struct {
		chrg			GSci_CharacterRange
		lpstrText		*byte
	}

	GSciStyle byte
	GSciCell	struct {
		charByte		GSciChar
		styleByte		GSciStyle
	}

	GSci_TextToFind struct {
		chrg			GSci_CharacterRange		//搜索的区域
		lpstrText		*byte					//搜索的内容,0结尾
		chrgText		GSci_CharacterRange		//返回匹配到的字符的区域
	}

	GSCNotification struct {
		NotifyHeader			WinApi.GNMHDR
		// SCN_STYLENEEDED, SCN_DOUBLECLICK, SCN_MODIFIED, SCN_MARGINCLICK,
		// SCN_NEEDSHOWN, SCN_DWELLSTART, SCN_DWELLEND, SCN_CALLTIPCLICK,
		// SCN_HOTSPOTCLICK, SCN_HOTSPOTDOUBLECLICK, SCN_HOTSPOTRELEASECLICK,
		// SCN_INDICATORCLICK, SCN_INDICATORRELEASE,
		// SCN_USERLISTSELECTION, SCN_AUTOCSELECTION
		Position				int
		Char					int						//SCN_CHARADDED, SCN_KEY

		// SCN_KEY, SCN_DOUBLECLICK, SCN_HOTSPOTCLICK, SCN_HOTSPOTDOUBLECLICK,
		// SCN_HOTSPOTRELEASECLICK, SCN_INDICATORCLICK, SCN_INDICATORRELEASE,
		Modifiers				int
		ModificationType		int						// SCN_MODIFIED
		Text				 	*byte					// SCN_MODIFIED
		Length					int						// SCN_MODIFIED
		LinesAdded				int						// SCN_MODIFIED
		Message					int						// SCN_MACRORECORD
		WParam					uint32					// SCN_MACRORECORD
		LParam					int32					// SCN_MACRORECORD
		Line					int						// SCN_MODIFIED
		FoldLevelNow			int						// SCN_MODIFIED
		FoldLevelPrev			int						// SCN_MODIFIED
		Margin					int						// SCN_MARGINCLICK
		ListType				int						// SCN_USERLISTSELECTION
		X						int						// SCN_DWELLSTART, SCN_DWELLEND
		Y						int                     // SCN_DWELLSTART, SCN_DWELLEND
		Token					int						// SCN_MODIFIED with SC_MOD_CONTAINER
		AnnotationLinesAdded	int						// SCN_MODIFIED with SC_MOD_CHANGEANNOTATION
		Updated					int						// SCN_UPDATEUI
	}
)

const(
	INVALID_POSITION = -1

	SCI_SETLINEINDENTATION = 2126


	//缩进的结束位置
	SCI_GETLINEINDENTPOSITION = 2128
	//获取缩进
	SCI_GETLINEINDENTATION = 2127
	//设置Tab字符宽度
	SCI_SETTABWIDTH = 2036

	//是否用空格代替Tab
	SCI_SETUSETABS = 2124
	SCI_GETUSETABS = 2125

	SCI_SETTABINDENTS = 2260
	SCI_SETBACKSPACEUNINDENTS = 2262

	SC_IV_NONE = 0
	SC_IV_REAL = 1
	SC_IV_LOOKFORWARD = 2
	SC_IV_LOOKBOTH = 3

	/// <summary>Show or hide indentation guides.</summary>
	SCI_SETINDENTATIONGUIDES = 2132

	/// <summary>Are the indentation guides visible?</summary>
	SCI_GETINDENTATIONGUIDES = 2133


	SCI_CALLTIPSHOW = 2200
	SCI_CALLTIPACTIVE = 2202
	SCI_CALLTIPSETHLT = 2204
	SCI_CALLTIPCANCEL = 2201

	SCI_GETDIRECTFUNCTION = 2184
	SCI_GETDIRECTPOINTER = 2185


	//获得整个文档的内容长度，字节数
	//SCI_GETTEXT(int length, char *text) → int
	SCI_GETLENGTH = 2006

	//可以通过本状态设定来判定文档是否修改
	SCI_SETSAVEPOINT = 2014

	//SCI_GETLINE(int line, char *text) → int
	SCI_GETLINE = 2153

	//替换选择的文字
	//SCI_REPLACESEL(<unused>, const char *text)
	SCI_REPLACESEL = 2170

	//SCI_SETREADONLY(bool readOnly)
	SCI_SETREADONLY = 2171

	//SCI_GETREADONLY → bool
	SCI_GETREADONLY = 2140

	//SCI_GETTEXTRANGE(<unused>, Sci_TextRange *tr) → int
	//如果tr的cpMax=-1，则一直获取到文档末尾
	//获取的文字会以#0结尾，所以长度应该多加1个字节
	//返回实际长度，不包括0
	SCI_GETTEXTRANGE = 2162

	//设置文档内容
	//SCI_SETTEXT(<unused>, const char *text)
	SCI_SETTEXT = 2181
	//获取指定的内容长度的数据信息到给定的内存中去
	//获取的文字会以#0结尾，所以长度应该多加1个字节
	SCI_GETTEXT = 2182

	//SCI_ALLOCATE(int bytes)
	//分配一个文档内容缓存区
	SCI_ALLOCATE = 2446

	//SCI_ADDTEXT(int length, const char *text)
	//插入的数据内容，碰到0会终止，并且不会滚动视图
	SCI_ADDTEXT = 2001


	//SCI_ADDSTYLEDTEXT(int length, cell *c)
	//和SCI_AddText类似，但是插入的是样式文字
	SCI_ADDSTYLEDTEXT = 2002

	//SCI_APPENDTEXT(int length, const char *text)
	//增加到文档末尾，其他同插入一致，也不会滚动视图
	SCI_APPENDTEXT = 2282

	//SCI_INSERTTEXT(int pos, const char *text)
	//在Pos位置插入Text，如果pos=-1，则插入到当前位置
	//会有视图移动
	SCI_INSERTTEXT = 2003

	//SCI_CHANGEINSERTION(int length, const char *text)
	//本消息主要是在SC_MOD_INSERTCHECK的消息通知要插入文字的时候
	//如果要用来修改插入的文字内容，可以使用用来修改插入的文字内容
	SCI_CHANGEINSERTION = 2672

	//除非文档是只读的，否则将删除所有文本。
	SCI_CLEARALL = 2004

	//SCI_DELETERANGE(int start, int lengthDelete)
	//删除指定开始位置和长度的文字
	SCI_DELETERANGE = 2645

	//清楚文档的样式，比如设定了Pascal的语法样式
	//调用之后，则变成常规的txt文字数据，无语法样式
	SCI_CLEARDOCUMENTSTYLE = 2005

	//SCI_GETCHARAT(int pos) → int
	//返回Pos位置所在的字符,如果为负或者文档末尾，返回0
	SCI_GETCHARAT = 2007

	//SCI_GETSTYLEAT(int pos) → int
	//获取Pos位置所在的样式
	SCI_GETSTYLEAT = 2010

	//SCI_GETSTYLEDTEXT(<unused>, Sci_TextRange *tr) → int
	//和SCI_GETTEXTRANGE差不多，区别是，这个返回的是Cell，包含有2个字节
	//SCI_GETTEXTRANGE只包含一个字节，所以这个长度应该是 2*(cpMax-cpMin)+2
	SCI_GETSTYLEDTEXT = 2015


	//Release all extended (&gt255) style numbers
	SCI_RELEASEALLEXTENDEDSTYLES = 2552

	//SCI_ALLOCATEEXTENDEDSTYLES(int numberStyles) → int
	//Allocate some extended (&gt255) style numbers and return the start of the range
	SCI_ALLOCATEEXTENDEDSTYLES = 2553

	//SCI_TARGETASUTF8(<unused>, char *s) → int
	//Returns the target converted to UTF8.Return the length in bytes
	SCI_TARGETASUTF8 = 2447

	//SCI_SETLENGTHFORENCODE(int bytes)
	//Set the length of the utf8 argument for calling EncodedFromUTF8.
	//Set to -1 and the string will be measured to the first nul.
	SCI_SETLENGTHFORENCODE = 2448

	//SCI_ENCODEDFROMUTF8(const char *utf8, char *encoded) → int
	//将UTF8字符串到文件编码。
	//Return the length of the result in bytes.
	//On error return 0.
	SCI_ENCODEDFROMUTF8 = 2449

	//SCI_SETTARGETSTART(int start)
	//用来设置搜索的目标开始的位置
	SCI_SETTARGETSTART = 2190

	//SCI_GETTARGETSTART → position
	//获得搜索开始的位置
	SCI_GETTARGETSTART = 2191

	//SCI_SETTARGETEND(int end)
	//搜索目标的结束位置
	SCI_SETTARGETEND = 2192

	//SCI_GETTARGETEND → position
	//搜索结束的位置
	SCI_GETTARGETEND = 2193

	//SCI_SETTARGETRANGE(int start, int end)
	//设置搜索开始结束位置
	SCI_SETTARGETRANGE = 2686


	//设置在选中的内容中搜索
	SCI_TARGETFROMSELECTION = 2287

	//设置为在整个文档搜索
	SCI_TARGETWHOLEDOCUMENT = 2690

	//SCI_SETSEARCHFLAGS(int searchFlags)
	//设置搜索标志 searchFlags
	SCI_SETSEARCHFLAGS = 2198

	//SCI_GETSEARCHFLAGS → int
	//返回搜索标志
	SCI_GETSEARCHFLAGS = 2199

	//SCI_SEARCHINTARGET(int length, const char *text) → int
	//通过搜索标志以及设置的搜索区域进行text的内容搜索，length指定text长度
	//搜索成功返回搜索内容的开始位置，否则返回-1
	SCI_SEARCHINTARGET = 2197

	//SCI_GETTARGETTEXT(<unused>, char *text) → int
	//获取设定区域的文字到text
	SCI_GETTARGETTEXT = 2687

	//SCI_REPLACETARGET(int length, const char *text) → int
	//如果length=-1，text为0结尾的字符串,否则是替换text中length长度的数据
	//替换完成之后，搜索区域为替换之后的文字区域，返回的是替换内容的长度
	//删除一个区域内的文字，可以设定区域，然后指定一个''字符串，然后执行命令
	SCI_REPLACETARGET = 2194

	//SCI_REPLACETARGETRE(int length, const char *text) → int
	//This replaces the target using regular expressions. If length is -1, text is a zero terminated string,
	//otherwise length is the number of characters to use. The replacement string is formed from the text string
	// with any sequences of \1 through \9 replaced by tagged matches from the most recent regular expression search.
	//\0 is replaced with all the matched text from the most recent search. After replacement, the target range refers
	//to the replacement text. The return value is the length of the replacement string.
	SCI_REPLACETARGETRE = 2195

	//SCI_GETTAG(int tagNumber, char *tagValue NUL-terminated) → int
	SCI_GETTAG = 2616

	//SCI_FINDTEXT(int searchFlags, Sci_TextToFind *ft) → position
	//在整个文档中搜索指定的内容，搜索到内容之后不会选中搜索到的内容
	//可以设置开始的位置大于结束的位置，可以从后往前搜索
	//搜索失败返回-1，成功返回匹配到的开始位置
	SCI_FINDTEXT = 2150

	//在执行SCI_SEARCHNEXT和SCI_SEARCHPREV消息前先设定本消息
	//用来设定当前位置为开始搜索的位置
	SCI_SEARCHANCHOR = 2366

	//SCI_SEARCHNEXT(int searchFlags, const char *text) → int
	//搜索下一个位置，调用之前，需要先调用SCI_SEARCHANCHOR消息
	SCI_SEARCHNEXT = 2367

	//SCI_SEARCHPREV(int searchFlags, const char *text) → int
	//搜索上一个位置
	SCI_SEARCHPREV = 2368


	//searchFlags搜索标志
	//是否整词匹配,SCI_SETWORDCHARS.设置词
	SCFIND_WHOLEWORD = 0x2
	//是否区分大小写的标志
	SCFIND_MATCHCASE = 0x4
	//单词的开始匹配 SCI_SETWORDCHARS.设置词
	SCFIND_WORDSTART = 0x00100000
	//正则表达式匹配，默认使用Sci自己的正则表达式方式
	SCFIND_REGEXP = 0x00200000
	//设置本标记之后，SCFIND_CXX11REGEX无效
	SCFIND_POSIX = 0x00400000
	//使用C++11的正则代替Sci自身的
	SCFIND_CXX11REGEX = 0x00800000

	//获取代码行数
	SCI_GETLINECOUNT = 2154

	//SCI_GETTEXTLENGTH → int
	//和SCI_GETLENGTH一样，都是返回文档长度
	SCI_GETTEXTLENGTH = 2183

	//SCI_LINESONSCREEN → int
	//返回控件区域的可见行数
	SCI_LINESONSCREEN = 2370
	//SCI_GETMODIFY → bool
	//判定文档是否修改了，根据文档的设定的保存点SCI_SETSAVEPOINT的情况来判定
	SCI_GETMODIFY = 2159
	//SCI_SETSEL(int anchor, int caret)
	//设置选中文本，如果anchor为负数，则取消选择，如果caret为负数，则到文本末尾
	SCI_SETSEL = 2160
	//SCI_GOTOPOS(int caret)
	//设置光标位置，相当于 SCI_SETSEL(caret, caret)
	SCI_GOTOPOS = 2025
	//SCI_GOTOLINE(int line)
	//调到某行的开头位置
	SCI_GOTOLINE = 2024
	//SCI_SETCURRENTPOS(int caret)
	//改变选择区域的光标的结束位置来执行加选，但是不滚动视图
	SCI_SETCURRENTPOS = 2141
	//SCI_GETCURRENTPOS → position
	SCI_GETCURRENTPOS = 2008

	//SCI_SETANCHOR(int anchor)
	//和SCI_SETCURRENTPOS相对，设定选择区域的光标的开始位置
	SCI_SETANCHOR = 2026
	//SCI_GETANCHOR → position
	SCI_GETANCHOR = 2009

	//SCI_SETSELECTIONSTART(int anchor)
	//SCI_SETSELECTIONEND(int caret)
	//
	SCI_SETSELECTIONSTART = 2142
	SCI_SETSELECTIONEND = 2144
	//SCI_GETSELECTIONSTART → position
	//SCI_GETSELECTIONEND → position
	SCI_GETSELECTIONSTART = 2143
	SCI_GETSELECTIONEND = 2145
	//SCI_SETEMPTYSELECTION(int caret)
	//移除选中，并把光标设置到caret,滚动视图
	SCI_SETEMPTYSELECTION = 2556

	//全选
	SCI_SELECTALL = 2013
	//SCI_LINEFROMPOSITION(int pos) → int
	//返回指定的pos位置所在的行
	SCI_LINEFROMPOSITION = 2166

	//SCI_POSITIONFROMLINE(int line) → position
	//返回所在行在文本中的位置
	SCI_POSITIONFROMLINE = 2167

	//SCI_GETLINEENDPOSITION(int line) → position
	//返回某行的末尾位置
	SCI_GETLINEENDPOSITION = 2136

	//SCI_LINELENGTH(int line) → int
	//返回一行内容的长度，包括行字符，如果不想包括行字符，可以使用
	// SCI_GETLINEENDPOSITION(line) - SCI_POSITIONFROMLINE(line)获得.
	SCI_LINELENGTH = 2350

	//SCI_GETSELTEXT(<unused>, char *text NUL-terminated) → int
	//获取选中的文本信息，如果要获取选中文本长度，text=nil
	SCI_GETSELTEXT = 2161

	//SCI_GETCURLINE(int length, char *text NUL-terminated) → int
	//获取当前光标位置所在的内容到text中
	SCI_GETCURLINE = 2027

	//SCI_SELECTIONISRECTANGLE → bool
	//判断选择是否是区域性选择
	SCI_SELECTIONISRECTANGLE = 2372

	//SCI_SETSELECTIONMODE(int selectionMode)
	//SCI_GETSELECTIONMODE → int
	//设置选择模式
	SCI_SETSELECTIONMODE = 2422
	SCI_GETSELECTIONMODE = 2423

	//SCI_GETLINESELSTARTPOSITION(int line) → position
	//SCI_GETLINESELENDPOSITION(int line) → position
	//Retrieve the position of the start and end of the selection at the given line with INVALID_POSITION returned if no selection on this line.
	SCI_GETLINESELSTARTPOSITION = 2424
	SCI_GETLINESELENDPOSITION = 2425

	//SCI_MOVECARETINSIDEVIEW
	//使光标在视图范围之内
	SCI_MOVECARETINSIDEVIEW = 2401

	//SCI_POSITIONBEFORE(int pos) → position
	//SCI_POSITIONAFTER(int pos) → position
	SCI_POSITIONBEFORE = 2417
	SCI_POSITIONAFTER = 2418


	//使视图滚动到光标可见
	SCI_SCROLLCARET = 2169

	//撤销重做
	SCI_BEGINUNDOACTION = 2078
	SCI_ENDUNDOACTION = 2079

	//LineEndings
	//SCI_SETEOLMODE(int eolMode)
	SCI_SETEOLMODE = 2031
	//SCI_GETEOLMODE → int
	SCI_GETEOLMODE = 2030

	SC_EOL_CRLF = 0
	SC_EOL_CR = 1
	SC_EOL_LF = 2
	//SCI_CONVERTEOLS(int eolMode)
	//有效值SC_EOL_CRLF (0), SC_EOL_CR (1), or SC_EOL_LF (2).
	SCI_CONVERTEOLS = 2029

	//SCI_SETVIEWEOL(bool visible)
	SCI_SETVIEWEOL = 2356

	//SCI_GETVIEWEOL → bool
	SCI_GETVIEWEOL = 2355

	//SCI_GETLINEENDTYPESSUPPORTED → int
	SCI_GETLINEENDTYPESSUPPORTED = 4018

	//SCI_SETLINEENDTYPESALLOWED(int lineEndBitSet)
	SCI_SETLINEENDTYPESALLOWED = 2656
	//SCI_GETLINEENDTYPESALLOWED → int
	SCI_GETLINEENDTYPESALLOWED = 2657

	//SCI_GETLINEENDTYPESACTIVE → int
	SCI_GETLINEENDTYPESACTIVE = 2658
	//颜色相关
	SC_ALPHA_TRANSPARENT = 0
	SC_ALPHA_OPAQUE = 255
	SC_ALPHA_NOALPHA = 256

	//SCI_SETSELFORE(bool useSetting, colour fore)
	//选择前景色
	SCI_SETSELFORE = 2067
	//SCI_SETSELBACK(bool useSetting, colour back)
	//背景色
	SCI_SETSELBACK = 2068

	//设置选择区域的透明度
	SCI_SETSELALPHA = 2478
	SCI_GETSELALPHA = 2477

	//h活动行是否开启
	SCI_GETCARETLINEVISIBLE = 2095
	SCI_SETCARETLINEVISIBLE = 2096

	SCI_SETCARETLINEBACK = 2098
	SCI_GETCARETLINEBACK = 2097

	SCI_GETCARETLINEVISIBLEALWAYS = 2654
	SCI_SETCARETLINEVISIBLEALWAYS = 2655

	/// <summary>Set background alpha of the caret line.</summary>
	SCI_SETCARETLINEBACKALPHA = 2470

	/// <summary>Get the background alpha of the caret line.</summary>
	SCI_GETCARETLINEBACKALPHA = 2471
	//代码折叠
	SC_FOLDLEVELBASE = 0x400
	SC_FOLDLEVELWHITEFLAG = 0x1000
	SC_FOLDLEVELHEADERFLAG = 0x2000
	SC_FOLDLEVELNUMBERMASK = 0x0FFF

	//SCI_VISIBLEFROMDOCLINE(int docLine) → int
	SCI_VISIBLEFROMDOCLINE = 2220
	//SCI_DOCLINEFROMVISIBLE(int displayLine) → int
	SCI_DOCLINEFROMVISIBLE = 2221

	//SCI_SETFOLDLEVEL(int line, int level)
	SCI_SETFOLDLEVEL = 2222
	//SCI_GETFOLDLEVEL(int line) → int
	SCI_GETFOLDLEVEL = 2223
	//SCI_SHOWLINES(int lineStart, int lineEnd)
	SCI_SHOWLINES = 2226
	//SCI_HIDELINES(int lineStart, int lineEnd)
	SCI_HIDELINES = 2227
	//SCI_GETLINEVISIBLE(int line) → bool
	SCI_GETLINEVISIBLE = 2228
	//SCI_GETALLLINESVISIBLE → bool
	SCI_GETALLLINESVISIBLE = 2236

	//SCI_SETFOLDFLAGS(int flags)
	//设置折叠区域的线条显示标记
	SCI_SETFOLDFLAGS = 2233
	//标记值
	SC_FOLDFLAG_LINEBEFORE_EXPANDED = 0x0002  //如果折叠展开，在代码区域上方绘制一个线条
	SC_FOLDFLAG_LINEBEFORE_CONTRACTED = 0x0004 //如果折叠，在上方绘制线条
	SC_FOLDFLAG_LINEAFTER_EXPANDED = 0x0008 //展开时则代码下方绘制线条
	SC_FOLDFLAG_LINEAFTER_CONTRACTED = 0x0010 //折叠时则代码下方绘制线条
	SC_FOLDFLAG_LEVELNUMBERS = 0x0040 //在线条边距中显示十六进制折叠级别以帮助调试折叠
	SC_FOLDFLAG_LINESTATE = 0x0080  //在行边距中显示十六进制行状态, 以帮助调试词法分析和折叠，不可与SC_FOLDFLAG_LEVELNUMBERS同时使用


	//SCI_GETLASTCHILD(int line, int level) → int
	SCI_GETLASTCHILD = 2224
	//SCI_GETFOLDPARENT(int line) → int
	SCI_GETFOLDPARENT = 2225

	//SCI_TOGGLEFOLD(int line)
	//本消息在点击折叠行代码点的时候，用来触发设定代码是折叠还是展开
	SCI_TOGGLEFOLD = 2231
	//SCI_TOGGLEFOLDSHOWTEXT(int line, const char *text)
	//本消息用来指定代码折叠之后，在折叠区域显示的字符串text信息
	SCI_TOGGLEFOLDSHOWTEXT = 2700

	//SCI_FOLDDISPLAYTEXTSETSTYLE(int style),指定折叠文本的显示外观
	SCI_FOLDDISPLAYTEXTSETSTYLE = 2701
	SC_FOLDDISPLAYTEXT_HIDDEN = 0 //不显示折叠文本，默认
	SC_FOLDDISPLAYTEXT_STANDARD = 1 //显示折叠文本标记
	SC_FOLDDISPLAYTEXT_BOXED = 2//用周围绘制的方框来显示文本标记


	//SCI_FOLDALL(int action)
	//折叠或者展开所有的
	//取值，SC_FOLDACTION_CONTRACT..SC_FOLDACTION_TOGGLE
	SCI_FOLDALL = 2662
	SC_FOLDACTION_CONTRACT = 0
	SC_FOLDACTION_EXPAND = 1
	SC_FOLDACTION_TOGGLE = 2
	//SCI_FOLDCHILDREN(int line, int action)
	SCI_FOLDCHILDREN=2238


	STYLE_DEFAULT = 32
	STYLE_LINENUMBER = 33
	STYLE_BRACELIGHT = 34
	STYLE_BRACEBAD = 35
	STYLE_CONTROLCHAR = 36
	STYLE_INDENTGUIDE = 37
	STYLE_CALLTIP = 38
	STYLE_LASTPREDEFINED = 39
	STYLE_MAX = 255

	//页边设置
	SC_MAX_MARGIN = 4
	//设置编辑器的页边数量
	//SCI_SETMARGINS(int margins)
	//不可超过5个
	SCI_SETMARGINS = 2252

	//SCI_GETMARGINS → int
	SCI_GETMARGINS = 2253


	SCI_MARKERDEFINE = 2040

	MARKER_MAX = 31
	SC_MARK_CIRCLE = 0
	SC_MARK_ROUNDRECT = 1
	SC_MARK_ARROW = 2
	SC_MARK_SMALLRECT = 3
	SC_MARK_SHORTARROW = 4
	SC_MARK_EMPTY = 5
	SC_MARK_ARROWDOWN = 6
	SC_MARK_MINUS = 7
	SC_MARK_PLUS = 8

	/// <summary>Shapes used for outlining column.</summary>
	SC_MARK_VLINE = 9
	SC_MARK_LCORNER = 10
	SC_MARK_TCORNER = 11
	SC_MARK_BOXPLUS = 12
	SC_MARK_BOXPLUSCONNECTED = 13
	SC_MARK_BOXMINUS = 14
	SC_MARK_BOXMINUSCONNECTED = 15
	SC_MARK_LCORNERCURVE = 16
	SC_MARK_TCORNERCURVE = 17
	SC_MARK_CIRCLEPLUS = 18
	SC_MARK_CIRCLEPLUSCONNECTED = 19
	SC_MARK_CIRCLEMINUS = 20
	SC_MARK_CIRCLEMINUSCONNECTED = 21

	/// <summary>Invisible mark that only sets the line background colour.</summary>
	SC_MARK_BACKGROUND = 22
	SC_MARK_DOTDOTDOT = 23
	SC_MARK_ARROWS = 24
	SC_MARK_PIXMAP = 25
	SC_MARK_FULLRECT = 26
	SC_MARK_LEFTRECT = 27
	SC_MARK_AVAILABLE = 28
	SC_MARK_UNDERLINE = 29
	SC_MARK_RGBAIMAGE = 30
	SC_MARK_BOOKMARK = 31

	SC_MARK_CHARACTER = 10000

	/// <summary>Markers used for outlining column.</summary>
	SC_MARKNUM_FOLDEREND = 25
	SC_MARKNUM_FOLDEROPENMID = 26
	SC_MARKNUM_FOLDERMIDTAIL = 27
	SC_MARKNUM_FOLDERTAIL = 28
	SC_MARKNUM_FOLDERSUB = 29
	SC_MARKNUM_FOLDER = 30
	SC_MARKNUM_FOLDEROPEN = 31

	SC_MARGIN_SYMBOL = 0
	SC_MARGIN_NUMBER = 1
	SC_MARGIN_BACK = 2
	SC_MARGIN_FORE = 3
	SC_MARGIN_TEXT = 4
	SC_MARGIN_RTEXT = 5

	//SCI_SETMARGINTYPEN(int margin, int marginType)
	//margin=0,1,2,3,4
	//设定页边距的类型，比如SC_MARGIN_SYMBOL=0符号页边距
	//SC_MARGIN_NUMBER (1)行号页边距等
	SCI_SETMARGINTYPEN = 2240
	//SCI_GETMARGINTYPEN(int margin) → int
	SCI_GETMARGINTYPEN = 2241

	SCI_TEXTWIDTH = 2276
	//SCI_SETMARGINWIDTHN(int margin, int pixelWidth)
	//设置页边宽度,0宽度边距不可见，默认情况下，1号页边宽16
	//如果设定行号宽度，可以使用SCI_TEXTWIDTH (STYLE_LINENUMBER, "_99999") 获得宽度再做设定
	SCI_SETMARGINWIDTHN = 2242
	//SCI_GETMARGINWIDTHN(int margin) → int
	//获取页边宽度
	SCI_GETMARGINWIDTHN = 2243

	SCI_MARKERSYMBOLDEFINED =2529



	SC_MASK_FOLDERS = 0xFE000000
	//SCI_SETMARGINMASKN(int margin, int mask)
	//SCI_SETMARGINMASKN(1, ~SC_MASK_FOLDERS)去掉折叠
	//SCI_SETMARGINMASKN(2, SC_MASK_FOLDERS)设置折叠
	SCI_SETMARGINMASKN = 2244
	//SCI_GETMARGINMASKN(int margin) → int
	SCI_GETMARGINMASKN = 2245

	//SCI_SETMARGINSENSITIVEN(int margin, bool sensitive)
	//使页边可以接收鼠标事件，默认不可接收
	SCI_SETMARGINSENSITIVEN = 2246
	//SCI_GETMARGINSENSITIVEN(int margin) → bool
	//是否支持接收鼠标事件
	SCI_GETMARGINSENSITIVEN = 2247


	SC_CURSORNORMAL = -1
	SC_CURSORARROW = 2
	SC_CURSORWAIT = 4
	SC_CURSORREVERSEARROW = 7

	SCI_MARKERLINEFROMHANDLE = 2017

	//SCI_SETMARGINCURSORN(int margin, int cursor)
	//SCI_SETMARGINCURSORN(margin, SC_CURSORARROW) SCI_SETMARGINCURSORN(margin, SC_CURSORREVERSEARROW).
	SCI_SETMARGINCURSORN = 2248
	//SCI_GETMARGINCURSORN(int margin) → int
	//反向光标
	SCI_GETMARGINCURSORN = 2249

	//SCI_SETMARGINBACKN(int margin, colour back)
	//页边类型设置为SC_MARGIN_COLOUR的时候，可以设置背景色
	SCI_SETMARGINBACKN = 2250
	//SCI_GETMARGINBACKN(int margin) → colour
	SCI_GETMARGINBACKN = 2251

	//设定页边文字左右间隔的像素，默认是1个像素
	SCI_SETMARGINLEFT = 2155
	SCI_GETMARGINLEFT = 2156
	SCI_SETMARGINRIGHT = 2157
	SCI_GETMARGINRIGHT = 2158

	//设定折叠状态栏的颜色和高亮颜色
	//SCI_SETFOLDMARGINCOLOUR(bool useSetting, colour back)
	//默认是GetSysColor(COLOR_3DFACE)
	SCI_SETFOLDMARGINCOLOUR = 2290
	//SCI_SETFOLDMARGINHICOLOUR(bool useSetting, colour fore)
	//默认是GetSysColor(COLOR_3DHIGHLIGHT).
	SCI_SETFOLDMARGINHICOLOUR = 2291


	//文字边距的相关消息//，边距类型为SC_MARGIN_TEXT
	//SCI_MARGINSETTEXT(int line, const char *text)
	SCI_MARGINSETTEXT = 2530
	//SCI_MARGINGETTEXT(int line, char *text) → int
	SCI_MARGINGETTEXT = 2531
	//SCI_MARGINSETSTYLE(int line, int style)
	SCI_MARGINSETSTYLE = 2532
	//SCI_MARGINGETSTYLE(int line) → int
	SCI_MARGINGETSTYLE = 2533
	//SCI_MARGINSETSTYLES(int line, const char *styles)
	SCI_MARGINSETSTYLES = 2534
	//SCI_MARGINGETSTYLES(int line, char *styles) → int
	SCI_MARGINGETSTYLES = 2535
	//SCI_MARGINTEXTCLEARALL
	//清空所有文字信息
	SCI_MARGINTEXTCLEARALL = 2536

	SCI_MARKERDEFINEPIXMAP = 2049

	SCI_MARKERDELETEHANDLE = 2018

	//设置相关
	SCI_SETFOCUS = 2380
	SCI_GETFOCUS = 2381

	SCI_GETBUFFEREDDRAW = 2034
	SCI_SETBUFFEREDDRAW = 2035


	SC_EFF_QUALITY_DEFAULT = 0
	SC_EFF_QUALITY_NON_ANTIALIASED = 1
	SC_EFF_QUALITY_ANTIALIASED = 2
	SC_EFF_QUALITY_LCD_OPTIMIZED = 3
	//文字质量,值为SC_EFF_QUALITY_DEFAULT (向后兼容)、SC_EFF_QUALITY_NON_ANTIALIASED、
	//SC_EFF_QUALITY_ANTIALIASED、SC_EFF_QUALITY_LCD_OPTIMIZED。
	SCI_SETFONTQUALITY = 2611
	SCI_GETFONTQUALITY = 2612

	//设定编辑器的编码格式
	SCI_SETCODEPAGE = 2037
	SCI_GETCODEPAGE = 2137
	//utf8编码
	SC_CP_UTF8 = 65001
	SC_CP_GBK = 936
	SC_CP_GB18030 = 54936
	SC_CP_Big5 = 950 //繁体
	SC_CP_JP = 932 //日语
	SC_CP_Han = 949 //韩文朝鲜
	SC_CP_Korean = 1361//韩文


	SC_IME_WINDOWED = 0
	SC_IME_INLINE=1
	//设置输入法窗口，SCI_SETIMEINTERACTION(int imeInteraction)取值
	//SC_IME_WINDOWED=0，使用默认的外挂输入法窗口
	//SC_IME_INLINE=1输入法在内部
	SCI_GETIMEINTERACTION = 2678
	SCI_SETIMEINTERACTION = 2679

	//显示相关设置
	SCI_STYLECLEARALL = 2050
	SCI_STYLESETFORE = 2051
	SCI_STYLESETBACK = 2052
	SCI_STYLESETBOLD = 2053
	SCI_STYLESETITALIC = 2054
	SCI_STYLESETSIZE = 2055
	SCI_STYLESETFONT = 2056
	SCI_STYLESETEOLFILLED = 2057
	SCI_STYLERESETDEFAULT = 2058
	SCI_STYLESETUNDERLINE = 2059

	SCI_STYLESETWEIGHT = 2063
	SCI_STYLEGETWEIGHT = 2064
	SCI_STYLESETCHARACTERSET = 2066

	SCI_SETKEYWORDS = 4005

	//事件消息相关
	SCEN_CHANGE = 768
	SCEN_SETFOCUS = 512
	SCEN_KILLFOCUS = 256


	SCI_GETENDSTYLED =2028

	SCN_STYLENEEDED = 2000
	SCN_CHARADDED = 2001
	SCN_SAVEPOINTREACHED = 2002
	SCN_SAVEPOINTLEFT = 2003
	SCN_MODIFYATTEMPTRO = 2004
	SCN_KEY = 2005
	SCN_DOUBLECLICK = 2006
	SCN_UPDATEUI = 2007
	SCN_MODIFIED = 2008
	SCN_MACRORECORD = 2009
	SCN_MARGINCLICK = 2010
	SCN_NEEDSHOWN = 2011
	SCN_PAINTED = 2013
	SCN_USERLISTSELECTION = 2014
	SCN_URIDROPPED = 2015
	SCN_DWELLSTART = 2016
	SCN_DWELLEND = 2017
	SCN_ZOOM = 2018
	SCN_HOTSPOTCLICK = 2019
	SCN_HOTSPOTDOUBLECLICK = 2020
	SCN_CALLTIPCLICK = 2021
	SCN_AUTOCSELECTION = 2022
	SCN_INDICATORCLICK = 2023
	SCN_INDICATORRELEASE = 2024
	SCN_AUTOCCANCELLED = 2025
	SCN_AUTOCCHARDELETED = 2026
	SCN_HOTSPOTRELEASECLICK = 2027
	SCN_FOCUSIN = 2028
	SCN_FOCUSOUT = 2029
	SCN_AUTOCCOMPLETED = 2030
	SCN_MARGINRIGHTCLICK = 2031

	//自动完成
	SCI_AUTOCSHOW = 2100
	SCI_AUTOCCANCEL = 2101
	SCI_AUTOCACTIVE = 2102
	SCI_AUTOCPOSSTART = 2103
	SCI_AUTOCCOMPLETE = 2104
	SCI_AUTOCSETSEPARATOR = 2106
	SCI_AUTOCGETSEPARATOR = 2107
	SCI_AUTOCSELECT = 2108
	SCI_AUTOCSETCANCELATSTART = 2110
	SCI_AUTOCGETCANCELATSTART = 2111
	SCI_AUTOCSETFILLUPS = 2112 //自动完成之后后面自动填充字符
	SCI_AUTOCSETCHOOSESINGLE = 2113
	SCI_AUTOCGETCHOOSESINGLE = 2114
	SCI_AUTOCSETIGNORECASE = 2115
	SCI_AUTOCGETIGNORECASE = 2116

	SCI_AUTOCSETAUTOHIDE = 2118
	SCI_AUTOCGETAUTOHIDE = 2119

	SCI_AUTOCGETCURRENT = 2445
	SCI_AUTOCGETCURRENTTEXT = 2610

	//设置碰到啥字符，自动停掉自动完成
	SCI_AUTOCSTOPS = 2105

	//排序
	SCI_AUTOCSETORDER = 2660
	SCI_AUTOCGETORDER = 2661

	SC_ORDER_PRESORTED = 0//需要自己在外面按照顺序排序
	SC_ORDER_PERFORMSORT = 1//编辑器内部排序
	SC_ORDER_CUSTOM = 2 //自己处理排序

	//自动完成的最大宽度
	SCI_AUTOCSETMAXWIDTH = 2208
	SCI_AUTOCGETMAXWIDTH = 2209
	SCI_AUTOCSETMAXHEIGHT = 2210
	SCI_AUTOCGETMAXHEIGHT = 2211
	SCI_REGISTERIMAGE = 2405

	//设定书签的前景背景
	SCI_MARKERSETFORE = 2041
	SCI_MARKERSETBACK = 2042

	SCI_MARKERADD = 2043
	SCI_MARKERDELETE = 2044
	SCI_MARKERDELETEALL = 2045
	SCI_MARKERGET = 2046
	SCI_MARKERNEXT = 2047
	SCI_MARKERPREVIOUS = 2048


	//设置底部的一个行间距效果
	SCI_SETEXTRAASCENT = 2525
	SCI_GETEXTRAASCENT = 2526
	//设置距离上部分的行间距效果
	SCI_SETEXTRADESCENT = 2527
	SCI_GETEXTRADESCENT = 2528


	//文档相关
	//获取当前文档内容缓存的指针指向位置
	SCI_GETDOCPOINTER = 2357
	SCI_SETDOCPOINTER = 2358

	//文档缓存增加引用计数
	SCI_ADDREFDOCUMENT = 2376
	SCI_RELEASEDOCUMENT = 2377

	//编辑相关
	//输入设置为Unicode
	SCI_SETKEYSUNICODE = 2521
	SCI_GETKEYSUNICODE = 2522


	SC_CHARSET_ANSI = 0
	SC_CHARSET_DEFAULT = 1
	SC_CHARSET_BALTIC = 186
	SC_CHARSET_CHINESEBIG5 = 136
	SC_CHARSET_EASTEUROPE = 238
	SC_CHARSET_GB2312 = 134
	SC_CHARSET_GREEK = 161
	SC_CHARSET_HANGUL = 129
	SC_CHARSET_MAC = 77
	SC_CHARSET_OEM = 255
	SC_CHARSET_RUSSIAN = 204
	SC_CHARSET_CYRILLIC = 1251
	SC_CHARSET_SHIFTJIS = 128
	SC_CHARSET_SYMBOL = 2
	SC_CHARSET_TURKISH = 162
	SC_CHARSET_JOHAB = 130
	SC_CHARSET_HEBREW = 177
	SC_CHARSET_ARABIC = 178
	SC_CHARSET_VIETNAMESE = 163
	SC_CHARSET_THAI = 222
	SC_CHARSET_8859_15 = 1000


	SCI_SETLEXER = 4001

	/// <summary>Retrieve the lexing language of the document.</summary>
	SCI_GETLEXER = 4002
	SCI_SETPROPERTY = 4004

	SCI_SETLEXERLANGUAGE = 4006

	//词法分析器
	SCLEX_CONTAINER = 0 //如果SCI_SETLEXER设置为本分析器，则为用户自定义词法分析器，每次都会触发 SCN_STYLENEEDED
	SCLEX_NULL = 1
	SCLEX_PYTHON = 2
	SCLEX_CPP = 3
	SCLEX_HTML = 4
	SCLEX_XML = 5
	SCLEX_PERL = 6
	SCLEX_SQL = 7
	SCLEX_VB = 8
	SCLEX_PROPERTIES = 9
	SCLEX_ERRORLIST = 10
	SCLEX_MAKEFILE = 11
	SCLEX_BATCH = 12
	SCLEX_XCODE = 13
	SCLEX_LATEX = 14
	SCLEX_LUA = 15
	SCLEX_DIFF = 16
	SCLEX_CONF = 17
	SCLEX_PASCAL = 18
	SCLEX_AVE = 19
	SCLEX_ADA = 20
	SCLEX_LISP = 21
	SCLEX_RUBY = 22
	SCLEX_EIFFEL = 23
	SCLEX_EIFFELKW = 24
	SCLEX_TCL = 25
	SCLEX_NNCRONTAB = 26
	SCLEX_BULLANT = 27
	SCLEX_VBSCRIPT = 28
	SCLEX_BAAN = 31
	SCLEX_MATLAB = 32
	SCLEX_SCRIPTOL = 33
	SCLEX_ASM = 34
	SCLEX_CPPNOCASE = 35
	SCLEX_FORTRAN = 36
	SCLEX_F77 = 37
	SCLEX_CSS = 38
	SCLEX_POV = 39
	SCLEX_LOUT = 40
	SCLEX_ESCRIPT = 41
	SCLEX_PS = 42
	SCLEX_NSIS = 43
	SCLEX_MMIXAL = 44
	SCLEX_CLW = 45
	SCLEX_CLWNOCASE = 46
	SCLEX_LOT = 47
	SCLEX_YAML = 48
	SCLEX_TEX = 49
	SCLEX_METAPOST = 50
	SCLEX_POWERBASIC = 51
	SCLEX_FORTH = 52
	SCLEX_ERLANG = 53
	SCLEX_OCTAVE = 54
	SCLEX_MSSQL = 55
	SCLEX_VERILOG = 56
	SCLEX_KIX = 57
	SCLEX_GUI4CLI = 58
	SCLEX_SPECMAN = 59
	SCLEX_AU3 = 60
	SCLEX_APDL = 61
	SCLEX_BASH = 62
	SCLEX_ASN1 = 63
	SCLEX_VHDL = 64
	SCLEX_CAML = 65
	SCLEX_BLITZBASIC = 66
	SCLEX_PUREBASIC = 67
	SCLEX_HASKELL = 68
	SCLEX_PHPSCRIPT = 69
	SCLEX_TADS3 = 70
	SCLEX_REBOL = 71
	SCLEX_SMALLTALK = 72
	SCLEX_FLAGSHIP = 73
	SCLEX_CSOUND = 74
	SCLEX_FREEBASIC = 75
	SCLEX_INNOSETUP = 76
	SCLEX_OPAL = 77
	SCLEX_SPICE = 78
	SCLEX_D = 79
	SCLEX_CMAKE = 80
	SCLEX_GAP = 81
	SCLEX_PLM = 82
	SCLEX_PROGRESS = 83
	SCLEX_ABAQUS = 84
	SCLEX_ASYMPTOTE = 85
	SCLEX_R = 86
	SCLEX_MAGIK = 87
	SCLEX_POWERSHELL = 88
	SCLEX_MYSQL = 89
	SCLEX_PO = 90
	SCLEX_TAL = 91
	SCLEX_COBOL = 92
	SCLEX_TACL = 93
	SCLEX_SORCUS = 94
	SCLEX_POWERPRO = 95
	SCLEX_NIMROD = 96
	SCLEX_SML = 97
	SCLEX_MARKDOWN = 98
	SCLEX_TXT2TAGS = 99
	SCLEX_A68K = 100
	SCLEX_MODULA = 101
	SCLEX_COFFEESCRIPT = 102
	SCLEX_TCMD = 103
	SCLEX_AVS = 104
	SCLEX_ECL = 105
	SCLEX_OSCRIPT = 106
	SCLEX_VISUALPROLOG = 107
	SCLEX_LITERATEHASKELL = 108
	SCLEX_STTXT = 109
	SCLEX_KVIRC = 110
	SCLEX_RUST = 111
	SCLEX_DMAP = 112
	SCLEX_AS = 113
	SCLEX_DMIS = 114

	/// <summary>When a lexer specifies its language as SCLEX_AUTOMATIC it receives a
	/// value assigned in sequence from SCLEX_AUTOMATIC+1.</summary>
	SCLEX_AUTOMATIC = 1000

	SCI_GETCOLUMN =  2129

	//复制粘贴
	SCI_CUT = 2177
	SCI_COPY  = 2178
	SCI_PASTE = 2179
	SCI_CLEAR = 2180
	SCI_CANPASTE = 2173
	SCI_COPYALLOWLINE = 2519
	SCI_GETSELECTIONEMPTY = 2650 //"can copy" or "can cut"用这个获取
	SCI_GETSELECTIONS = 2570
	//SCI_COPYRANGE(int start, int end)
	SCI_COPYRANGE = 2419
	//SCI_COPYTEXT(int length, const char *text)
	SCI_COPYTEXT = 2420

	SC_STATUS_OK = 0
	SC_STATUS_FAILURE=1
	SC_STATUS_BADALLOC=2
	SC_STATUS_WARN_REGEX  = 1001
	//错误处理，如果发生错误，可以使用SCI_GETSTATUS获取内部的错误编码，可以使用SCI_SETSTATUS(0)，清除错误编码
	//Status values from 1 to 999 are errors and status SC_STATUS_WARN_START (1000) and above are warnings. The currently defined statuses are:
	//SC_STATUS_OK=0,SC_STATUS_FAILURE=1，SC_STATUS_BADALLOC=2，SC_STATUS_WARN_REGEX  = 1001
	//SCI_SETSTATUS(int status)
	SCI_SETSTATUS = 2382
	//SCI_GETSTATUS → int
	SCI_GETSTATUS = 2383

	//撤销，重做
	SCI_CANUNDO = 2174
	SCI_UNDO = 2176
	SCI_EMPTYUNDOBUFFER = 2175
	SCI_REDO =2011
	SCI_CANREDO = 2016


	//单词，返回单词的开头和结尾
	//SCI_WORDSTARTPOSITION(int pos, bool onlyWordCharacters) → int
	SCI_WORDSTARTPOSITION = 2266
	//SCI_WORDENDPOSITION(int pos, bool onlyWordCharacters) → int
	SCI_WORDENDPOSITION = 2267

	//SCI_ISRANGEWORD(int start, int end) → bool
	//检查一个指定区域start和end之间的内容，是否开始于一个单词，并且结束于一个单词，不检查中间的空格
	SCI_ISRANGEWORD = 2691

	//SCI_SETWORDCHARS(0, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	//定义组成单词的字符集
	SCI_SETWORDCHARS = 2077
	//SCI_GETWORDCHARS(<unused>, char *characters) → int
	SCI_GETWORDCHARS = 2646

	//SCI_SETWHITESPACECHARS(<unused>, const char *characters)
	//定义将哪些字符认定作为空格字符对待，这样Ctrl+向右或者向左键的时候可以直接跳过这些字符
	SCI_SETWHITESPACECHARS = 2443
	SCI_GETWHITESPACECHARS = 2647

	//SCI_SETPUNCTUATIONCHARS(<unused>, const char *characters)
	//SCI_GETPUNCTUATIONCHARS(<unused>, char *characters) → int
	//将哪些字符作为标点符号字符
	SCI_SETPUNCTUATIONCHARS = 2648
	SCI_GETPUNCTUATIONCHARS = 2649
	//以上都应该在 SCI_SETWORDCHARS之后调用

	//使用默认的 word 和空白字符组
	SCI_SETCHARSDEFAULT =2444

	/// <summary>Move caret left one word.</summary>
	SCI_WORDLEFT = 2308

	/// <summary>Move caret left one word extending selection to new caret position.</summary>
	SCI_WORDLEFTEXTEND = 2309

	/// <summary>Move caret right one word.</summary>
	SCI_WORDRIGHT = 2310

	/// <summary>Move caret right one word extending selection to new caret position.</summary>
	SCI_WORDRIGHTEXTEND = 2311

	/// <summary>Move caret to first position on line.</summary>
	SCI_HOME = 2312

	/// <summary>Move caret to first position on line extending selection to new caret position.</summary>
	SCI_HOMEEXTEND = 2313

	/// <summary>Move caret to last position on line.</summary>
	SCI_LINEEND = 2314

	/// <summary>Move caret to last position on line extending selection to new caret position.</summary>
	SCI_LINEENDEXTEND = 2315

	/// <summary>Move caret to first position in document.</summary>
	SCI_DOCUMENTSTART = 2316

	/// <summary>Move caret to first position in document extending selection to new caret position.</summary>
	SCI_DOCUMENTSTARTEXTEND = 2317

	/// <summary>Move caret to last position in document.</summary>
	SCI_DOCUMENTEND = 2318

	/// <summary>Move caret to last position in document extending selection to new caret position.</summary>
	SCI_DOCUMENTENDEXTEND = 2319

	/// <summary>Move caret one page up.</summary>
	SCI_PAGEUP = 2320

	/// <summary>Move caret one page up extending selection to new caret position.</summary>
	SCI_PAGEUPEXTEND = 2321

	/// <summary>Move caret one page down.</summary>
	SCI_PAGEDOWN = 2322

	/// <summary>Move caret one page down extending selection to new caret position.</summary>
	SCI_PAGEDOWNEXTEND = 2323

	/// <summary>Move caret left one word, position cursor at end of word.</summary>
	SCI_WORDLEFTEND = 2439

	/// <summary>Move caret left one word, position cursor at end of word, extending selection to new caret position.</summary>
	SCI_WORDLEFTENDEXTEND = 2440

	/// <summary>Move caret right one word, position cursor at end of word.</summary>
	SCI_WORDRIGHTEND = 2441

	/// <summary>Move caret right one word, position cursor at end of word, extending selection to new caret position.</summary>
	SCI_WORDRIGHTENDEXTEND = 2442

	/// <summary>Move to the previous change in capitalisation.</summary>
	SCI_WORDPARTLEFT = 2390

	/// <summary>Move to the previous change in capitalisation extending selection
	/// to new caret position.</summary>
	SCI_WORDPARTLEFTEXTEND = 2391

	/// <summary>Move to the change next in capitalisation.</summary>
	SCI_WORDPARTRIGHT = 2392

	/// <summary>Move to the next change in capitalisation extending selection
	/// to new caret position.</summary>
	SCI_WORDPARTRIGHTEXTEND = 2393

	/// <summary>Delete the word to the left of the caret.</summary>
	SCI_DELWORDLEFT = 2335

	/// <summary>Delete the word to the right of the caret.</summary>
	SCI_DELWORDRIGHT = 2336

	/// <summary>Delete the word to the right of the caret, but not the trailing non-word characters.</summary>
	SCI_DELWORDRIGHTEND = 2518

	//Annotations批注
	//批注是每行可编辑文本下方的只读文本行.批注可以由多个由 "\ n" 分隔的行组成
	//可以用作代码信息的提示
	//SCI_ANNOTATIONSETTEXT(int line, const char *text)
	SCI_ANNOTATIONSETTEXT = 2540

	//SCI_ANNOTATIONGETTEXT(int line, char *text) → int
	SCI_ANNOTATIONGETTEXT = 2541

	//SCI_ANNOTATIONSETSTYLE(int line, int style)
	SCI_ANNOTATIONSETSTYLE = 2542

	//SCI_ANNOTATIONGETSTYLE(int line) → int
	SCI_ANNOTATIONGETSTYLE = 2543

	//SCI_ANNOTATIONSETSTYLES(int line, const char *styles)
	SCI_ANNOTATIONSETSTYLES = 2544

	//SCI_ANNOTATIONGETSTYLES(int line, char *styles) → int
	SCI_ANNOTATIONGETSTYLES = 2545

	//SCI_ANNOTATIONGETLINES(int line) → int
	SCI_ANNOTATIONGETLINES = 2546
	SCI_ANNOTATIONCLEARALL = 2547
	ANNOTATION_HIDDEN = 0
	ANNOTATION_STANDARD = 1
	ANNOTATION_BOXED = 2
	ANNOTATION_INDENTED = 3
	SCI_ANNOTATIONSETVISIBLE = 2548
	SCI_ANNOTATIONGETVISIBLE = 2549
	SCI_ANNOTATIONSETSTYLEOFFSET = 2550
	SCI_ANNOTATIONGETSTYLEOFFSET = 2551
	//PascalLexer
	SCE_PAS_DEFAULT = 0
	SCE_PAS_IDENTIFIER = 1
	SCE_PAS_COMMENT = 2
	SCE_PAS_COMMENT2 = 3
	SCE_PAS_COMMENTLINE = 4
	SCE_PAS_PREPROCESSOR = 5
	SCE_PAS_PREPROCESSOR2 = 6
	SCE_PAS_NUMBER = 7
	SCE_PAS_HEXNUMBER = 8
	SCE_PAS_WORD = 9
	SCE_PAS_STRING = 10
	SCE_PAS_STRINGEOL = 11
	SCE_PAS_CHARACTER = 12
	SCE_PAS_OPERATOR = 13
	SCE_PAS_ASM = 14


	//CppLexer
	SCE_C_DEFAULT = 0
	SCE_C_COMMENT = 1
	SCE_C_COMMENTLINE = 2
	SCE_C_COMMENTDOC = 3
	SCE_C_NUMBER = 4
	SCE_C_WORD = 5
	SCE_C_STRING = 6
	SCE_C_CHARACTER = 7
	SCE_C_UUID = 8
	SCE_C_PREPROCESSOR = 9
	SCE_C_OPERATOR = 10
	SCE_C_IDENTIFIER = 11
	SCE_C_STRINGEOL = 12
	SCE_C_VERBATIM = 13
	SCE_C_REGEX = 14
	SCE_C_COMMENTLINEDOC = 15
	SCE_C_WORD2 = 16
	SCE_C_COMMENTDOCKEYWORD = 17
	SCE_C_COMMENTDOCKEYWORDERROR = 18
	SCE_C_GLOBALCLASS = 19
	SCE_C_STRINGRAW = 20
	SCE_C_TRIPLEVERBATIM = 21
	SCE_C_HASHQUOTEDSTRING = 22
	SCE_C_PREPROCESSORCOMMENT = 23
	SCE_C_PREPROCESSORCOMMENTDOC = 24
	SCE_C_USERLITERAL = 25
	SCE_C_TASKMARKER = 26
	SCE_C_ESCAPESEQUENCE = 27

	//Modifiers常量
	SCMOD_NORM = 0
	SCMOD_SHIFT = 1
	SCMOD_CTRL = 2
	SCMOD_ALT = 4
	SCMOD_SUPER = 8
	SCMOD_META = 16


	//UpdateUI的Update
	SC_UPDATE_CONTENT = 1
	SC_UPDATE_SELECTION = 2
	SC_UPDATE_V_SCROLL = 4
	SC_UPDATE_H_SCROLL = 8


	//
	SC_MOD_INSERTTEXT = 0x1
	SC_MOD_DELETETEXT = 0x2
	SC_MOD_CHANGESTYLE = 0x4
	SC_MOD_CHANGEFOLD = 0x8
	SC_PERFORMED_USER = 0x10
	SC_PERFORMED_UNDO = 0x20
	SC_PERFORMED_REDO = 0x40
	SC_MULTISTEPUNDOREDO = 0x80
	SC_LASTSTEPINUNDOREDO = 0x100
	SC_MOD_CHANGEMARKER = 0x200
	SC_MOD_BEFOREINSERT = 0x400
	SC_MOD_BEFOREDELETE = 0x800
	SC_MULTILINEUNDOREDO = 0x1000
	SC_STARTACTION = 0x2000
	SC_MOD_CHANGEINDICATOR = 0x4000
	SC_MOD_CHANGELINESTATE = 0x8000
	SC_MOD_CHANGEMARGIN = 0x10000
	SC_MOD_CHANGEANNOTATION = 0x20000
	SC_MOD_CONTAINER = 0x40000
	SC_MOD_LEXERSTATE = 0x80000
	SC_MOD_INSERTCHECK = 0x100000
	SC_MOD_CHANGETABSTOPS = 0x200000
	SC_MODEVENTMASKALL = 0x3FFFFF
)

