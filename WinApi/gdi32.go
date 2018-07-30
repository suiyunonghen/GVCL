package WinApi

import (
	"syscall"
	"unsafe"
)

var (
	libgdi32                       syscall.Handle
	fnAbortDoc                     uintptr
	fnAbortPath                    uintptr
	fnAddFontMemResourceEx         uintptr
	fnAddFontResourceW             uintptr
	fnAddFontResourceA             uintptr
	fnAddFontResourceExW           uintptr
	fnAddFontResourceExA           uintptr
	fnAngleArc                     uintptr
	fnAnimatePalette               uintptr
	fnArc                          uintptr
	fnArcTo                        uintptr
	fnBeginPath                    uintptr
	fnBitBlt                       uintptr
	fnCancelDC                     uintptr
	fnCheckColorsInGamut           uintptr
	fnChoosePixelFormat            uintptr
	fnChord                        uintptr
	fnCloseEnhMetaFile             uintptr
	fnCloseFigure                  uintptr
	fnCloseMetaFile                uintptr
	fnColorCorrectPalette          uintptr
	fnColorMatchToTarget           uintptr
	fnCombineRgn                   uintptr
	fnCombineTransform             uintptr
	fnCopyEnhMetaFileW             uintptr
	fnCopyEnhMetaFileA             uintptr
	fnCopyMetaFileW                uintptr
	fnCopyMetaFileA                uintptr
	fnCreateBitmap                 uintptr
	fnCreateBitmapIndirect         uintptr
	fnCreateBrushIndirect          uintptr
	fnCreateColorSpaceW            uintptr
	fnCreateColorSpaceA            uintptr
	fnCreateCompatibleBitmap       uintptr
	fnCreateCompatibleDC           uintptr
	fnCreateDCW                    uintptr
	fnCreateDCA                    uintptr
	fnCreateDIBPatternBrush        uintptr
	fnCreateDIBPatternBrushPt      uintptr
	fnCreateDIBSection             uintptr
	fnCreateDIBitmap               uintptr
	fnCreateDiscardableBitmap      uintptr
	fnCreateEllipticRgn            uintptr
	fnCreateEllipticRgnIndirect    uintptr
	fnCreateEnhMetaFileW           uintptr
	fnCreateEnhMetaFileA           uintptr
	fnCreateFontW                  uintptr
	fnCreateFontA                  uintptr
	fnCreateFontIndirectW          uintptr
	fnCreateFontIndirectA          uintptr
	fnCreateFontIndirectExW        uintptr
	fnCreateFontIndirectExA        uintptr
	fnCreateHalftonePalette        uintptr
	fnCreateHatchBrush             uintptr
	fnCreateICW                    uintptr
	fnCreateICA                    uintptr
	fnCreateMetaFileW              uintptr
	fnCreateMetaFileA              uintptr
	fnCreatePalette                uintptr
	fnCreatePatternBrush           uintptr
	fnCreatePen                    uintptr
	fnCreatePenIndirect            uintptr
	fnCreatePolyPolygonRgn         uintptr
	fnCreatePolygonRgn             uintptr
	fnCreateRectRgn                uintptr
	fnCreateRectRgnIndirect        uintptr
	fnCreateRoundRectRgn           uintptr
	fnCreateScalableFontResourceW  uintptr
	fnCreateScalableFontResourceA  uintptr
	fnCreateSolidBrush             uintptr
	fnDPtoLP                       uintptr
	fnDeleteColorSpace             uintptr
	fnDeleteDC                     uintptr
	fnDeleteEnhMetaFile            uintptr
	fnDeleteMetaFile               uintptr
	fnDeleteObject                 uintptr
	fnDescribePixelFormat          uintptr
	fnDeviceCapabilitiesW          uintptr
	fnDeviceCapabilitiesA          uintptr
	fnDrawEscape                   uintptr
	fnEllipse                      uintptr
	fnEndDoc                       uintptr
	fnEndPage                      uintptr
	fnEndPath                      uintptr
	fnEnumEnhMetaFile              uintptr
	fnEnumFontFamiliesW            uintptr
	fnEnumFontFamiliesA            uintptr
	fnEnumFontFamiliesExW          uintptr
	fnEnumFontFamiliesExA          uintptr
	fnEnumICMProfilesW             uintptr
	fnEnumICMProfilesA             uintptr
	fnEnumMetaFile                 uintptr
	fnEnumObjects                  uintptr
	fnEqualRgn                     uintptr
	fnEscape                       uintptr
	fnExcludeClipRect              uintptr
	fnExtCreatePen                 uintptr
	fnExtCreateRegion              uintptr
	fnExtEscape                    uintptr
	fnExtFloodFill                 uintptr
	fnExtSelectClipRgn             uintptr
	fnFillPath                     uintptr
	fnFillRgn                      uintptr
	fnFlattenPath                  uintptr
	fnFloodFill                    uintptr
	fnFrameRgn                     uintptr
	fnGdiComment                   uintptr
	fnGdiFlush                     uintptr
	fnGdiGetBatchLimit             uintptr
	fnGdiSetBatchLimit             uintptr
	fnGetArcDirection              uintptr
	fnGetAspectRatioFilterEx       uintptr
	fnGetBitmapBits                uintptr
	fnGetBitmapDimensionEx         uintptr
	fnGetBkColor                   uintptr
	fnGetBkMode                    uintptr
	fnGetBoundsRect                uintptr
	fnGetBrushOrgEx                uintptr
	fnGetCharABCWidthsW            uintptr
	fnGetCharABCWidthsA            uintptr
	fnGetCharABCWidthsI            uintptr
	fnGetCharABCWidthsFloatW       uintptr
	fnGetCharABCWidthsFloatA       uintptr
	fnGetCharWidth32W              uintptr
	fnGetCharWidth32A              uintptr
	fnGetCharWidthW                uintptr
	fnGetCharWidthA                uintptr
	fnGetCharWidthFloatW           uintptr
	fnGetCharWidthFloatA           uintptr
	fnGetCharWidthI                uintptr
	fnGetCharacterPlacementW       uintptr
	fnGetCharacterPlacementA       uintptr
	fnGetClipBox                   uintptr
	fnGetClipRgn                   uintptr
	fnGetColorAdjustment           uintptr
	fnGetColorSpace                uintptr
	fnGetCurrentObject             uintptr
	fnGetCurrentPositionEx         uintptr
	fnGetDCBrushColor              uintptr
	fnGetDCPenColor                uintptr
	fnGetDCOrgEx                   uintptr
	fnGetDIBColorTable             uintptr
	fnGetDIBits                    uintptr
	fnGetDeviceCaps                uintptr
	fnGetDeviceGammaRamp           uintptr
	fnGetEnhMetaFileW              uintptr
	fnGetEnhMetaFileA              uintptr
	fnGetEnhMetaFileBits           uintptr
	fnGetEnhMetaFileDescriptionW   uintptr
	fnGetEnhMetaFileDescriptionA   uintptr
	fnGetEnhMetaFileHeader         uintptr
	fnGetEnhMetaFilePaletteEntries uintptr
	fnGetEnhMetaFilePixelFormat    uintptr
	fnGetFontData                  uintptr
	fnGetFontLanguageInfo          uintptr
	fnGetFontUnicodeRanges         uintptr
	fnGetGlyphIndicesW             uintptr
	fnGetGlyphIndicesA             uintptr
	fnGetGlyphOutlineW             uintptr
	fnGetGlyphOutlineA             uintptr
	fnGetGraphicsMode              uintptr
	fnGetICMProfileW               uintptr
	fnGetICMProfileA               uintptr
	fnGetKerningPairs              uintptr
	fnGetLogColorSpaceW            uintptr
	fnGetLogColorSpaceA            uintptr
	fnGetMapMode                   uintptr
	fnGetMetaFileW                 uintptr
	fnGetMetaFileA                 uintptr
	fnGetMetaFileBitsEx            uintptr
	fnGetMetaRgn                   uintptr
	fnGetMiterLimit                uintptr
	fnGetNearestColor              uintptr
	fnGetNearestPaletteIndex       uintptr
	fnGetObjectW                   uintptr
	fnGetObjectA                   uintptr
	fnGetObjectType                uintptr
	fnGetOutlineTextMetricsW       uintptr
	fnGetOutlineTextMetricsA       uintptr
	fnGetPaletteEntries            uintptr
	fnGetPath                      uintptr
	fnGetPixel                     uintptr
	fnGetPixelFormat               uintptr
	fnGetPolyFillMode              uintptr
	fnGetROP2                      uintptr
	fnGetRandomRgn                 uintptr
	fnGetRasterizerCaps            uintptr
	fnGetRegionData                uintptr
	fnGetRgnBox                    uintptr
	fnGetStockObject               uintptr
	fnGetStretchBltMode            uintptr
	fnGetSystemPaletteEntries      uintptr
	fnGetSystemPaletteUse          uintptr
	fnGetTextAlign                 uintptr
	fnGetTextCharacterExtra        uintptr
	fnGetTextCharset               uintptr
	fnGetTextCharsetInfo           uintptr
	fnGetTextColor                 uintptr
	fnGetTextExtentExPointW        uintptr
	fnGetTextExtentExPointA        uintptr
	fnGetTextExtentExPointI        uintptr
	fnGetTextExtentPointW          uintptr
	fnGetTextExtentPointA          uintptr
	fnGetTextExtentPointI          uintptr
	fnGetTextFaceW                 uintptr
	fnGetTextFaceA                 uintptr
	fnGetTextMetricsW              uintptr
	fnGetTextMetricsA              uintptr
	fnGetViewportExtEx             uintptr
	fnGetViewportOrgEx             uintptr
	fnGetWinMetaFileBits           uintptr
	fnGetWindowExtEx               uintptr
	fnGetWindowOrgEx               uintptr
	fnGetWorldTransform            uintptr
	fnIntersectClipRect            uintptr
	fnInvertRgn                    uintptr
	fnLPtoDP                       uintptr
	fnLineDDA                      uintptr
	fnLineTo                       uintptr
	fnMaskBlt                      uintptr
	fnModifyWorldTransform         uintptr
	fnMoveToEx                     uintptr
	fnOffsetClipRgn                uintptr
	fnOffsetRgn                    uintptr
	fnOffsetViewportOrgEx          uintptr
	fnOffsetWindowOrgEx            uintptr
	fnPaintRgn                     uintptr
	fnPatBlt                       uintptr
	fnPathToRegion                 uintptr
	fnPie                          uintptr
	fnPlayEnhMetaFile              uintptr
	fnPlayEnhMetaFileRecord        uintptr
	fnPlayMetaFile                 uintptr
	fnPlgBlt                       uintptr
	fnPolyBezier                   uintptr
	fnPolyBezierTo                 uintptr
	fnPolyDraw                     uintptr
	fnPolyPolygon                  uintptr
	fnPolyPolyline                 uintptr
	fnPolyTextOutW                 uintptr
	fnPolyTextOutA                 uintptr
	fnPolygon                      uintptr
	fnPolyline                     uintptr
	fnPolylineTo                   uintptr
	fnPtInRegion                   uintptr
	fnPtVisible                    uintptr
	fnRealizePalette               uintptr
	fnRectInRegion                 uintptr
	fnRectVisible                  uintptr
	fnRectangle                    uintptr
	fnRemoveFontMemResourceEx      uintptr
	fnRemoveFontResourceW          uintptr
	fnRemoveFontResourceA          uintptr
	fnRemoveFontResourceExW        uintptr
	fnRemoveFontResourceExA        uintptr
	fnResetDCW                     uintptr
	fnResetDCA                     uintptr
	fnResizePalette                uintptr
	fnRestoreDC                    uintptr
	fnRoundRect                    uintptr
	fnSaveDC                       uintptr
	fnScaleViewportExtEx           uintptr
	fnScaleWindowExtEx             uintptr
	fnSelectClipPath               uintptr
	fnSelectClipRgn                uintptr
	fnSelectObject                 uintptr
	fnSelectPalette                uintptr
	fnSetAbortProc                 uintptr
	fnSetArcDirection              uintptr
	fnSetBitmapBits                uintptr
	fnSetBitmapDimensionEx         uintptr
	fnSetBkColor                   uintptr
	fnSetBkMode                    uintptr
	fnSetBoundsRect                uintptr
	fnSetBrushOrgEx                uintptr
	fnSetColorAdjustment           uintptr
	fnSetColorSpace                uintptr
	fnSetDCBrushColor              uintptr
	fnSetDCPenColor                uintptr
	fnSetDIBColorTable             uintptr
	fnSetDIBits                    uintptr
	fnSetDIBitsToDevice            uintptr
	fnSetDeviceGammaRamp           uintptr
	fnSetEnhMetaFileBits           uintptr
	fnSetGraphicsMode              uintptr
	fnSetICMMode                   uintptr
	fnSetICMProfileW               uintptr
	fnSetICMProfileA               uintptr
	fnSetMapMode                   uintptr
	fnSetMapperFlags               uintptr
	fnSetMetaFileBitsEx            uintptr
	fnSetMetaRgn                   uintptr
	fnSetMiterLimit                uintptr
	fnSetPaletteEntries            uintptr
	fnSetPixel                     uintptr
	fnSetPixelFormat               uintptr
	fnSetPixelV                    uintptr
	fnSetPolyFillMode              uintptr
	fnSetROP2                      uintptr
	fnSetRectRgn                   uintptr
	fnSetStretchBltMode            uintptr
	fnSetSystemPaletteUse          uintptr
	fnSetTextAlign                 uintptr
	fnSetTextColor                 uintptr
	fnSetTextCharacterExtra        uintptr
	fnSetTextJustification         uintptr
	fnSetViewportExtEx             uintptr
	fnSetViewportOrgEx             uintptr
	fnSetWinMetaFileBits           uintptr
	fnSetWindowExtEx               uintptr
	fnSetWindowOrgEx               uintptr
	fnSetWorldTransform            uintptr
	fnStartDocW                    uintptr
	fnStartDocA                    uintptr
	fnStartPage                    uintptr
	fnStretchBlt                   uintptr
	fnStretchDIBits                uintptr
	fnStrokeAndFillPath            uintptr
	fnStrokePath                   uintptr
	fnSwapBuffers                  uintptr
	fnTextOutW                     uintptr
	fnTextOutA                     uintptr
	fnTransparentDIBits            uintptr
	fnUnrealizeObject              uintptr
	fnUpdateColors                 uintptr
	fnUpdateICMRegKeyW             uintptr
	fnUpdateICMRegKeyA             uintptr
	fnWidenPath                    uintptr
)

