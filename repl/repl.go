package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/emilkloeden/monkey/evaluator"
	"github.com/emilkloeden/monkey/lexer"
	"github.com/emilkloeden/monkey/object"
	"github.com/emilkloeden/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := parser.New(l, cwd)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
