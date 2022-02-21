package b

import (
	"fmt"

	dep "github.com/flipb/govendor-dep-1"
)

const moduleName = "b (app-2)"

type SomeExposedType struct {
	N uint
}

func (s *SomeExposedType) DoStuff() {
	fmt.Printf("Will greet from %s\n", moduleName)
	for i := 0; i <= int(s.N); i++ {
		dep.Hello()
		println("")
	}
}
