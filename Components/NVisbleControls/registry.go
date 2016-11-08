package NVisbleControls

import (
	"suiyunonghen/GVCL/WinApi"
	"suiyunonghen/GVCL/Graphics"
	"strings"
	"syscall"
	"unsafe"
)

type GRegKeyInfo struct {
	NumSubKeys uint32
	MaxSubKeyLen uint32
	NumValues uint32
	MaxValueLen uint32
	MaxDataLen uint32
	FileTime WinApi.GFileTime
}
type GRegistry struct {
	Graphics.GObject
	fCurrentKey WinApi.HKEY
	fRootKey  WinApi.HKEY
	fLazyWrite bool
	fCurrentPath string
	fCloseRootKey bool
	fAccess uint32
	fLastError int32
}

type GRegDataType byte
const (
	_ GRegDataType=iota
		RdUnknown=iota
		RdString
		RdExpandString
		RdInteger
		RdBinary
)

type GRegDataInfo struct {
	RegData GRegDataType
	DataSize uint32
}

func (registry *GRegistry) SubInit() {
	registry.GObject.SubInit(registry)
	registry.SetRootKey(WinApi.HKEY_CURRENT_USER)
	registry.fAccess = WinApi.KEY_ALL_ACCESS
	registry.fLazyWrite = true
}
func (registry *GRegistry)Destroy()  {
	registry.CloseKey()
}

func (reg *GRegistry)SetRootKey(rkey WinApi.HKEY)  {
	if reg.fRootKey != rkey{
		if reg.fCloseRootKey{
			WinApi.RegCloseKey(reg.fRootKey);
			reg.fCloseRootKey = false
		}
		reg.fRootKey = rkey
		reg.CloseKey()
	}
}

func (reg *GRegistry)CloseKey()  {
	if reg.fCurrentKey != 0{
		if !reg.fLazyWrite{
			WinApi.RegFlushKey(reg.fCurrentKey)
		}
		WinApi.RegCloseKey(reg.fCurrentKey)
		reg.fCurrentKey = 0
		reg.fCurrentPath = ""
	}
}

func (reg *GRegistry)CurrentKey()WinApi.HKEY  {
	return reg.fCurrentKey
}

func (reg *GRegistry)CheckResult(RetVal int32)bool  {
	reg.fLastError = RetVal
	return RetVal == WinApi.ERROR_SUCCESS
}

func (reg *GRegistry)getBaseKey(Relative bool)WinApi.HKEY  {
	if reg.fCurrentKey == 0 || !Relative{
		return reg.fRootKey
	}else{
		return reg.fCurrentKey
	}
}

func (reg *GRegistry)ChangeKey(Value WinApi.HKEY,Path string)  {
	reg.CloseKey()
	reg.fCurrentKey = Value
	reg.fCurrentPath = Path
}

func (reg *GRegistry)OpenKey(keypath string,Cancreate bool)(result bool)  {
	Relative := !(keypath != "" && keypath[0]==92)
	if !Relative{
		keypath = strings.Trim(keypath,"\\")
	}
	var TempKey WinApi.HKEY = 0
	if !Cancreate || keypath == ""{
		result = reg.CheckResult(WinApi.RegOpenKeyEx(reg.getBaseKey(Relative), WinApi.PChar(keypath), 0,
			WinApi.REGSAM(reg.fAccess), &TempKey))
	}else{
		var lpdwDisposition uint32=0
		result = reg.CheckResult(WinApi.RegCreateKeyEx(reg.getBaseKey(Relative), WinApi.PChar(keypath), 0, nil,
			WinApi.REG_OPTION_NON_VOLATILE, WinApi.REGSAM(reg.fAccess), nil, &TempKey, &lpdwDisposition))
	}
	if result{
		if reg.fCurrentKey != 0 && Relative{
			keypath = reg.fCurrentPath + "\\" + keypath;
		}
		reg.ChangeKey(TempKey, keypath)
	}
	return
}

