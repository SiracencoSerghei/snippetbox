# Early Returns & Guard Clauses in Go

Go supports **early returns**, which allow a function to return before reaching the end of its body. This feature is especially powerful when used to create **guard clauses**.

Guard clauses reduce nesting and make code easier to read by handling exceptional or terminating conditions early.

---

## Early Returns

An **early return** exits a function as soon as a condition is met.

```go
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return dividend / divisor, nil
}
```

### Why this is good

- Error case handled immediately
- Happy path is clear and uncluttered
- No unnecessary `else` blocks

This style is idiomatic Go, especially for error handling.

---

## Guard Clauses

A **guard clause** is a conditional check that exits the function early when a condition is met.

Instead of nesting logic deeply using `if / else`, guard clauses keep logic **one-dimensional**.

---

## Example: Nested Conditionals (Hard to Read)

```go
func getInsuranceAmount(status insuranceStatus) int {
	amount := 0
	if !status.hasInsurance() {
		amount = 1
	} else {
		if status.isTotaled() {
			amount = 10000
		} else {
			if status.isDented() {
				amount = 160
				if status.isBigDent() {
					amount = 270
				}
			} else {
				amount = 0
			}
		}
	}
	return amount
}
```

### Problems with this approach

- Deep nesting
- Hard to follow logic paths
- High cognitive load
- Difficult to see when a specific value is returned

---

## Same Logic Using Guard Clauses (Recommended)

```go
func getInsuranceAmount(status insuranceStatus) int {
	if !status.hasInsurance() {
		return 1
	}
	if status.isTotaled() {
		return 10000
	}
	if !status.isDented() {
		return 0
	}
	if status.isBigDent() {
		return 270
	}
	return 160
}
```

### Why this is better

- Flat structure (no deep nesting)
- Each rule is evaluated independently
- Easy to follow top-to-bottom
- Much easier to maintain

---

## Cognitive Load Explained

Cognitive load refers to how much mental effort is required to understand code.

### Nested version:

To understand when `270` is returned, you must:
- Track multiple `if / else` branches
- Remember which conditions apply
- Mentally reconstruct the logic tree

### Guard clause version:

To understand when `270` is returned, you simply:
- Read the function top to bottom
- Stop at the first matching condition

---

## Guard Clauses in Error Handling

Go error handling naturally encourages guard clauses:

```go
if err != nil {
	return err
}
```

This pattern keeps the happy path clean and readable.

---

## When to Use Guard Clauses

✅ Use guard clauses when:
- There are clear exit conditions
- You want to avoid deep nesting
- Logic can be evaluated sequentially

❌ Avoid when:
- All conditions must be evaluated
- Control flow is genuinely complex

---

## Key Takeaways

- Early returns simplify control flow
- Guard clauses reduce nesting
- One-dimensional code is easier to read
- Idiomatic Go favors clarity over cleverness

---

## Summary

| Technique | Benefit |
|--------|---------|
| Early return | Clear exit paths |
| Guard clauses | Reduced nesting |
| Flat structure | Lower cognitive load |

Write code for humans first — the compiler is already very smart.

