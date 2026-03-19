<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Advanced Go Closures</title>
<style>
body {
    font-family: Arial;
    background: #0f172a;
    color: #e2e8f0;
    margin: 40px;
}
h1, h2 { color: #38bdf8; }
pre {
    background: #1e293b;
    padding: 15px;
    border-radius: 8px;
}
.box {
    background: #1e293b;
    border-left: 4px solid #f87171;
    padding: 15px;
    margin: 20px 0;
}
.good {
    border-left: 4px solid #4ade80;
}
</style>
</head>
<body>

<h1>Advanced Closures in Go</h1>

<p>
Closures are not about syntax. They are about <strong>state + behavior</strong>.
Used correctly, they replace structs, encapsulate logic, and power real systems.
</p>

---

<h2>1. Stateful Counter (Function as Object)</h2>

<pre><code>func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}</code></pre>

<pre><code>c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
fmt.Println(c()) // 3
</code></pre>

<div class="good">
Each closure has its own memory. This is basically a lightweight object.
</div>

---

<h2>2. Closure as Factory</h2>

<pre><code>func multiplier(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}</code></pre>

<pre><code>double := multiplier(2)
triple := multiplier(3)

fmt.Println(double(5)) // 10
fmt.Println(triple(5)) // 15
</code></pre>

---

<h2>3. Middleware Pattern (Real Backend Use)</h2>

<pre><code>func logger(next func(string)) func(string) {
    return func(msg string) {
        fmt.Println("LOG:", msg)
        next(msg)
    }
}</code></pre>

<pre><code>func handler(msg string) {
    fmt.Println("Handling:", msg)
}

wrapped := logger(handler)
wrapped("Hello")
</code></pre>

<div class="good">
Closures allow wrapping behavior dynamically — core concept in web frameworks.
</div>

---

<h2>4. Lazy Evaluation / Caching</h2>

<pre><code>func lazySum(a, b int) func() int {
    var result *int

    return func() int {
        if result == nil {
            r := a + b
            result = &r
        }
        return *result
    }
}</code></pre>

---

<h2>5. Capturing Loop Variables (Classic Bug)</h2>

<div class="box">
This is where most developers write broken code.
</div>

❌ Wrong:

<pre><code>func bad() {
    for i := 0; i < 3; i++ {
        go func() {
            fmt.Println(i)
        }()
    }
}</code></pre>

👉 Output (unexpected):
333

---

✅ Correct:

<pre><code>func good() {
    for i := 0; i < 3; i++ {
        i := i // capture value
        go func() {
            fmt.Println(i)
        }()
    }
}</code></pre>

---

<h2>6. Closure vs Struct (When NOT to use)</h2>

<div class="box">
If your logic grows — closures become spaghetti.
</div>

❌ Bad:

<pre><code>func complex() func(int) int {
    state1 := 0
    state2 := 100
    state3 := "hello"
    // too much hidden state
    return func(x int) int {
        state1 += x
        return state1 + state2
    }
}</code></pre>

---

✅ Better:

<pre><code>type Processor struct {
    state1 int
    state2 int
}

func (p *Processor) Process(x int) int {
    p.state1 += x
    return p.state1 + p.state2
}
</code></pre>

---

<h2>7. When Closures Are Powerful</h2>

<ul>
<li>Function factories</li>
<li>Middleware</li>
<li>Encapsulation without structs</li>
<li>Short-lived state</li>
</ul>

---

<h2>8. When Closures Are Trash</h2>

<ul>
<li>Too much state</li>
<li>Hard to debug</li>
<li>Hidden side effects</li>
<li>Complex business logic</li>
</ul>

---

<h2>Final Insight</h2>

<div class="box">
Closures are not magic.  
They are just functions with memory.
</div>

<p>
Use them when they simplify the design.  
Avoid them when they hide complexity.
</p>

</body>
</html>