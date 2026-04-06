<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Go Structs — Deep Understanding</title>
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

<h1>Structs in Go — Real Understanding</h1>

<p>
Structs are used to group related data into a single unit.
Think of them as a custom data type.
</p>

<div class="box">
<strong>Mental model:</strong><br>
Struct = "your own type with fields"
</div>

---

<h2>Basic Definition</h2>

<pre><code>type Car struct {
    Brand   string
    Model   string
    Doors   int
    Mileage int
}</code></pre>

---

<h2>Creating a Struct</h2>

<pre><code>c := Car{
    Brand: "Toyota",
    Model: "Corolla",
    Doors: 4,
    Mileage: 120000,
}</code></pre>

---

<h2>Accessing Fields</h2>

<pre><code>fmt.Println(c.Brand)   // Toyota
c.Mileage += 100
</code></pre>

---

<h2>Why Structs Exist</h2>

<ul>
<li>Group related data</li>
<li>Replace maps/dictionaries</li>
<li>Improve readability</li>
</ul>

---

<h2>Struct vs Map</h2>

<div class="box">
Struct:
<ul>
<li>Typed</li>
<li>Fast</li>
<li>Safe</li>
</ul>

Map:
<ul>
<li>Dynamic</li>
<li>Flexible</li>
<li>Less safe</li>
</ul>
</div>

---

<h2>Zero Values</h2>

<pre><code>var c Car</code></pre>

<ul>
<li>string → ""</li>
<li>int → 0</li>
</ul>

---

<h2>Pointers to Struct</h2>

<pre><code>c := &Car{Brand: "BMW"}

c.Mileage = 1000</code></pre>

<div class="box">
Go automatically dereferences pointers for structs.
</div>

---

<h2>Methods on Struct</h2>

<pre><code>func (c Car) Drive(km int) {
    c.Mileage += km
}</code></pre>

<div class="warn">
This WON'T modify original struct (copy!)
</div>

---

<h2>Correct Version (Pointer Receiver)</h2>

<pre><code>func (c *Car) Drive(km int) {
    c.Mileage += km
}</code></pre>

<div class="good">
Now it modifies the original struct.
</div>

---

<h2>Struct Embedding (Composition)</h2>

<pre><code>type Engine struct {
    Power int
}

type Car struct {
    Brand string
    Engine
}</code></pre>

<pre><code>c.Engine.Power
c.Power // shorthand</code></pre>

---

<h2>Anonymous Structs</h2>

<pre><code>user := struct {
    Name string
    Age  int
}{
    Name: "John",
    Age: 30,
}</code></pre>

---

<h2>When Structs Become Trash</h2>

<div class="warn">
<ul>
<li>Too many fields</li>
<li>Mixed responsibilities</li>
<li>Used like a dump object</li>
</ul>
</div>

---

<h2>Best Practices</h2>

<div class="good">
<ul>
<li>Keep structs small</li>
<li>One responsibility</li>
<li>Use methods for behavior</li>
<li>Use pointers when modifying</li>
</ul>
</div>

---

<h2>Real Example</h2>

<pre><code>type User struct {
    Name string
    Balance float64
}

func (u *User) Deposit(amount float64) {
    u.Balance += amount
}</code></pre>

---

<h2>Summary</h2>

<ul>
<li>Struct = custom data type</li>
<li>Groups related fields</li>
<li>Used everywhere in Go</li>
<li>Pointer receivers are critical</li>
</ul>

</body>
</html>