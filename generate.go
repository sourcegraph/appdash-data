//go:generate go-bindata -pkg=data -ignore=.*\.go -o data.go . ./benkeen/d3pie/ ./jiahuang/d3-timeline/ ./krisk/fuse/ ./zeroclipboard/zeroclipboard/
//go:generate go fmt .

package data
