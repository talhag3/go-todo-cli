# Todo CLI

A simple command-line todo application written in Go.

## Description

This is a learning project to understand Go fundamentals including:
- Structs and custom types
- Slices and pointers
- User input handling
- Basic CLI interactions

## Features

- Add new tasks with auto-generated IDs
- List all tasks with completion status
- Edit existing tasks
- Toggle task completion status
- Remove tasks by ID

## Usage

Run the application:
```bash
go run main.go
```

## Commands

| Command | Alias | Description |
|---------|-------|-------------|
| `list`  | `l`   | Show all tasks |
| `add`   | `a`   | Add a new task |
| `edit`  | `e`   | Edit a task by ID |
| `toggle`| `t`   | Toggle task completion status |
| `remove`| `r`   | Remove a task by ID |
| `q`     |       | Quit application |

## Data Structure

```go
type Todo struct {
    ID   int
    Task string
    Done bool
}
```

## Learning Progress

See [LEARNING.md](./LEARNING.md) for the step-by-step learning guide.

## Author

[talhag3](https://github.com/talhag3)

## License

MIT
