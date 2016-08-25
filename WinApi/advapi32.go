package WinApi

import (
	"syscall"
)

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
	fnRegNotifyChangeKeyValue                   uintptr
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
)

func init() {
	libadvapi32, _ = syscall.LoadLibrary("advapi32.dll")
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
	fnRegNotifyChangeKeyValue, _ = syscall.GetProcAddress(libadvapi32, "RegNotifyChangeKeyValue")
	fnRegQueryMultipleValuesW, _ = syscall.GetProcAddress(libadvapi32, "RegQueryMultipleValuesW")
	fnRegQueryMultipleValuesA, _ = syscall.GetProcAddress(libadvapi32, "RegQueryMultipleValuesA")
	fnRegisterEventSourceW, _ = syscall.GetProcAddress(libadvapi32, "RegisterEventSourceW")
	fnRegisterEventSourceA, _ = syscall.GetProcAddress(libadvapi32, "RegisterEventSourceA")
	fnReportEventW, _ = syscall.GetProcAddress(libadvapi32, "ReportEventW")
	fnReportEventA, _ = syscall.GetProcAddress(libadvapi32, "ReportEventA")
	fnRevertToSelf, _ = syscall.GetProcAddress(libadvapi32, "RevertToSelf")
	fnSetAclInformation, _ = syscall.GetProcAddress(libadvapi32, "SetAclInformation")
	fnSetFileSecurityW, _ = syscall.GetProcAddress(libadvapi32, "SetFileSecurityW")
	fnSetFileSecurityA, _ = syscall.GetProcAddress(libadvapi32, "SetFileSecurityA")
	fnSetKernelObjectSecurity, _ = syscall.GetProcAddress(libadvapi32, "SetKernelObjectSecurity")
	fnSetPrivateObjectSecurity, _ = syscall.GetProcAddress(libadvapi32, "SetPrivateObjectSecurity")
	fnSetPrivateObjectSecurityEx, _ = syscall.GetProcAddress(libadvapi32, "SetPrivateObjectSecurityEx")
	fnSetSecurityDescriptorControl, _ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorControl")
	fnSetSecurityDescriptorDacl, _ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorDacl")
	fnSetSecurityDescriptorGroup, _ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorGroup")
	fnSetSecurityDescriptorOwner, _ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorOwner")
	fnSetSecurityDescriptorSacl, _ = syscall.GetProcAddress(libadvapi32, "SetSecurityDescriptorSacl")
	fnSetThreadToken, _ = syscall.GetProcAddress(libadvapi32, "SetThreadToken")
	fnSetTokenInformation, _ = syscall.GetProcAddress(libadvapi32, "SetTokenInformation")
}
