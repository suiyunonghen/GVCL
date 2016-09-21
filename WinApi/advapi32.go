package WinApi

import (
	"syscall"
	"unsafe"
)

type HKEY uintptr
type REGSAM uint32
type PSID uintptr
type WORD uint16
type DWORD uint32
type GSecurityAttributes  struct{
	NLength uint32
	SecurityDescriptor uintptr
	InheritHandle bool
}

type GFileTime struct {
	LowDateTime uint32
	HighDateTime uint32
}

type GACL struct {
	AclRevision byte
	Sbz1 byte
	AclSize uint16
	AceCount uint16
	Sbz2 uint16
}

type GSecurityDescriptor struct{
	Revision byte
	Sbz1 byte
	Control uint16
	Owner PSID
	Group PSID
	Sacl *GACL
	Dacl *GACL

}
var (
	libadvapi32                                 syscall.Handle
	fnAbortSystemShutdownW                      uintptr
	fnAbortSystemShutdownA                      uintptr
	fnAddAuditAccessAce                         uintptr
	fnAddAuditAccessAceEx                       uintptr
	fnAddAuditAccessObjectAce                   uintptr
	fnAdjustTokenGroups                         uintptr
	fnAllocateLocallyUniqueId                   uintptr
	fnAreAllAccessesGranted                     uintptr
	fnAreAnyAccessesGranted                     uintptr
	fnBackupEventLogW                           uintptr
	fnBackupEventLogA                           uintptr
	fnClearEventLogW                            uintptr
	fnClearEventLogA                            uintptr
	fnCloseEventLog                             uintptr
	fnConvertToAutoInheritPrivateObjectSecurity uintptr
	fnCopySid                                   uintptr
	fnConvertSidToStringSidW                    uintptr
	fnConvertSidToStringSidA                    uintptr
	fnConvertStringSidToSidW                    uintptr
	fnConvertStringSidToSidA                    uintptr
	fnCreatePrivateObjectSecurity               uintptr
	fnCreatePrivateObjectSecurityEx             uintptr
	fnCreateProcessAsUserW                      uintptr
	fnCreateProcessAsUserA                      uintptr
	fnDecryptFileW                              uintptr
	fnDecryptFileA                              uintptr
	fnDeleteAce                                 uintptr
	fnDeregisterEventSource                     uintptr
	fnDestroyPrivateObjectSecurity              uintptr
	fnDuplicateToken                            uintptr
	fnDuplicateTokenEx                          uintptr
	fnEncryptFileW                              uintptr
	fnEncryptFileA                              uintptr
	fnEqualPrefixSid                            uintptr
	fnEqualSid                                  uintptr
	fnFindFirstFreeAce                          uintptr
	fnFreeSid                                   uintptr
	fnGetAce                                    uintptr
	fnGetCurrentHwProfileW                      uintptr
	fnGetCurrentHwProfileA                      uintptr
	fnGetFileSecurityW                          uintptr
	fnGetFileSecurityA                          uintptr
	fnGetKernelObjectSecurity                   uintptr
	fnGetLengthSid                              uintptr
	fnGetNumberOfEventLogRecords                uintptr
	fnGetOldestEventLogRecord                   uintptr
	fnGetPrivateObjectSecurity                  uintptr
	fnGetSecurityDescriptorControl              uintptr
	fnGetSecurityDescriptorDacl                 uintptr
	fnGetSecurityDescriptorGroup                uintptr
	fnGetSecurityDescriptorLength               uintptr
	fnGetSecurityDescriptorOwner                uintptr
	fnGetSecurityDescriptorSacl                 uintptr
	fnGetSidIdentifierAuthority                 uintptr
	fnGetSidLengthRequired                      uintptr
	fnGetSidSubAuthority                        uintptr
	fnGetSidSubAuthorityCount                   uintptr
	fnGetTokenInformation                       uintptr
	fnGetUserNameW                              uintptr
	fnGetUserNameA                              uintptr
	fnImpersonateLoggedOnUser                   uintptr
	fnImpersonateNamedPipeClient                uintptr
	fnImpersonateSelf                           uintptr
	fnInitializeAcl                             uintptr
	fnInitializeSecurityDescriptor              uintptr
	fnInitiateSystemShutdownW                   uintptr
	fnInitiateSystemShutdownA                   uintptr
	fnIsTextUnicode                             uintptr
	fnIsValidAcl                                uintptr
	fnIsValidSecurityDescriptor                 uintptr
	fnIsValidSid                                uintptr
	fnLogonUserW                                uintptr
	fnLogonUserA                                uintptr
	fnLookupAccountNameW                        uintptr
	fnLookupAccountNameA                        uintptr
	fnLookupAccountSidW                         uintptr
	fnLookupAccountSidA                         uintptr
	fnLookupPrivilegeDisplayNameW               uintptr
	fnLookupPrivilegeDisplayNameA               uintptr
	fnLookupPrivilegeNameW                      uintptr
	fnLookupPrivilegeNameA                      uintptr
	fnLookupPrivilegeValueW                     uintptr
	fnLookupPrivilegeValueA                     uintptr
	fnMakeAbsoluteSD                            uintptr
	fnMakeSelfRelativeSD                        uintptr
	fnNotifyChangeEventLog                      uintptr
	fnObjectCloseAuditAlarmW                    uintptr
	fnObjectCloseAuditAlarmA                    uintptr
	fnObjectDeleteAuditAlarmW                   uintptr
	fnObjectDeleteAuditAlarmA                   uintptr
	fnObjectOpenAuditAlarmW                     uintptr
	fnObjectOpenAuditAlarmA                     uintptr
	fnObjectPrivilegeAuditAlarmW                uintptr
	fnObjectPrivilegeAuditAlarmA                uintptr
	fnOpenBackupEventLogW                       uintptr
	fnOpenBackupEventLogA                       uintptr
	fnOpenEventLogW                             uintptr
	fnOpenEventLogA                             uintptr
	fnOpenProcessToken                          uintptr
	fnOpenThreadToken                           uintptr
	fnPrivilegeCheck                            uintptr
	fnPrivilegedServiceAuditAlarmW              uintptr
	fnPrivilegedServiceAuditAlarmA              uintptr
	fnReadEventLogW                             uintptr
	fnReadEventLogA                             uintptr
	fnRegOverridePredefKey                      uintptr
	fnRegOpenUserClassesRoot                    uintptr
	fnRegOpenCurrentUser                        uintptr
	fnRegDisablePredefinedCache                 uintptr
	fnRegQueryMultipleValuesW                   uintptr
	fnRegQueryMultipleValuesA                   uintptr
	fnRegisterEventSourceW                      uintptr
	fnRegisterEventSourceA                      uintptr
	fnReportEventW                              uintptr
	fnReportEventA                              uintptr
	fnRevertToSelf                              uintptr
	fnSetAclInformation                         uintptr
	fnSetFileSecurityW                          uintptr
	fnSetFileSecurityA                          uintptr
	fnSetKernelObjectSecurity                   uintptr
	fnSetPrivateObjectSecurity                  uintptr
	fnSetPrivateObjectSecurityEx                uintptr
	fnSetSecurityDescriptorControl              uintptr
	fnSetSecurityDescriptorDacl                 uintptr
	fnSetSecurityDescriptorGroup                uintptr
	fnSetSecurityDescriptorOwner                uintptr
	fnSetSecurityDescriptorSacl                 uintptr
	fnSetThreadToken                            uintptr
	fnSetTokenInformation                       uintptr
	fnRegCloseKey				    uintptr
	fnRegOpenKeyEx				    uintptr
	fnRegFlushKey				    uintptr
	fnRegCreateKeyEx			    uintptr
	fnRegQueryInfoKey			    uintptr
	fnRegEnumKeyEx				    uintptr
	fnRegEnumValue				    uintptr
	fnRegLoadKey				    uintptr
	fnRegUnLoadKey				    uintptr
	fnRegQueryValueEx		            uintptr
	fnRegQueryValue			   	    uintptr
	fnRegOpenKey				    uintptr
	fnRegConnectRegistry			    uintptr
	fnRegDeleteKey			   	    uintptr
	fnRegDeleteKeyEx			    uintptr
	fnRegDeleteValue			    uintptr
	fnRegGetKeySecurity			    uintptr
	fnRegNotifyChangeKeyValue		    uintptr
	fnRegQueryMultipleValues		    uintptr
	fnRegReplaceKey				    uintptr
	fnRegRestoreKey				    uintptr
	fnRegSaveKeyEx				    uintptr
	fnRegSaveKey				    uintptr
	fnRegSetKeySecurity			    uintptr
	fnRegSetValue				    uintptr
	fnRegSetValueEx				    uintptr
	fnRegisterEventSource			    uintptr
	fnReportEvent				    uintptr
	fnSetFileSecurity			    uintptr
)

