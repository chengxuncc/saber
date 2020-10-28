package main

import (
	"fmt"

	. "github.com/chengxuncc/saber"
)

func main() {
	Cat(`.gitignore`).EchoN().Run()
	fmt.Println(Cat(`.gitignore`).Output())
}
