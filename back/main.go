package main

import (
	_ "back/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"back/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
