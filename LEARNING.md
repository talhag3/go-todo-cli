# Learning Guide: Structs and Maps

A step-by-step guide to practice Go structs and maps by converting the todo application.

---

## Step 1: Struct

### Concept
A struct groups related fields together (like an object in other languages). Each todo item has properties: ID, Title, and Done status.

### What to Do

1. Define a `Todo` struct with these fields:
   - `ID` (int) - unique identifier
   - `Title` (string) - task description
   - `Done` (bool) - completion status

2. Change `Todos` from `[]string` to `[]Todo`

3. Add a counter variable (e.g., `nextID`) to generate unique IDs

4. Update your functions:
   - `PrintList` - access `todo.ID` and `todo.Title`
   - `AddTodo` - create `Todo{ID: nextID, Title: ..., Done: false}`
   - `Remove` - find todo by ID, remove from slice

### Practice Exercises

1. Add a `ToggleDone` function to mark task complete/incomplete
2. Update `PrintList` to show `[x]` or `[ ]` based on Done status
3. Add an `edit` command to change a task title
4. Add a `complete` command that marks a task as done

### Hints

- Access struct fields with dot notation: `todo.Title`
- Loop through slice to find by ID:
  ```go
  for i, todo := range Todos {
      if todo.ID == id {
          // found it at index i
      }
  }
  ```

---

## Step 2: Map

### Concept
A map is a collection of key-value pairs. Great for lookups by ID. Makes delete operations simpler.

### What to Do

1. Change `Todos` from slice to `map[int]Todo`

2. Initialize with `make(map[int]Todo)` or `Todos = map[int]Todo{}`

3. Use ID as the key:
   - Add: `Todos[id] = Todo{...}`
   - Get: `todo := Todos[id]`
   - Delete: `delete(Todos, id)`

4. Update `PrintList` to iterate with `range`

### Practice Exercises

1. Check if ID exists before accessing: `todo, ok := Todos[id]`
2. Add a function to count completed vs pending tasks
3. Add a search function to find tasks by title
4. List only incomplete tasks (filter by Done status)

### Hints

- Check key exists: `todo, exists := Todos[id]`
- `exists` is `true` if key found, `false` otherwise
- Iterate map:
  ```go
  for id, todo := range Todos {
      // use id and todo
  }
  ```
- Map iteration order is random - sort keys if you need order

---

## Summary

| Feature | Struct | Map |
|---------|--------|-----|
| Best for | Grouping data | Fast lookups |
| Access by | Index (slice) | Key |
| Delete | Slice manipulation | `delete(map, key)` |
| Order | Preserved | Random |

---

Choose one step to start with. Complete all exercises before moving to the next step.
