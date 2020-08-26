package main

import (
	"fmt"
	"go-test/util"
)

func main() {
	go util.Test(1,2)
	fmt.Print(11111)
}
