package WinApi

import (
	"syscall"
	"unsafe"
)

var (
	libkernel32                         syscall.Handle
	fnAddAtomW                          uintptr
	fnAddAtomA                          uintptr
	fnAllocConsole                      uintptr
	fnAreFileApisANSI                   uintptr
	fnBackupRead                        uintptr
	fnBackupSeek                        uintptr
	fnBackupWrite                       uintptr
	fnBeep                              uintptr
	fnBeginUpdateResourceW              uintptr
	fnBeginUpdateResourceA              uintptr
	fnBindIoCompletionCallback          uintptr
	fnBuildCommDCBW                     uintptr
	fnBuildCommDCBA                     uintptr
	fnBuildCommDCBAndTimeoutsW          uintptr
	fnBuildCommDCBAndTimeoutsA          uintptr
	fnCallNamedPipeW                    uintptr
	fnCallNamedPipeA                    uintptr
	fnCancelIo                          uintptr
	fnCancelTimerQueueTimer             uintptr
	fnCancelWaitableTimer               uintptr
	fnChangeTimerQueueTimer             uintptr
	fnCheckNameLegalDOS8Dot3W           uintptr
	fnCheckNameLegalDOS8Dot3A           uintptr
	fnCheckRemoteDebuggerPresent        uintptr
	fnClearCommBreak                    uintptr
	fnClearCommError                    uintptr
	fnCommConfigDialogW                 uintptr
	fnCommConfigDialogA                 uintptr
	fnCompareFileTime                   uintptr
	fnConnectNamedPipe                  uintptr
	fnContinueDebugEvent                uintptr
	fnConvertThreadToFiber              uintptr
	fnConvertFiberToThread              uintptr
	fnConvertDefaultLocale              uintptr
	fnCopyFileTransactedW               uintptr
	fnCopyFileTransactedA               uintptr
	fnReplaceFileW                      uintptr
	fnReplaceFileA                      uintptr
	fnCreateConsoleScreenBuffer         uintptr
	fnCreateDirectoryTransactedW        uintptr
	fnCreateDirectoryTransactedA        uintptr
	fnCreateEventW                      uintptr
	fnCreateEventA                      uintptr
	fnCreateFiber                       uintptr
	fnCreateFileTransactedW             uintptr
	fnCreateFileTransactedA             uintptr
	fnCreateHardLinkW                   uintptr
	fnCreateHardLinkA                   uintptr
	fnCreateHardLinkTransactedW         uintptr
	fnCreateHardLinkTransactedA         uintptr
	fnCreateIoCompletionPort            uintptr
	fnCreateMailslotW                   uintptr
	fnCreateMailslotA                   uintptr
	fnCreateRemoteThread                uintptr
	fnCreateSemaphoreW                  uintptr
	fnCreateSemaphoreA                  uintptr
	fnCreateTapePartition               uintptr
	fnCreateThread                      uintptr
	fnCreateTimerQueue                  uintptr
	fnCreateTimerQueueTimer             uintptr
	fnCreateWaitableTimerW              uintptr
	fnCreateWaitableTimerA              uintptr
	fnDebugActiveProcess                uintptr
	fnDeleteAtom                        uintptr
	fnDeleteFiber                       uintptr
	fnDeleteFileTransactedW             uintptr
	fnDeleteFileTransactedA             uintptr
	fnDeleteTimerQueue                  uintptr
	fnDeleteTimerQueueEx                uintptr
	fnDeleteTimerQueueTimer             uintptr
	fnDeviceIoControl                   uintptr
	fnDisableThreadLibraryCalls         uintptr
	fnDisconnectNamedPipe               uintptr
	fnDosDateTimeToFileTime             uintptr
	fnDuplicateHandle                   uintptr
	fnEndUpdateResourceW                uintptr
	fnEndUpdateResourceA                uintptr
	fnEnumCalendarInfoW                 uintptr
	fnEnumCalendarInfoA                 uintptr
	fnEnumCalendarInfoExW               uintptr
	fnEnumCalendarInfoExA               uintptr
	fnEnumCalendarInfoExEx              uintptr
	fnEnumDateFormatsW                  uintptr
	fnEnumDateFormatsA                  uintptr
	fnEnumDateFormatsExW                uintptr
	fnEnumDateFormatsExA                uintptr
	fnEnumDateFormatsExEx               uintptr
	fnEnumLanguageGroupLocalesW         uintptr
	fnEnumLanguageGroupLocalesA         uintptr
	fnEnumResourceLanguagesW            uintptr
	fnEnumResourceLanguagesA            uintptr
	fnEnumResourceNamesW                uintptr
	fnEnumResourceNamesA                uintptr
	fnEnumResourceTypesW                uintptr
	fnEnumResourceTypesA                uintptr
	fnEnumSystemCodePagesW              uintptr
	fnEnumSystemCodePagesA              uintptr
	fnEnumSystemGeoID                   uintptr
	fnEnumSystemLanguageGroupsW         uintptr
	fnEnumSystemLanguageGroupsA         uintptr
	fnEnumSystemLocalesW                uintptr
	fnEnumSystemLocalesA                uintptr
	fnEnumSystemLocalesEx               uintptr
	fnEnumTimeFormatsW                  uintptr
	fnEnumTimeFormatsA                  uintptr
	fnEnumTimeFormatsEx                 uintptr
	fnEnumUILanguagesW                  uintptr
	fnEnumUILanguagesA                  uintptr
	fnEraseTape                         uintptr
	fnEscapeCommFunction                uintptr
	fnExpandEnvironmentStringsW         uintptr
	fnExpandEnvironmentStringsA         uintptr
	fnFileTimeToDosDateTime             uintptr
	fnFillConsoleOutputAttribute        uintptr
	fnFillConsoleOutputCharacterW       uintptr
	fnFillConsoleOutputCharacterA       uintptr
	fnFindAtomW                         uintptr
	fnFindAtomA                         uintptr
	fnFindFirstFileTransactedW          uintptr
	fnFindFirstFileTransactedA          uintptr
	fnFindNLSString                     uintptr
	fnFindNLSStringEx                   uintptr
	fnFindStringOrdinal                 uintptr
	fnFlushConsoleInputBuffer           uintptr
	fnFlushFileBuffers                  uintptr
	fnFlushInstructionCache             uintptr
	fnFoldStringW                       uintptr
	fnFoldStringA                       uintptr
	fnFormatMessageW                    uintptr
	fnFormatMessageA                    uintptr
	fnFreeConsole                       uintptr
	fnFreeEnvironmentStringsW           uintptr
	fnFreeEnvironmentStringsA           uintptr
	fnFreeLibrary                       uintptr
	fnInterlockedCompareExchange        uintptr
	fnInterlockedCompareExchange64      uintptr
	fnInterlockedDecrement              uintptr
	fnInterlockedExchange               uintptr
	fnInterlockedExchangeAdd            uintptr
	fnInterlockedIncrement              uintptr
	fnGenerateConsoleCtrlEvent          uintptr
	fnGetACP                            uintptr
	fnGetAtomNameW                      uintptr
	fnGetAtomNameA                      uintptr
	fnGetBinaryTypeW                    uintptr
	fnGetBinaryTypeA                    uintptr
	fnGetCalendarInfoW                  uintptr
	fnGetCalendarInfoA                  uintptr
	fnGetCalendarInfoEx                 uintptr
	fnGetCPInfo                         uintptr
	fnGetCPInfoExW                      uintptr
	fnGetCPInfoExA                      uintptr
	fnGetCommandLineW                   uintptr
	fnGetCommandLineA                   uintptr
	fnGetCommConfig                     uintptr
	fnGetCommMask                       uintptr
	fnGetCommModemStatus                uintptr
	fnGetCommProperties                 uintptr
	fnGetCommState                      uintptr
	fnGetCommTimeouts                   uintptr
	fnGetCompressedFileSizeW            uintptr
	fnGetCompressedFileSizeA            uintptr
	fnGetCompressedFileSizeTransactedW  uintptr
	fnGetCompressedFileSizeTransactedA  uintptr
	fnGetComputerNameW                  uintptr
	fnGetComputerNameA                  uintptr
	fnGetComputerNameExW                uintptr
	fnGetComputerNameExA                uintptr
	fnGetConsoleCP                      uintptr
	fnGetConsoleCursorInfo              uintptr
	fnGetConsoleMode                    uintptr
	fnGetConsoleOutputCP                uintptr
	fnGetConsoleScreenBufferInfo        uintptr
	fnGetConsoleTitleW                  uintptr
	fnGetConsoleTitleA                  uintptr
	fnGetCurrencyFormatW                uintptr
	fnGetCurrencyFormatA                uintptr
	fnGetCurrencyFormatEx               uintptr
	fnGetCurrentDirectoryW              uintptr
	fnGetCurrentDirectoryA              uintptr
	fnGetCurrentProcess                 uintptr
	fnGetCurrentProcessId               uintptr
	fnGetCurrentThread                  uintptr
	fnGetCurrentThreadId                uintptr
	fnGetDateFormatW                    uintptr
	fnGetDateFormatA                    uintptr
	fnGetDateFormatEx                   uintptr
	fnGetDefaultCommConfigW             uintptr
	fnGetDefaultCommConfigA             uintptr
	fnGetDurationFormat                 uintptr
	fnGetDurationFormatEx               uintptr
	fnGetDllDirectoryW                  uintptr
	fnGetDllDirectoryA                  uintptr
	fnGetEnvironmentStringsW            uintptr
	fnGetEnvironmentStringsA            uintptr
	fnGetExitCodeProcess                uintptr
	fnGetExitCodeThread                 uintptr
	fnGetFileAttributesTransactedW      uintptr
	fnGetFileAttributesTransactedA      uintptr
	fnGetFileInformationByHandle        uintptr
	fnGetFileMUIInfo                    uintptr
	fnGetFileMUIPath                    uintptr
	fnGetFullPathNameW                  uintptr
	fnGetFullPathNameA                  uintptr
	fnGetFullPathNameTransactedW        uintptr
	fnGetFullPathNameTransactedA        uintptr
	fnGetGeoInfoW                       uintptr
	fnGetGeoInfoA                       uintptr
	fnGetHandleInformation              uintptr
	fnGetLargestConsoleWindowSize       uintptr
	fnGetLastError                      uintptr
	fnGetLocaleInfoW                    uintptr
	fnGetLocaleInfoA                    uintptr
	fnGetLogicalDrives                  uintptr
	fnGetNLSVersion                     uintptr
	fnGetNLSVersionEx                   uintptr
	fnGetNumberFormatEx                 uintptr
	fnGetMailslotInfo                   uintptr
	fnGetModuleFileNameW                uintptr
	fnGetModuleFileNameA                uintptr
	fnGetModuleHandleW                  uintptr
	fnGetModuleHandleA                  uintptr
	fnGetNamedPipeHandleStateW          uintptr
	fnGetNamedPipeHandleStateA          uintptr
	fnGetNamedPipeInfo                  uintptr
	fnGetNumberFormatW                  uintptr
	fnGetNumberFormatA                  uintptr
	fnGetNumberOfConsoleInputEvents     uintptr
	fnGetNumberOfConsoleMouseButtons    uintptr
	fnGetOEMCP                          uintptr
	fnGetOverlappedResult               uintptr
	fnGetPriorityClass                  uintptr
	fnGetPrivateProfileIntW             uintptr
	fnGetPrivateProfileIntA             uintptr
	fnGetPrivateProfileSectionW         uintptr
	fnGetPrivateProfileSectionA         uintptr
	fnGetPrivateProfileSectionNamesW    uintptr
	fnGetPrivateProfileSectionNamesA    uintptr
	fnGetPrivateProfileStringW          uintptr
	fnGetPrivateProfileStringA          uintptr
	fnGetProcAddress                    uintptr
	fnGetProcessAffinityMask            uintptr
	fnGetProcessHandleCount             uintptr
	fnGetProcessHeap                    uintptr
	fnGetProcessHeaps                   uintptr
	fnGetProcessId                      uintptr
	fnGetProcessPreferredUILanguages    uintptr
	fnGetProcessPriorityBoost           uintptr
	fnGetProcessShutdownParameters      uintptr
	fnGetProcessTimes                   uintptr
	fnGetProcessVersion                 uintptr
	fnGetProcessWorkingSetSize          uintptr
	fnGetProfileIntW                    uintptr
	fnGetProfileIntA                    uintptr
	fnGetProfileSectionW                uintptr
	fnGetProfileSectionA                uintptr
	fnGetProfileStringW                 uintptr
	fnGetProfileStringA                 uintptr
	fnGetQueuedCompletionStatus         uintptr
	fnGetShortPathNameW                 uintptr
	fnGetShortPathNameA                 uintptr
	fnGetLongPathNameW                  uintptr
	fnGetLongPathNameA                  uintptr
	fnGetLongPathNameTransactedW        uintptr
	fnGetLongPathNameTransactedA        uintptr
	fnGetStdHandle                      uintptr
	fnGetStringScripts                  uintptr
	fnGetStringTypeExW                  uintptr
	fnGetStringTypeExA                  uintptr
	fnGetStringTypeA                    uintptr
	fnGetStringTypeW                    uintptr
	fnGetSystemDefaultLCID              uintptr
	fnGetSystemDefaultLangID            uintptr
	fnGetSystemDefaultLocaleName        uintptr
	fnGetSystemDefaultUILanguage        uintptr
	fnGetSystemDirectoryW               uintptr
	fnGetSystemDirectoryA               uintptr
	fnGetProductInfo                    uintptr
	fnGetSystemRegistryQuota            uintptr
	fnGetSystemTimes                    uintptr
	fnGetSystemPreferredUILanguages     uintptr
	fnGetSystemPowerStatus              uintptr
	fnGetSystemTimeAdjustment           uintptr
	fnGetTapeParameters                 uintptr
	fnGetTapePosition                   uintptr
	fnGetTapeStatus                     uintptr
	fnGetThreadContext                  uintptr
	fnGetThreadIOPendingFlag            uintptr
	fnGetThreadLocale                   uintptr
	fnGetThreadPreferredUILanguages     uintptr
	fnGetThreadPriority                 uintptr
	fnGetThreadPriorityBoost            uintptr
	fnGetThreadSelectorEntry            uintptr
	fnGetThreadTimes                    uintptr
	fnGetThreadUILanguage               uintptr
	fnGetTickCount                      uintptr
	fnGetTimeFormatW                    uintptr
	fnGetTimeFormatA                    uintptr
	fnGetTimeFormatEx                   uintptr
	fnGetUILanguageInfo                 uintptr
	fnGetUserDefaultLocaleName          uintptr
	fnGetUserDefaultUILanguage          uintptr
	fnGetUserGeoID                      uintptr
	fnGetUserPreferredUILanguages       uintptr
	fnGetDynamicTimeZoneInformation     uintptr
	fnSetDynamicTimeZoneInformation     uintptr
	fnGetUserDefaultLCID                uintptr
	fnGetUserDefaultLangID              uintptr
	fnLCIDToLocaleName                  uintptr
	fnLocaleNameToLCID                  uintptr
	fnGetLocaleInfoEx                   uintptr
	fnGetVersion                        uintptr
	fnGetVersionExW                     uintptr
	fnGetVersionExA                     uintptr
	fnGetVolumePathNameW                uintptr
	fnGetVolumePathNameA                uintptr
	fnGetVolumePathNamesForVolumeNameW  uintptr
	fnGetVolumePathNamesForVolumeNameA  uintptr
	fnGetVolumeNameForVolumeMountPointW uintptr
	fnGetVolumeNameForVolumeMountPointA uintptr
	fnFindFirstVolumeW                  uintptr
	fnFindFirstVolumeA                  uintptr
	fnFindNextVolumeW                   uintptr
	fnFindNextVolumeA                   uintptr
	fnFindVolumeClose                   uintptr
	fnFindFirstVolumeMountPointW        uintptr
	fnFindFirstVolumeMountPointA        uintptr
	fnFindNextVolumeMountPointW         uintptr
	fnFindNextVolumeMountPointA         uintptr
	fnFindVolumeMountPointClose         uintptr
	fnCreateSymbolicLinkW               uintptr
	fnCreateSymbolicLinkA               uintptr
	fnCreateSymbolicLinkTransactedW     uintptr
	fnCreateSymbolicLinkTransactedA     uintptr
	fnGetFinalPathNameByHandleW         uintptr
	fnGetFinalPathNameByHandleA         uintptr
	fnGetWindowsDirectoryW              uintptr
	fnGetWindowsDirectoryA              uintptr
	fnGlobalAddAtomW                    uintptr
	fnGlobalAddAtomA                    uintptr
	fnGlobalAlloc                       uintptr
	fnGlobalCompact                     uintptr
	fnGlobalDeleteAtom                  uintptr
	fnGlobalFindAtomW                   uintptr
	fnGlobalFindAtomA                   uintptr
	fnGlobalFlags                       uintptr
	fnGlobalFree                        uintptr
	fnGlobalGetAtomNameW                uintptr
	fnGlobalGetAtomNameA                uintptr
	fnGlobalLock                        uintptr
	fnGlobalHandle                      uintptr
	fnGlobalMemoryStatusEx              uintptr
	fnGlobalReAlloc                     uintptr
	fnGlobalSize                        uintptr
	fnGlobalUnWire                      uintptr
	fnGlobalUnlock                      uintptr
	fnGlobalWire                        uintptr
	fnHeapAlloc                         uintptr
	fnHeapCompact                       uintptr
	fnHeapCreate                        uintptr
	fnHeapDestroy                       uintptr
	fnHeapFree                          uintptr
	fnHeapLock                          uintptr
	fnHeapReAlloc                       uintptr
	fnHeapSize                          uintptr
	fnHeapUnlock                        uintptr
	fnHeapValidate                      uintptr
	fnHeapWalk                          uintptr
	fnIdnToAscii                        uintptr
	fnIdnToNameprepUnicode              uintptr
	fnIdnToUnicode                      uintptr
	fnInitAtomTable                     uintptr
	fnIsBadCodePtr                      uintptr
	fnIsBadHugeReadPtr                  uintptr
	fnIsBadHugeWritePtr                 uintptr
	fnIsBadReadPtr                      uintptr
	fnIsBadStringPtrW                   uintptr
	fnIsBadStringPtrA                   uintptr
	fnIsBadWritePtr                     uintptr
	fnIsDBCSLeadByte                    uintptr
	fnIsDBCSLeadByteEx                  uintptr
	fnIsNLSDefinedString                uintptr
	fnIsNormalizedString                uintptr
	fnIsProcessorFeaturePresent         uintptr
	fnIsValidCodePage                   uintptr
	fnIsValidLanguageGroup              uintptr
	fnIsValidLocale                     uintptr
	fnIsValidLocaleName                 uintptr
	fnLCMapStringW                      uintptr
	fnLCMapStringA                      uintptr
	fnLCMapStringEx                     uintptr
	fnLoadModule                        uintptr
	fnLocalAlloc                        uintptr
	fnLocalCompact                      uintptr
	fnLocalFlags                        uintptr
	fnLocalFree                         uintptr
	fnLocalLock                         uintptr
	fnLocalReAlloc                      uintptr
	fnLocalShrink                       uintptr
	fnLocalSize                         uintptr
	fnLocalUnlock                       uintptr
	fnLockFile                          uintptr
	fnLockFileEx                        uintptr
	fnMoveFileW                         uintptr
	fnMoveFileA                         uintptr
	fnMoveFileExW                       uintptr
	fnMoveFileExA                       uintptr
	fnMoveFileTransactedW               uintptr
	fnMoveFileTransactedA               uintptr
	fnMoveFileWithProgressW             uintptr
	fnMoveFileWithProgressA             uintptr
	fnMulDiv                            uintptr
	fnMultiByteToWideChar               uintptr
	fnNormalizeString                   uintptr
	fnOpenEventW                        uintptr
	fnOpenEventA                        uintptr
	fnOpenFile                          uintptr
	fnOpenFileMappingW                  uintptr
	fnOpenFileMappingA                  uintptr
	fnOpenMutexW                        uintptr
	fnOpenMutexA                        uintptr
	fnOpenProcess                       uintptr
	fnOpenRawW                          uintptr
	fnOpenRawA                          uintptr
	fnOpenSemaphoreW                    uintptr
	fnOpenWaitableTimerW                uintptr
	fnOpenSemaphoreA                    uintptr
	fnOpenWaitableTimerA                uintptr
	fnIsDebuggerPresent                 uintptr
	fnPeekConsoleInputW                 uintptr
	fnPeekConsoleInputA                 uintptr
	fnPeekNamedPipe                     uintptr
	fnPostQueuedCompletionStatus        uintptr
	fnPrepareTape                       uintptr
	fnPulseEvent                        uintptr
	fnPurgeComm                         uintptr
	fnQueryPerformanceCounter           uintptr
	fnQueryPerformanceFrequency         uintptr
	fnQueryRecoveryAgentsW              uintptr
	fnQueryRecoveryAgentsA              uintptr
	fnQueueUserAPC                      uintptr
	fnQueueUserWorkItem                 uintptr
	fnReadConsoleW                      uintptr
	fnReadConsoleA                      uintptr
	fnReadConsoleInputW                 uintptr
	fnReadConsoleInputA                 uintptr
	fnReadConsoleOutputW                uintptr
	fnReadConsoleOutputA                uintptr
	fnReadConsoleOutputAttribute        uintptr
	fnReadConsoleOutputCharacterW       uintptr
	fnReadConsoleOutputCharacterA       uintptr
	fnReadDirectoryChangesW             uintptr
	fnReadFile                          uintptr
	fnReadFileEx                        uintptr
	fnReadProcessMemory                 uintptr
	fnReadRaw                           uintptr
	fnRegisterWaitForSingleObject       uintptr
	fnRegisterWaitForSingleObjectEx     uintptr
	fnReleaseMutex                      uintptr
	fnReleaseSemaphore                  uintptr
	fnRemoveDirectoryW                  uintptr
	fnRemoveDirectoryA                  uintptr
	fnRemoveDirectoryTransactedW        uintptr
	fnRemoveDirectoryTransactedA        uintptr
	fnResetEvent                        uintptr
	fnResolveLocaleName                 uintptr
	fnResumeThread                      uintptr
	fnSearchPathW                       uintptr
	fnSearchPathA                       uintptr
	fnSetCalendarInfoW                  uintptr
	fnSetCalendarInfoA                  uintptr
	fnSetCommBreak                      uintptr
	fnSetCommConfig                     uintptr
	fnSetCommMask                       uintptr
	fnSetCommState                      uintptr
	fnSetCommTimeouts                   uintptr
	fnSetComputerNameW                  uintptr
	fnSetComputerNameA                  uintptr
	fnSetConsoleActiveScreenBuffer      uintptr
	fnSetConsoleCP                      uintptr
	fnSetConsoleCtrlHandler             uintptr
	fnSetConsoleCursorPosition          uintptr
	fnSetConsoleMode                    uintptr
	fnSetConsoleOutputCP                uintptr
	fnSetConsoleScreenBufferSize        uintptr
	fnSetConsoleTextAttribute           uintptr
	fnSetConsoleTitleW                  uintptr
	fnSetConsoleTitleA                  uintptr
	fnSetCriticalSectionSpinCount       uintptr
	fnSetCurrentDirectoryW              uintptr
	fnSetCurrentDirectoryA              uintptr
	fnSetDefaultCommConfigW             uintptr
	fnSetDefaultCommConfigA             uintptr
	fnSetDllDirectoryW                  uintptr
	fnSetDllDirectoryA                  uintptr
	fnSetEndOfFile                      uintptr
	fnSetEnvironmentVariableW           uintptr
	fnSetEnvironmentVariableA           uintptr
	fnGetThreadErrorMode                uintptr
	fnSetThreadErrorMode                uintptr
	fnSetEvent                          uintptr
	fnSetFileAttributesW                uintptr
	fnSetFileAttributesA                uintptr
	fnSetFileAttributesTransactedW      uintptr
	fnSetFileAttributesTransactedA      uintptr
	fnSetFilePointer                    uintptr
	fnSetFilePointerEx                  uintptr
	fnSetFileTime                       uintptr
	fnSetHandleCount                    uintptr
	fnSetHandleInformation              uintptr
	fnSetLocalTime                      uintptr
	fnSetLocaleInfoW                    uintptr
	fnSetLocaleInfoA                    uintptr
	fnSetMailslotInfo                   uintptr
	fnSetNamedPipeHandleState           uintptr
	fnSetPriorityClass                  uintptr
	fnSetProcessAffinityMask            uintptr
	fnSetProcessPreferredUILanguages    uintptr
	fnSetProcessPriorityBoost           uintptr
	fnSetProcessShutdownParameters      uintptr
	fnSetProcessWorkingSetSize          uintptr
	fnSetStdHandle                      uintptr
	fnSetSystemPowerState               uintptr
	fnSetSystemTime                     uintptr
	fnSetSystemTimeAdjustment           uintptr
	fnSetTapeParameters                 uintptr
	fnSetTapePosition                   uintptr
	fnSetThreadAffinityMask             uintptr
	fnSetThreadContext                  uintptr
	fnSetThreadExecutionState           uintptr
	fnSetThreadIdealProcessor           uintptr
	fnSetThreadLocale                   uintptr
	fnSetThreadPreferredUILanguages     uintptr
	fnSetThreadPriority                 uintptr
	fnSetThreadPriorityBoost            uintptr
	fnSetThreadUILanguage               uintptr
	fnSetTimerQueueTimer                uintptr
	fnSetTimeZoneInformation            uintptr
	fnSetUserGeoID                      uintptr
	fnSetUnhandledExceptionFilter       uintptr
	fnSetVolumeLabelW                   uintptr
	fnSetVolumeLabelA                   uintptr
	fnSetWaitableTimer                  uintptr
	fnSetupComm                         uintptr
	fnSignalObjectAndWait               uintptr
	fnSleepEx                           uintptr
	fnSuspendThread                     uintptr
	fnSwitchToThread                    uintptr
	fnSystemTimeToFileTime              uintptr
	fnSystemTimeToTzSpecificLocalTime   uintptr
	fnTerminateProcess                  uintptr
	fnTerminateThread                   uintptr
	fnTlsAlloc                          uintptr
	fnTlsFree                           uintptr
	fnTlsGetValue                       uintptr
	fnTlsSetValue                       uintptr
	fnTransactNamedPipe                 uintptr
	fnTransmitCommChar                  uintptr
	fnUnhandledExceptionFilter          uintptr
	fnUnlockFile                        uintptr
	fnUnlockFileEx                      uintptr
	fnUnregisterWait                    uintptr
	fnUnregisterWaitEx                  uintptr
	fnUpdateResourceW                   uintptr
	fnUpdateResourceA                   uintptr
	fnVerLanguageNameW                  uintptr
	fnVerLanguageNameA                  uintptr
	fnVerifyScripts                     uintptr
	fnVerifyVersionInfoW                uintptr
	fnVerifyVersionInfoA                uintptr
	fnVerSetConditionMask               uintptr
	fnVirtualAlloc                      uintptr
	fnVirtualAllocEx                    uintptr
	fnVirtualFree                       uintptr
	fnVirtualFreeEx                     uintptr
	fnVirtualLock                       uintptr
	fnVirtualQuery                      uintptr
	fnVirtualQueryEx                    uintptr
	fnVirtualUnlock                     uintptr
	fnWaitCommEvent                     uintptr
	fnWaitForDebugEvent                 uintptr
	fnWaitForMultipleObjects            uintptr
	fnWaitForMultipleObjectsEx          uintptr
	fnWaitForSingleObject               uintptr
	fnWaitForSingleObjectEx             uintptr
	fnSleepConditionVariableCS          uintptr
	fnWaitNamedPipeW                    uintptr
	fnWaitNamedPipeA                    uintptr
	fnWideCharToMultiByte               uintptr
	fnWinExec                           uintptr
	fnWriteConsoleW                     uintptr
	fnWriteConsoleA                     uintptr
	fnWriteConsoleInputW                uintptr
	fnWriteConsoleInputA                uintptr
	fnWriteConsoleOutputW               uintptr
	fnWriteConsoleOutputA               uintptr
	fnWriteConsoleOutputAttribute       uintptr
	fnWriteConsoleOutputCharacterW      uintptr
	fnWriteConsoleOutputCharacterA      uintptr
	fnWriteFile                         uintptr
	fnWriteFileEx                       uintptr
	fnWritePrivateProfileSectionW       uintptr
	fnWritePrivateProfileSectionA       uintptr
	fnWritePrivateProfileStringW        uintptr
	fnWritePrivateProfileStringA        uintptr
	fnWriteProcessMemory                uintptr
	fnWriteProfileSectionW              uintptr
	fnWriteProfileSectionA              uintptr
	fnWriteProfileStringW               uintptr
	fnWriteProfileStringA               uintptr
	fnWriteRaw                          uintptr
	fnWriteTapemark                     uintptr
	fn_hread                            uintptr
	fn_hwrite                           uintptr
	fn_lclose                           uintptr
	fn_lcreat                           uintptr
	fn_llseek                           uintptr
	fn_lopen                            uintptr
	fn_lread                            uintptr
	fn_lwrite                           uintptr
	fnlstrcatW                          uintptr
	fnlstrcatA                          uintptr
	fnlstrcmpW                          uintptr
	fnlstrcmpA                          uintptr
	fnlstrcmpiW                         uintptr
	fnlstrcmpiA                         uintptr
	fnlstrcpyW                          uintptr
	fnlstrcpyA                          uintptr
	fnlstrcpynW                         uintptr
	fnlstrcpynA                         uintptr
	fnlstrlenW                          uintptr
	fnlstrlenA                          uintptr
	fnGetPrivateProfileStructW          uintptr
	fnGetPrivateProfileStructA          uintptr
	fnWritePrivateProfileStructW        uintptr
	fnWritePrivateProfileStructA        uintptr
	fnCreateActCtxW                     uintptr
	fnCreateActCtxA                     uintptr
	fnZombifyActCtx                     uintptr
	fnActivateActCtx                    uintptr
	fnDeactivateActCtx                  uintptr
	fnGetCurrentActCtx                  uintptr
	fnProcessIdToSessionId              uintptr
	fnWTSGetActiveConsoleSessionId      uintptr
	fnIsWow64Process                    uintptr
	fnGetLogicalProcessorInformation    uintptr
	fnGetLogicalProcessorInformationEx  uintptr
	fnGetNumaHighestNodeNumber          uintptr
	fnGetNumaProcessorNode              uintptr
	fnGetNumaAvailableMemoryNode        uintptr
	fnGetNumaNodeProcessorMask          uintptr
	fnSetProcessDEPPolicy               uintptr
	fnGetNamedPipeServerSessionId       uintptr
	fnGetNamedPipeServerProcessId       uintptr
	fnGetNamedPipeClientProcessId       uintptr
	fnGetNamedPipeClientSessionId       uintptr
	fnGetNamedPipeClientComputerNameW   uintptr
	fnGetNamedPipeClientComputerNameA   uintptr
	fnMoveMemory						uintptr
)

