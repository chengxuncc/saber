package main

import (
	. "github.com/chengxuncc/saber"
)

func main() {
	Run(
		Grep("idea"),
		Cat(".gitignore").Grep(`idea`),
	)
}