func init() {
	libgdi32, _ = syscall.LoadLibrary("gdi32.dll")
	fnAbortDoc, _ = syscall.GetProcAddress(libgdi32, "AbortDoc")
	fnAbortPath, _ = syscall.GetProcAddress(libgdi32, "AbortPath")
	fnAddFontMemResourceEx, _ = syscall.GetProcAddress(libgdi32, "AddFontMemResourceEx")
	fnAddFontResourceW, _ = syscall.GetProcAddress(libgdi32, "AddFontResourceW")
	fnAddFontResourceA, _ = syscall.GetProcAddress(libgdi32, "AddFontResourceA")
	fnAddFontResourceExW, _ = syscall.GetProcAddress(libgdi32, "AddFontResourceExW")
	fnAddFontResourceExA, _ = syscall.GetProcAddress(libgdi32, "AddFontResourceExA")
	fnAngleArc, _ = syscall.GetProcAddress(libgdi32, "AngleArc")
	fnAnimatePalette, _ = syscall.GetProcAddress(libgdi32, "AnimatePalette")
	fnArc, _ = syscall.GetProcAddress(libgdi32, "Arc")
	fnArcTo, _ = syscall.GetProcAddress(libgdi32, "ArcTo")
	fnBeginPath, _ = syscall.GetProcAddress(libgdi32, "BeginPath")
	fnBitBlt, _ = syscall.GetProcAddress(libgdi32, "BitBlt")
	fnCancelDC, _ = syscall.GetProcAddress(libgdi32, "CancelDC")
	fnCheckColorsInGamut, _ = syscall.GetProcAddress(libgdi32, "CheckColorsInGamut")
	fnChoosePixelFormat, _ = syscall.GetProcAddress(libgdi32, "ChoosePixelFormat")
	fnChord, _ = syscall.GetProcAddress(libgdi32, "Chord")
	fnCloseEnhMetaFile, _ = syscall.GetProcAddress(libgdi32, "CloseEnhMetaFile")
	fnCloseFigure, _ = syscall.GetProcAddress(libgdi32, "CloseFigure")
	fnCloseMetaFile, _ = syscall.GetProcAddress(libgdi32, "CloseMetaFile")
	fnColorCorrectPalette, _ = syscall.GetProcAddress(libgdi32, "ColorCorrectPalette")
	fnColorMatchToTarget, _ = syscall.GetProcAddress(libgdi32, "ColorMatchToTarget")
	fnCombineRgn, _ = syscall.GetProcAddress(libgdi32, "CombineRgn")
	fnCombineTransform, _ = syscall.GetProcAddress(libgdi32, "CombineTransform")
	fnCopyEnhMetaFileW, _ = syscall.GetProcAddress(libgdi32, "CopyEnhMetaFileW")
	fnCopyEnhMetaFileA, _ = syscall.GetProcAddress(libgdi32, "CopyEnhMetaFileA")
	fnCopyMetaFileW, _ = syscall.GetProcAddress(libgdi32, "CopyMetaFileW")
	fnCopyMetaFileA, _ = syscall.GetProcAddress(libgdi32, "CopyMetaFileA")
	fnCreateBitmap, _ = syscall.GetProcAddress(libgdi32, "CreateBitmap")
	fnCreateBitmapIndirect, _ = syscall.GetProcAddress(libgdi32, "CreateBitmapIndirect")
	fnCreateBrushIndirect, _ = syscall.GetProcAddress(libgdi32, "CreateBrushIndirect")
	fnCreateColorSpaceW, _ = syscall.GetProcAddress(libgdi32, "CreateColorSpaceW")
	fnCreateColorSpaceA, _ = syscall.GetProcAddress(libgdi32, "CreateColorSpaceA")
	fnCreateCompatibleBitmap, _ = syscall.GetProcAddress(libgdi32, "CreateCompatibleBitmap")
	fnCreateCompatibleDC, _ = syscall.GetProcAddress(libgdi32, "CreateCompatibleDC")
	fnCreateDCW, _ = syscall.GetProcAddress(libgdi32, "CreateDCW")
	fnCreateDCA, _ = syscall.GetProcAddress(libgdi32, "CreateDCA")
	fnCreateDIBPatternBrush, _ = syscall.GetProcAddress(libgdi32, "CreateDIBPatternBrush")
	fnCreateDIBPatternBrushPt, _ = syscall.GetProcAddress(libgdi32, "CreateDIBPatternBrushPt")
	fnCreateDIBSection, _ = syscall.GetProcAddress(libgdi32, "CreateDIBSection")
	fnCreateDIBitmap, _ = syscall.GetProcAddress(libgdi32, "CreateDIBitmap")
	fnCreateDiscardableBitmap, _ = syscall.GetProcAddress(libgdi32, "CreateDiscardableBitmap")
	fnCreateEllipticRgn, _ = syscall.GetProcAddress(libgdi32, "CreateEllipticRgn")
	fnCreateEllipticRgnIndirect, _ = syscall.GetProcAddress(libgdi32, "CreateEllipticRgnIndirect")
	fnCreateEnhMetaFileW, _ = syscall.GetProcAddress(libgdi32, "CreateEnhMetaFileW")
	fnCreateEnhMetaFileA, _ = syscall.GetProcAddress(libgdi32, "CreateEnhMetaFileA")
	fnCreateFontW, _ = syscall.GetProcAddress(libgdi32, "CreateFontW")
	fnCreateFontA, _ = syscall.GetProcAddress(libgdi32, "CreateFontA")
	fnCreateFontIndirectW, _ = syscall.GetProcAddress(libgdi32, "CreateFontIndirectW")
	fnCreateFontIndirectA, _ = syscall.GetProcAddress(libgdi32, "CreateFontIndirectA")
	fnCreateFontIndirectExW, _ = syscall.GetProcAddress(libgdi32, "CreateFontIndirectExW")
	fnCreateFontIndirectExA, _ = syscall.GetProcAddress(libgdi32, "CreateFontIndirectExA")
	fnCreateHalftonePalette, _ = syscall.GetProcAddress(libgdi32, "CreateHalftonePalette")
	fnCreateHatchBrush, _ = syscall.GetProcAddress(libgdi32, "CreateHatchBrush")
	fnCreateICW, _ = syscall.GetProcAddress(libgdi32, "CreateICW")
	fnCreateICA, _ = syscall.GetProcAddress(libgdi32, "CreateICA")
	fnCreateMetaFileW, _ = syscall.GetProcAddress(libgdi32, "CreateMetaFileW")
	fnCreateMetaFileA, _ = syscall.GetProcAddress(libgdi32, "CreateMetaFileA")
	fnCreatePalette, _ = syscall.GetProcAddress(libgdi32, "CreatePalette")
	fnCreatePatternBrush, _ = syscall.GetProcAddress(libgdi32, "CreatePatternBrush")
	fnCreatePen, _ = syscall.GetProcAddress(libgdi32, "CreatePen")
	fnCreatePenIndirect, _ = syscall.GetProcAddress(libgdi32, "CreatePenIndirect")
	fnCreatePolyPolygonRgn, _ = syscall.GetProcAddress(libgdi32, "CreatePolyPolygonRgn")
	fnCreatePolygonRgn, _ = syscall.GetProcAddress(libgdi32, "CreatePolygonRgn")
	fnCreateRectRgn, _ = syscall.GetProcAddress(libgdi32, "CreateRectRgn")
	fnCreateRectRgnIndirect, _ = syscall.GetProcAddress(libgdi32, "CreateRectRgnIndirect")
	fnCreateRoundRectRgn, _ = syscall.GetProcAddress(libgdi32, "CreateRoundRectRgn")
	fnCreateScalableFontResourceW, _ = syscall.GetProcAddress(libgdi32, "CreateScalableFontResourceW")
	fnCreateScalableFontResourceA, _ = syscall.GetProcAddress(libgdi32, "CreateScalableFontResourceA")
	fnCreateSolidBrush, _ = syscall.GetProcAddress(libgdi32, "CreateSolidBrush")
	fnDPtoLP, _ = syscall.GetProcAddress(libgdi32, "DPtoLP")
	fnDeleteColorSpace, _ = syscall.GetProcAddress(libgdi32, "DeleteColorSpace")
	fnDeleteDC, _ = syscall.GetProcAddress(libgdi32, "DeleteDC")
	fnDeleteEnhMetaFile, _ = syscall.GetProcAddress(libgdi32, "DeleteEnhMetaFile")
	fnDeleteMetaFile, _ = syscall.GetProcAddress(libgdi32, "DeleteMetaFile")
	fnDeleteObject, _ = syscall.GetProcAddress(libgdi32, "DeleteObject")
	fnDescribePixelFormat, _ = syscall.GetProcAddress(libgdi32, "DescribePixelFormat")
	fnDeviceCapabilitiesW, _ = syscall.GetProcAddress(libgdi32, "DeviceCapabilitiesW")
	fnDeviceCapabilitiesA, _ = syscall.GetProcAddress(libgdi32, "DeviceCapabilitiesA")
	fnDrawEscape, _ = syscall.GetProcAddress(libgdi32, "DrawEscape")
	fnEllipse, _ = syscall.GetProcAddress(libgdi32, "Ellipse")
	fnEndDoc, _ = syscall.GetProcAddress(libgdi32, "EndDoc")
	fnEndPage, _ = syscall.GetProcAddress(libgdi32, "EndPage")
	fnEndPath, _ = syscall.GetProcAddress(libgdi32, "EndPath")
	fnEnumEnhMetaFile, _ = syscall.GetProcAddress(libgdi32, "EnumEnhMetaFile")
	fnEnumFontFamiliesW, _ = syscall.GetProcAddress(libgdi32, "EnumFontFamiliesW")
	fnEnumFontFamiliesA, _ = syscall.GetProcAddress(libgdi32, "EnumFontFamiliesA")
	fnEnumFontFamiliesExW, _ = syscall.GetProcAddress(libgdi32, "EnumFontFamiliesExW")
	fnEnumFontFamiliesExA, _ = syscall.GetProcAddress(libgdi32, "EnumFontFamiliesExA")
	fnEnumICMProfilesW, _ = syscall.GetProcAddress(libgdi32, "EnumICMProfilesW")
	fnEnumICMProfilesA, _ = syscall.GetProcAddress(libgdi32, "EnumICMProfilesA")
	fnEnumMetaFile, _ = syscall.GetProcAddress(libgdi32, "EnumMetaFile")
	fnEnumObjects, _ = syscall.GetProcAddress(libgdi32, "EnumObjects")
	fnEqualRgn, _ = syscall.GetProcAddress(libgdi32, "EqualRgn")
	fnEscape, _ = syscall.GetProcAddress(libgdi32, "Escape")
	fnExcludeClipRect, _ = syscall.GetProcAddress(libgdi32, "ExcludeClipRect")
	fnExtCreatePen, _ = syscall.GetProcAddress(libgdi32, "ExtCreatePen")
	fnExtCreateRegion, _ = syscall.GetProcAddress(libgdi32, "ExtCreateRegion")
	fnExtEscape, _ = syscall.GetProcAddress(libgdi32, "ExtEscape")
	fnExtFloodFill, _ = syscall.GetProcAddress(libgdi32, "ExtFloodFill")
	fnExtSelectClipRgn, _ = syscall.GetProcAddress(libgdi32, "ExtSelectClipRgn")
	fnFillPath, _ = syscall.GetProcAddress(libgdi32, "FillPath")
	fnFillRgn, _ = syscall.GetProcAddress(libgdi32, "FillRgn")
	fnFlattenPath, _ = syscall.GetProcAddress(libgdi32, "FlattenPath")
	fnFloodFill, _ = syscall.GetProcAddress(libgdi32, "FloodFill")
	fnFrameRgn, _ = syscall.GetProcAddress(libgdi32, "FrameRgn")
	fnGdiComment, _ = syscall.GetProcAddress(libgdi32, "GdiComment")
	fnGdiFlush, _ = syscall.GetProcAddress(libgdi32, "GdiFlush")
	fnGdiGetBatchLimit, _ = syscall.GetProcAddress(libgdi32, "GdiGetBatchLimit")
	fnGdiSetBatchLimit, _ = syscall.GetProcAddress(libgdi32, "GdiSetBatchLimit")
	fnGetArcDirection, _ = syscall.GetProcAddress(libgdi32, "GetArcDirection")
	fnGetAspectRatioFilterEx, _ = syscall.GetProcAddress(libgdi32, "GetAspectRatioFilterEx")
	fnGetBitmapBits, _ = syscall.GetProcAddress(libgdi32, "GetBitmapBits")
	fnGetBitmapDimensionEx, _ = syscall.GetProcAddress(libgdi32, "GetBitmapDimensionEx")
	fnGetBkColor, _ = syscall.GetProcAddress(libgdi32, "GetBkColor")
	fnGetBkMode, _ = syscall.GetProcAddress(libgdi32, "GetBkMode")
	fnGetBoundsRect, _ = syscall.GetProcAddress(libgdi32, "GetBoundsRect")
	fnGetBrushOrgEx, _ = syscall.GetProcAddress(libgdi32, "GetBrushOrgEx")
	fnGetCharABCWidthsW, _ = syscall.GetProcAddress(libgdi32, "GetCharABCWidthsW")
	fnGetCharABCWidthsA, _ = syscall.GetProcAddress(libgdi32, "GetCharABCWidthsA")
	fnGetCharABCWidthsI, _ = syscall.GetProcAddress(libgdi32, "GetCharABCWidthsI")
	fnGetCharABCWidthsFloatW, _ = syscall.GetProcAddress(libgdi32, "GetCharABCWidthsFloatW")
	fnGetCharABCWidthsFloatA, _ = syscall.GetProcAddress(libgdi32, "GetCharABCWidthsFloatA")
	fnGetCharWidth32W, _ = syscall.GetProcAddress(libgdi32, "GetCharWidth32W")
	fnGetCharWidth32A, _ = syscall.GetProcAddress(libgdi32, "GetCharWidth32A")
	fnGetCharWidthW, _ = syscall.GetProcAddress(libgdi32, "GetCharWidthW")
	fnGetCharWidthA, _ = syscall.GetProcAddress(libgdi32, "GetCharWidthA")
	fnGetCharWidthFloatW, _ = syscall.GetProcAddress(libgdi32, "GetCharWidthFloatW")
	fnGetCharWidthFloatA, _ = syscall.GetProcAddress(libgdi32, "GetCharWidthFloatA")
	fnGetCharWidthI, _ = syscall.GetProcAddress(libgdi32, "GetCharWidthI")
	fnGetCharacterPlacementW, _ = syscall.GetProcAddress(libgdi32, "GetCharacterPlacementW")
	fnGetCharacterPlacementA, _ = syscall.GetProcAddress(libgdi32, "GetCharacterPlacementA")
	fnGetClipBox, _ = syscall.GetProcAddress(libgdi32, "GetClipBox")
	fnGetClipRgn, _ = syscall.GetProcAddress(libgdi32, "GetClipRgn")
	fnGetColorAdjustment, _ = syscall.GetProcAddress(libgdi32, "GetColorAdjustment")
	fnGetColorSpace, _ = syscall.GetProcAddress(libgdi32, "GetColorSpace")
	fnGetCurrentObject, _ = syscall.GetProcAddress(libgdi32, "GetCurrentObject")
	fnGetCurrentPositionEx, _ = syscall.GetProcAddress(libgdi32, "GetCurrentPositionEx")
	fnGetDCBrushColor, _ = syscall.GetProcAddress(libgdi32, "GetDCBrushColor")
	fnGetDCPenColor, _ = syscall.GetProcAddress(libgdi32, "GetDCPenColor")
	fnGetDCOrgEx, _ = syscall.GetProcAddress(libgdi32, "GetDCOrgEx")
	fnGetDIBColorTable, _ = syscall.GetProcAddress(libgdi32, "GetDIBColorTable")
	fnGetDIBits, _ = syscall.GetProcAddress(libgdi32, "GetDIBits")
	fnGetDeviceCaps, _ = syscall.GetProcAddress(libgdi32, "GetDeviceCaps")
	fnGetDeviceGammaRamp, _ = syscall.GetProcAddress(libgdi32, "GetDeviceGammaRamp")
	fnGetEnhMetaFileW, _ = syscall.GetProcAddress(libgdi32, "GetEnhMetaFileW")
	fnGetEnhMetaFileA, _ = syscall.GetProcAddress(libgdi32, "GetEnhMetaFileA")
	fnGetEnhMetaFileBits, _ = syscall.GetProcAddress(libgdi32, "GetEnhMetaFileBits")
	fnGetEnhMetaFileDescriptionW, _ = syscall.GetProcAddress(libgdi32, "GetEnhMetaFileDescriptionW")
	fnGetEnhMetaFileDescriptionA, _ = syscall.GetProcAddress(libgdi32, "GetEnhMetaFileDescriptionA")
	fnGetEnhMetaFileHeader, _ = syscall.GetProcAddress(libgdi32, "GetEnhMetaFileHeader")
	fnGetEnhMetaFilePaletteEntries, _ = syscall.GetProcAddress(libgdi32, "GetEnhMetaFilePaletteEntries")
	fnGetEnhMetaFilePixelFormat, _ = syscall.GetProcAddress(libgdi32, "GetEnhMetaFilePixelFormat")
	fnGetFontData, _ = syscall.GetProcAddress(libgdi32, "GetFontData")
	fnGetFontLanguageInfo, _ = syscall.GetProcAddress(libgdi32, "GetFontLanguageInfo")
	fnGetFontUnicodeRanges, _ = syscall.GetProcAddress(libgdi32, "GetFontUnicodeRanges")
	fnGetGlyphIndicesW, _ = syscall.GetProcAddress(libgdi32, "GetGlyphIndicesW")
	fnGetGlyphIndicesA, _ = syscall.GetProcAddress(libgdi32, "GetGlyphIndicesA")
	fnGetGlyphOutlineW, _ = syscall.GetProcAddress(libgdi32, "GetGlyphOutlineW")
	fnGetGlyphOutlineA, _ = syscall.GetProcAddress(libgdi32, "GetGlyphOutlineA")
	fnGetGraphicsMode, _ = syscall.GetProcAddress(libgdi32, "GetGraphicsMode")
	fnGetICMProfileW, _ = syscall.GetProcAddress(libgdi32, "GetICMProfileW")
	fnGetICMProfileA, _ = syscall.GetProcAddress(libgdi32, "GetICMProfileA")
	fnGetKerningPairs, _ = syscall.GetProcAddress(libgdi32, "GetKerningPairs")
	fnGetLogColorSpaceW, _ = syscall.GetProcAddress(libgdi32, "GetLogColorSpaceW")
	fnGetLogColorSpaceA, _ = syscall.GetProcAddress(libgdi32, "GetLogColorSpaceA")
	fnGetMapMode, _ = syscall.GetProcAddress(libgdi32, "GetMapMode")
	fnGetMetaFileW, _ = syscall.GetProcAddress(libgdi32, "GetMetaFileW")
	fnGetMetaFileA, _ = syscall.GetProcAddress(libgdi32, "GetMetaFileA")
	fnGetMetaFileBitsEx, _ = syscall.GetProcAddress(libgdi32, "GetMetaFileBitsEx")
	fnGetMetaRgn, _ = syscall.GetProcAddress(libgdi32, "GetMetaRgn")
	fnGetMiterLimit, _ = syscall.GetProcAddress(libgdi32, "GetMiterLimit")
	fnGetNearestColor, _ = syscall.GetProcAddress(libgdi32, "GetNearestColor")
	fnGetNearestPaletteIndex, _ = syscall.GetProcAddress(libgdi32, "GetNearestPaletteIndex")
	fnGetObjectW, _ = syscall.GetProcAddress(libgdi32, "GetObjectW")
	fnGetObjectA, _ = syscall.GetProcAddress(libgdi32, "GetObjectA")
	fnGetObjectType, _ = syscall.GetProcAddress(libgdi32, "GetObjectType")
	fnGetOutlineTextMetricsW, _ = syscall.GetProcAddress(libgdi32, "GetOutlineTextMetricsW")
	fnGetOutlineTextMetricsA, _ = syscall.GetProcAddress(libgdi32, "GetOutlineTextMetricsA")
	fnGetPaletteEntries, _ = syscall.GetProcAddress(libgdi32, "GetPaletteEntries")
	fnGetPath, _ = syscall.GetProcAddress(libgdi32, "GetPath")
	fnGetPixel, _ = syscall.GetProcAddress(libgdi32, "GetPixel")
	fnGetPixelFormat, _ = syscall.GetProcAddress(libgdi32, "GetPixelFormat")
	fnGetPolyFillMode, _ = syscall.GetProcAddress(libgdi32, "GetPolyFillMode")
	fnGetROP2, _ = syscall.GetProcAddress(libgdi32, "GetROP2")
	fnGetRandomRgn, _ = syscall.GetProcAddress(libgdi32, "GetRandomRgn")
	fnGetRasterizerCaps, _ = syscall.GetProcAddress(libgdi32, "GetRasterizerCaps")
	fnGetRegionData, _ = syscall.GetProcAddress(libgdi32, "GetRegionData")
	fnGetRgnBox, _ = syscall.GetProcAddress(libgdi32, "GetRgnBox")
	fnGetStockObject, _ = syscall.GetProcAddress(libgdi32, "GetStockObject")
	fnGetStretchBltMode, _ = syscall.GetProcAddress(libgdi32, "GetStretchBltMode")
	fnGetSystemPaletteEntries, _ = syscall.GetProcAddress(libgdi32, "GetSystemPaletteEntries")
	fnGetSystemPaletteUse, _ = syscall.GetProcAddress(libgdi32, "GetSystemPaletteUse")
	fnGetTextAlign, _ = syscall.GetProcAddress(libgdi32, "GetTextAlign")
	fnGetTextCharacterExtra, _ = syscall.GetProcAddress(libgdi32, "GetTextCharacterExtra")
	fnGetTextCharset, _ = syscall.GetProcAddress(libgdi32, "GetTextCharset")
	fnGetTextCharsetInfo, _ = syscall.GetProcAddress(libgdi32, "GetTextCharsetInfo")
	fnGetTextColor, _ = syscall.GetProcAddress(libgdi32, "GetTextColor")
	fnGetTextExtentExPointW, _ = syscall.GetProcAddress(libgdi32, "GetTextExtentExPointW")
	fnGetTextExtentExPointA, _ = syscall.GetProcAddress(libgdi32, "GetTextExtentExPointA")
	fnGetTextExtentExPointI, _ = syscall.GetProcAddress(libgdi32, "GetTextExtentExPointI")
	fnGetTextExtentPointW, _ = syscall.GetProcAddress(libgdi32, "GetTextExtentPointW")
	fnGetTextExtentPointA, _ = syscall.GetProcAddress(libgdi32, "GetTextExtentPointA")
	fnGetTextExtentPointI, _ = syscall.GetProcAddress(libgdi32, "GetTextExtentPointI")
	fnGetTextFaceW, _ = syscall.GetProcAddress(libgdi32, "GetTextFaceW")
	fnGetTextFaceA, _ = syscall.GetProcAddress(libgdi32, "GetTextFaceA")
	fnGetTextMetricsW, _ = syscall.GetProcAddress(libgdi32, "GetTextMetricsW")
	fnGetTextMetricsA, _ = syscall.GetProcAddress(libgdi32, "GetTextMetricsA")
	fnGetViewportExtEx, _ = syscall.GetProcAddress(libgdi32, "GetViewportExtEx")
	fnGetViewportOrgEx, _ = syscall.GetProcAddress(libgdi32, "GetViewportOrgEx")
	fnGetWinMetaFileBits, _ = syscall.GetProcAddress(libgdi32, "GetWinMetaFileBits")
	fnGetWindowExtEx, _ = syscall.GetProcAddress(libgdi32, "GetWindowExtEx")
	fnGetWindowOrgEx, _ = syscall.GetProcAddress(libgdi32, "GetWindowOrgEx")
	fnGetWorldTransform, _ = syscall.GetProcAddress(libgdi32, "GetWorldTransform")
	fnIntersectClipRect, _ = syscall.GetProcAddress(libgdi32, "IntersectClipRect")
	fnInvertRgn, _ = syscall.GetProcAddress(libgdi32, "InvertRgn")
	fnLPtoDP, _ = syscall.GetProcAddress(libgdi32, "LPtoDP")
	fnLineDDA, _ = syscall.GetProcAddress(libgdi32, "LineDDA")
	fnLineTo, _ = syscall.GetProcAddress(libgdi32, "LineTo")
	fnMaskBlt, _ = syscall.GetProcAddress(libgdi32, "MaskBlt")
	fnModifyWorldTransform, _ = syscall.GetProcAddress(libgdi32, "ModifyWorldTransform")
	fnMoveToEx, _ = syscall.GetProcAddress(libgdi32, "MoveToEx")
	fnOffsetClipRgn, _ = syscall.GetProcAddress(libgdi32, "OffsetClipRgn")
	fnOffsetRgn, _ = syscall.GetProcAddress(libgdi32, "OffsetRgn")
	fnOffsetViewportOrgEx, _ = syscall.GetProcAddress(libgdi32, "OffsetViewportOrgEx")
	fnOffsetWindowOrgEx, _ = syscall.GetProcAddress(libgdi32, "OffsetWindowOrgEx")
	fnPaintRgn, _ = syscall.GetProcAddress(libgdi32, "PaintRgn")
	fnPatBlt, _ = syscall.GetProcAddress(libgdi32, "PatBlt")
	fnPathToRegion, _ = syscall.GetProcAddress(libgdi32, "PathToRegion")
	fnPie, _ = syscall.GetProcAddress(libgdi32, "Pie")
	fnPlayEnhMetaFile, _ = syscall.GetProcAddress(libgdi32, "PlayEnhMetaFile")
	fnPlayEnhMetaFileRecord, _ = syscall.GetProcAddress(libgdi32, "PlayEnhMetaFileRecord")
	fnPlayMetaFile, _ = syscall.GetProcAddress(libgdi32, "PlayMetaFile")
	fnPlgBlt, _ = syscall.GetProcAddress(libgdi32, "PlgBlt")
	fnPolyBezier, _ = syscall.GetProcAddress(libgdi32, "PolyBezier")
	fnPolyBezierTo, _ = syscall.GetProcAddress(libgdi32, "PolyBezierTo")
	fnPolyDraw, _ = syscall.GetProcAddress(libgdi32, "PolyDraw")
	fnPolyPolygon, _ = syscall.GetProcAddress(libgdi32, "PolyPolygon")
	fnPolyPolyline, _ = syscall.GetProcAddress(libgdi32, "PolyPolyline")
	fnPolyTextOutW, _ = syscall.GetProcAddress(libgdi32, "PolyTextOutW")
	fnPolyTextOutA, _ = syscall.GetProcAddress(libgdi32, "PolyTextOutA")
	fnPolygon, _ = syscall.GetProcAddress(libgdi32, "Polygon")
	fnPolyline, _ = syscall.GetProcAddress(libgdi32, "Polyline")
	fnPolylineTo, _ = syscall.GetProcAddress(libgdi32, "PolylineTo")
	fnPtInRegion, _ = syscall.GetProcAddress(libgdi32, "PtInRegion")
	fnPtVisible, _ = syscall.GetProcAddress(libgdi32, "PtVisible")
	fnRealizePalette, _ = syscall.GetProcAddress(libgdi32, "RealizePalette")
	fnRectInRegion, _ = syscall.GetProcAddress(libgdi32, "RectInRegion")
	fnRectVisible, _ = syscall.GetProcAddress(libgdi32, "RectVisible")
	fnRectangle, _ = syscall.GetProcAddress(libgdi32, "Rectangle")
	fnRemoveFontMemResourceEx, _ = syscall.GetProcAddress(libgdi32, "RemoveFontMemResourceEx")
	fnRemoveFontResourceW, _ = syscall.GetProcAddress(libgdi32, "RemoveFontResourceW")
	fnRemoveFontResourceA, _ = syscall.GetProcAddress(libgdi32, "RemoveFontResourceA")
	fnRemoveFontResourceExW, _ = syscall.GetProcAddress(libgdi32, "RemoveFontResourceExW")
	fnRemoveFontResourceExA, _ = syscall.GetProcAddress(libgdi32, "RemoveFontResourceExA")
	fnResetDCW, _ = syscall.GetProcAddress(libgdi32, "ResetDCW")
	fnResetDCA, _ = syscall.GetProcAddress(libgdi32, "ResetDCA")
	fnResizePalette, _ = syscall.GetProcAddress(libgdi32, "ResizePalette")
	fnRestoreDC, _ = syscall.GetProcAddress(libgdi32, "RestoreDC")
	fnRoundRect, _ = syscall.GetProcAddress(libgdi32, "RoundRect")
	fnSaveDC, _ = syscall.GetProcAddress(libgdi32, "SaveDC")
	fnScaleViewportExtEx, _ = syscall.GetProcAddress(libgdi32, "ScaleViewportExtEx")
	fnScaleWindowExtEx, _ = syscall.GetProcAddress(libgdi32, "ScaleWindowExtEx")
	fnSelectClipPath, _ = syscall.GetProcAddress(libgdi32, "SelectClipPath")
	fnSelectClipRgn, _ = syscall.GetProcAddress(libgdi32, "SelectClipRgn")
	fnSelectObject, _ = syscall.GetProcAddress(libgdi32, "SelectObject")
	fnSelectPalette, _ = syscall.GetProcAddress(libgdi32, "SelectPalette")
	fnSetAbortProc, _ = syscall.GetProcAddress(libgdi32, "SetAbortProc")
	fnSetArcDirection, _ = syscall.GetProcAddress(libgdi32, "SetArcDirection")
	fnSetBitmapBits, _ = syscall.GetProcAddress(libgdi32, "SetBitmapBits")
	fnSetBitmapDimensionEx, _ = syscall.GetProcAddress(libgdi32, "SetBitmapDimensionEx")
	fnSetBkColor, _ = syscall.GetProcAddress(libgdi32, "SetBkColor")
	fnSetBkMode, _ = syscall.GetProcAddress(libgdi32, "SetBkMode")
	fnSetBoundsRect, _ = syscall.GetProcAddress(libgdi32, "SetBoundsRect")
	fnSetBrushOrgEx, _ = syscall.GetProcAddress(libgdi32, "SetBrushOrgEx")
	fnSetColorAdjustment, _ = syscall.GetProcAddress(libgdi32, "SetColorAdjustment")
	fnSetColorSpace, _ = syscall.GetProcAddress(libgdi32, "SetColorSpace")
	fnSetDCBrushColor, _ = syscall.GetProcAddress(libgdi32, "SetDCBrushColor")
	fnSetDCPenColor, _ = syscall.GetProcAddress(libgdi32, "SetDCPenColor")
	fnSetDIBColorTable, _ = syscall.GetProcAddress(libgdi32, "SetDIBColorTable")
	fnSetDIBits, _ = syscall.GetProcAddress(libgdi32, "SetDIBits")
	fnSetDIBitsToDevice, _ = syscall.GetProcAddress(libgdi32, "SetDIBitsToDevice")
	fnSetDeviceGammaRamp, _ = syscall.GetProcAddress(libgdi32, "SetDeviceGammaRamp")
	fnSetEnhMetaFileBits, _ = syscall.GetProcAddress(libgdi32, "SetEnhMetaFileBits")
	fnSetGraphicsMode, _ = syscall.GetProcAddress(libgdi32, "SetGraphicsMode")
	fnSetICMMode, _ = syscall.GetProcAddress(libgdi32, "SetICMMode")
	fnSetICMProfileW, _ = syscall.GetProcAddress(libgdi32, "SetICMProfileW")
	fnSetICMProfileA, _ = syscall.GetProcAddress(libgdi32, "SetICMProfileA")
	fnSetMapMode, _ = syscall.GetProcAddress(libgdi32, "SetMapMode")
	fnSetMapperFlags, _ = syscall.GetProcAddress(libgdi32, "SetMapperFlags")
	fnSetMetaFileBitsEx, _ = syscall.GetProcAddress(libgdi32, "SetMetaFileBitsEx")
	fnSetMetaRgn, _ = syscall.GetProcAddress(libgdi32, "SetMetaRgn")
	fnSetMiterLimit, _ = syscall.GetProcAddress(libgdi32, "SetMiterLimit")
	fnSetPaletteEntries, _ = syscall.GetProcAddress(libgdi32, "SetPaletteEntries")
	fnSetPixel, _ = syscall.GetProcAddress(libgdi32, "SetPixel")
	fnSetPixelFormat, _ = syscall.GetProcAddress(libgdi32, "SetPixelFormat")
	fnSetPixelV, _ = syscall.GetProcAddress(libgdi32, "SetPixelV")
	fnSetPolyFillMode, _ = syscall.GetProcAddress(libgdi32, "SetPolyFillMode")
	fnSetROP2, _ = syscall.GetProcAddress(libgdi32, "SetROP2")
	fnSetRectRgn, _ = syscall.GetProcAddress(libgdi32, "SetRectRgn")
	fnSetStretchBltMode, _ = syscall.GetProcAddress(libgdi32, "SetStretchBltMode")
	fnSetSystemPaletteUse, _ = syscall.GetProcAddress(libgdi32, "SetSystemPaletteUse")
	fnSetTextAlign, _ = syscall.GetProcAddress(libgdi32, "SetTextAlign")
	fnSetTextColor, _ = syscall.GetProcAddress(libgdi32, "SetTextColor")
	fnSetTextCharacterExtra, _ = syscall.GetProcAddress(libgdi32, "SetTextCharacterExtra")
	fnSetTextJustification, _ = syscall.GetProcAddress(libgdi32, "SetTextJustification")
	fnSetViewportExtEx, _ = syscall.GetProcAddress(libgdi32, "SetViewportExtEx")
	fnSetViewportOrgEx, _ = syscall.GetProcAddress(libgdi32, "SetViewportOrgEx")
	fnSetWinMetaFileBits, _ = syscall.GetProcAddress(libgdi32, "SetWinMetaFileBits")
	fnSetWindowExtEx, _ = syscall.GetProcAddress(libgdi32, "SetWindowExtEx")
	fnSetWindowOrgEx, _ = syscall.GetProcAddress(libgdi32, "SetWindowOrgEx")
	fnSetWorldTransform, _ = syscall.GetProcAddress(libgdi32, "SetWorldTransform")
	fnStartDocW, _ = syscall.GetProcAddress(libgdi32, "StartDocW")
	fnStartDocA, _ = syscall.GetProcAddress(libgdi32, "StartDocA")
	fnStartPage, _ = syscall.GetProcAddress(libgdi32, "StartPage")
	fnStretchBlt, _ = syscall.GetProcAddress(libgdi32, "StretchBlt")
	fnStretchDIBits, _ = syscall.GetProcAddress(libgdi32, "StretchDIBits")
	fnStrokeAndFillPath, _ = syscall.GetProcAddress(libgdi32, "StrokeAndFillPath")
	fnStrokePath, _ = syscall.GetProcAddress(libgdi32, "StrokePath")
	fnSwapBuffers, _ = syscall.GetProcAddress(libgdi32, "SwapBuffers")
	fnTextOutW, _ = syscall.GetProcAddress(libgdi32, "TextOutW")
	fnTextOutA, _ = syscall.GetProcAddress(libgdi32, "TextOutA")
	fnTransparentDIBits, _ = syscall.GetProcAddress(libgdi32, "TransparentDIBits")
	fnUnrealizeObject, _ = syscall.GetProcAddress(libgdi32, "UnrealizeObject")
	fnUpdateColors, _ = syscall.GetProcAddress(libgdi32, "UpdateColors")
	fnUpdateICMRegKeyW, _ = syscall.GetProcAddress(libgdi32, "UpdateICMRegKeyW")
	fnUpdateICMRegKeyA, _ = syscall.GetProcAddress(libgdi32, "UpdateICMRegKeyA")
	fnWidenPath, _ = syscall.GetProcAddress(libgdi32, "WidenPath")
}

