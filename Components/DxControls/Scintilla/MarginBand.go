/*
Scintilla编辑器控件的Margin封装
Autor: 不得闲
QQ:75492895
 */
package Scintilla

import(
	"github.com/suiyunonghen/GVCL/Graphics"
	"github.com/suiyunonghen/DxCommonLib"
	"unsafe"
	"fmt"
)

type (
	//边条单击事件
	GMarginClick func(sender *GDxMarginBand,pos int,MarginIndex int,modifiers int)
	//编辑器边条
	GDxMarginBand struct {
		fcodeEditor					*GScintilla
		ShowLineNum					bool	//是否显示行号
		ShowCodeFlod				bool	//是否代码折叠
		TextMargin					bool	//是否显示文字边
		fBookmarks					[10]int //保存书签句柄，只保存10个
		ShowBookMark				bool	//是否显示书签
		fCodeFlodInited				bool	//折叠初始化
		fLineNumIndex				byte	//行边索引
		fTextIndex					byte	//文字边索引
		fFoldIndex					byte	//折叠边索引
		fBookmarkIndex				byte
		Color						Graphics.GColorValue	//边条颜色
		OnClick						GMarginClick	//单击事件
		OnRightClick				GMarginClick	//右键单击
		fupcount					int
		TextMarginWidth				int				//文字边宽度
		BookMarkBack				Graphics.GColorValue
		BookMarkFont				Graphics.GFont
		fBookMarkStyle				int				//书签的文字样式，默认指定254
		LineNumFont					Graphics.GFont
	}
)

func (band *GDxMarginBand)BeginUpdate()  {
	band.fupcount++
}

func (band *GDxMarginBand)EndUpdate()  {
	band.fupcount--
	if band.fupcount<=0{
		band.fupcount=0
		band.Update()
	}
}

func (band *GDxMarginBand)ClearMarks()  {
	if band.fcodeEditor == nil || !band.fcodeEditor.HandleAllocated(){
		return
	}
	band.fcodeEditor.SendEditor(SCI_MARKERDELETEALL,-1,0)
	band.fcodeEditor.SendEditor(SCI_MARGINTEXTCLEARALL,0,0) //清空所有标记
	for i := 0;i<10;i++{
		band.fBookmarks[i] = -1
	}
}

func (band *GDxMarginBand)GotoBookmark(bkindex int)  {
	if band.fcodeEditor == nil || !band.fcodeEditor.HandleAllocated(){
		return
	}
	band.fcodeEditor.GoToLine(band.fBookmarks[bkindex])
}

func (band *GDxMarginBand)UpdateBookMark()  {
	//指定书签的显示样式
	if band.ShowBookMark{
		bold,italic,underline,_ := band.BookMarkFont.FontStyle().StyleInfo()
		b := ([]byte)(band.BookMarkFont.FontName)
		b = append(b,0)
		band.fcodeEditor.SendEditor(SCI_STYLESETFONT,band.fBookMarkStyle,int(uintptr(unsafe.Pointer(&b[0]))))
		band.fcodeEditor.SendEditor(SCI_STYLESETITALIC,band.fBookMarkStyle,int(DxCommonLib.Ord(italic)))
		band.fcodeEditor.SendEditor(SCI_STYLESETBOLD,band.fBookMarkStyle,int(DxCommonLib.Ord(bold)))
		band.fcodeEditor.SendEditor(SCI_STYLESETUNDERLINE,band.fBookMarkStyle,int(DxCommonLib.Ord(underline)))
		band.fcodeEditor.SendEditor(SCI_STYLESETFORE,band.fBookMarkStyle,int(band.BookMarkFont.Color))
		band.fcodeEditor.SendEditor(SCI_STYLESETBACK,band.fBookMarkStyle,int(band.BookMarkBack))
	}
	if band.ShowLineNum{
		bold,italic,underline,_ := band.LineNumFont.FontStyle().StyleInfo()
		b := ([]byte)(band.LineNumFont.FontName)
		b = append(b,0)

		band.fcodeEditor.SendEditor(SCI_STYLESETFONT,STYLE_LINENUMBER,int(uintptr(unsafe.Pointer(&b[0]))))
		band.fcodeEditor.SendEditor(SCI_STYLESETITALIC,STYLE_LINENUMBER,int(DxCommonLib.Ord(italic)))
		band.fcodeEditor.SendEditor(SCI_STYLESETBOLD,STYLE_LINENUMBER,int(DxCommonLib.Ord(bold)))
		band.fcodeEditor.SendEditor(SCI_STYLESETUNDERLINE,STYLE_LINENUMBER,int(DxCommonLib.Ord(underline)))
		band.fcodeEditor.SendEditor(SCI_STYLESETFORE,STYLE_LINENUMBER,int(band.LineNumFont.Color))
		//band.fcodeEditor.SendEditor(SCI_STYLESETBACK,STYLE_LINENUMBER,int(band.BookMarkBack))
	}
}

