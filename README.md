# Appdash-data

This repository holds binary and external asset files (images, CSS, JS etc) that [Appdash](https://sourcegraph.com/sourcegraph/appdash) relies on.

The data is embedded into the final Appdash binary using [go-bindata](https://github.com/jteeuwen/go-bindata) and served via HTTP using [go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs).

# Regenerating

Simply run `go generate` in this directory to regenerate the `data.go` file using [go-bindata](https://github.com/jteeuwen/go-bindata).

# Licenses

All assets, source code files, etc in the top-level directory are under the [Appdash License](https://github.com/sourcegraph/appdash/blob/master/LICENSE).

The following directories are under their own licenses:

- `benkeen/d3pie` - [d3pie license](https://github.com/benkeen/d3pie/blob/master/LICENSE).
- `jiahuang/d3-timeline` - [d3-timeline license](https://github.com/jiahuang/d3-timeline#license).
- `krisk/fuse` - [Fuse license](https://github.com/krisk/Fuse/blob/master/LICENSE).
- `zeroclipboard/zeroclipboard` - [Zeroclipboard license](https://github.com/zeroclipboard/zeroclipboard/blob/master/LICENSE).
