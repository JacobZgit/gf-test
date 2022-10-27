package main

import (
	_ "gf-test/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-test/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