func (band *GDxMarginBand)SetBookmark(bkindex int)  {
	if band.fcodeEditor == nil || !band.fcodeEditor.HandleAllocated(){
		return
	}
	if band.fBookmarks[bkindex] != -1{
		band.fcodeEditor.SendEditor(SCI_MARGINSETTEXT,band.fBookmarks[bkindex],0)
	}
	bt := DxCommonLib.FastString2Byte(fmt.Sprintf("%d",bkindex))
	band.fcodeEditor.SendEditor(SCI_MARGINSETTEXT,band.fcodeEditor.CaretPos.Line,int(uintptr(unsafe.Pointer(&bt[0]))))
	band.fcodeEditor.SendEditor(SCI_MARGINSETSTYLE,band.fcodeEditor.CaretPos.Line,band.fBookMarkStyle) //设置书签样式
	band.fBookmarks[bkindex] = band.fcodeEditor.CaretPos.Line
}

func (band *GDxMarginBand)Update()  {
	if band.fcodeEditor == nil || !band.fcodeEditor.HandleAllocated(){
		return
	}
	band.fcodeEditor.SendEditor(SCI_STYLESETBACK,33, int(band.Color)) //设置页边颜色
	//0号位置行号，1号位置是书签，2号位置文字页边，3号位置折叠
	var idx byte = 0
	if band.ShowLineNum {
		band.fLineNumIndex = 0
		idx++
	}else{
		band.fLineNumIndex = 10
	}
	if band.ShowBookMark{
		band.fBookmarkIndex = idx
		idx++
	}else{
		band.fBookmarkIndex = 10
	}
	if band.TextMargin && band.TextMarginWidth > 8{
		band.fTextIndex = idx
		idx++
	}else{
		band.fTextIndex = 10
	}
	if band.ShowCodeFlod{
		band.fFoldIndex = idx
		idx++
	}else{
		band.fFoldIndex = 10
	}
	//设置页边个数
	band.fcodeEditor.SendEditor(SCI_SETMARGINS,int(idx),0)
	if band.ShowLineNum{//行号
		band.fcodeEditor.SendEditor(SCI_SETMARGINTYPEN,int(band.fLineNumIndex), SC_MARGIN_NUMBER)
		str := fmt.Sprintf("_%d",band.fcodeEditor.CodeLines.Count())
		band.fcodeEditor.SendEditor(SCI_SETMARGINWIDTHN,int(band.fLineNumIndex), band.MarginTextLen(str)) //页边长度
		if band.ShowBookMark{
			band.fcodeEditor.SendEditor(SCI_SETMARGINCURSORN,int(band.fLineNumIndex), SC_CURSORARROW)
			band.fcodeEditor.SendEditor(SCI_SETMARGINSENSITIVEN,int(band.fLineNumIndex),1) //接受鼠标点击
		}else{
			band.fcodeEditor.SendEditor(SCI_SETMARGINCURSORN,int(band.fLineNumIndex), SC_CURSORREVERSEARROW)
		}
	}

	if band.ShowBookMark{
		band.fcodeEditor.SendEditor(SCI_SETMARGINTYPEN,int(band.fBookmarkIndex), SC_MARGIN_TEXT) //文字边距
		band.fcodeEditor.SendEditor(SCI_SETMARGINWIDTHN,int(band.fBookmarkIndex), 16)
		//这个是0-9位显示，其他的不显示,$03FF设定显示掩码
		band.fcodeEditor.SendEditor(SCI_SETMARGINMASKN, int(band.fBookmarkIndex), 0x3FF)
		band.fcodeEditor.SendEditor(SCI_SETMARGINSENSITIVEN,int(band.fBookmarkIndex),1); //接受鼠标点击
		band.fcodeEditor.SendEditor(SCI_SETMARGINCURSORN,int(band.fBookmarkIndex), SC_CURSORARROW)

		band.UpdateBookMark()
	}

	if band.TextMargin{
		band.fcodeEditor.SendEditor(SCI_SETMARGINTYPEN,int(band.fTextIndex), SC_MARGIN_TEXT) //文字边距
		band.fcodeEditor.SendEditor(SCI_SETMARGINWIDTHN,int(band.fTextIndex), band.TextMarginWidth) //页边长度
	}

	//代码折叠
	if band.ShowCodeFlod{
		band.fcodeEditor.SendEditor(SCI_SETMARGINTYPEN, int(band.fFoldIndex), SC_MARGIN_SYMBOL)
		band.fcodeEditor.SendEditor(SCI_SETMARGINMASKN, int(band.fFoldIndex), SC_MASK_FOLDERS)
		band.fcodeEditor.SendEditor(SCI_SETMARGINWIDTHN, int(band.fFoldIndex), 11) //页宽
		band.fcodeEditor.SendEditor(SCI_SETMARGINSENSITIVEN, int(band.fFoldIndex), 1)
		if !band.fCodeFlodInited{
			band.fCodeFlodInited = true
			band.fcodeEditor.SendEditor(SCI_MARKERDEFINE, SC_MARKNUM_FOLDER, SC_MARK_CIRCLEPLUS);
			band.fcodeEditor.SendEditor(SCI_MARKERDEFINE, SC_MARKNUM_FOLDEROPEN, SC_MARK_CIRCLEMINUS);
			band.fcodeEditor.SendEditor(SCI_MARKERDEFINE, SC_MARKNUM_FOLDEREND,  SC_MARK_CIRCLEPLUSCONNECTED);
			band.fcodeEditor.SendEditor(SCI_MARKERDEFINE, SC_MARKNUM_FOLDEROPENMID, SC_MARK_CIRCLEMINUSCONNECTED);
			band.fcodeEditor.SendEditor(SCI_MARKERDEFINE, SC_MARKNUM_FOLDERMIDTAIL, SC_MARK_TCORNERCURVE);
			band.fcodeEditor.SendEditor(SCI_MARKERDEFINE, SC_MARKNUM_FOLDERSUB, SC_MARK_VLINE);
			band.fcodeEditor.SendEditor(SCI_MARKERDEFINE, SC_MARKNUM_FOLDERTAIL, SC_MARK_LCORNERCURVE);

			band.fcodeEditor.SendEditor(SCI_SETFOLDFLAGS, 16 | 4, 0) //如果折叠就在折叠行的上下各画一条横线
			band.fcodeEditor.SendEditor(SCI_MARKERSETBACK, SC_MARKNUM_FOLDERSUB, 0xa0a0a0)
			band.fcodeEditor.SendEditor(SCI_MARKERSETBACK, SC_MARKNUM_FOLDERMIDTAIL, 0xa0a0a0)
			band.fcodeEditor.SendEditor(SCI_MARKERSETBACK, SC_MARKNUM_FOLDERTAIL, 0xa0a0a0)
			if band.fcodeEditor.fInitLexer{
				band.fcodeEditor.SetProperty("fold","1")
				//band.fcodeEditor.SetProperty("fold.compact","0")
				band.fcodeEditor.SetProperty("fold.comment","1")
				band.fcodeEditor.SetProperty("fold.preprocessor","1")
			}
		}
	}else{
		band.fcodeEditor.SetProperty("fold","0")
		band.fcodeEditor.SendEditor(SCI_SETMARGINSENSITIVEN, 2, 0)//去掉鼠标事件
	}


}

