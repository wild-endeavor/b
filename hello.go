package hello

import cc "github.com/wild-endeavor/c"

func HelloB() string {
	return "hello from b"
}

func HelloC() string {
	return cc.Hello2()
}
