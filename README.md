# The Interpreter

A tree-walking interpreter implementation written in Go, built by following the book "Writing An Interpreter In Go" by Thorsten Ball.

## About

This project implements a fully functional interpreter for a dynamically-typed programming language. The interpreter parses and evaluates code through lexical analysis, parsing, and direct AST evaluation.

## Features

- **Lexer**: Tokenizes source code into meaningful tokens
- **Parser**: Builds an Abstract Syntax Tree (AST) from tokens using Pratt parsing
- **Evaluator**: Tree-walking interpreter that executes the AST
- **Data Types**: Integers, booleans, strings, arrays, and hash maps
- **Functions**: First-class and higher-order functions with closures
- **Built-in Functions**: Common operations like `len()`, `println()`, `first()`, `last()`, etc.
- **Control Flow**: If-else expressions
- **Operators**: Arithmetic, comparison, and logical operators

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

```bash
git clone https://github.com/MTZACK9/the-interpreter
cd the-interpreter
go build
```

### Running the REPL

```bash
./the-interpreter
```

### Example Code

```
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      1
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};

fibonacci(10);
```

## Project Structure

```
.
├── token/          # Tokens
├── lexer/          # Tokenization
├── parser/         # Parsing and AST construction
├── ast/            # Abstract Syntax Tree definitions
├── evaluator/      # Tree-walking interpreter
├── object/         # Runtime object system
├── repl/           # Read-Eval-Print Loop
└── main.go         # Entry point
```

## Learning Journey

This interpreter was built as a learning project following "Writing An Interpreter In Go". The process involved:

1. Implementing a lexer to tokenize source code
2. Building a recursive descent parser with Pratt parsing techniques
3. Defining an AST to represent the program structure
4. Creating an evaluator that walks the AST and executes code
5. Adding support for various data types and operations
6. Implementing first-class functions with closures

## Testing

Run the test suite:

```bash
go test ./...
```

## Acknowledgments

- Thorsten Ball for the excellent book "Writing An Interpreter In Go"
- The Go community for great documentation and tools

## License

MIT License - feel free to use this code for learning purposes.

## What's Next?

Potential improvements and extensions:
- Add more built-in functions
- Implement a standard library
- Add error handling improvements
- Create a compiler (see "Writing A Compiler In Go")
