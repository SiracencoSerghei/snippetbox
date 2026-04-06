<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Currying in Go — Real Understanding</title>
<style>
body {
    font-family: Arial;
    background: #0f172a;
    color: #e2e8f0;
    margin: 40px;
    line-height: 1.6;
}
h1, h2 {
    color: #38bdf8;
}
pre {
    background: #1e293b;
    padding: 15px;
    border-radius: 8px;
}
.box {
    background: #1e293b;
    border-left: 4px solid #38bdf8;
    padding: 15px;
    margin: 20px 0;
}
.warn {
    border-left: 4px solid #f87171;
}
.good {
    border-left: 4px solid #4ade80;
}
</style>
</head>
<body>

<h1>Currying in Go — Real Understanding</h1>

<p>
Currying is not magic. It is just a way to <strong>break a function with multiple arguments into steps</strong>.
</p>

<div class="box">
<strong>Simple idea:</strong><br>
Instead of:
<pre><code>f(x, y)</code></pre>
You do:
<pre><code>f(x)(y)</code></pre>
</div>

---

<h2>Step 1 — Normal Function</h2>

<pre><code>func multiply(x, y int) int {
    return x * y
}</code></pre>

<p>
You must pass both arguments at once.
</p>

---

<h2>Step 2 — Currying Version</h2>

<pre><code>func multiply(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}</code></pre>

<div class="good">
Now the function returns another function.<br>
The first argument (<code>x</code>) is stored inside a closure.
</div>

---

<h2>Usage</h2>

<pre><code>double := multiply(2)
result := double(5)

fmt.Println(result) // 10</code></pre>

---

<h2>What Actually Happens</h2>

<div class="box">
Step-by-step:
<ul>
<li>multiply(2) → returns function</li>
<li>that function "remembers" x = 2</li>
<li>then you call it with 5 → 2 * 5</li>
</ul>
</div>

---

<h2>Your Example Explained</h2>

<pre><code>func selfMath(mathFunc func(int, int) int) func(int) int {
    return func(x int) int {
        return mathFunc(x, x)
    }
}</code></pre>

---

<h2>How It Works</h2>

<ul>
<li>You pass a function: multiply or add</li>
<li>selfMath returns a new function</li>
<li>This function calls original with (x, x)</li>
</ul>

---

<h2>Execution Flow</h2>

<pre><code>squareFunc := selfMath(multiply)

squareFunc(5)</code></pre>

<div class="box">
Flow:
<ul>
<li>selfMath(multiply)</li>
<li>returns func(x)</li>
<li>call func(5)</li>
<li>multiply(5, 5)</li>
<li>= 25</li>
</ul>
</div>

---

<h2>Key Insight</h2>

<div class="good">
Currying = storing part of the arguments for later.
</div>

---

<h2>Another Example (Real Currying)</h2>

<pre><code>func add(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}</code></pre>

<pre><code>add10 := add(10)
fmt.Println(add10(5)) // 15</code></pre>

---

<h2>Why This Exists</h2>

<ul>
<li>Create reusable logic</li>
<li>Pre-configure functions</li>
<li>Avoid repeating arguments</li>
</ul>

---

<h2>When This Becomes Trash</h2>

<div class="box warn">
<ul>
<li>Too many nested functions</li>
<li>Hard to read</li>
<li>No real benefit</li>
</ul>
</div>

---

<h2>Mental Model</h2>

<div class="box">
Think like this:
<pre><code>multiply(2) → "function that always multiplies by 2"</code></pre>
</div>

---

<h2>Final Summary</h2>

<ul>
<li>Currying = function returning function</li>
<li>Uses closures to store arguments</li>
<li>Breaks execution into steps</li>
<li>Powerful but easy to overuse</li>
</ul>

</body>
</html>