func init() {
	libadvapi32, _ = syscall.LoadLibrary("advapi32.dll")
	fnRegCloseKey, _ = syscall.GetProcAddress(libadvapi32, "RegCloseKey")
	fnRegFlushKey, _ = syscall.GetProcAddress(libadvapi32, "RegFlushKey")
	fnRegCreateKeyEx, _ = syscall.GetProcAddress(libadvapi32, "RegCreateKeyExW")
	fnRegQueryInfoKey, _ = syscall.GetProcAddress(libadvapi32, "RegQueryInfoKeyW")
	fnRegEnumKeyEx, _ = syscall.GetProcAddress(libadvapi32, "RegEnumKeyExW")
	fnRegEnumValue, _ = syscall.GetProcAddress(libadvapi32, "RegEnumValueW")
	fnRegLoadKey, _ = syscall.GetProcAddress(libadvapi32, "RegLoadKeyW")
	fnRegUnLoadKey, _ = syscall.GetProcAddress(libadvapi32, "RegUnLoadKeyW")
	fnRegQueryValueEx, _ = syscall.GetProcAddress(libadvapi32, "RegQueryValueExW")
	fnRegQueryValue, _ = syscall.GetProcAddress(libadvapi32, "RegQueryValueW")
	fnRegOpenKeyEx, _ = syscall.GetProcAddress(libadvapi32, "RegOpenKeyExW")
	fnRegOpenKey, _ = syscall.GetProcAddress(libadvapi32, "RegOpenKeyW")
	fnRegConnectRegistry, _ = syscall.GetProcAddress(libadvapi32, "RegConnectRegistryW")
	fnRegDeleteKey, _ = syscall.GetProcAddress(libadvapi32, "RegDeleteKeyW")
	fnRegDeleteKeyEx, _ = syscall.GetProcAddress(libadvapi32, "RegDeleteKeyExW")
	fnRegDeleteValue, _ = syscall.GetProcAddress(libadvapi32, "RegDeleteValueW")
	fnRegGetKeySecurity, _ = syscall.GetProcAddress(libadvapi32, "RegGetKeySecurity")
	fnRegNotifyChangeKeyValue, _ = syscall.GetProcAddress(libadvapi32, "RegNotifyChangeKeyValue")
	fnRegQueryMultipleValues, _ = syscall.GetProcAddress(libadvapi32, "RegQueryMultipleValuesW")
	fnRegReplaceKey, _ = syscall.GetProcAddress(libadvapi32, "RegReplaceKeyW")
	fnRegRestoreKey, _ = syscall.GetProcAddress(libadvapi32, "RegRestoreKeyW")
	fnRegSaveKeyEx, _ = syscall.GetProcAddress(libadvapi32, "RegSaveKeyExW")
	fnRegSaveKeyEx, _ = syscall.GetProcAddress(libadvapi32, "RegSaveKeyW")
	fnRegSetKeySecurity, _ = syscall.GetProcAddress(libadvapi32, "RegSetKeySecurity")
	fnRegSetValue, _ = syscall.GetProcAddress(libadvapi32, "RegSetValueW")
	fnRegSetValueEx, _ = syscall.GetProcAddress(libadvapi32, "RegSetValueExW")
	fnRegisterEventSource,_ = syscall.GetProcAddress(libadvapi32, "RegisterEventSourceW")
	fnReportEvent,_ = syscall.GetProcAddress(libadvapi32, "ReportEventW")
	fnRevertToSelf,_ = syscall.GetProcAddress(libadvapi32, "RevertToSelf")
	fnSetAclInformation,_ = syscall.GetProcAddress(libadvapi32, "SetAclInformation")
	fnSetFileSecurity,_ = syscall.GetProcAddress(libadvapi32, "SetFileSecurityW")
	fnSetKernelObjectSecurity, _ = syscall.GetProcAddress(libadvapi32, "SetKernelObjectSecurity")
	fnSetPrivateObjectSecurity, _ = syscall.GetProcAddress(libadvapi32, "SetPrivateObjectSecurity")
	fnSetPrivateObjectSecurityEx, _ = syscall.GetProcAddress(libadvapi32, "SetPrivateObjectSecurityEx")
	fnSetSecurityDescriptorControl, _ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorControl")
	fnSetSecurityDescriptorDacl, _ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorDacl")
	fnSetSecurityDescriptorGroup,_ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorGroup")
	fnSetSecurityDescriptorOwner,_ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorOwner")
	fnSetSecurityDescriptorSacl,_ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorSacl")
	fnSetThreadToken,_ = syscall.GetProcAddress(libadvapi32, "SetThreadToken")
	fnSetTokenInformation,_ = syscall.GetProcAddress(libadvapi32, "SetTokenInformation")



	fnAbortSystemShutdownW, _ = syscall.GetProcAddress(libadvapi32, "AbortSystemShutdownW")
	fnAbortSystemShutdownA, _ = syscall.GetProcAddress(libadvapi32, "AbortSystemShutdownA")
	fnAddAuditAccessAce, _ = syscall.GetProcAddress(libadvapi32, "AddAuditAccessAce")
	fnAddAuditAccessAceEx, _ = syscall.GetProcAddress(libadvapi32, "AddAuditAccessAceEx")
	fnAddAuditAccessObjectAce, _ = syscall.GetProcAddress(libadvapi32, "AddAuditAccessObjectAce")
	fnAdjustTokenGroups, _ = syscall.GetProcAddress(libadvapi32, "AdjustTokenGroups")
	fnAllocateLocallyUniqueId, _ = syscall.GetProcAddress(libadvapi32, "AllocateLocallyUniqueId")
	fnAreAllAccessesGranted, _ = syscall.GetProcAddress(libadvapi32, "AreAllAccessesGranted")
	fnAreAnyAccessesGranted, _ = syscall.GetProcAddress(libadvapi32, "AreAnyAccessesGranted")
	fnBackupEventLogW, _ = syscall.GetProcAddress(libadvapi32, "BackupEventLogW")
	fnBackupEventLogA, _ = syscall.GetProcAddress(libadvapi32, "BackupEventLogA")
	fnClearEventLogW, _ = syscall.GetProcAddress(libadvapi32, "ClearEventLogW")
	fnClearEventLogA, _ = syscall.GetProcAddress(libadvapi32, "ClearEventLogA")
	fnCloseEventLog, _ = syscall.GetProcAddress(libadvapi32, "CloseEventLog")
	fnConvertToAutoInheritPrivateObjectSecurity, _ = syscall.GetProcAddress(libadvapi32, "ConvertToAutoInheritPrivateObjectSecurity")
	fnCopySid, _ = syscall.GetProcAddress(libadvapi32, "CopySid")
	fnConvertSidToStringSidW, _ = syscall.GetProcAddress(libadvapi32, "ConvertSidToStringSidW")
	fnConvertSidToStringSidA, _ = syscall.GetProcAddress(libadvapi32, "ConvertSidToStringSidA")
	fnConvertStringSidToSidW, _ = syscall.GetProcAddress(libadvapi32, "ConvertStringSidToSidW")
	fnConvertStringSidToSidA, _ = syscall.GetProcAddress(libadvapi32, "ConvertStringSidToSidA")
	fnCreatePrivateObjectSecurity, _ = syscall.GetProcAddress(libadvapi32, "CreatePrivateObjectSecurity")
	fnCreatePrivateObjectSecurityEx, _ = syscall.GetProcAddress(libadvapi32, "CreatePrivateObjectSecurityEx")
	fnCreateProcessAsUserW, _ = syscall.GetProcAddress(libadvapi32, "CreateProcessAsUserW")
	fnCreateProcessAsUserA, _ = syscall.GetProcAddress(libadvapi32, "CreateProcessAsUserA")
	fnDecryptFileW, _ = syscall.GetProcAddress(libadvapi32, "DecryptFileW")
	fnDecryptFileA, _ = syscall.GetProcAddress(libadvapi32, "DecryptFileA")
	fnDeleteAce, _ = syscall.GetProcAddress(libadvapi32, "DeleteAce")
	fnDeregisterEventSource, _ = syscall.GetProcAddress(libadvapi32, "DeregisterEventSource")
	fnDestroyPrivateObjectSecurity, _ = syscall.GetProcAddress(libadvapi32, "DestroyPrivateObjectSecurity")
	fnDuplicateToken, _ = syscall.GetProcAddress(libadvapi32, "DuplicateToken")
	fnDuplicateTokenEx, _ = syscall.GetProcAddress(libadvapi32, "DuplicateTokenEx")
	fnEncryptFileW, _ = syscall.GetProcAddress(libadvapi32, "EncryptFileW")
	fnEncryptFileA, _ = syscall.GetProcAddress(libadvapi32, "EncryptFileA")
	fnEqualPrefixSid, _ = syscall.GetProcAddress(libadvapi32, "EqualPrefixSid")
	fnEqualSid, _ = syscall.GetProcAddress(libadvapi32, "EqualSid")
	fnFindFirstFreeAce, _ = syscall.GetProcAddress(libadvapi32, "FindFirstFreeAce")
	fnFreeSid, _ = syscall.GetProcAddress(libadvapi32, "FreeSid")
	fnGetAce, _ = syscall.GetProcAddress(libadvapi32, "GetAce")
	fnGetCurrentHwProfileW, _ = syscall.GetProcAddress(libadvapi32, "GetCurrentHwProfileW")
	fnGetCurrentHwProfileA, _ = syscall.GetProcAddress(libadvapi32, "GetCurrentHwProfileA")
	fnGetFileSecurityW, _ = syscall.GetProcAddress(libadvapi32, "GetFileSecurityW")
	fnGetFileSecurityA, _ = syscall.GetProcAddress(libadvapi32, "GetFileSecurityA")
	fnGetKernelObjectSecurity, _ = syscall.GetProcAddress(libadvapi32, "GetKernelObjectSecurity")
	fnGetLengthSid, _ = syscall.GetProcAddress(libadvapi32, "GetLengthSid")
	fnGetNumberOfEventLogRecords, _ = syscall.GetProcAddress(libadvapi32, "GetNumberOfEventLogRecords")
	fnGetOldestEventLogRecord, _ = syscall.GetProcAddress(libadvapi32, "GetOldestEventLogRecord")
	fnGetPrivateObjectSecurity, _ = syscall.GetProcAddress(libadvapi32, "GetPrivateObjectSecurity")
	fnGetSecurityDescriptorControl, _ = syscall.GetProcAddress(libadvapi32, "GetSecurityDescriptorControl")
	fnGetSecurityDescriptorDacl, _ = syscall.GetProcAddress(libadvapi32, "GetSecurityDescriptorDacl")
	fnGetSecurityDescriptorGroup, _ = syscall.GetProcAddress(libadvapi32, "GetSecurityDescriptorGroup")
	fnGetSecurityDescriptorLength, _ = syscall.GetProcAddress(libadvapi32, "GetSecurityDescriptorLength")
	fnGetSecurityDescriptorOwner, _ = syscall.GetProcAddress(libadvapi32, "GetSecurityDescriptorOwner")
	fnGetSecurityDescriptorSacl, _ = syscall.GetProcAddress(libadvapi32, "GetSecurityDescriptorSacl")
	fnGetSidIdentifierAuthority, _ = syscall.GetProcAddress(libadvapi32, "GetSidIdentifierAuthority")
	fnGetSidLengthRequired, _ = syscall.GetProcAddress(libadvapi32, "GetSidLengthRequired")
	fnGetSidSubAuthority, _ = syscall.GetProcAddress(libadvapi32, "GetSidSubAuthority")
	fnGetSidSubAuthorityCount, _ = syscall.GetProcAddress(libadvapi32, "GetSidSubAuthorityCount")
	fnGetTokenInformation, _ = syscall.GetProcAddress(libadvapi32, "GetTokenInformation")
	fnGetUserNameW, _ = syscall.GetProcAddress(libadvapi32, "GetUserNameW")
	fnGetUserNameA, _ = syscall.GetProcAddress(libadvapi32, "GetUserNameA")
	fnImpersonateLoggedOnUser, _ = syscall.GetProcAddress(libadvapi32, "ImpersonateLoggedOnUser")
	fnImpersonateNamedPipeClient, _ = syscall.GetProcAddress(libadvapi32, "ImpersonateNamedPipeClient")
	fnImpersonateSelf, _ = syscall.GetProcAddress(libadvapi32, "ImpersonateSelf")
	fnInitializeAcl, _ = syscall.GetProcAddress(libadvapi32, "InitializeAcl")
	fnInitializeSecurityDescriptor, _ = syscall.GetProcAddress(libadvapi32, "InitializeSecurityDescriptor")
	fnInitiateSystemShutdownW, _ = syscall.GetProcAddress(libadvapi32, "InitiateSystemShutdownW")
	fnInitiateSystemShutdownA, _ = syscall.GetProcAddress(libadvapi32, "InitiateSystemShutdownA")
	fnIsTextUnicode, _ = syscall.GetProcAddress(libadvapi32, "IsTextUnicode")
	fnIsValidAcl, _ = syscall.GetProcAddress(libadvapi32, "IsValidAcl")
	fnIsValidSecurityDescriptor, _ = syscall.GetProcAddress(libadvapi32, "IsValidSecurityDescriptor")
	fnIsValidSid, _ = syscall.GetProcAddress(libadvapi32, "IsValidSid")
	fnLogonUserW, _ = syscall.GetProcAddress(libadvapi32, "LogonUserW")
	fnLogonUserA, _ = syscall.GetProcAddress(libadvapi32, "LogonUserA")
	fnLookupAccountNameW, _ = syscall.GetProcAddress(libadvapi32, "LookupAccountNameW")
	fnLookupAccountNameA, _ = syscall.GetProcAddress(libadvapi32, "LookupAccountNameA")
	fnLookupAccountSidW, _ = syscall.GetProcAddress(libadvapi32, "LookupAccountSidW")
	fnLookupAccountSidA, _ = syscall.GetProcAddress(libadvapi32, "LookupAccountSidA")
	fnLookupPrivilegeDisplayNameW, _ = syscall.GetProcAddress(libadvapi32, "LookupPrivilegeDisplayNameW")
	fnLookupPrivilegeDisplayNameA, _ = syscall.GetProcAddress(libadvapi32, "LookupPrivilegeDisplayNameA")
	fnLookupPrivilegeNameW, _ = syscall.GetProcAddress(libadvapi32, "LookupPrivilegeNameW")
	fnLookupPrivilegeNameA, _ = syscall.GetProcAddress(libadvapi32, "LookupPrivilegeNameA")
	fnLookupPrivilegeValueW, _ = syscall.GetProcAddress(libadvapi32, "LookupPrivilegeValueW")
	fnLookupPrivilegeValueA, _ = syscall.GetProcAddress(libadvapi32, "LookupPrivilegeValueA")
	fnMakeAbsoluteSD, _ = syscall.GetProcAddress(libadvapi32, "MakeAbsoluteSD")
	fnMakeSelfRelativeSD, _ = syscall.GetProcAddress(libadvapi32, "MakeSelfRelativeSD")
	fnNotifyChangeEventLog, _ = syscall.GetProcAddress(libadvapi32, "NotifyChangeEventLog")
	fnObjectCloseAuditAlarmW, _ = syscall.GetProcAddress(libadvapi32, "ObjectCloseAuditAlarmW")
	fnObjectCloseAuditAlarmA, _ = syscall.GetProcAddress(libadvapi32, "ObjectCloseAuditAlarmA")
	fnObjectDeleteAuditAlarmW, _ = syscall.GetProcAddress(libadvapi32, "ObjectDeleteAuditAlarmW")
	fnObjectDeleteAuditAlarmA, _ = syscall.GetProcAddress(libadvapi32, "ObjectDeleteAuditAlarmA")
	fnObjectOpenAuditAlarmW, _ = syscall.GetProcAddress(libadvapi32, "ObjectOpenAuditAlarmW")
	fnObjectOpenAuditAlarmA, _ = syscall.GetProcAddress(libadvapi32, "ObjectOpenAuditAlarmA")
	fnObjectPrivilegeAuditAlarmW, _ = syscall.GetProcAddress(libadvapi32, "ObjectPrivilegeAuditAlarmW")
	fnObjectPrivilegeAuditAlarmA, _ = syscall.GetProcAddress(libadvapi32, "ObjectPrivilegeAuditAlarmA")
	fnOpenBackupEventLogW, _ = syscall.GetProcAddress(libadvapi32, "OpenBackupEventLogW")
	fnOpenBackupEventLogA, _ = syscall.GetProcAddress(libadvapi32, "OpenBackupEventLogA")
	fnOpenEventLogW, _ = syscall.GetProcAddress(libadvapi32, "OpenEventLogW")
	fnOpenEventLogA, _ = syscall.GetProcAddress(libadvapi32, "OpenEventLogA")
	fnOpenProcessToken, _ = syscall.GetProcAddress(libadvapi32, "OpenProcessToken")
	fnOpenThreadToken, _ = syscall.GetProcAddress(libadvapi32, "OpenThreadToken")
	fnPrivilegeCheck, _ = syscall.GetProcAddress(libadvapi32, "PrivilegeCheck")
	fnPrivilegedServiceAuditAlarmW, _ = syscall.GetProcAddress(libadvapi32, "PrivilegedServiceAuditAlarmW")
	fnPrivilegedServiceAuditAlarmA, _ = syscall.GetProcAddress(libadvapi32, "PrivilegedServiceAuditAlarmA")
	fnReadEventLogW, _ = syscall.GetProcAddress(libadvapi32, "ReadEventLogW")
	fnReadEventLogA, _ = syscall.GetProcAddress(libadvapi32, "ReadEventLogA")
	fnRegOverridePredefKey, _ = syscall.GetProcAddress(libadvapi32, "RegOverridePredefKey")
	fnRegOpenUserClassesRoot, _ = syscall.GetProcAddress(libadvapi32, "RegOpenUserClassesRoot")
	fnRegOpenCurrentUser, _ = syscall.GetProcAddress(libadvapi32, "RegOpenCurrentUser")
	fnRegDisablePredefinedCache, _ = syscall.GetProcAddress(libadvapi32, "RegDisablePredefinedCache")
	fnRegQueryMultipleValuesW, _ = syscall.GetProcAddress(libadvapi32, "RegQueryMultipleValuesW")
	fnRegQueryMultipleValuesA, _ = syscall.GetProcAddress(libadvapi32, "RegQueryMultipleValuesA")
	fnRegisterEventSourceW, _ = syscall.GetProcAddress(libadvapi32, "RegisterEventSourceW")
	fnRegisterEventSourceA, _ = syscall.GetProcAddress(libadvapi32, "RegisterEventSourceA")
	fnReportEventW, _ = syscall.GetProcAddress(libadvapi32, "ReportEventW")
	fnReportEventA, _ = syscall.GetProcAddress(libadvapi32, "ReportEventA")
	fnSetFileSecurityW, _ = syscall.GetProcAddress(libadvapi32, "SetFileSecurityW")
	fnSetFileSecurityA, _ = syscall.GetProcAddress(libadvapi32, "SetFileSecurityA")
}

