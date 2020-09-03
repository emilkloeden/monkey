package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/emilkloeden/monkey/evaluator"
	"github.com/emilkloeden/monkey/lexer"
	"github.com/emilkloeden/monkey/object"
	"github.com/emilkloeden/monkey/parser"
	"github.com/emilkloeden/monkey/repl"
)

func main() {
	interactive := flag.Bool("i", false, "interactive mode")
	flag.Parse()
	if *interactive {
		fmt.Printf("interactive mode: %t\n", *interactive)
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else {
		if len(flag.Args()) != 1 {
			fmt.Println("Incorrect usage. Usage `monkey filePath`")
		}

		fileName := flag.Args()[0]
		absolutePath, err := filepath.Abs(fileName)

		if err != nil {
			panic(err)
		}
		fmt.Printf("Loading... %s\n\n", absolutePath)
		fileDir := path.Dir(absolutePath)
		fmt.Printf("From... %s\n\n", fileDir)

		contents, err := ioutil.ReadFile(fileName)

		if err != nil {
			fmt.Printf("Failure to read file '%s'. Err: '%s'",
				string(contents), err)
		}
		l := lexer.New(string(contents))
		p := parser.New(l, fileDir)
		program := p.ParseProgram()

		env := object.NewEnvironment()

		// hacky way to handle not returning anything
		// on the final line of a program
		result, ok := evaluator.Eval(program, env).(object.Object)
		if ok && result.Type() != "NULL" {
			fmt.Println(result.Inspect())
		}
	}
}
