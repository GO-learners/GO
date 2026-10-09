# Control Flow in Go

This section covers control flow structures in Go, including conditional statements, loops, and switch statements.

## 📂 Directory Structure

```
02-control-flow/
├── if-else/
│   ├── if-else.go
│   └── README.md
├── loops/
│   ├── loops.go
│   └── README.md
├── switch/
│   ├── switch.go
│   └── README.md
└── README.md
```

## 📚 Topics Covered

### 1. If-Else Statements (`if-else/`)
Learn about conditional branching in Go:
- Basic `if`, `else if`, and `else` syntax
- Conditions without parentheses
- If statements with short variable declarations
- Logical operators (`&&`, `||`, `!`)

**Key Feature:** Go requires braces even for single-line statements, making code more consistent and reducing bugs.

### 2. Loops (`loops/`)
Master Go's unified looping construct:
- Traditional `for` loop with three components
- While-style loops using `for` with only a condition
- Infinite loops with `break` and `continue`
- Loop variable scoping

**Key Feature:** Go simplifies looping by having only `for` loops, which can be configured to behave like `while` or infinite loops.

### 3. Switch Statements (`switch/`)
Explore Go's powerful switch capabilities:
- Basic switch with value matching
- Switch without expression (condition-based)
- Automatic break (no fall-through by default)
- Multiple values per case
- Type switches

**Key Feature:** Unlike C/Java, Go's switch cases don't fall through automatically, making them safer and more intuitive.

## 🚀 Getting Started

Each subdirectory contains:
- A `.go` file with example code
- A `README.md` with detailed explanations

To run any example:

```bash
# Navigate to a specific folder
cd if-else/

# Run the example
go run if-else.go
```

Or run directly from the control-flow directory:

```bash
go run if-else/if-else.go
go run loops/loops.go
go run switch/switch.go
```

## 🎯 Key Takeaways

1. **Simplicity:** Go's control flow is minimal and consistent
2. **No Parentheses:** Conditions don't require parentheses, but braces are mandatory
3. **One Loop to Rule Them All:** The `for` loop handles all iteration needs
4. **Safe Switches:** Automatic break prevents common fall-through bugs
5. **Short Statements:** Initialize variables in if/switch for better scoping

## 📖 Learn More

- [Go by Example: If/Else](https://gobyexample.com/if-else)
- [Go by Example: For](https://gobyexample.com/for)
- [Go by Example: Switch](https://gobyexample.com/switch)
- [Effective Go: Control Structures](https://go.dev/doc/effective_go#control-structures)

## 🔄 What's Next?

After mastering control flow, move on to:
- Functions and methods
- Arrays and slices
- Maps and structs
- Error handling patterns
---

## ?? Best Practices

- **Write Clean Code:** Always format your code using `go fmt`.
- **Use Meaningful Names:** Choose descriptive names for variables and functions.
- **Keep it Simple:** Avoid overly complex logic when a simpler solution exists.
- **Document Your Code:** Add comments to explain *why*, not just *what*.
- **Handle Errors:** Always check for and handle errors gracefully.
