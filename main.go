package main

import (
	_ "gf-decoration/boot"
	_ "gf-decoration/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
