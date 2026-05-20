package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

func main() {
	checkerModule := Checker()

	pgs.Init(pgs.DebugEnv("DEBUG_PG_CHECK")).
		RegisterModule(checkerModule).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
	checkerModule.ExitCheck()
}
