package main

import (
	"fmt"

	. "github.com/chengxuncc/saber"
)

func main() {
	Run(
		Cmd(`go`, `version`),
	)
	fmt.Println(Cmd(`go`, `version`).Output())
}