func init() {
	initKernel32()
}

func initKernel32() {
	if libkernel32 != 0 {
		return
	}
	libkernel32, _ = syscall.LoadLibrary("kernel32.dll")
	fnGetModuleHandleW, _ = syscall.GetProcAddress(libkernel32, "GetModuleHandleW")
	fnGetModuleHandleA, _ = syscall.GetProcAddress(libkernel32, "GetModuleHandleA")
	fnAddAtomW, _ = syscall.GetProcAddress(libkernel32, "AddAtomW")
	fnAddAtomA, _ = syscall.GetProcAddress(libkernel32, "AddAtomA")
	fnAllocConsole, _ = syscall.GetProcAddress(libkernel32, "AllocConsole")
	fnAreFileApisANSI, _ = syscall.GetProcAddress(libkernel32, "AreFileApisANSI")
	fnBackupRead, _ = syscall.GetProcAddress(libkernel32, "BackupRead")
	fnBackupSeek, _ = syscall.GetProcAddress(libkernel32, "BackupSeek")
	fnBackupWrite, _ = syscall.GetProcAddress(libkernel32, "BackupWrite")
	fnBeep, _ = syscall.GetProcAddress(libkernel32, "Beep")
	fnBeginUpdateResourceW, _ = syscall.GetProcAddress(libkernel32, "BeginUpdateResourceW")
	fnBeginUpdateResourceA, _ = syscall.GetProcAddress(libkernel32, "BeginUpdateResourceA")
	fnBindIoCompletionCallback, _ = syscall.GetProcAddress(libkernel32, "BindIoCompletionCallback")
	fnBuildCommDCBW, _ = syscall.GetProcAddress(libkernel32, "BuildCommDCBW")
	fnBuildCommDCBA, _ = syscall.GetProcAddress(libkernel32, "BuildCommDCBA")
	fnBuildCommDCBAndTimeoutsW, _ = syscall.GetProcAddress(libkernel32, "BuildCommDCBAndTimeoutsW")
	fnBuildCommDCBAndTimeoutsA, _ = syscall.GetProcAddress(libkernel32, "BuildCommDCBAndTimeoutsA")
	fnCallNamedPipeW, _ = syscall.GetProcAddress(libkernel32, "CallNamedPipeW")
	fnCallNamedPipeA, _ = syscall.GetProcAddress(libkernel32, "CallNamedPipeA")
	fnCancelIo, _ = syscall.GetProcAddress(libkernel32, "CancelIo")
	fnCancelTimerQueueTimer, _ = syscall.GetProcAddress(libkernel32, "CancelTimerQueueTimer")
	fnCancelWaitableTimer, _ = syscall.GetProcAddress(libkernel32, "CancelWaitableTimer")
	fnChangeTimerQueueTimer, _ = syscall.GetProcAddress(libkernel32, "ChangeTimerQueueTimer")
	fnCheckNameLegalDOS8Dot3W, _ = syscall.GetProcAddress(libkernel32, "CheckNameLegalDOS8Dot3W")
	fnCheckNameLegalDOS8Dot3A, _ = syscall.GetProcAddress(libkernel32, "CheckNameLegalDOS8Dot3A")
	fnCheckRemoteDebuggerPresent, _ = syscall.GetProcAddress(libkernel32, "CheckRemoteDebuggerPresent")
	fnClearCommBreak, _ = syscall.GetProcAddress(libkernel32, "ClearCommBreak")
	fnClearCommError, _ = syscall.GetProcAddress(libkernel32, "ClearCommError")
	fnCommConfigDialogW, _ = syscall.GetProcAddress(libkernel32, "CommConfigDialogW")
	fnCommConfigDialogA, _ = syscall.GetProcAddress(libkernel32, "CommConfigDialogA")
	fnCompareFileTime, _ = syscall.GetProcAddress(libkernel32, "CompareFileTime")
	fnConnectNamedPipe, _ = syscall.GetProcAddress(libkernel32, "ConnectNamedPipe")
	fnContinueDebugEvent, _ = syscall.GetProcAddress(libkernel32, "ContinueDebugEvent")
	fnConvertThreadToFiber, _ = syscall.GetProcAddress(libkernel32, "ConvertThreadToFiber")
	fnConvertFiberToThread, _ = syscall.GetProcAddress(libkernel32, "ConvertFiberToThread")
	fnConvertDefaultLocale, _ = syscall.GetProcAddress(libkernel32, "ConvertDefaultLocale")
	fnCopyFileTransactedW, _ = syscall.GetProcAddress(libkernel32, "CopyFileTransactedW")
	fnCopyFileTransactedA, _ = syscall.GetProcAddress(libkernel32, "CopyFileTransactedA")
	fnReplaceFileW, _ = syscall.GetProcAddress(libkernel32, "ReplaceFileW")
	fnReplaceFileA, _ = syscall.GetProcAddress(libkernel32, "ReplaceFileA")
	fnCreateConsoleScreenBuffer, _ = syscall.GetProcAddress(libkernel32, "CreateConsoleScreenBuffer")
	fnCreateDirectoryTransactedW, _ = syscall.GetProcAddress(libkernel32, "CreateDirectoryTransactedW")
	fnCreateDirectoryTransactedA, _ = syscall.GetProcAddress(libkernel32, "CreateDirectoryTransactedA")
	fnCreateEventW, _ = syscall.GetProcAddress(libkernel32, "CreateEventW")
	fnCreateEventA, _ = syscall.GetProcAddress(libkernel32, "CreateEventA")
	fnCreateFiber, _ = syscall.GetProcAddress(libkernel32, "CreateFiber")
	fnCreateFileTransactedW, _ = syscall.GetProcAddress(libkernel32, "CreateFileTransactedW")
	fnCreateFileTransactedA, _ = syscall.GetProcAddress(libkernel32, "CreateFileTransactedA")
	fnCreateHardLinkW, _ = syscall.GetProcAddress(libkernel32, "CreateHardLinkW")
	fnCreateHardLinkA, _ = syscall.GetProcAddress(libkernel32, "CreateHardLinkA")
	fnCreateHardLinkTransactedW, _ = syscall.GetProcAddress(libkernel32, "CreateHardLinkTransactedW")
	fnCreateHardLinkTransactedA, _ = syscall.GetProcAddress(libkernel32, "CreateHardLinkTransactedA")
	fnCreateIoCompletionPort, _ = syscall.GetProcAddress(libkernel32, "CreateIoCompletionPort")
	fnCreateMailslotW, _ = syscall.GetProcAddress(libkernel32, "CreateMailslotW")
	fnCreateMailslotA, _ = syscall.GetProcAddress(libkernel32, "CreateMailslotA")
	fnCreateRemoteThread, _ = syscall.GetProcAddress(libkernel32, "CreateRemoteThread")
	fnCreateSemaphoreW, _ = syscall.GetProcAddress(libkernel32, "CreateSemaphoreW")
	fnCreateSemaphoreA, _ = syscall.GetProcAddress(libkernel32, "CreateSemaphoreA")
	fnCreateTapePartition, _ = syscall.GetProcAddress(libkernel32, "CreateTapePartition")
	fnCreateThread, _ = syscall.GetProcAddress(libkernel32, "CreateThread")
	fnCreateTimerQueue, _ = syscall.GetProcAddress(libkernel32, "CreateTimerQueue")
	fnCreateTimerQueueTimer, _ = syscall.GetProcAddress(libkernel32, "CreateTimerQueueTimer")
	fnCreateWaitableTimerW, _ = syscall.GetProcAddress(libkernel32, "CreateWaitableTimerW")
	fnCreateWaitableTimerA, _ = syscall.GetProcAddress(libkernel32, "CreateWaitableTimerA")
	fnDebugActiveProcess, _ = syscall.GetProcAddress(libkernel32, "DebugActiveProcess")
	fnDeleteAtom, _ = syscall.GetProcAddress(libkernel32, "DeleteAtom")
	fnDeleteFiber, _ = syscall.GetProcAddress(libkernel32, "DeleteFiber")
	fnDeleteFileTransactedW, _ = syscall.GetProcAddress(libkernel32, "DeleteFileTransactedW")
	fnDeleteFileTransactedA, _ = syscall.GetProcAddress(libkernel32, "DeleteFileTransactedA")
	fnDeleteTimerQueue, _ = syscall.GetProcAddress(libkernel32, "DeleteTimerQueue")
	fnDeleteTimerQueueEx, _ = syscall.GetProcAddress(libkernel32, "DeleteTimerQueueEx")
	fnDeleteTimerQueueTimer, _ = syscall.GetProcAddress(libkernel32, "DeleteTimerQueueTimer")
	fnDeviceIoControl, _ = syscall.GetProcAddress(libkernel32, "DeviceIoControl")
	fnDisableThreadLibraryCalls, _ = syscall.GetProcAddress(libkernel32, "DisableThreadLibraryCalls")
	fnDisconnectNamedPipe, _ = syscall.GetProcAddress(libkernel32, "DisconnectNamedPipe")
	fnDosDateTimeToFileTime, _ = syscall.GetProcAddress(libkernel32, "DosDateTimeToFileTime")
	fnDuplicateHandle, _ = syscall.GetProcAddress(libkernel32, "DuplicateHandle")
	fnEndUpdateResourceW, _ = syscall.GetProcAddress(libkernel32, "EndUpdateResourceW")
	fnEndUpdateResourceA, _ = syscall.GetProcAddress(libkernel32, "EndUpdateResourceA")
	fnEnumCalendarInfoW, _ = syscall.GetProcAddress(libkernel32, "EnumCalendarInfoW")
	fnEnumCalendarInfoA, _ = syscall.GetProcAddress(libkernel32, "EnumCalendarInfoA")
	fnEnumCalendarInfoExW, _ = syscall.GetProcAddress(libkernel32, "EnumCalendarInfoExW")
	fnEnumCalendarInfoExA, _ = syscall.GetProcAddress(libkernel32, "EnumCalendarInfoExA")
	fnEnumCalendarInfoExEx, _ = syscall.GetProcAddress(libkernel32, "EnumCalendarInfoExEx")
	fnEnumDateFormatsW, _ = syscall.GetProcAddress(libkernel32, "EnumDateFormatsW")
	fnEnumDateFormatsA, _ = syscall.GetProcAddress(libkernel32, "EnumDateFormatsA")
	fnEnumDateFormatsExW, _ = syscall.GetProcAddress(libkernel32, "EnumDateFormatsExW")
	fnEnumDateFormatsExA, _ = syscall.GetProcAddress(libkernel32, "EnumDateFormatsExA")
	fnEnumDateFormatsExEx, _ = syscall.GetProcAddress(libkernel32, "EnumDateFormatsExEx")
	fnEnumLanguageGroupLocalesW, _ = syscall.GetProcAddress(libkernel32, "EnumLanguageGroupLocalesW")
	fnEnumLanguageGroupLocalesA, _ = syscall.GetProcAddress(libkernel32, "EnumLanguageGroupLocalesA")
	fnEnumResourceLanguagesW, _ = syscall.GetProcAddress(libkernel32, "EnumResourceLanguagesW")
	fnEnumResourceLanguagesA, _ = syscall.GetProcAddress(libkernel32, "EnumResourceLanguagesA")
	fnEnumResourceNamesW, _ = syscall.GetProcAddress(libkernel32, "EnumResourceNamesW")
	fnEnumResourceNamesA, _ = syscall.GetProcAddress(libkernel32, "EnumResourceNamesA")
	fnEnumResourceTypesW, _ = syscall.GetProcAddress(libkernel32, "EnumResourceTypesW")
	fnEnumResourceTypesA, _ = syscall.GetProcAddress(libkernel32, "EnumResourceTypesA")
	fnEnumSystemCodePagesW, _ = syscall.GetProcAddress(libkernel32, "EnumSystemCodePagesW")
	fnEnumSystemCodePagesA, _ = syscall.GetProcAddress(libkernel32, "EnumSystemCodePagesA")
	fnEnumSystemGeoID, _ = syscall.GetProcAddress(libkernel32, "EnumSystemGeoID")
	fnEnumSystemLanguageGroupsW, _ = syscall.GetProcAddress(libkernel32, "EnumSystemLanguageGroupsW")
	fnEnumSystemLanguageGroupsA, _ = syscall.GetProcAddress(libkernel32, "EnumSystemLanguageGroupsA")
	fnEnumSystemLocalesW, _ = syscall.GetProcAddress(libkernel32, "EnumSystemLocalesW")
	fnEnumSystemLocalesA, _ = syscall.GetProcAddress(libkernel32, "EnumSystemLocalesA")
	fnEnumSystemLocalesEx, _ = syscall.GetProcAddress(libkernel32, "EnumSystemLocalesEx")
	fnEnumTimeFormatsW, _ = syscall.GetProcAddress(libkernel32, "EnumTimeFormatsW")
	fnEnumTimeFormatsA, _ = syscall.GetProcAddress(libkernel32, "EnumTimeFormatsA")
	fnEnumTimeFormatsEx, _ = syscall.GetProcAddress(libkernel32, "EnumTimeFormatsEx")
	fnEnumUILanguagesW, _ = syscall.GetProcAddress(libkernel32, "EnumUILanguagesW")
	fnEnumUILanguagesA, _ = syscall.GetProcAddress(libkernel32, "EnumUILanguagesA")
	fnEraseTape, _ = syscall.GetProcAddress(libkernel32, "EraseTape")
	fnEscapeCommFunction, _ = syscall.GetProcAddress(libkernel32, "EscapeCommFunction")
	fnExpandEnvironmentStringsW, _ = syscall.GetProcAddress(libkernel32, "ExpandEnvironmentStringsW")
	fnExpandEnvironmentStringsA, _ = syscall.GetProcAddress(libkernel32, "ExpandEnvironmentStringsA")
	fnFileTimeToDosDateTime, _ = syscall.GetProcAddress(libkernel32, "FileTimeToDosDateTime")
	fnFillConsoleOutputAttribute, _ = syscall.GetProcAddress(libkernel32, "FillConsoleOutputAttribute")
	fnFillConsoleOutputCharacterW, _ = syscall.GetProcAddress(libkernel32, "FillConsoleOutputCharacterW")
	fnFillConsoleOutputCharacterA, _ = syscall.GetProcAddress(libkernel32, "FillConsoleOutputCharacterA")
	fnFindAtomW, _ = syscall.GetProcAddress(libkernel32, "FindAtomW")
	fnFindAtomA, _ = syscall.GetProcAddress(libkernel32, "FindAtomA")
	fnFindFirstFileTransactedW, _ = syscall.GetProcAddress(libkernel32, "FindFirstFileTransactedW")
	fnFindFirstFileTransactedA, _ = syscall.GetProcAddress(libkernel32, "FindFirstFileTransactedA")
	fnFindNLSString, _ = syscall.GetProcAddress(libkernel32, "FindNLSString")
	fnFindNLSStringEx, _ = syscall.GetProcAddress(libkernel32, "FindNLSStringEx")
	fnFindStringOrdinal, _ = syscall.GetProcAddress(libkernel32, "FindStringOrdinal")
	fnFlushConsoleInputBuffer, _ = syscall.GetProcAddress(libkernel32, "FlushConsoleInputBuffer")
	fnFlushFileBuffers, _ = syscall.GetProcAddress(libkernel32, "FlushFileBuffers")
	fnFlushInstructionCache, _ = syscall.GetProcAddress(libkernel32, "FlushInstructionCache")
	fnFoldStringW, _ = syscall.GetProcAddress(libkernel32, "FoldStringW")
	fnFoldStringA, _ = syscall.GetProcAddress(libkernel32, "FoldStringA")
	fnFormatMessageW, _ = syscall.GetProcAddress(libkernel32, "FormatMessageW")
	fnFormatMessageA, _ = syscall.GetProcAddress(libkernel32, "FormatMessageA")
	fnFreeConsole, _ = syscall.GetProcAddress(libkernel32, "FreeConsole")
	fnFreeEnvironmentStringsW, _ = syscall.GetProcAddress(libkernel32, "FreeEnvironmentStringsW")
	fnFreeEnvironmentStringsA, _ = syscall.GetProcAddress(libkernel32, "FreeEnvironmentStringsA")
	fnFreeLibrary, _ = syscall.GetProcAddress(libkernel32, "FreeLibrary")
	fnInterlockedCompareExchange, _ = syscall.GetProcAddress(libkernel32, "InterlockedCompareExchange")
	fnInterlockedCompareExchange64, _ = syscall.GetProcAddress(libkernel32, "InterlockedCompareExchange64")
	fnInterlockedDecrement, _ = syscall.GetProcAddress(libkernel32, "InterlockedDecrement")
	fnInterlockedExchange, _ = syscall.GetProcAddress(libkernel32, "InterlockedExchange")
	fnInterlockedExchangeAdd, _ = syscall.GetProcAddress(libkernel32, "InterlockedExchangeAdd")
	fnInterlockedIncrement, _ = syscall.GetProcAddress(libkernel32, "InterlockedIncrement")
	fnGenerateConsoleCtrlEvent, _ = syscall.GetProcAddress(libkernel32, "GenerateConsoleCtrlEvent")
	fnGetACP, _ = syscall.GetProcAddress(libkernel32, "GetACP")
	fnGetAtomNameW, _ = syscall.GetProcAddress(libkernel32, "GetAtomNameW")
	fnGetAtomNameA, _ = syscall.GetProcAddress(libkernel32, "GetAtomNameA")
	fnGetBinaryTypeW, _ = syscall.GetProcAddress(libkernel32, "GetBinaryTypeW")
	fnGetBinaryTypeA, _ = syscall.GetProcAddress(libkernel32, "GetBinaryTypeA")
	fnGetCalendarInfoW, _ = syscall.GetProcAddress(libkernel32, "GetCalendarInfoW")
	fnGetCalendarInfoA, _ = syscall.GetProcAddress(libkernel32, "GetCalendarInfoA")
	fnGetCalendarInfoEx, _ = syscall.GetProcAddress(libkernel32, "GetCalendarInfoEx")
	fnGetCPInfo, _ = syscall.GetProcAddress(libkernel32, "GetCPInfo")
	fnGetCPInfoExW, _ = syscall.GetProcAddress(libkernel32, "GetCPInfoExW")
	fnGetCPInfoExA, _ = syscall.GetProcAddress(libkernel32, "GetCPInfoExA")
	fnGetCommandLineW, _ = syscall.GetProcAddress(libkernel32, "GetCommandLineW")
	fnGetCommandLineA, _ = syscall.GetProcAddress(libkernel32, "GetCommandLineA")
	fnGetCommConfig, _ = syscall.GetProcAddress(libkernel32, "GetCommConfig")
	fnGetCommMask, _ = syscall.GetProcAddress(libkernel32, "GetCommMask")
	fnGetCommModemStatus, _ = syscall.GetProcAddress(libkernel32, "GetCommModemStatus")
	fnGetCommProperties, _ = syscall.GetProcAddress(libkernel32, "GetCommProperties")
	fnGetCommState, _ = syscall.GetProcAddress(libkernel32, "GetCommState")
	fnGetCommTimeouts, _ = syscall.GetProcAddress(libkernel32, "GetCommTimeouts")
	fnGetCompressedFileSizeW, _ = syscall.GetProcAddress(libkernel32, "GetCompressedFileSizeW")
	fnGetCompressedFileSizeA, _ = syscall.GetProcAddress(libkernel32, "GetCompressedFileSizeA")
	fnGetCompressedFileSizeTransactedW, _ = syscall.GetProcAddress(libkernel32, "GetCompressedFileSizeTransactedW")
	fnGetCompressedFileSizeTransactedA, _ = syscall.GetProcAddress(libkernel32, "GetCompressedFileSizeTransactedA")
	fnGetComputerNameW, _ = syscall.GetProcAddress(libkernel32, "GetComputerNameW")
	fnGetComputerNameA, _ = syscall.GetProcAddress(libkernel32, "GetComputerNameA")
	fnGetComputerNameExW, _ = syscall.GetProcAddress(libkernel32, "GetComputerNameExW")
	fnGetComputerNameExA, _ = syscall.GetProcAddress(libkernel32, "GetComputerNameExA")
	fnGetConsoleCP, _ = syscall.GetProcAddress(libkernel32, "GetConsoleCP")
	fnGetConsoleCursorInfo, _ = syscall.GetProcAddress(libkernel32, "GetConsoleCursorInfo")
	fnGetConsoleMode, _ = syscall.GetProcAddress(libkernel32, "GetConsoleMode")
	fnGetConsoleOutputCP, _ = syscall.GetProcAddress(libkernel32, "GetConsoleOutputCP")
	fnGetConsoleScreenBufferInfo, _ = syscall.GetProcAddress(libkernel32, "GetConsoleScreenBufferInfo")
	fnGetConsoleTitleW, _ = syscall.GetProcAddress(libkernel32, "GetConsoleTitleW")
	fnGetConsoleTitleA, _ = syscall.GetProcAddress(libkernel32, "GetConsoleTitleA")
	fnGetCurrencyFormatW, _ = syscall.GetProcAddress(libkernel32, "GetCurrencyFormatW")
	fnGetCurrencyFormatA, _ = syscall.GetProcAddress(libkernel32, "GetCurrencyFormatA")
	fnGetCurrencyFormatEx, _ = syscall.GetProcAddress(libkernel32, "GetCurrencyFormatEx")
	fnGetCurrentDirectoryW, _ = syscall.GetProcAddress(libkernel32, "GetCurrentDirectoryW")
	fnGetCurrentDirectoryA, _ = syscall.GetProcAddress(libkernel32, "GetCurrentDirectoryA")
	fnGetCurrentProcess, _ = syscall.GetProcAddress(libkernel32, "GetCurrentProcess")
	fnGetCurrentProcessId, _ = syscall.GetProcAddress(libkernel32, "GetCurrentProcessId")
	fnGetCurrentThread, _ = syscall.GetProcAddress(libkernel32, "GetCurrentThread")
	fnGetCurrentThreadId, _ = syscall.GetProcAddress(libkernel32, "GetCurrentThreadId")
	fnGetDateFormatW, _ = syscall.GetProcAddress(libkernel32, "GetDateFormatW")
	fnGetDateFormatA, _ = syscall.GetProcAddress(libkernel32, "GetDateFormatA")
	fnGetDateFormatEx, _ = syscall.GetProcAddress(libkernel32, "GetDateFormatEx")
	fnGetDefaultCommConfigW, _ = syscall.GetProcAddress(libkernel32, "GetDefaultCommConfigW")
	fnGetDefaultCommConfigA, _ = syscall.GetProcAddress(libkernel32, "GetDefaultCommConfigA")
	fnGetDurationFormat, _ = syscall.GetProcAddress(libkernel32, "GetDurationFormat")
	fnGetDurationFormatEx, _ = syscall.GetProcAddress(libkernel32, "GetDurationFormatEx")
	fnGetDllDirectoryW, _ = syscall.GetProcAddress(libkernel32, "GetDllDirectoryW")
	fnGetDllDirectoryA, _ = syscall.GetProcAddress(libkernel32, "GetDllDirectoryA")
	fnGetEnvironmentStringsW, _ = syscall.GetProcAddress(libkernel32, "GetEnvironmentStringsW")
	fnGetEnvironmentStringsA, _ = syscall.GetProcAddress(libkernel32, "GetEnvironmentStringsA")
	fnGetExitCodeProcess, _ = syscall.GetProcAddress(libkernel32, "GetExitCodeProcess")
	fnGetExitCodeThread, _ = syscall.GetProcAddress(libkernel32, "GetExitCodeThread")
	fnGetFileAttributesTransactedW, _ = syscall.GetProcAddress(libkernel32, "GetFileAttributesTransactedW")
	fnGetFileAttributesTransactedA, _ = syscall.GetProcAddress(libkernel32, "GetFileAttributesTransactedA")
	fnGetFileInformationByHandle, _ = syscall.GetProcAddress(libkernel32, "GetFileInformationByHandle")
	fnGetFileMUIInfo, _ = syscall.GetProcAddress(libkernel32, "GetFileMUIInfo")
	fnGetFileMUIPath, _ = syscall.GetProcAddress(libkernel32, "GetFileMUIPath")
	fnGetFullPathNameW, _ = syscall.GetProcAddress(libkernel32, "GetFullPathNameW")
	fnGetFullPathNameA, _ = syscall.GetProcAddress(libkernel32, "GetFullPathNameA")
	fnGetFullPathNameTransactedW, _ = syscall.GetProcAddress(libkernel32, "GetFullPathNameTransactedW")
	fnGetFullPathNameTransactedA, _ = syscall.GetProcAddress(libkernel32, "GetFullPathNameTransactedA")
	fnGetGeoInfoW, _ = syscall.GetProcAddress(libkernel32, "GetGeoInfoW")
	fnGetGeoInfoA, _ = syscall.GetProcAddress(libkernel32, "GetGeoInfoA")
	fnGetHandleInformation, _ = syscall.GetProcAddress(libkernel32, "GetHandleInformation")
	fnGetLargestConsoleWindowSize, _ = syscall.GetProcAddress(libkernel32, "GetLargestConsoleWindowSize")
	fnGetLastError, _ = syscall.GetProcAddress(libkernel32, "GetLastError")
	fnGetLocaleInfoW, _ = syscall.GetProcAddress(libkernel32, "GetLocaleInfoW")
	fnGetLocaleInfoA, _ = syscall.GetProcAddress(libkernel32, "GetLocaleInfoA")
	fnGetLogicalDrives, _ = syscall.GetProcAddress(libkernel32, "GetLogicalDrives")
	fnGetNLSVersion, _ = syscall.GetProcAddress(libkernel32, "GetNLSVersion")
	fnGetNLSVersionEx, _ = syscall.GetProcAddress(libkernel32, "GetNLSVersionEx")
	fnGetNumberFormatEx, _ = syscall.GetProcAddress(libkernel32, "GetNumberFormatEx")
	fnGetMailslotInfo, _ = syscall.GetProcAddress(libkernel32, "GetMailslotInfo")
	fnGetModuleFileNameW, _ = syscall.GetProcAddress(libkernel32, "GetModuleFileNameW")
	fnGetModuleFileNameA, _ = syscall.GetProcAddress(libkernel32, "GetModuleFileNameA")
	fnGetNamedPipeHandleStateW, _ = syscall.GetProcAddress(libkernel32, "GetNamedPipeHandleStateW")
	fnGetNamedPipeHandleStateA, _ = syscall.GetProcAddress(libkernel32, "GetNamedPipeHandleStateA")
	fnGetNamedPipeInfo, _ = syscall.GetProcAddress(libkernel32, "GetNamedPipeInfo")
	fnGetNumberFormatW, _ = syscall.GetProcAddress(libkernel32, "GetNumberFormatW")
	fnGetNumberFormatA, _ = syscall.GetProcAddress(libkernel32, "GetNumberFormatA")
	fnGetNumberOfConsoleInputEvents, _ = syscall.GetProcAddress(libkernel32, "GetNumberOfConsoleInputEvents")
	fnGetNumberOfConsoleMouseButtons, _ = syscall.GetProcAddress(libkernel32, "GetNumberOfConsoleMouseButtons")
	fnGetOEMCP, _ = syscall.GetProcAddress(libkernel32, "GetOEMCP")
	fnGetOverlappedResult, _ = syscall.GetProcAddress(libkernel32, "GetOverlappedResult")
	fnGetPriorityClass, _ = syscall.GetProcAddress(libkernel32, "GetPriorityClass")
	fnGetPrivateProfileIntW, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileIntW")
	fnGetPrivateProfileIntA, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileIntA")
	fnGetPrivateProfileSectionW, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileSectionW")
	fnGetPrivateProfileSectionA, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileSectionA")
	fnGetPrivateProfileSectionNamesW, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileSectionNamesW")
	fnGetPrivateProfileSectionNamesA, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileSectionNamesA")
	fnGetPrivateProfileStringW, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileStringW")
	fnGetPrivateProfileStringA, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileStringA")
	fnGetProcAddress, _ = syscall.GetProcAddress(libkernel32, "GetProcAddress")
	fnGetProcessAffinityMask, _ = syscall.GetProcAddress(libkernel32, "GetProcessAffinityMask")
	fnGetProcessHandleCount, _ = syscall.GetProcAddress(libkernel32, "GetProcessHandleCount")
	fnGetProcessHeap, _ = syscall.GetProcAddress(libkernel32, "GetProcessHeap")
	fnGetProcessHeaps, _ = syscall.GetProcAddress(libkernel32, "GetProcessHeaps")
	fnGetProcessId, _ = syscall.GetProcAddress(libkernel32, "GetProcessId")
	fnGetProcessPreferredUILanguages, _ = syscall.GetProcAddress(libkernel32, "GetProcessPreferredUILanguages")
	fnGetProcessPriorityBoost, _ = syscall.GetProcAddress(libkernel32, "GetProcessPriorityBoost")
	fnGetProcessShutdownParameters, _ = syscall.GetProcAddress(libkernel32, "GetProcessShutdownParameters")
	fnGetProcessTimes, _ = syscall.GetProcAddress(libkernel32, "GetProcessTimes")
	fnGetProcessVersion, _ = syscall.GetProcAddress(libkernel32, "GetProcessVersion")
	fnGetProcessWorkingSetSize, _ = syscall.GetProcAddress(libkernel32, "GetProcessWorkingSetSize")
	fnGetProfileIntW, _ = syscall.GetProcAddress(libkernel32, "GetProfileIntW")
	fnGetProfileIntA, _ = syscall.GetProcAddress(libkernel32, "GetProfileIntA")
	fnGetProfileSectionW, _ = syscall.GetProcAddress(libkernel32, "GetProfileSectionW")
	fnGetProfileSectionA, _ = syscall.GetProcAddress(libkernel32, "GetProfileSectionA")
	fnGetProfileStringW, _ = syscall.GetProcAddress(libkernel32, "GetProfileStringW")
	fnGetProfileStringA, _ = syscall.GetProcAddress(libkernel32, "GetProfileStringA")
	fnGetQueuedCompletionStatus, _ = syscall.GetProcAddress(libkernel32, "GetQueuedCompletionStatus")
	fnGetShortPathNameW, _ = syscall.GetProcAddress(libkernel32, "GetShortPathNameW")
	fnGetShortPathNameA, _ = syscall.GetProcAddress(libkernel32, "GetShortPathNameA")
	fnGetLongPathNameW, _ = syscall.GetProcAddress(libkernel32, "GetLongPathNameW")
	fnGetLongPathNameA, _ = syscall.GetProcAddress(libkernel32, "GetLongPathNameA")
	fnGetLongPathNameTransactedW, _ = syscall.GetProcAddress(libkernel32, "GetLongPathNameTransactedW")
	fnGetLongPathNameTransactedA, _ = syscall.GetProcAddress(libkernel32, "GetLongPathNameTransactedA")
	fnGetStdHandle, _ = syscall.GetProcAddress(libkernel32, "GetStdHandle")
	fnGetStringScripts, _ = syscall.GetProcAddress(libkernel32, "GetStringScripts")
	fnGetStringTypeExW, _ = syscall.GetProcAddress(libkernel32, "GetStringTypeExW")
	fnGetStringTypeExA, _ = syscall.GetProcAddress(libkernel32, "GetStringTypeExA")
	fnGetStringTypeA, _ = syscall.GetProcAddress(libkernel32, "GetStringTypeA")
	fnGetStringTypeW, _ = syscall.GetProcAddress(libkernel32, "GetStringTypeW")
	fnGetSystemDefaultLCID, _ = syscall.GetProcAddress(libkernel32, "GetSystemDefaultLCID")
	fnGetSystemDefaultLangID, _ = syscall.GetProcAddress(libkernel32, "GetSystemDefaultLangID")
	fnGetSystemDefaultLocaleName, _ = syscall.GetProcAddress(libkernel32, "GetSystemDefaultLocaleName")
	fnGetSystemDefaultUILanguage, _ = syscall.GetProcAddress(libkernel32, "GetSystemDefaultUILanguage")
	fnGetSystemDirectoryW, _ = syscall.GetProcAddress(libkernel32, "GetSystemDirectoryW")
	fnGetSystemDirectoryA, _ = syscall.GetProcAddress(libkernel32, "GetSystemDirectoryA")
	fnGetProductInfo, _ = syscall.GetProcAddress(libkernel32, "GetProductInfo")
	fnGetSystemRegistryQuota, _ = syscall.GetProcAddress(libkernel32, "GetSystemRegistryQuota")
	fnGetSystemTimes, _ = syscall.GetProcAddress(libkernel32, "GetSystemTimes")
	fnGetSystemPreferredUILanguages, _ = syscall.GetProcAddress(libkernel32, "GetSystemPreferredUILanguages")
	fnGetSystemPowerStatus, _ = syscall.GetProcAddress(libkernel32, "GetSystemPowerStatus")
	fnGetSystemTimeAdjustment, _ = syscall.GetProcAddress(libkernel32, "GetSystemTimeAdjustment")
	fnGetTapeParameters, _ = syscall.GetProcAddress(libkernel32, "GetTapeParameters")
	fnGetTapePosition, _ = syscall.GetProcAddress(libkernel32, "GetTapePosition")
	fnGetTapeStatus, _ = syscall.GetProcAddress(libkernel32, "GetTapeStatus")
	fnGetThreadContext, _ = syscall.GetProcAddress(libkernel32, "GetThreadContext")
	fnGetThreadIOPendingFlag, _ = syscall.GetProcAddress(libkernel32, "GetThreadIOPendingFlag")
	fnGetThreadLocale, _ = syscall.GetProcAddress(libkernel32, "GetThreadLocale")
	fnGetThreadPreferredUILanguages, _ = syscall.GetProcAddress(libkernel32, "GetThreadPreferredUILanguages")
	fnGetThreadPriority, _ = syscall.GetProcAddress(libkernel32, "GetThreadPriority")
	fnGetThreadPriorityBoost, _ = syscall.GetProcAddress(libkernel32, "GetThreadPriorityBoost")
	fnGetThreadSelectorEntry, _ = syscall.GetProcAddress(libkernel32, "GetThreadSelectorEntry")
	fnGetThreadTimes, _ = syscall.GetProcAddress(libkernel32, "GetThreadTimes")
	fnGetThreadUILanguage, _ = syscall.GetProcAddress(libkernel32, "GetThreadUILanguage")
	fnGetTickCount, _ = syscall.GetProcAddress(libkernel32, "GetTickCount")
	fnGetTimeFormatW, _ = syscall.GetProcAddress(libkernel32, "GetTimeFormatW")
	fnGetTimeFormatA, _ = syscall.GetProcAddress(libkernel32, "GetTimeFormatA")
	fnGetTimeFormatEx, _ = syscall.GetProcAddress(libkernel32, "GetTimeFormatEx")
	fnGetUILanguageInfo, _ = syscall.GetProcAddress(libkernel32, "GetUILanguageInfo")
	fnGetUserDefaultLocaleName, _ = syscall.GetProcAddress(libkernel32, "GetUserDefaultLocaleName")
	fnGetUserDefaultUILanguage, _ = syscall.GetProcAddress(libkernel32, "GetUserDefaultUILanguage")
	fnGetUserGeoID, _ = syscall.GetProcAddress(libkernel32, "GetUserGeoID")
	fnGetUserPreferredUILanguages, _ = syscall.GetProcAddress(libkernel32, "GetUserPreferredUILanguages")
	fnGetDynamicTimeZoneInformation, _ = syscall.GetProcAddress(libkernel32, "GetDynamicTimeZoneInformation")
	fnSetDynamicTimeZoneInformation, _ = syscall.GetProcAddress(libkernel32, "SetDynamicTimeZoneInformation")
	fnGetUserDefaultLCID, _ = syscall.GetProcAddress(libkernel32, "GetUserDefaultLCID")
	fnGetUserDefaultLangID, _ = syscall.GetProcAddress(libkernel32, "GetUserDefaultLangID")
	fnLCIDToLocaleName, _ = syscall.GetProcAddress(libkernel32, "LCIDToLocaleName")
	fnLocaleNameToLCID, _ = syscall.GetProcAddress(libkernel32, "LocaleNameToLCID")
	fnGetLocaleInfoEx, _ = syscall.GetProcAddress(libkernel32, "GetLocaleInfoEx")
	fnGetVersion, _ = syscall.GetProcAddress(libkernel32, "GetVersion")
	fnGetVersionExW, _ = syscall.GetProcAddress(libkernel32, "GetVersionExW")
	fnGetVersionExA, _ = syscall.GetProcAddress(libkernel32, "GetVersionExA")
	fnGetVolumePathNameW, _ = syscall.GetProcAddress(libkernel32, "GetVolumePathNameW")
	fnGetVolumePathNameA, _ = syscall.GetProcAddress(libkernel32, "GetVolumePathNameA")
	fnGetVolumePathNamesForVolumeNameW, _ = syscall.GetProcAddress(libkernel32, "GetVolumePathNamesForVolumeNameW")
	fnGetVolumePathNamesForVolumeNameA, _ = syscall.GetProcAddress(libkernel32, "GetVolumePathNamesForVolumeNameA")
	fnGetVolumeNameForVolumeMountPointW, _ = syscall.GetProcAddress(libkernel32, "GetVolumeNameForVolumeMountPointW")
	fnGetVolumeNameForVolumeMountPointA, _ = syscall.GetProcAddress(libkernel32, "GetVolumeNameForVolumeMountPointA")
	fnFindFirstVolumeW, _ = syscall.GetProcAddress(libkernel32, "FindFirstVolumeW")
	fnFindFirstVolumeA, _ = syscall.GetProcAddress(libkernel32, "FindFirstVolumeA")
	fnFindNextVolumeW, _ = syscall.GetProcAddress(libkernel32, "FindNextVolumeW")
	fnFindNextVolumeA, _ = syscall.GetProcAddress(libkernel32, "FindNextVolumeA")
	fnFindVolumeClose, _ = syscall.GetProcAddress(libkernel32, "FindVolumeClose")
	fnFindFirstVolumeMountPointW, _ = syscall.GetProcAddress(libkernel32, "FindFirstVolumeMountPointW")
	fnFindFirstVolumeMountPointA, _ = syscall.GetProcAddress(libkernel32, "FindFirstVolumeMountPointA")
	fnFindNextVolumeMountPointW, _ = syscall.GetProcAddress(libkernel32, "FindNextVolumeMountPointW")
	fnFindNextVolumeMountPointA, _ = syscall.GetProcAddress(libkernel32, "FindNextVolumeMountPointA")
	fnFindVolumeMountPointClose, _ = syscall.GetProcAddress(libkernel32, "FindVolumeMountPointClose")
	fnCreateSymbolicLinkW, _ = syscall.GetProcAddress(libkernel32, "CreateSymbolicLinkW")
	fnCreateSymbolicLinkA, _ = syscall.GetProcAddress(libkernel32, "CreateSymbolicLinkA")
	fnCreateSymbolicLinkTransactedW, _ = syscall.GetProcAddress(libkernel32, "CreateSymbolicLinkTransactedW")
	fnCreateSymbolicLinkTransactedA, _ = syscall.GetProcAddress(libkernel32, "CreateSymbolicLinkTransactedA")
	fnGetFinalPathNameByHandleW, _ = syscall.GetProcAddress(libkernel32, "GetFinalPathNameByHandleW")
	fnGetFinalPathNameByHandleA, _ = syscall.GetProcAddress(libkernel32, "GetFinalPathNameByHandleA")
	fnGetWindowsDirectoryW, _ = syscall.GetProcAddress(libkernel32, "GetWindowsDirectoryW")
	fnGetWindowsDirectoryA, _ = syscall.GetProcAddress(libkernel32, "GetWindowsDirectoryA")
	fnGlobalAddAtomW, _ = syscall.GetProcAddress(libkernel32, "GlobalAddAtomW")
	fnGlobalAddAtomA, _ = syscall.GetProcAddress(libkernel32, "GlobalAddAtomA")
	fnGlobalAlloc, _ = syscall.GetProcAddress(libkernel32, "GlobalAlloc")
	fnGlobalCompact, _ = syscall.GetProcAddress(libkernel32, "GlobalCompact")
	fnGlobalDeleteAtom, _ = syscall.GetProcAddress(libkernel32, "GlobalDeleteAtom")
	fnGlobalFindAtomW, _ = syscall.GetProcAddress(libkernel32, "GlobalFindAtomW")
	fnGlobalFindAtomA, _ = syscall.GetProcAddress(libkernel32, "GlobalFindAtomA")
	fnGlobalFlags, _ = syscall.GetProcAddress(libkernel32, "GlobalFlags")
	fnGlobalFree, _ = syscall.GetProcAddress(libkernel32, "GlobalFree")
	fnGlobalGetAtomNameW, _ = syscall.GetProcAddress(libkernel32, "GlobalGetAtomNameW")
	fnGlobalGetAtomNameA, _ = syscall.GetProcAddress(libkernel32, "GlobalGetAtomNameA")
	fnGlobalLock, _ = syscall.GetProcAddress(libkernel32, "GlobalLock")
	fnGlobalHandle, _ = syscall.GetProcAddress(libkernel32, "GlobalHandle")
	fnGlobalMemoryStatusEx, _ = syscall.GetProcAddress(libkernel32, "GlobalMemoryStatusEx")
	fnGlobalReAlloc, _ = syscall.GetProcAddress(libkernel32, "GlobalReAlloc")
	fnGlobalSize, _ = syscall.GetProcAddress(libkernel32, "GlobalSize")
	fnGlobalUnWire, _ = syscall.GetProcAddress(libkernel32, "GlobalUnWire")
	fnGlobalUnlock, _ = syscall.GetProcAddress(libkernel32, "GlobalUnlock")
	fnGlobalWire, _ = syscall.GetProcAddress(libkernel32, "GlobalWire")
	fnHeapAlloc, _ = syscall.GetProcAddress(libkernel32, "HeapAlloc")
	fnHeapCompact, _ = syscall.GetProcAddress(libkernel32, "HeapCompact")
	fnHeapCreate, _ = syscall.GetProcAddress(libkernel32, "HeapCreate")
	fnHeapDestroy, _ = syscall.GetProcAddress(libkernel32, "HeapDestroy")
	fnHeapFree, _ = syscall.GetProcAddress(libkernel32, "HeapFree")
	fnHeapLock, _ = syscall.GetProcAddress(libkernel32, "HeapLock")
	fnHeapReAlloc, _ = syscall.GetProcAddress(libkernel32, "HeapReAlloc")
	fnHeapSize, _ = syscall.GetProcAddress(libkernel32, "HeapSize")
	fnHeapUnlock, _ = syscall.GetProcAddress(libkernel32, "HeapUnlock")
	fnHeapValidate, _ = syscall.GetProcAddress(libkernel32, "HeapValidate")
	fnHeapWalk, _ = syscall.GetProcAddress(libkernel32, "HeapWalk")
	fnIdnToAscii, _ = syscall.GetProcAddress(libkernel32, "IdnToAscii")
	fnIdnToNameprepUnicode, _ = syscall.GetProcAddress(libkernel32, "IdnToNameprepUnicode")
	fnIdnToUnicode, _ = syscall.GetProcAddress(libkernel32, "IdnToUnicode")
	fnInitAtomTable, _ = syscall.GetProcAddress(libkernel32, "InitAtomTable")
	fnIsBadCodePtr, _ = syscall.GetProcAddress(libkernel32, "IsBadCodePtr")
	fnIsBadHugeReadPtr, _ = syscall.GetProcAddress(libkernel32, "IsBadHugeReadPtr")
	fnIsBadHugeWritePtr, _ = syscall.GetProcAddress(libkernel32, "IsBadHugeWritePtr")
	fnIsBadReadPtr, _ = syscall.GetProcAddress(libkernel32, "IsBadReadPtr")
	fnIsBadStringPtrW, _ = syscall.GetProcAddress(libkernel32, "IsBadStringPtrW")
	fnIsBadStringPtrA, _ = syscall.GetProcAddress(libkernel32, "IsBadStringPtrA")
	fnIsBadWritePtr, _ = syscall.GetProcAddress(libkernel32, "IsBadWritePtr")
	fnIsDBCSLeadByte, _ = syscall.GetProcAddress(libkernel32, "IsDBCSLeadByte")
	fnIsDBCSLeadByteEx, _ = syscall.GetProcAddress(libkernel32, "IsDBCSLeadByteEx")
	fnIsNLSDefinedString, _ = syscall.GetProcAddress(libkernel32, "IsNLSDefinedString")
	fnIsNormalizedString, _ = syscall.GetProcAddress(libkernel32, "IsNormalizedString")
	fnIsProcessorFeaturePresent, _ = syscall.GetProcAddress(libkernel32, "IsProcessorFeaturePresent")
	fnIsValidCodePage, _ = syscall.GetProcAddress(libkernel32, "IsValidCodePage")
	fnIsValidLanguageGroup, _ = syscall.GetProcAddress(libkernel32, "IsValidLanguageGroup")
	fnIsValidLocale, _ = syscall.GetProcAddress(libkernel32, "IsValidLocale")
	fnIsValidLocaleName, _ = syscall.GetProcAddress(libkernel32, "IsValidLocaleName")
	fnLCMapStringW, _ = syscall.GetProcAddress(libkernel32, "LCMapStringW")
	fnLCMapStringA, _ = syscall.GetProcAddress(libkernel32, "LCMapStringA")
	fnLCMapStringEx, _ = syscall.GetProcAddress(libkernel32, "LCMapStringEx")
	fnLoadModule, _ = syscall.GetProcAddress(libkernel32, "LoadModule")
	fnLocalAlloc, _ = syscall.GetProcAddress(libkernel32, "LocalAlloc")
	fnLocalCompact, _ = syscall.GetProcAddress(libkernel32, "LocalCompact")
	fnLocalFlags, _ = syscall.GetProcAddress(libkernel32, "LocalFlags")
	fnLocalFree, _ = syscall.GetProcAddress(libkernel32, "LocalFree")
	fnLocalLock, _ = syscall.GetProcAddress(libkernel32, "LocalLock")
	fnLocalReAlloc, _ = syscall.GetProcAddress(libkernel32, "LocalReAlloc")
	fnLocalShrink, _ = syscall.GetProcAddress(libkernel32, "LocalShrink")
	fnLocalSize, _ = syscall.GetProcAddress(libkernel32, "LocalSize")
	fnLocalUnlock, _ = syscall.GetProcAddress(libkernel32, "LocalUnlock")
	fnLockFile, _ = syscall.GetProcAddress(libkernel32, "LockFile")
	fnLockFileEx, _ = syscall.GetProcAddress(libkernel32, "LockFileEx")
	fnMoveFileW, _ = syscall.GetProcAddress(libkernel32, "MoveFileW")
	fnMoveFileA, _ = syscall.GetProcAddress(libkernel32, "MoveFileA")
	fnMoveFileExW, _ = syscall.GetProcAddress(libkernel32, "MoveFileExW")
	fnMoveFileExA, _ = syscall.GetProcAddress(libkernel32, "MoveFileExA")
	fnMoveFileTransactedW, _ = syscall.GetProcAddress(libkernel32, "MoveFileTransactedW")
	fnMoveFileTransactedA, _ = syscall.GetProcAddress(libkernel32, "MoveFileTransactedA")
	fnMoveFileWithProgressW, _ = syscall.GetProcAddress(libkernel32, "MoveFileWithProgressW")
	fnMoveFileWithProgressA, _ = syscall.GetProcAddress(libkernel32, "MoveFileWithProgressA")
	fnMulDiv, _ = syscall.GetProcAddress(libkernel32, "MulDiv")
	fnMultiByteToWideChar, _ = syscall.GetProcAddress(libkernel32, "MultiByteToWideChar")
	fnNormalizeString, _ = syscall.GetProcAddress(libkernel32, "NormalizeString")
	fnOpenEventW, _ = syscall.GetProcAddress(libkernel32, "OpenEventW")
	fnOpenEventA, _ = syscall.GetProcAddress(libkernel32, "OpenEventA")
	fnOpenFile, _ = syscall.GetProcAddress(libkernel32, "OpenFile")
	fnOpenFileMappingW, _ = syscall.GetProcAddress(libkernel32, "OpenFileMappingW")
	fnOpenFileMappingA, _ = syscall.GetProcAddress(libkernel32, "OpenFileMappingA")
	fnOpenMutexW, _ = syscall.GetProcAddress(libkernel32, "OpenMutexW")
	fnOpenMutexA, _ = syscall.GetProcAddress(libkernel32, "OpenMutexA")
	fnOpenProcess, _ = syscall.GetProcAddress(libkernel32, "OpenProcess")
	fnOpenRawW, _ = syscall.GetProcAddress(libkernel32, "OpenRawW")
	fnOpenRawA, _ = syscall.GetProcAddress(libkernel32, "OpenRawA")
	fnOpenSemaphoreW, _ = syscall.GetProcAddress(libkernel32, "OpenSemaphoreW")
	fnOpenWaitableTimerW, _ = syscall.GetProcAddress(libkernel32, "OpenWaitableTimerW")
	fnOpenSemaphoreA, _ = syscall.GetProcAddress(libkernel32, "OpenSemaphoreA")
	fnOpenWaitableTimerA, _ = syscall.GetProcAddress(libkernel32, "OpenWaitableTimerA")
	fnIsDebuggerPresent, _ = syscall.GetProcAddress(libkernel32, "IsDebuggerPresent")
	fnPeekConsoleInputW, _ = syscall.GetProcAddress(libkernel32, "PeekConsoleInputW")
	fnPeekConsoleInputA, _ = syscall.GetProcAddress(libkernel32, "PeekConsoleInputA")
	fnPeekNamedPipe, _ = syscall.GetProcAddress(libkernel32, "PeekNamedPipe")
	fnPostQueuedCompletionStatus, _ = syscall.GetProcAddress(libkernel32, "PostQueuedCompletionStatus")
	fnPrepareTape, _ = syscall.GetProcAddress(libkernel32, "PrepareTape")
	fnPulseEvent, _ = syscall.GetProcAddress(libkernel32, "PulseEvent")
	fnPurgeComm, _ = syscall.GetProcAddress(libkernel32, "PurgeComm")
	fnQueryPerformanceCounter, _ = syscall.GetProcAddress(libkernel32, "QueryPerformanceCounter")
	fnQueryPerformanceFrequency, _ = syscall.GetProcAddress(libkernel32, "QueryPerformanceFrequency")
	fnQueryRecoveryAgentsW, _ = syscall.GetProcAddress(libkernel32, "QueryRecoveryAgentsW")
	fnQueryRecoveryAgentsA, _ = syscall.GetProcAddress(libkernel32, "QueryRecoveryAgentsA")
	fnQueueUserAPC, _ = syscall.GetProcAddress(libkernel32, "QueueUserAPC")
	fnQueueUserWorkItem, _ = syscall.GetProcAddress(libkernel32, "QueueUserWorkItem")
	fnReadConsoleW, _ = syscall.GetProcAddress(libkernel32, "ReadConsoleW")
	fnReadConsoleA, _ = syscall.GetProcAddress(libkernel32, "ReadConsoleA")
	fnReadConsoleInputW, _ = syscall.GetProcAddress(libkernel32, "ReadConsoleInputW")
	fnReadConsoleInputA, _ = syscall.GetProcAddress(libkernel32, "ReadConsoleInputA")
	fnReadConsoleOutputW, _ = syscall.GetProcAddress(libkernel32, "ReadConsoleOutputW")
	fnReadConsoleOutputA, _ = syscall.GetProcAddress(libkernel32, "ReadConsoleOutputA")
	fnReadConsoleOutputAttribute, _ = syscall.GetProcAddress(libkernel32, "ReadConsoleOutputAttribute")
	fnReadConsoleOutputCharacterW, _ = syscall.GetProcAddress(libkernel32, "ReadConsoleOutputCharacterW")
	fnReadConsoleOutputCharacterA, _ = syscall.GetProcAddress(libkernel32, "ReadConsoleOutputCharacterA")
	fnReadDirectoryChangesW, _ = syscall.GetProcAddress(libkernel32, "ReadDirectoryChangesW")
	fnReadFile, _ = syscall.GetProcAddress(libkernel32, "ReadFile")
	fnReadFileEx, _ = syscall.GetProcAddress(libkernel32, "ReadFileEx")
	fnReadProcessMemory, _ = syscall.GetProcAddress(libkernel32, "ReadProcessMemory")
	fnReadRaw, _ = syscall.GetProcAddress(libkernel32, "ReadRaw")
	fnRegisterWaitForSingleObject, _ = syscall.GetProcAddress(libkernel32, "RegisterWaitForSingleObject")
	fnRegisterWaitForSingleObjectEx, _ = syscall.GetProcAddress(libkernel32, "RegisterWaitForSingleObjectEx")
	fnReleaseMutex, _ = syscall.GetProcAddress(libkernel32, "ReleaseMutex")
	fnReleaseSemaphore, _ = syscall.GetProcAddress(libkernel32, "ReleaseSemaphore")
	fnRemoveDirectoryW, _ = syscall.GetProcAddress(libkernel32, "RemoveDirectoryW")
	fnRemoveDirectoryA, _ = syscall.GetProcAddress(libkernel32, "RemoveDirectoryA")
	fnRemoveDirectoryTransactedW, _ = syscall.GetProcAddress(libkernel32, "RemoveDirectoryTransactedW")
	fnRemoveDirectoryTransactedA, _ = syscall.GetProcAddress(libkernel32, "RemoveDirectoryTransactedA")
	fnResetEvent, _ = syscall.GetProcAddress(libkernel32, "ResetEvent")
	fnResolveLocaleName, _ = syscall.GetProcAddress(libkernel32, "ResolveLocaleName")
	fnResumeThread, _ = syscall.GetProcAddress(libkernel32, "ResumeThread")
	fnSearchPathW, _ = syscall.GetProcAddress(libkernel32, "SearchPathW")
	fnSearchPathA, _ = syscall.GetProcAddress(libkernel32, "SearchPathA")
	fnSetCalendarInfoW, _ = syscall.GetProcAddress(libkernel32, "SetCalendarInfoW")
	fnSetCalendarInfoA, _ = syscall.GetProcAddress(libkernel32, "SetCalendarInfoA")
	fnSetCommBreak, _ = syscall.GetProcAddress(libkernel32, "SetCommBreak")
	fnSetCommConfig, _ = syscall.GetProcAddress(libkernel32, "SetCommConfig")
	fnSetCommMask, _ = syscall.GetProcAddress(libkernel32, "SetCommMask")
	fnSetCommState, _ = syscall.GetProcAddress(libkernel32, "SetCommState")
	fnSetCommTimeouts, _ = syscall.GetProcAddress(libkernel32, "SetCommTimeouts")
	fnSetComputerNameW, _ = syscall.GetProcAddress(libkernel32, "SetComputerNameW")
	fnSetComputerNameA, _ = syscall.GetProcAddress(libkernel32, "SetComputerNameA")
	fnSetConsoleActiveScreenBuffer, _ = syscall.GetProcAddress(libkernel32, "SetConsoleActiveScreenBuffer")
	fnSetConsoleCP, _ = syscall.GetProcAddress(libkernel32, "SetConsoleCP")
	fnSetConsoleCtrlHandler, _ = syscall.GetProcAddress(libkernel32, "SetConsoleCtrlHandler")
	fnSetConsoleCursorPosition, _ = syscall.GetProcAddress(libkernel32, "SetConsoleCursorPosition")
	fnSetConsoleMode, _ = syscall.GetProcAddress(libkernel32, "SetConsoleMode")
	fnSetConsoleOutputCP, _ = syscall.GetProcAddress(libkernel32, "SetConsoleOutputCP")
	fnSetConsoleScreenBufferSize, _ = syscall.GetProcAddress(libkernel32, "SetConsoleScreenBufferSize")
	fnSetConsoleTextAttribute, _ = syscall.GetProcAddress(libkernel32, "SetConsoleTextAttribute")
	fnSetConsoleTitleW, _ = syscall.GetProcAddress(libkernel32, "SetConsoleTitleW")
	fnSetConsoleTitleA, _ = syscall.GetProcAddress(libkernel32, "SetConsoleTitleA")
	fnSetCriticalSectionSpinCount, _ = syscall.GetProcAddress(libkernel32, "SetCriticalSectionSpinCount")
	fnSetCurrentDirectoryW, _ = syscall.GetProcAddress(libkernel32, "SetCurrentDirectoryW")
	fnSetCurrentDirectoryA, _ = syscall.GetProcAddress(libkernel32, "SetCurrentDirectoryA")
	fnSetDefaultCommConfigW, _ = syscall.GetProcAddress(libkernel32, "SetDefaultCommConfigW")
	fnSetDefaultCommConfigA, _ = syscall.GetProcAddress(libkernel32, "SetDefaultCommConfigA")
	fnSetDllDirectoryW, _ = syscall.GetProcAddress(libkernel32, "SetDllDirectoryW")
	fnSetDllDirectoryA, _ = syscall.GetProcAddress(libkernel32, "SetDllDirectoryA")
	fnSetEndOfFile, _ = syscall.GetProcAddress(libkernel32, "SetEndOfFile")
	fnSetEnvironmentVariableW, _ = syscall.GetProcAddress(libkernel32, "SetEnvironmentVariableW")
	fnSetEnvironmentVariableA, _ = syscall.GetProcAddress(libkernel32, "SetEnvironmentVariableA")
	fnGetThreadErrorMode, _ = syscall.GetProcAddress(libkernel32, "GetThreadErrorMode")
	fnSetThreadErrorMode, _ = syscall.GetProcAddress(libkernel32, "SetThreadErrorMode")
	fnSetEvent, _ = syscall.GetProcAddress(libkernel32, "SetEvent")
	fnSetFileAttributesW, _ = syscall.GetProcAddress(libkernel32, "SetFileAttributesW")
	fnSetFileAttributesA, _ = syscall.GetProcAddress(libkernel32, "SetFileAttributesA")
	fnSetFileAttributesTransactedW, _ = syscall.GetProcAddress(libkernel32, "SetFileAttributesTransactedW")
	fnSetFileAttributesTransactedA, _ = syscall.GetProcAddress(libkernel32, "SetFileAttributesTransactedA")
	fnSetFilePointer, _ = syscall.GetProcAddress(libkernel32, "SetFilePointer")
	fnSetFilePointerEx, _ = syscall.GetProcAddress(libkernel32, "SetFilePointerEx")
	fnSetFileTime, _ = syscall.GetProcAddress(libkernel32, "SetFileTime")
	fnSetHandleCount, _ = syscall.GetProcAddress(libkernel32, "SetHandleCount")
	fnSetHandleInformation, _ = syscall.GetProcAddress(libkernel32, "SetHandleInformation")
	fnSetLocalTime, _ = syscall.GetProcAddress(libkernel32, "SetLocalTime")
	fnSetLocaleInfoW, _ = syscall.GetProcAddress(libkernel32, "SetLocaleInfoW")
	fnSetLocaleInfoA, _ = syscall.GetProcAddress(libkernel32, "SetLocaleInfoA")
	fnSetMailslotInfo, _ = syscall.GetProcAddress(libkernel32, "SetMailslotInfo")
	fnSetNamedPipeHandleState, _ = syscall.GetProcAddress(libkernel32, "SetNamedPipeHandleState")
	fnSetPriorityClass, _ = syscall.GetProcAddress(libkernel32, "SetPriorityClass")
	fnSetProcessAffinityMask, _ = syscall.GetProcAddress(libkernel32, "SetProcessAffinityMask")
	fnSetProcessPreferredUILanguages, _ = syscall.GetProcAddress(libkernel32, "SetProcessPreferredUILanguages")
	fnSetProcessPriorityBoost, _ = syscall.GetProcAddress(libkernel32, "SetProcessPriorityBoost")
	fnSetProcessShutdownParameters, _ = syscall.GetProcAddress(libkernel32, "SetProcessShutdownParameters")
	fnSetProcessWorkingSetSize, _ = syscall.GetProcAddress(libkernel32, "SetProcessWorkingSetSize")
	fnSetStdHandle, _ = syscall.GetProcAddress(libkernel32, "SetStdHandle")
	fnSetSystemPowerState, _ = syscall.GetProcAddress(libkernel32, "SetSystemPowerState")
	fnSetSystemTime, _ = syscall.GetProcAddress(libkernel32, "SetSystemTime")
	fnSetSystemTimeAdjustment, _ = syscall.GetProcAddress(libkernel32, "SetSystemTimeAdjustment")
	fnSetTapeParameters, _ = syscall.GetProcAddress(libkernel32, "SetTapeParameters")
	fnSetTapePosition, _ = syscall.GetProcAddress(libkernel32, "SetTapePosition")
	fnSetThreadAffinityMask, _ = syscall.GetProcAddress(libkernel32, "SetThreadAffinityMask")
	fnSetThreadContext, _ = syscall.GetProcAddress(libkernel32, "SetThreadContext")
	fnSetThreadExecutionState, _ = syscall.GetProcAddress(libkernel32, "SetThreadExecutionState")
	fnSetThreadIdealProcessor, _ = syscall.GetProcAddress(libkernel32, "SetThreadIdealProcessor")
	fnSetThreadLocale, _ = syscall.GetProcAddress(libkernel32, "SetThreadLocale")
	fnSetThreadPreferredUILanguages, _ = syscall.GetProcAddress(libkernel32, "SetThreadPreferredUILanguages")
	fnSetThreadPriority, _ = syscall.GetProcAddress(libkernel32, "SetThreadPriority")
	fnSetThreadPriorityBoost, _ = syscall.GetProcAddress(libkernel32, "SetThreadPriorityBoost")
	fnSetThreadUILanguage, _ = syscall.GetProcAddress(libkernel32, "SetThreadUILanguage")
	fnSetTimerQueueTimer, _ = syscall.GetProcAddress(libkernel32, "SetTimerQueueTimer")
	fnSetTimeZoneInformation, _ = syscall.GetProcAddress(libkernel32, "SetTimeZoneInformation")
	fnSetUserGeoID, _ = syscall.GetProcAddress(libkernel32, "SetUserGeoID")
	fnSetUnhandledExceptionFilter, _ = syscall.GetProcAddress(libkernel32, "SetUnhandledExceptionFilter")
	fnSetVolumeLabelW, _ = syscall.GetProcAddress(libkernel32, "SetVolumeLabelW")
	fnSetVolumeLabelA, _ = syscall.GetProcAddress(libkernel32, "SetVolumeLabelA")
	fnSetWaitableTimer, _ = syscall.GetProcAddress(libkernel32, "SetWaitableTimer")
	fnSetupComm, _ = syscall.GetProcAddress(libkernel32, "SetupComm")
	fnSignalObjectAndWait, _ = syscall.GetProcAddress(libkernel32, "SignalObjectAndWait")
	fnSleepEx, _ = syscall.GetProcAddress(libkernel32, "SleepEx")
	fnSuspendThread, _ = syscall.GetProcAddress(libkernel32, "SuspendThread")
	fnSwitchToThread, _ = syscall.GetProcAddress(libkernel32, "SwitchToThread")
	fnSystemTimeToFileTime, _ = syscall.GetProcAddress(libkernel32, "SystemTimeToFileTime")
	fnSystemTimeToTzSpecificLocalTime, _ = syscall.GetProcAddress(libkernel32, "SystemTimeToTzSpecificLocalTime")
	fnTerminateProcess, _ = syscall.GetProcAddress(libkernel32, "TerminateProcess")
	fnTerminateThread, _ = syscall.GetProcAddress(libkernel32, "TerminateThread")
	fnTlsAlloc, _ = syscall.GetProcAddress(libkernel32, "TlsAlloc")
	fnTlsFree, _ = syscall.GetProcAddress(libkernel32, "TlsFree")
	fnTlsGetValue, _ = syscall.GetProcAddress(libkernel32, "TlsGetValue")
	fnTlsSetValue, _ = syscall.GetProcAddress(libkernel32, "TlsSetValue")
	fnTransactNamedPipe, _ = syscall.GetProcAddress(libkernel32, "TransactNamedPipe")
	fnTransmitCommChar, _ = syscall.GetProcAddress(libkernel32, "TransmitCommChar")
	fnUnhandledExceptionFilter, _ = syscall.GetProcAddress(libkernel32, "UnhandledExceptionFilter")
	fnUnlockFile, _ = syscall.GetProcAddress(libkernel32, "UnlockFile")
	fnUnlockFileEx, _ = syscall.GetProcAddress(libkernel32, "UnlockFileEx")
	fnUnregisterWait, _ = syscall.GetProcAddress(libkernel32, "UnregisterWait")
	fnUnregisterWaitEx, _ = syscall.GetProcAddress(libkernel32, "UnregisterWaitEx")
	fnUpdateResourceW, _ = syscall.GetProcAddress(libkernel32, "UpdateResourceW")
	fnUpdateResourceA, _ = syscall.GetProcAddress(libkernel32, "UpdateResourceA")
	fnVerLanguageNameW, _ = syscall.GetProcAddress(libkernel32, "VerLanguageNameW")
	fnVerLanguageNameA, _ = syscall.GetProcAddress(libkernel32, "VerLanguageNameA")
	fnVerifyScripts, _ = syscall.GetProcAddress(libkernel32, "VerifyScripts")
	fnVerifyVersionInfoW, _ = syscall.GetProcAddress(libkernel32, "VerifyVersionInfoW")
	fnVerifyVersionInfoA, _ = syscall.GetProcAddress(libkernel32, "VerifyVersionInfoA")
	fnVerSetConditionMask, _ = syscall.GetProcAddress(libkernel32, "VerSetConditionMask")
	fnVirtualAlloc, _ = syscall.GetProcAddress(libkernel32, "VirtualAlloc")
	fnVirtualAllocEx, _ = syscall.GetProcAddress(libkernel32, "VirtualAllocEx")
	fnVirtualFree, _ = syscall.GetProcAddress(libkernel32, "VirtualFree")
	fnVirtualFreeEx, _ = syscall.GetProcAddress(libkernel32, "VirtualFreeEx")
	fnVirtualLock, _ = syscall.GetProcAddress(libkernel32, "VirtualLock")
	fnVirtualQuery, _ = syscall.GetProcAddress(libkernel32, "VirtualQuery")
	fnVirtualQueryEx, _ = syscall.GetProcAddress(libkernel32, "VirtualQueryEx")
	fnVirtualUnlock, _ = syscall.GetProcAddress(libkernel32, "VirtualUnlock")
	fnWaitCommEvent, _ = syscall.GetProcAddress(libkernel32, "WaitCommEvent")
	fnWaitForDebugEvent, _ = syscall.GetProcAddress(libkernel32, "WaitForDebugEvent")
	fnWaitForMultipleObjects, _ = syscall.GetProcAddress(libkernel32, "WaitForMultipleObjects")
	fnWaitForMultipleObjectsEx, _ = syscall.GetProcAddress(libkernel32, "WaitForMultipleObjectsEx")
	fnWaitForSingleObject, _ = syscall.GetProcAddress(libkernel32, "WaitForSingleObject")
	fnWaitForSingleObjectEx, _ = syscall.GetProcAddress(libkernel32, "WaitForSingleObjectEx")
	fnSleepConditionVariableCS, _ = syscall.GetProcAddress(libkernel32, "SleepConditionVariableCS")
	fnWaitNamedPipeW, _ = syscall.GetProcAddress(libkernel32, "WaitNamedPipeW")
	fnWaitNamedPipeA, _ = syscall.GetProcAddress(libkernel32, "WaitNamedPipeA")
	fnWideCharToMultiByte, _ = syscall.GetProcAddress(libkernel32, "WideCharToMultiByte")
	fnWinExec, _ = syscall.GetProcAddress(libkernel32, "WinExec")
	fnWriteConsoleW, _ = syscall.GetProcAddress(libkernel32, "WriteConsoleW")
	fnWriteConsoleA, _ = syscall.GetProcAddress(libkernel32, "WriteConsoleA")
	fnWriteConsoleInputW, _ = syscall.GetProcAddress(libkernel32, "WriteConsoleInputW")
	fnWriteConsoleInputA, _ = syscall.GetProcAddress(libkernel32, "WriteConsoleInputA")
	fnWriteConsoleOutputW, _ = syscall.GetProcAddress(libkernel32, "WriteConsoleOutputW")
	fnWriteConsoleOutputA, _ = syscall.GetProcAddress(libkernel32, "WriteConsoleOutputA")
	fnWriteConsoleOutputAttribute, _ = syscall.GetProcAddress(libkernel32, "WriteConsoleOutputAttribute")
	fnWriteConsoleOutputCharacterW, _ = syscall.GetProcAddress(libkernel32, "WriteConsoleOutputCharacterW")
	fnWriteConsoleOutputCharacterA, _ = syscall.GetProcAddress(libkernel32, "WriteConsoleOutputCharacterA")
	fnWriteFile, _ = syscall.GetProcAddress(libkernel32, "WriteFile")
	fnWriteFileEx, _ = syscall.GetProcAddress(libkernel32, "WriteFileEx")
	fnWritePrivateProfileSectionW, _ = syscall.GetProcAddress(libkernel32, "WritePrivateProfileSectionW")
	fnWritePrivateProfileSectionA, _ = syscall.GetProcAddress(libkernel32, "WritePrivateProfileSectionA")
	fnWritePrivateProfileStringW, _ = syscall.GetProcAddress(libkernel32, "WritePrivateProfileStringW")
	fnWritePrivateProfileStringA, _ = syscall.GetProcAddress(libkernel32, "WritePrivateProfileStringA")
	fnWriteProcessMemory, _ = syscall.GetProcAddress(libkernel32, "WriteProcessMemory")
	fnWriteProfileSectionW, _ = syscall.GetProcAddress(libkernel32, "WriteProfileSectionW")
	fnWriteProfileSectionA, _ = syscall.GetProcAddress(libkernel32, "WriteProfileSectionA")
	fnWriteProfileStringW, _ = syscall.GetProcAddress(libkernel32, "WriteProfileStringW")
	fnWriteProfileStringA, _ = syscall.GetProcAddress(libkernel32, "WriteProfileStringA")
	fnWriteRaw, _ = syscall.GetProcAddress(libkernel32, "WriteRaw")
	fnWriteTapemark, _ = syscall.GetProcAddress(libkernel32, "WriteTapemark")
	fn_hread, _ = syscall.GetProcAddress(libkernel32, "_hread")
	fn_hwrite, _ = syscall.GetProcAddress(libkernel32, "_hwrite")
	fn_lclose, _ = syscall.GetProcAddress(libkernel32, "_lclose")
	fn_lcreat, _ = syscall.GetProcAddress(libkernel32, "_lcreat")
	fn_llseek, _ = syscall.GetProcAddress(libkernel32, "_llseek")
	fn_lopen, _ = syscall.GetProcAddress(libkernel32, "_lopen")
	fn_lread, _ = syscall.GetProcAddress(libkernel32, "_lread")
	fn_lwrite, _ = syscall.GetProcAddress(libkernel32, "_lwrite")
	fnlstrcatW, _ = syscall.GetProcAddress(libkernel32, "lstrcatW")
	fnlstrcatA, _ = syscall.GetProcAddress(libkernel32, "lstrcatA")
	fnlstrcmpW, _ = syscall.GetProcAddress(libkernel32, "lstrcmpW")
	fnlstrcmpA, _ = syscall.GetProcAddress(libkernel32, "lstrcmpA")
	fnlstrcmpiW, _ = syscall.GetProcAddress(libkernel32, "lstrcmpiW")
	fnlstrcmpiA, _ = syscall.GetProcAddress(libkernel32, "lstrcmpiA")
	fnlstrcpyW, _ = syscall.GetProcAddress(libkernel32, "lstrcpyW")
	fnlstrcpyA, _ = syscall.GetProcAddress(libkernel32, "lstrcpyA")
	fnlstrcpynW, _ = syscall.GetProcAddress(libkernel32, "lstrcpynW")
	fnlstrcpynA, _ = syscall.GetProcAddress(libkernel32, "lstrcpynA")
	fnlstrlenW, _ = syscall.GetProcAddress(libkernel32, "lstrlenW")
	fnlstrlenA, _ = syscall.GetProcAddress(libkernel32, "lstrlenA")
	fnGetPrivateProfileStructW, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileStructW")
	fnGetPrivateProfileStructA, _ = syscall.GetProcAddress(libkernel32, "GetPrivateProfileStructA")
	fnWritePrivateProfileStructW, _ = syscall.GetProcAddress(libkernel32, "WritePrivateProfileStructW")
	fnWritePrivateProfileStructA, _ = syscall.GetProcAddress(libkernel32, "WritePrivateProfileStructA")
	fnCreateActCtxW, _ = syscall.GetProcAddress(libkernel32, "CreateActCtxW")
	fnCreateActCtxA, _ = syscall.GetProcAddress(libkernel32, "CreateActCtxA")
	fnZombifyActCtx, _ = syscall.GetProcAddress(libkernel32, "ZombifyActCtx")
	fnActivateActCtx, _ = syscall.GetProcAddress(libkernel32, "ActivateActCtx")
	fnDeactivateActCtx, _ = syscall.GetProcAddress(libkernel32, "DeactivateActCtx")
	fnGetCurrentActCtx, _ = syscall.GetProcAddress(libkernel32, "GetCurrentActCtx")
	fnProcessIdToSessionId, _ = syscall.GetProcAddress(libkernel32, "ProcessIdToSessionId")
	fnWTSGetActiveConsoleSessionId, _ = syscall.GetProcAddress(libkernel32, "WTSGetActiveConsoleSessionId")
	fnIsWow64Process, _ = syscall.GetProcAddress(libkernel32, "IsWow64Process")
	fnGetLogicalProcessorInformation, _ = syscall.GetProcAddress(libkernel32, "GetLogicalProcessorInformation")
	fnGetLogicalProcessorInformationEx, _ = syscall.GetProcAddress(libkernel32, "GetLogicalProcessorInformationEx")
	fnGetNumaHighestNodeNumber, _ = syscall.GetProcAddress(libkernel32, "GetNumaHighestNodeNumber")
	fnGetNumaProcessorNode, _ = syscall.GetProcAddress(libkernel32, "GetNumaProcessorNode")
	fnGetNumaAvailableMemoryNode, _ = syscall.GetProcAddress(libkernel32, "GetNumaAvailableMemoryNode")
	fnGetNumaNodeProcessorMask, _ = syscall.GetProcAddress(libkernel32, "GetNumaNodeProcessorMask")
	fnSetProcessDEPPolicy, _ = syscall.GetProcAddress(libkernel32, "SetProcessDEPPolicy")
	fnGetNamedPipeServerSessionId, _ = syscall.GetProcAddress(libkernel32, "GetNamedPipeServerSessionId")
	fnGetNamedPipeServerProcessId, _ = syscall.GetProcAddress(libkernel32, "GetNamedPipeServerProcessId")
	fnGetNamedPipeClientProcessId, _ = syscall.GetProcAddress(libkernel32, "GetNamedPipeClientProcessId")
	fnGetNamedPipeClientSessionId, _ = syscall.GetProcAddress(libkernel32, "GetNamedPipeClientSessionId")
	fnGetNamedPipeClientComputerNameW, _ = syscall.GetProcAddress(libkernel32, "GetNamedPipeClientComputerNameW")
	fnGetNamedPipeClientComputerNameA, _ = syscall.GetProcAddress(libkernel32, "GetNamedPipeClientComputerNameA")
	fnMoveMemory, _ = syscall.GetProcAddress(libkernel32, "RtlMoveMemory")
}

