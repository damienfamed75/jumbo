package main

// // #include "./lib/hello.c"

/*
#cgo LDFLAGS: -L./lib -lhello
#include "./lib/hello.h"
*/
import "C"

import (
	"fmt"
)

func main() {
	orig := "Gopher"
	str := C.CString(orig)

	C.hello(str, C.uint(len(orig)))

	fmt.Printf("%s\n", C.GoString(str))
}
