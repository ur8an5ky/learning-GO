# Calculator

A simple command-line calculator written in Go. Built as the first exercise
in my Go learning journey.

## What it does

Takes numbers and an operator from the command line, performs the calculation,
and prints the result. Uses postfix (RPN-like) notation: operator goes **last**.

## Building

From the project directory:

```bash
go build .
```

This creates an executable named `calculator` in the current directory.

Alternatively, you can run without building:

```bash
go run . <args>
```

## Usage
```bash
./calculator <number1> <number2> [<number3> ...] <operator>
```

### Examples

```bash
$ ./calculator 12 15 +
Whole operation: 12 + 15 = 27

$ ./calculator 100 25 -
Whole operation: 100 - 25 = 75

$ ./calculator 2 3 4 +
Whole operation: 2 + 3 + 4 = 9

$ ./calculator 2 3 4 5 '*'
Whole operation: 2 * 3 * 4 * 5 = 120

$ ./calculator 20 4 /
Whole operation: 20 / 4 = 5
```

> **Note:** The `*` character must be quoted (`'*'`) or escaped (`\*`) to prevent
> the shell from expanding it into a list of files in the current directory.

## Supported operations

| Operator | Operation      | Number of operands |
|:--------:|----------------|--------------------|
| `+`      | Addition       | 2 or more          |
| `-`      | Subtraction    | Exactly 2          |
| `*`      | Multiplication | 2 or more          |
| `/`      | Division       | Exactly 2          |

## Limitations and behavior

- **Integer-only output.** All calculations produce integer results. Floating-point
  inputs are accepted but silently truncated (with a warning).
- **No operator precedence.** Each command runs a single operation; there's
  no support for expressions like `2 + 3 * 4`.
- **Division by zero** terminates the program with an error message.
- **Non-numeric arguments** (other than the trailing operator) terminate
  the program with an error message.
- **Wrong number of arguments** for `-` and `/` terminates the program
  with an error message.

## Exit codes

| Code | Meaning                                        |
|:----:|------------------------------------------------|
| `0`  | Success                                        |
| `1`  | Any error (bad input, division by zero, etc.) |

## Project structure
```
calculator/
├── go.mod        # Go module definition
└── main.go       # All program logic (single-file project)
```

## Future improvements ideas

- Floating-point arithmetic support
- Interactive mode (REPL)
- Expression parsing with operator precedence
- Unit tests
- Refactoring calculation logic into a separate package