package main

import (
	. "github.com/chengxuncc/saber"
)

func main() {
	Curl(`http://google.com`).Run()
}