func (reg *GRegistry)CreateKey(Key string)bool  {
	Relative := !(Key != "" && Key[0]==92)
	if !Relative{
		Key = strings.Trim(Key,"\\")
	}
	var TempKey WinApi.HKEY = 0
	WOWFlags := reg.fAccess & WinApi.KEY_WOW64_RES
	var lpdwDisposition uint32=0
	result := reg.CheckResult(WinApi.RegCreateKeyEx(reg.getBaseKey(Relative), WinApi.PChar(Key), 0, nil,
		WinApi.REG_OPTION_NON_VOLATILE, WinApi.REGSAM(WinApi.KEY_ALL_ACCESS | WOWFlags), nil, &TempKey, &lpdwDisposition))
	if result{
		WinApi.RegCloseKey(TempKey)
	}
	return result
}

func (reg *GRegistry)getKey(key string)(Result WinApi.HKEY)  {
	Relative := !(key != "" && key[0]==92)
	if !Relative{
		key = strings.Trim(key,"\\")
	}
	Result = 0
	WinApi.RegOpenKeyEx(reg.getBaseKey(Relative), WinApi.PChar(key), 0, WinApi.REGSAM(reg.fAccess), &Result)
	return Result
}

func (reg *GRegistry)getKeyInfo()(keyinfo *GRegKeyInfo,ok bool)  {
	keyinfo = new(GRegKeyInfo)
	ok = reg.CheckResult(WinApi.RegQueryInfoKey(reg.fCurrentKey, nil, nil, 0, &keyinfo.NumSubKeys,
		&keyinfo.MaxSubKeyLen, nil, &keyinfo.NumValues, &keyinfo.MaxValueLen,
		&keyinfo.MaxDataLen, nil, &keyinfo.FileTime))
	if ok{
		keyinfo.MaxSubKeyLen += keyinfo.MaxSubKeyLen
		keyinfo.MaxValueLen += keyinfo.MaxValueLen
	}
	return
}

func (reg *GRegistry)DeleteKey(Key string)bool  {
	Relative := !(Key != "" && Key[0]==92)
	if !Relative{
		Key = strings.Trim(Key,"\\")
	}
	OldKey := reg.fCurrentKey
	fDeleteKey := reg.getKey(Key)
	if fDeleteKey != 0{
		reg.fCurrentKey = fDeleteKey
		if keyinfo,ok := reg.getKeyInfo();ok{
			KeyName := make([]uint16,keyinfo.MaxSubKeyLen+1)
			for i := keyinfo.NumSubKeys - 1;i>=0;i--{
				Len := keyinfo.MaxSubKeyLen + 1
				if reg.CheckResult(WinApi.RegEnumKeyEx(fDeleteKey, uint(i), &KeyName[0], &Len, 0, nil, nil, nil)){
					reg.DeleteKey(syscall.UTF16ToString(KeyName))
				}
			}
		}
		reg.fCurrentKey = OldKey
		WinApi.RegCloseKey(fDeleteKey)
	}
	if WinApi.RegDeleteKeyEx != nil{
		return reg.CheckResult(WinApi.RegDeleteKeyEx(reg.getBaseKey(Relative), WinApi.PChar(Key), WinApi.REGSAM(reg.fAccess & WinApi.KEY_WOW64_RES), 0))
	}
	return reg.CheckResult(WinApi.RegDeleteKey(reg.getBaseKey(Relative), WinApi.PChar(Key)))
}

func (reg *GRegistry)DeleteValue(Name string)bool  {
	return reg.CheckResult(WinApi.RegDeleteValue(reg.fCurrentKey, WinApi.PChar(Name)));
}

func (reg *GRegistry)KeyExists(Key string)bool  {
	OldAccess := reg.fAccess
	reg.fAccess = WinApi.STANDARD_RIGHTS_READ | WinApi.KEY_QUERY_VALUE |
		WinApi.KEY_ENUMERATE_SUB_KEYS | (OldAccess & WinApi.KEY_WOW64_RES)
	TempKey := reg.getKey(Key)
	if TempKey != 0 {
		WinApi.RegCloseKey(TempKey)
		return true
	}
	reg.fAccess = OldAccess
	return false
}

func dataTypeToRegDatatype(DataType uint32)GRegDataType  {
	switch DataType {
	case WinApi.REG_SZ: return RdString
	case WinApi.REG_EXPAND_SZ: return RdExpandString
	case WinApi.REG_DWORD: return RdInteger
	case WinApi.REG_BINARY: return RdBinary
	default:
		return RdUnknown
	}
}

