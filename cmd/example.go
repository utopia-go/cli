package main

import (
	"fmt"
	"github.com/utopia-go/cli"
)

func main() {
	c := cli.NewCLI(nil)
	fmt.Println(c)
}
