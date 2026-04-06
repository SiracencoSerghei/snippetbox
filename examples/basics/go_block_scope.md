<hr>

<h1>Go Block Scope — Deep Dive</h1>

<p>
Go is <strong>block-scoped</strong>, not function-scoped like Python.
</p>

<p>
A variable exists only within the block <code>{}</code> where it is declared 
(and any nested blocks).
</p>

<hr>

<h2>Basic Example</h2>

<pre><code>package main

var age = 19 // package-level (global)

func sendEmail() {
    name := "Jon Snow" // function scope

    for i := 0; i < 5; i++ {
        email := "snow@winterfell.net" // loop scope
        _ = email
    }

    // email is NOT accessible here
}
</code></pre>

<hr>

<h2>What Creates a New Scope?</h2>

<ul>
  <li>Functions</li>
  <li>Loops (<code>for</code>)</li>
  <li>If statements</li>
  <li>Switch statements</li>
  <li>Select statements</li>
  <li>Explicit blocks <code>{}</code></li>
</ul>

<hr>

<h2>Explicit Block (Rare but Useful)</h2>

<pre><code>func main() {
    {
        age := 19
        fmt.Println(age) // OK
    }

    fmt.Println(age) // ❌ compile error
}
</code></pre>

<p>
Explicit blocks are used to:
</p>

<ul>
  <li>Limit variable lifetime</li>
  <li>Avoid name conflicts</li>
  <li>Control memory usage</li>
</ul>

<hr>

<h2>Shadowing (Critical Trap)</h2>

<p>
Go allows redeclaring variables in inner scopes. This is called <strong>shadowing</strong>.
</p>

<pre><code>func main() {
    x := 10

    if true {
        x := 20 // NEW variable, shadows outer x
        fmt.Println(x) // 20
    }

    fmt.Println(x) // 10
}
</code></pre>

<p>
🚨 This is a common source of bugs.
</p>

<hr>

<h2>Shadowing with Errors (Real Bug)</h2>

<pre><code>func example() error {
    var err error

    if true {
        err := doSomething() // ❌ new err, NOT the outer one
        if err != nil {
            return err
        }
    }

    return err // always nil → BUG
}
</code></pre>

<p>
You think you're updating <code>err</code>, but you're not.
</p>

<hr>

<h2>Fix Shadowing Bug</h2>

<pre><code>func example() error {
    var err error

    if true {
        err = doSomething() // correct: assign, don't redeclare
        if err != nil {
            return err
        }
    }

    return err
}
</code></pre>

<hr>

<h2>Short Variable Declaration Rule (<code>:=</code>)</h2>

<p>
<code>:=</code> creates a new variable if at least one variable is new.
</p>

<pre><code>func main() {
    x := 10

    x, y := 20, 30 // x reassigned, y is new
    fmt.Println(x, y)
}
</code></pre>

<p>
🚨 Easy to accidentally introduce shadowing.
</p>

<hr>

<h2>If Statement Scope</h2>

<pre><code>if x := 10; x > 5 {
    fmt.Println(x) // OK
}

fmt.Println(x) // ❌ undefined
</code></pre>

<p>
Variables declared in <code>if</code> are scoped to that statement only.
</p>

<hr>

<h2>Loop Scope + Closure Trap</h2>

<pre><code>for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)
    }()
}
</code></pre>

<p>
🚨 Output is unpredictable (often 3,3,3)
</p>

<p>
Because all goroutines capture the same variable.
</p>

<h3>Fix</h3>

<pre><code>for i := 0; i < 3; i++ {
    i := i
    go func() {
        fmt.Println(i)
    }()
}
</code></pre>

<hr>

<h2>Memory & Lifetime Insight</h2>

<p>
Smaller scope = shorter lifetime.
</p>

<ul>
  <li>Helps garbage collector</li>
  <li>Reduces memory pressure</li>
  <li>Makes code safer</li>
</ul>

<hr>

<h2>Best Practices</h2>

<ul>
  <li>Avoid unnecessary shadowing</li>
  <li>Use <code>=</code> when reassigning</li>
  <li>Keep variables in the smallest possible scope</li>
  <li>Be careful with <code>:=</code> in complex blocks</li>
</ul>

<hr>

<h2>Interview Traps</h2>

<h3>1. What does this print?</h3>

<pre><code>x := 10

if true {
    x := 20
}

fmt.Println(x)
</code></pre>

<p><strong>Answer:</strong> 10</p>

<hr>

<h3>2. Is this a bug?</h3>

<pre><code>if err := doSomething(); err != nil {
    return err
}
</code></pre>

<p>
✔️ No — this is idiomatic Go (scoped error handling)
</p>

<hr>

<h2>Bottom Line</h2>

<p>
If you ignore scope rules, you will:
</p>

<ul>
  <li>Shadow variables accidentally</li>
  <li>Lose errors silently</li>
  <li>Create race conditions</li>
</ul>

<p>
If you understand it — your code becomes predictable and safe.
</p>