func (reg *GRegistry)GetDataInfo(ValueName string)(result *GRegDataInfo,ok bool)  {
	var DataType,datasize uint32
	ok = reg.CheckResult(WinApi.RegQueryValueEx(reg.fCurrentKey, WinApi.PChar(ValueName), 0, &DataType, nil,
	  &datasize));
	if ok{
		result = new(GRegDataInfo)
		result.DataSize = datasize
		result.RegData = dataTypeToRegDatatype(DataType)
	}else{
		result = nil
	}
	return
}

func (reg *GRegistry)ValueExists(name string)bool  {
	_,ok := reg.GetDataInfo(name)
	return ok
}

func (reg *GRegistry)OpenKeyReadOnly(Key string)bool  {
	Relative := !(Key != "" && Key[0]==92)
	if !Relative{
		Key = strings.Trim(Key,"\\")
	}
	var TempKey WinApi.HKEY = 0
	WOWFlags := reg.fAccess & WinApi.KEY_WOW64_RES
	Result := reg.CheckResult(WinApi.RegOpenKeyEx(reg.getBaseKey(Relative), WinApi.PChar(Key), 0,
		WinApi.REGSAM(WinApi.KEY_READ | WOWFlags), &TempKey))
	if Result{
		reg.fAccess = WinApi.KEY_READ | WOWFlags
		if reg.fCurrentKey != 0 && Relative {
			Key = reg.fCurrentPath + "\\" + Key
		}
		reg.ChangeKey(TempKey, Key)
	}else {
		Result = reg.CheckResult(WinApi.RegOpenKeyEx(reg.getBaseKey(Relative), WinApi.PChar(Key), 0,
		 	WinApi.REGSAM(WinApi.STANDARD_RIGHTS_READ | WinApi.KEY_QUERY_VALUE | WinApi.KEY_ENUMERATE_SUB_KEYS | WOWFlags),
			&TempKey))
		if Result{
			reg.fAccess = WinApi.STANDARD_RIGHTS_READ | WinApi.KEY_QUERY_VALUE | WinApi.KEY_ENUMERATE_SUB_KEYS | WOWFlags
			if reg.fCurrentKey != 0  && Relative {
				Key = reg.fCurrentPath + "\\" + Key
			}
			reg.ChangeKey(TempKey, Key)
		}else{
			Result = reg.CheckResult(WinApi.RegOpenKeyEx(reg.getBaseKey(Relative), WinApi.PChar(Key), 0,
				WinApi.REGSAM(WinApi.KEY_QUERY_VALUE | WOWFlags), &TempKey))
			if Result {
				reg.fAccess = WinApi.KEY_QUERY_VALUE | WOWFlags
				if reg.fCurrentKey != 0  && Relative {
					Key = reg.fCurrentPath + "\\" + Key
				}
				reg.ChangeKey(TempKey, Key)
			}
		}
	}
	return Result
}

func (reg *GRegistry)getData(Name string,Buffer  uintptr,BufSize uint32)(GRegDataType,int)  {
	var Datatype uint32=WinApi.REG_NONE
	if reg.CheckResult(WinApi.RegQueryValueEx(reg.fCurrentKey, WinApi.PChar(Name), 0, &Datatype, (*byte)(unsafe.Pointer(Buffer)), &BufSize)){
		return dataTypeToRegDatatype(Datatype),int(BufSize)
	}
	return RdUnknown,-1
}

func (reg *GRegistry)ReadInteger(Name string)(Result int)  {
	if regdata,ret :=reg.getData(Name, uintptr(unsafe.Pointer(&Result)), uint32(unsafe.Sizeof(Result)));ret>0 && regdata==RdInteger{
		return
	}else{
		return 0
	}
}

func (reg *GRegistry)ReadBool(name string)bool  {
	return reg.ReadInteger(name) != 0
}

