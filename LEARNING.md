# Learning Guide: Structs and Maps

A step-by-step guide to practice Go structs and maps by converting the todo application.

---

## Step 1: Struct

### Concept
A struct groups related fields together (like an object in other languages). Each todo item has properties: ID, Task, and Done status.

### What to Do

1. Define a `Todo` struct with these fields:
   - `ID` (int) - unique identifier
   - `Task` (string) - task description
   - `Done` (bool) - completion status

2. Change `Todos` from `[]string` to `[]Todo`

3. Add a counter variable (e.g., `nextID`) to generate unique IDs

4. Update your functions:
   - `PrintList` - access `todo.ID` and `todo.Task`
   - `AddTodo` - create `Todo{ID: nextID, Task: ..., Done: false}`
   - `Remove` - find todo by ID, remove from slice
   - `ToggleDone` - find by ID and flip Done status
   - `EditTask` - find by ID and update Task field

### Practice Exercises

1. Add a `ToggleDone` function to mark task complete/incomplete
2. Update `PrintList` to show `[x]` or `[ ]` based on Done status
3. Add an `edit` command to change a task title
4. Add a `toggle` command that flips the Done status

### Hints

- Access struct fields with dot notation: `todo.Task`
- Use `slices.IndexFunc` to find by ID:
  ```go
  index := slices.IndexFunc(Todos, func(t Todo) bool {
      return t.ID == id
  })
  ```
- Use `slices.DeleteFunc` to remove by condition

---

## Step 1.5: Method Receivers (Like PHP Classes)

### Concept
In Go, you can attach functions to structs using **method receivers**. This is similar to methods in PHP classes.

### Comparison

**PHP Class:**
```php
class Todo {
    public $ID;
    public $Task;
    public $Done;

    public function toggle() {
        $this->Done = !$this->Done;
    }
}
```

**Go Method Receiver:**
```go
func (t *Todo) Toggle() {
    t.Done = !t.Done
}
```

### What to Do

1. Create methods on the `Todo` struct:
   - `Toggle()` - flip Done status
   - `Edit(newTask string)` - update Task field

2. Create methods on `[]Todo` (slice type):
   - Define a new type: `type TodoList []Todo`
   - Add methods: `Add`, `Remove`, `FindByID`, `Print`

3. Call methods like: `todo.Toggle()` or `todoList.Add(task)`

### Practice Exercises

1. Convert `ToggleDone` to a method on `Todo`
2. Convert `AddTodo`, `Remove`, `PrintList` to methods on `TodoList`
3. Add a `FindByID(id int) *Todo` method
4. Refactor `Init` to use method calls

### Hints

- Receiver type `*Todo` means pointer (can modify)
- Receiver type `Todo` means copy (read-only)
- Define custom type for slice:
  ```go
  type TodoList []Todo

  func (list *TodoList) Add(task string) {
      // implementation
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

Complete each step before moving to the next. Step 1 → Step 1.5 → Step 2.
