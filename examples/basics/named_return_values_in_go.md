# Named Return Values (Go)

Go functions can declare **named return values**. When return values are named, they behave like variables that are **implicitly declared and initialized** at the start of the function.

Named return values are most useful as **documentation**—they make it clear what each returned value represents.

---

## Basic Idea

```go
func getCoords() (x, y int) {
	// x and y are created here and initialized to zero values
	return // naked return
}
```

- `x` and `y` are **return variables**
- They start with their **zero values** (`0` for `int`)
- `return` without arguments returns the current values of `x` and `y`

This is called a **naked return**.

---

## Equivalent Version (Without Named Returns)

```go
func getCoords() (int, int) {
	var x int
	var y int
	return x, y
}
```

Both functions behave the same way.

The difference is **readability and intent**, not behavior.

---

## Naked Returns

From *A Tour of Go*:

> A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.

### Good use (short & obvious):

```go
func sum(a, b int) (result int) {
	result = a + b
	return
}
```

### Bad use (long or complex logic):

```go
func complicated() (ok bool) {
	// many lines of logic...
	return // hard to see what is being returned
}
```

---

## When to Use Named Return Values

✅ Use them when:
- The function is **short**
- The meaning of return values is **obvious**
- You want to **document intent**

❌ Avoid them when:
- The function is **long or complex**
- It’s unclear what variables are returned
- Naked returns reduce readability

---

## Named Returns Without Naked Return

You can still explicitly return values:

```go
func divide(a, b int) (quotient int, ok bool) {
	if b == 0 {
		return 0, false
	}
	quotient = a / b
	return quotient, true
}
```

Named return values do **not** force naked returns.

---

## Key Takeaways

- Named return values are **real variables**
- They are initialized to **zero values**
- `return` alone returns all named values
- Best used as **documentation**, not a shortcut
- Prefer clarity over cleverness

---

## Summary

| Feature | Named Returns |
|------|--------------|
| Improves readability | ✅ |
| Required in Go | ❌ |
| Allows naked return | ✅ |
| Good for long functions | ❌ |

Use them when they make the code easier to understand—otherwise, return values explicitly.

