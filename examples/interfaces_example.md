<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Go Interfaces — From Zero to Real Understanding</title>
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
    overflow-x: auto;
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

<h1>Go Interfaces — Senior Mindset, Simple Explanation</h1>

<p>
If struct is "data", then interface is "behavior".
</p>

<div class="box">
<strong>Core idea:</strong><br>
Interface describes WHAT something can do, not WHAT it is.
</div>

---

<h2>1. Basic Interface</h2>

<pre><code>type Speaker interface {
    Speak() string
}</code></pre>

<p>
This means: anything that has Speak() string → is a Speaker.
</p>

---

<h2>2. No "implements" keyword</h2>

<pre><code>type Dog struct{}

func (d Dog) Speak() string {
    return "woof"
}</code></pre>

<div class="good">
Dog automatically implements Speaker.
</div>

<div class="warn">
If you think "where is implements?" → wrong mindset.
</div>

---

<h2>3. Using Interface</h2>

<pre><code>func makeSound(s Speaker) {
    fmt.Println(s.Speak())
}</code></pre>

<pre><code>dog := Dog{}
makeSound(dog)</code></pre>

---

<h2>4. Why Interfaces Exist</h2>

<div class="box">
Without interface:
<ul>
<li>tight coupling</li>
<li>hard to test</li>
</ul>

With interface:
<ul>
<li>flexibility</li>
<li>mocking</li>
<li>clean architecture</li>
</ul>
</div>

---

<h2>5. Real Example (Dependency Injection)</h2>

<pre><code>type Mailer interface {
    Send(message string) error
}</code></pre>

<pre><code>type EmailService struct{}

func (e EmailService) Send(message string) error {
    fmt.Println("Sending:", message)
    return nil
}</code></pre>

<pre><code>func Notify(m Mailer) {
    m.Send("Hello")
}</code></pre>

---

<h2>6. Why This is Powerful</h2>

<div class="good">
You can swap implementation without changing code.
</div>

---

<h2>7. Empty Interface (⚠️ Dangerous)</h2>

<pre><code>var x interface{}</code></pre>

<div class="warn">
This is basically "any" — use carefully.
</div>

---

<h2>8. Type Assertion</h2>

<pre><code>val, ok := x.(string)</code></pre>

---

<h2>9. Interface + Struct = Real Go</h2>

<div class="box">
Struct = data<br>
Interface = behavior<br>
Together = architecture
</div>

---

<h2>10. Common Mistakes</h2>

<div class="warn">
<ul>
<li>Creating interface too early</li>
<li>Making huge interfaces</li>
<li>Using interface{} everywhere</li>
<li>Thinking in OOP (Java style)</li>
</ul>
</div>

---

<h2>11. Senior Rules</h2>

<div class="good">
<ul>
<li>Interfaces should be small (1-2 methods)</li>
<li>Define interfaces where they are used</li>
<li>Prefer composition over inheritance</li>
<li>Use interfaces for testing</li>
</ul>
</div>

---

<h2>12. Mental Model</h2>

<pre><code>
struct = WHAT it is
interface = WHAT it can do
</code></pre>

---

<h2>Final Thought</h2>

<div class="box">
If you don’t understand interfaces → you don’t understand Go.
</div>

</body>
</html>