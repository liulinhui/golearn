package main

import (
	"fmt"
	"github.com/liulinhui/golearn/http/fileServer"
	"github.com/liulinhui/golearn/math"
)

func init() {
	fmt.Println("init func done!")
}

func main() {
	fmt.Println("hello world!")
	math.Math()
	math.Float()
	fileServer.FileServer()
}
