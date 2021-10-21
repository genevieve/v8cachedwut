package main

import (
	"fmt"
	"log"
	"os/exec"
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

	cmd := exec.Command("go", "run", "child/main.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
