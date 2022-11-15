module github.com/suiyunonghen/GVCL

go 1.12

require (
	github.com/suiyunonghen/DxCommonLib v0.5.2
	golang.org/x/image v0.1.0
)

replace (
	//github.com/suiyunonghen/DxCommonLib => /../DxCommonLib
	golang.org/x/crypto => github.com/golang/crypto v0.2.0
	golang.org/x/image => github.com/golang/image v0.1.0
	golang.org/x/net => github.com/golang/net v0.2.0
	golang.org/x/sync => github.com/golang/sync v0.1.0
	golang.org/x/sys => github.com/golang/sys v0.2.0
	golang.org/x/text => github.com/golang/text v0.4.0
	golang.org/x/tools => github.com/golang/tools v0.3.0
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20220907171357-04be3eba64a2
)
