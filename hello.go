package hello

import cc "github.com/wild-endeavor/c"
import "fmt"

func HelloB() string {
	return "hello from b"
}

func HelloC() string {
	return cc.Hello2()
}