func RegCloseKey(hKey HKEY)int32  {
	ret,_,_ := syscall.Syscall(fnRegCloseKey,1,uintptr(hKey),0,0)
	return int32(ret)
}

func RegOpenKeyEx(hkey HKEY,lpSubKey *uint16,ulOptions uint32,samDesired REGSAM,phkResult *HKEY) int32 {
	ret,_,_ := syscall.Syscall6(fnRegOpenKeyEx,5,uintptr(hkey),uintptr(unsafe.Pointer(lpSubKey)),uintptr(ulOptions),
		uintptr(samDesired),uintptr(unsafe.Pointer(phkResult)),0)
	return int32(ret)
}


func RegFlushKey(hkey HKEY)int32  {
	ret,_,_ := syscall.Syscall(fnRegFlushKey,1,uintptr(hkey),0,0)
	return int32(ret)
}


func RegCreateKeyEx(hkey HKEY,lpSubKey *uint16,Reserved uint32,lpClass *uint16,dwOptions uint32,samDesired REGSAM,
	lpSecurityAttributes *GSecurityAttributes,phkResult *HKEY,lpdwDisposition *uint32)int32  {
	ret,_,_ := syscall.Syscall9(fnRegCreateKeyEx,9,uintptr(hkey),uintptr(unsafe.Pointer(lpSubKey)),uintptr(Reserved),
		uintptr(unsafe.Pointer(lpClass)),uintptr(dwOptions),uintptr(samDesired),uintptr(unsafe.Pointer(lpSecurityAttributes)),
		uintptr(unsafe.Pointer(phkResult)),uintptr(unsafe.Pointer(lpdwDisposition)))
	return int32(ret)
}

