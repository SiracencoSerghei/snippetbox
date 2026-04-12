<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Go Slices — Combined Deep Dive</title>
  <style>
    body { font-family: Arial, sans-serif; line-height: 1.7; max-width: 1100px; margin: 40px auto; padding: 0 20px; color: #222; }
    h1, h2, h3 { color: #111; }
    pre { background: #f4f4f4; padding: 16px; border-radius: 10px; overflow-x: auto; }
    code { font-family: Consolas, monospace; }
    .card { border: 1px solid #ddd; border-radius: 16px; padding: 20px; margin: 20px 0; }
    .warn { background: #fff8e1; }
    .good { background: #eefbf0; }
    .deep { background: #eef4ff; }
    table { width: 100%; border-collapse: collapse; margin: 16px 0; }
    th, td { border: 1px solid #ddd; padding: 10px; text-align: left; }
    th { background: #f7f7f7; }
  </style>
</head>
<body>
  <h1>Go Slices — Combined Production + Memory Deep Dive</h1>
  <p>This is the <strong>combined version</strong>: slice internals, <code>var</code> vs <code>make</code>, growth, reallocations, ownership, and production memory behavior.</p>

  <div class="card deep">
    <h2>1) Slice Header Mental Model</h2>
    <pre><code>type slice struct {
    ptr *T
    len int
    cap int
}</code></pre>
    <p>The slice is only a descriptor. The real data lives in the backing array.</p>
  </div>

  <div class="card">
    <h2>2) var vs Literal vs make</h2>
    <pre><code>var a []int          // nil
b := []int{}        // empty non-nil
c := make([]int, 0) // empty non-nil
d := make([]int, 0, 100)</code></pre>
    <table>
      <tr><th>Form</th><th>len</th><th>cap</th><th>nil</th><th>Best use</th></tr>
      <tr><td>var a []int</td><td>0</td><td>0</td><td>true</td><td>zero-value APIs</td></tr>
      <tr><td>[]int{}</td><td>0</td><td>0</td><td>false</td><td>JSON []</td></tr>
      <tr><td>make([]int,0)</td><td>0</td><td>0</td><td>false</td><td>explicit init</td></tr>
      <tr><td>make([]int,0,100)</td><td>0</td><td>100</td><td>false</td><td>performance</td></tr>
    </table>
  </div>

  <div class="card good">
    <h2>3) len, cap, append Growth</h2>
    <pre><code>s := make([]int, 0, 2)
s = append(s, 10)
s = append(s, 20)
s = append(s, 30) // may trigger reallocation</code></pre>
    <p>When <code>len == cap</code>, Go allocates a new bigger array and copies data.</p>
  </div>

  <div class="card warn">
    <h2>4) Reallocation Cost in Production</h2>
    <ul>
      <li>new heap allocation</li>
      <li>copy old elements</li>
      <li>more GC pressure</li>
      <li>pointer may change</li>
    </ul>
    <pre><code>users := make([]User, 0, expectedRows)</code></pre>
    <p>Preallocation removes avoidable reallocations in hot paths.</p>
  </div>

  <div class="card warn">
    <h2>5) Memory Ownership</h2>
    <pre><code>base := []int{1,2,3,4}
a := base[:2]
b := base[1:3]

b[0] = 99 // mutates shared array</code></pre>
    <p>Two slices can own the same memory view. This is the source of many hidden bugs.</p>
    <pre><code>clone := append([]int(nil), base...)</code></pre>
  </div>

  <div class="card warn">
    <h2>6) Large Array Retention Bug</h2>
    <pre><code>buf := make([]byte, 10_000_000)
small := buf[:10]</code></pre>
    <p><code>small</code> still retains the entire 10MB backing array.</p>
    <pre><code>small = append([]byte(nil), small...)</code></pre>
  </div>

  <div class="card good">
    <h2>7) Safe Goroutine Ownership</h2>
    <pre><code>payloadCopy := append([]byte(nil), payload...)
go worker(payloadCopy)</code></pre>
    <p>Never pass shared mutable slice memory into async code without copying.</p>
  </div>

  <div class="card deep">
    <h2>8) Production Rules</h2>
    <ul>
      <li>Prefer <code>var s []T</code> as zero value</li>
      <li>Use <code>make(..., cap)</code> when size is predictable</li>
      <li>Always reassign append result</li>
      <li>Copy before sharing across goroutines</li>
      <li>Never keep tiny slices from giant buffers</li>
    </ul>
  </div>
  <div class="card good">
    <h2>9) Core Slice Operations You Use Every Day</h2>
    <h3>Append</h3>
    <pre><code>s = append(s, 40)
s = append(s, 50, 60)</code></pre>
    <p>Always capture the returned slice because reallocation may move the backing array.</p>

    <h3>Copy</h3>
    <pre><code>dst := make([]int, len(s))
copy(dst, s)</code></pre>
    <p><code>copy()</code> duplicates elements into a different backing array, which is critical for ownership safety.</p>

    <h3>Insert at Index</h3>
    <pre><code>i := 2
v := 99
s = append(s[:i], append([]int{v}, s[i:]...)...)</code></pre>
    <p>Useful for ordered lists, middleware chains, priority queues, and in-memory indexes.</p>

    <h3>Delete by Index</h3>
    <pre><code>i := 2
s = append(s[:i], s[i+1:]...)</code></pre>
    <p>This preserves order. Common in caches, task lists, and request pipelines.</p>

    <h3>Delete Without Preserving Order</h3>
    <pre><code>i := 2
s[i] = s[len(s)-1]
s = s[:len(s)-1]</code></pre>
    <p>O(1) deletion for high-performance systems where order does not matter.</p>
  </div>

  <div class="card deep">
    <h2>10) Final Senior-Level Mental Model</h2>
    <p>A slice is a <strong>window + ownership contract</strong> over an array.</p>
    <p>Every production decision with slices is about:</p>
    <ul>
      <li>who owns the backing memory</li>
      <li>who may mutate it</li>
      <li>whether growth reallocates</li>
      <li>whether retained capacity wastes RAM</li>
    </ul>
    <p>Master these, and slices stop being syntax and become a performance tool.</p>
  </div>
</body>
</html>