func GetCurrentPositionEx(dc HDC,pt *POINT)bool{
	ret,_,_ := syscall.Syscall(fnGetCurrentPositionEx,2,uintptr(dc),uintptr(unsafe.Pointer(pt)),0)
	return ret !=0
}

func MoveToEx(dc HDC,p1,p2 int32,pt *POINT)bool  {
	ret,_,_ := syscall.Syscall6(fnMoveToEx,4,uintptr(dc),uintptr(p1),uintptr(p2),uintptr(unsafe.Pointer(pt)),0,0)
	return ret !=0
}

func GetClipBox(dc HDC,rect *Rect)int  {
	ret,_,_ := syscall.Syscall(fnGetClipBox,2,uintptr(dc),uintptr(unsafe.Pointer(rect)),0)
	return int(ret)
}

func SelectObject(dc HDC,gdiobj uintptr) uintptr {
	ret,_,_ := syscall.Syscall(fnSelectObject,2,uintptr(dc),gdiobj,0)
	return ret
}

func LineTo(dc HDC,x,y int) bool {
	ret,_,_ := syscall.Syscall(fnLineTo,3,uintptr(dc),uintptr(x),uintptr(y))
	return ret !=0
}

func IntersectClipRect(dc HDC, X1, Y1, X2, Y2 int32) int  {
	ret,_,_ := syscall.Syscall6(fnIntersectClipRect,5,uintptr(dc),uintptr(X1),uintptr(Y1),uintptr(X2),uintptr(Y2),0)
	return int(ret)
}

