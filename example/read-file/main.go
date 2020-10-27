package main

import (
	"fmt"

	. "github.com/chengxuncc/saber"
)

func main() {
	Cat(`.gitignore`).Println().Run()
	fmt.Println(Cat(`.gitignore`).Output())
}
