// +build generate

package main

import (
	"log"
	"net/http"
	"os"
	pathpkg "path"

	"github.com/shurcooL/go/vfs/httpfs/filter"
	"github.com/shurcooL/vfsgen"
)

func main() {
	ignore := func(fi os.FileInfo, path string) bool {
		switch {
		case len(fi.Name()) >= 2 && fi.Name()[0] == '.': // Skip .files/.folders (but not ".").
			return true
		case pathpkg.Ext(fi.Name()) == ".go":
			return true
		case path == "/README.md":
			return true
		default:
			return false
		}
	}
	var data http.FileSystem = filter.NewIgnore(http.Dir("."), ignore)

	err := vfsgen.Generate(data, vfsgen.Options{
		Filename:     "data.go",
		PackageName:  "data",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
