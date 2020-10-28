package main

import (
	"fmt"

	. "github.com/chengxuncc/saber"
)

func main() {
	Cat(`.gitignore`).Echon().Run()
	fmt.Println(Cat(`.gitignore`).Output())
}
