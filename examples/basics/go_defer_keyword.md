<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Go defer — Complete Guide</title>
</head>
<body>

<h1>Go <code>defer</code> — Complete Guide</h1>

<p>
The <code>defer</code> keyword schedules a function call to run 
<strong>right before the surrounding function returns</strong>.
</p>

<p>
It is mainly used for cleanup: closing files, DB connections, releasing resources.
</p>

<hr>

<h2>Core Rules</h2>

<ul>
  <li>Arguments are evaluated immediately</li>
  <li>Execution happens at function return</li>
  <li>Multiple defers run in LIFO order (Last-In, First-Out)</li>
</ul>

<hr>

<h2>Basic Example</h2>

<pre><code>func GetUsername(dstName, srcName string) (username string, err error) {
    conn, _ := db.Open(srcName)

    defer conn.Close()

    username, err = db.FetchUser()
    if err != nil {
        return "", err
    }

    return username, nil
}
</code></pre>

<p>
<code>conn.Close()</code> will execute no matter where the function returns.
</p>

<hr>

<h2>Multiple Defers (LIFO)</h2>

<pre><code>func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
</code></pre>

<p><strong>Output:</strong></p>

<pre><code>3
2
1
</code></pre>

<hr>

<h2>Order Matters</h2>

<pre><code>func CreateTempFile() {
    f, _ := os.Create("temp.txt")

    defer os.Remove(f.Name()) // runs second
    defer f.Close()           // runs first
}
</code></pre>

<hr>

<h2>Argument Evaluation Trap</h2>

<pre><code>func main() {
    x := 10
    defer fmt.Println(x)
    x = 20
}
</code></pre>

<p><strong>Output:</strong> 10</p>

<p>
Arguments are evaluated when <code>defer</code> is declared.
</p>

<hr>

<h2>Closure Behavior</h2>

<pre><code>func main() {
    x := 10
    defer func() {
        fmt.Println(x)
    }()
    x = 20
}
</code></pre>

<p><strong>Output:</strong> 20</p>

<p>
Closures capture variables, not values.
</p>

<hr>

<h2>Named Return Modification</h2>

<pre><code>func test() (result int) {
    defer func() {
        result++
    }()
    return 5
}
</code></pre>

<p><strong>Output:</strong> 6</p>

<p>
Deferred function runs after return value is assigned.
</p>

<hr>

<h2>Defer in Loops (Common Bug)</h2>

<pre><code>for i := 0; i < 3; i++ {
    defer fmt.Println(i)
}
</code></pre>

<p><strong>Output:</strong></p>

<pre><code>2
1
0
</code></pre>

<p>
All defers execute after the loop finishes.
</p>

<h3>Real Problem (Resource Leak)</h3>

<pre><code>for _, f := range files {
    file, _ := os.Open(f)
    defer file.Close()
}
</code></pre>

<p>
🚨 Files stay open until function exits → can crash system.
</p>

<h3>Fix</h3>

<pre><code>for _, f := range files {
    func() {
        file, _ := os.Open(f)
        defer file.Close()
    }()
}
</code></pre>

<hr>

<h2>Loop + Closure Trap (Interview Classic)</h2>

<pre><code>for i := 0; i < 3; i++ {
    defer func() {
        fmt.Println(i)
    }()
}
</code></pre>

<p><strong>Output:</strong></p>

<pre><code>3
3
3
</code></pre>

<p>
Because closure captures <code>i</code>.
</p>

<h3>Fix</h3>

<pre><code>for i := 0; i < 3; i++ {
    i := i
    defer func() {
        fmt.Println(i)
    }()
}
</code></pre>

<hr>

<h2>Panic + Defer</h2>

<pre><code>func main() {
    defer fmt.Println("cleanup")
    panic("boom")
}
</code></pre>

<p><strong>Output:</strong></p>

<pre><code>cleanup
panic: boom
</code></pre>

<p>
Defers execute even during panic.
</p>

<hr>

<h2>Recover (Advanced)</h2>

<pre><code>func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()

    panic("boom")
}
</code></pre>

<p>
<code>recover()</code> only works inside deferred functions.
</p>

<hr>

<h2>Performance Consideration</h2>

<ul>
  <li><code>defer</code> adds overhead (~20–80ns per call)</li>
  <li>Fine for normal code</li>
  <li>Avoid in hot loops or critical paths</li>
</ul>

<hr>

<h2>Key Takeaways</h2>

<ul>
  <li>Runs at function exit, not where written</li>
  <li>Arguments evaluated immediately</li>
  <li>LIFO execution order</li>
  <li>Closures capture variables</li>
  <li>Named returns can be modified</li>
  <li>Dangerous inside loops</li>
  <li>Has performance cost</li>
</ul>

<hr>

<h2>Bottom Line</h2>

<p>
If you misunderstand <code>defer</code>, you will:
</p>

<ul>
  <li>Leak resources</li>
  <li>Get wrong values</li>
  <li>Introduce hidden bugs</li>
</ul>

<p>
If you understand it — your Go code becomes predictable and safe.
</p>

</body>
</html>