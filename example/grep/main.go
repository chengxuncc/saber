package main

import (
	. "github.com/chengxuncc/saber"
)

func main() {
	Run(
		Cat(".gitignore").Grep(`idea`),
	)
}
