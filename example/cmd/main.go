package main

import (
	"fmt"

	. "github.com/chengxuncc/saber"
)

func main() {
	Run(
		Exec(`go`, `version`),
	)
	fmt.Println(Exec(`go`, `version`).Output())
}
