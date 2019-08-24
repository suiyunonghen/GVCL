module github.com/suiyunonghen/GVCL

go 1.12

require (
	github.com/suiyunonghen/DxCommonLib v0.1.4
	golang.org/x/image v0.0.1
)

replace (
	//github.com/suiyunonghen/DxCommonLib => /../DxCommonLib
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/image => github.com/golang/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/net => github.com/golang/net v0.0.0-20190813141303-74dc4d7220e7
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190813064441-fde4db37ae7a
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190822000311-fc82fb2afd64
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
)
