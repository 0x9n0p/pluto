package main

import (
	"fmt"
	"pluto"
)

func main() {
	fmt.Printf("%s %s\n", pluto.Name, pluto.Version)
	select {}
}
