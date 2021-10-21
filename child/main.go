package main

import (
	"fmt"
	"time"

	v8 "rogchap.com/v8go"
)

func main() {
	s := "function foo() { return 'bar'; }; foo()"

	iso := v8.NewIsolate()
	defer iso.Dispose()

	ctx := v8.NewContext(iso)
	defer ctx.Close()

	start := time.Now()
	ctx.RunScript(s, "script.js")
	fmt.Printf("\nduration to run script %s\n", time.Since(start))
}
