package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mtzack9/the-interpreter/evaluator"
	"github.com/mtzack9/the-interpreter/lexer"
	"github.com/mtzack9/the-interpreter/parser"
)

// REPL stands for "Read Eval Print Loop"

const PROMPT = ">> "
const INTERPRETER_FACE = `
_________________________________________________________________
|  _______ _    _ _____   _____ _   _ _______ ______ _____  _____  |
| |__   __| |  | |  ___| |_   _| \\ | |__   __|  ____|  __ \\|  __ \\ |
|    | |  | |__| | |__     | | |  \\| |  | |  | |__  | |__) | |__) ||
|    | |  |  __  |  __|    | | | .    |  | |  |  __| |  _  /|  ___/ |
|    | |  | |  | | |____  _| |_| |\\  |  | |  | |____| | \\ \\| |     |
|    |_|  |_|  |_|______||_____|_| \\_|  |_|  |______|_|  \\_\\_|     |
|_________________________________________________________________|
|                                                                  |
|   _____  _____  ______ _______ ______ _____                      |
|  |  __ \\|  __ \\|  ____|__   __|  ____|  __ \\                     |
|  | |__) | |__) | |__     | |  | |__  | |__) |                    |
|  |  ___/|  _  /|  __|    | |  |  __| |  _  /                     |
|  | |    | | \\ \\| |____   | |  | |____| | \\ \\                     |
|  |_|    |_|  \\_\\______|  |_|  |______|_|  \\_\\                    |
|                                                                  |
|__________________________________________________________________|`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