func RegQueryInfoKey(hkey HKEY,lpClass *uint16,lpcbClass *uint32,lpReserved uintptr,lpcSubKeys, lpcbMaxSubKeyLen, lpcbMaxClassLen, lpcValues *uint32,
	lpcbMaxValueNameLen, lpcbMaxValueLen, lpcbSecurityDescriptor *uint32,lpftLastWriteTime *GFileTime) int32 {
	ret,_,_ := syscall.Syscall12(fnRegQueryInfoKey,12,uintptr(hkey),uintptr(unsafe.Pointer(lpClass)),
		uintptr(unsafe.Pointer(lpcbClass)),uintptr(lpReserved),
		uintptr(unsafe.Pointer(lpcSubKeys)),uintptr(unsafe.Pointer(lpcbMaxSubKeyLen)),
		uintptr(unsafe.Pointer(lpcbMaxClassLen)),uintptr(unsafe.Pointer(lpcValues)),
		uintptr(unsafe.Pointer(lpcbMaxValueNameLen)),uintptr(unsafe.Pointer(lpcbMaxValueLen)),
		uintptr(unsafe.Pointer(lpcbSecurityDescriptor)), uintptr(unsafe.Pointer(lpftLastWriteTime)))
	return int32(ret)
}


func RegEnumKeyEx(hKey HKEY,dwIndex uint,lpName *uint16,lpcbName *uint32,lpReserved uintptr,lpClass *uint16,
	lpcbClass *uint32,lpftLastWriteTime *GFileTime)int32  {
	ret,_,_ := syscall.Syscall9(fnRegEnumKeyEx,8,uintptr(hKey),uintptr(dwIndex),uintptr(unsafe.Pointer(lpName)),
		uintptr(unsafe.Pointer(lpcbName)),lpReserved,uintptr(unsafe.Pointer(lpClass)),uintptr(unsafe.Pointer(lpcbClass)),
		uintptr(unsafe.Pointer(lpftLastWriteTime)),0)
	return int32(ret)
}

