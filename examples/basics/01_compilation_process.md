# The Compilation Process

Computers execute machine code. They do not understand natural language or high-level programming languages such as Go. To run a Go program, the source code must first be translated into machine instructions that the target hardware (your CPU) can execute.

This translation is performed by the Go compiler.

The result of the compilation process is a standalone executable file:
- On **Windows**, an `.exe` file
- On **macOS** and **Linux**, a standard executable binary

---

## Go Program Structure

A minimal Go program consists of several required elements. Each has a specific purpose in the compilation and execution process.

### `package main`

The `package main` declaration indicates that the program is intended to be compiled as an executable. Programs in the `main` package produce a runnable binary, rather than a library meant to be imported by other Go programs.

### `import`

The `import` statement allows a program to use functionality from other packages. For example:

```go
import "fmt"
```
### `import "fmt"`

The `fmt` package is part of the Go standard library. It provides formatted input and output functions, such as `fmt.Println`.

### `func main()`

The `main` function is the entry point of a Go program. Execution begins in this function when the compiled program is run.

A program must define `func main()` in the `main` package in order to be executable.

---

## Types of Errors

Programming errors generally fall into two categories.

### Compilation Errors

Compilation errors occur during the compilation step, before a program is executed. These errors prevent the compiler from producing an executable binary.

Compilation errors are beneficial in that they cannot reach production. A program with compilation errors cannot be built or distributed.

Examples include:
- Syntax errors
- Type mismatches
- Undefined variables or functions

### Runtime Errors

Runtime errors occur while a program is running. These errors can cause a program to crash or behave unexpectedly.

Runtime errors are more dangerous because they may only appear under certain conditions and can affect users of a deployed application.

Examples include:
- Division by zero
- Accessing invalid memory
- Nil pointer dereferences

---

## Compilation and Execution in the Browser

In browser-based environments, compilation and execution often happen in a single step. This can make it less obvious whether an error is a compilation error or a runtime error.

Despite this, the same distinction applies: code must successfully compile before it can be executed.

When using online Go playgrounds or similar tools, the environment typically handles both compilation and execution behind the scenes. If an error occurs, it will indicate whether it is a compilation error (for example, a syntax error) or a runtime error (for example, a panic during execution).


### Compilation errors are beneficial in that they cannot reach production. A program with compilation errors cannot be built or distributed.

#### Examples include:
- Syntax errors
- Type mismatches
- Undefined variables or functions

