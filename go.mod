module github.com/suiyunonghen/GVCL

go 1.12

require (
	github.com/suiyunonghen/DxCommonLib v0.2.7
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8
)

replace (
	//github.com/suiyunonghen/DxCommonLib => /../DxCommonLib
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191128160524-b544559bb6d1
	golang.org/x/image => github.com/golang/image v0.0.0-20191009234506-e7c1f5e7dbb8
	golang.org/x/net => github.com/golang/net v0.0.0-20191126235420-ef20fe5d7933
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20191128015809-6d18c012aee9
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20191130070609-6e064ea0cf2d
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191011141410-1b5146add898
)
