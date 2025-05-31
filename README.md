# Brainfuck Runner

> A simple, clean Brainfuck interpreter written in Go with REPL support.

**Brainfuck Runner** is a minimalist yet powerful interpreter for the [Brainfuck programming language](https://en.wikipedia.org/wiki/Brainfuck), implemented in Go. This tool is designed for both educational and practical use—allowing you to experiment with Brainfuck interactively via a REPL (Read-Eval-Print Loop), or execute static Brainfuck programs.

Whether you're a curious programmer exploring esoteric languages or someone building tooling for programming language education, this project offers a clean and well-structured Brainfuck interpreter that’s easy to extend, understand, and test.

![GitHub](https://img.shields.io/github/license/BaseMax/brain-fuck-runner)
![Go Version](https://img.shields.io/badge/Go-1.22.4+-brightgreen)

## Features

- Fully functional Brainfuck interpreter
- REPL (interactive mode)
- Supports loops, input/output, and all Brainfuck commands
- Unit-tested core engine
- Clean, modular design

## Usage

### REPL

Run the program and start entering Brainfuck code directly:

```bash
go run main.go
```

Example:

```brainfuck
>> ++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.
```

**Output:** Hello World!

## Run with input file (optional)

If you prefer to load code from a file:

```go
// Example to load from a file
data, _ := os.ReadFile("program.bf")
interpreter := bf.New(string(data))
output, _ := interpreter.Execute()
```

## Project Structure

```bash
brain-fuck-runner/
├── bf/
│   ├── bf.go         # Brainfuck interpreter core
│   └── bf_test.go    # Unit tests
├── main.go           # REPL interface
├── go.mod            # Go module file
├── LICENSE           # MIT License
└── README.md         # This file
```

## License

This project is licensed under the MIT License – see the LICENSE file for details.

Made with ❤️ by Max Base

Copyright (c) 2025 Max Base