func SetViewportOrgEx(dc HDC, X, Y int, Point *POINT)bool  {
	ret,_,_ := syscall.Syscall6(fnIntersectClipRect,4,uintptr(dc),uintptr(X),uintptr(Y),uintptr(unsafe.Pointer(Point)),0,0)
	return ret != 0
}

func UnrealizeObject(gdihandle uintptr)bool  {
	ret,_,_ := syscall.Syscall(fnUnrealizeObject,1,gdihandle,0,0)
	return ret !=0
}

func SetBkColor(dc HDC,color uint32) uint32 {
	ret,_,_ := syscall.Syscall(fnSetBkColor,2,uintptr(dc),uintptr(color),0)
	return uint32(ret)
}

func SetBkMode(dc HDC,BkMode int32) int32 {
	ret,_,_ := syscall.Syscall(fnSetBkMode,2,uintptr(dc),uintptr(BkMode),0)
	return int32(ret)
}

func SetTextColor(dc HDC,color uint32) uint32 {
	ret,_,_ := syscall.Syscall(fnSetTextColor,2,uintptr(dc),uintptr(color),0)
	return uint32(ret)
}

func SetWindowOrgEx(dc HDC,x,y int,oldpt *POINT)bool  {
	ret,_,_ := syscall.Syscall6(fnSetWindowOrgEx,4,uintptr(dc),uintptr(x),uintptr(y),uintptr(unsafe.Pointer(oldpt)),0,0)
	return ret != 0
}