func RegEnumValue(hKey HKEY,dwIndex uint32,lpValueName *uint16,lpcbValueName *uint32,lpReserved uintptr,
	lpType *uint32,lpData *byte,lpcbData *uint32)int32  {
	ret,_,_ := syscall.Syscall9(fnRegEnumValue,8,uintptr(hKey),uintptr(dwIndex),uintptr(unsafe.Pointer(lpValueName)),
		uintptr(unsafe.Pointer(lpcbValueName)),lpReserved,uintptr(unsafe.Pointer(lpType)),uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(lpcbData)),0)
	return int32(ret)
}

func RegLoadKey(hKey HKEY, lpSubKey, lpFile *uint16)int32  {
	ret,_,_ := syscall.Syscall(fnRegLoadKey,3,uintptr(hKey),uintptr(unsafe.Pointer(lpSubKey)),uintptr(unsafe.Pointer(lpFile)))
	return int32(ret)
}

func RegUnLoadKey(hKey HKEY, lpSubKey *uint16)int32  {
	ret,_,_ := syscall.Syscall(fnRegUnLoadKey,2,uintptr(hKey),uintptr(unsafe.Pointer(lpSubKey)),0)
	return int32(ret)
}

func RegQueryValueEx(hKey HKEY, lpValueName *uint16,
lpReserved uintptr, lpType *uint32, lpData *byte, lpcbData *uint32)int32  {
	ret,_,_ := syscall.Syscall6(fnRegQueryValueEx,6,uintptr(hKey),uintptr(unsafe.Pointer(lpValueName)),lpReserved,
		uintptr(unsafe.Pointer(lpType)),uintptr(unsafe.Pointer(lpData)),uintptr(unsafe.Pointer(lpcbData)))
	return int32(ret)
}

