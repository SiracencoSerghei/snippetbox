# How Go's Return-Type Rules Shape API Design

This document explains how Go's strict return-type system directly
influences API design decisions. It is intended for developers coming
from Python/JavaScript backgrounds.

------------------------------------------------------------------------

## Core Rule

In Go, **a function must always return values of the same static type**.

This is enforced at **compile time**, not runtime.

``` go
func f(x int) ???
```

There is **no valid return type** that allows: - sometimes `string` -
sometimes `number`

------------------------------------------------------------------------

## Why This Matters for API Design

Because Go forbids dynamic return types, API designers must be
**explicit** about:

-   Success vs failure
-   Presence vs absence of a value
-   Valid vs invalid input

This leads to predictable, self-documenting APIs.

------------------------------------------------------------------------

## Pattern 1: `(value, error)` --- The Standard Go API

### Function signature

``` go
func ReadFile(path string) ([]byte, error)
```

### Why this exists

Instead of returning: - data OR string error - data OR nil OR false

Go forces: - a value of known type - an explicit `error`

### Usage

``` go
data, err := ReadFile("file.txt")
if err != nil {
    return err
}
```

✔ No ambiguity\
✔ No exceptions\
✔ Caller must handle failure

------------------------------------------------------------------------

## Pattern 2: `(value, ok)` --- Lookup APIs

### Example: map access

``` go
value, ok := myMap["key"]
```

### API design principle

Instead of returning: - value OR nil

Go returns: - value - boolean indicating presence

### Custom API example

``` go
func FindUser(id int) (User, bool)
```

------------------------------------------------------------------------

## Pattern 3: Sentinel Errors

### Definition

``` go
var ErrNotFound = errors.New("not found")
```

### API usage

``` go
func LoadConfig() error
```

Caller can check:

``` go
if errors.Is(err, ErrNotFound) {
    // handle specific case
}
```

✔ Typed control flow\
✔ No string matching

------------------------------------------------------------------------

## Pattern 4: Single Return Type with Encoded State

Sometimes APIs return a single struct that contains state:

``` go
type Result struct {
    Value int
    Valid bool
}
```

Used when: - multiple fields matter together - extensibility is expected

------------------------------------------------------------------------

## What Go APIs Avoid (By Design)

❌ Returning `interface{}`\
❌ Returning different types per condition\
❌ Throwing exceptions\
❌ Hidden control flow

These patterns are common in dynamic languages but rejected in Go.

------------------------------------------------------------------------

## Comparison with Python / JavaScript

### Python

``` python
def load():
    if fail:
        return "error"
    return 42
```

❌ Ambiguous\
❌ Runtime-only failure

### Go

``` go
func Load() (int, error)
```

✔ Compile-time guarantees\
✔ Forced error handling

------------------------------------------------------------------------

## Resulting Benefits

-   APIs are predictable
-   Error handling is explicit
-   Tooling (IDE, refactor, grep) works better
-   Code reviews are easier
-   Fewer production surprises

------------------------------------------------------------------------

## Mental Model

> In Go, **function signatures tell the truth**.

If an operation can fail, the type system makes that impossible to
ignore.

------------------------------------------------------------------------

## Summary

Go's ban on mixed return types:

-   shapes how APIs are designed
-   forces clarity at boundaries
-   trades flexibility for reliability

This is not a limitation --- it is a deliberate design choice.

------------------------------------------------------------------------

**Recommended reading:** 
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Proverbs](https://go.dev/doc/proverbs)
- [Go Blog: Error Handling and Go](https://go.dev/blog/error-handling-and-go)
