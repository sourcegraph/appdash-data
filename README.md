# Appdash-data

This repository holds binary and external asset files (images, CSS, JS etc) that [Appdash](https://sourcegraph.com/sourcegraph/appdash) relies on.

The data is embedded into the final Appdash binary using [go-bindata](https://github.com/jteeuwen/go-bindata) and served via HTTP using [go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs).

