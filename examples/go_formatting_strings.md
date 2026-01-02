# Formatting Strings in Go

Go follows the **printf tradition** from the C language.  
Compared to Python's f-strings, Go’s string formatting is more explicit and type-driven.

---

## Main Formatting Functions

### `fmt.Printf`
- Prints a formatted string directly to [standard output](https://stackoverflow.com/questions/3385201/confused-about-stdin-stdout-and-stderr).
- Does **not** return a string.

```go
fmt.Printf("I am %d years old", 10)
```
- [Official documentation](https://pkg.go.dev/fmt#Printf)
---

### `fmt.Sprintf`
- Formats a string and **returns** it.
- Very useful when you need to store or reuse the result.

```go
s := fmt.Sprintf("I am %d years old", 10)
```
- [Official documentation](https://pkg.go.dev/fmt#Sprintf)
---

## Formatting Verbs

### Default Representation — `%v`
The `%v` verb prints values in their **default format**.
It works with almost any type and is often used as a catch‑all.

```go
s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old
```

---

### String — `%s`
Used for string values.

```go
s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old
```

---

### Integer — `%d`
Used for signed integers (`int`, `int32`, `int64`, etc.).

```go
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old
```

---

### Float — `%f`
Used for floating‑point numbers.

```go
s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old
```

#### Rounding Floats
You can control decimal precision using `.N`.

```go
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old
```

---

## Common Mistakes

❌ Wrong type for formatting verb:
```go
fmt.Sprintf("Age: %d", "ten") // compile error
```

✅ Correct usage:
```go
fmt.Sprintf("Age: %s", "ten")
```

---

## When to Use What

| Situation | Recommended |
|---------|-------------|
| Debugging / quick output | `%v` |
| User‑facing strings | `%s`, `%d`, `%f` |
| Precise numeric formatting | `%.2f`, `%.3f`, etc. |
| Building strings | `fmt.Sprintf` |
| Printing only | `fmt.Printf` |

---

📘 For a full list of formatting options, see the official Go [fmt](https://pkg.go.dev/fmt#hdr-Printing) package documentation.