func GetWindowOrgEx(dc HDC,pt *POINT)bool  {
	ret,_,_ := syscall.Syscall(fnGetWindowOrgEx,2,uintptr(dc),uintptr(unsafe.Pointer(pt)),0)
	return ret !=0
}

func MoveWindowOrg(dc HDC,dx,dy int32)  {
	p := new(POINT)
	syscall.Syscall(fnGetWindowOrgEx,2,uintptr(dc),uintptr(unsafe.Pointer(p)),0)
	syscall.Syscall6(fnSetWindowOrgEx,4,uintptr(dc),uintptr(p.X - dx),uintptr(p.Y-dy),0,0,0)
}

func RectVisible(dc HDC,rect *Rect)bool  {
	ret,_,_ := syscall.Syscall(fnRectVisible,2,uintptr(dc),uintptr(unsafe.Pointer(rect)),0)
	return ret !=0
}

func BitBlt(DestDC HDC,X, Y, Width, Height int,SrcDC HDC,XSrc, YSrc int,Rop uint32)bool {
	ret,_,_ := syscall.Syscall9(fnBitBlt,9,uintptr(DestDC),uintptr(X),uintptr(Y),uintptr(Width),
		uintptr(Height),uintptr(SrcDC),uintptr(XSrc),uintptr(YSrc),uintptr(Rop))
	return ret !=0
}


