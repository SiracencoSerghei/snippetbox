# Basic Variables (Go)

This lesson introduces **basic variable types in Go**, explains how variables are declared, and demonstrates **zero values**.

---

## Basic Variable Types

### `bool`
A boolean value. It can be either:
- `true`
- `false`

**Zero value:** `false`

---

### `string`
A sequence of characters (text).

Examples:
- `"hello"`
- `"Go is fun"`

**Zero value:** empty string `""`

---

### `int`
A signed integer number.

Examples:
- `-3`
- `0`
- `42`

**Zero value:** `0`

---

### `float64`
A floating‑point number with double precision.

Examples:
- `3.14`
- `0.5`
- `-12.75`

**Zero value:** `0.0`

---

### `byte`
Represents **8 bits of data** (alias for `uint8`).
Often used for raw data or ASCII characters.

Examples:
- `65` (represents `'A'`)
- `0xFF`

**Zero value:** `0`

---

## Declaring a Variable ("The Sad Way")

```go
var mySkillIssues int
mySkillIssues = 42
```

### What happens here?
1. `var mySkillIssues int`
   - Declares a variable named `mySkillIssues`
   - Assigns it the **zero value** of `int` → `0`

2. `mySkillIssues = 42`
   - Overwrites the zero value with `42`

This works, but it is **verbose and not idiomatic Go**.

👉 A better way will be shown in the next lesson.

---

## Declaring a Variable (Idiomatic Go Way)

```go
mySkillIssues := 42
```

### Why this is preferred
- Uses **short variable declaration** (`:=`)
- Go **infers the type automatically** (`int` in this case)
- Combines declaration and initialization in one step
- Cleaner, shorter, and easier to read

This form is the **standard and idiomatic way** to declare and initialize variables inside functions in Go.

> Note: `:=` can only be used **inside functions**. At package level, `var` is still required.

---

## Assignment

Initialize the variables used in the `print` statement **with their zero values**.

### Required types:
- `int`
- `float64`
- `bool`
- `string`

### Example solution:

```go
var i int
var f float64
var b bool
var s string
```

These variables will automatically have the following values:

| Variable | Type     | Zero Value |
|--------|----------|------------|
| `i`    | `int`    | `0`        |
| `f`    | `float64`| `0.0`      |
| `b`    | `bool`   | `false`    |
| `s`    | `string` | `""`       |

---

## Key Takeaway

> **Every variable in Go always has a value.**  
> If you don’t assign one, Go assigns the **zero value** automatically.

This is one of the reasons Go code is safer and easier to reason about.