func RegQueryValue(hKey HKEY, lpSubKey,lpValue *uint16, lpcbValue *int32)int32  {
	ret,_,_ := syscall.Syscall6(fnRegQueryValue,4,uintptr(hKey),uintptr(unsafe.Pointer(lpSubKey)),
		uintptr(unsafe.Pointer(lpValue)),uintptr(unsafe.Pointer(lpcbValue)),0,0)
	return int32(ret)
}

func RegOpenKey(hKey HKEY, lpSubKey *uint16,  phkResult *HKEY)int32{
	ret,_,_ := syscall.Syscall(fnRegOpenKey,3,uintptr(hKey),uintptr(unsafe.Pointer(lpSubKey)),uintptr(unsafe.Pointer(phkResult)))
	return int32(ret)
}
func RegConnectRegistry(lpMachineName *uint16, hKey HKEY, phkResult *HKEY)int32{
	ret,_,_ := syscall.Syscall(fnRegConnectRegistry,3,uintptr(unsafe.Pointer(lpMachineName)),uintptr(hKey),uintptr(unsafe.Pointer(phkResult)))
	return int32(ret)
}

func RegDeleteKey(hKey HKEY, lpSubKey *uint16)int32{
	ret,_,_ := syscall.Syscall(fnRegDeleteKey,2,uintptr(hKey),uintptr(unsafe.Pointer(lpSubKey)),0)
	return int32(ret)
}