func CreateCompatibleDC(dc HDC)HDC  {
	ret,_,_ := syscall.Syscall(fnCreateCompatibleDC,1,uintptr(dc),0,0)
	return HDC(ret)
}

func DeleteDC(dc HDC)bool  {
	ret,_,_ := syscall.Syscall(fnDeleteDC,1,uintptr(dc),0,0)
	return ret !=0
}

func DeleteObject(p HGDIOBJ)bool  {
	ret,_,_ := syscall.Syscall(fnDeleteObject,1,uintptr(p),0,0)
	return ret !=0
}

/*function SetDIBitsToDevice(DC: HDC; DestX, DestY: Integer; Width, Height: DWORD;
SrcX, SrcY: Integer; nStartScan, NumScans: UINT; Bits: Pointer;
var BitsInfo: TBitmapInfo; Usage: UINT): Integer; stdcall;*/

func SetDIBitsToDevice(dc HDC,DestX, DestY int,Width, Height uint32,SrcX, SrcY int,nStartScan, NumScans uint,bits uintptr,BitsInfo *BITMAPINFO,usage uint)int  {
	ret,_,_ := syscall.Syscall12(fnSetDIBitsToDevice,12,uintptr(dc),
		uintptr(DestX),uintptr(DestY),uintptr(Width),uintptr(Height),
		uintptr(SrcX),uintptr(SrcY),uintptr(nStartScan),uintptr(NumScans),bits,
			uintptr(unsafe.Pointer(BitsInfo)),uintptr(usage))
	return int(ret)
}