func (band *GDxMarginBand)MarginTextLen(text string)int  {
	if band.fcodeEditor != nil && text != "" {
		bt := DxCommonLib.FastString2Byte(text)
		return band.fcodeEditor.SendEditor(SCI_TEXTWIDTH,STYLE_LINENUMBER,int(uintptr(unsafe.Pointer(&bt[0]))))
	}else{
		return  4
	}
}

func (band *GDxMarginBand)FindValidBookmark(ClickLine int)(CurBookmark int,result int)  {
	result = -1
	CurBookmark = -1
	for i := 0;i<10;i++{
		if band.fBookmarks[i] == -1{
			if result == -1{
				result = i
			}
		}else {
			bt := [1]byte{16}
			if band.fcodeEditor.SendEditor(SCI_MARGINGETTEXT,ClickLine,int(uintptr(unsafe.Pointer(&bt[0])))) > 0{
				CurBookmark = int(bt[0]) -48
			}
		}
	}
	return
}

func (band *GDxMarginBand)BandClick(pos,MarginIndex,modifiers int)  {
	lineNumber := band.fcodeEditor.SendEditor(SCI_LINEFROMPOSITION,pos,0)
	if MarginIndex == int(band.fBookmarkIndex) || MarginIndex == int(band.fLineNumIndex){
		if band.ShowBookMark{
			//先判断ClickLine有没有书签
			var bt []byte
			CurBookmark,ValidBkIndex := band.FindValidBookmark(lineNumber)
			if CurBookmark == -1{
				//没有书签，加入书签
				if ValidBkIndex == -1{
					//然后将0位置的标签去掉
					band.fcodeEditor.SendEditor(SCI_MARGINSETTEXT,band.fBookmarks[0],0)
					ValidBkIndex = 0
				}
				bt = DxCommonLib.FastString2Byte(fmt.Sprintf("%d",ValidBkIndex))
				band.fcodeEditor.SendEditor(SCI_MARGINSETTEXT,lineNumber,int(uintptr(unsafe.Pointer(&bt[0]))))
				band.fcodeEditor.SendEditor(SCI_MARGINSETSTYLE,lineNumber,band.fBookMarkStyle) //设置书签样式
				band.fBookmarks[ValidBkIndex] = lineNumber
			}else{
				//移除书签
				band.fcodeEditor.SendEditor(SCI_MARGINSETTEXT,lineNumber,0)
				band.fBookmarks[CurBookmark] = -1
			}
		}
	}
}

func newMarginBand(codeEditor *GScintilla)*GDxMarginBand  {
	result := new(GDxMarginBand)
	result.ShowLineNum = true
	result.fBookmarks = [10]int{-1,-1,-1,-1,-1,-1,-1,-1,-1,-1}
	result.ShowBookMark = true
	result.Color = Graphics.ClBtnFace
	result.fcodeEditor = codeEditor
	result.fBookMarkStyle = STYLE_MAX - 1
	result.BookMarkFont.Color = Graphics.RGB(43,43,43)
	result.BookMarkBack = Graphics.RGB(169,183,198)
	result.BookMarkFont.FontName = "宋体"
	result.BookMarkFont.SetSize(9)
	result.BookMarkFont.SetBold(true)

	result.LineNumFont.Color = Graphics.RGB(43,43,43)
	result.LineNumFont.FontName = "宋体"
	result.LineNumFont.SetSize(9)
	result.LineNumFont.Color = Graphics.ClRed
	result.LineNumFont.SetBold(false)
	return result
}
