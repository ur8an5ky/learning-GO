# To-do List CLI

An interactive command-line to-do list manager written in Go.
Data lives only in memory and is lost when the program exits.
Built as the third exercise in my Go learning journey.

## What it does

Provides a REPL (read-eval-print loop) for managing tasks: add, list,
mark as done, and delete. Tasks have auto-generated incremental IDs
and a done/not-done status displayed in the list.

The project is split into two packages: `main` (CLI parsing and I/O)
and `tasks` (the actual task domain logic).

## Building

From the project directory:

```bash
go build .
```

This creates an executable named `to-do-list` (or whatever your directory
is named) in the current directory.

Alternatively, run without building:

```bash
go run .
```

## Usage

The program is fully interactive — no command-line arguments needed.

```
$ ./to-do-list
Welcome to to-do! Type 'help' to see the available commands.
> add Buy groceries
Added task #1: Buy groceries
> add Take out the trash
Added task #2: Take out the trash
> add Learn Go
Added task #3: Learn Go
> list
List of tasks:
	[ ]  #1: Buy groceries
	[ ]  #2: Take out the trash
	[ ]  #3: Learn Go
> done 2
Task #2 marked as done
> list
List of tasks:
	[ ]  #1: Buy groceries
	[X]  #2: Take out the trash
	[ ]  #3: Learn Go
> delete 1
Task #1 deleted
> quit
See you later!
```

## Commands

| Command         | Description                              |
|-----------------|------------------------------------------|
| `add <text>`    | Add a new task with the given text       |
| `list`          | Display all tasks with their status      |
| `done <id>`     | Mark task with the given ID as done      |
| `delete <id>`   | Remove the task with the given ID        |
| `help`          | Show the list of available commands      |
| `quit` / `exit` | Exit the program                         |
| `Ctrl+D` (EOF)  | Exit gracefully with a goodbye message   |

Commands are case-insensitive (`ADD`, `Add`, `add` all work).

## Behavior and validation

- **Task IDs** start at 1 and increase monotonically. Deleting a task
  does *not* reuse its ID — subsequent tasks keep their original numbering.
- **Non-numeric ID** in `done` or `delete` (e.g. `done abc`) prints an error
  and waits for the next command.
- **Missing arguments** (`add` without text, `done`/`delete` without an ID)
  print a usage hint and wait for the next command.
- **Non-existent ID** in `done` or `delete` (e.g. `delete 999` when only
  tasks 1–3 exist) prints an error and waits for the next command.
- **Empty input** (just pressing Enter) is silently ignored.
- **Unknown commands** print an error suggesting `help`.
- **Whitespace** around input is automatically trimmed.

## Project structure

```
to-do-list/
├── go.mod              # Go module definition
├── main.go             # CLI layer: REPL, command parsing, output
└── tasks/
    └── tasks.go        # domain layer: Task and TaskList types and methods
```

### Separation of concerns

- **`main` package** handles user interaction: reading input, parsing
  commands, printing results.
- **`tasks` package** contains the domain logic: the `Task` and `TaskList`
  types and operations on them. It exposes a public API
  (`NewTaskList`, `Add`, `List`, `Done`, `Delete`) and keeps internal
  state (the task slice, the next ID counter) private to the package.

This separation makes the `tasks` package reusable — the same logic
could power a web server or a GUI without modification.

## Known limitations

- **No persistence.** All tasks are lost when the program exits.
  A future iteration will add JSON-based persistence (planned as a separate project).
- **No editing.** A task's text cannot be changed after creation —
  only deleted and re-added.
- **No `undone` command.** A task marked as done cannot be reverted to
  not-done.
- **No filtering.** `list` always shows all tasks, with no way to view
  only completed or pending ones.
- **Basic line editing.** Input uses `bufio.Scanner`, which means arrow
  keys, command history, and other readline-style features are not
  supported. See "Future improvements" below.

## Exit codes

| Code | Meaning                |
|------|------------------------|
| `0`  | Normal exit            |

The program currently never exits with a non-zero code since all errors
are recoverable within the loop.

## Future improvements

- **Persistence via JSON.** Save tasks to a file on every change and
  load them at startup. Will be implemented as a separate, larger project
  to focus on `encoding/json` and file I/O.
- **`edit <id> <text>`** command to change a task's description.
- **`undone <id>`** command to revert a completed task to not-done.
- **Filtering in `list`:** `list done` and `list todo` to show only
  completed or pending tasks.
- **Priorities** for tasks (e.g. `add !high Important task`) with
  sorting in `list`.
- **Tags / categories** for grouping (e.g. `add Buy milk #shopping`).
- **Due dates** for tasks.
- **Counter** at the end of `list` ("3 tasks, 1 done, 2 pending").
- **Colored output** — green for completed tasks, red for pending —
  using a library like `fatih/color`.
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
- **Unit tests** for the `tasks` package, exercising add/list/done/delete
  and their error cases.
- **Subcommands** (in style of `git`/`docker`) — e.g.
  `to-do-list add "task"` from the shell, without entering REPL mode.