func (reg *GRegistry)ReadString(name string)string  {
	if regdata,ok := reg.GetDataInfo(name);ok && regdata.DataSize > 0{
		result := make([]uint16,regdata.DataSize / 2)
		if regdata,ret := reg.getData(name,uintptr(unsafe.Pointer(&result[0])),regdata.DataSize);
			ret>0 && regdata == RdString || regdata == RdExpandString{
			return syscall.UTF16ToString(result)
		}
	}
	return  ""
}

func (reg *GRegistry)ReadFloat(name string)(Result float64)  {
	var RLen uint32=uint32(unsafe.Sizeof(Result))
	if regdata,ret :=reg.getData(name, uintptr(unsafe.Pointer(&Result)), RLen);uint32(ret)!=RLen || regdata!=RdBinary{
		Result = 0
	}
	return
}

func regDataToDataType(Value GRegDataType) uint32  {
	switch Value {
	case RdString: return WinApi.REG_SZ
	case RdExpandString: return WinApi.REG_EXPAND_SZ
	case RdBinary: return WinApi.REG_BINARY
	case RdInteger: return WinApi.REG_DWORD
	default:
		return WinApi.REG_NONE
	}
}

func (reg *GRegistry)putData(Name string,Buffer  uintptr,BufSize uint32,RegData GRegDataType)bool  {
	datatype := regDataToDataType(RegData)
	return reg.CheckResult(WinApi.RegSetValueEx(reg.fCurrentKey, WinApi.PChar(Name), 0, datatype, Buffer, BufSize))
}

func (reg *GRegistry)WriteInteger(Name string,value int)bool  {
	return reg.putData(Name, uintptr(unsafe.Pointer(&value)), uint32(unsafe.Sizeof(value)), RdInteger)
}

func (reg *GRegistry)WriteBool(Name string,v bool)bool  {
	return reg.WriteInteger(Name,int(WinApi.BoolToUint(v)))
}

func (reg *GRegistry)WriteFloat(Name string,v float64)bool  {
	return reg.putData(Name, uintptr(unsafe.Pointer(&v)), uint32(unsafe.Sizeof(v)), RdBinary)
}

func (reg *GRegistry)WriteString(Name string,v string)bool  {
	vbyte := syscall.StringToUTF16(v)
	return reg.putData(Name, uintptr(unsafe.Pointer(&vbyte[0])), uint32(2*(len(vbyte)+1)), RdString)
}

func (reg *GRegistry)WriteExpandString(Name string,v string)bool  {
	vbyte := syscall.StringToUTF16(v)
	return reg.putData(Name, uintptr(unsafe.Pointer(&vbyte[0])), uint32(2*(len(vbyte)+1)), RdExpandString)
}

func (reg *GRegistry)CurrentPath()string  {
	return reg.fCurrentPath
}

func (reg *GRegistry)RootKey()WinApi.HKEY  {
	return reg.fRootKey
}

func (reg *GRegistry)LoadKey(Key,FileName string)bool  {
	Relative := !(Key != "" && Key[0]==92)
	if !Relative{
		Key = strings.Trim(Key,"\\")
	}
	return reg.CheckResult(WinApi.RegLoadKey(reg.fRootKey, WinApi.PChar(Key), WinApi.PChar(FileName)))
}

func (reg *GRegistry)UnLoadKey(key string)bool  {
	Relative := !(key != "" && key[0]==92)
	if !Relative{
		key = strings.Trim(key,"\\")
	}
	return reg.CheckResult(WinApi.RegUnLoadKey(reg.fRootKey, WinApi.PChar(key)))
}

func (reg *GRegistry)WriteBinaryData(Name string,buffer uintptr,BufSize uint32)bool  {
	return reg.putData(Name, buffer, BufSize, RdBinary)
}

func (reg *GRegistry)ReadBinaryData(Name string,buffer uintptr,BufSize uint32)(result int)  {
	result = -1
	if regdata,ok:=reg.GetDataInfo(Name);ok{
		result = int(regdata.DataSize)
		if (regdata.RegData == RdBinary || regdata.RegData == RdUnknown) && uint32(result) <= BufSize{
			reg.getData(Name, buffer, uint32(result))
		}
	}
	return
}
func NewRegistry(access uint32)*GRegistry  {
	reg := new(GRegistry)
	reg.SubInit()
	if access != 0{
		reg.fAccess = access
	}
	return reg
}