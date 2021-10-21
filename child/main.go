package main

import (
	"fmt"
	"io"
	"os"
	"time"

	v8 "rogchap.com/v8go"
)

func main() {
	line, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nchild process receives: ")
	fmt.Println(line)

	s := "function foo() { return 'bar'; }; foo()"

	iso := v8.NewIsolate()
	defer iso.Dispose()

	ctx := v8.NewContext(iso)
	defer ctx.Close()

	start := time.Now()
	v, err := ctx.RunCompiledScript(s, line, "script.js")
	fmt.Printf("\nduration to run compiled script %s\n", time.Since(start))
	if err != nil {
		panic(err)
	}
	if v.String() != "bar" {
		panic("wrong result")
	}
}