func SetMapMode(dc HDC,p2 int)int  {
	ret,_,_ := syscall.Syscall(fnSetMapMode,2,uintptr(dc),uintptr(p2),0)
	return int(ret)
}

func SetGraphicsMode(dc HDC,imode int)int  {
	ret,_,_ := syscall.Syscall(fnSetGraphicsMode,2,uintptr(dc),uintptr(imode),0)
	return int(ret)
}

func CreateDIBSection(dc HDC,p2 *BITMAPINFO,p3 uint,p4 *uintptr,p5 syscall.Handle,p6 uint32) HBITMAP  {
	ret,_,_ := syscall.Syscall6(fnCreateDIBSection,6,uintptr(dc),uintptr(unsafe.Pointer(p2)),uintptr(p3),uintptr(unsafe.Pointer(p4)),uintptr(p5),uintptr(p6))
	return HBITMAP(ret)
}


func CreatePalette(LogPalette *LOGPALETTE)HPALETTE  {
	ret,_,_ := syscall.Syscall(fnCreatePalette,1,uintptr(unsafe.Pointer(LogPalette)),0,0)
	return HPALETTE(ret)
}

func SelectPalette(dc HDC,palete HPALETTE,ForceBackground bool) HPALETTE  {
	var ret uintptr
	if ForceBackground{
		ret,_,_ = syscall.Syscall(fnSelectPalette,3,uintptr(dc),uintptr(palete),1)
	}else{
		ret,_,_ = syscall.Syscall(fnSelectPalette,3,uintptr(dc),uintptr(palete),0)
	}

	return HPALETTE(ret)
}

func RealizePalette(dc HDC) uint  {
	ret,_,_ := syscall.Syscall(fnRealizePalette,1,uintptr(dc),0,0)
	return uint(ret)
}


func GetDIBColorTable(dc HDC,p2,p3 uint, RGBQuadStructs *uintptr)  uint{
	ret,_,_ := syscall.Syscall6(fnGetDIBColorTable,4,uintptr(dc),uintptr(p2),
		uintptr(p3),uintptr(unsafe.Pointer(RGBQuadStructs)),0,0,)
	return uint(ret)
}

func PatBlt(dc HDC,x,y,width,height int,Rop uint32)bool  {
	ret,_,_ := syscall.Syscall6(fnPatBlt,6,uintptr(dc),uintptr(x),
		uintptr(y),uintptr(width),uintptr(height),uintptr(Rop))
	return ret != 0
}