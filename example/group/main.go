package main

import (
	"fmt"
	"strconv"

	. "github.com/chengxuncc/saber"
)

func main() {
	Main.Debug = true
	Group(
		Echo("Group Run"),
		Echo(1),
		Echo(2),
	).Run()
	fmt.Println(strconv.Quote(
		Group(
			Echo("Group Output"),
			Echo(3),
			Echo(4),
		).Output(),
	))
}
