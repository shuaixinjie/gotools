package flag

import (
	"flag"
	"fmt"
	"os"
)

// OsArgs 最简单粗暴的读取参数
func OsArgs() []string {
	args := os.Args
	return args[1:]
}

func Flag() {
	var a string
	flag.StringVar(&a, "a", "haha", "is string")
	flag.Parse()
	fmt.Println("==================>>>", a)
}
