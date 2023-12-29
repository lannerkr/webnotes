package main

import "github.com/zserge/lorca"

func genAsset() {
	err := lorca.Embed("main", "assets.go", "www")
	if err != nil {
		panic(err)
	}
}
