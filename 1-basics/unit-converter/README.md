# Unit Converter

An interactive command-line unit converter written in Go. Converts between
units of temperature, length, and weight. Built as the fourth exercise in
my Go learning journey.

## What it does

Provides a REPL (read-eval-print loop) where you type expressions like
`5 km to mi` and get back the conversion. Supports three categories of
units — temperature, length, and weight — with built-in help showing
the unit codes for each category.

The project is split into three packages: `main` (CLI parsing and I/O),
`units` (the conversion domain logic), and `help` (help text).

## Building

From the project directory:

```bash
go build .
```

This creates an executable named `unit-converter` (or whatever your directory
is named) in the current directory.

Alternatively, run without building:

```bash
go run .
```

## Usage

The program is fully interactive — no command-line arguments needed.

```
$ ./unit-converter
Welcome to Unit Converter! Type 'help' to see the available commands.
> 25 c to f
25 c = 77 f
> 10 km to mi
10 km = 6.2137119223733395 mi
> 500 g to lb
500 g = 1.1023113109243879 lb
> help temperature
Temperature units:
	c    - Celsius
	f    - Fahrenheit
	k    - Kelvin

Example: 25 c to f
> quit
See you later!
```

## Commands

| Command                       | Description                                   |
|-------------------------------|-----------------------------------------------|
| `<value> <unit1> to <unit2>`  | Convert `<value>` from `<unit1>` to `<unit2>` |
| `help`                        | Show the list of available commands           |
| `help <category>`             | Show the unit codes for a category            |
| `quit` / `exit`               | Exit the program                              |
| `Ctrl+D` (EOF)                | Exit gracefully with a goodbye message        |

Input is case-insensitive (`5 KM TO MI`, `5 km to mi`, `5 Km to Mi` all work).

## Supported categories and units

### Temperature

| Code | Unit       |
|------|------------|
| `c`  | Celsius    |
| `f`  | Fahrenheit |
| `k`  | Kelvin     |

### Length (base unit: metre)

| Code | Unit       |
|------|------------|
| `m`  | metre      |
| `km` | kilometre  |
| `cm` | centimetre |
| `mi` | mile       |
| `ft` | foot       |

### Weight (base unit: gram)

| Code | Unit      |
|------|-----------|
| `g`  | gram      |
| `kg` | kilogram  |
| `mg` | milligram |
| `t`  | tonne     |
| `lb` | pound     |
| `oz` | ounce     |

## Behavior and validation

- **Case-insensitive input.** The entire input line is lowercased before
  parsing, so `5 KM to MI` works the same as `5 km to mi`.
- **Empty input** (just pressing Enter) is silently ignored.
- **Malformed commands** (wrong number of tokens, missing `to` separator)
  print a usage hint and wait for the next command.
- **Invalid number** in the value position (e.g. `abc km to mi`) prints
  an error and waits for the next command.
- **Unknown unit codes** print an error suggesting `help`.
- **Mixed-category conversion** (e.g. `5 km to f` — length to temperature)
  is rejected with an error.
- **Leading and trailing whitespace** is automatically trimmed.

## Project structure

```
unit-converter/
├── go.mod              # Go module definition
├── main.go             # CLI layer: REPL, command parsing, output
├── help/
│   └── help.go         # help text functions
└── units/
    └── units.go        # domain layer: Units type and conversion logic
```

### Separation of concerns

- **`main` package** handles user interaction: reading input, parsing
  commands, dispatching to the right handler, printing results.
- **`units` package** contains the domain logic: the `Units` type with
  its conversion methods, the `Category` enum, and the unit definitions.
  It exposes a public API (`NewUnits`, `Convert`, `TypeOfCategory`)
  and keeps helpers like `bothInMap`, `toCelsius`, `fromCelsius` and the
  individual `convertX` methods private to the package.
- **`help` package** contains only the help-text functions, separated
  out so the lengthy `fmt.Println` blocks don't clutter the conversion
  logic.

The `units` package is reusable — the same logic could power a web API
or GUI application without modification, since it has no dependencies on
the CLI layer.

## Known limitations

- **No persistence.** The converter is stateless and keeps no history
  between sessions.
- **Three categories only.** No support for currency, time, area, volume,
  speed, energy, data size, or other units.
- **No precision control.** Results are printed with Go's default `%g`
  formatting, which can show up to 16 significant digits. There's no way
  to round to e.g. 2 decimal places from the CLI.
- **Negative-Kelvin temperatures are accepted.** Physically impossible
  (absolute zero is 0 K), but the program does not validate this.
- **Basic line editing.** Input uses `bufio.Scanner`, which means arrow
  keys, command history, and other readline-style features are not
  supported. See "Future improvements" below.

## Exit codes

| Code | Meaning     |
|------|-------------|
| `0`  | Normal exit |

The program currently never exits with a non-zero code since all errors
are recoverable within the REPL loop.

## Future improvements

- **Better line editing — arrow keys, history, Ctrl+R search.**
  The current implementation uses `bufio.Scanner`, which reads input
  line-by-line and does not interpret special keys. Replacing it with
  a readline-style library would give:
    - command history navigation with up/down arrows,
    - in-line cursor movement with left/right arrows,
    - line editing shortcuts (`Ctrl+A`, `Ctrl+E`, `Ctrl+W`, `Ctrl+U`),
    - reverse history search (`Ctrl+R`).

  Candidates worth investigating: `github.com/chzyer/readline`,
  `github.com/peterh/liner`, `github.com/c-bata/go-prompt`.
- **More categories.** Currency (with live exchange rates from an API),
  time, area, volume, speed, energy, data size, fuel economy.
- **Non-interactive mode.** Accept a one-off conversion from CLI arguments:
  `./unit-converter 5 km to mi` — useful for shell scripts and pipelines.
- **Precision control.** A `precision` command to set the number of
  decimal places shown (e.g. `precision 2` → `5 km = 3.11 mi`).
- **Output formatting options** — show as fraction (`1/2 mi`), scientific
  notation, or with thousand separators.
- **Reverse-direction conversion shortcut** like `5 km in mi` as an alias
  for `to`, for a more conversational feel.
- **Unit tests** for the `units` package, especially for the conversion
  formulas (temperature offsets, length/weight multipliers) and edge cases.
- **Negative-Kelvin validation** — return an error for physically
  impossible temperatures.
- **Refactor to an interface-based design** if the number of categories
  grows past 5–6. A `Converter` interface (with `Convert` and `HasUnits`
  methods) would let each category live in its own file and register
  itself at startup, removing the central `switch` in `Convert`.