func RegDeleteKeyEx(hKey HKEY, lpSubKey *uint16, samDesired REGSAM, Reserved uint32)int32{
	ret,_,_ := syscall.Syscall6(fnRegDeleteKeyEx,4,uintptr(hKey),uintptr(unsafe.Pointer(lpSubKey)),
		uintptr(samDesired),uintptr(Reserved),0,0)
	return int32(ret)
}

func RegDeleteValue(hKey HKEY, lpValueName *uint16)int32{
	ret,_,_ := syscall.Syscall(fnRegDeleteValue,2,uintptr(hKey),uintptr(unsafe.Pointer(lpValueName)),0)
	return int32(ret)
}

func RegGetKeySecurity(hKey HKEY, SecurityInformation uint32,
pSecurityDescriptor *GSecurityDescriptor, lpcbSecurityDescriptor  *uint32) int32{
	ret,_,_ := syscall.Syscall6(fnRegGetKeySecurity,4,uintptr(hKey),uintptr(SecurityInformation),
		uintptr(unsafe.Pointer(pSecurityDescriptor)),uintptr(unsafe.Pointer(lpcbSecurityDescriptor)),0,0)
	return int32(ret)
}

func RegNotifyChangeKeyValue(hKey HKEY, bWatchSubtree bool, dwNotifyFilter uint32, hEvent syscall.Handle, fAsynchronus bool)int32{
	ret,_,_ := syscall.Syscall6(fnRegNotifyChangeKeyValue,5,uintptr(hKey),uintptr(BoolToUint(bWatchSubtree)),
		uintptr(dwNotifyFilter),uintptr(hEvent),uintptr(BoolToUint(fAsynchronus)),0)
	return int32(ret)
}

