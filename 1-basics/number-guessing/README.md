# Number Guessing Game

A simple command-line "higher or lower" guessing game written in Go.
The computer picks a random number between 1 and 100, and you try to guess it.
Built as the second exercise in my Go learning journey.

## What it does

On start, the program picks a random integer between 1 and 100. You enter
guesses one by one, and the program responds with `Higher!` or `Lower!`
until you find the number. When you do, it tells you how many tries it took.

You can give up at any time by typing `quit`.

## Building

From the project directory:

```bash
go build .
```

This creates an executable named `guessing-game` (or whatever your directory
is named) in the current directory.

Alternatively, run without building:

```bash
go run .
```

## Usage

The game is fully interactive — no command-line arguments needed.

```bash
$ ./guessing-game
I have a randomly generated number in my memory (between 1 and 100).
Let's play a game – guess what the number is – I'll tell you whether it's 'higher' or 'lower'.
Good luck! [Type 'quit' if you're giving up]
>  50
Lower!
Try again!
>  25
Higher!
Try again!
>  37
Higher!
Try again!
>  42
You guessed it in 4 tries!
```

### Special commands

| Input            | Action                                      |
|------------------|---------------------------------------------|
| Any number 1-100 | Submit as a guess                           |
| `quit`           | Give up and exit the game                   |
| `Ctrl+D` (EOF)   | Exit gracefully with a goodbye message      |

## Behavior and validation

- **Non-numeric input** (e.g. `abc`) prints a hint and asks for another input.
  The attempt counter is *not* incremented.
- **Numbers outside the 1-100 range** print a warning. (Note: in the current
  version, this still counts toward the attempt count — see "Known limitations".)
- **Whitespace around input** is automatically trimmed.
- The random number is freshly generated on every run.

## Known limitations

- **Out-of-range guesses still increment the attempt counter.** A guess of `200`
  is rejected as out of range, but the internal counter goes up. To be fixed in
  a future iteration.
- **Hardcoded range** of 1-100. Not configurable without code changes.
- **No upper limit on attempts** — you can keep guessing forever.
- **No persistence** — statistics aren't kept between game sessions.

## Exit codes

| Code | Meaning                |
|------|------------------------|
| `0`  | Normal exit (win/quit) |

The game currently never exits with a non-zero code since all errors
(invalid input, etc.) are recoverable within the loop.

## Project structure

```
guessing-game/
├── go.mod        # Go module definition
└── main.go       # All program logic (single-file project
```

## Future improvements

- **Refactor into a `Game` struct.** Currently all state lives as loose variables
  in `main()` — `secret`, `attempts`, `min`, `max`. Wrapping these in a `Game`
  struct with methods like `Play()`, `checkGuess()`, `readGuess()` would make
  the code more organized and easier to extend.
- **Configurable range via CLI arguments.** For example: `./guessing-game 1 1000`.
- **Difficulty levels.** Easy (1-50), medium (1-100), hard (1-1000).
- **Attempt limits.** Lose the game after N failed guesses.
- **"Hot/cold" hints** based on distance from the target instead of just
  higher/lower.
- **Guess history** displayed in the prompt (e.g. `[tried: 50, 25, 37] > `).
- **Multiple rounds in one session** with summary statistics (best score,
  average attempts).
- **Reverse mode** — you pick a number, the computer guesses it using a
  binary-search strategy.
- **Unit tests** for the comparison logic, especially once it's refactored
  out of `main()`.
- **Fix the attempt-counter bug** for out-of-range guesses (see limitations).