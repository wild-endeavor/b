package main

import cc "github.com/wild-endeavor/c"
import "fmt"

func HelloB() string {
	return "hello from b"
}

func main() {
	fmt.Println(cc.Hello2())
}
