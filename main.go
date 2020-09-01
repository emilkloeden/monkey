package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"github.com/emilkloeden/monkey/evaluator"
	"github.com/emilkloeden/monkey/lexer"
	"github.com/emilkloeden/monkey/object"
	"github.com/emilkloeden/monkey/parser"
	"github.com/emilkloeden/monkey/repl"
	"os"
	"os/user"
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

		filePath := flag.Args()[0]
		contents, err := ioutil.ReadFile(filePath)

		if err != nil {
			fmt.Printf("Failure to read file '%s'. Err: '%s'",
				string(contents), err)
		}
		l := lexer.New(string(contents))
		p := parser.New(l)
		program := p.ParseProgram()

		env := object.NewEnvironment()
		result := evaluator.Eval(program, env)

		fmt.Println(result.Inspect())
	}

}
