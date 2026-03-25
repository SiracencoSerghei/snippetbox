<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Go Testing Guide</title>
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

<h1>🧪 Testing Guide (Go)</h1>

<p>
This is how you actually control tests in Go.
</p>

---

<h2>Run All Tests</h2>

<pre><code>go test ./... -v</code></pre>

<div class="box">
Runs all packages recursively.
</div>

---

<h2>Run Specific Package</h2>

<pre><code>go test ./examples -v</code></pre>

---

<h2>Run Specific Test</h2>

<pre><code>go test ./examples -run TestCarDrive -v</code></pre>

<div class="good">
-run uses regex
</div>

---

<h2>Run Multiple Tests (Regex)</h2>

<pre><code>go test ./examples -run 'TestCar|TestMap' -v</code></pre>

---

<h2>Disable Cache (CRITICAL)</h2>

<pre><code>go test ./examples -v -count=1</code></pre>

<div class="warn">
Without this, Go may NOT rerun tests!
</div>

---

<h2>Run Subtests</h2>

<pre><code>go test ./examples -run 'TestSplitEmail/valid_simple' -v</code></pre>

---

<h2>Show Coverage</h2>

<pre><code>go test ./examples -cover</code></pre>

---

<h2>Detailed Coverage</h2>

<pre><code>go test ./examples -coverprofile=cover.out
go tool cover -html=cover.out</code></pre>

<div class="good">
Opens visual coverage report
</div>

---

<h2>Benchmark Tests</h2>

<pre><code>go test ./examples -bench .</code></pre>

---

<h2>Benchmark with Memory</h2>

<pre><code>go test -bench . -benchmem</code></pre>

---

<h2>Run Only Benchmarks</h2>

<pre><code>go test -run=^$ -bench .</code></pre>

---

<h2>Race Condition Detection</h2>

<pre><code>go test -race ./...</code></pre>

<div class="warn">
Must use in concurrent code
</div>

---

<h2>Verbose Logging</h2>

<pre><code>go test -v</code></pre>

---

<h2>Fail Fast</h2>

<pre><code>go test -failfast</code></pre>

---

<h2>Timeout Control</h2>

<pre><code>go test -timeout 5s</code></pre>

---

<h2>Debug with Print</h2>

<pre><code>t.Log("debug info")
t.Logf("value: %d", x)</code></pre>

---

<h2>Golden Rule</h2>

<div class="box">
Always use:
<pre><code>go test -v -count=1</code></pre>
</div>

---

<h2>Common Mistakes</h2>

<div class="warn">
<ul>
<li>Running single file instead of package</li>
<li>Forgetting -count=1</li>
<li>Wrong test names (must start with Test)</li>
<li>Package mismatch</li>
</ul>
</div>

---

<h2>Pro Workflow</h2>

<pre><code># run only one test
go test ./examples -run TestCarDrive -v -count=1

# debug coverage
go test ./examples -coverprofile=cover.out
go tool cover -html=cover.out

# performance
go test -bench . -benchmem

# concurrency safety
go test -race ./...</code></pre>

---

<h2>Final Insight</h2>

<div class="good">
Go testing is not about writing tests.<br>
It's about controlling execution precisely.
</div>

</body>
</html>