package main

import (
	"fmt"

	. "github.com/chengxuncc/saber"
)

func main() {
	Main.Debug = true
	Cat(`.gitignore`).Run()
	fmt.Println(Cat(`.gitignore`).Output())
}
