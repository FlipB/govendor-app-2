package main

import (
	"fmt"

	"github.com/flipb/govendor-app-2/pkg/b"
)

func main() {
	fmt.Printf("Starting b app-2\n")
	x := b.SomeExposedType{
		N: 0,
	}
	x.DoStuff()
}