func RegReplaceKey(hKey HKEY, lpSubKey,lpNewFile,lpOldFile *uint16)int32{
	ret,_,_ := syscall.Syscall6(fnRegReplaceKey,4,uintptr(hKey),uintptr(unsafe.Pointer(lpSubKey)),
		uintptr(unsafe.Pointer(lpNewFile)),uintptr(unsafe.Pointer(lpOldFile)),0,0)
	return int32(ret)
}

func RegRestoreKey(hKey HKEY, lpFile *uint16, dwFlags uint32) int32{
	ret,_,_ := syscall.Syscall(fnRegRestoreKey,3,uintptr(hKey),uintptr(unsafe.Pointer(lpFile)),uintptr(dwFlags))
	return int32(ret)
}

func RegSaveKeyEx(hKey HKEY, lpFile *uint16, lpSecurityAttributes *GSecurityAttributes, Flags uint32)int32{
	ret,_,_ := syscall.Syscall6(fnRegSaveKeyEx,4,uintptr(hKey),uintptr(unsafe.Pointer(lpFile)),
		uintptr(unsafe.Pointer(lpSecurityAttributes)),uintptr(Flags),0,0)
	return int32(ret)
}

func RegSaveKey(hKey HKEY, lpFile *uint16, lpSecurityAttributes *GSecurityAttributes)int32{
	ret,_,_ := syscall.Syscall(fnRegSaveKey,3,uintptr(hKey),uintptr(unsafe.Pointer(lpFile)),uintptr(unsafe.Pointer(lpSecurityAttributes)))
	return int32(ret)
}

func RegSetValue(hKey HKEY, lpSubKey *uint16, dwType uint32, lpData *uint16, cbData uint32)int32{
	ret,_,_ := syscall.Syscall6(fnRegSetValue,5,uintptr(hKey),uintptr(unsafe.Pointer(lpSubKey)),uintptr(dwType),
		uintptr(unsafe.Pointer(lpData)),uintptr(cbData),0)
	return int32(ret)
}

func RegSetValueEx(hKey HKEY, lpValueName *uint16, Reserved,dwType uint32, lpData uintptr, cbData uint32)int32{
	ret,_,_ := syscall.Syscall6(fnRegSetValueEx,6,uintptr(hKey),uintptr(unsafe.Pointer(lpValueName)),uintptr(Reserved),
		uintptr(dwType),lpData,uintptr(cbData))
	return int32(ret)
}

func RegisterEventSource(lpUNCServerName, lpSourceName *uint16) syscall.Handle{
	ret,_,_ := syscall.Syscall(fnRegisterEventSource,2,uintptr(unsafe.Pointer(lpUNCServerName)),uintptr(unsafe.Pointer(lpSourceName)),0)
	return syscall.Handle(ret)
}