@REM @Author: anchen
@REM @Date:   2016-09-14 09:58:47
@REM @Last Modified by:   anchen
@REM Modified time: 2016-09-14 11:50:23
@REM rsrc.exe -ico="dota.ico,Dota2.ico"  -manifest msgbox.manifest  -bmp GVCL.bmp -o msgbox.syso
@REM syso -o msgbox.syso
rsrc.exe -ico="dota.ico,Dota2.ico"  -manifest msgbox.manifest  -bmp GVCL.bmp -o msgbox.syso
go build -ldflags="-H windowsgui"
@REM go build