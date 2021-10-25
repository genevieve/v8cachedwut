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

	start := time.Now()
	data, err := iso.CompileScript(s, "script.js", v8.ScriptCompilerCompileOptionEagerCompile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nduration to compile script %s\n", time.Since(start))
	fmt.Println("parent process sends: ")
	fmt.Println(data)

	cmd := exec.Command("go", "run", "child/main.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	_, err = stdin.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	stdin.Close()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