func MoveMemory(destination, source unsafe.Pointer, length uintptr) {
	syscall.Syscall(fnMoveMemory, 3,
		uintptr(unsafe.Pointer(destination)),
		uintptr(source),
		uintptr(length))
}

func GetModuleHandle(lpModuleName string) syscall.Handle {
	initKernel32()
	var lpName uintptr
	if lpModuleName != "" {
		lpName = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpModuleName)))
	} else {
		lpName = 0
	}

	r1, _, _ := syscall.Syscall(fnGetModuleHandleW, 1, lpName, 0, 0)
	return syscall.Handle(r1)
}

func GetCurrentProcessId() uint32 {
	initKernel32()
	r1, _, _ := syscall.Syscall(fnGetCurrentProcessId, 0, 0, 0, 0)
	return uint32(r1)
}


func GetCurrentProcess() syscall.Handle {
	initKernel32()
	r1, _, _ := syscall.Syscall(fnGetCurrentProcess, 0, 0, 0, 0)
	return syscall.Handle(r1)
}

func GlobalAddAtom(atomtag string) ATOM {
	initKernel32()
	r1, _, _ := syscall.Syscall(fnGlobalAddAtomW, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(atomtag))), 0, 0)
	return ATOM(r1)
}

func GlobalFindAtom(atomtag string) ATOM {
	initKernel32()
	r1, _, _ := syscall.Syscall(fnGlobalFindAtomW, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(atomtag))), 0, 0)
	return ATOM(r1)
}

func GlobalDeleteAtom(atom ATOM) bool {
	initKernel32()
	r1, _, _ := syscall.Syscall(fnGlobalDeleteAtom, 1, uintptr(atom), 0, 0)
	return r1 != 0
}

func MulDiv(nNumber, nNumerator, nDenominator int32)int32  {
	initKernel32()
	r1, _, _ := syscall.Syscall(fnMulDiv, 3, uintptr(nNumber), uintptr(nNumerator), uintptr(nDenominator))
	return int32(r1)
}

func GetLastError()uint32  {
	initKernel32()
	r1, _, _ := syscall.Syscall(fnGetLastError, 0, 0,0,0)
	return uint32(r1)
}